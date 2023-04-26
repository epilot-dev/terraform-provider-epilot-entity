// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
)

type EntitySchemaGroupSettingsInfoTooltipTitle struct {
	Default *string `json:"default,omitempty"`
	Key     *string `json:"key,omitempty"`
}

type EntitySchemaGroupSettings struct {
	Purpose  []string `json:"_purpose,omitempty"`
	Expanded *bool    `json:"expanded,omitempty"`
	// This group should only be active when the feature flag is enabled
	FeatureFlag      *string                                    `json:"feature_flag,omitempty"`
	ID               string                                     `json:"id"`
	InfoTooltipTitle *EntitySchemaGroupSettingsInfoTooltipTitle `json:"info_tooltip_title,omitempty"`
	Label            string                                     `json:"label"`
	// Render order of the group
	Order           *int64  `json:"order,omitempty"`
	RenderCondition *string `json:"render_condition,omitempty"`
	// This group should only be active when the setting is enabled
	SettingFlag *string `json:"setting_flag,omitempty"`
}

// EntitySchemaLayoutSettings - Custom grid definitions for the layout. These settings are composed by managed and un-managed properties:
// - Managed Properties: are interpreted and transformed into layout styles
// - Un-managed Properties: are appended as styles into the attribute mounting node
type EntitySchemaLayoutSettings struct {
	// Defines the grid gap for the mounting node of the attribute.
	GridGap *string `json:"grid_gap,omitempty"`
	// Defines the grid column template for the mounting node of the attribute.
	GridTemplateColumns *string `json:"grid_template_columns,omitempty"`

	AdditionalProperties map[string]interface{} `json:"-"`
}
type _EntitySchemaLayoutSettings EntitySchemaLayoutSettings

func (c *EntitySchemaLayoutSettings) UnmarshalJSON(bs []byte) error {
	data := _EntitySchemaLayoutSettings{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = EntitySchemaLayoutSettings(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "grid_gap")
	delete(additionalFields, "grid_template_columns")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c EntitySchemaLayoutSettings) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_EntitySchemaLayoutSettings(c))
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

type EntitySchemaUIConfigCreateViewType string

const (
	EntitySchemaUIConfigCreateViewTypeEntityDefaultCreate EntitySchemaUIConfigCreateViewType = "EntityDefaultCreate"
	EntitySchemaUIConfigCreateViewTypeRedirectEntityView  EntitySchemaUIConfigCreateViewType = "RedirectEntityView"
	EntitySchemaUIConfigCreateViewTypeEntityViewDisabled  EntitySchemaUIConfigCreateViewType = "EntityViewDisabled"
)

type EntitySchemaUIConfigCreateView struct {
	EntityDefaultCreate *EntityDefaultCreate
	RedirectEntityView  *RedirectEntityView
	EntityViewDisabled  *EntityViewDisabled

	Type EntitySchemaUIConfigCreateViewType
}

func CreateEntitySchemaUIConfigCreateViewEntityDefaultCreate(entityDefaultCreate EntityDefaultCreate) EntitySchemaUIConfigCreateView {
	typ := EntitySchemaUIConfigCreateViewTypeEntityDefaultCreate

	return EntitySchemaUIConfigCreateView{
		EntityDefaultCreate: &entityDefaultCreate,
		Type:                typ,
	}
}

func CreateEntitySchemaUIConfigCreateViewRedirectEntityView(redirectEntityView RedirectEntityView) EntitySchemaUIConfigCreateView {
	typ := EntitySchemaUIConfigCreateViewTypeRedirectEntityView

	return EntitySchemaUIConfigCreateView{
		RedirectEntityView: &redirectEntityView,
		Type:               typ,
	}
}

func CreateEntitySchemaUIConfigCreateViewEntityViewDisabled(entityViewDisabled EntityViewDisabled) EntitySchemaUIConfigCreateView {
	typ := EntitySchemaUIConfigCreateViewTypeEntityViewDisabled

	return EntitySchemaUIConfigCreateView{
		EntityViewDisabled: &entityViewDisabled,
		Type:               typ,
	}
}

func (u *EntitySchemaUIConfigCreateView) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	entityDefaultCreate := new(EntityDefaultCreate)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&entityDefaultCreate); err == nil {
		u.EntityDefaultCreate = entityDefaultCreate
		u.Type = EntitySchemaUIConfigCreateViewTypeEntityDefaultCreate
		return nil
	}

	redirectEntityView := new(RedirectEntityView)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&redirectEntityView); err == nil {
		u.RedirectEntityView = redirectEntityView
		u.Type = EntitySchemaUIConfigCreateViewTypeRedirectEntityView
		return nil
	}

	entityViewDisabled := new(EntityViewDisabled)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&entityViewDisabled); err == nil {
		u.EntityViewDisabled = entityViewDisabled
		u.Type = EntitySchemaUIConfigCreateViewTypeEntityViewDisabled
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u EntitySchemaUIConfigCreateView) MarshalJSON() ([]byte, error) {
	if u.EntityDefaultCreate != nil {
		return json.Marshal(u.EntityDefaultCreate)
	}

	if u.RedirectEntityView != nil {
		return json.Marshal(u.RedirectEntityView)
	}

	if u.EntityViewDisabled != nil {
		return json.Marshal(u.EntityViewDisabled)
	}

	return nil, nil
}

type EntitySchemaUIConfigEditViewType string

const (
	EntitySchemaUIConfigEditViewTypeEntityDefaultEdit  EntitySchemaUIConfigEditViewType = "EntityDefaultEdit"
	EntitySchemaUIConfigEditViewTypeRedirectEntityView EntitySchemaUIConfigEditViewType = "RedirectEntityView"
	EntitySchemaUIConfigEditViewTypeEntityViewDisabled EntitySchemaUIConfigEditViewType = "EntityViewDisabled"
)

type EntitySchemaUIConfigEditView struct {
	EntityDefaultEdit  *EntityDefaultEdit
	RedirectEntityView *RedirectEntityView
	EntityViewDisabled *EntityViewDisabled

	Type EntitySchemaUIConfigEditViewType
}

func CreateEntitySchemaUIConfigEditViewEntityDefaultEdit(entityDefaultEdit EntityDefaultEdit) EntitySchemaUIConfigEditView {
	typ := EntitySchemaUIConfigEditViewTypeEntityDefaultEdit

	return EntitySchemaUIConfigEditView{
		EntityDefaultEdit: &entityDefaultEdit,
		Type:              typ,
	}
}

func CreateEntitySchemaUIConfigEditViewRedirectEntityView(redirectEntityView RedirectEntityView) EntitySchemaUIConfigEditView {
	typ := EntitySchemaUIConfigEditViewTypeRedirectEntityView

	return EntitySchemaUIConfigEditView{
		RedirectEntityView: &redirectEntityView,
		Type:               typ,
	}
}

func CreateEntitySchemaUIConfigEditViewEntityViewDisabled(entityViewDisabled EntityViewDisabled) EntitySchemaUIConfigEditView {
	typ := EntitySchemaUIConfigEditViewTypeEntityViewDisabled

	return EntitySchemaUIConfigEditView{
		EntityViewDisabled: &entityViewDisabled,
		Type:               typ,
	}
}

func (u *EntitySchemaUIConfigEditView) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	entityDefaultEdit := new(EntityDefaultEdit)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&entityDefaultEdit); err == nil {
		u.EntityDefaultEdit = entityDefaultEdit
		u.Type = EntitySchemaUIConfigEditViewTypeEntityDefaultEdit
		return nil
	}

	redirectEntityView := new(RedirectEntityView)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&redirectEntityView); err == nil {
		u.RedirectEntityView = redirectEntityView
		u.Type = EntitySchemaUIConfigEditViewTypeRedirectEntityView
		return nil
	}

	entityViewDisabled := new(EntityViewDisabled)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&entityViewDisabled); err == nil {
		u.EntityViewDisabled = entityViewDisabled
		u.Type = EntitySchemaUIConfigEditViewTypeEntityViewDisabled
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u EntitySchemaUIConfigEditView) MarshalJSON() ([]byte, error) {
	if u.EntityDefaultEdit != nil {
		return json.Marshal(u.EntityDefaultEdit)
	}

	if u.RedirectEntityView != nil {
		return json.Marshal(u.RedirectEntityView)
	}

	if u.EntityViewDisabled != nil {
		return json.Marshal(u.EntityViewDisabled)
	}

	return nil, nil
}

type EntitySchemaUIConfigListItemSummaryAttributesType string

const (
	EntitySchemaUIConfigListItemSummaryAttributesTypeSummaryAttribute EntitySchemaUIConfigListItemSummaryAttributesType = "SummaryAttribute"
	EntitySchemaUIConfigListItemSummaryAttributesTypeStr              EntitySchemaUIConfigListItemSummaryAttributesType = "str"
)

type EntitySchemaUIConfigListItemSummaryAttributes struct {
	SummaryAttribute *SummaryAttribute
	Str              *string

	Type EntitySchemaUIConfigListItemSummaryAttributesType
}

func CreateEntitySchemaUIConfigListItemSummaryAttributesSummaryAttribute(summaryAttribute SummaryAttribute) EntitySchemaUIConfigListItemSummaryAttributes {
	typ := EntitySchemaUIConfigListItemSummaryAttributesTypeSummaryAttribute

	return EntitySchemaUIConfigListItemSummaryAttributes{
		SummaryAttribute: &summaryAttribute,
		Type:             typ,
	}
}

func CreateEntitySchemaUIConfigListItemSummaryAttributesStr(str string) EntitySchemaUIConfigListItemSummaryAttributes {
	typ := EntitySchemaUIConfigListItemSummaryAttributesTypeStr

	return EntitySchemaUIConfigListItemSummaryAttributes{
		Str:  &str,
		Type: typ,
	}
}

func (u *EntitySchemaUIConfigListItemSummaryAttributes) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	summaryAttribute := new(SummaryAttribute)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&summaryAttribute); err == nil {
		u.SummaryAttribute = summaryAttribute
		u.Type = EntitySchemaUIConfigListItemSummaryAttributesTypeSummaryAttribute
		return nil
	}

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = EntitySchemaUIConfigListItemSummaryAttributesTypeStr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u EntitySchemaUIConfigListItemSummaryAttributes) MarshalJSON() ([]byte, error) {
	if u.SummaryAttribute != nil {
		return json.Marshal(u.SummaryAttribute)
	}

	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	return nil, nil
}

type EntitySchemaUIConfigListItem struct {
	SummaryAttributes []EntitySchemaUIConfigListItemSummaryAttributes `json:"summary_attributes,omitempty"`
}

type EntitySchemaUIConfigSharing struct {
	// Show the sharing button in entity detail view
	ShowSharingButton *bool `json:"show_sharing_button,omitempty"`
}

type EntitySchemaUIConfigSingleViewType string

const (
	EntitySchemaUIConfigSingleViewTypeEntityDefaultEdit  EntitySchemaUIConfigSingleViewType = "EntityDefaultEdit"
	EntitySchemaUIConfigSingleViewTypeRedirectEntityView EntitySchemaUIConfigSingleViewType = "RedirectEntityView"
	EntitySchemaUIConfigSingleViewTypeEntityViewDisabled EntitySchemaUIConfigSingleViewType = "EntityViewDisabled"
)

type EntitySchemaUIConfigSingleView struct {
	EntityDefaultEdit  *EntityDefaultEdit
	RedirectEntityView *RedirectEntityView
	EntityViewDisabled *EntityViewDisabled

	Type EntitySchemaUIConfigSingleViewType
}

func CreateEntitySchemaUIConfigSingleViewEntityDefaultEdit(entityDefaultEdit EntityDefaultEdit) EntitySchemaUIConfigSingleView {
	typ := EntitySchemaUIConfigSingleViewTypeEntityDefaultEdit

	return EntitySchemaUIConfigSingleView{
		EntityDefaultEdit: &entityDefaultEdit,
		Type:              typ,
	}
}

func CreateEntitySchemaUIConfigSingleViewRedirectEntityView(redirectEntityView RedirectEntityView) EntitySchemaUIConfigSingleView {
	typ := EntitySchemaUIConfigSingleViewTypeRedirectEntityView

	return EntitySchemaUIConfigSingleView{
		RedirectEntityView: &redirectEntityView,
		Type:               typ,
	}
}

func CreateEntitySchemaUIConfigSingleViewEntityViewDisabled(entityViewDisabled EntityViewDisabled) EntitySchemaUIConfigSingleView {
	typ := EntitySchemaUIConfigSingleViewTypeEntityViewDisabled

	return EntitySchemaUIConfigSingleView{
		EntityViewDisabled: &entityViewDisabled,
		Type:               typ,
	}
}

func (u *EntitySchemaUIConfigSingleView) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	entityDefaultEdit := new(EntityDefaultEdit)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&entityDefaultEdit); err == nil {
		u.EntityDefaultEdit = entityDefaultEdit
		u.Type = EntitySchemaUIConfigSingleViewTypeEntityDefaultEdit
		return nil
	}

	redirectEntityView := new(RedirectEntityView)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&redirectEntityView); err == nil {
		u.RedirectEntityView = redirectEntityView
		u.Type = EntitySchemaUIConfigSingleViewTypeRedirectEntityView
		return nil
	}

	entityViewDisabled := new(EntityViewDisabled)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&entityViewDisabled); err == nil {
		u.EntityViewDisabled = entityViewDisabled
		u.Type = EntitySchemaUIConfigSingleViewTypeEntityViewDisabled
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u EntitySchemaUIConfigSingleView) MarshalJSON() ([]byte, error) {
	if u.EntityDefaultEdit != nil {
		return json.Marshal(u.EntityDefaultEdit)
	}

	if u.RedirectEntityView != nil {
		return json.Marshal(u.RedirectEntityView)
	}

	if u.EntityViewDisabled != nil {
		return json.Marshal(u.EntityViewDisabled)
	}

	return nil, nil
}

type EntitySchemaUIConfigTableViewType string

const (
	EntitySchemaUIConfigTableViewTypeEntityDefaultTable EntitySchemaUIConfigTableViewType = "EntityDefaultTable"
	EntitySchemaUIConfigTableViewTypeRedirectEntityView EntitySchemaUIConfigTableViewType = "RedirectEntityView"
	EntitySchemaUIConfigTableViewTypeEntityViewDisabled EntitySchemaUIConfigTableViewType = "EntityViewDisabled"
)

type EntitySchemaUIConfigTableView struct {
	EntityDefaultTable *EntityDefaultTable
	RedirectEntityView *RedirectEntityView
	EntityViewDisabled *EntityViewDisabled

	Type EntitySchemaUIConfigTableViewType
}

func CreateEntitySchemaUIConfigTableViewEntityDefaultTable(entityDefaultTable EntityDefaultTable) EntitySchemaUIConfigTableView {
	typ := EntitySchemaUIConfigTableViewTypeEntityDefaultTable

	return EntitySchemaUIConfigTableView{
		EntityDefaultTable: &entityDefaultTable,
		Type:               typ,
	}
}

func CreateEntitySchemaUIConfigTableViewRedirectEntityView(redirectEntityView RedirectEntityView) EntitySchemaUIConfigTableView {
	typ := EntitySchemaUIConfigTableViewTypeRedirectEntityView

	return EntitySchemaUIConfigTableView{
		RedirectEntityView: &redirectEntityView,
		Type:               typ,
	}
}

func CreateEntitySchemaUIConfigTableViewEntityViewDisabled(entityViewDisabled EntityViewDisabled) EntitySchemaUIConfigTableView {
	typ := EntitySchemaUIConfigTableViewTypeEntityViewDisabled

	return EntitySchemaUIConfigTableView{
		EntityViewDisabled: &entityViewDisabled,
		Type:               typ,
	}
}

func (u *EntitySchemaUIConfigTableView) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	entityDefaultTable := new(EntityDefaultTable)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&entityDefaultTable); err == nil {
		u.EntityDefaultTable = entityDefaultTable
		u.Type = EntitySchemaUIConfigTableViewTypeEntityDefaultTable
		return nil
	}

	redirectEntityView := new(RedirectEntityView)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&redirectEntityView); err == nil {
		u.RedirectEntityView = redirectEntityView
		u.Type = EntitySchemaUIConfigTableViewTypeRedirectEntityView
		return nil
	}

	entityViewDisabled := new(EntityViewDisabled)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&entityViewDisabled); err == nil {
		u.EntityViewDisabled = entityViewDisabled
		u.Type = EntitySchemaUIConfigTableViewTypeEntityViewDisabled
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u EntitySchemaUIConfigTableView) MarshalJSON() ([]byte, error) {
	if u.EntityDefaultTable != nil {
		return json.Marshal(u.EntityDefaultTable)
	}

	if u.RedirectEntityView != nil {
		return json.Marshal(u.RedirectEntityView)
	}

	if u.EntityViewDisabled != nil {
		return json.Marshal(u.EntityViewDisabled)
	}

	return nil, nil
}

type EntitySchemaUIConfig struct {
	CreateView *EntitySchemaUIConfigCreateView `json:"create_view,omitempty"`
	EditView   *EntitySchemaUIConfigEditView   `json:"edit_view,omitempty"`
	ListItem   *EntitySchemaUIConfigListItem   `json:"list_item,omitempty"`
	Sharing    *EntitySchemaUIConfigSharing    `json:"sharing,omitempty"`
	SingleView *EntitySchemaUIConfigSingleView `json:"single_view,omitempty"`
	TableView  *EntitySchemaUIConfigTableView  `json:"table_view,omitempty"`
}

// EntitySchema - The "type" of an Entity. Describes the shape. Includes Entity Attributes, Relations and Capabilities.
type EntitySchema struct {
	// An ordered list of attributes the entity contains
	Attributes []Attribute `json:"attributes,omitempty"`
	// Reference to blueprint
	Blueprint    *string                `json:"blueprint,omitempty"`
	Capabilities []EntityCapability     `json:"capabilities,omitempty"`
	DialogConfig map[string]interface{} `json:"dialog_config,omitempty"`
	Draft        *bool                  `json:"draft,omitempty"`
	// This schema should only be active when one of the organization settings is enabled
	EnableSetting []string `json:"enable_setting,omitempty"`
	// Advanced: explicit Elasticsearch index mapping definitions for entity data
	//
	ExplicitSearchMappings map[string]SearchMappings `json:"explicit_search_mappings,omitempty"`
	// This schema should only be active when the feature flag is enabled
	FeatureFlag *string `json:"feature_flag,omitempty"`
	// A dictionary of Group Titles and associated settings if present.
	GroupSettings []EntitySchemaGroupSettings `json:"group_settings,omitempty"`
	Icon          *string                     `json:"icon,omitempty"`
	// Custom grid definitions for the layout. These settings are composed by managed and un-managed properties:
	// - Managed Properties: are interpreted and transformed into layout styles
	// - Un-managed Properties: are appended as styles into the attribute mounting node
	//
	LayoutSettings *EntitySchemaLayoutSettings `json:"layout_settings,omitempty"`
	// User-friendly identifier for the entity schema
	Name      string `json:"name"`
	Plural    string `json:"plural"`
	Published *bool  `json:"published,omitempty"`
	// URL-friendly identifier for the entity schema
	Slug string `json:"slug"`
	// Template for rendering the title field. Uses handlebars
	TitleTemplate *string               `json:"title_template,omitempty"`
	UIConfig      *EntitySchemaUIConfig `json:"ui_config,omitempty"`
	Version       *int64                `json:"version,omitempty"`
}
