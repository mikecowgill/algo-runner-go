package main

import (
	"algo-runner-go/swagger"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

func main() {

	configFilePtr := flag.String("config", "./config.json", "JSON config file to load")
	kafkaServersPtr := flag.String("kafka-servers", "localhost:9092", "Kafka broker addresses separated by a comma")

	flag.Parse()

	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	config := loadConfig(*configFilePtr)

	// Launch the server if not started
	if config.Serverless == false {

		var serverTerminated bool
		go func() {
			serverTerminated = startServer(config, kafkaServersPtr)
			if serverTerminated {
				os.Exit(1)
			}
		}()

	}

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":               *kafkaServersPtr,
		"group.id":                        "myGroup",
		"auto.offset.reset":               "earliest",
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Consumer %v\n", c)

	topicRoutes := make(map[string]swagger.PipelineRouteModel)

	for _, route := range config.PipelineRoutes {

		if route.DestAlgoOwnerName == config.AlgoOwnerUserName &&
			route.DestAlgoUrlName == config.AlgoUrlName {

			switch routeType := route.RouteType; routeType {
			case "Algo":

				topic := strings.ToLower(fmt.Sprintf("algorun.%s.%s.algo.%s.%s",
					config.EndpointOwnerUserName,
					config.EndpointUrlName,
					route.SourceAlgoOwnerName,
					route.SourceAlgoUrlName))

				topicRoutes[topic] = route

				err = c.Subscribe(topic, nil)

			case "DataSource":

				topic := strings.ToLower(fmt.Sprintf("algorun.%s.%s.connector.%s",
					config.EndpointOwnerUserName,
					config.EndpointUrlName,
					route.PipelineDataSource.DataConnector.Name))

				topicRoutes[topic] = route

				err = c.Subscribe(topic, nil)

			case "EndpointSource":

				topic := strings.ToLower(fmt.Sprintf("algorun.%s.%s.output.%s",
					config.EndpointOwnerUserName,
					config.EndpointUrlName,
					route.PipelineEndpointSourceOutputName))

				topicRoutes[topic] = route

				err = c.Subscribe(topic, nil)

			}

		}

	}

	data := make(map[string]map[*swagger.AlgoInputModel][]InputData)

	waiting := true

	for waiting == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			waiting = false

		case ev := <-c.Events():
			switch e := ev.(type) {
			case *kafka.Message:

				fmt.Printf("%% Message on %s:\n%s\n",
					e.TopicPartition, string(e.Value))

				route := topicRoutes[*e.TopicPartition.Topic]
				runID, inputData, run := processMessage(e, route, config)

				inputMap := make(map[*swagger.AlgoInputModel][]InputData)

				inputMap[route.DestAlgoInput] = append(inputMap[route.DestAlgoInput], inputData)

				data[runID] = inputMap

				if run {

					if config.Serverless {
						runExec(config, data[runID])
					} else {
						runHTTP(config, data[runID])
					}

					delete(data, runID)
				}

			case kafka.AssignedPartitions:
				fmt.Fprintf(os.Stderr, "%% %v\n", e)
				c.Assign(e.Partitions)
			case kafka.RevokedPartitions:
				fmt.Fprintf(os.Stderr, "%% %v\n", e)
				c.Unassign()
			case kafka.PartitionEOF:
				fmt.Printf("%% Reached %v\n", e)
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
				waiting = false
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	c.Close()

}

func processMessage(msg *kafka.Message,
	route swagger.PipelineRouteModel,
	config swagger.RunnerConfig) (runID string, inputData InputData, run bool) {

	// Parse the headers
	var fileName string
	for _, header := range msg.Headers {
		if header.Key == "runId" {
			runID = string(header.Value)
		}
		if header.Key == "fileName" {
			fileName = string(header.Value)
		}
		if header.Key == "run" {
			b, _ := strconv.ParseBool(string(header.Value))
			run = b
		}
	}

	// Save the data based on the delivery type
	inputData = InputData{}

	if route.DestAlgoInput.InputDeliveryType == "StdIn" ||
		route.DestAlgoInput.InputDeliveryType == "Http" ||
		route.DestAlgoInput.InputDeliveryType == "Https" {

		inputData.isFile = false
		inputData.data = msg.Value

	} else {

		inputData.isFile = true
		if _, err := os.Stat("/data"); os.IsNotExist(err) {
			os.MkdirAll("/data", os.ModePerm)
		}
		file := fmt.Sprintf("/data/%s", fileName)
		err := ioutil.WriteFile(file, msg.Value, 0644)
		if err != nil {
			// TODO: Log error
		}

		inputData.data = []byte(file)

	}

	return

}

func getCommand(config swagger.RunnerConfig) []string {

	cmd := strings.Split(config.Entrypoint, " ")

	for _, param := range config.AlgoParams {
		cmd = append(cmd, param.Name)
		if param.DataType.Name != "switch" {
			cmd = append(cmd, param.Value)
		}
	}

	return cmd
}

func getEnvironment(config swagger.RunnerConfig) []string {

	env := strings.Split(config.Entrypoint, " ")

	return env
}

func runExec(config swagger.RunnerConfig,
	inputMap map[*swagger.AlgoInputModel][]InputData) {

	startTime := time.Now()

	command := getCommand(config)

	targetCmd := exec.Command(command[0], command[1:]...)

	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	// setup termination on kill signals
	go func() {
		sig := <-sigchan
		fmt.Printf("Caught signal %v. Killing server process: %s\n", sig, config.Entrypoint)
		if targetCmd != nil && targetCmd.Process != nil {
			val := targetCmd.Process.Kill()
			if val != nil {
				fmt.Printf("Killed algo process: %s - error %s\n", config.Entrypoint, val.Error())
			}
		}
	}()

	envs := getEnvironment(config)
	if len(envs) > 0 {
		//targetCmd.Env = envs
	}

	var out []byte
	var cmdErr error

	var wg sync.WaitGroup

	// TODO: Write to the topic as error if no value
	if inputMap == nil {

		// 	return
	}

	var timer *time.Timer

	// Set the timeout routine
	if config.TimeoutSeconds > 0 {
		timer = time.NewTimer(time.Duration(config.TimeoutSeconds) * time.Second)

		go func() {
			<-timer.C
			fmt.Printf("Killing process: %s\n", config.Entrypoint)
			if targetCmd != nil && targetCmd.Process != nil {
				val := targetCmd.Process.Kill()
				if val != nil {
					fmt.Printf("Killed process: %s - error %s\n", config.Entrypoint, val.Error())
				}
			}
		}()
	}

	for input, inputData := range inputMap {

		switch inputDeliveryType := input.InputDeliveryType; inputDeliveryType {
		case "StdIn":
			for _, data := range inputData {

				// get the writer for stdin
				writer, _ := targetCmd.StdinPipe()

				wg.Add(1)

				// Write to pipe in separate go-routine to prevent blocking
				go func(stdInData []byte) {
					defer wg.Done()

					writer.Write(stdInData)
					writer.Close()
				}(data.data)
			}

		case "Parameter":

			if input.Parameter != "" {
				targetCmd.Args = append(targetCmd.Args, input.Parameter)
			}
			for _, data := range inputData {
				targetCmd.Args = append(targetCmd.Args, string(data.data))
			}

		case "RepeatedParameter":

			for _, data := range inputData {
				if input.Parameter != "" {
					targetCmd.Args = append(targetCmd.Args, input.Parameter)
				}
				targetCmd.Args = append(targetCmd.Args, string(data.data))
			}

		case "DelimitedParameter":

			if input.Parameter != "" {
				targetCmd.Args = append(targetCmd.Args, input.Parameter)
			}
			var buffer bytes.Buffer
			for i := 0; i < len(inputData); i++ {
				buffer.WriteString(string(inputData[i].data))
				if i != len(inputData)-1 {
					buffer.WriteString(input.ParameterDelimiter)
				}
			}
			targetCmd.Args = append(targetCmd.Args, buffer.String())
		}
	}

	wg.Add(1)

	go func() {
		var b bytes.Buffer
		targetCmd.Stderr = &b

		defer wg.Done()

		out, cmdErr = targetCmd.Output()
		if b.Len() > 0 {
			fmt.Printf("stderr: %s", b.Bytes())
		}
		b.Reset()
	}()

	wg.Wait()

	if timer != nil {
		timer.Stop()
	}

	if cmdErr != nil {

		fmt.Printf("Success=%t, Error=%s\n", targetCmd.ProcessState.Success(), cmdErr.Error())
		fmt.Printf("Out=%s\n", out)

		// TODO: Write error to output topic

		return
	}

	var bytesWritten string
	// if config.writeDebug == true {
	os.Stdout.Write(out)
	// } else {
	bytesWritten = fmt.Sprintf("Wrote %d Bytes", len(out))
	//}

	execDuration := time.Since(startTime).Seconds()

	// TODO: Write to output topic

	if len(bytesWritten) > 0 {
		fmt.Printf("%s - Duration: %f seconds", bytesWritten, execDuration)
	} else {
		fmt.Printf("Duration: %f seconds", execDuration)
	}

}

func runHTTP(config swagger.RunnerConfig,
	inputMap map[*swagger.AlgoInputModel][]InputData) {

	startTime := time.Now()

	// TODO: Write to the topic as error if no value
	if inputMap == nil {

		// 	return
	}

	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	// Set the timeout routine
	if config.TimeoutSeconds > 0 {
		netClient.Timeout = time.Second * time.Duration(config.TimeoutSeconds)
	}

	for input, inputData := range inputMap {

		u, _ := url.Parse("http://localhost")

		u.Scheme = strings.ToLower(input.InputDeliveryType)
		if input.HttpPort > 0 {
			u.Host = fmt.Sprintf("http://localhost:%d", input.HttpPort)
		}
		u.Path = input.HttpPath

		q := u.Query()
		for _, param := range config.AlgoParams {
			q.Set(param.Name, param.Value)
		}
		u.RawQuery = q.Encode()

		for _, data := range inputData {
			request, err := http.NewRequest(strings.ToLower(input.HttpVerb), u.String(), bytes.NewReader(data.data))
			if err != nil {

			}
			response, err := netClient.Do(request)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error getting response from http server: %s\n", err)
			} else {
				defer response.Body.Close()
				contents, err := ioutil.ReadAll(response.Body)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error reading response from http server: %s\n", err)
				}
				fmt.Println("The calculated length is:", len(string(contents)), "for the url:", u.String())
				fmt.Println("   ", response.StatusCode)
				hdr := response.Header
				for key, value := range hdr {
					fmt.Println("   ", key, ":", value)
				}
				fmt.Println(contents)
			}
		}
		fmt.Println(u)

	}

	execDuration := time.Since(startTime).Seconds()

	// TODO: Write to output topic

	//if len(bytesWritten) > 0 {
	//		fmt.Printf("%s - Duration: %f seconds", bytesWritten, execDuration)
	//	} else {
	fmt.Printf("Duration: %f seconds", execDuration)
	//	}

}

func produceLogMessage(topic string, kafkaServers *string, logMessage swagger.LogMessage) {

	topic = strings.ToLower(fmt.Sprintf("algorun.%s.%s.algo.%s.%s",
		logMessage.EndpointOwnerUserName,
		logMessage.EndpointUrlName,
		logMessage.AlgoOwnerUserName,
		logMessage.AlgoUrlName))

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": *kafkaServers})

	if err != nil {
		fmt.Printf("Failed to create server message producer: %s\n", err)
		os.Exit(1)
	}

	doneChan := make(chan bool)

	go func() {
		defer close(doneChan)
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				m := ev
				if m.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
				} else {
					fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
						*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
				}
				return

			default:
				fmt.Printf("Ignored event: %s\n", ev)
			}
		}
	}()

	logMessageBytes, err := json.Marshal(logMessage)

	p.ProduceChannel() <- &kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value: logMessageBytes}

	// wait for delivery report goroutine to finish
	_ = <-doneChan

	p.Close()

}
