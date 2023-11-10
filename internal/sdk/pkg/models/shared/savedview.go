// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-entity/internal/sdk/pkg/utils"
)

type SavedViewSource string

const (
	SavedViewSourceSystem    SavedViewSource = "SYSTEM"
	SavedViewSourceBlueprint SavedViewSource = "BLUEPRINT"
)

func (e SavedViewSource) ToPointer() *SavedViewSource {
	return &e
}

func (e *SavedViewSource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SYSTEM":
		fallthrough
	case "BLUEPRINT":
		*e = SavedViewSource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SavedViewSource: %v", v)
	}
}

// SavedView2 - A system-created view
type SavedView2 struct {
	AdditionalProperties interface{}      `additionalProperties:"true" json:"-"`
	Source               *SavedViewSource `json:"source,omitempty"`
}

func (s SavedView2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SavedView2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SavedView2) GetAdditionalProperties() interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *SavedView2) GetSource() *SavedViewSource {
	if o == nil {
		return nil
	}
	return o.Source
}

// SavedView1 - A user that created the view
type SavedView1 struct {
	UserID *string `json:"user_id,omitempty"`
}

func (o *SavedView1) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

type CreatedByType string

const (
	CreatedByTypeSavedView1 CreatedByType = "SavedView_1"
	CreatedByTypeSavedView2 CreatedByType = "SavedView_2"
)

type CreatedBy struct {
	SavedView1 *SavedView1
	SavedView2 *SavedView2

	Type CreatedByType
}

func CreateCreatedBySavedView1(savedView1 SavedView1) CreatedBy {
	typ := CreatedByTypeSavedView1

	return CreatedBy{
		SavedView1: &savedView1,
		Type:       typ,
	}
}

func CreateCreatedBySavedView2(savedView2 SavedView2) CreatedBy {
	typ := CreatedByTypeSavedView2

	return CreatedBy{
		SavedView2: &savedView2,
		Type:       typ,
	}
}

func (u *CreatedBy) UnmarshalJSON(data []byte) error {

	savedView1 := new(SavedView1)
	if err := utils.UnmarshalJSON(data, &savedView1, "", true, true); err == nil {
		u.SavedView1 = savedView1
		u.Type = CreatedByTypeSavedView1
		return nil
	}

	savedView2 := new(SavedView2)
	if err := utils.UnmarshalJSON(data, &savedView2, "", true, true); err == nil {
		u.SavedView2 = savedView2
		u.Type = CreatedByTypeSavedView2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreatedBy) MarshalJSON() ([]byte, error) {
	if u.SavedView1 != nil {
		return utils.MarshalJSON(u.SavedView1, "", true)
	}

	if u.SavedView2 != nil {
		return utils.MarshalJSON(u.SavedView2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// SavedView - A saved entity view
type SavedView struct {
	CreatedBy CreatedBy `json:"created_by"`
	// User-friendly identifier for the saved view
	Name string `json:"name"`
	// Organisation ID a view belongs to
	Org *string `json:"org,omitempty"`
	// boolean property for if a view is shared with organisation
	Shared *bool `json:"shared,omitempty"`
	// list of schemas a view can belong to
	Slug     []string               `json:"slug"`
	UIConfig map[string]interface{} `json:"ui_config"`
}

func (o *SavedView) GetCreatedBy() CreatedBy {
	if o == nil {
		return CreatedBy{}
	}
	return o.CreatedBy
}

func (o *SavedView) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SavedView) GetOrg() *string {
	if o == nil {
		return nil
	}
	return o.Org
}

func (o *SavedView) GetShared() *bool {
	if o == nil {
		return nil
	}
	return o.Shared
}

func (o *SavedView) GetSlug() []string {
	if o == nil {
		return []string{}
	}
	return o.Slug
}

func (o *SavedView) GetUIConfig() map[string]interface{} {
	if o == nil {
		return map[string]interface{}{}
	}
	return o.UIConfig
}
