terraform {
  required_providers {
    epilot-entity = {
      source  = "epilot-dev/epilot-entity"
      version = "0.2.0"
    }
  }
}

provider "epilot-entity" {
  epilot_auth = var.epilot_api_key
}

variable "epilot_api_key" {
  type = string
}

resource "epilot-entity_schema" "test" {
  slug   = "contact"
  name   = "contact from terraform"
  plural = "contacts"
  attributes = [
    {
      text_attribute = {
        label = "test"
        name  = "test"
      }
    }
  ]
  capabilities = [
    {
      name  = "customer_messaging"
      title = "Messaging"
      attributes = [
        {
          text_attribute = {
            name          = "example"
            label         = "example"
            placeholder   = "example placeholder"
            show_in_table = true
          }
        }
      ]
      ui_hooks = [
        {
          hook = "EntityDetailsV2:Tab"
        }
      ]
    }
  ]
}

resource "epilot-entity_entity" "contact_test" {
  slug = epilot-entity_schema.test.slug

  entity = jsonencode({
    "test" = "test"
  })
}