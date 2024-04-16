// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"os"

	casaos "github.com/rstuhlmuller/terraform-provider-casaos/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// // Ensure casaosProvider satisfies various provider interfaces.
var _ provider.Provider = &casaosProvider{}

// var _ provider.ProviderWithFunctions = &casaosProvider{}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &casaosProvider{
			version: version,
		}
	}
}

// casaosProvider defines the provider implementation.
type casaosProvider struct {
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

func (p *casaosProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "casaos"
	resp.Version = p.version
}

func (p *casaosProvider) Schema(_ context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
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

func (p *casaosProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
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

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing CasaOS API Host",
			"The provider cannot create the CasaOS API client as there is a missing or empty value for the CasaOS API host. "+
				"Set the host value in the configuration or use the CASAOS_HOST environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}
	if username == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Missing CasaOS API Username",
			"The provider cannot create the CasaOS API client as there is a missing or empty value for the CasaOS API username. "+
				"Set the username value in the configuration or use the CASAOS_USERNAME environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}
	if password == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing CasaOS API Password",
			"The provider cannot create the CasaOS API client as there is a missing or empty value for the CasaOS API password. "+
				"Set the password value in the configuration or use the CASAOS_PASSWORD environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	client, err := casaos.NewClient(&host, &username, &password)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create CasaOS API Client",
			"An unexpected error occurred when creating the CasaOS API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"CasaOS Client Error: "+err.Error(),
		)
		return
	}

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *casaosProvider) Resources(ctx context.Context) []func() resource.Resource {
	// return []func() resource.Resource{
	// 	NewExampleResource,
	// }
	return nil
}

func (p *casaosProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewAppManagementWebAppGridDataSource,
	}
}

func (p *casaosProvider) Functions(ctx context.Context) []func() function.Function {
	// return []func() function.Function{
	// 	NewExampleFunction,
	// }
	return nil
}
