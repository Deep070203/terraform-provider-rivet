package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"rivet/blob/main/sdks/api/full/go/client"
	http "net/http"
)

var _ provider.Provider = &RivetProvider{}

type RivetProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// RivetProviderModel describes the provider data model.
type RivetProviderModel struct {
	ServerURL types.String `tfsdk:"server_url"`
}

func (p *RivetProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "rivet"
	resp.Version = p.version
}

func (p *RivetProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: Rivet API generated from the OpenAPI specification.,
		Attributes: map[string]schema.Attribute{
			"server_url": schema.StringAttribute{
				MarkdownDescription: "Server URL",
				Optional:            true,
				Required:            false,
			},
		},
	}
}

func (p *RivetProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data RivetProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ServerURL := data.ServerURL.ValueString()

	if ServerURL == "" {
		ServerURL = "https://api.rivet.gg"
	}

	var ClientOptions := {
		BaseURL: ServerURL,
		HTTPClient: http.DefaultClient,
	}

	client, err := client.NewClient(ClientOptions)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *RivetProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		
		NewActorsResource,
		
	}
}

func (p *RivetProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		
		NewActorsDataSource,
		
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &RivetProvider{
			version: version,
		}
	}
}