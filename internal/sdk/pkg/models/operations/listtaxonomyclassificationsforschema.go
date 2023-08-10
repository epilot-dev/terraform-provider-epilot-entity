// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"epilot-entity/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListTaxonomyClassificationsForSchemaRequest struct {
	Query        *string  `queryParam:"style=form,explode=true,name=query"`
	Size         *float64 `queryParam:"style=form,explode=true,name=size"`
	Slug         string   `pathParam:"style=simple,explode=false,name=slug"`
	TaxonomySlug string   `pathParam:"style=simple,explode=false,name=taxonomySlug"`
}

// ListTaxonomyClassificationsForSchema200ApplicationJSON - List of taxonomy classifications
type ListTaxonomyClassificationsForSchema200ApplicationJSON struct {
	Results []shared.TaxonomyClassification `json:"results,omitempty"`
}

type ListTaxonomyClassificationsForSchemaResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// List of taxonomy classifications
	ListTaxonomyClassificationsForSchema200ApplicationJSONObject *ListTaxonomyClassificationsForSchema200ApplicationJSON
}
