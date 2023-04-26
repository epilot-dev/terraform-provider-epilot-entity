// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type RelationAttribute1 struct {
	Purpose                  []types.String                   `tfsdk:"purpose"`
	Actions                  []RelationAttributeActions1      `tfsdk:"actions"`
	AddButtonLabel           types.String                     `tfsdk:"add_button_label"`
	AllowedSchemas           []types.String                   `tfsdk:"allowed_schemas"`
	Constraints              map[string]types.String          `tfsdk:"constraints"`
	DefaultValue             types.String                     `tfsdk:"default_value"`
	Deprecated               types.Bool                       `tfsdk:"deprecated"`
	DetailsViewModeEnabled   types.Bool                       `tfsdk:"details_view_mode_enabled"`
	DrawerSize               types.String                     `tfsdk:"drawer_size"`
	EditMode                 types.String                     `tfsdk:"edit_mode"`
	EnableRelationPicker     types.Bool                       `tfsdk:"enable_relation_picker"`
	EnableRelationTags       types.Bool                       `tfsdk:"enable_relation_tags"`
	EntityBuilderDisableEdit types.Bool                       `tfsdk:"entity_builder_disable_edit"`
	FeatureFlag              types.String                     `tfsdk:"feature_flag"`
	Group                    types.String                     `tfsdk:"group"`
	HasPrimary               types.Bool                       `tfsdk:"has_primary"`
	Hidden                   types.Bool                       `tfsdk:"hidden"`
	HideLabel                types.Bool                       `tfsdk:"hide_label"`
	Icon                     types.String                     `tfsdk:"icon"`
	Label                    types.String                     `tfsdk:"label"`
	Layout                   types.String                     `tfsdk:"layout"`
	Name                     types.String                     `tfsdk:"name"`
	Order                    types.Int64                      `tfsdk:"order"`
	Placeholder              types.String                     `tfsdk:"placeholder"`
	PreviewValueFormatter    types.String                     `tfsdk:"preview_value_formatter"`
	Protected                types.Bool                       `tfsdk:"protected"`
	Readonly                 types.Bool                       `tfsdk:"readonly"`
	RelationAffinityMode     types.String                     `tfsdk:"relation_affinity_mode"`
	RelationType             types.String                     `tfsdk:"relation_type"`
	RenderCondition          types.String                     `tfsdk:"render_condition"`
	Required                 types.Bool                       `tfsdk:"required"`
	ReverseAttributes        map[string]types.String          `tfsdk:"reverse_attributes"`
	SearchPlaceholder        types.String                     `tfsdk:"search_placeholder"`
	SettingFlag              types.String                     `tfsdk:"setting_flag"`
	ShowInTable              types.Bool                       `tfsdk:"show_in_table"`
	SummaryFields            []RelationAttributeSummaryFields `tfsdk:"summary_fields"`
	Type                     types.String                     `tfsdk:"type"`
	ValueFormatter           types.String                     `tfsdk:"value_formatter"`
}
