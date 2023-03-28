// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"epilotentity/internal/sdk/pkg/models/shared"
	"net/http"
)

type TaxonomyAutocompleteRequest struct {
	// Input to autocomplete
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// Minimum number of results to return
	Size *float64 `queryParam:"style=form,explode=true,name=size"`
	// Limit results to slug
	TaxonomySlug string `pathParam:"style=simple,explode=false,name=taxonomySlug"`
}

// TaxonomyAutocomplete200ApplicationJSON - Taxonomy classifications
type TaxonomyAutocomplete200ApplicationJSON struct {
	Results []shared.TaxonomyClassification `json:"results,omitempty"`
}

type TaxonomyAutocompleteResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Taxonomy classifications
	TaxonomyAutocomplete200ApplicationJSONObject *TaxonomyAutocomplete200ApplicationJSON
}
