// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-entity/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateClassificationsForTaxonomyRequest struct {
	ClassificationsUpdate *shared.ClassificationsUpdate `request:"mediaType=application/json"`
	// Taxonomy slug
	TaxonomySlug string `pathParam:"style=simple,explode=false,name=taxonomySlug"`
}

func (o *UpdateClassificationsForTaxonomyRequest) GetClassificationsUpdate() *shared.ClassificationsUpdate {
	if o == nil {
		return nil
	}
	return o.ClassificationsUpdate
}

func (o *UpdateClassificationsForTaxonomyRequest) GetTaxonomySlug() string {
	if o == nil {
		return ""
	}
	return o.TaxonomySlug
}

type Deleted struct {
}

// UpdateClassificationsForTaxonomyResponseBody - Taxonomies classifications
type UpdateClassificationsForTaxonomyResponseBody struct {
	Created []shared.TaxonomyClassification `json:"created,omitempty"`
	Deleted *Deleted                        `json:"deleted,omitempty"`
	Updated []shared.TaxonomyClassification `json:"updated,omitempty"`
}

func (o *UpdateClassificationsForTaxonomyResponseBody) GetCreated() []shared.TaxonomyClassification {
	if o == nil {
		return nil
	}
	return o.Created
}

func (o *UpdateClassificationsForTaxonomyResponseBody) GetDeleted() *Deleted {
	if o == nil {
		return nil
	}
	return o.Deleted
}

func (o *UpdateClassificationsForTaxonomyResponseBody) GetUpdated() []shared.TaxonomyClassification {
	if o == nil {
		return nil
	}
	return o.Updated
}

type UpdateClassificationsForTaxonomyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Taxonomies classifications
	Object *UpdateClassificationsForTaxonomyResponseBody
}

func (o *UpdateClassificationsForTaxonomyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateClassificationsForTaxonomyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateClassificationsForTaxonomyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateClassificationsForTaxonomyResponse) GetObject() *UpdateClassificationsForTaxonomyResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
