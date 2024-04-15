terraform {
  required_providers {
    casaos = {
      source = "hashicorp.local/rstuhlmuller/casaos"
    }
  }
}

provider "casaos" {
}

data "casaos_hardware" "main" {
}
