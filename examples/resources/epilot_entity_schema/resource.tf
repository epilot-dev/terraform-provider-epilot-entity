resource "epilot_entity_schema" "my_schema" {
    attributes = [
        {
            purpose = [
                "c2ddf7cc-78ca-41ba-928f-c816742cb739",
            ]
            constraints = {}
            default_value = "{ \"see\": \"documentation\" }"
            deprecated = true
            entity_builder_disable_edit = false
            feature_flag = "FF_MY_FEATURE_FLAG"
            group = "...my_group..."
            has_primary = false
            hidden = true
            hide_label = false
            icon = "...my_icon..."
            label = "...my_label..."
            layout = "full_width"
            name = "Curtis Morissette"
            order = 0
            placeholder = "...my_placeholder..."
            preview_value_formatter = "...my_preview_value_formatter..."
            protected = true
            readonly = true
            render_condition = "...my_render_condition..."
            required = true
            setting_flag = "MY_SETTING"
            show_in_table = true
            type = "relation_address"
            value_formatter = "...my_value_formatter..."
        },
    ]
            blueprint = "6eb10faa-a235-42c5-9559-07aff1a3a2fa"
            capabilities = [
        {
            purpose = [
                "94677392-51aa-452c-bf5a-d019da1ffe78",
            ]
            attributes = [
                {
                    purpose = [
                        "f097b007-4f15-4471-b5e6-e13b99d488e1",
                    ]
                    constraints = {}
                    default_value = "{ \"see\": \"documentation\" }"
                    deprecated = true
                    entity_builder_disable_edit = false
                    feature_flag = "FF_MY_FEATURE_FLAG"
                    group = "...my_group..."
                    has_primary = true
                    hidden = false
                    hide_label = true
                    icon = "...my_icon..."
                    label = "...my_label..."
                    layout = "full_width"
                    name = "Elizabeth Orn"
                    order = 0
                    placeholder = "...my_placeholder..."
                    preview_value_formatter = "...my_preview_value_formatter..."
                    protected = false
                    readonly = true
                    render_condition = "...my_render_condition..."
                    required = true
                    setting_flag = "MY_SETTING"
                    show_in_table = true
                    type = "relation_address"
                    value_formatter = "...my_value_formatter..."
                },
            ]
            feature_flag = "FF_MY_FEATURE_FLAG"
            legacy = false
            name = "customer_messaging"
            setting_flag = "MY_SETTING"
            title = "Messaging"
            ui_hooks = [
                {
                    component = "PricingItems"
                    disabled = false
                    group_expanded = false
                    header = false
                    hook = "EntityDetailsV2:Tab"
                    icon = "email"
                    import = "@epilot360/notes"
                    order = 10
                    render_condition = "_is_composite_price = \"false\""
                    required_permission = {
                        action = "note:view"
                        resource = 123
                    }
                    route = "notes"
                    title = "Notes"
                },
            ]
        },
    ]
            dialog_config = {
        "facilis" = "{ \"see\": \"documentation\" }"
        "tempore" = "{ \"see\": \"documentation\" }"
    }
            draft = false
            enable_setting = [
        "360_features",
    ]
            explicit_search_mappings = {
        "delectus" = {
            fields = {
                "eum" = "{ \"see\": \"documentation\" }"
                "non" = "{ \"see\": \"documentation\" }"
            }
            index = false
            type = "float"
        }
        "aliquid" = {
            fields = {
                "provident" = "{ \"see\": \"documentation\" }"
                "necessitatibus" = "{ \"see\": \"documentation\" }"
            }
            index = true
            type = "float"
        }
    }
            feature_flag = "FF_MY_FEATURE_FLAG"
            group_settings = [
        {
            purpose = [
                "3efa77df-b14c-4d66-ae39-5efb9ba88f3a",
            ]
            expanded = false
            feature_flag = "FF_MY_FEATURE_FLAG"
            id = "6997074b-a446-49b6-a214-1959890afa56"
            info_tooltip_title = {
                default = "...my_default..."
                key = "...my_key..."
            }
            label = "...my_label..."
            order = 2
            render_condition = "_is_composite_price = \"false\""
            setting_flag = "MY_SETTING"
        },
    ]
            icon = "person"
            layout_settings = {
        grid_gap = "...my_grid_gap..."
        grid_template_columns = "...my_grid_template_columns..."
    }
            name = "Contact"
            plural = "Contacts"
            published = false
            slug = "contact"
            title_template = "{{first_name}} {{last_name}}"
            ui_config = {
        create_view = {
            entity_default_create =     {
                    search_params = {
                        "eius" = "..."
                        "maxime" = "..."
                    }
                    table_menu_options = {
                        icon = "...my_icon..."
                        label = "...my_label..."
                    }
                    view_type = "default"
                }
        }
        edit_view = {
            entity_default_edit =     {
                    search_params = {
                        "deleniti" = "..."
                        "facilis" = "..."
                    }
                    table_menu_options = {
                        icon = "...my_icon..."
                        label = "...my_label..."
                    }
                    view_type = "default"
                }
        }
        list_item = {
            summary_attributes = [
                {
                    feature_flag = "...my_feature_flag..."
                    label = "...my_label..."
                    render_condition = "...my_render_condition..."
                    setting_flag = "...my_setting_flag..."
                    show_as_tag = true
                    tag_color = "...my_tag_color..."
                    value = "...my_value..."
                },
            ]
        }
        sharing = {
            show_sharing_button = true
        }
        single_view = {
            entity_default_edit =     {
                    search_params = {
                        "architecto" = "..."
                        "repudiandae" = "..."
                    }
                    table_menu_options = {
                        icon = "...my_icon..."
                        label = "...my_label..."
                    }
                    view_type = "default"
                }
        }
        table_view = {
            entity_default_table =     {
                    classic_view = "...my_classic_view..."
                    dropdown_items = [
                        {
                            entity = "contact"
                            feature_flag = "FF_MY_FEATURE_FLAG"
                            legacy = true
                            type = "entity"
                        },
                    ]
                    enable_thumbnails = true
                    navbar_actions = [
                        {
                            options = [
                                {
                                    label = "...my_label..."
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
            version = 5
        }