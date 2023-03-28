// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type CreateEntityRequest struct {
	RequestBody map[string]interface{} `request:"mediaType=application/json"`
	// Activity to include in event feed
	ActivityID *string `queryParam:"style=form,explode=true,name=activity_id"`
	// Don't wait for updated entity to become available in Search API. Useful for large migrations
	Async *bool `queryParam:"style=form,explode=true,name=async"`
	// Entity Schema
	Slug string `pathParam:"style=simple,explode=false,name=slug"`
}

type CreateEntityResponse struct {
	ContentType string
	// Success
	EntityItem  map[string]interface{}
	StatusCode  int
	RawResponse *http.Response
}
