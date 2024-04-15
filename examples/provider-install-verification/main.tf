terraform {
  required_providers {
    casaos = {
      source = "hashicorp.local/rstuhlmuller/casaos"
    }
  }
}

provider "casaos" {
  # host     = "http://casaos.local:80"
  # username = "themanofrod"
  # password = "pU$$ylUvr69"
}

# data "casaos_hardware" "main" {
# }
