package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &NodeTypeDataSource{}

func NewExampleDataSource() datasource.DataSource {
	return &NodeTypeDataSource{}
}

// NodeTypeDataSource defines the data source implementation.
type NodeTypeDataSource struct {
	client *altinityCloudClient
}

// NodeTypeDataSourceModel describes the data source data model.
type NodeTypeDataSourceModel struct {
	envId types.String `tfsdk:"env_id"`
	Id    types.String `tfsdk:"id"`
}

type NodeTypes struct {
	Data []struct {
		ID            string `json:"id"`
		Scope         string `json:"scope"`
		Code          string `json:"code"`
		Name          string `json:"name"`
		Pool          string `json:"pool"`
		StorageClass  string `json:"storageClass"`
		CPU           string `json:"cpu"`
		Memory        string `json:"memory"`
		IDEnvironment string `json:"id_environment"`
		ExtraSpec     string `json:"extraSpec"`
		Tolerations   []struct {
			Key      string `json:"key"`
			Operator string `json:"operator"`
			Effect   string `json:"effect"`
			Value    string `json:"value"`
		} `json:"tolerations"`
		NodeSelector string `json:"nodeSelector"`
		CPUAlloc     string `json:"cpu_alloc"`
		MemoryAlloc  string `json:"memory_alloc"`
	} `json:"data"`
}

func (d *NodeTypeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_node_type"
}

func (d *NodeTypeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Altinity.Cloud Node Type data source",

		Attributes: map[string]schema.Attribute{
			"env_id": schema.StringAttribute{
				MarkdownDescription: "Altinity.Cloud environment ID",
				Optional:            false,
			},
		},
	}

	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"env_id": schema.StringAttribute{
				MarkdownDescription: "Altinity.Cloud environment ID",
				Optional:            false,
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
						"nodeSelector": schema.StringAttribute{
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

func (d *NodeTypeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*altinityCloudClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *altinityCloudClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *NodeTypeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NodeTypeDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := d.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
	//     return
	// }

	// For the purposes of this example code, hardcoding a response value to
	// save into the Terraform state.
	data.Id = types.StringValue("example-id")

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
