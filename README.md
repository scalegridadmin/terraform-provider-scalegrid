# Terraform Provider ScaleGrid

Official Terraform provider for the [ScaleGrid](https://scalegrid.ai) platform.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 1.5+
- A ScaleGrid account

## Installation

Add the provider to your Terraform configuration:

```hcl
terraform {
  required_providers {
    scalegrid = {
      source  = "scalegridadmin/scalegrid"
      version = "~> 0.1"
    }
  }
}
```

Then run:

```sh
terraform init
```

## Authentication

Configure exactly one credential shape in the `provider "scalegrid"` block:

```hcl
provider "scalegrid" {
  bearer_token = var.scalegrid_bearer_token
}
```

Or use username and password (the provider mints a bearer via ScaleGrid identity):

```hcl
provider "scalegrid" {
  username = var.scalegrid_username
  password = var.scalegrid_password
}
```

## Usage

Call sg-api operations through the generic `scalegrid_operation` data source:

```hcl
data "scalegrid_operation" "pg_versions" {
  operation_id = "get_scalegrid_service_versions"
  body_json = jsonencode({
    pageIndex = 0
    pageSize  = 10
  })
}
```

See the ScaleGrid platform documentation for the full list of operation IDs.

## Development

Provider source is generated from the ScaleGrid OpenAPI spec in Bitbucket (`scalegrid_admin/mcp`). This public repository contains registry documentation and release artifacts only; generated Go code is built in CI and is not published on `main`.
