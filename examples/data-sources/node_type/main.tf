terraform {
  required_providers {
    altinitycloud = {
      source = "tatari.tv/dev/altinitycloud"
    }
  }
}

provider "altinitycloud" {}

data "altinitycloud_node_type" "example" {
  env_id = "652"
}

output "node_type" {
  value = data.altinitycloud_node_type.example
}