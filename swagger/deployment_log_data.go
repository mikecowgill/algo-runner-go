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

type DeploymentLogData struct {

	OrchEventType string `json:"orchEventType,omitempty"`

	EndpointOwnerUserName string `json:"endpointOwnerUserName,omitempty"`

	EndpointName string `json:"endpointName,omitempty"`

	AlgoOwnerName string `json:"algoOwnerName,omitempty"`

	AlgoName string `json:"algoName,omitempty"`

	AlgoVersionTag string `json:"algoVersionTag,omitempty"`

	AlgoIndex int32 `json:"algoIndex,omitempty"`

	Name string `json:"name,omitempty"`

	Desired int32 `json:"desired,omitempty"`

	Current int32 `json:"current,omitempty"`

	UpToDate int32 `json:"upToDate,omitempty"`

	Available int32 `json:"available,omitempty"`

	CreatedTimestamp time.Time `json:"createdTimestamp,omitempty"`
}