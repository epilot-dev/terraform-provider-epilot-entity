terraform {
  required_providers {
    epilot-entity = {
      source  = "epilot-dev/epilot-entity"
      version = "0.2.0"
    }
  }
}


variable "epilot_auth" {
  type = string
}


provider "epilot-entity" {
  # Configuration options

  epilot_auth = var.epilot_auth
}


resource "epilot-entity_schema" "imported_schema" {

}
