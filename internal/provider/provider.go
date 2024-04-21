package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/rstuhlmuller/terraform-provider-casaos/internal/casaos"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &casaosProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &casaosProvider{
			version: version,
		}
	}
}

// casaosProvider is the provider implementation.
type casaosProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

type casaosProviderModel struct {
	Host     types.String `tfsdk:"host"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

// Metadata returns the provider type name.
func (p *casaosProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "casaos"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *casaosProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Optional: true,
			},
			"username": schema.StringAttribute{
				Optional: true,
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

// Configure prepares a Casaos API client for data sources and resources.
func (p *casaosProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config casaosProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown Casaos API Host",
			"The provider cannot create the Casaos API client as there is an unknown configuration value for the Casaos API host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the CASAOS_HOST environment variable.",
		)
	}

	if config.Username.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Unknown Casaos API Username",
			"The provider cannot create the Casaos API client as there is an unknown configuration value for the Casaos API username. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the CASAOS_USERNAME environment variable.",
		)
	}

	if config.Password.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown Casaos API Password",
			"The provider cannot create the Casaos API client as there is an unknown configuration value for the Casaos API password. "+
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

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing Casaos API Host",
			"The provider cannot create the Casaos API client as there is a missing or empty value for the Casaos API host. "+
				"Set the host value in the configuration or use the CASAOS_HOST environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if username == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Missing Casaos API Username",
			"The provider cannot create the Casaos API client as there is a missing or empty value for the Casaos API username. "+
				"Set the username value in the configuration or use the CASAOS_USERNAME environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if password == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing Casaos API Password",
			"The provider cannot create the Casaos API client as there is a missing or empty value for the Casaos API password. "+
				"Set the password value in the configuration or use the CASAOS_PASSWORD environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Create a new Casaos client using the configuration values
	client, err := casaos.NewClient(&host, &username, &password)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Casaos API Client",
			"An unexpected error occurred when creating the Casaos API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Casaos Client Error: "+err.Error(),
		)
		return
	}

	// Make the Casaos client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = client
	resp.ResourceData = client
}

// DataSources defines the data sources implemented in the provider.
func (p *casaosProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

// Resources defines the resources implemented in the provider.
func (p *casaosProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}
