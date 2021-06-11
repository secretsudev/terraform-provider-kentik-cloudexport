terraform {
  required_providers {
    kentik-cloudexport = {
      version = ">= 0.3.0"
      source = "kentik/kentik-cloudexport" # use provider from registry.terraform.io
      # source  = "kentik/automation/kentik-cloudexport" # use locally-built provider (make install)
    }
  }
}

provider "kentik-cloudexport" {
  # email, token and apiurl are read from KTAPI_AUTH_EMAIL, KTAPI_AUTH_TOKEN, KTAPI_URL env variables
}