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

type ScheduleResponse struct {
	Uid string `json:"uid,omitempty"`
	Name string `json:"name,omitempty"`
	Schedule string `json:"schedule,omitempty"`
	Ttl string `json:"ttl,omitempty"`
	Labels Labels `json:"labels,omitempty"`
	Options BackupOptions `json:"options,omitempty"`
	Status string `json:"status,omitempty"`
	LastBackup string `json:"lastBackup,omitempty"`
}
