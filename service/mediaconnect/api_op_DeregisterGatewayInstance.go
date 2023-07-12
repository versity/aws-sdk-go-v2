// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deregisters an instance. Before you deregister an instance, all bridges running
// on the instance must be stopped. If you want to deregister an instance without
// stopping the bridges, you must use the --force option.
func (c *Client) DeregisterGatewayInstance(ctx context.Context, params *DeregisterGatewayInstanceInput, optFns ...func(*Options)) (*DeregisterGatewayInstanceOutput, error) {
	if params == nil {
		params = &DeregisterGatewayInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterGatewayInstance", params, optFns, c.addOperationDeregisterGatewayInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterGatewayInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterGatewayInstanceInput struct {

	// The Amazon Resource Name (ARN) of the gateway that contains the instance that
	// you want to deregister.
	//
	// This member is required.
	GatewayInstanceArn *string

	// Force the deregistration of an instance. Force will deregister an instance,
	// even if there are bridges running on it.
	Force bool

	noSmithyDocumentSerde
}

type DeregisterGatewayInstanceOutput struct {

	// The Amazon Resource Name (ARN) of the instance.
	GatewayInstanceArn *string

	// The status of the instance.
	InstanceState types.InstanceState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeregisterGatewayInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeregisterGatewayInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeregisterGatewayInstance{}, middleware.After)
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
	if err = addOpDeregisterGatewayInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterGatewayInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeregisterGatewayInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "DeregisterGatewayInstance",
	}
}
