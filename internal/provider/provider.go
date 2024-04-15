// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"os"

	// 	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"

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

// casaosProviderModel describes the provider data model.
type casaosProviderModel struct {
	Host     types.String `tfsdk:"host"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

func (p *CasaOSProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "casaos"
	resp.Version = p.version
}

func (p *CasaOSProvider) Schema(_ context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				MarkdownDescription: "Host IP or DNS name.",
				Optional:            true,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "Username of casaos.",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Password of casaos.",
				Optional:            true,
				Sensitive:           true,
			},
		},
	}
}

func (p *CasaOSProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Where we authenticate using username - password etc.

	var config casaosProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown CasaOS API Host",
			"The provider cannot create the CasaOS API client as there is a missing or empty value for the CasaOS API host."+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the CASAOS_HOST environment variable.",
		)
	}

	if config.Username.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Unknown CasaOS API Username",
			"The provider cannot create the CasaOS API client as there is an unknown configuration value for the CasaOS API username. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the CASAOS_USERNAME environment variable.",
		)
	}

	if config.Password.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown CasaOS API Password",
			"The provider cannot create the CasaOS API client as there is an unknown configuration value for the CasaOS API password. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the CASAOS_PASSWORD environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	host := os.Getenv("CASAOS_HOST")
	username := os.Getenv("CASAOS_USERNAME")
	password := os.Getenv("CASAOS_PASSWORD")

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}
	if !config.Username.IsNull() {
		username = config.Username.ValueString()
	}
	if !config.Password.IsNull() {
		password = config.Password.ValueString()
	}
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
