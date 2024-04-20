package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	casaos "github.com/rstuhlmuller/terraform-provider-casaos/internal/client"
)

var (
	_ datasource.DataSource              = &appManagementWebAppGridDataSource{}
	_ datasource.DataSourceWithConfigure = &appManagementWebAppGridDataSource{}
)

func NewAppManagementWebAppGridDataSource() datasource.DataSource {
	return &appManagementWebAppGridDataSource{}
}

type appManagementWebAppGridDataSource struct {
	client *casaos.Client
}

// Configure adds the provider configured client to the data source.
func (d *appManagementWebAppGridDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *appManagementWebAppGridDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app_management_web_app_grid"
}

func (d *appManagementWebAppGridDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"data": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"app_type": schema.StringAttribute{
							Description: "The type of the app.",
							Required:    true,
						},
						"author_type": schema.StringAttribute{
							Description: "The author of the app.",
							Required:    true,
						},
						"hostname": schema.StringAttribute{
							Description: "The hostname of the app.",
							Required:    true,
						},
						"icon": schema.StringAttribute{
							Description: "The icon of the app.",
							Required:    true,
						},
						"image": schema.StringAttribute{
							Description: "The image of the app.",
							Required:    true,
						},
						"index": schema.StringAttribute{
							Description: "The index of the app.",
							Required:    true,
						},
						"is_uncontrolled": schema.BoolAttribute{
							Description: "Whether the app is uncontrolled.",
							Required:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the app.",
							Required:    true,
						},
						"port": schema.StringAttribute{
							Description: "The port of the app.",
							Required:    true,
						},
						"scheme": schema.StringAttribute{
							Description: "The scheme of the app.",
							Required:    true,
						},
						"status": schema.StringAttribute{
							Description: "The status of the app.",
							Required:    true,
						},
						"store_app_id": schema.StringAttribute{
							Description: "The store app ID of the app.",
							Required:    true,
						},
						"title": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"custom": schema.StringAttribute{
										Description: "The custom title of the app.",
										Required:    true,
									},
									"en_us": schema.StringAttribute{
										Description: "The English title of the app.",
										Required:    true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// appManagementWebAppGridDataSourceModel maps the data source schema data.
type appManagementWebAppGridDataSourceModel struct {
	AppManagementWebAppGrid []appManagementWebAppGridModel `tfsdk:"data"`
}

// coffeesModel maps coffees schema data.
type appManagementWebAppGridModel struct {
	ID             types.String                        `tfsdk:"store_app_id"`
	AppType        types.String                        `tfsdk:"app_type"`
	AuthorType     types.String                        `tfsdk:"author_type"`
	Hostname       types.String                        `tfsdk:"hostname"`
	Icon           types.String                        `tfsdk:"icon"`
	Index          types.String                        `tfsdk:"index"`
	IsUncontrolled types.Bool                          `tfsdk:"is_uncontrolled"`
	Port           types.String                        `tfsdk:"port"`
	Scheme         types.String                        `tfsdk:"scheme"`
	Status         types.String                        `tfsdk:"status"`
	Title          []appManagementWebAppGridTitleModel `tfsdk:"title"`
}

type appManagementWebAppGridTitleModel struct {
	Custom types.String `tfsdk:"custom"`
	EnUS   types.String `tfsdk:"en_us"`
}

// coffeesIngredientsModel maps coffee ingredients data
type coffeesIngredientsModel struct {
	ID types.Int64 `tfsdk:"id"`
}

func (d *appManagementWebAppGridDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state appManagementWebAppGridDataSourceModel

	apps, err := GetAppManagementWebAppGrid()
	if err != nil {
		resp.Diagnostics.AddError("Unable to retrieve app management web app grid", err.Error())
		return
	}

	for _, app := range apps {
		AppManagementWebAppGridState := appManagementWebAppGridModel{
			ID:             types.StringValue(app.ID),
			AppType:        types.StringValue(app.AppType),
			AuthorType:     types.StringValue(app.AuthorType),
			Hostname:       types.StringValue(app.Hostname),
			Icon:           types.StringValue(app.Icon),
			Index:          types.StringValue(app.Index),
			IsUncontrolled: types.Bool(app.IsUncontrolled),
			Port:           types.StringValue(app.Port),
			Scheme:         types.StringValue(app.Scheme),
			Status:         types.StringValue(app.Status),
		}
		for _, title := range app.Title {
			AppManagementWebAppGridState.Title = append(AppManagementWebAppGridState.Title, appManagementWebAppGridTitleModel{
				Custom: types.StringValue(title.Custom),
				EnUS:   types.StringValue(title.EnUS),
			})
		}
		state.AppManagementWebAppGrid = append(state.AppManagementWebAppGrid, AppManagementWebAppGridState)
	}
	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func GetAppManagementWebAppGrid() ([]casaos.AppManagementWebAppGrid, error) {
	apps, err := casaos.GetAppManagementWebAppGrid()
	if err != nil {
		return nil, err
	}

	return apps, nil
}
