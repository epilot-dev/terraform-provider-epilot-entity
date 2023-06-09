// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"epilot-entity/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSavedViewResponse struct {
	ContentType string
	// Success
	SavedViewItem *shared.SavedViewItem
	StatusCode    int
	RawResponse   *http.Response
}
