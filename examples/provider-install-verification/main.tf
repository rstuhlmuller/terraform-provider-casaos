terraform {
  required_providers {
    casaos = {
      source = "hashicorp.local/rstuhlmuller/casaos"
    }
  }
}

provider "casaos" {
  host     = "http://casaos.local"
  username = "themanofrod"
  password = "pU$$ylUvr69"
}

data "casaos_app_management_web_app_grid" "main" {}
