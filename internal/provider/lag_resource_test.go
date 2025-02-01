package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestLagResource(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:                 func() { testPreCheck(t) },
		ProtoV6ProviderFactories: testProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testLagResourceConfig(1000, 1000, 1000, 1000, "one"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("testlagger_lag.test", "create_delay", "1000"),
					resource.TestCheckResourceAttr("testlagger_lag.test", "read_delay", "1000"),
					resource.TestCheckResourceAttr("testlagger_lag.test", "update_delay", "1000"),
					resource.TestCheckResourceAttr("testlagger_lag.test", "delete_delay", "1000"),
					resource.TestCheckResourceAttr("testlagger_lag.test", "input", "one"),
					resource.TestCheckResourceAttr("testlagger_lag.test", "output", "one"),
				),
			},
			// ImportState testing
			{
				ResourceName:                         "testlagger_lag.test",
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "id",
				ImportStateVerifyIgnore: []string{
					"create_delay",
					"delete_delay",
					"read_delay",
					"update_delay",
				},
			},
			// Update and Read testing
			{
				Config: testLagResourceConfig(1000, 1000, 1000, 1000, "two"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("testlagger_lag.test", "output", "two"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testLagResourceConfig(createDelay int64, readDelay int64, updateDelay int64, deleteDelay int64, input string) string {
	return fmt.Sprintf(`
resource "testlagger_lag" "test" {
	create_delay = %d
	read_delay = %d
	update_delay = %d
	delete_delay = %d
	input = "%s"
}
`, createDelay, readDelay, updateDelay, deleteDelay, input)
}
