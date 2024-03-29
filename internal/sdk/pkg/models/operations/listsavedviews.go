// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-entity/internal/sdk/pkg/models/shared"
	"net/http"
)

// ListSavedViewsResponseBody - Success
type ListSavedViewsResponseBody struct {
	Results []shared.SavedViewItem `json:"results,omitempty"`
}

func (o *ListSavedViewsResponseBody) GetResults() []shared.SavedViewItem {
	if o == nil {
		return nil
	}
	return o.Results
}

type ListSavedViewsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	Object *ListSavedViewsResponseBody
}

func (o *ListSavedViewsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListSavedViewsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListSavedViewsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListSavedViewsResponse) GetObject() *ListSavedViewsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
