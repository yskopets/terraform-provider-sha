// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestBase64Sha1_Known(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::sha::base64sha1("testvalue")
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "/Dz+K09VTCdSRE8ea3utHo1czPg="),
				),
			},
		},
	})
}

func TestBase64Sha1_Null(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::sha::base64sha1(null)
				}
				`,
				// The parameter does not enable AllowNullValue
				ExpectError: regexp.MustCompile(`argument must not be null`),
			},
		},
	})
}

func TestBase64Sha1_Unknown(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "terraform_data" "test" {
					input = "testvalue"
				}
				
				output "test" {
					value = provider::sha::base64sha1(terraform_data.test.output)
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "/Dz+K09VTCdSRE8ea3utHo1czPg="),
				),
			},
		},
	})
}
