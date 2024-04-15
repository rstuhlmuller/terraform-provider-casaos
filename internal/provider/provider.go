// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	// 	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	// "github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	// "github.com/hashicorp/terraform-plugin-framework/types"
)

// // Ensure CasaOSProvider satisfies various provider interfaces.
var _ provider.Provider = &CasaOSProvider{}

// var _ provider.ProviderWithFunctions = &CasaOSProvider{}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CasaOSProvider{
			version: version,
		}
	}
}

// CasaOSProvider defines the provider implementation.
type CasaOSProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// // CasaOSProviderModel describes the provider data model.
// type CasaOSProviderModel struct {
// 	Endpoint types.String `tfsdk:"endpoint"`
// }

func (p *CasaOSProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "casaos"
	resp.Version = p.version
}

func (p *CasaOSProvider) Schema(_ context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
	// Attributes: map[string]schema.Attribute{
	// 	"endpoint": schema.StringAttribute{
	// 		MarkdownDescription: "Example provider attribute",
	// 		Optional:            true,
	// 	},
	// },
	// }
}

func (p *CasaOSProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Where we authenticate using username - password etc.

	// var data CasaOSProviderModel

	// resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	// if resp.Diagnostics.HasError() {
	// 	return
	// }

	// // Configuration values are now available.
	// // if data.Endpoint.IsNull() { /* ... */ }

	// // Example client configuration for data sources and resources
	// client := http.DefaultClient
	// resp.DataSourceData = client
	// resp.ResourceData = client
}

func (p *CasaOSProvider) Resources(ctx context.Context) []func() resource.Resource {
	// return []func() resource.Resource{
	// 	NewExampleResource,
	// }
	return nil
}

func (p *CasaOSProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	// return []func() datasource.DataSource{
	// 	NewExampleDataSource,
	// }
	return nil
}

func (p *CasaOSProvider) Functions(ctx context.Context) []func() function.Function {
	// return []func() function.Function{
	// 	NewExampleFunction,
	// }
	return nil
}
