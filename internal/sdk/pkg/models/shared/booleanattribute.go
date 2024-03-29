// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-entity/internal/sdk/pkg/utils"
)

// BooleanAttributeConstraints - A set of constraints applicable to the attribute.
// These constraints should and will be enforced by the attribute renderer.
type BooleanAttributeConstraints struct {
}

type BooleanAttributeType string

const (
	BooleanAttributeTypeBoolean BooleanAttributeType = "boolean"
)

func (e BooleanAttributeType) ToPointer() *BooleanAttributeType {
	return &e
}

func (e *BooleanAttributeType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "boolean":
		*e = BooleanAttributeType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BooleanAttributeType: %v", v)
	}
}

// BooleanAttribute - Yes / No Toggle
type BooleanAttribute struct {
	Purpose []string `json:"_purpose,omitempty"`
	// A set of constraints applicable to the attribute.
	// These constraints should and will be enforced by the attribute renderer.
	//
	Constraints  *BooleanAttributeConstraints `json:"constraints,omitempty"`
	DefaultValue interface{}                  `json:"default_value,omitempty"`
	Deprecated   *bool                        `default:"false" json:"deprecated"`
	// Setting to `true` disables editing the attribute on the entity builder UI
	EntityBuilderDisableEdit *bool `default:"false" json:"entity_builder_disable_edit"`
	// This attribute should only be active when the feature flag is enabled
	FeatureFlag *string `json:"feature_flag,omitempty"`
	// Which group the attribute should appear in. Accepts group ID or group name
	Group *string `json:"group,omitempty"`
	// Do not render attribute in entity views
	Hidden *bool `default:"false" json:"hidden"`
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
	Protected *bool `default:"true" json:"protected"`
	Readonly  *bool `default:"false" json:"readonly"`
	// Defines the conditional rendering expression for showing this field.
	// When a valid expression is parsed, their evaluation defines the visibility of this attribute.
	// Note: Empty or invalid expression have no effect on the field visibility.
	//
	RenderCondition *string `json:"render_condition,omitempty"`
	Required        *bool   `default:"false" json:"required"`
	// This attribute should only be active when the setting is enabled
	SettingFlag *string `json:"setting_flag,omitempty"`
	// Render as a column in table views. When defined, overrides `hidden`
	ShowInTable    *bool                 `json:"show_in_table,omitempty"`
	Type           *BooleanAttributeType `json:"type,omitempty"`
	ValueFormatter *string               `json:"value_formatter,omitempty"`
}

func (b BooleanAttribute) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BooleanAttribute) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *BooleanAttribute) GetPurpose() []string {
	if o == nil {
		return nil
	}
	return o.Purpose
}

func (o *BooleanAttribute) GetConstraints() *BooleanAttributeConstraints {
	if o == nil {
		return nil
	}
	return o.Constraints
}

func (o *BooleanAttribute) GetDefaultValue() interface{} {
	if o == nil {
		return nil
	}
	return o.DefaultValue
}

func (o *BooleanAttribute) GetDeprecated() *bool {
	if o == nil {
		return nil
	}
	return o.Deprecated
}

func (o *BooleanAttribute) GetEntityBuilderDisableEdit() *bool {
	if o == nil {
		return nil
	}
	return o.EntityBuilderDisableEdit
}

func (o *BooleanAttribute) GetFeatureFlag() *string {
	if o == nil {
		return nil
	}
	return o.FeatureFlag
}

func (o *BooleanAttribute) GetGroup() *string {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *BooleanAttribute) GetHidden() *bool {
	if o == nil {
		return nil
	}
	return o.Hidden
}

func (o *BooleanAttribute) GetHideLabel() *bool {
	if o == nil {
		return nil
	}
	return o.HideLabel
}

func (o *BooleanAttribute) GetIcon() *string {
	if o == nil {
		return nil
	}
	return o.Icon
}

func (o *BooleanAttribute) GetLabel() string {
	if o == nil {
		return ""
	}
	return o.Label
}

func (o *BooleanAttribute) GetLayout() *string {
	if o == nil {
		return nil
	}
	return o.Layout
}

func (o *BooleanAttribute) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *BooleanAttribute) GetOrder() *int64 {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *BooleanAttribute) GetPlaceholder() *string {
	if o == nil {
		return nil
	}
	return o.Placeholder
}

func (o *BooleanAttribute) GetPreviewValueFormatter() *string {
	if o == nil {
		return nil
	}
	return o.PreviewValueFormatter
}

func (o *BooleanAttribute) GetProtected() *bool {
	if o == nil {
		return nil
	}
	return o.Protected
}

func (o *BooleanAttribute) GetReadonly() *bool {
	if o == nil {
		return nil
	}
	return o.Readonly
}

func (o *BooleanAttribute) GetRenderCondition() *string {
	if o == nil {
		return nil
	}
	return o.RenderCondition
}

func (o *BooleanAttribute) GetRequired() *bool {
	if o == nil {
		return nil
	}
	return o.Required
}

func (o *BooleanAttribute) GetSettingFlag() *string {
	if o == nil {
		return nil
	}
	return o.SettingFlag
}

func (o *BooleanAttribute) GetShowInTable() *bool {
	if o == nil {
		return nil
	}
	return o.ShowInTable
}

func (o *BooleanAttribute) GetType() *BooleanAttributeType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *BooleanAttribute) GetValueFormatter() *string {
	if o == nil {
		return nil
	}
	return o.ValueFormatter
}
