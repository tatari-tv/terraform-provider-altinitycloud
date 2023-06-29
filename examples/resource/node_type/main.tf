terraform {
  required_providers {
    altinitycloud = {
      source = "tatari.tv/dev/altinitycloud"
    }
  }
}

provider "altinitycloud" {
  api_endpoint = "https://acm.altinity.cloud/api"
}

resource "altinitycloud_node_type" "example" {
  env_id = "648"
  node_type = {
    name          = "exampleTFProvider"
    scope         = "clickhouse"
    code          = "exampleTFProvider"
    storage_class = "gp3"
    memory        = "8192"
    cpu           = "4"
    pool          = "m6a.large"
    tolerations = [
      {
        key      = "key"
        operator = "Equal"
        effect   = "NoSchedule"
        value    = "dedicated=altinity-clickhouse"
      }
    ]
  }
}
