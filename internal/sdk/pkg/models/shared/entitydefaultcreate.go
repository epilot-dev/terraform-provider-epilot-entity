// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type EntityDefaultCreateTableMenuOptions struct {
	Icon  *string `json:"icon,omitempty"`
	Label *string `json:"label,omitempty"`
}

type EntityDefaultCreateViewTypeEnum string

const (
	EntityDefaultCreateViewTypeEnumDefault EntityDefaultCreateViewTypeEnum = "default"
)

func (e EntityDefaultCreateViewTypeEnum) ToPointer() *EntityDefaultCreateViewTypeEnum {
	return &e
}

func (e *EntityDefaultCreateViewTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "default":
		*e = EntityDefaultCreateViewTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EntityDefaultCreateViewTypeEnum: %v", v)
	}
}

type EntityDefaultCreate struct {
	SearchParams     map[string]string                    `json:"search_params,omitempty"`
	TableMenuOptions *EntityDefaultCreateTableMenuOptions `json:"table_menu_options,omitempty"`
	ViewType         *EntityDefaultCreateViewTypeEnum     `json:"view_type,omitempty"`
}
