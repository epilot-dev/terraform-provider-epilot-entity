// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"epilotentity/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetSchemaVersionsRequest struct {
	Slug string `pathParam:"style=simple,explode=false,name=slug"`
}

// GetSchemaVersions200ApplicationJSON - Success
type GetSchemaVersions200ApplicationJSON struct {
	Drafts   []shared.EntitySchemaItem `json:"drafts,omitempty"`
	Versions []shared.EntitySchemaItem `json:"versions,omitempty"`
}

type GetSchemaVersionsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Success
	GetSchemaVersions200ApplicationJSONObject *GetSchemaVersions200ApplicationJSON
}
