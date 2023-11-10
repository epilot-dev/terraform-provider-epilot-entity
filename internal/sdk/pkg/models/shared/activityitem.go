// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/epilot-dev/terraform-provider-epilot-entity/internal/sdk/pkg/utils"
	"time"
)

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

func (a ActivityItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *ActivityItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ActivityItem) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ActivityItem) GetCaller() *ActivityCallerContext {
	if o == nil {
		return nil
	}
	return o.Caller
}

func (o *ActivityItem) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *ActivityItem) GetOperations() []EntityOperation {
	if o == nil {
		return nil
	}
	return o.Operations
}

func (o *ActivityItem) GetOperationsTotal() *int64 {
	if o == nil {
		return nil
	}
	return o.OperationsTotal
}

func (o *ActivityItem) GetPayload() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Payload
}

func (o *ActivityItem) GetTimestamp() *time.Time {
	if o == nil {
		return nil
	}
	return o.Timestamp
}

func (o *ActivityItem) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *ActivityItem) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}
