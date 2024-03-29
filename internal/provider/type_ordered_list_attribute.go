// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type OrderedListAttribute struct {
	Purpose                  []types.String `tfsdk:"purpose"`
	Constraints              *Constraints   `tfsdk:"constraints"`
	DefaultValue             types.String   `tfsdk:"default_value"`
	Deprecated               types.Bool     `tfsdk:"deprecated"`
	EntityBuilderDisableEdit types.Bool     `tfsdk:"entity_builder_disable_edit"`
	FeatureFlag              types.String   `tfsdk:"feature_flag"`
	Group                    types.String   `tfsdk:"group"`
	Hidden                   types.Bool     `tfsdk:"hidden"`
	HideLabel                types.Bool     `tfsdk:"hide_label"`
	Icon                     types.String   `tfsdk:"icon"`
	Label                    types.String   `tfsdk:"label"`
	Layout                   types.String   `tfsdk:"layout"`
	Name                     types.String   `tfsdk:"name"`
	Order                    types.Int64    `tfsdk:"order"`
	Placeholder              types.String   `tfsdk:"placeholder"`
	PreviewValueFormatter    types.String   `tfsdk:"preview_value_formatter"`
	Protected                types.Bool     `tfsdk:"protected"`
	Readonly                 types.Bool     `tfsdk:"readonly"`
	RenderCondition          types.String   `tfsdk:"render_condition"`
	Required                 types.Bool     `tfsdk:"required"`
	SettingFlag              types.String   `tfsdk:"setting_flag"`
	ShowInTable              types.Bool     `tfsdk:"show_in_table"`
	Type                     types.String   `tfsdk:"type"`
	ValueFormatter           types.String   `tfsdk:"value_formatter"`
}
