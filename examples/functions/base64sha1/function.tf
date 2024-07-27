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
