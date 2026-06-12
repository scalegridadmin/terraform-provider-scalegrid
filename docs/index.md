---
page_title: "ScaleGrid Provider"
description: |-
  Terraform provider for the ScaleGrid platform API.
---

# ScaleGrid Provider

Use this provider to call ScaleGrid API operations from Terraform via the generic `scalegrid_operation` resource.

## Authentication

Configure exactly one credential shape in the provider block:

- `bearer_token`, or
- `username` and `password` (the provider mints a bearer via ScaleGrid identity)

## API operations reference

Every supported ScaleGrid API operation is documented with a Terraform example on the [`scalegrid_operation` resource page](./resources/operation):

- [Full operations table and examples](./resources/operation) — all `operation_id` values, HTTP methods, paths, and copy-paste Terraform blocks
- [Operations index](./guides/index) — quick links to individual operation summaries

`terraform plan` validates your configuration without calling sg-api. API requests run only on `terraform apply`.

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

resource "scalegrid_operation" "example" {
  operation_id = "get_scalegrid_service_versions"
  body_json = jsonencode({
    pageIndex = 0
    pageSize  = 10
  })
}
```
