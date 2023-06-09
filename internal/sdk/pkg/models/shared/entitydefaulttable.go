// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type EntityDefaultTableDropdownItems2Type string

const (
	EntityDefaultTableDropdownItems2TypeLink EntityDefaultTableDropdownItems2Type = "link"
)

func (e EntityDefaultTableDropdownItems2Type) ToPointer() *EntityDefaultTableDropdownItems2Type {
	return &e
}

func (e *EntityDefaultTableDropdownItems2Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "link":
		*e = EntityDefaultTableDropdownItems2Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EntityDefaultTableDropdownItems2Type: %v", v)
	}
}

type EntityDefaultTableDropdownItems2 struct {
	// This dropdown item should only be active when the feature flag is enabled
	FeatureFlag *string `json:"feature_flag,omitempty"`
	// Only show item for legacy tenants (ivy)
	Legacy *bool                                 `json:"legacy,omitempty"`
	Title  *string                               `json:"title,omitempty"`
	Type   *EntityDefaultTableDropdownItems2Type `json:"type,omitempty"`
	URI    *string                               `json:"uri,omitempty"`
}

type EntityDefaultTableDropdownItems1Type string

const (
	EntityDefaultTableDropdownItems1TypeEntity EntityDefaultTableDropdownItems1Type = "entity"
)

func (e EntityDefaultTableDropdownItems1Type) ToPointer() *EntityDefaultTableDropdownItems1Type {
	return &e
}

func (e *EntityDefaultTableDropdownItems1Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "entity":
		*e = EntityDefaultTableDropdownItems1Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EntityDefaultTableDropdownItems1Type: %v", v)
	}
}

type EntityDefaultTableDropdownItems1 struct {
	// URL-friendly identifier for the entity schema
	Entity *string `json:"entity,omitempty"`
	// This dropdown item should only be active when the feature flag is enabled
	FeatureFlag *string `json:"feature_flag,omitempty"`
	// Only show item for legacy tenants (ivy)
	Legacy *bool                                 `json:"legacy,omitempty"`
	Type   *EntityDefaultTableDropdownItems1Type `json:"type,omitempty"`
}

type EntityDefaultTableDropdownItemsType string

const (
	EntityDefaultTableDropdownItemsTypeEntityDefaultTableDropdownItems1 EntityDefaultTableDropdownItemsType = "EntityDefaultTable_dropdown_items_1"
	EntityDefaultTableDropdownItemsTypeEntityDefaultTableDropdownItems2 EntityDefaultTableDropdownItemsType = "EntityDefaultTable_dropdown_items_2"
)

type EntityDefaultTableDropdownItems struct {
	EntityDefaultTableDropdownItems1 *EntityDefaultTableDropdownItems1
	EntityDefaultTableDropdownItems2 *EntityDefaultTableDropdownItems2

	Type EntityDefaultTableDropdownItemsType
}

func CreateEntityDefaultTableDropdownItemsEntityDefaultTableDropdownItems1(entityDefaultTableDropdownItems1 EntityDefaultTableDropdownItems1) EntityDefaultTableDropdownItems {
	typ := EntityDefaultTableDropdownItemsTypeEntityDefaultTableDropdownItems1

	return EntityDefaultTableDropdownItems{
		EntityDefaultTableDropdownItems1: &entityDefaultTableDropdownItems1,
		Type:                             typ,
	}
}

func CreateEntityDefaultTableDropdownItemsEntityDefaultTableDropdownItems2(entityDefaultTableDropdownItems2 EntityDefaultTableDropdownItems2) EntityDefaultTableDropdownItems {
	typ := EntityDefaultTableDropdownItemsTypeEntityDefaultTableDropdownItems2

	return EntityDefaultTableDropdownItems{
		EntityDefaultTableDropdownItems2: &entityDefaultTableDropdownItems2,
		Type:                             typ,
	}
}

func (u *EntityDefaultTableDropdownItems) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	entityDefaultTableDropdownItems1 := new(EntityDefaultTableDropdownItems1)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&entityDefaultTableDropdownItems1); err == nil {
		u.EntityDefaultTableDropdownItems1 = entityDefaultTableDropdownItems1
		u.Type = EntityDefaultTableDropdownItemsTypeEntityDefaultTableDropdownItems1
		return nil
	}

	entityDefaultTableDropdownItems2 := new(EntityDefaultTableDropdownItems2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&entityDefaultTableDropdownItems2); err == nil {
		u.EntityDefaultTableDropdownItems2 = entityDefaultTableDropdownItems2
		u.Type = EntityDefaultTableDropdownItemsTypeEntityDefaultTableDropdownItems2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u EntityDefaultTableDropdownItems) MarshalJSON() ([]byte, error) {
	if u.EntityDefaultTableDropdownItems1 != nil {
		return json.Marshal(u.EntityDefaultTableDropdownItems1)
	}

	if u.EntityDefaultTableDropdownItems2 != nil {
		return json.Marshal(u.EntityDefaultTableDropdownItems2)
	}

	return nil, nil
}

type EntityDefaultTableNavbarActionsOptions struct {
	Label  string                 `json:"label"`
	Params map[string]interface{} `json:"params,omitempty"`
}

type EntityDefaultTableNavbarActions struct {
	Options []EntityDefaultTableNavbarActionsOptions `json:"options,omitempty"`
	Type    string                                   `json:"type"`
}

type EntityDefaultTableViewType string

const (
	EntityDefaultTableViewTypeDefault EntityDefaultTableViewType = "default"
)

func (e EntityDefaultTableViewType) ToPointer() *EntityDefaultTableViewType {
	return &e
}

func (e *EntityDefaultTableViewType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "default":
		*e = EntityDefaultTableViewType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EntityDefaultTableViewType: %v", v)
	}
}

type EntityDefaultTable struct {
	ClassicView   *string                           `json:"classic_view,omitempty"`
	DropdownItems []EntityDefaultTableDropdownItems `json:"dropdown_items,omitempty"`
	// Enable the thumbnail column
	EnableThumbnails *bool                             `json:"enable_thumbnails,omitempty"`
	NavbarActions    []EntityDefaultTableNavbarActions `json:"navbar_actions,omitempty"`
	RowActions       []string                          `json:"row_actions,omitempty"`
	ViewType         *EntityDefaultTableViewType       `json:"view_type,omitempty"`
}
