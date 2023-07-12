// Code generated by smithy-go-codegen DO NOT EDIT.

package kafkaconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kafkaconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// A summary description of the custom plugin.
func (c *Client) DescribeCustomPlugin(ctx context.Context, params *DescribeCustomPluginInput, optFns ...func(*Options)) (*DescribeCustomPluginOutput, error) {
	if params == nil {
		params = &DescribeCustomPluginInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCustomPlugin", params, optFns, c.addOperationDescribeCustomPluginMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCustomPluginOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCustomPluginInput struct {

	// Returns information about a custom plugin.
	//
	// This member is required.
	CustomPluginArn *string

	noSmithyDocumentSerde
}

type DescribeCustomPluginOutput struct {

	// The time that the custom plugin was created.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the custom plugin.
	CustomPluginArn *string

	// The state of the custom plugin.
	CustomPluginState types.CustomPluginState

	// The description of the custom plugin.
	Description *string

	// The latest successfully created revision of the custom plugin. If there are no
	// successfully created revisions, this field will be absent.
	LatestRevision *types.CustomPluginRevisionSummary

	// The name of the custom plugin.
	Name *string

	// Details about the state of a custom plugin.
	StateDescription *types.StateDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeCustomPluginMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeCustomPlugin{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeCustomPlugin{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeCustomPluginValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCustomPlugin(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeCustomPlugin(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kafkaconnect",
		OperationName: "DescribeCustomPlugin",
	}
}
