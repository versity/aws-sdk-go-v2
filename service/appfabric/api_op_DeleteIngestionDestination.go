// Code generated by smithy-go-codegen DO NOT EDIT.

package appfabric

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an ingestion destination. This deletes the association between an
// ingestion and it's destination. It doesn't delete previously ingested data or
// the storage destination, such as the Amazon S3 bucket where the data is
// delivered. If the ingestion destination is deleted while the associated
// ingestion is enabled, the ingestion will fail and is eventually disabled.
func (c *Client) DeleteIngestionDestination(ctx context.Context, params *DeleteIngestionDestinationInput, optFns ...func(*Options)) (*DeleteIngestionDestinationOutput, error) {
	if params == nil {
		params = &DeleteIngestionDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteIngestionDestination", params, optFns, c.addOperationDeleteIngestionDestinationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteIngestionDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteIngestionDestinationInput struct {

	// The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app
	// bundle to use for the request.
	//
	// This member is required.
	AppBundleIdentifier *string

	// The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the
	// ingestion destination to use for the request.
	//
	// This member is required.
	IngestionDestinationIdentifier *string

	// The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the
	// ingestion to use for the request.
	//
	// This member is required.
	IngestionIdentifier *string

	noSmithyDocumentSerde
}

type DeleteIngestionDestinationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteIngestionDestinationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteIngestionDestination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteIngestionDestination{}, middleware.After)
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
	if err = addOpDeleteIngestionDestinationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteIngestionDestination(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteIngestionDestination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appfabric",
		OperationName: "DeleteIngestionDestination",
	}
}
