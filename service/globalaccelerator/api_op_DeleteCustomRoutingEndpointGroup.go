// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Delete an endpoint group from a listener for a custom routing accelerator.
func (c *Client) DeleteCustomRoutingEndpointGroup(ctx context.Context, params *DeleteCustomRoutingEndpointGroupInput, optFns ...func(*Options)) (*DeleteCustomRoutingEndpointGroupOutput, error) {
	if params == nil {
		params = &DeleteCustomRoutingEndpointGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteCustomRoutingEndpointGroup", params, optFns, c.addOperationDeleteCustomRoutingEndpointGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteCustomRoutingEndpointGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteCustomRoutingEndpointGroupInput struct {

	// The Amazon Resource Name (ARN) of the endpoint group to delete.
	//
	// This member is required.
	EndpointGroupArn *string

	noSmithyDocumentSerde
}

type DeleteCustomRoutingEndpointGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteCustomRoutingEndpointGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteCustomRoutingEndpointGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteCustomRoutingEndpointGroup{}, middleware.After)
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
	if err = addOpDeleteCustomRoutingEndpointGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCustomRoutingEndpointGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteCustomRoutingEndpointGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "DeleteCustomRoutingEndpointGroup",
	}
}
