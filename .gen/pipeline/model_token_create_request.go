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
import (
	"time"
)

type TokenCreateRequest struct {
	Name string `json:"name,omitempty"`
	VirtualUser string `json:"virtualUser,omitempty"`
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
}
