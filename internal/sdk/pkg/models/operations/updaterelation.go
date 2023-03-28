// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"epilotentity/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateRelationPathParams struct {
	// The attribute that express meaning
	Attribute string `pathParam:"style=simple,explode=false,name=attribute"`
	// The attribute that express meaning
	EntityID string `pathParam:"style=simple,explode=false,name=entity_id"`
	// Entity id
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Entity Type
	Slug string `pathParam:"style=simple,explode=false,name=slug"`
}

type UpdateRelationQueryParams struct {
	// Don't wait for updated entity to become available in Search API. Useful for large migrations
	Async *bool `queryParam:"style=form,explode=true,name=async"`
}

type UpdateRelationRequestBody struct {
	Tags []string `json:"_tags,omitempty"`
}

type UpdateRelationRequest struct {
	PathParams  UpdateRelationPathParams
	QueryParams UpdateRelationQueryParams
	Request     *UpdateRelationRequestBody `request:"mediaType=application/json"`
}

type UpdateRelationResponse struct {
	ContentType string
	// Success
	RelationItem *shared.RelationItem
	StatusCode   int
	RawResponse  *http.Response
}
