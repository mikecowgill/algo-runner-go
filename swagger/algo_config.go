/*
 * Algo.Run API 1.0
 *
 * API for the Algo.Run Engine
 *
 * API version: 1.0
 * Contact: support@algohub.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type AlgoConfig struct {

	Applied bool `json:"applied,omitempty"`

	DeploymentName string `json:"deploymentName,omitempty"`

	AlgoOwnerUserName string `json:"algoOwnerUserName,omitempty"`

	AlgoName string `json:"algoName,omitempty"`

	AlgoVersionTag string `json:"algoVersionTag,omitempty"`

	AlgoIndex int32 `json:"algoIndex,omitempty"`

	Container string `json:"container,omitempty"`

	Entrypoint string `json:"entrypoint,omitempty"`

	ServerType string `json:"serverType,omitempty"`

	AlgoParams []AlgoParamModel `json:"algoParams,omitempty"`

	Inputs []AlgoInputModel `json:"inputs,omitempty"`

	Outputs []AlgoOutputModel `json:"outputs,omitempty"`

	WriteAllOutputs bool `json:"writeAllOutputs,omitempty"`

	GpuEnabled bool `json:"gpuEnabled,omitempty"`

	TimeoutSeconds int32 `json:"timeoutSeconds,omitempty"`

	MemoryRequestBytes int64 `json:"memoryRequestBytes,omitempty"`

	MemoryLimitBytes int64 `json:"memoryLimitBytes,omitempty"`

	CpuRequestUnits float64 `json:"cpuRequestUnits,omitempty"`

	CpuLimitUnits float64 `json:"cpuLimitUnits,omitempty"`

	GpuLimitUnits float64 `json:"gpuLimitUnits,omitempty"`

	Instances int32 `json:"instances,omitempty"`

	AutoScale bool `json:"autoScale,omitempty"`

	MinInstances int32 `json:"minInstances,omitempty"`

	MaxInstances int32 `json:"maxInstances,omitempty"`

	AutoScaleCPUPercent int32 `json:"autoScaleCPUPercent,omitempty"`
}
