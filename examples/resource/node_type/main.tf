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

resource "altinitycloud_node_type" "example" {
  env_id = "648"
  node_type = {
    name          = "tf_example"
    scope         = "clickhouse"
    code          = "danmahoneyexample"
    storage_class = "gp3"
    memory        = "8192"
    cpu           = "4"
    pool          = "m6a.xlarge"
    tolerations = [
      {
        key      = "dedicated"
        operator = "Equal"
        effect   = "NoSchedule"
        value    = "altinity-example"
      }
    ]
  }
}
