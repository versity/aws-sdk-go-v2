// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Purchase an offering and create a reservation.
func (c *Client) PurchaseOffering(ctx context.Context, params *PurchaseOfferingInput, optFns ...func(*Options)) (*PurchaseOfferingOutput, error) {
	if params == nil {
		params = &PurchaseOfferingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PurchaseOffering", params, optFns, c.addOperationPurchaseOfferingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PurchaseOfferingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for PurchaseOfferingRequest
type PurchaseOfferingInput struct {

	// Number of resources
	//
	// This member is required.
	Count int32

	// Offering to purchase, e.g. '87654321'
	//
	// This member is required.
	OfferingId *string

	// Name for the new reservation
	Name *string

	// Renewal settings for the reservation
	RenewalSettings *types.RenewalSettings

	// Unique request ID to be specified. This is needed to prevent retries from
	// creating multiple resources.
	RequestId *string

	// Requested reservation start time (UTC) in ISO-8601 format. The specified time
	// must be between the first day of the current month and one year from now. If no
	// value is given, the default is now.
	Start *string

	// A collection of key-value pairs
	Tags map[string]string

	noSmithyDocumentSerde
}

// Placeholder documentation for PurchaseOfferingResponse
type PurchaseOfferingOutput struct {

	// Reserved resources available to use
	Reservation *types.Reservation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPurchaseOfferingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPurchaseOffering{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPurchaseOffering{}, middleware.After)
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
	if err = addIdempotencyToken_opPurchaseOfferingMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpPurchaseOfferingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPurchaseOffering(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpPurchaseOffering struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpPurchaseOffering) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpPurchaseOffering) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*PurchaseOfferingInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *PurchaseOfferingInput ")
	}

	if input.RequestId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.RequestId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opPurchaseOfferingMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpPurchaseOffering{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opPurchaseOffering(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "PurchaseOffering",
	}
}
