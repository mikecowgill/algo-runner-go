#!/bin/bash

# Local instance must be running to pull the swagger.json file
java -jar ./openapi-generator-cli-4.2.3.jar generate -i http://localhost:5000/swagger/v1-beta1/swagger.json \
-g go \
-p enumClassPrefix=true \
-t openapi-template \
-o algorun-go-client

mkdir -p ./pkg/openapi/
cp ./algorun-go-client/model_algo_runner_config.go ./pkg/openapi/
cp ./algorun-go-client/model_container_image_model.go ./pkg/openapi/
cp ./algorun-go-client/model_resource_requirements_v1.go ./pkg/openapi/
cp ./algorun-go-client/model_auto_scaling_spec.go ./pkg/openapi/
cp ./algorun-go-client/model_scale_metric_model.go ./pkg/openapi/
cp ./algorun-go-client/model_algo_param_spec.go ./pkg/openapi/
cp ./algorun-go-client/model_algo_input_spec.go ./pkg/openapi/
cp ./algorun-go-client/model_algo_output_spec.go ./pkg/openapi/
cp ./algorun-go-client/model_topic_config_model.go ./pkg/openapi/
cp ./algorun-go-client/model_topic_param_model.go ./pkg/openapi/
cp ./algorun-go-client/model_topic_retry_strategy_model.go ./pkg/openapi/
cp ./algorun-go-client/model_topic_retry_step_model.go ./pkg/openapi/
cp ./algorun-go-client/model_data_type_model.go ./pkg/openapi/
cp ./algorun-go-client/model_content_type_model.go ./pkg/openapi/
cp ./algorun-go-client/model_pipe_model.go ./pkg/openapi/
cp ./algorun-go-client/model_pipeline_data_connector_model.go ./pkg/openapi/
cp ./algorun-go-client/model_data_connector_model.go ./pkg/openapi/
cp ./algorun-go-client/model_data_connector_version_model.go ./pkg/openapi/
cp ./algorun-go-client/model_data_connector_option_model.go ./pkg/openapi/
cp ./algorun-go-client/model_algo_spec.go ./pkg/openapi/
cp ./algorun-go-client/model_log_entry_model.go ./pkg/openapi/
cp ./algorun-go-client/model_file_reference.go ./pkg/openapi/
cp ./algorun-go-client/model_config_mount_model.go ./pkg/openapi/
cp ./algorun-go-client/model_probe_v1.go ./pkg/openapi/
cp ./algorun-go-client/model_exec_action_v1.go ./pkg/openapi/
cp ./algorun-go-client/model_http_get_action_v1.go ./pkg/openapi/
cp ./algorun-go-client/model_http_header_v1.go ./pkg/openapi/
cp ./algorun-go-client/model_tcp_socket_action_v1.go ./pkg/openapi/
cp ./algorun-go-client/model_int32_or_string_v1.go ./pkg/openapi/

cp ./algorun-go-client/model_image_pull_policies.go ./pkg/openapi/
cp ./algorun-go-client/model_retry_strategies.go ./pkg/openapi/
cp ./algorun-go-client/model_log_levels.go ./pkg/openapi/
cp ./algorun-go-client/model_log_types.go ./pkg/openapi/
cp ./algorun-go-client/model_executors.go ./pkg/openapi/
cp ./algorun-go-client/model_input_delivery_types.go ./pkg/openapi/
cp ./algorun-go-client/model_output_delivery_types.go ./pkg/openapi/
cp ./algorun-go-client/model_data_connector_types.go ./pkg/openapi/
cp ./algorun-go-client/model_data_types.go ./pkg/openapi/
cp ./algorun-go-client/model_component_types.go ./pkg/openapi/
cp ./algorun-go-client/model_message_data_types.go ./pkg/openapi/
cp ./algorun-go-client/model_metric_source_types.go ./pkg/openapi/
cp ./algorun-go-client/model_metric_target_types.go ./pkg/openapi/

rm -rf ./algorun-go-client/

