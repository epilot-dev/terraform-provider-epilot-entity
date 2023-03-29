terraform {
  required_providers {
    epilot-entity = {
      source  = "epilot-dev/epilot-entity"
      version = "0.0.2"
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
  name   = "contact"
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
      name = "customer_messaging"
      title = "Messaging"
      attributes = [
        {
          text_attribute = {
            name             = "example"
            label            = "example"
            placeholder      = "example placeholder"
              show_in_table    = true
          }
        }
      ]
      ui_hooks = [
        {
          hook             = "EntityDetailsV2:Tab"
        }
      ]
    }
  ]
}