---
page_title: "ScaleGrid Provider"
description: |-
  Terraform provider for the ScaleGrid platform API.
---

# ScaleGrid Provider

Use this provider to call ScaleGrid API operations from Terraform via the generic `scalegrid_operation` data source.

## Authentication

Configure exactly one credential shape in the provider block:

- `bearer_token`, or
- `username` and `password` (the provider mints a bearer via ScaleGrid identity)

## Example Usage

```terraform
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

data "scalegrid_operation" "example" {
  operation_id = "get_scalegrid_service_versions"
  body_json = jsonencode({
    pageIndex = 0
    pageSize  = 10
  })
}
```
