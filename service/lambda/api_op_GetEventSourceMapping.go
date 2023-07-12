// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns details about an event source mapping. You can get the identifier of a
// mapping from the output of ListEventSourceMappings .
func (c *Client) GetEventSourceMapping(ctx context.Context, params *GetEventSourceMappingInput, optFns ...func(*Options)) (*GetEventSourceMappingOutput, error) {
	if params == nil {
		params = &GetEventSourceMappingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEventSourceMapping", params, optFns, c.addOperationGetEventSourceMappingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEventSourceMappingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEventSourceMappingInput struct {

	// The identifier of the event source mapping.
	//
	// This member is required.
	UUID *string

	noSmithyDocumentSerde
}

// A mapping between an Amazon Web Services resource and a Lambda function. For
// details, see CreateEventSourceMapping .
type GetEventSourceMappingOutput struct {

	// Specific configuration settings for an Amazon Managed Streaming for Apache
	// Kafka (Amazon MSK) event source.
	AmazonManagedKafkaEventSourceConfig *types.AmazonManagedKafkaEventSourceConfig

	// The maximum number of records in each batch that Lambda pulls from your stream
	// or queue and sends to your function. Lambda passes all of the records in the
	// batch to the function in a single call, up to the payload limit for synchronous
	// invocation (6 MB). Default value: Varies by service. For Amazon SQS, the default
	// is 10. For all other services, the default is 100. Related setting: When you set
	// BatchSize to a value greater than 10, you must set
	// MaximumBatchingWindowInSeconds to at least 1.
	BatchSize *int32

	// (Kinesis and DynamoDB Streams only) If the function returns an error, split the
	// batch in two and retry. The default value is false.
	BisectBatchOnFunctionError *bool

	// (Kinesis and DynamoDB Streams only) An Amazon SQS queue or Amazon SNS topic
	// destination for discarded records.
	DestinationConfig *types.DestinationConfig

	// Specific configuration settings for a DocumentDB event source.
	DocumentDBEventSourceConfig *types.DocumentDBEventSourceConfig

	// The Amazon Resource Name (ARN) of the event source.
	EventSourceArn *string

	// An object that defines the filter criteria that determine whether Lambda should
	// process an event. For more information, see Lambda event filtering (https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html)
	// .
	FilterCriteria *types.FilterCriteria

	// The ARN of the Lambda function.
	FunctionArn *string

	// (Kinesis, DynamoDB Streams, and Amazon SQS) A list of current response type
	// enums applied to the event source mapping.
	FunctionResponseTypes []types.FunctionResponseType

	// The date that the event source mapping was last updated or that its state
	// changed.
	LastModified *time.Time

	// The result of the last Lambda invocation of your function.
	LastProcessingResult *string

	// The maximum amount of time, in seconds, that Lambda spends gathering records
	// before invoking the function. You can configure MaximumBatchingWindowInSeconds
	// to any value from 0 seconds to 300 seconds in increments of seconds. For streams
	// and Amazon SQS event sources, the default batching window is 0 seconds. For
	// Amazon MSK, Self-managed Apache Kafka, Amazon MQ, and DocumentDB event sources,
	// the default batching window is 500 ms. Note that because you can only change
	// MaximumBatchingWindowInSeconds in increments of seconds, you cannot revert back
	// to the 500 ms default batching window after you have changed it. To restore the
	// default batching window, you must create a new event source mapping. Related
	// setting: For streams and Amazon SQS event sources, when you set BatchSize to a
	// value greater than 10, you must set MaximumBatchingWindowInSeconds to at least
	// 1.
	MaximumBatchingWindowInSeconds *int32

	// (Kinesis and DynamoDB Streams only) Discard records older than the specified
	// age. The default value is -1, which sets the maximum age to infinite. When the
	// value is set to infinite, Lambda never discards old records. The minimum valid
	// value for maximum record age is 60s. Although values less than 60 and greater
	// than -1 fall within the parameter's absolute range, they are not allowed
	MaximumRecordAgeInSeconds *int32

	// (Kinesis and DynamoDB Streams only) Discard records after the specified number
	// of retries. The default value is -1, which sets the maximum number of retries to
	// infinite. When MaximumRetryAttempts is infinite, Lambda retries failed records
	// until the record expires in the event source.
	MaximumRetryAttempts *int32

	// (Kinesis and DynamoDB Streams only) The number of batches to process
	// concurrently from each shard. The default value is 1.
	ParallelizationFactor *int32

	// (Amazon MQ) The name of the Amazon MQ broker destination queue to consume.
	Queues []string

	// (Amazon SQS only) The scaling configuration for the event source. For more
	// information, see Configuring maximum concurrency for Amazon SQS event sources (https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#events-sqs-max-concurrency)
	// .
	ScalingConfig *types.ScalingConfig

	// The self-managed Apache Kafka cluster for your event source.
	SelfManagedEventSource *types.SelfManagedEventSource

	// Specific configuration settings for a self-managed Apache Kafka event source.
	SelfManagedKafkaEventSourceConfig *types.SelfManagedKafkaEventSourceConfig

	// An array of the authentication protocol, VPC components, or virtual host to
	// secure and define your event source.
	SourceAccessConfigurations []types.SourceAccessConfiguration

	// The position in a stream from which to start reading. Required for Amazon
	// Kinesis and Amazon DynamoDB Stream event sources. AT_TIMESTAMP is supported
	// only for Amazon Kinesis streams, Amazon DocumentDB, Amazon MSK, and self-managed
	// Apache Kafka.
	StartingPosition types.EventSourcePosition

	// With StartingPosition set to AT_TIMESTAMP , the time from which to start
	// reading. StartingPositionTimestamp cannot be in the future.
	StartingPositionTimestamp *time.Time

	// The state of the event source mapping. It can be one of the following: Creating
	// , Enabling , Enabled , Disabling , Disabled , Updating , or Deleting .
	State *string

	// Indicates whether a user or Lambda made the last change to the event source
	// mapping.
	StateTransitionReason *string

	// The name of the Kafka topic.
	Topics []string

	// (Kinesis and DynamoDB Streams only) The duration in seconds of a processing
	// window for DynamoDB and Kinesis Streams event sources. A value of 0 seconds
	// indicates no tumbling window.
	TumblingWindowInSeconds *int32

	// The identifier of the event source mapping.
	UUID *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEventSourceMappingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetEventSourceMapping{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetEventSourceMapping{}, middleware.After)
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
	if err = addOpGetEventSourceMappingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEventSourceMapping(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetEventSourceMapping(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "GetEventSourceMapping",
	}
}
