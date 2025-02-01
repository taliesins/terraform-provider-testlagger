package provider

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"time"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &LagResource{}
var _ resource.ResourceWithImportState = &LagResource{}

func NewLagResource() resource.Resource {
	return &LagResource{}
}

type LagResource struct {
	Id     string
	client *TestLaggerClient
}

type LagResourceModel struct {
	Id          types.String `tfsdk:"id"`
	CreateDelay types.Int64  `tfsdk:"create_delay"`
	ReadDelay   types.Int64  `tfsdk:"read_delay"`
	UpdateDelay types.Int64  `tfsdk:"update_delay"`
	DeleteDelay types.Int64  `tfsdk:"delete_delay"`
	Input       types.String `tfsdk:"input"`
	Output      types.String `tfsdk:"output"`
}

func (r *LagResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lag"
}

func (r *LagResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Echos the given input after a delay.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Unique identifier",
				Computed:            true,
			},
			"create_delay": schema.Int64Attribute{
				MarkdownDescription: "Amount of time in milliseconds to delay before create function returns",
				Optional:            true,
			},
			"read_delay": schema.Int64Attribute{
				MarkdownDescription: "Amount of time in milliseconds to delay before read function returns",
				Optional:            true,
			},
			"update_delay": schema.Int64Attribute{
				MarkdownDescription: "Amount of time in milliseconds to delay before update function returns",
				Optional:            true,
			},
			"delete_delay": schema.Int64Attribute{
				MarkdownDescription: "Amount of time in milliseconds to delay before delete function returns",
				Optional:            true,
			},
			"input": schema.StringAttribute{
				MarkdownDescription: "Input string to echo",
				Required:            true,
			},
			"output": schema.StringAttribute{
				MarkdownDescription: "Output string echoed",
				Computed:            true,
			},
		},
	}
}

func (r *LagResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*TestLaggerClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *TestLaggerClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	id := uuid.New().String()

	// Client does work to initialize
	if client.ResourceConfigureDelay > 0 {
		startMessage := fmt.Sprintf("Resource Lag Configure (%s/%s): Start sleeping for %d seconds...\n", client.Id, id, client.ResourceConfigureDelay)
		tflog.Trace(ctx, startMessage)

		time.Sleep(time.Duration(client.ResourceConfigureDelay) * time.Millisecond)

		finishMessage := fmt.Sprintf("Resource Lag Configure (%s/%s): Finished sleeping for %d seconds...\n", client.Id, id, client.ResourceConfigureDelay)
		tflog.Trace(ctx, finishMessage)
	}

	r.client = client
	r.Id = id
}

func (r *LagResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plannedState LagResourceModel

	// Read Terraform plan plannedState into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plannedState)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read input values
	var createDelay int64
	var input string

	if plannedState.CreateDelay.IsNull() || plannedState.CreateDelay.IsUnknown() {
		createDelay = 0
	} else {
		createDelay = plannedState.CreateDelay.ValueInt64()
	}

	if plannedState.Input.IsNull() || plannedState.Input.IsUnknown() {
		input = ""
	} else {
		input = plannedState.Input.ValueString()
	}

	// Client does work against API
	if createDelay > 0 {
		startMessage := fmt.Sprintf("Resource Lag Create (%s/%s): Start sleeping for %d seconds...\n", r.client.Id, r.Id, createDelay)
		tflog.Trace(ctx, startMessage)

		time.Sleep(time.Duration(createDelay) * time.Millisecond)

		finishMessage := fmt.Sprintf("Resource Lag Create (%s/%s): Finished sleeping for %d seconds...\n", r.client.Id, r.Id, createDelay)
		tflog.Trace(ctx, finishMessage)
	}

	// Set state
	plannedState.Output = types.StringValue(input)
	plannedState.Id = types.StringValue(input)

	// Save plannedState into Terraform state
	diags := resp.State.Set(ctx, &plannedState)
	resp.Diagnostics.Append(diags...)
}

func (r *LagResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state LagResourceModel

	// Read state state into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read input values
	var readDelay int64

	if state.ReadDelay.IsNull() || state.ReadDelay.IsUnknown() {
		readDelay = 0
	} else {
		readDelay = state.ReadDelay.ValueInt64()
	}

	// Client does work against API
	if readDelay > 0 {
		startMessage := fmt.Sprintf("Resource Lag Read (%s/%s): Start sleeping for %d seconds...\n", r.client.Id, r.Id, readDelay)
		tflog.Trace(ctx, startMessage)

		time.Sleep(time.Duration(readDelay) * time.Millisecond)

		finishMessage := fmt.Sprintf("Resource Lag Read (%s/%s): Finished sleeping for %d seconds...\n", r.client.Id, r.Id, readDelay)
		tflog.Trace(ctx, finishMessage)
	}

	// Save updated state into Terraform state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *LagResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Read Terraform plan plannedState into the model
	var state LagResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var plannedState LagResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plannedState)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read input values
	var updateDelay int64
	var input string

	if plannedState.UpdateDelay.IsNull() || plannedState.UpdateDelay.IsUnknown() {
		updateDelay = 0
	} else {
		updateDelay = plannedState.UpdateDelay.ValueInt64()
	}

	if plannedState.Input.IsNull() || plannedState.Input.IsUnknown() {
		input = ""
	} else {
		input = plannedState.Input.ValueString()
	}

	// Client does work against API
	if updateDelay > 0 {
		startMessage := fmt.Sprintf("Resource Lag Update (%s/%s): Start sleeping for %d seconds...\n", r.client.Id, r.Id, updateDelay)
		tflog.Trace(ctx, startMessage)

		time.Sleep(time.Duration(updateDelay) * time.Millisecond)

		finishMessage := fmt.Sprintf("Resource Lag Update (%s/%s): Finished sleeping for %d seconds...\n", r.client.Id, r.Id, updateDelay)
		tflog.Trace(ctx, finishMessage)
	}

	// Set state
	state.Id = types.StringValue(input)
	state.Input = plannedState.Input
	state.Output = types.StringValue(input)
	state.CreateDelay = plannedState.CreateDelay
	state.ReadDelay = plannedState.ReadDelay
	state.UpdateDelay = plannedState.UpdateDelay
	state.DeleteDelay = plannedState.DeleteDelay

	// Save updated plannedState into Terraform state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *LagResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LagResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read input values
	var deleteDelay int64

	if data.DeleteDelay.IsNull() || data.DeleteDelay.IsUnknown() {
		deleteDelay = 0
	} else {
		deleteDelay = data.DeleteDelay.ValueInt64()
	}

	// Client does work against API
	if deleteDelay > 0 {
		startMessage := fmt.Sprintf("Resource Lag Delete (%s/%s): Start sleeping for %d seconds...\n", r.client.Id, r.Id, deleteDelay)
		tflog.Trace(ctx, startMessage)

		time.Sleep(time.Duration(deleteDelay) * time.Millisecond)

		finishMessage := fmt.Sprintf("Resource Lag Delete (%s/%s): Finished sleeping for %d seconds...\n", r.client.Id, r.Id, deleteDelay)
		tflog.Trace(ctx, finishMessage)
	}
}

func (r *LagResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	// Client does work against API
	if r.client.ResourceImportStateDelay > 0 {
		startMessage := fmt.Sprintf("Resource Lag Import State (%s/%s): Start sleeping for %d seconds...\n", r.client.Id, r.Id, r.client.ResourceImportStateDelay)
		tflog.Trace(ctx, startMessage)

		time.Sleep(time.Duration(r.client.ResourceConfigureDelay) * time.Millisecond)

		finishedMessage := fmt.Sprintf("Resource Lag Import State (%s/%s): Finished sleeping for %d seconds...\n", r.client.Id, r.Id, r.client.ResourceImportStateDelay)
		tflog.Trace(ctx, finishedMessage)
	}

	model := &LagResourceModel{
		Id:     types.StringValue(req.ID),
		Input:  types.StringValue(req.ID),
		Output: types.StringValue(req.ID),
	}

	resp.State.Set(ctx, model)
}
