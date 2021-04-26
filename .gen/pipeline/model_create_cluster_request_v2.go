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
// CreateClusterRequestV2 struct for CreateClusterRequestV2
type CreateClusterRequestV2 struct {
	Name string `json:"name"`
	SecretId string `json:"secretId,omitempty"`
	SecretName string `json:"secretName,omitempty"`
	SshSecretId string `json:"sshSecretId,omitempty"`
	Type string `json:"type"`
}
