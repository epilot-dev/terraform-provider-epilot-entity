// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type EntityImportParamsS3Reference struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

type EntityImportParams struct {
	S3Reference EntityImportParamsS3Reference `json:"S3Reference"`
	Schema      string                        `json:"schema"`
}
