// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// RelationAttributeActionsActionTypeEnum - The action type. Currently supported actions:
//
// | action | description |
// |--------|-------------|
// | add_existing | Enables the user to pick an existing entity to link as relation |
// | create_new | Enables the user to create a new entity using the first/main `allowed_schemas` schema
// | create_from_existing | Enables the user to pick an existing entity to clone from, while creating a blank new entity to link as relation |
type RelationAttributeActionsActionTypeEnum string

const (
	RelationAttributeActionsActionTypeEnumAddExisting        RelationAttributeActionsActionTypeEnum = "add_existing"
	RelationAttributeActionsActionTypeEnumCreateNew          RelationAttributeActionsActionTypeEnum = "create_new"
	RelationAttributeActionsActionTypeEnumCreateFromExisting RelationAttributeActionsActionTypeEnum = "create_from_existing"
)

func (e *RelationAttributeActionsActionTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "add_existing":
		fallthrough
	case "create_new":
		fallthrough
	case "create_from_existing":
		*e = RelationAttributeActionsActionTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for RelationAttributeActionsActionTypeEnum: %s", s)
	}
}

type RelationAttributeActionsNewEntityItem struct {
	CreatedAt time.Time `json:"_created_at"`
	ID        string    `json:"_id"`
	// Organization Id the entity belongs to
	Org string `json:"_org"`
	// URL-friendly identifier for the entity schema
	Schema string   `json:"_schema"`
	Tags   []string `json:"_tags,omitempty"`
	// Title of entity
	Title     string    `json:"_title"`
	UpdatedAt time.Time `json:"_updated_at"`

	AdditionalProperties map[string]interface{} `json:"-"`
}
type _RelationAttributeActionsNewEntityItem RelationAttributeActionsNewEntityItem

func (c *RelationAttributeActionsNewEntityItem) UnmarshalJSON(bs []byte) error {
	data := _RelationAttributeActionsNewEntityItem{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = RelationAttributeActionsNewEntityItem(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "_created_at")
	delete(additionalFields, "_id")
	delete(additionalFields, "_org")
	delete(additionalFields, "_schema")
	delete(additionalFields, "_tags")
	delete(additionalFields, "_title")
	delete(additionalFields, "_updated_at")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c RelationAttributeActionsNewEntityItem) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_RelationAttributeActionsNewEntityItem(c))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	bs, err = json.Marshal(c.AdditionalProperties)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	return json.Marshal(out)
}

type RelationAttributeActions struct {
	// The action type. Currently supported actions:
	//
	// | action | description |
	// |--------|-------------|
	// | add_existing | Enables the user to pick an existing entity to link as relation |
	// | create_new | Enables the user to create a new entity using the first/main `allowed_schemas` schema
	// | create_from_existing | Enables the user to pick an existing entity to clone from, while creating a blank new entity to link as relation |
	//
	ActionType *RelationAttributeActionsActionTypeEnum `json:"action_type,omitempty"`
	// Sets the action as the default action, visible as the main action button.
	Default *bool `json:"default,omitempty"`
	// Name of the feature flag that enables this action
	FeatureFlag *string `json:"feature_flag,omitempty"`
	// The action label or action translation key (i18n)
	Label         *string                                `json:"label,omitempty"`
	NewEntityItem *RelationAttributeActionsNewEntityItem `json:"new_entity_item,omitempty"`
	// Name of the setting flag that enables this action
	SettingFlag *string `json:"setting_flag,omitempty"`
}

type RelationAttributeDrawerSizeEnum string

const (
	RelationAttributeDrawerSizeEnumSmall  RelationAttributeDrawerSizeEnum = "small"
	RelationAttributeDrawerSizeEnumMedium RelationAttributeDrawerSizeEnum = "medium"
	RelationAttributeDrawerSizeEnumLarge  RelationAttributeDrawerSizeEnum = "large"
)

func (e *RelationAttributeDrawerSizeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "small":
		fallthrough
	case "medium":
		fallthrough
	case "large":
		*e = RelationAttributeDrawerSizeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for RelationAttributeDrawerSizeEnum: %s", s)
	}
}

type RelationAttributeEditModeEnum string

const (
	RelationAttributeEditModeEnumListView RelationAttributeEditModeEnum = "list-view"
)

func (e *RelationAttributeEditModeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "list-view":
		*e = RelationAttributeEditModeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for RelationAttributeEditModeEnum: %s", s)
	}
}

// RelationAttributeRelationAffinityModeEnum - Weak relation attributes are kept when duplicating an entity. Strong relation attributes are discarded when duplicating an entity.
type RelationAttributeRelationAffinityModeEnum string

const (
	RelationAttributeRelationAffinityModeEnumWeak   RelationAttributeRelationAffinityModeEnum = "weak"
	RelationAttributeRelationAffinityModeEnumStrong RelationAttributeRelationAffinityModeEnum = "strong"
)

func (e *RelationAttributeRelationAffinityModeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "weak":
		fallthrough
	case "strong":
		*e = RelationAttributeRelationAffinityModeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for RelationAttributeRelationAffinityModeEnum: %s", s)
	}
}

type RelationAttributeRelationTypeEnum string

const (
	RelationAttributeRelationTypeEnumHasMany RelationAttributeRelationTypeEnum = "has_many"
	RelationAttributeRelationTypeEnumHasOne  RelationAttributeRelationTypeEnum = "has_one"
)

func (e *RelationAttributeRelationTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "has_many":
		fallthrough
	case "has_one":
		*e = RelationAttributeRelationTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for RelationAttributeRelationTypeEnum: %s", s)
	}
}

type RelationAttributeSummaryFieldsType string

const (
	RelationAttributeSummaryFieldsTypeStr          RelationAttributeSummaryFieldsType = "str"
	RelationAttributeSummaryFieldsTypeSummaryField RelationAttributeSummaryFieldsType = "SummaryField"
)

type RelationAttributeSummaryFields struct {
	Str          *string
	SummaryField *SummaryField

	Type RelationAttributeSummaryFieldsType
}

func CreateRelationAttributeSummaryFieldsStr(str string) RelationAttributeSummaryFields {
	typ := RelationAttributeSummaryFieldsTypeStr

	return RelationAttributeSummaryFields{
		Str:  &str,
		Type: typ,
	}
}

func CreateRelationAttributeSummaryFieldsSummaryField(summaryField SummaryField) RelationAttributeSummaryFields {
	typ := RelationAttributeSummaryFieldsTypeSummaryField

	return RelationAttributeSummaryFields{
		SummaryField: &summaryField,
		Type:         typ,
	}
}

func (u *RelationAttributeSummaryFields) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = RelationAttributeSummaryFieldsTypeStr
		return nil
	}

	summaryField := new(SummaryField)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&summaryField); err == nil {
		u.SummaryField = summaryField
		u.Type = RelationAttributeSummaryFieldsTypeSummaryField
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u RelationAttributeSummaryFields) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.SummaryField != nil {
		return json.Marshal(u.SummaryField)
	}

	return nil, nil
}

type RelationAttributeTypeEnum string

const (
	RelationAttributeTypeEnumRelation RelationAttributeTypeEnum = "relation"
)

func (e *RelationAttributeTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "relation":
		*e = RelationAttributeTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for RelationAttributeTypeEnum: %s", s)
	}
}

// RelationAttribute - Entity Relationship
type RelationAttribute struct {
	Purpose []string                   `json:"_purpose,omitempty"`
	Actions []RelationAttributeActions `json:"actions,omitempty"`
	// Optional label for the add button. The translated value for add_button_lable is used, if found else the string is used as is.
	AddButtonLabel *string     `json:"add_button_label,omitempty"`
	AllowedSchemas []string    `json:"allowedSchemas,omitempty"`
	DefaultValue   interface{} `json:"default_value,omitempty"`
	Deprecated     *bool       `json:"deprecated,omitempty"`
	// Enables the preview, edition, and creation of relation items on a Master-Details view mode.
	DetailsViewModeEnabled *bool                            `json:"details_view_mode_enabled,omitempty"`
	DrawerSize             *RelationAttributeDrawerSizeEnum `json:"drawer_size,omitempty"`
	EditMode               *RelationAttributeEditModeEnum   `json:"edit_mode,omitempty"`
	// When enable_relation_picker is set to true the user will be able to pick existing relations as values. Otherwise, the user will need to create new relation to link.
	EnableRelationPicker *bool `json:"enable_relation_picker,omitempty"`
	// When enable_relation_tags is set to true the user will be able to set tags(labels) in each relation item.
	EnableRelationTags *bool `json:"enable_relation_tags,omitempty"`
	// Setting to `true` disables editing the attribute on the entity builder UI
	EntityBuilderDisableEdit *bool `json:"entity_builder_disable_edit,omitempty"`
	// This attribute should only be active when the feature flag is enabled
	FeatureFlag *string `json:"feature_flag,omitempty"`
	// Which group the attribute should appear in. Accepts group ID or group name
	Group      *string `json:"group,omitempty"`
	HasPrimary *bool   `json:"has_primary,omitempty"`
	// Do not render attribute in entity views
	Hidden *bool `json:"hidden,omitempty"`
	// When set to true, will hide the label of the field.
	HideLabel *bool   `json:"hide_label,omitempty"`
	Icon      *string `json:"icon,omitempty"`
	Label     string  `json:"label"`
	Layout    *string `json:"layout,omitempty"`
	Name      string  `json:"name"`
	// Attribute sort order (ascending) in group
	Order                 *int64  `json:"order,omitempty"`
	Placeholder           *string `json:"placeholder,omitempty"`
	PreviewValueFormatter *string `json:"preview_value_formatter,omitempty"`
	// Setting to `true` prevents the attribute from being modified / deleted
	Protected *bool `json:"protected,omitempty"`
	Readonly  *bool `json:"readonly,omitempty"`
	// Weak relation attributes are kept when duplicating an entity. Strong relation attributes are discarded when duplicating an entity.
	RelationAffinityMode *RelationAttributeRelationAffinityModeEnum `json:"relation_affinity_mode,omitempty"`
	RelationType         *RelationAttributeRelationTypeEnum         `json:"relation_type,omitempty"`
	// Defines the conditional rendering expression for showing this field.
	// When a valid expression is parsed, their evaluation defines the visibility of this attribute.
	// Note: Empty or invalid expression have no effect on the field visibility.
	//
	RenderCondition *string `json:"render_condition,omitempty"`
	Required        *bool   `json:"required,omitempty"`
	// Map of schema slug to target relation attribute
	ReverseAttributes map[string]string `json:"reverse_attributes,omitempty"`
	// Optional placeholder text for the relation search input. The translated value for search_placeholder is used, if found else the string is used as is.
	SearchPlaceholder *string `json:"search_placeholder,omitempty"`
	// This attribute should only be active when the setting is enabled
	SettingFlag *string `json:"setting_flag,omitempty"`
	// Render as a column in table views. When defined, overrides `hidden`
	ShowInTable    *bool                            `json:"show_in_table,omitempty"`
	SummaryFields  []RelationAttributeSummaryFields `json:"summary_fields,omitempty"`
	Type           *RelationAttributeTypeEnum       `json:"type,omitempty"`
	ValueFormatter *string                          `json:"value_formatter,omitempty"`
}
