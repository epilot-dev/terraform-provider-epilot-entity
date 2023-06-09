// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SearchMappingsType string

const (
	SearchMappingsTypeKeyword   SearchMappingsType = "keyword"
	SearchMappingsTypeText      SearchMappingsType = "text"
	SearchMappingsTypeBoolean   SearchMappingsType = "boolean"
	SearchMappingsTypeInteger   SearchMappingsType = "integer"
	SearchMappingsTypeLong      SearchMappingsType = "long"
	SearchMappingsTypeFloat     SearchMappingsType = "float"
	SearchMappingsTypeDate      SearchMappingsType = "date"
	SearchMappingsTypeFlattened SearchMappingsType = "flattened"
	SearchMappingsTypeNested    SearchMappingsType = "nested"
)

func (e SearchMappingsType) ToPointer() *SearchMappingsType {
	return &e
}

func (e *SearchMappingsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "keyword":
		fallthrough
	case "text":
		fallthrough
	case "boolean":
		fallthrough
	case "integer":
		fallthrough
	case "long":
		fallthrough
	case "float":
		fallthrough
	case "date":
		fallthrough
	case "flattened":
		fallthrough
	case "nested":
		*e = SearchMappingsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SearchMappingsType: %v", v)
	}
}

type SearchMappings struct {
	Fields map[string]interface{} `json:"fields,omitempty"`
	Index  *bool                  `json:"index,omitempty"`
	Type   *SearchMappingsType    `json:"type,omitempty"`
}
