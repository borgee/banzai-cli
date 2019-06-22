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

type CreateAzureObjectStoreBucketProperties struct {
	StorageAccount string `json:"storageAccount"`
	Location string `json:"location"`
	ResourceGroup string `json:"resourceGroup"`
}
