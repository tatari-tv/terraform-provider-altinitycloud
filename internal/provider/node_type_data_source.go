package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tatari-tv/terraform-provider-altinitycloud/cmd/client"
	"strconv"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &nodeTypeDataSource{}
	_ datasource.DataSourceWithConfigure = &nodeTypeDataSource{}
)

func NewNodeTypesDataSource() datasource.DataSource {
	return &nodeTypeDataSource{}
}

// nodeTypeDataSource defines the Data source implementation.
type nodeTypeDataSource struct {
	client *client.AltinityCloudClient
}

// Metadata - returns the altinitycloud_node_type type name.
func (d *nodeTypeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_node_type"
}

// Schema - defines the node_type schema.
func (d *nodeTypeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Altinity.Cloud environment ID",
				Optional:            true,
			},
			"env_id": schema.StringAttribute{
				MarkdownDescription: "Altinity.Cloud environment ID",
				Required:            true,
			},
			"data": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed: true,
						},
						"scope": schema.StringAttribute{
							Computed: true,
						},
						"code": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"pool": schema.StringAttribute{
							Computed: true,
						},
						"storage_class": schema.StringAttribute{
							Computed: true,
						},
						"cpu": schema.StringAttribute{
							Computed: true,
						},
						"memory": schema.StringAttribute{
							Computed: true,
						},
						"id_environment": schema.StringAttribute{
							Computed: true,
						},
						"extra_spec": schema.StringAttribute{
							Computed: true,
						},
						"tolerations": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Computed: true,
									},
									"operator": schema.StringAttribute{
										Computed: true,
									},
									"effect": schema.StringAttribute{
										Computed: true,
									},
									"value": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
						"node_selector": schema.StringAttribute{
							Computed: true,
						},
						"cpu_alloc": schema.StringAttribute{
							Computed: true,
						},
						"memory_alloc": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

// Configure - bootstraps node type datasource with Altinity.Cloud client.
func (d *nodeTypeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

// Read - implements the read functionality for node type.
func (d *nodeTypeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state NodeTypeDataSourceModel

	// Read Terraform configuration state into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// set an ID
	state.Id = types.StringValue("node_type_" + state.EnvId.ValueString())

	// initialize provider client state and make a call using it.
	nodeTypes, err := d.client.GetNodeTypes(state.EnvId.String())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
		return
	}

	for _, nt := range nodeTypes.Data {
		id, err := strconv.ParseInt(nt.ID, 10, 64)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("fetch node types from Altinity.Cloud API in environment %v", err))
		}
		nodeTypeState := NodeTypeModel{
			ID:            types.Int64Value(id),
			Scope:         types.StringValue(nt.Scope),
			Code:          types.StringValue(nt.Code),
			Name:          types.StringValue(nt.Name),
			Pool:          types.StringValue(nt.Pool),
			StorageClass:  types.StringValue(nt.StorageClass),
			CPU:           types.StringValue(nt.CPU),
			Memory:        types.StringValue(nt.Memory),
			IDEnvironment: types.StringValue(nt.IDEnvironment),
			ExtraSpec:     types.StringValue(nt.ExtraSpec),
			NodeSelector:  types.StringValue(nt.NodeSelector),
			CPUAlloc:      types.StringValue(nt.CPUAlloc),
			MemoryAlloc:   types.StringValue(nt.MemoryAlloc),
		}
		for _, t := range nt.Tolerations {
			nodeTypeState.Tolerations = append(nodeTypeState.Tolerations, NodeTypeTolerationModel{
				Key:      types.StringValue(t.Key),
				Operator: types.StringValue(t.Operator),
				Effect:   types.StringValue(t.Effect),
				Value:    types.StringValue(t.Value),
			})
		}

		state.Data = append(state.Data, nodeTypeState)
	}

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, fmt.Sprintf("fetch node types from Altinity.Cloud API in environment %v", state.EnvId))

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
