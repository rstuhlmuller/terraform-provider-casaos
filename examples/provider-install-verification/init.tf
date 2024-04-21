terraform {
  required_providers {
    casaos = {
      source = "hashicorp.com/rstuhlmuller/casaos"
    }
  }
}

provider "casaos" {}

data "casaos_app_grid" "example" {}
