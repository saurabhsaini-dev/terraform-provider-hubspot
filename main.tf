terraform {
  required_providers {
    hubspot = {
      version = "0.1.0"
      source  = "hashicorp.com/edu/hubspot"
    }
  }
}

provider "hubspot" { 
  client_id     = "433d25ae-2157-4962-95ee-fe1bf7b4046b"
  client_secret = "12d5c888-67a2-4bc0-9b21-9f414d401e1b"
  refresh_token =  "5872de09-a2b9-4da1-92b4-37bf4a653bb7"
}