/*
 * Pipeline API
 *
 * Pipeline is a feature rich application platform, built for containers on top of Kubernetes to automate the DevOps experience, continuous application development and the lifecycle of deployments. 
 *
 * API version: latest
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline
// ClusterNotFound struct for ClusterNotFound
type ClusterNotFound struct {
	Code int32 `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Error string `json:"error,omitempty"`
}
