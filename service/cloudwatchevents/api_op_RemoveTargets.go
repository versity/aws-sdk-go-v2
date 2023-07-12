// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the specified targets from the specified rule. When the rule is
// triggered, those targets are no longer be invoked. When you remove a target,
// when the associated rule triggers, removed targets might continue to be invoked.
// Allow a short period of time for changes to take effect. This action can
// partially fail if too many requests are made at the same time. If that happens,
// FailedEntryCount is non-zero in the response and each entry in FailedEntries
// provides the ID of the failed target and the error code.
func (c *Client) RemoveTargets(ctx context.Context, params *RemoveTargetsInput, optFns ...func(*Options)) (*RemoveTargetsOutput, error) {
	if params == nil {
		params = &RemoveTargetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveTargets", params, optFns, c.addOperationRemoveTargetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveTargetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveTargetsInput struct {

	// The IDs of the targets to remove from the rule.
	//
	// This member is required.
	Ids []string

	// The name of the rule.
	//
	// This member is required.
	Rule *string

	// The name or ARN of the event bus associated with the rule. If you omit this,
	// the default event bus is used.
	EventBusName *string

	// If this is a managed rule, created by an Amazon Web Services service on your
	// behalf, you must specify Force as True to remove targets. This parameter is
	// ignored for rules that are not managed rules. You can check whether a rule is a
	// managed rule by using DescribeRule or ListRules and checking the ManagedBy
	// field of the response.
	Force bool

	noSmithyDocumentSerde
}

type RemoveTargetsOutput struct {

	// The failed target entries.
	FailedEntries []types.RemoveTargetsResultEntry

	// The number of failed entries.
	FailedEntryCount int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRemoveTargetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRemoveTargets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRemoveTargets{}, middleware.After)
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
	if err = addOpRemoveTargetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveTargets(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRemoveTargets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "RemoveTargets",
	}
}
