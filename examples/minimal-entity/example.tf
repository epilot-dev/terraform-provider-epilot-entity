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
  slug   = "test_contact"
  name   = "contact"
  plural = "contacts"
  attributes = [
    {
      text_attribute = {
        label = "foo"
        name  = "bar"
      }
    }
  ]
}