// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"time"
)

// This action is useful to determine if a bucket exists and you have permission to
// access it. The action returns a 200 OK if the bucket exists and you have
// permission to access it. If the bucket does not exist or you do not have
// permission to access it, the HEAD request returns a generic 404 Not Found or 403
// Forbidden code. A message body is not included, so you cannot determine the
// exception beyond these error codes. To use this operation, you must have
// permissions to perform the s3:ListBucket action. The bucket owner has this
// permission by default and can grant this permission to others. For more
// information about permissions, see Permissions Related to Bucket Subresource
// Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources
// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html).
// To use this API against an access point, you must provide the alias of the
// access point in place of the bucket name or specify the access point ARN. When
// using the access point ARN, you must direct requests to the access point
// hostname. The access point hostname takes the form
// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using the
// Amazon Web Services SDKs, you provide the ARN in place of the bucket name. For
// more information see, Using access points
// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-access-points.html).
func (c *Client) HeadBucket(ctx context.Context, params *HeadBucketInput, optFns ...func(*Options)) (*HeadBucketOutput, error) {
	if params == nil {
		params = &HeadBucketInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "HeadBucket", params, optFns, c.addOperationHeadBucketMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*HeadBucketOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HeadBucketInput struct {

	// The bucket name. When using this action with an access point, you must direct
	// requests to the access point hostname. The access point hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// action with an access point through the Amazon Web Services SDKs, you provide
	// the access point ARN in place of the bucket name. For more information about
	// access point ARNs, see Using access points
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-access-points.html)
	// in the Amazon S3 User Guide. When using this action with Amazon S3 on Outposts,
	// you must direct requests to the S3 on Outposts hostname. The S3 on Outposts
	// hostname takes the form
	// AccessPointName-AccountId.outpostID.s3-outposts.Region.amazonaws.com. When using
	// this action using S3 on Outposts through the Amazon Web Services SDKs, you
	// provide the Outposts bucket ARN in place of the bucket name. For more
	// information about S3 on Outposts ARNs, see Using S3 on Outposts
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html) in the
	// Amazon S3 User Guide.
	//
	// This member is required.
	Bucket *string

	// The account ID of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string

	noSmithyDocumentSerde
}

type HeadBucketOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationHeadBucketMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpHeadBucket{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpHeadBucket{}, middleware.After)
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
	if err = addOpHeadBucketValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opHeadBucket(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addHeadBucketUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// HeadBucketAPIClient is a client that implements the HeadBucket operation.
type HeadBucketAPIClient interface {
	HeadBucket(context.Context, *HeadBucketInput, ...func(*Options)) (*HeadBucketOutput, error)
}

var _ HeadBucketAPIClient = (*Client)(nil)

// BucketExistsWaiterOptions are waiter options for BucketExistsWaiter
type BucketExistsWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// BucketExistsWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, BucketExistsWaiter will use default max delay of 120 seconds. Note that
	// MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *HeadBucketInput, *HeadBucketOutput, error) (bool, error)
}

// BucketExistsWaiter defines the waiters for BucketExists
type BucketExistsWaiter struct {
	client HeadBucketAPIClient

	options BucketExistsWaiterOptions
}

// NewBucketExistsWaiter constructs a BucketExistsWaiter.
func NewBucketExistsWaiter(client HeadBucketAPIClient, optFns ...func(*BucketExistsWaiterOptions)) *BucketExistsWaiter {
	options := BucketExistsWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = bucketExistsStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &BucketExistsWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for BucketExists waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *BucketExistsWaiter) Wait(ctx context.Context, params *HeadBucketInput, maxWaitDur time.Duration, optFns ...func(*BucketExistsWaiterOptions)) error {
	if maxWaitDur <= 0 {
		return fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.HeadBucket(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return err
		}
		if !retryable {
			return nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return fmt.Errorf("exceeded max wait time for BucketExists waiter")
}

func bucketExistsStateRetryable(ctx context.Context, input *HeadBucketInput, output *HeadBucketOutput, err error) (bool, error) {

	if err == nil {
		return false, nil
	}

	if err != nil {
		var errorType *types.NotFound
		if errors.As(err, &errorType) {
			return true, nil
		}
	}

	return true, nil
}

// BucketNotExistsWaiterOptions are waiter options for BucketNotExistsWaiter
type BucketNotExistsWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// BucketNotExistsWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, BucketNotExistsWaiter will use default max delay of 120 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *HeadBucketInput, *HeadBucketOutput, error) (bool, error)
}

// BucketNotExistsWaiter defines the waiters for BucketNotExists
type BucketNotExistsWaiter struct {
	client HeadBucketAPIClient

	options BucketNotExistsWaiterOptions
}

// NewBucketNotExistsWaiter constructs a BucketNotExistsWaiter.
func NewBucketNotExistsWaiter(client HeadBucketAPIClient, optFns ...func(*BucketNotExistsWaiterOptions)) *BucketNotExistsWaiter {
	options := BucketNotExistsWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = bucketNotExistsStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &BucketNotExistsWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for BucketNotExists waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *BucketNotExistsWaiter) Wait(ctx context.Context, params *HeadBucketInput, maxWaitDur time.Duration, optFns ...func(*BucketNotExistsWaiterOptions)) error {
	if maxWaitDur <= 0 {
		return fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.HeadBucket(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return err
		}
		if !retryable {
			return nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return fmt.Errorf("exceeded max wait time for BucketNotExists waiter")
}

func bucketNotExistsStateRetryable(ctx context.Context, input *HeadBucketInput, output *HeadBucketOutput, err error) (bool, error) {

	if err != nil {
		var errorType *types.NotFound
		if errors.As(err, &errorType) {
			return false, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opHeadBucket(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "HeadBucket",
	}
}

// getHeadBucketBucketMember returns a pointer to string denoting a provided bucket
// member valueand a boolean indicating if the input has a modeled bucket name,
func getHeadBucketBucketMember(input interface{}) (*string, bool) {
	in := input.(*HeadBucketInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addHeadBucketUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getHeadBucketBucketMember,
		},
		UsePathStyle:            options.UsePathStyle,
		UseAccelerate:           options.UseAccelerate,
		SupportsAccelerate:      true,
		TargetS3ObjectLambda:    false,
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseDualstack:            options.UseDualstack,
		UseARNRegion:            options.UseARNRegion,
	})
}
