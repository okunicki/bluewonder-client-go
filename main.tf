terraform {
  required_providers {
    bluewonder = {
      source  = "my-org/bluewonder"
      version = "0.0.1"
    }
  }
}

provider "bluewonder" {}

resource "bluewonder_me" "example" {}
