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

type RestoreResponse struct {
	Id int32 `json:"id,omitempty"`
	Uid string `json:"uid,omitempty"`
	Name string `json:"name,omitempty"`
	BackupName string `json:"backupName,omitempty"`
	Options BackupOptions `json:"options,omitempty"`
	Status string `json:"status,omitempty"`
	Warnings int32 `json:"warnings,omitempty"`
	Errors int32 `json:"errors,omitempty"`
}
