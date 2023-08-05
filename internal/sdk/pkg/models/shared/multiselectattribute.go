// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// MultiSelectAttributeConstraints - A set of constraints applicable to the attribute.
// These constraints should and will be enforced by the attribute renderer.
type MultiSelectAttributeConstraints struct {
}

type MultiSelectAttributeOptions2 struct {
	Title *string `json:"title,omitempty"`
	Value string  `json:"value"`
}

type MultiSelectAttributeOptionsType string

const (
	MultiSelectAttributeOptionsTypeStr                          MultiSelectAttributeOptionsType = "str"
	MultiSelectAttributeOptionsTypeMultiSelectAttributeOptions2 MultiSelectAttributeOptionsType = "MultiSelectAttribute_options_2"
)

type MultiSelectAttributeOptions struct {
	Str                          *string
	MultiSelectAttributeOptions2 *MultiSelectAttributeOptions2

	Type MultiSelectAttributeOptionsType
}

func CreateMultiSelectAttributeOptionsStr(str string) MultiSelectAttributeOptions {
	typ := MultiSelectAttributeOptionsTypeStr

	return MultiSelectAttributeOptions{
		Str:  &str,
		Type: typ,
	}
}

func CreateMultiSelectAttributeOptionsMultiSelectAttributeOptions2(multiSelectAttributeOptions2 MultiSelectAttributeOptions2) MultiSelectAttributeOptions {
	typ := MultiSelectAttributeOptionsTypeMultiSelectAttributeOptions2

	return MultiSelectAttributeOptions{
		MultiSelectAttributeOptions2: &multiSelectAttributeOptions2,
		Type:                         typ,
	}
}

func (u *MultiSelectAttributeOptions) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = MultiSelectAttributeOptionsTypeStr
		return nil
	}

	multiSelectAttributeOptions2 := new(MultiSelectAttributeOptions2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&multiSelectAttributeOptions2); err == nil {
		u.MultiSelectAttributeOptions2 = multiSelectAttributeOptions2
		u.Type = MultiSelectAttributeOptionsTypeMultiSelectAttributeOptions2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u MultiSelectAttributeOptions) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.MultiSelectAttributeOptions2 != nil {
		return json.Marshal(u.MultiSelectAttributeOptions2)
	}

	return nil, nil
}

type MultiSelectAttributeType string

const (
	MultiSelectAttributeTypeMultiselect MultiSelectAttributeType = "multiselect"
	MultiSelectAttributeTypeCheckbox    MultiSelectAttributeType = "checkbox"
)

func (e MultiSelectAttributeType) ToPointer() *MultiSelectAttributeType {
	return &e
}

func (e *MultiSelectAttributeType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "multiselect":
		fallthrough
	case "checkbox":
		*e = MultiSelectAttributeType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MultiSelectAttributeType: %v", v)
	}
}

// MultiSelectAttribute - Multi Choice Selection
type MultiSelectAttribute struct {
	Purpose []string `json:"_purpose,omitempty"`
	// Allow arbitrary input values in addition to provided options
	AllowAny *bool `json:"allow_any,omitempty"`
	// controls if the 360 ui will allow the user to enter a value which is not defined by the options
	AllowExtraOptions *bool `json:"allow_extra_options,omitempty"`
	// A set of constraints applicable to the attribute.
	// These constraints should and will be enforced by the attribute renderer.
	//
	Constraints  *MultiSelectAttributeConstraints `json:"constraints,omitempty"`
	DefaultValue interface{}                      `json:"default_value,omitempty"`
	Deprecated   *bool                            `json:"deprecated,omitempty"`
	// controls if the matching of values against the options is case sensitive or not
	DisableCaseSensitive *bool `json:"disable_case_sensitive,omitempty"`
	// Setting to `true` disables editing the attribute on the entity builder UI
	EntityBuilderDisableEdit *bool `json:"entity_builder_disable_edit,omitempty"`
	// This attribute should only be active when the feature flag is enabled
	FeatureFlag *string `json:"feature_flag,omitempty"`
	// Which group the attribute should appear in. Accepts group ID or group name
	Group *string `json:"group,omitempty"`
	// Do not render attribute in entity views
	Hidden *bool `json:"hidden,omitempty"`
	// When set to true, will hide the label of the field.
	HideLabel *bool `json:"hide_label,omitempty"`
	// Code name of the icon to used to represent this attribute.
	// The value must be a valid @epilot/base-elements Icon name
	//
	Icon    *string                       `json:"icon,omitempty"`
	Label   string                        `json:"label"`
	Layout  *string                       `json:"layout,omitempty"`
	Name    string                        `json:"name"`
	Options []MultiSelectAttributeOptions `json:"options,omitempty"`
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
	ShowInTable    *bool                     `json:"show_in_table,omitempty"`
	Type           *MultiSelectAttributeType `json:"type,omitempty"`
	ValueFormatter *string                   `json:"value_formatter,omitempty"`
}
