// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"crypto/sha1"
	"encoding/base64"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = Base64Sha1{}
)

func NewBase64Sha1Function() function.Function {
	return Base64Sha1{}
}

type Base64Sha1 struct{}

func (r Base64Sha1) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "base64sha1"
}

func (r Base64Sha1) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Compute the SHA-1 hash of a given string and encode it with Base64",
		MarkdownDescription: "Compute the SHA-1 hash of a given string and encode it with Base64.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "input",
				MarkdownDescription: "String to compute the SHA1 hash of",
			},
		},
		Return: function.StringReturn{},
	}
}

func (r Base64Sha1) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &input))

	if resp.Error != nil {
		return
	}

	sha1sum := sha1.Sum([]byte(input))
	output := base64.StdEncoding.EncodeToString(sha1sum[:])

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, output))
}
