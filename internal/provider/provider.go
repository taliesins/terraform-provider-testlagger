// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"time"
)

// Ensure TestLaggerProvider satisfies various provider interfaces.
var _ provider.Provider = &TestLaggerProvider{}
var _ provider.ProviderWithFunctions = &TestLaggerProvider{}

// TestLaggerProvider defines the provider implementation.
type TestLaggerProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// TestLaggerProviderModel describes the provider data model.
type TestLaggerProviderModel struct {
	ClientInitializeDelay    types.Int64 `tfsdk:"client_initialize_delay"`
	DatasourceConfigureDelay types.Int64 `tfsdk:"datasource_configure_delay"`
	ResourceConfigureDelay   types.Int64 `tfsdk:"resource_configure_delay"`
	ResourceImportStateDelay types.Int64 `tfsdk:"resource_import_state_delay"`
}

func (p *TestLaggerProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "testlagger"
	resp.Version = p.version
}

func (p *TestLaggerProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"client_initialize_delay": schema.Int64Attribute{
				MarkdownDescription: "Amount of time to delay before client is created",
				Optional:            true,
			},
			"datasource_configure_delay": schema.Int64Attribute{
				MarkdownDescription: "Amount of time to delay before datasource configure function returns",
				Optional:            true,
			},
			"resource_configure_delay": schema.Int64Attribute{
				MarkdownDescription: "Amount of time to delay before resource configure function returns",
				Optional:            true,
			},
			"resource_import_state_delay": schema.Int64Attribute{
				MarkdownDescription: "Amount of time to delay before resource import state function returns",
				Optional:            true,
			},
		},
	}
}

type TestLaggerClient struct {
	Id                       string
	DatasourceConfigureDelay int64
	ResourceConfigureDelay   int64
	ResourceImportStateDelay int64
}

func (p *TestLaggerProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data TestLaggerProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := uuid.New().String()

	var clientInitializeDelay int64
	if !(data.ClientInitializeDelay.IsNull() || data.ClientInitializeDelay.IsUnknown()) {
		clientInitializeDelay = data.ClientInitializeDelay.ValueInt64()
	} else {
		clientInitializeDelay = 0
	}

	var datasourceConfigureDelay int64
	if !(data.DatasourceConfigureDelay.IsNull() || data.DatasourceConfigureDelay.IsUnknown()) {
		datasourceConfigureDelay = data.DatasourceConfigureDelay.ValueInt64()
	} else {
		datasourceConfigureDelay = 0
	}

	var resourceConfigureDelay int64
	if !(data.ResourceConfigureDelay.IsNull() || data.ResourceConfigureDelay.IsUnknown()) {
		resourceConfigureDelay = data.ResourceConfigureDelay.ValueInt64()
	} else {
		resourceConfigureDelay = 0
	}

	var resourceImportStateDelay int64
	if !(data.ResourceImportStateDelay.IsNull() || data.ResourceImportStateDelay.IsUnknown()) {
		resourceImportStateDelay = data.ResourceImportStateDelay.ValueInt64()
	} else {
		resourceImportStateDelay = 0
	}

	if clientInitializeDelay > 0 {
		startMessage := fmt.Sprintf("Provider Configure (%s): Start sleeping for %d seconds...", id, clientInitializeDelay)
		tflog.Trace(ctx, startMessage)

		time.Sleep(time.Duration(clientInitializeDelay) * time.Millisecond)

		finishMessage := fmt.Sprintf("Provider Configure (%s): Finished sleeping for %d seconds...", id, clientInitializeDelay)
		tflog.Trace(ctx, finishMessage)
	}

	client := &TestLaggerClient{
		Id:                       id,
		DatasourceConfigureDelay: datasourceConfigureDelay,
		ResourceConfigureDelay:   resourceConfigureDelay,
		ResourceImportStateDelay: resourceImportStateDelay,
	}

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *TestLaggerProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewLagResource,
	}
}

func (p *TestLaggerProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewLagDataSource,
	}
}

func (p *TestLaggerProvider) Functions(_ context.Context) []func() function.Function {
	return []func() function.Function{
		NewLagFunction,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &TestLaggerProvider{
			version: version,
		}
	}
}
