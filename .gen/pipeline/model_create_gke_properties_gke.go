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

type CreateGkePropertiesGke struct {
	ProjectId string `json:"projectId,omitempty"`
	Master CreateGkePropertiesGkeMaster `json:"master,omitempty"`
	NodeVersion string `json:"nodeVersion,omitempty"`
	// Name of the GCP Network (VPC) to deploy the cluster to. If omitted than the \"default\" VPC is used.
	Vpc string `json:"vpc,omitempty"`
	// Name of the GCP Subnet to deploy the cluster to. If \"default\" VPC is used this field can be omitted. The subnet must be in the same region as the location of the cluster.
	Subnet string `json:"subnet,omitempty"`
	NodePools map[string]NodePoolsGoogle `json:"nodePools"`
}
