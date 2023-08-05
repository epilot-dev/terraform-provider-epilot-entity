// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
)

type EntitySchemaItemGroupSettingsInfoTooltipTitle struct {
	Default *string `json:"default,omitempty"`
	Key     *string `json:"key,omitempty"`
}

type EntitySchemaItemGroupSettings struct {
	Purpose  []string `json:"_purpose,omitempty"`
	Expanded *bool    `json:"expanded,omitempty"`
	// This group should only be active when the feature flag is enabled
	FeatureFlag      *string                                        `json:"feature_flag,omitempty"`
	ID               string                                         `json:"id"`
	InfoTooltipTitle *EntitySchemaItemGroupSettingsInfoTooltipTitle `json:"info_tooltip_title,omitempty"`
	Label            string                                         `json:"label"`
	// Render order of the group
	Order           *int64  `json:"order,omitempty"`
	RenderCondition *string `json:"render_condition,omitempty"`
	// This group should only be active when the setting is enabled
	SettingFlag *string `json:"setting_flag,omitempty"`
}

// EntitySchemaItemLayoutSettings - Custom grid definitions for the layout. These settings are composed by managed and un-managed properties:
// - Managed Properties: are interpreted and transformed into layout styles
// - Un-managed Properties: are appended as styles into the attribute mounting node
type EntitySchemaItemLayoutSettings struct {
	// Defines the grid gap for the mounting node of the attribute.
	GridGap *string `json:"grid_gap,omitempty"`
	// Defines the grid column template for the mounting node of the attribute.
	GridTemplateColumns *string `json:"grid_template_columns,omitempty"`

	AdditionalProperties interface{} `json:"-"`
}
type _EntitySchemaItemLayoutSettings EntitySchemaItemLayoutSettings

func (c *EntitySchemaItemLayoutSettings) UnmarshalJSON(bs []byte) error {
	data := _EntitySchemaItemLayoutSettings{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = EntitySchemaItemLayoutSettings(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "grid_gap")
	delete(additionalFields, "grid_template_columns")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c EntitySchemaItemLayoutSettings) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_EntitySchemaItemLayoutSettings(c))
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

type EntitySchemaItemSource struct {
	ID   *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
}

type EntitySchemaItemUIConfigCreateViewType string

const (
	EntitySchemaItemUIConfigCreateViewTypeEntityDefaultCreate EntitySchemaItemUIConfigCreateViewType = "EntityDefaultCreate"
	EntitySchemaItemUIConfigCreateViewTypeRedirectEntityView  EntitySchemaItemUIConfigCreateViewType = "RedirectEntityView"
	EntitySchemaItemUIConfigCreateViewTypeEntityViewDisabled  EntitySchemaItemUIConfigCreateViewType = "EntityViewDisabled"
)

type EntitySchemaItemUIConfigCreateView struct {
	EntityDefaultCreate *EntityDefaultCreate
	RedirectEntityView  *RedirectEntityView
	EntityViewDisabled  *EntityViewDisabled

	Type EntitySchemaItemUIConfigCreateViewType
}

func CreateEntitySchemaItemUIConfigCreateViewEntityDefaultCreate(entityDefaultCreate EntityDefaultCreate) EntitySchemaItemUIConfigCreateView {
	typ := EntitySchemaItemUIConfigCreateViewTypeEntityDefaultCreate

	return EntitySchemaItemUIConfigCreateView{
		EntityDefaultCreate: &entityDefaultCreate,
		Type:                typ,
	}
}

func CreateEntitySchemaItemUIConfigCreateViewRedirectEntityView(redirectEntityView RedirectEntityView) EntitySchemaItemUIConfigCreateView {
	typ := EntitySchemaItemUIConfigCreateViewTypeRedirectEntityView

	return EntitySchemaItemUIConfigCreateView{
		RedirectEntityView: &redirectEntityView,
		Type:               typ,
	}
}

func CreateEntitySchemaItemUIConfigCreateViewEntityViewDisabled(entityViewDisabled EntityViewDisabled) EntitySchemaItemUIConfigCreateView {
	typ := EntitySchemaItemUIConfigCreateViewTypeEntityViewDisabled

	return EntitySchemaItemUIConfigCreateView{
		EntityViewDisabled: &entityViewDisabled,
		Type:               typ,
	}
}

func (u *EntitySchemaItemUIConfigCreateView) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	entityDefaultCreate := new(EntityDefaultCreate)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&entityDefaultCreate); err == nil {
		u.EntityDefaultCreate = entityDefaultCreate
		u.Type = EntitySchemaItemUIConfigCreateViewTypeEntityDefaultCreate
		return nil
	}

	redirectEntityView := new(RedirectEntityView)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&redirectEntityView); err == nil {
		u.RedirectEntityView = redirectEntityView
		u.Type = EntitySchemaItemUIConfigCreateViewTypeRedirectEntityView
		return nil
	}

	entityViewDisabled := new(EntityViewDisabled)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&entityViewDisabled); err == nil {
		u.EntityViewDisabled = entityViewDisabled
		u.Type = EntitySchemaItemUIConfigCreateViewTypeEntityViewDisabled
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u EntitySchemaItemUIConfigCreateView) MarshalJSON() ([]byte, error) {
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

type EntitySchemaItemUIConfigEditViewType string

const (
	EntitySchemaItemUIConfigEditViewTypeEntityDefaultEdit  EntitySchemaItemUIConfigEditViewType = "EntityDefaultEdit"
	EntitySchemaItemUIConfigEditViewTypeRedirectEntityView EntitySchemaItemUIConfigEditViewType = "RedirectEntityView"
	EntitySchemaItemUIConfigEditViewTypeEntityViewDisabled EntitySchemaItemUIConfigEditViewType = "EntityViewDisabled"
)

type EntitySchemaItemUIConfigEditView struct {
	EntityDefaultEdit  *EntityDefaultEdit
	RedirectEntityView *RedirectEntityView
	EntityViewDisabled *EntityViewDisabled

	Type EntitySchemaItemUIConfigEditViewType
}

func CreateEntitySchemaItemUIConfigEditViewEntityDefaultEdit(entityDefaultEdit EntityDefaultEdit) EntitySchemaItemUIConfigEditView {
	typ := EntitySchemaItemUIConfigEditViewTypeEntityDefaultEdit

	return EntitySchemaItemUIConfigEditView{
		EntityDefaultEdit: &entityDefaultEdit,
		Type:              typ,
	}
}

func CreateEntitySchemaItemUIConfigEditViewRedirectEntityView(redirectEntityView RedirectEntityView) EntitySchemaItemUIConfigEditView {
	typ := EntitySchemaItemUIConfigEditViewTypeRedirectEntityView

	return EntitySchemaItemUIConfigEditView{
		RedirectEntityView: &redirectEntityView,
		Type:               typ,
	}
}

func CreateEntitySchemaItemUIConfigEditViewEntityViewDisabled(entityViewDisabled EntityViewDisabled) EntitySchemaItemUIConfigEditView {
	typ := EntitySchemaItemUIConfigEditViewTypeEntityViewDisabled

	return EntitySchemaItemUIConfigEditView{
		EntityViewDisabled: &entityViewDisabled,
		Type:               typ,
	}
}

func (u *EntitySchemaItemUIConfigEditView) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	entityDefaultEdit := new(EntityDefaultEdit)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&entityDefaultEdit); err == nil {
		u.EntityDefaultEdit = entityDefaultEdit
		u.Type = EntitySchemaItemUIConfigEditViewTypeEntityDefaultEdit
		return nil
	}

	redirectEntityView := new(RedirectEntityView)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&redirectEntityView); err == nil {
		u.RedirectEntityView = redirectEntityView
		u.Type = EntitySchemaItemUIConfigEditViewTypeRedirectEntityView
		return nil
	}

	entityViewDisabled := new(EntityViewDisabled)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&entityViewDisabled); err == nil {
		u.EntityViewDisabled = entityViewDisabled
		u.Type = EntitySchemaItemUIConfigEditViewTypeEntityViewDisabled
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u EntitySchemaItemUIConfigEditView) MarshalJSON() ([]byte, error) {
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

type EntitySchemaItemUIConfigListItem struct {
	SummaryAttributes []SummaryAttribute `json:"summary_attributes,omitempty"`
}

type EntitySchemaItemUIConfigSharing struct {
	// Show the sharing button in entity detail view
	ShowSharingButton *bool `json:"show_sharing_button,omitempty"`
}

type EntitySchemaItemUIConfigSingleViewType string

const (
	EntitySchemaItemUIConfigSingleViewTypeEntityDefaultEdit  EntitySchemaItemUIConfigSingleViewType = "EntityDefaultEdit"
	EntitySchemaItemUIConfigSingleViewTypeRedirectEntityView EntitySchemaItemUIConfigSingleViewType = "RedirectEntityView"
	EntitySchemaItemUIConfigSingleViewTypeEntityViewDisabled EntitySchemaItemUIConfigSingleViewType = "EntityViewDisabled"
)

type EntitySchemaItemUIConfigSingleView struct {
	EntityDefaultEdit  *EntityDefaultEdit
	RedirectEntityView *RedirectEntityView
	EntityViewDisabled *EntityViewDisabled

	Type EntitySchemaItemUIConfigSingleViewType
}

func CreateEntitySchemaItemUIConfigSingleViewEntityDefaultEdit(entityDefaultEdit EntityDefaultEdit) EntitySchemaItemUIConfigSingleView {
	typ := EntitySchemaItemUIConfigSingleViewTypeEntityDefaultEdit

	return EntitySchemaItemUIConfigSingleView{
		EntityDefaultEdit: &entityDefaultEdit,
		Type:              typ,
	}
}

func CreateEntitySchemaItemUIConfigSingleViewRedirectEntityView(redirectEntityView RedirectEntityView) EntitySchemaItemUIConfigSingleView {
	typ := EntitySchemaItemUIConfigSingleViewTypeRedirectEntityView

	return EntitySchemaItemUIConfigSingleView{
		RedirectEntityView: &redirectEntityView,
		Type:               typ,
	}
}

func CreateEntitySchemaItemUIConfigSingleViewEntityViewDisabled(entityViewDisabled EntityViewDisabled) EntitySchemaItemUIConfigSingleView {
	typ := EntitySchemaItemUIConfigSingleViewTypeEntityViewDisabled

	return EntitySchemaItemUIConfigSingleView{
		EntityViewDisabled: &entityViewDisabled,
		Type:               typ,
	}
}

func (u *EntitySchemaItemUIConfigSingleView) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	entityDefaultEdit := new(EntityDefaultEdit)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&entityDefaultEdit); err == nil {
		u.EntityDefaultEdit = entityDefaultEdit
		u.Type = EntitySchemaItemUIConfigSingleViewTypeEntityDefaultEdit
		return nil
	}

	redirectEntityView := new(RedirectEntityView)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&redirectEntityView); err == nil {
		u.RedirectEntityView = redirectEntityView
		u.Type = EntitySchemaItemUIConfigSingleViewTypeRedirectEntityView
		return nil
	}

	entityViewDisabled := new(EntityViewDisabled)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&entityViewDisabled); err == nil {
		u.EntityViewDisabled = entityViewDisabled
		u.Type = EntitySchemaItemUIConfigSingleViewTypeEntityViewDisabled
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u EntitySchemaItemUIConfigSingleView) MarshalJSON() ([]byte, error) {
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

type EntitySchemaItemUIConfigTableViewType string

const (
	EntitySchemaItemUIConfigTableViewTypeEntityDefaultTable EntitySchemaItemUIConfigTableViewType = "EntityDefaultTable"
	EntitySchemaItemUIConfigTableViewTypeRedirectEntityView EntitySchemaItemUIConfigTableViewType = "RedirectEntityView"
	EntitySchemaItemUIConfigTableViewTypeEntityViewDisabled EntitySchemaItemUIConfigTableViewType = "EntityViewDisabled"
)

type EntitySchemaItemUIConfigTableView struct {
	EntityDefaultTable *EntityDefaultTable
	RedirectEntityView *RedirectEntityView
	EntityViewDisabled *EntityViewDisabled

	Type EntitySchemaItemUIConfigTableViewType
}

func CreateEntitySchemaItemUIConfigTableViewEntityDefaultTable(entityDefaultTable EntityDefaultTable) EntitySchemaItemUIConfigTableView {
	typ := EntitySchemaItemUIConfigTableViewTypeEntityDefaultTable

	return EntitySchemaItemUIConfigTableView{
		EntityDefaultTable: &entityDefaultTable,
		Type:               typ,
	}
}

func CreateEntitySchemaItemUIConfigTableViewRedirectEntityView(redirectEntityView RedirectEntityView) EntitySchemaItemUIConfigTableView {
	typ := EntitySchemaItemUIConfigTableViewTypeRedirectEntityView

	return EntitySchemaItemUIConfigTableView{
		RedirectEntityView: &redirectEntityView,
		Type:               typ,
	}
}

func CreateEntitySchemaItemUIConfigTableViewEntityViewDisabled(entityViewDisabled EntityViewDisabled) EntitySchemaItemUIConfigTableView {
	typ := EntitySchemaItemUIConfigTableViewTypeEntityViewDisabled

	return EntitySchemaItemUIConfigTableView{
		EntityViewDisabled: &entityViewDisabled,
		Type:               typ,
	}
}

func (u *EntitySchemaItemUIConfigTableView) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	entityDefaultTable := new(EntityDefaultTable)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&entityDefaultTable); err == nil {
		u.EntityDefaultTable = entityDefaultTable
		u.Type = EntitySchemaItemUIConfigTableViewTypeEntityDefaultTable
		return nil
	}

	redirectEntityView := new(RedirectEntityView)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&redirectEntityView); err == nil {
		u.RedirectEntityView = redirectEntityView
		u.Type = EntitySchemaItemUIConfigTableViewTypeRedirectEntityView
		return nil
	}

	entityViewDisabled := new(EntityViewDisabled)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&entityViewDisabled); err == nil {
		u.EntityViewDisabled = entityViewDisabled
		u.Type = EntitySchemaItemUIConfigTableViewTypeEntityViewDisabled
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u EntitySchemaItemUIConfigTableView) MarshalJSON() ([]byte, error) {
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

type EntitySchemaItemUIConfig struct {
	CreateView *EntitySchemaItemUIConfigCreateView `json:"create_view,omitempty"`
	EditView   *EntitySchemaItemUIConfigEditView   `json:"edit_view,omitempty"`
	ListItem   *EntitySchemaItemUIConfigListItem   `json:"list_item,omitempty"`
	Sharing    *EntitySchemaItemUIConfigSharing    `json:"sharing,omitempty"`
	SingleView *EntitySchemaItemUIConfigSingleView `json:"single_view,omitempty"`
	TableView  *EntitySchemaItemUIConfigTableView  `json:"table_view,omitempty"`
}

// EntitySchemaItem - The "type" of an Entity. Describes the shape. Includes Entity Attributes, Relations and Capabilities.
type EntitySchemaItem struct {
	// An ordered list of attributes the entity contains
	Attributes []Attribute `json:"attributes,omitempty"`
	// Reference to blueprint
	Blueprint    *string                `json:"blueprint,omitempty"`
	Capabilities []EntityCapability     `json:"capabilities,omitempty"`
	Comment      *string                `json:"comment,omitempty"`
	CreatedAt    *string                `json:"created_at,omitempty"`
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
	GroupSettings []EntitySchemaItemGroupSettings `json:"group_settings,omitempty"`
	Icon          *string                         `json:"icon,omitempty"`
	// Generated uuid for schema
	ID *string `json:"id,omitempty"`
	// Custom grid definitions for the layout. These settings are composed by managed and un-managed properties:
	// - Managed Properties: are interpreted and transformed into layout styles
	// - Un-managed Properties: are appended as styles into the attribute mounting node
	//
	LayoutSettings *EntitySchemaItemLayoutSettings `json:"layout_settings,omitempty"`
	// User-friendly identifier for the entity schema
	Name      string `json:"name"`
	Plural    string `json:"plural"`
	Published *bool  `json:"published,omitempty"`
	// URL-friendly identifier for the entity schema
	Slug   string                  `json:"slug"`
	Source *EntitySchemaItemSource `json:"source,omitempty"`
	// Template for rendering the title field. Uses handlebars
	TitleTemplate *string                   `json:"title_template,omitempty"`
	UIConfig      *EntitySchemaItemUIConfig `json:"ui_config,omitempty"`
	UpdatedAt     *string                   `json:"updated_at,omitempty"`
	Version       *int64                    `json:"version,omitempty"`
}
