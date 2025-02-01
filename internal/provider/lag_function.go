// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = LagFunction{}
)

func NewLagFunction() function.Function {
	return LagFunction{}
}

type LagFunction struct{}

func (r LagFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "lag"
}

func (r LagFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Lag function",
		MarkdownDescription: "Echos the given input after a delay.",
		Parameters: []function.Parameter{
			function.Int64Parameter{
				AllowUnknownValues:  false,
				AllowNullValue:      false,
				MarkdownDescription: "Amount of time in milliseconds to delay before function returns",
				Name:                "delay",
			},
			function.StringParameter{
				AllowUnknownValues:  false,
				AllowNullValue:      false,
				Name:                "input",
				MarkdownDescription: "String to echo",
			},
		},
		Return: function.StringReturn{},
	}
}

func (r LagFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var delay int64
	var input string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &delay, &input))

	if resp.Error != nil {
		return
	}

	if delay > 0 {
		id := uuid.New().String()

		startMessage := fmt.Sprintf("Lag Function (%s): Start sleeping for %d seconds...\n", id, delay)
		tflog.Trace(ctx, startMessage)

		time.Sleep(time.Duration(delay) * time.Millisecond)

		finishMessage := fmt.Sprintf("Lag Function (%s): Finished sleeping for %d seconds...\n", id, delay)
		tflog.Trace(ctx, finishMessage)
	}
	result := input

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}
