terraform {
  required_providers {
    scalegrid = {
      source  = "scalegridadmin/scalegrid"
      version = "~> 0.1"
    }
  }
}

provider "scalegrid" {
  bearer_token = var.scalegrid_bearer_token
}

variable "scalegrid_bearer_token" {
  type      = string
  sensitive = true
}

data "scalegrid_operation" "pg_versions" {
  operation_id = "get_scalegrid_service_versions"
  body_json = jsonencode({
    pageIndex = 0
    pageSize  = 10
  })
}

output "pg_versions_response" {
  value     = data.scalegrid_operation.pg_versions.response
  sensitive = true
}
