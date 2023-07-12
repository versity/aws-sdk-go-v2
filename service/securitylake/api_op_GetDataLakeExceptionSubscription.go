// Code generated by smithy-go-codegen DO NOT EDIT.

package securitylake

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the details of exception notifications for the account in Amazon
// Security Lake.
func (c *Client) GetDataLakeExceptionSubscription(ctx context.Context, params *GetDataLakeExceptionSubscriptionInput, optFns ...func(*Options)) (*GetDataLakeExceptionSubscriptionOutput, error) {
	if params == nil {
		params = &GetDataLakeExceptionSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDataLakeExceptionSubscription", params, optFns, c.addOperationGetDataLakeExceptionSubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDataLakeExceptionSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDataLakeExceptionSubscriptionInput struct {
	noSmithyDocumentSerde
}

type GetDataLakeExceptionSubscriptionOutput struct {

	// The expiration period and time-to-live (TTL).
	ExceptionTimeToLive *int64

	// The Amazon Web Services account where you receive exception notifications.
	NotificationEndpoint *string

	// The subscription protocol to which exception notifications are posted.
	SubscriptionProtocol *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDataLakeExceptionSubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDataLakeExceptionSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDataLakeExceptionSubscription{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDataLakeExceptionSubscription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDataLakeExceptionSubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securitylake",
		OperationName: "GetDataLakeExceptionSubscription",
	}
}
