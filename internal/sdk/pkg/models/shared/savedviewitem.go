// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SavedViewItemCreatedBy2SourceEnum string

const (
	SavedViewItemCreatedBy2SourceEnumSystem    SavedViewItemCreatedBy2SourceEnum = "SYSTEM"
	SavedViewItemCreatedBy2SourceEnumBlueprint SavedViewItemCreatedBy2SourceEnum = "BLUEPRINT"
)

func (e SavedViewItemCreatedBy2SourceEnum) ToPointer() *SavedViewItemCreatedBy2SourceEnum {
	return &e
}

func (e *SavedViewItemCreatedBy2SourceEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SYSTEM":
		fallthrough
	case "BLUEPRINT":
		*e = SavedViewItemCreatedBy2SourceEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SavedViewItemCreatedBy2SourceEnum: %v", v)
	}
}

// SavedViewItemCreatedBy2 - A system-created view
type SavedViewItemCreatedBy2 struct {
	Source *SavedViewItemCreatedBy2SourceEnum `json:"source,omitempty"`

	AdditionalProperties map[string]interface{} `json:"-"`
}
type _SavedViewItemCreatedBy2 SavedViewItemCreatedBy2

func (c *SavedViewItemCreatedBy2) UnmarshalJSON(bs []byte) error {
	data := _SavedViewItemCreatedBy2{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = SavedViewItemCreatedBy2(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "source")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c SavedViewItemCreatedBy2) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_SavedViewItemCreatedBy2(c))
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

// SavedViewItemCreatedBy1 - A user that created the view
type SavedViewItemCreatedBy1 struct {
	UserID *string `json:"user_id,omitempty"`
}

// SavedViewItem - A saved entity view
type SavedViewItem struct {
	CreatedAt *string     `json:"created_at,omitempty"`
	CreatedBy interface{} `json:"created_by"`
	// Generated uuid for a saved view
	ID *string `json:"id,omitempty"`
	// User-friendly identifier for the saved view
	Name string `json:"name"`
	// Organisation ID a view belongs to
	Org *string `json:"org,omitempty"`
	// boolean property for if a view is shared with organisation
	Shared *bool `json:"shared,omitempty"`
	// list of schemas a view can belong to
	Slug      []string               `json:"slug"`
	UIConfig  map[string]interface{} `json:"ui_config"`
	UpdatedAt *string                `json:"updated_at,omitempty"`
}
