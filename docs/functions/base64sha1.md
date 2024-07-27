---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "base64sha1 function - sha"
subcategory: ""
description: |-
  Compute the SHA-1 hash of a given string and encode it with Base64
---

# function: base64sha1

Compute the SHA-1 hash of a given string and encode it with Base64.

## Example Usage

```terraform
terraform {
  required_providers {
    sha = {
      source = "registry.terraform.io/yskopets/sha"
    }
  }
  required_version = ">= 1.8.0"
}

provider "sha" {}

output "base64sha1" {
  value = provider :: sha :: base64sha1("secret")
}
```

## Signature

<!-- signature generated by tfplugindocs -->
```text
base64sha1(input string) string
```

## Arguments

<!-- arguments generated by tfplugindocs -->
1. `input` (String) String to compute the SHA1 hash of
