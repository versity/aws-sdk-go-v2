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

// Updates an existing bridge source.
func (c *Client) UpdateBridgeSource(ctx context.Context, params *UpdateBridgeSourceInput, optFns ...func(*Options)) (*UpdateBridgeSourceOutput, error) {
	if params == nil {
		params = &UpdateBridgeSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateBridgeSource", params, optFns, c.addOperationUpdateBridgeSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateBridgeSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The fields that you want to update in the bridge source.
type UpdateBridgeSourceInput struct {

	// The ARN of the bridge that you want to update.
	//
	// This member is required.
	BridgeArn *string

	// The name of the source that you want to update.
	//
	// This member is required.
	SourceName *string

	// Update the flow source of the bridge.
	FlowSource *types.UpdateBridgeFlowSourceRequest

	// Update the network source of the bridge.
	NetworkSource *types.UpdateBridgeNetworkSourceRequest

	noSmithyDocumentSerde
}

type UpdateBridgeSourceOutput struct {

	// The Amazon Resource Number (ARN) of the bridge.
	BridgeArn *string

	// The bridge's source.
	Source *types.BridgeSource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateBridgeSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateBridgeSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateBridgeSource{}, middleware.After)
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
	if err = addOpUpdateBridgeSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateBridgeSource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateBridgeSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "UpdateBridgeSource",
	}
}
