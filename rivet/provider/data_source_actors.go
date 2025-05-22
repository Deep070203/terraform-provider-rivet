package provider

import (
	"context"
	"fmt"
	"github.com/UPDATE_PATH_TO_MODELS/models"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"https://github.com/rivet-gg/rivet/blob/main/sdks/api/runtime/go/client"
	"https://github.com/rivet-gg/rivet/blob/main/sdks/api/runtime/go/actors/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &ActorsDataSource{}
var _ datasource.DataSourceWithConfigure = &ActorsDataSource{}

func NewActorsDataSource() datasource.DataSource {
	return &ActorsDataSource{}
}

// ActorsDataSource is the data source implementation.
type ActorsDataSource struct {
	client *Client.Actors
}

// ActorsDataSourceModel describes the data model.
type ActorsDataSourceModel struct {
	Actor  types.String  `tfsdk:"actor"`
    Project  types.String  `tfsdk:"project"`
    Environment  types.String  `tfsdk:"environment"`
    EndpointType  models.actorsendpointtype  `tfsdk:"endpoint_type"`
    TagsJson  types.String  `tfsdk:"tags_json"`
    IncludeDestroyed  types.Bool  `tfsdk:"include_destroyed"`
    Cursor  types.String  `tfsdk:"cursor"`
    Actors  types.Array  `tfsdk:"actors"`
    Pagination  models.pagination  `tfsdk:"pagination"`
    
}

// Metadata returns the data source type name.
func (r *ActorsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_Actors"
}

// Schema defines the schema for the data source.
func (r *ActorsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Actors DataSource",

		Attributes: map[string]schema.Attribute{
			"actor": schema.StringAttribute{
				Required: true,
			},
			"project": schema.StringAttribute{
				Optional: true,
			},
			"environment": schema.StringAttribute{
				Optional: true,
			},
			"endpoint_type": schema.SingleNestedAttribute{
				Optional: true,
				schema.TypeAttribute{
				Required: true,
			}
			},
			"tags_json": schema.StringAttribute{
				Optional: true,
			},
			"include_destroyed": schema.BoolAttribute{
				Optional: true,
			},
			"cursor": schema.StringAttribute{
				Optional: true,
			},
			"actors": schema.ArrayAttribute{
				Required: true,
			},
			"pagination": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					cursor: schema.StringAttribute{
						Optional: true,
					}
				}
			},
			
		},
	}
}

func (r *ActorsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*Client.Actors)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *Client.Actors, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}


func (r *ActorsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	
	var data *ActorsResourceModel
	
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	
    actor := data.Actor.ValueString()
    
    project := new(string)
	if !data.Project.IsUnknown() && !data.Project.IsNull() {
		*project = data.Project.ValueString()
	} else {
		project = nil
	}
    
    environment := new(string)
	if !data.Environment.IsUnknown() && !data.Environment.IsNull() {
		*environment = data.Environment.ValueString()
	} else {
		environment = nil
	}
    

	request := actors.ActorsGetRequest{
		Actor: actor,
		Project: project,
		Environment: environment,
		
	}
	res, err := r.client.Get(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	
}

func (r *ActorsResourceModel) RefreshFromSharedActorsGetResponse(resp *shared.ActorsGetActorResponse) {
	if resp != nil {
		actor = types.ActorsactorValue(resp.Actor)
		
	}
}