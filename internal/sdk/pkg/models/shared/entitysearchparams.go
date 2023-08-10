// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// EntitySearchParamsAggs - Aggregation supported by ElasticSearch allows summarizing data as metrics, statistics, or other analytics.
type EntitySearchParamsAggs struct {
}

type EntitySearchParams struct {
	// Aggregation supported by ElasticSearch allows summarizing data as metrics, statistics, or other analytics.
	Aggs *EntitySearchParamsAggs `json:"aggs,omitempty"`
	// List of entity fields to include in search results
	Fields []string `json:"fields,omitempty"`
	From   *int64   `json:"from,omitempty"`
	// When true, enables entity hydration to resolve nested $relation & $relation_ref references in-place.
	Hydrate *bool `json:"hydrate,omitempty"`
	// Adds a `_score` number field to results that can be used to rank by match score
	IncludeScores *bool `json:"include_scores,omitempty"`
	// Lucene queries supported with ElasticSearch
	Q string `json:"q"`
	// Max search size is 1000 with higher values defaulting to 1000
	Size *int64  `json:"size,omitempty"`
	Sort *string `json:"sort,omitempty"`
}
