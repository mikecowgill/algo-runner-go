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
// ResourceRequirementsV1 struct for ResourceRequirementsV1
type ResourceRequirementsV1 struct {
	Limits map[string]string `json:"limits,omitempty"`
	Requests map[string]string `json:"requests,omitempty"`
}