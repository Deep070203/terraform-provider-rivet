// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccActorsResource(t *testing.T) {
	rName := acctest.RandomWithPrefix("tftest")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccActorsResourceConfig(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("Rivet_Actors.test", "actor", "UPDATE ME!"),
					resource.TestCheckResourceAttr("Rivet_Actors.test", "project", "UPDATE ME!"),
					resource.TestCheckResourceAttr("Rivet_Actors.test", "environment", "UPDATE ME!"),
					resource.TestCheckResourceAttr("Rivet_Actors.test", "endpoint_type", "UPDATE ME!"),
					resource.TestCheckResourceAttr("Rivet_Actors.test", "actor", "UPDATE ME!"),
					resource.TestCheckResourceAttr("Rivet_Actors.test", "project", "UPDATE ME!"),
					resource.TestCheckResourceAttr("Rivet_Actors.test", "environment", "UPDATE ME!"),
					resource.TestCheckResourceAttr("Rivet_Actors.test", "endpoint_type", "UPDATE ME!"),
					resource.TestCheckResourceAttr("Rivet_Actors.test", "actor", "UPDATE ME!"),
					resource.TestCheckResourceAttr("Rivet_Actors.test", "region", "UPDATE ME!"),
					resource.TestCheckResourceAttr("Rivet_Actors.test", "tags", "UPDATE ME!"),
					resource.TestCheckResourceAttr("Rivet_Actors.test", "build", "UPDATE ME!"),
					resource.TestCheckResourceAttr("Rivet_Actors.test", "build_tags", "UPDATE ME!"),
					resource.TestCheckResourceAttr("Rivet_Actors.test", "runtime", "UPDATE ME!"),
					resource.TestCheckResourceAttr("Rivet_Actors.test", "network", "UPDATE ME!"),
					resource.TestCheckResourceAttr("Rivet_Actors.test", "resources", "UPDATE ME!"),
					resource.TestCheckResourceAttr("Rivet_Actors.test", "lifecycle", "UPDATE ME!"),
					resource.TestCheckResourceAttr("Rivet_Actors.test", "actor", "UPDATE ME!"),
					resource.TestCheckResourceAttr("Rivet_Actors.test", "project", "UPDATE ME!"),
					resource.TestCheckResourceAttr("Rivet_Actors.test", "environment", "UPDATE ME!"),
					resource.TestCheckResourceAttr("Rivet_Actors.test", "build", "UPDATE ME!"),
					resource.TestCheckResourceAttr("Rivet_Actors.test", "build_tags", "UPDATE ME!"),
					resource.TestCheckResourceAttr("Rivet_Actors.test", "actor", "UPDATE ME!"),
					resource.TestCheckResourceAttr("Rivet_Actors.test", "project", "UPDATE ME!"),
					resource.TestCheckResourceAttr("Rivet_Actors.test", "environment", "UPDATE ME!"),
					resource.TestCheckResourceAttr("Rivet_Actors.test", "override_kill_timeout", "UPDATE ME!"),
					
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

func testAccActorsResourceConfig(name string) string {
	return fmt.Sprintf(`
provider "Rivet" {
}

resource "Rivet_Actors" "test" {
  "UPDATE ME!"

}`)
}