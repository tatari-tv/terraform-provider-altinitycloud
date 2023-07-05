package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tatari-tv/terraform-provider-altinitycloud/cmd/client"
	"os"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &altinityCloudProvider{}
)

// New - helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &altinityCloudProvider{
			version: version,
		}
	}
}

// altinityCloudProvider - provider implementation.
type altinityCloudProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// altinityCloudProviderModel - maps provider schema NodeTypes to a Go type.
type altinityCloudProviderModel struct {
	APIEndpoint types.String `tfsdk:"api_endpoint"`
	APIToken    types.String `tfsdk:"api_token"`
}

// Metadata - returns the provider type name.
func (p *altinityCloudProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "altinitycloud"
	resp.Version = p.version
}

// Schema - defines the provider-level schema for configuration NodeTypes.
func (p *altinityCloudProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_endpoint": schema.StringAttribute{
				Optional: true,
			},
			"api_token": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

// Configure - bootstraps provider with configurations needed to run.
func (p *altinityCloudProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Altinity.Cloud client")
	// Retrieve provider NodeTypes from configuration
	var config altinityCloudProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	if config.APIEndpoint.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_endpoint"),
			"Unknown Altinit.Cloud API Endpoint",
			"The provider cannot create the Altinity.Cloud API client as there is an unknown configuration value for the API Endpoint. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the ALTINITY_CLOUD_ENDPOINT environment variable.",
		)
	}

	if config.APIToken.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_token"),
			"Unknown Altinity.Cloud API Token",
			"The provider cannot create the Altinity.Cloud API client as there is an unknown configuration value for the Altinity.Cloud API token. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the ALTINITY_CLOUD_TOKEN environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	endpoint := os.Getenv("ALTINITY_CLOUD_ENDPOINT")
	token := os.Getenv("ALTINITY_CLOUD_TOKEN")

	if !config.APIEndpoint.IsNull() {
		endpoint = config.APIEndpoint.ValueString()
	}

	if !config.APIToken.IsNull() {
		token = config.APIToken.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if endpoint == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_endpoint"),
			"Missing Altinity.Cloud API Endpoint",
			"The provider cannot create the Altinity.Cloud API client as there is a missing or empty value for the API endpoint. "+
				"Set the host value in the configuration or use the ALTINITY_CLOUD_ENDPOINT environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if token == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_token"),
			"Missing Altinity.Cloud API token",
			"The provider cannot create the Altinity.Cloud API client as there is a missing or empty value for the Altinity.Cloud API token. "+
				"Set the username value in the configuration or use the ALTINITY_CLOUD_TOKEN environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "api_endpoint", endpoint)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "api_token", token)

	tflog.Debug(ctx, "Creating Altiniy.Cloud client")

	// Create a new Altinity.Cloud client using the configuration values
	client, err := client.NewClient(&endpoint, &token)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Altinity.Cloud API Client",
			"An unexpected error occurred when creating the Altinity.Cloud API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Altinity.Cloud Client Error: "+err.Error(),
		)
		return
	}

	// Make the Altinity.Cloud client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = client
	resp.ResourceData = client
}

// DataSources - defines the NodeTypes sources implemented in the provider.
func (p *altinityCloudProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewNodeTypesDataSource,
	}
}

// Resources - defines the resources implemented in the provider.
func (p *altinityCloudProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewNodeTypeResource,
	}
}
