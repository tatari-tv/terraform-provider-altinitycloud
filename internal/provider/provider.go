package provider

import (
	"context"
	"net/http"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &altinityCloudProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &altinityCloudProvider{
			version: version,
		}
	}
}

// altinityCloudClient is wrapper for http client and configs.
type altinityCloudClient struct {
	client      *http.Client
	APIEndpoint string
	APIToken    string
}

// altinityCloudProvider is the provider implementation.
type altinityCloudProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// altinityCloudProviderModel maps provider schema data to a Go type.
type altinityCloudProviderModel struct {
	APIEndpoint types.String `tfsdk:"api_endpoint"`
	APIToken    types.String `tfsdk:"api_token"`
}

// Metadata returns the provider type name.
func (p *altinityCloudProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "altinitycloud"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
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

func (p *altinityCloudProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
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
		endpoint = config.APIToken.ValueString()
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

	client := altinityCloudClient{
		client:      http.DefaultClient,
		APIEndpoint: endpoint,
		APIToken:    token,
	}

	// Make the Altinity.Cloud client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = client
	resp.ResourceData = client
}

// DataSources defines the data sources implemented in the provider.
func (p *altinityCloudProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

// Resources defines the resources implemented in the provider.
func (p *altinityCloudProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}
