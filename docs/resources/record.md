---
page_title: "zone Resource - terraform-provider-pdd"
subcategory: "dns"
description: |-
  dns zone resource.
---

## Example Usage

```terraform
resource "pdd_dns_zone_record" "a_a" {
  zone = pdd_dns_zone.test_com.domain
  host            = "A"
  type            = "TXT"
  value           = "11.22.33.44"
  ttl             = 10
  external_id     = ""
  additional_info = ""
}
```

## Argument Reference

- `host` - (Required) Subdomain.
- `type` - (Required) Record type.
- `value` - (Required) Record value.
- `ttl` - (Optional) Record ttl.
- `external_id` - (Optional).
- `additional_info` - (Optional) Text description.
