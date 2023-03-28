// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"epilotentity/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetRelationsV2Request struct {
	// List of entity fields to include in results
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// Starting page number
	From *int64 `queryParam:"style=form,explode=true,name=from"`
	// When true, expand relation items with full blown entities.
	Hydrate *bool `queryParam:"style=form,explode=true,name=hydrate"`
	// Entity id
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// When true, includes reverse relations in response (other entities pointing to this entity)
	IncludeReverse *bool `queryParam:"style=form,explode=true,name=include_reverse"`
	// Input to filter search results
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// Number of results to return per page
	Size *int64 `queryParam:"style=form,explode=true,name=size"`
	// Entity Type
	Slug string `pathParam:"style=simple,explode=false,name=slug"`
}

type GetRelationsV2Response struct {
	ContentType string
	// Success
	GetRelationsRespWithPagination *shared.GetRelationsRespWithPagination
	StatusCode                     int
	RawResponse                    *http.Response
}
