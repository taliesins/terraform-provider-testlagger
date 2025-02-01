package provider

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"time"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &LagDataSource{}

func NewLagDataSource() datasource.DataSource {
	return &LagDataSource{}
}

type LagDataSource struct {
	Id     string
	client *TestLaggerClient
}

type lagDataSourceModel struct {
	ReadDelay types.Int64  `tfsdk:"read_delay"`
	Input     types.String `tfsdk:"input"`
	Output    types.String `tfsdk:"output"`
}

func (d *LagDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lag"
}

func (d *LagDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Returns the given input after a delay.",
		Attributes: map[string]schema.Attribute{
			"read_delay": schema.Int64Attribute{
				MarkdownDescription: "Amount of time to delay before read function returns",
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

func (d *LagDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*TestLaggerClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *TestLaggerClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	id := uuid.New().String()

	if client.DatasourceConfigureDelay > 0 {
		startMessage := fmt.Sprintf("Datasource Lag Configure (%s/%s): Start sleeping for %d seconds...\n", client.Id, id, client.DatasourceConfigureDelay)
		tflog.Trace(ctx, startMessage)

		time.Sleep(time.Duration(client.DatasourceConfigureDelay) * time.Millisecond)

		finishMessage := fmt.Sprintf("Datasource Lag Configure (%s/%s): Finished sleeping for %d seconds...\n", client.Id, id, client.DatasourceConfigureDelay)
		tflog.Trace(ctx, finishMessage)
	}

	d.client = client
	d.Id = id
}

func (d *LagDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data lagDataSourceModel

	// Read configuration data into the model
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read input values
	var readDelay int64
	var input string

	if data.ReadDelay.IsNull() || data.ReadDelay.IsUnknown() {
		readDelay = 0
	} else {
		readDelay = data.ReadDelay.ValueInt64()
	}

	if data.Input.IsNull() || data.Input.IsUnknown() {
		input = ""
	} else {
		input = data.Input.ValueString()
	}

	if readDelay > 0 {
		startMessage := fmt.Sprintf("Datasource Lag Read (%s/%s): Start sleeping for %d seconds...\n", d.client.Id, d.Id, readDelay)
		tflog.Trace(ctx, startMessage)

		time.Sleep(time.Duration(readDelay) * time.Millisecond)

		finishMessage := fmt.Sprintf("Datasource Lag Read (%s/%s): Finished sleeping for %d seconds...\n", d.client.Id, d.Id, readDelay)
		tflog.Trace(ctx, finishMessage)
	}

	// Set output values
	data.Output = types.StringValue(input)

	// Save updated data into Terraform state
	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}
