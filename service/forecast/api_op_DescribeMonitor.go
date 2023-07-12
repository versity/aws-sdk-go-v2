// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes a monitor resource. In addition to listing the properties provided in
// the CreateMonitor request, this operation lists the following properties:
//   - Baseline
//   - CreationTime
//   - LastEvaluationTime
//   - LastEvaluationState
//   - LastModificationTime
//   - Message
//   - Status
func (c *Client) DescribeMonitor(ctx context.Context, params *DescribeMonitorInput, optFns ...func(*Options)) (*DescribeMonitorOutput, error) {
	if params == nil {
		params = &DescribeMonitorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMonitor", params, optFns, c.addOperationDescribeMonitorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMonitorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMonitorInput struct {

	// The Amazon Resource Name (ARN) of the monitor resource to describe.
	//
	// This member is required.
	MonitorArn *string

	noSmithyDocumentSerde
}

type DescribeMonitorOutput struct {

	// Metrics you can use as a baseline for comparison purposes. Use these values you
	// interpret monitoring results for an auto predictor.
	Baseline *types.Baseline

	// The timestamp for when the monitor resource was created.
	CreationTime *time.Time

	// The estimated number of minutes remaining before the monitor resource finishes
	// its current evaluation.
	EstimatedEvaluationTimeRemainingInMinutes *int64

	// The state of the monitor's latest evaluation.
	LastEvaluationState *string

	// The timestamp of the latest evaluation completed by the monitor.
	LastEvaluationTime *time.Time

	// The timestamp of the latest modification to the monitor.
	LastModificationTime *time.Time

	// An error message, if any, for the monitor.
	Message *string

	// The Amazon Resource Name (ARN) of the monitor resource described.
	MonitorArn *string

	// The name of the monitor.
	MonitorName *string

	// The Amazon Resource Name (ARN) of the auto predictor being monitored.
	ResourceArn *string

	// The status of the monitor resource.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeMonitorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMonitor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMonitor{}, middleware.After)
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
	if err = addOpDescribeMonitorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMonitor(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeMonitor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "DescribeMonitor",
	}
}
