// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates a subscription filter and associates it with the specified
// log group. With subscription filters, you can subscribe to a real-time stream of
// log events ingested through PutLogEvents (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutLogEvents.html)
// and have them delivered to a specific destination. When log events are sent to
// the receiving service, they are Base64 encoded and compressed with the GZIP
// format. The following destinations are supported for subscription filters:
//   - An Amazon Kinesis data stream belonging to the same account as the
//     subscription filter, for same-account delivery.
//   - A logical destination created with PutDestination (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDestination.html)
//     that belongs to a different account, for cross-account delivery. We currently
//     support Kinesis Data Streams and Kinesis Data Firehose as logical destinations.
//   - An Amazon Kinesis Data Firehose delivery stream that belongs to the same
//     account as the subscription filter, for same-account delivery.
//   - An Lambda function that belongs to the same account as the subscription
//     filter, for same-account delivery.
//
// Each log group can have up to two subscription filters associated with it. If
// you are updating an existing filter, you must specify the correct name in
// filterName . To perform a PutSubscriptionFilter operation for any destination
// except a Lambda function, you must also have the iam:PassRole permission.
func (c *Client) PutSubscriptionFilter(ctx context.Context, params *PutSubscriptionFilterInput, optFns ...func(*Options)) (*PutSubscriptionFilterOutput, error) {
	if params == nil {
		params = &PutSubscriptionFilterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutSubscriptionFilter", params, optFns, c.addOperationPutSubscriptionFilterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutSubscriptionFilterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutSubscriptionFilterInput struct {

	// The ARN of the destination to deliver matching log events to. Currently, the
	// supported destinations are:
	//   - An Amazon Kinesis stream belonging to the same account as the subscription
	//   filter, for same-account delivery.
	//   - A logical destination (specified using an ARN) belonging to a different
	//   account, for cross-account delivery. If you're setting up a cross-account
	//   subscription, the destination must have an IAM policy associated with it. The
	//   IAM policy must allow the sender to send logs to the destination. For more
	//   information, see PutDestinationPolicy (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDestinationPolicy.html)
	//   .
	//   - A Kinesis Data Firehose delivery stream belonging to the same account as
	//   the subscription filter, for same-account delivery.
	//   - A Lambda function belonging to the same account as the subscription filter,
	//   for same-account delivery.
	//
	// This member is required.
	DestinationArn *string

	// A name for the subscription filter. If you are updating an existing filter, you
	// must specify the correct name in filterName . To find the name of the filter
	// currently associated with a log group, use DescribeSubscriptionFilters (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeSubscriptionFilters.html)
	// .
	//
	// This member is required.
	FilterName *string

	// A filter pattern for subscribing to a filtered stream of log events.
	//
	// This member is required.
	FilterPattern *string

	// The name of the log group.
	//
	// This member is required.
	LogGroupName *string

	// The method used to distribute log data to the destination. By default, log data
	// is grouped by log stream, but the grouping can be set to random for a more even
	// distribution. This property is only applicable when the destination is an Amazon
	// Kinesis data stream.
	Distribution types.Distribution

	// The ARN of an IAM role that grants CloudWatch Logs permissions to deliver
	// ingested log events to the destination stream. You don't need to provide the ARN
	// when you are working with a logical destination for cross-account delivery.
	RoleArn *string

	noSmithyDocumentSerde
}

type PutSubscriptionFilterOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutSubscriptionFilterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutSubscriptionFilter{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutSubscriptionFilter{}, middleware.After)
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
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpPutSubscriptionFilterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutSubscriptionFilter(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutSubscriptionFilter(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "PutSubscriptionFilter",
	}
}
