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

type ApiClusterGroup struct {
	EnabledFeatures []string `json:"enabledFeatures,omitempty"`
	Id int32 `json:"id,omitempty"`
	Members []ApiMember `json:"members,omitempty"`
	Name string `json:"name,omitempty"`
	OrganizationId int32 `json:"organizationId,omitempty"`
	Uid string `json:"uid,omitempty"`
}
