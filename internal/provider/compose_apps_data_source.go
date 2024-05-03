package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/rstuhlmuller/terraform-provider-casaos/internal/casaos"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &composeAppsDataSource{}
	_ datasource.DataSourceWithConfigure = &composeAppsDataSource{}
)

// NewcomposeAppsDataSource is a helper function to simplify the provider implementation.
func NewcomposeAppsDataSource() datasource.DataSource {
	return &composeAppsDataSource{}
}

// composeAppsDataSource is the data source implementation.
type composeAppsDataSource struct {
	client *casaos.Client
}

// Metadata returns the data source type name.
func (d *composeAppsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_compose_apps"
}

// Schema defines the schema for the data source.
func (d *composeAppsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

// Read refreshes the Terraform state with the latest data.
func (d *composeAppsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
}

// Configure adds the provider configured client to the data source.
func (d *composeAppsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*casaos.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *hashicups.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = client
}
