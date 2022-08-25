package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTeamResourceDefault(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccTeamResourceConfigDefault("Accounting"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("test_team.test", "id", "Accounting"),
					resource.TestCheckResourceAttr("test_team.test", "bool_empty_default", "false"),
					resource.TestCheckResourceAttr("test_team.test", "bool_known_default", "true"),
					resource.TestCheckResourceAttr("test_team.test", "string_empty_default", ""),
					resource.TestCheckResourceAttr("test_team.test", "string_known_default", "One"),
					resource.TestCheckNoResourceAttr("test_team.test", "nullable_string"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "test_team.test",
				ImportState:       true,
				ImportStateId:     "Accounting",
				ImportStateVerify: true,
			},
			// Update with same config again
			{
				Config: testAccTeamResourceConfigDefault("Accounting"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("test_team.test", "id", "Accounting"),
					resource.TestCheckResourceAttr("test_team.test", "bool_empty_default", "false"),
					resource.TestCheckResourceAttr("test_team.test", "bool_known_default", "true"),
					resource.TestCheckResourceAttr("test_team.test", "string_empty_default", ""),
					resource.TestCheckResourceAttr("test_team.test", "string_known_default", "One"),
					resource.TestCheckNoResourceAttr("test_team.test", "nullable_string"),
				),
			},
			// Update and Read testing
			{
				Config: testAccTeamResourceConfigNonDefault("Accounting"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("test_team.test", "id", "Accounting"),
					resource.TestCheckResourceAttr("test_team.test", "bool_empty_default", "true"),
					resource.TestCheckResourceAttr("test_team.test", "bool_known_default", "false"),
					resource.TestCheckResourceAttr("test_team.test", "string_empty_default", "Image"),
					resource.TestCheckResourceAttr("test_team.test", "string_known_default", "Bank"),
					resource.TestCheckResourceAttr("test_team.test", "nullable_string", "London"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "test_team.test",
				ImportState:       true,
				ImportStateId:     "Accounting",
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccTeamResourceNonDefault(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccTeamResourceConfigNonDefault("DevOps"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("test_team.test", "id", "DevOps"),
					resource.TestCheckResourceAttr("test_team.test", "bool_empty_default", "true"),
					resource.TestCheckResourceAttr("test_team.test", "bool_known_default", "false"),
					resource.TestCheckResourceAttr("test_team.test", "string_empty_default", "Image"),
					resource.TestCheckResourceAttr("test_team.test", "string_known_default", "Bank"),
					resource.TestCheckResourceAttr("test_team.test", "nullable_string", "London"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "test_team.test",
				ImportState:       true,
				ImportStateId:     "DevOps",
				ImportStateVerify: true,
			},
			// Update with same config again
			{
				Config: testAccTeamResourceConfigNonDefault("DevOps"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("test_team.test", "id", "DevOps"),
					resource.TestCheckResourceAttr("test_team.test", "bool_empty_default", "true"),
					resource.TestCheckResourceAttr("test_team.test", "bool_known_default", "false"),
					resource.TestCheckResourceAttr("test_team.test", "string_empty_default", "Image"),
					resource.TestCheckResourceAttr("test_team.test", "string_known_default", "Bank"),
					resource.TestCheckResourceAttr("test_team.test", "nullable_string", "London"),
				),
			},
			// Update with null config
			{
				Config: testAccTeamResourceConfigDefault("DevOps"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("test_team.test", "id", "DevOps"),
					resource.TestCheckResourceAttr("test_team.test", "bool_empty_default", "false"),
					resource.TestCheckResourceAttr("test_team.test", "bool_known_default", "true"),
					resource.TestCheckResourceAttr("test_team.test", "string_empty_default", ""),
					resource.TestCheckResourceAttr("test_team.test", "string_known_default", "One"),
					resource.TestCheckNoResourceAttr("test_team.test", "nullable_string"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccTeamResourceConfigDefault(name string) string {
	return fmt.Sprintf(`
resource "test_team" "test" {
  id = "%s"
}
`, name)
}

func testAccTeamResourceConfigNonDefault(name string) string {
	return fmt.Sprintf(`
resource "test_team" "test" {
  id = "%s"
  bool_empty_default = true
  bool_known_default = false
  string_empty_default = "Image"
  string_known_default = "Bank"
  nullable_string = "London"
}
`, name)
}
