// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// ActivityItem - Success
type ActivityItem struct {
	// See https://github.com/ulid/spec
	ID     *string                `json:"_id,omitempty"`
	Caller *ActivityCallerContext `json:"caller,omitempty"`
	// Message for activity. Supports handlebars syntax.
	Message    string            `json:"message"`
	Operations []EntityOperation `json:"operations,omitempty"`
	// Count of total operations attached to this activity
	OperationsTotal *int64                 `json:"operations_total,omitempty"`
	Payload         map[string]interface{} `json:"payload,omitempty"`
	Timestamp       *time.Time             `json:"timestamp,omitempty"`
	// Title for activity. Supports handlebars syntax.
	Title string `json:"title"`
	Type  string `json:"type"`
}
