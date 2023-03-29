// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"epilot-entity/internal/sdk/pkg/models/shared"
	"net/http"
)

type ExportEntitiesRequest struct {
	EntitySearchParams *shared.EntitySearchParams `request:"mediaType=application/json"`
	// Pass 'true' to generate import template
	IsTemplate *bool `queryParam:"style=form,explode=true,name=is_template"`
	// Export Job Id to get the result
	JobID *string `queryParam:"style=form,explode=true,name=job_id"`
	// Export headers translation language
	Language *string `queryParam:"style=form,explode=true,name=language"`
}

type ExportEntitiesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
