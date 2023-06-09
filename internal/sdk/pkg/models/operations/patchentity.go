// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"epilot-entity/internal/sdk/pkg/models/shared"
	"net/http"
)

type PatchEntityRequest struct {
	Entity *shared.Entity `request:"mediaType=application/json"`
	// Activity to include in event feed
	ActivityID *string `queryParam:"style=form,explode=true,name=activity_id"`
	// Don't wait for the patch entity to become available in Search API. Useful for large migrations
	Async *bool `queryParam:"style=form,explode=true,name=async"`
	// Dry Run mode = returns the patch result but doesn't perform the patch.
	DryRun *bool `queryParam:"style=form,explode=true,name=dry_run"`
	// Entity Id
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Entity Schema
	Slug string `pathParam:"style=simple,explode=false,name=slug"`
}

type PatchEntityResponse struct {
	ContentType string
	// Entity was updated
	EntityItem  *shared.EntityItem
	StatusCode  int
	RawResponse *http.Response
}
