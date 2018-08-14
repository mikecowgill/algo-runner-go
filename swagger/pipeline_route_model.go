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

import (
	"time"
)

type PipelineRouteModel struct {

	RouteType string `json:"routeType,omitempty"`

	PipelineEndpointSourceOutputName string `json:"pipelineEndpointSourceOutputName,omitempty"`

	PipelineDataSourceName string `json:"pipelineDataSourceName,omitempty"`

	PipelineDataSourceIndex int32 `json:"pipelineDataSourceIndex,omitempty"`

	SourceAlgoOwnerName string `json:"sourceAlgoOwnerName,omitempty"`

	SourceAlgoName string `json:"sourceAlgoName,omitempty"`

	SourceAlgoIndex int32 `json:"sourceAlgoIndex,omitempty"`

	SourceAlgoOutputName string `json:"sourceAlgoOutputName,omitempty"`

	SourceAlgoOutputMessageDataType string `json:"sourceAlgoOutputMessageDataType,omitempty"`

	PipelineDataSinkName string `json:"pipelineDataSinkName,omitempty"`

	PipelineDataSinkIndex int32 `json:"pipelineDataSinkIndex,omitempty"`

	DestAlgoOwnerName string `json:"destAlgoOwnerName,omitempty"`

	DestAlgoName string `json:"destAlgoName,omitempty"`

	DestAlgoIndex int32 `json:"destAlgoIndex,omitempty"`

	DestAlgoInputName string `json:"destAlgoInputName,omitempty"`

	CreatedTimestamp time.Time `json:"createdTimestamp,omitempty"`

	UpdatedTimestamp time.Time `json:"updatedTimestamp,omitempty"`
}
