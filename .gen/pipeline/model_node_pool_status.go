/*
 * Pipeline API
 *
 * Pipeline v0.3.0 swagger
 *
 * API version: pke-oidc-refactor
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline

type NodePoolStatus struct {
	InstanceType string `json:"instanceType,omitempty"`
	SpotPrice string `json:"spot_price,omitempty"`
	Autoscaling bool `json:"autoscaling,omitempty"`
	Count int32 `json:"count,omitempty"`
	MinCount int32 `json:"minCount,omitempty"`
	MaxCount int32 `json:"maxCount,omitempty"`
	Image string `json:"image,omitempty"`
	Labels map[string]string `json:"labels,omitempty"`
}
