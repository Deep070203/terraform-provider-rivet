// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccActorsDataSource(t *testing.T) {
	rName := acctest.RandomWithPrefix("tftest")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccActorsDataSourceConfig(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					
				),
			},
			// ImportState testing
			{
				ResourceName:      "Rivet_Actors.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccActorsDataSourceConfig(name string) string {
	return fmt.Sprintf(`
provider "Rivet" {
}

resource "Rivet_Actors" "test" {
  "UPDATE ME!"

}`)
}