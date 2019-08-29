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

type EksVpc struct {
	// The identifier of existing VPC to be used for creating the EKS cluster. If not provided a new VPC is created for the cluster.
	VpcId string `json:"vpcId,omitempty"`
	// The CIDR range for the VPC in case new VPC is created.
	Cidr string `json:"cidr,omitempty"`
}
