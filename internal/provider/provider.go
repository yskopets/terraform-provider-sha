// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure ShaProvider satisfies various provider interfaces.
var _ provider.Provider = &ShaProvider{}
var _ provider.ProviderWithFunctions = &ShaProvider{}

// ShaProvider defines the provider implementation.
type ShaProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

func (p *ShaProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "sha"
	resp.Version = p.version
}

func (p *ShaProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (p *ShaProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *ShaProvider) Resources(ctx context.Context) []func() resource.Resource {
	return nil
}

func (p *ShaProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return nil
}

func (p *ShaProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewBase64Sha1Function,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ShaProvider{
			version: version,
		}
	}
}
