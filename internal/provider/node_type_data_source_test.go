package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestNodeTypeDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testNodeTypeProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: altinityCloudNodeTypeExampleDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.altinitycloud_node_type.test", "id", "node_type_111"),
				),
			},
		},
	})
}

const altinityCloudNodeTypeExampleDataSourceConfig = `
data "altinitycloud_node_type" "test" {
  env_id = "111"
}
`
