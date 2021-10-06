---
page_title: "Provider: Yandex pdd"
subcategory: "dns"
description: |-
  Terraform provider for interacting with dns records hosted on Yandex pdd.
---

# Adman DNS Provider

## Example Usage

Do not keep your authentication password in HCL for production environments, use Terraform environment variables.

```terraform
provider "pdd" {
  token = "token"
}
```

## Schema

### Required

- **token** (String, Optional) User token to authenticate to API
