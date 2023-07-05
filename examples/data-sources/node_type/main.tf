terraform {
  required_providers {
    altinitycloud = {
      source = "tatari.tv/dev/altinitycloud"
    }
  }
}

provider "altinitycloud" {
  api_endpoint = "https://acm.altinity.cloud/api"
  // set API token via environment variable ALTINITYCLOUD_API_TOKEN
}

data "altinitycloud_node_type" "example" {
  env_id = "648"
  name   = "r6a.xlarge"
}

output "node_type" {
  value = data.altinitycloud_node_type.example
}