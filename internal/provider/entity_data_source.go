// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-entity/internal/sdk"
	"github.com/epilot-dev/terraform-provider-epilot-entity/internal/sdk/pkg/models/operations"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &EntityDataSource{}
var _ datasource.DataSourceWithConfigure = &EntityDataSource{}

func NewEntityDataSource() datasource.DataSource {
	return &EntityDataSource{}
}

// EntityDataSource is the data source implementation.
type EntityDataSource struct {
	client *sdk.SDK
}

// EntityDataSourceModel describes the data model.
type EntityDataSourceModel struct {
	CreatedAt types.String   `tfsdk:"created_at"`
	Entity    types.String   `tfsdk:"entity"`
	ID        types.String   `tfsdk:"id"`
	Org       types.String   `tfsdk:"org"`
	Schema    types.String   `tfsdk:"schema"`
	Slug      types.String   `tfsdk:"slug"`
	Tags      []types.String `tfsdk:"tags"`
	Title     types.String   `tfsdk:"title"`
	UpdatedAt types.String   `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *EntityDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_entity"
}

// Schema defines the schema for the data source.
func (r *EntityDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Entity DataSource",

		Attributes: map[string]schema.Attribute{
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"entity": schema.StringAttribute{
				Computed:    true,
				Description: `Parsed as JSON.`,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `Entity id`,
			},
			"org": schema.StringAttribute{
				Computed:    true,
				Description: `Organization Id the entity belongs to`,
			},
			"schema": schema.StringAttribute{
				Computed:    true,
				Description: `URL-friendly identifier for the entity schema`,
			},
			"slug": schema.StringAttribute{
				Required:    true,
				Description: `Entity Type`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
			"title": schema.StringAttribute{
				Computed:    true,
				Description: `Title of entity`,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *EntityDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *EntityDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *EntityDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.ID.ValueString()
	slug := data.Slug.ValueString()
	request := operations.GetEntityRequest{
		ID:   id,
		Slug: slug,
	}
	res, err := r.client.Entities.GetEntity(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.Object == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedEntityItem(res.Object.Entity)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
