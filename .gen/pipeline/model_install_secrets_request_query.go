/*
 * Pipeline API
 *
 * Pipeline v0.3.0 swagger
 *
 * API version: 0.23.0
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline

type InstallSecretsRequestQuery struct {
	Type string `json:"type,omitempty"`
	Ids []string `json:"ids,omitempty"`
	Tags []string `json:"tags,omitempty"`
}
