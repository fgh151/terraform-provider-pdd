---
page_title: "zone Resource - terraform-provider-pdd"
subcategory: "dns"
description: |-
  dns zone resource.
---

## Example Usage

```terraform
resource "pdd_dns_zone" "test_com" {
  domain  = "test.com"
}
```

## Argument Reference

- `domain` - (Required) Domain name.
