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

type NodeItemStatusNodeInfo struct {
	MachineID string `json:"machineID,omitempty"`
	SystemUUID string `json:"systemUUID,omitempty"`
	BootID string `json:"bootID,omitempty"`
	KernelVersion string `json:"kernelVersion,omitempty"`
	OsImage string `json:"osImage,omitempty"`
	ContainerRuntimeVersion string `json:"containerRuntimeVersion,omitempty"`
	KubeletVersion string `json:"kubeletVersion,omitempty"`
	KubeProxyVersion string `json:"kubeProxyVersion,omitempty"`
	OperatingSystem string `json:"operatingSystem,omitempty"`
	Architecture string `json:"architecture,omitempty"`
}
