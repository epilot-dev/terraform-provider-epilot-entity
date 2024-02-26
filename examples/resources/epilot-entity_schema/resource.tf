resource "epilot-entity_schema" "my_schema" {
  attributes = [
    {
      address_relation_attribute = {
        purpose = [
          "caf7351b-352e-4984-aecb-2b1eab0d73fe",
        ]
        constraints                 = {}
        default_value               = "{ \"see\": \"documentation\" }"
        deprecated                  = true
        entity_builder_disable_edit = true
        feature_flag                = "FF_MY_FEATURE_FLAG"
        group                       = "...my_group..."
        has_primary                 = false
        hidden                      = false
        hide_label                  = true
        icon                        = "...my_icon..."
        label                       = "...my_label..."
        layout                      = "full_width"
        name                        = "Miss Mamie Dooley"
        order                       = 0
        placeholder                 = "...my_placeholder..."
        preview_value_formatter     = "...my_preview_value_formatter..."
        protected                   = false
        readonly                    = true
        render_condition            = "...my_render_condition..."
        required                    = false
        setting_flag                = "MY_SETTING"
        show_in_table               = true
        type                        = "relation_address"
        value_formatter             = "...my_value_formatter..."
      }
    },
  ]
  blueprint = "8aa711d9-eed9-4891-8711-fd8d9d547bbc"
  capabilities = [
    {
      purpose = [
        "819d9639-5fa7-49a1-b280-372c96b7bf2f",
      ]
      attributes = [
        {
          address_relation_attribute = {
            purpose = [
              "f67fd353-b397-4e7b-9d76-40e6321a646e",
            ]
            constraints                 = {}
            default_value               = "{ \"see\": \"documentation\" }"
            deprecated                  = false
            entity_builder_disable_edit = false
            feature_flag                = "FF_MY_FEATURE_FLAG"
            group                       = "...my_group..."
            has_primary                 = true
            hidden                      = true
            hide_label                  = false
            icon                        = "...my_icon..."
            label                       = "...my_label..."
            layout                      = "full_width"
            name                        = "Jeannette Marvin Jr."
            order                       = 0
            placeholder                 = "...my_placeholder..."
            preview_value_formatter     = "...my_preview_value_formatter..."
            protected                   = false
            readonly                    = true
            render_condition            = "...my_render_condition..."
            required                    = false
            setting_flag                = "MY_SETTING"
            show_in_table               = true
            type                        = "relation_address"
            value_formatter             = "...my_value_formatter..."
          }
        },
      ]
      feature_flag = "FF_MY_FEATURE_FLAG"
      legacy       = false
      name         = "customer_messaging"
      setting_flag = "MY_SETTING"
      title        = "Messaging"
      ui_hooks = [
        {
          component        = "PricingItems"
          disabled         = true
          group_expanded   = false
          header           = false
          hook             = "EntityDetailsV2:Tab"
          icon             = "email"
          import           = "@epilot360/notes"
          order            = 10
          render_condition = "_is_composite_price = \"false\""
          required_permission = {
            action   = "note:view"
            resource = 123
          }
          route = "notes"
          title = "Notes"
        },
      ]
    },
  ]
  dialog_config = {
    "male" = "{ \"see\": \"documentation\" }"
    "3rd"  = "{ \"see\": \"documentation\" }"
  }
  draft = false
  enable_setting = [
    "360_features",
  ]
  explicit_search_mappings = {
    "indeed" = {
      fields = {
        "Interactions" = "{ \"see\": \"documentation\" }"
        "Unbranded"    = "{ \"see\": \"documentation\" }"
      }
      index = false
      type  = "integer"
    }
    "Elegant" = {
      fields = {
        "Unbranded"  = "{ \"see\": \"documentation\" }"
        "Electronic" = "{ \"see\": \"documentation\" }"
      }
      index = true
      type  = "text"
    }
  }
  feature_flag = "FF_MY_FEATURE_FLAG"
  group_settings = [
    {
      purpose = [
        "bd8cdb99-20df-4010-8c1c-e693b06314bc",
      ]
      expanded     = false
      feature_flag = "FF_MY_FEATURE_FLAG"
      id           = "e26a5074-d671-44c0-bdac-b2ca16c72517"
      info_tooltip_title = {
        default = "...my_default..."
        key     = "...my_key..."
      }
      label            = "...my_label..."
      order            = 0
      render_condition = "_is_composite_price = \"false\""
      setting_flag     = "MY_SETTING"
    },
  ]
  icon = "person"
  layout_settings = {
    additional_properties = "{ \"see\": \"documentation\" }"
    grid_gap              = "...my_grid_gap..."
    grid_template_columns = "...my_grid_template_columns..."
  }
  name           = "Contact"
  plural         = "Contacts"
  published      = false
  slug           = "contact"
  title_template = "{{first_name}} {{last_name}}"
  ui_config = {
    create_view = {
      entity_default_create = {
        search_params = {
          "Concrete" = "..."
          "Sleek"    = "..."
        }
        table_menu_options = {
          icon  = "...my_icon..."
          label = "...my_label..."
        }
        view_type = "default"
      }
    }
    edit_view = {
      entity_default_edit = {
        search_params = {
          "Southwest" = "..."
          "Card"      = "..."
        }
        table_menu_options = {
          icon  = "...my_icon..."
          label = "...my_label..."
        }
        view_type = "default"
      }
    }
    list_item = {
      summary_attributes = [
        {
          str = "email"
        },
      ]
    }
    sharing = {
      show_sharing_button = true
    }
    single_view = {
      entity_default_edit = {
        search_params = {
          "International" = "..."
          "chew"          = "..."
        }
        table_menu_options = {
          icon  = "...my_icon..."
          label = "...my_label..."
        }
        view_type = "default"
      }
    }
    table_view = {
      entity_default_table = {
        classic_view = "...my_classic_view..."
        dropdown_items = [
          {
            one = {
              entity       = "contact"
              feature_flag = "FF_MY_FEATURE_FLAG"
              legacy       = true
              type         = "entity"
            }
          },
        ]
        enable_thumbnails = true
        navbar_actions = [
          {
            options = [
              {
                label  = "...my_label..."
                params = {}
              },
            ]
            type = "...my_type..."
          },
        ]
        row_actions = [
          "...",
        ]
        view_type = "default"
      }
    }
  }
  version = 8
}