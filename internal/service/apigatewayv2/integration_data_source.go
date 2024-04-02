// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigatewayv2"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	"github.com/hashicorp/terraform-provider-aws/internal/framework"
	"github.com/hashicorp/terraform-provider-aws/internal/framework/flex"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// TIP: ==== FILE STRUCTURE ====
// All data sources should follow this basic outline. Improve this data source's
// maintainability by sticking to it.
//
// 1. Package declaration
// 2. Imports
// 3. Main data source struct with schema method
// 4. Read method
// 5. Other functions (flatteners, expanders, waiters, finders, etc.)

// @FrameworkDataSource(name="Integration")
func newDataSourceIntegration(context.Context) (datasource.DataSourceWithConfigure, error) {
	return &dataSourceIntegration{}, nil
}

const (
	DSNameIntegration = "Integration Data Source"
)

type dataSourceIntegration struct {
	framework.DataSourceWithConfigure
}

func (d *dataSourceIntegration) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) { // nosemgrep:ci.meta-in-func-name
	resp.TypeName = "aws_apigatewayv2_integration"
}

// TIP: ==== SCHEMA ====
// In the schema, add each of the arguments and attributes in snake
// case (e.g., delete_automated_backups).
// * Alphabetize arguments to make them easier to find.
// * Do not add a blank line between arguments/attributes.
//
// Users can configure argument values while attribute values cannot be
// configured and are used as output. Arguments have either:
// Required: true,
// Optional: true,
//
// All attributes will be computed and some arguments. If users will
// want to read updated information or detect drift for an argument,
// it should be computed:
// Computed: true,
//
// You will typically find arguments in the input struct
// (e.g., CreateDBInstanceInput) for the create operation. Sometimes
// they are only in the input struct (e.g., ModifyDBInstanceInput) for
// the modify operation.
//
// For more about schema options, visit
// https://developer.hashicorp.com/terraform/plugin/framework/handling-data/schemas?page=schemas
func (d *dataSourceIntegration) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_id":                    schema.StringAttribute{Required: true},
			"connection_id":             schema.StringAttribute{Computed: true},
			"connection_type":           schema.StringAttribute{Computed: true},
			"content_handling_strategy": schema.StringAttribute{Computed: true},
			"credentials_arn":           schema.StringAttribute{Computed: true},
			"description":               schema.StringAttribute{Computed: true},
			"integration_id":            schema.StringAttribute{Optional: true, Computed: true},
			"integration_method":        schema.StringAttribute{Computed: true},
			"integration_response_selection_expression": schema.StringAttribute{Computed: true},
			"integration_subtype":                       schema.StringAttribute{Computed: true},
			"integration_type":                          schema.StringAttribute{Computed: true},
			"integration_uri":                           schema.StringAttribute{Computed: true},
			"passthrough_behavior":                      schema.StringAttribute{Computed: true},
			"payload_format_version":                    schema.StringAttribute{Computed: true},
			// "request_parameters":                        schema.StringAttribute{Computed: true},
			// "request_templates":                         schema.StringAttribute{Computed: true},
			// "response_parameters":                       schema.StringAttribute{Computed: true},
			"template_selection_expression": schema.StringAttribute{Computed: true},
			// "timeout_milliseconds":                      schema.StringAttribute{Computed: true},
			// "tls_config": schema.ObjectAttribute{
			// 	AttributeTypes: map[string]attr.Type{
			// 		"server_name_to_verify": types.StringType,
			// 	},
			// 	Computed: true,
			// },
		},
	}
}

func (d *dataSourceIntegration) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// TIP: ==== DATA SOURCE READ ====
	// Generally, the Read function should do the following things. Make
	// sure there is a good reason if you don't do one of these.
	//
	// 1. Get a client connection to the relevant service
	// 2. Fetch the config
	// 3. Get information about a resource from AWS
	// 4. Set the ID, arguments, and attributes
	// 5. Set the tags
	// 6. Set the state
	// TIP: -- 1. Get a client connection to the relevant service
	conn := d.Meta().APIGatewayV2Conn(ctx)

	// TIP: -- 2. Fetch the config
	var data dataSourceIntegrationData
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TIP: -- 3. Get information about a resource from AWS
	out, err := conn.GetIntegrationWithContext(ctx, &apigatewayv2.GetIntegrationInput{
		ApiId:         aws.String(data.ApiID.ValueString()),
		IntegrationId: aws.String(data.IntegrationID.ValueString()),
	})
	if err != nil {
		resp.Diagnostics.AddError(
			create.ProblemStandardMessage(names.APIGatewayV2, create.ErrActionReading, DSNameIntegration, data.IntegrationID.String(), err),
			err.Error(),
		)
		return
	}

	// TIP: -- 4. Set the ID, arguments, and attributes
	//
	// For simple data types (i.e., schema.StringAttribute, schema.BoolAttribute,
	// schema.Int64Attribute, and schema.Float64Attribue), simply setting the
	// appropriate data struct field is sufficient. The flex package implements
	// helpers for converting between Go and Plugin-Framework types seamlessly. No
	// error or nil checking is necessary.
	//
	// However, there are some situations where more handling is needed such as
	// complex data types (e.g., schema.ListAttribute, schema.SetAttribute). In
	// these cases the flatten function may have a diagnostics return value, which
	// should be appended to resp.Diagnostics.
	data.ConnectionId = flex.StringToFramework(ctx, out.ConnectionId)
	data.ConnectionType = flex.StringToFramework(ctx, out.ConnectionType)
	data.ContentHandlingStrategy = flex.StringToFramework(ctx, out.ContentHandlingStrategy)
	data.CredentialsArn = flex.StringToFramework(ctx, out.CredentialsArn)
	data.Description = flex.StringToFramework(ctx, out.Description)
	data.IntegrationMethod = flex.StringToFramework(ctx, out.IntegrationMethod)
	data.IntegrationResponseSelectionExpression = flex.StringToFramework(ctx, out.IntegrationResponseSelectionExpression)
	data.IntegrationSubtype = flex.StringToFramework(ctx, out.IntegrationSubtype)
	data.IntegrationType = flex.StringToFramework(ctx, out.IntegrationType)
	data.IntegrationUri = flex.StringToFramework(ctx, out.IntegrationUri)
	data.PassthroughBehavior = flex.StringToFramework(ctx, out.PassthroughBehavior)
	data.PayloadFormatVersion = flex.StringToFramework(ctx, out.PayloadFormatVersion)
	// data.RequestParameters = flex.FlattenFrameworkStringMap(ctx, out.RequestParameters)
	// data.RequestTemplates = flex.FlattenFrameworkStringMap(ctx, out.RequestTemplates)
	// data.ResponseParameters = flex.FlattenFrameworkStringMap(ctx, out.ResponseParameters)
	data.TemplateSelectionExpression = flex.StringToFramework(ctx, out.TemplateSelectionExpression)
	// data.TimeoutMilliseconds = flex.StringToFramework(ctx, out.TimeoutMilliseconds)
	// data.TlsConfig = flex.StringToFramework(ctx, out.TlsConfig)

	// TIP: -- 5. Set the tags

	// TIP: -- 6. Set the state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// TIP: ==== DATA STRUCTURES ====
// With Terraform Plugin-Framework configurations are deserialized into
// Go types, providing type safety without the need for type assertions.
// These structs should match the schema definition exactly, and the `tfsdk`
// tag value should match the attribute name.
//
// See more:
// https://developer.hashicorp.com/terraform/plugin/framework/handling-data/accessing-values
type dataSourceIntegrationData struct {
	ApiID                                  types.String `tfsdk:"api_id"`
	ConnectionId                           types.String `tfsdk:"connection_id"`
	ConnectionType                         types.String `tfsdk:"connection_type"`
	ContentHandlingStrategy                types.String `tfsdk:"content_handling_strategy"`
	CredentialsArn                         types.String `tfsdk:"credentials_arn"`
	Description                            types.String `tfsdk:"description"`
	IntegrationID                          types.String `tfsdk:"integration_id"`
	IntegrationMethod                      types.String `tfsdk:"integration_method"`
	IntegrationResponseSelectionExpression types.String `tfsdk:"integration_response_selection_expression"`
	IntegrationSubtype                     types.String `tfsdk:"integration_subtype"`
	IntegrationType                        types.String `tfsdk:"integration_type"`
	IntegrationUri                         types.String `tfsdk:"integration_uri"`
	PassthroughBehavior                    types.String `tfsdk:"passthrough_behavior"`
	PayloadFormatVersion                   types.String `tfsdk:"payload_format_version"`
	// RequestParameters                      types.String `tfsdk:"request_parameters"`
	// RequestTemplates                       types.String `tfsdk:"request_templates"`
	// ResponseParameters                     types.String `tfsdk:"response_parameters"`
	TemplateSelectionExpression types.String `tfsdk:"template_selection_expression"`
	// TimeoutMilliseconds                    types.String `tfsdk:"timeout_milliseconds"`
	// TlsConfig                              types.Map    `tfsdk:"tls_config"`
}
