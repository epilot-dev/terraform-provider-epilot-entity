// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-entity/internal/sdk/pkg/models/shared"
	"net/http"
)

// ListTaxonomiesResponseBody - Returns list of taxonomies in an organisation
type ListTaxonomiesResponseBody struct {
	Results []shared.Taxonomy `json:"results,omitempty"`
}

func (o *ListTaxonomiesResponseBody) GetResults() []shared.Taxonomy {
	if o == nil {
		return nil
	}
	return o.Results
}

type ListTaxonomiesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Returns list of taxonomies in an organisation
	Object *ListTaxonomiesResponseBody
}

func (o *ListTaxonomiesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListTaxonomiesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListTaxonomiesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListTaxonomiesResponse) GetObject() *ListTaxonomiesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
