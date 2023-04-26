// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SavedViewCreatedBy2SourceEnum string

const (
	SavedViewCreatedBy2SourceEnumSystem    SavedViewCreatedBy2SourceEnum = "SYSTEM"
	SavedViewCreatedBy2SourceEnumBlueprint SavedViewCreatedBy2SourceEnum = "BLUEPRINT"
)

func (e *SavedViewCreatedBy2SourceEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "SYSTEM":
		fallthrough
	case "BLUEPRINT":
		*e = SavedViewCreatedBy2SourceEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SavedViewCreatedBy2SourceEnum: %s", s)
	}
}

// SavedViewCreatedBy2 - A system-created view
type SavedViewCreatedBy2 struct {
	Source *SavedViewCreatedBy2SourceEnum `json:"source,omitempty"`

	AdditionalProperties map[string]interface{} `json:"-"`
}
type _SavedViewCreatedBy2 SavedViewCreatedBy2

func (c *SavedViewCreatedBy2) UnmarshalJSON(bs []byte) error {
	data := _SavedViewCreatedBy2{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = SavedViewCreatedBy2(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "source")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c SavedViewCreatedBy2) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_SavedViewCreatedBy2(c))
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

// SavedViewCreatedBy1 - A user that created the view
type SavedViewCreatedBy1 struct {
	UserID *string `json:"user_id,omitempty"`
}

// SavedView - A saved entity view
type SavedView struct {
	CreatedBy interface{} `json:"created_by"`
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
