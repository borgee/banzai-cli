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

type NodeItemStatusCapacity struct {
	Cpu string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
	Pods string `json:"pods,omitempty"`
}
