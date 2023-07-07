package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tatari-tv/terraform-provider-altinitycloud/cmd/client"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &usersDataSource{}
	_ datasource.DataSourceWithConfigure = &usersDataSource{}
)

func NewUsersDataSource() datasource.DataSource {
	return &usersDataSource{}
}

// nodeTypeDataSource - defines the Data source implementation.
type usersDataSource struct {
	client *client.AltinityCloudClient
}

// Metadata - returns the altinitycloud_users type name.
func (d *usersDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_users"
}

// Schema - defines the node_type schema.
func (d *usersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"cluster_id": schema.StringAttribute{
				MarkdownDescription: "Altinity.Cloud environment ID",
				Required:            true,
			},
			"users": schema.ListNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Altinity.Cloud Users.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Altinity.Cloud user ID.",
						},
						"login": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Altinity.Cloud user login.",
						},
						"password": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Altinity.Cloud user password.",
						},
						"networks": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Altinity.Cloud user networks (string representation e.g. 0.0.0.0/0\\n::/0).",
						},
						"databases": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Altinity.Cloud user databases (string representation e.g. default\\nanother_db).",
						},
						"profile_id": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Altinity.Cloud user profile ID.",
						},
						"quota_id": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "Altinity.Cloud user quota ID.",
						},
						"access_management": schema.BoolAttribute{
							Computed:            true,
							MarkdownDescription: "Altinity.Cloud user admin privilege.",
						},
						"system": schema.BoolAttribute{
							Computed:            true,
							MarkdownDescription: "Altinity.Cloud system user (e.g. operation ).",
						},
					},
				},
			},
		},
	}
}

// Configure - bootstraps users datasource with Altinity.Cloud client.
func (d *usersDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring users data source")
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.AltinityCloudClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *altinityCloudClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

// Read - implements the read functionality for users.
func (d *usersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Info(ctx, "Reading users state source")
	var state UsersDataSourceModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, fmt.Sprintf("><>>>>>>>>>>> %v", state.ClusterID.ValueString()))
	// initialize provider client state and make a call using it.
	users, err := d.client.GetUsers(state.ClusterID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
		return
	}

	for _, user := range users.Data {
		state.Users = append(state.Users, UserModel{
			ID:               types.StringValue(user.ID),
			Login:            types.StringValue(user.Login),
			Password:         types.StringValue(user.Password),
			Networks:         types.StringValue(user.Networks),
			Databases:        types.StringValue(user.Databases),
			ProfileID:        types.StringValue(user.IDProfile),
			QuotaID:          types.StringValue(user.IDQuota),
			AccessManagement: types.BoolValue(user.AccessManagement),
			System:           types.BoolValue(user.System),
		})
	}

	tflog.Trace(ctx, fmt.Sprintf("fetch users from Altinity.Cloud API in cluster %v", state.ClusterID.ValueString()))

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
