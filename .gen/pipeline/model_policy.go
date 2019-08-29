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

type Policy struct {
	Id string `json:"id"`
	Name string `json:"name,omitempty"`
	Comment string `json:"comment,omitempty"`
	Version string `json:"version"`
	Rules []PolicyRule `json:"rules,omitempty"`
}
