// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type NumberAttributeType string

const (
	NumberAttributeTypeNumber NumberAttributeType = "number"
)

func (e NumberAttributeType) ToPointer() *NumberAttributeType {
	return &e
}

func (e *NumberAttributeType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "number":
		*e = NumberAttributeType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NumberAttributeType: %v", v)
	}
}

// NumberAttribute - Numeric input
type NumberAttribute struct {
	Purpose []string `json:"_purpose,omitempty"`
	// A set of constraints applicable to the attribute.
	// These constraints should and will be enforced by the attribute renderer.
	//
	Constraints  map[string]interface{} `json:"constraints,omitempty"`
	DefaultValue interface{}            `json:"default_value,omitempty"`
	Deprecated   *bool                  `json:"deprecated,omitempty"`
	// Setting to `true` disables editing the attribute on the entity builder UI
	EntityBuilderDisableEdit *bool `json:"entity_builder_disable_edit,omitempty"`
	// This attribute should only be active when the feature flag is enabled
	FeatureFlag *string `json:"feature_flag,omitempty"`
	Format      *string `json:"format,omitempty"`
	// Which group the attribute should appear in. Accepts group ID or group name
	Group *string `json:"group,omitempty"`
	// Do not render attribute in entity views
	Hidden *bool `json:"hidden,omitempty"`
	// When set to true, will hide the label of the field.
	HideLabel *bool `json:"hide_label,omitempty"`
	// Code name of the icon to used to represent this attribute.
	// The value must be a valid @epilot/base-elements Icon name
	//
	Icon   *string `json:"icon,omitempty"`
	Label  string  `json:"label"`
	Layout *string `json:"layout,omitempty"`
	Name   string  `json:"name"`
	// Attribute sort order (ascending) in group
	Order                 *int64  `json:"order,omitempty"`
	Placeholder           *string `json:"placeholder,omitempty"`
	PreviewValueFormatter *string `json:"preview_value_formatter,omitempty"`
	// Setting to `true` prevents the attribute from being modified / deleted
	Protected *bool `json:"protected,omitempty"`
	Readonly  *bool `json:"readonly,omitempty"`
	// Defines the conditional rendering expression for showing this field.
	// When a valid expression is parsed, their evaluation defines the visibility of this attribute.
	// Note: Empty or invalid expression have no effect on the field visibility.
	//
	RenderCondition *string `json:"render_condition,omitempty"`
	Required        *bool   `json:"required,omitempty"`
	// This attribute should only be active when the setting is enabled
	SettingFlag *string `json:"setting_flag,omitempty"`
	// Render as a column in table views. When defined, overrides `hidden`
	ShowInTable    *bool                `json:"show_in_table,omitempty"`
	Type           *NumberAttributeType `json:"type,omitempty"`
	ValueFormatter *string              `json:"value_formatter,omitempty"`
}
