package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/UPDATE_PATH_TO_MODELS/models"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"https://github.com/rivet-gg/rivet/blob/main/sdks/api/runtime/go/client"
	"https://github.com/rivet-gg/rivet/blob/main/sdks/api/runtime/go/actors/client"
	"regexp"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ActorsResource{}
var _ resource.ResourceWithImportState = &ActorsResource{}

func NewActorsResource() resource.Resource {
	return &ActorsResource{}
}

// ActorsResource defines the resource implementation.
type ActorsResource struct {
	client *Client.Actors
}

// ActorsResourceModel describes the resource data model.
type ActorsResourceModel struct {
    Actor  types.String  `tfsdk:"actor"`
    Project  types.String  `tfsdk:"project"`
    Environment  types.String  `tfsdk:"environment"`
    EndpointType  models.actorsendpointtype  `tfsdk:"endpoint_type"`
    Region  types.String  `tfsdk:"region"`
    Tags  interface{}  `tfsdk:"tags"`
    Build  types.String  `tfsdk:"build"`
    BuildTags  interface{}  `tfsdk:"build_tags"`
    Runtime  models.actorscreateactorruntimerequest  `tfsdk:"runtime"`
    Network  models.actorscreateactornetworkrequest  `tfsdk:"network"`
    Resources  models.actorsresources  `tfsdk:"resources"`
    Lifecycle  models.actorslifecycle  `tfsdk:"lifecycle"`
    OverrideKillTimeout  types.Int64  `tfsdk:"override_kill_timeout"`
    
}

func (r *ActorsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_Actors"
}

func (r *ActorsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Actors Resource",
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
			"region": schema.StringAttribute{
				Optional: true,
			},
			"tags": schema.SingleNestedAttribute{
				Required: true,
				interface{}
			},
			"build": schema.StringAttribute{
				Optional: true,
			},
			"build_tags": schema.SingleNestedAttribute{
				Optional: true,
				interface{}
			},
			"runtime": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					environment: schema.ObjectAttribute{
						Optional: true,
					},
					network: schema.SingleNestedAttribute{
						Optional: true,
					Attributes: map[string]schema.Attribute{
					
				}
					}
				}
			},
			"network": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					ports: schema.ObjectAttribute{
						Optional: true,
					}
				}
			},
			"resources": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					cpu: schema.Int64Attribute{
						Optional: true,
					},
					memory: schema.Int64Attribute{
						Optional: true,
					}
				}
			},
			"lifecycle": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					kill_timeout: schema.Int64Attribute{
						Optional: true,
					},
					durable: schema.BoolAttribute{
						Optional: true,
					}
				}
			},
			"override_kill_timeout": schema.Int64Attribute{
				Optional: true,
			},
			
		},
	}
}

func (r *ActorsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*Client.Actors)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *Client.Actors, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ActorsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	
	var data *ActorsResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	
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
	
	
	req := *data.ToSharedActorsCreateActorRequestRequest()
	request := actors.ActorsCreateRequest{
		Project: project,
		Environment: environment,
		
		ActorsCreateActorRequest: req,
	}
	res, err := r.client.Create(ctx, request)
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
	data.actor = res.GetActor()
	
	data.Actors = res.GetActors.GetId()

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	
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


func (r *ActorsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	
	var data *ActorsResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	merge(ctx, req, resp, &data)

	
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
    
	req := *data.ToSharedActorsUpgradeActorRequestRequest()
	request := actors.ActorsUpgradeRequest{
		Actor: actor,
		Project: project,
		Environment: environment,
		
		ActorsUpgradeActorRequest: req,
	}
	res, err := r.client.Upgrade(ctx, request)
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
	
	data.Actors = res.GetActors.GetId()

    // Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	
}


func (r *ActorsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	
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
    
    override_kill_timeout := new(integer)
	if !data.Override_kill_timeout.IsUnknown() && !data.Override_kill_timeout.IsNull() {
		*override_kill_timeout = data.Override_kill_timeout.ValueInteger()
	} else {
		override_kill_timeout = nil
	}
    

	request := actors.ActorsDestroyRequest{
		Actor: actor,
		Actor: actor,
		Project: project,
		Environment: environment,
		Override_kill_timeout: override_kill_timeout,
		
	}
	res, err := r.client.Destroy(ctx, request)
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
	
}

func (r *IndexResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ActorsResourceModel) ToSharedActorsCreateActorRequestRequest() *models.ActorsCreateActorRequest {
	region := new(string)
	if !data.Region.IsUnknown() && !data.Region.IsNull() {
		*region = data.Region.ValueString()
	} else {
		region = nil
	}
	tags := interface{}
	build := new(string)
	if !data.Build.IsUnknown() && !data.Build.IsNull() {
		*build = data.Build.ValueString()
	} else {
		build = nil
	}
	build_tags := interface{}
	environment := data.Environment.ValueObject()
	runtime := models.actorscreateactorruntimerequest{
		Environment: environment,
	}

	ports := data.Ports.ValueObject()
	network := models.actorscreateactornetworkrequest{
		Ports: ports,
	}

	cpu := data.Cpu.ValueInt64()
	memory := data.Memory.ValueInt64()
	resources := models.actorsresources{
		Cpu: cpu,
		Memory: memory,
	}

	kill_timeout := data.Kill_timeout.ValueInt64()
	durable := data.Durable.ValueBool()
	lifecycle := models.actorslifecycle{
		Kill_timeout: kill_timeout,
		Durable: durable,
	}

	
	out := models.ActorsCreateActorRequest{
		Region: region,
        Tags: tags,
        Build: build,
        Build_tags: build_tags,
        Runtime: runtime,
        Network: network,
        Resources: resources,
        Lifecycle: lifecycle,
        
	}
	return &out
}

func (r *ActorsResourceModel) ToSharedActorsUpgradeActorRequestRequest() *models.ActorsUpgradeActorRequest {
	build := new(string)
	if !data.Build.IsUnknown() && !data.Build.IsNull() {
		*build = data.Build.ValueString()
	} else {
		build = nil
	}
	build_tags := interface{}
	
	out := models.ActorsUpgradeActorRequest{
		Build: build,
        Build_tags: build_tags,
        
	}
	return &out
}
