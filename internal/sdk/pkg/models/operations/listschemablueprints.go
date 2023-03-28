// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"epilotentity/internal/sdk/pkg/models/shared"
	"net/http"
)

// ListSchemaBlueprints200ApplicationJSON - Success
type ListSchemaBlueprints200ApplicationJSON struct {
	Results []shared.EntitySchemaItem `json:"results,omitempty"`
}

type ListSchemaBlueprintsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Success
	ListSchemaBlueprints200ApplicationJSONObject *ListSchemaBlueprints200ApplicationJSON
}
