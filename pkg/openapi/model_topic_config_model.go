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
// TopicConfigModel struct for TopicConfigModel
type TopicConfigModel struct {
	ComponentType *ComponentTypes `json:"componentType,omitempty"`
	SourceName string `json:"sourceName"`
	SourceOutputName string `json:"sourceOutputName"`
	TopicAutoName bool `json:"topicAutoName,omitempty"`
	TopicName string `json:"topicName,omitempty"`
	TopicAutoPartition bool `json:"topicAutoPartition,omitempty"`
	TopicPartitions int32 `json:"topicPartitions,omitempty"`
	TopicReplicationFactor int32 `json:"topicReplicationFactor,omitempty"`
	TopicParams []TopicParamModel `json:"topicParams,omitempty"`
	OverrideRetryStrategy bool `json:"overrideRetryStrategy,omitempty"`
	RetryStrategy *TopicRetryStrategyModel `json:"retryStrategy,omitempty"`
}
