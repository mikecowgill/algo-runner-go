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
// DataTypeModel struct for DataTypeModel
type DataTypeModel struct {
	Name DataTypes `json:"name"`
	Regex *string `json:"regex,omitempty"`
	Precision *int32 `json:"precision,omitempty"`
	Scale *int32 `json:"scale,omitempty"`
	Mask *string `json:"mask,omitempty"`
}
