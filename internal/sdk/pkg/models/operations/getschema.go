// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"epilot-entity/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetSchemaRequest struct {
	ID   *string `queryParam:"style=form,explode=true,name=id"`
	Slug string  `pathParam:"style=simple,explode=false,name=slug"`
}

type GetSchemaResponse struct {
	ContentType string
	// Success
	EntitySchemaItem *shared.EntitySchemaItem
	StatusCode       int
	RawResponse      *http.Response
}
