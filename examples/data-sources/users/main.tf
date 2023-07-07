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

data "altinitycloud_users" "example" {
  cluster_id = "3652"
}

output "node_type" {
  value = data.altinitycloud_users.example
}