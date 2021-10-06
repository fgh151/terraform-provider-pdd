terraform {
  required_providers {
    pdd = {
      source = "openitstudio.ru/dns/pdd"
      version = "0.0.1"
    }
  }
}

provider "pdd" {
  token = "token"
}

resource "pdd_dns_zone" "test_com" {
  domain  = "test.com"
}

resource "pdd_dns_zone_record" "a_a" {
  zone = pdd_dns_zone.test_com.domain
  host            = "A"
  type            = "TXT"
  value           = "11.22.33.44"
  ttl             = 10
  external_id     = ""
  additional_info = ""
}
