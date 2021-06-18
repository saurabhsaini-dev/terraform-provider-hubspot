terraform {
  required_providers {
    hubspot = {
      version = "0.1.0"
      source  = "hashicorp.com/edu/hubspot"
    }
  }
}

provider "hubspot" { 
  client_id     = "AAA"
  client_secret = "AAA"
  refresh_token = "AAA"
}