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

type ServerLogModel struct {
	EndpointOwnerUserName string `json:"endpointOwnerUserName,omitempty"`

	EndpointUrlName string `json:"endpointUrlName,omitempty"`

	AlgoOwnerUserName string `json:"algoOwnerUserName,omitempty"`

	AlgoUrlName string `json:"algoUrlName,omitempty"`

	AlgoVersionTag string `json:"algoVersionTag,omitempty"`

	Status string `json:"status,omitempty"`

	StdErr string `json:"stdErr,omitempty"`

	StdOut string `json:"stdOut,omitempty"`
}