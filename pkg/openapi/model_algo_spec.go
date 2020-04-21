/*
 * Algo.Run API 1.0-beta1
 *
 * API for the Algo.Run Engine
 *
 * API version: 1.0-beta1
 * Contact: support@algohub.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// AlgoSpec struct for AlgoSpec
type AlgoSpec struct {
	Owner string `json:"owner"`
	Name string `json:"name"`
	Version string `json:"version"`
	Index int32 `json:"index"`
	Image *ContainerImageModel `json:"image"`
	Entrypoint string `json:"entrypoint,omitempty"`
	Executor *Executors `json:"executor,omitempty"`
	ConfigMounts []ConfigMountModel `json:"configMounts,omitempty"`
	Parameters []AlgoParamSpec `json:"parameters,omitempty"`
	Inputs []AlgoInputSpec `json:"inputs,omitempty"`
	Outputs []AlgoOutputSpec `json:"outputs,omitempty"`
	RetryEnabled bool `json:"retryEnabled,omitempty"`
	RetryStrategy *TopicRetryStrategyModel `json:"retryStrategy,omitempty"`
	WriteAllOutputs bool `json:"writeAllOutputs,omitempty"`
	GpuEnabled bool `json:"gpuEnabled,omitempty"`
	TimeoutSeconds int32 `json:"timeoutSeconds,omitempty"`
	Replicas int32 `json:"replicas,omitempty"`
	Resources *ResourceRequirementsV1 `json:"resources,omitempty"`
	Autoscaling *AutoScalingSpec `json:"autoscaling,omitempty"`
	AlgoRunnerImage *ContainerImageModel `json:"algoRunnerImage,omitempty"`
	LivenessProbe *ProbeV1 `json:"livenessProbe,omitempty"`
	ReadinessProbe *ProbeV1 `json:"readinessProbe,omitempty"`
}
