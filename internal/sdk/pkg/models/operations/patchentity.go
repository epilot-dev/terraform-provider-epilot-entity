// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PatchEntityPathParams struct {
	// Entity Id
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Entity Schema
	Slug string `pathParam:"style=simple,explode=false,name=slug"`
}

type PatchEntityQueryParams struct {
	// Activity to include in event feed
	ActivityID *string `queryParam:"style=form,explode=true,name=activity_id"`
	// Don't wait for the patch entity to become available in Search API. Useful for large migrations
	Async *bool `queryParam:"style=form,explode=true,name=async"`
	// Dry Run mode = returns the patch result but doesn't perform the patch.
	DryRun *bool `queryParam:"style=form,explode=true,name=dry_run"`
}

type PatchEntityRequest struct {
	PathParams  PatchEntityPathParams
	QueryParams PatchEntityQueryParams
	Request     map[string]interface{} `request:"mediaType=application/json"`
}

type PatchEntityResponse struct {
	ContentType string
	// Entity was updated
	EntityItem  map[string]interface{}
	StatusCode  int
	RawResponse *http.Response
}
