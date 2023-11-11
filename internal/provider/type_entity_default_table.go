// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type EntityDefaultTable struct {
	ClassicView      types.String    `tfsdk:"classic_view"`
	DropdownItems    []DropdownItems `tfsdk:"dropdown_items"`
	EnableThumbnails types.Bool      `tfsdk:"enable_thumbnails"`
	NavbarActions    []NavbarActions `tfsdk:"navbar_actions"`
	RowActions       []types.String  `tfsdk:"row_actions"`
	ViewType         types.String    `tfsdk:"view_type"`
}
