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

type NodePoolStatusOracle struct {
	Count int32 `json:"count,omitempty"`
	MinCount int32 `json:"minCount,omitempty"`
	MaxCount int32 `json:"maxCount,omitempty"`
	InstanceType string `json:"instanceType,omitempty"`
	Image string `json:"image,omitempty"`
	Autoscaling bool `json:"autoscaling,omitempty"`
}
