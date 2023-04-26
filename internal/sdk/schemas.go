// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"context"
	"epilot-entity/internal/sdk/pkg/models/operations"
	"epilot-entity/internal/sdk/pkg/models/shared"
	"epilot-entity/internal/sdk/pkg/utils"
	"fmt"
	"net/http"
	"strings"
)

// schemas - Model Entities
type schemas struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newSchemas(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *schemas {
	return &schemas{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// DeleteSchema - deleteSchema
// Delete a schema, or a specific version of a schema
func (s *schemas) DeleteSchema(ctx context.Context, request operations.DeleteSchemaRequest) (*operations.DeleteSchemaResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/v1/entity/schemas/{slug}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.DeleteSchemaResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode >= 200 && httpRes.StatusCode < 300:
	}

	return res, nil
}

// GetSchema - getSchema
// By default gets the latest version of the Schema and to get the specific version of schema pass the id.
func (s *schemas) GetSchema(ctx context.Context, request operations.GetSchemaRequest) (*operations.GetSchemaResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/v1/entity/schemas/{slug}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetSchemaResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.EntitySchemaItem
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.EntitySchemaItem = out
		}
	}

	return res, nil
}

// GetSchemaVersions - getSchemaVersions
// Get all versions of this schema ordered by the latest versions including drafts.
func (s *schemas) GetSchemaVersions(ctx context.Context, request operations.GetSchemaVersionsRequest) (*operations.GetSchemaVersionsResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/v1/entity/schemas/{slug}/versions", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetSchemaVersionsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.GetSchemaVersions200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetSchemaVersions200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// ListSchemaBlueprints - listSchemaBlueprints
// List canonical versions of all available schemas
func (s *schemas) ListSchemaBlueprints(ctx context.Context) (*operations.ListSchemaBlueprintsResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/v1/entity/schemas/blueprints"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListSchemaBlueprintsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListSchemaBlueprints200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListSchemaBlueprints200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// ListSchemas - listSchemas
// Get the latest versions of all schemas
func (s *schemas) ListSchemas(ctx context.Context, request operations.ListSchemasRequest) (*operations.ListSchemasResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/v1/entity/schemas"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListSchemasResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListSchemas200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListSchemas200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// ListTaxonomyClassificationsForSchema - listTaxonomyClassificationsForSchema
// List taxonomy classifications for a given schema
func (s *schemas) ListTaxonomyClassificationsForSchema(ctx context.Context, request operations.ListTaxonomyClassificationsForSchemaRequest) (*operations.ListTaxonomyClassificationsForSchemaResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/v1/entity/schemas/{slug}/taxonomy/{taxonomySlug}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListTaxonomyClassificationsForSchemaResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListTaxonomyClassificationsForSchema200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListTaxonomyClassificationsForSchema200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// PutSchema - putSchema
// Create or update a schema with a new version
func (s *schemas) PutSchema(ctx context.Context, request operations.PutSchemaRequest) (*operations.PutSchemaResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/v1/entity/schemas/{slug}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "EntitySchema", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PutSchemaResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.PutSchema200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PutSchema200ApplicationJSONObject = out
		}
	}

	return res, nil
}
