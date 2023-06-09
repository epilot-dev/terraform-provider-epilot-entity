// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"epilot-entity/internal/sdk/pkg/models/shared"
	"net/http"
	"time"
)

type GetEntityActivityFeedRequest struct {
	// Get activities after this timestamp
	After *time.Time `queryParam:"style=form,explode=true,name=after"`
	// get activities before this timestamp
	Before *time.Time `queryParam:"style=form,explode=true,name=before"`
	// start from page
	From *int64 `queryParam:"style=form,explode=true,name=from"`
	// Entity id
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// max number of results to return
	Size *int64 `queryParam:"style=form,explode=true,name=size"`
	// Entity Type
	Slug string `pathParam:"style=simple,explode=false,name=slug"`
	// Activity type
	Type *string `queryParam:"style=form,explode=true,name=type"`
}

// GetEntityActivityFeed200ApplicationJSON - Success
type GetEntityActivityFeed200ApplicationJSON struct {
	Results []shared.ActivityItem `json:"results,omitempty"`
	Total   *int64                `json:"total,omitempty"`
}

type GetEntityActivityFeedResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Success
	GetEntityActivityFeed200ApplicationJSONObject *GetEntityActivityFeed200ApplicationJSON
}
