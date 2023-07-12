// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// An asynchronous API that deletes a stream’s existing edge configuration, as
// well as the corresponding media from the Edge Agent. When you invoke this API,
// the sync status is set to DELETING . A deletion process starts, in which active
// edge jobs are stopped and all media is deleted from the edge device. The time to
// delete varies, depending on the total amount of stored media. If the deletion
// process fails, the sync status changes to DELETE_FAILED . You will need to
// re-try the deletion. When the deletion process has completed successfully, the
// edge configuration is no longer accessible.
func (c *Client) DeleteEdgeConfiguration(ctx context.Context, params *DeleteEdgeConfigurationInput, optFns ...func(*Options)) (*DeleteEdgeConfigurationOutput, error) {
	if params == nil {
		params = &DeleteEdgeConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteEdgeConfiguration", params, optFns, c.addOperationDeleteEdgeConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteEdgeConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteEdgeConfigurationInput struct {

	// The Amazon Resource Name (ARN) of the stream. Specify either the StreamName or
	// the StreamARN .
	StreamARN *string

	// The name of the stream from which to delete the edge configuration. Specify
	// either the StreamName or the StreamARN .
	StreamName *string

	noSmithyDocumentSerde
}

type DeleteEdgeConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteEdgeConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteEdgeConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteEdgeConfiguration{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteEdgeConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteEdgeConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "DeleteEdgeConfiguration",
	}
}
