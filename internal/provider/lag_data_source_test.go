package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"testing"
)

func TestLagDataSource(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:                 func() { testPreCheck(t) },
		ProtoV6ProviderFactories: testProtoV6ProviderFactories,
		TerraformVersionChecks:   []tfversion.TerraformVersionCheck{
			//tfversion.SkipBelow(version.Must(version.NewVersion("1.8.0"))),
		},
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testLagDataSourceConfig(1000, "hello"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.testlagger_lag.test", "output", "hello"),
				),
			},
		},
	})
}

func testLagDataSourceConfig(readDelay int64, input string) string {
	return fmt.Sprintf(`
data "testlagger_lag" "test" {
	read_delay = %d
	input = "%s"
}
`, readDelay, input)
}
