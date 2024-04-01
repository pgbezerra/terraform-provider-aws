// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/apigatewayv2"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"

	// TIP: You will often need to import the package that this test file lives
	// in. Since it is in the "test" context, it must import the package to use
	// any normal context constants, variables, or functions.

	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccAPIGatewayV2IntegrationDataSource_basic(t *testing.T) {
	ctx := acctest.Context(t)

	var apiId string
	var v apigatewayv2.GetIntegrationOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	dataSourceName := "data.aws_apigatewayv2_integration.test"
	resourceName := "aws_apigatewayv2_integration.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.APIGatewayV2ServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckIntegrationDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccIntegrationDataSourceConfig_sqs(rName, 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIntegrationExists(ctx, dataSourceName, &apiId, &v),
					resource.TestCheckResourceAttrPair(dataSourceName, "connection_type", resourceName, "connection_type"),
					resource.TestCheckResourceAttrPair(dataSourceName, "content_handling_strategy", resourceName, "content_handling_strategy"),
					resource.TestCheckResourceAttrPair(dataSourceName, "credentials_arn", resourceName, "credentials_arn"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "integration_method", resourceName, "integration_method"),
					resource.TestCheckResourceAttrPair(dataSourceName, "integration_response_selection_expression", resourceName, "integration_response_selection_expression"),
					resource.TestCheckResourceAttrPair(dataSourceName, "integration_subtype", resourceName, "integration_subtype"),
					resource.TestCheckResourceAttrPair(dataSourceName, "integration_type", resourceName, "integration_type"),
					resource.TestCheckResourceAttrPair(dataSourceName, "integration_uri", resourceName, "integration_uri"),
					resource.TestCheckResourceAttrPair(dataSourceName, "passthrough_behavior", resourceName, "passthrough_behavior"),
					resource.TestCheckResourceAttrPair(dataSourceName, "payload_format_version", resourceName, "payload_format_version"),
					resource.TestCheckResourceAttrPair(dataSourceName, "request_parameters.%", resourceName, "request_parameters.%"),
					resource.TestCheckResourceAttrPair(dataSourceName, "request_parameters.MessageBody", resourceName, "request_parameters.MessageBody"),
					resource.TestCheckResourceAttrPair(dataSourceName, "request_parameters.MessageGroupId", resourceName, "request_parameters.MessageGroupId"),
					resource.TestCheckResourceAttrPair(dataSourceName, "request_parameters.QueueUrl", resourceName, "request_parameters.QueueUrl"),
					resource.TestCheckResourceAttrPair(dataSourceName, "request_templates.%", resourceName, "request_templates.%"),
					resource.TestCheckResourceAttrPair(dataSourceName, "response_parameters.#", resourceName, "response_parameters.#"),
					resource.TestCheckResourceAttrPair(dataSourceName, "template_selection_expression", resourceName, "template_selection_expression"),
					resource.TestCheckResourceAttrPair(dataSourceName, "timeout_milliseconds", resourceName, "timeout_milliseconds"),
					resource.TestCheckResourceAttrPair(dataSourceName, "tls_config.#", resourceName, "tls_config.#"),
				),
			},
		},
	})
}

func testAccIntegrationDataSourceConfig_sqs(rName string, queueIndex int) string {
	return acctest.ConfigCompose(
		testAccIntegrationConfig_sqs(rName, queueIndex),
		fmt.Sprintf(`
data "aws_apigatewayv2_integration" "test" {
  api_id           = aws_apigatewayv2_integration.test.api_id
  integration_type = aws_apigatewayv2_integration.test.integration_type
}
`, rName))
}
