// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// EntitySearchResults - Success
type EntitySearchResults struct {
	Aggregations map[string]interface{} `json:"aggregations,omitempty"`
	Hits         *float64               `json:"hits,omitempty"`
	Results      []EntityItem           `json:"results,omitempty"`
}
