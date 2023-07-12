// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Auto-termination is supported in Amazon EMR releases 5.30.0 and 6.1.0 and
// later. For more information, see Using an auto-termination policy (https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-auto-termination-policy.html)
// . Creates or updates an auto-termination policy for an Amazon EMR cluster. An
// auto-termination policy defines the amount of idle time in seconds after which a
// cluster automatically terminates. For alternative cluster termination options,
// see Control cluster termination (https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-plan-termination.html)
// .
func (c *Client) PutAutoTerminationPolicy(ctx context.Context, params *PutAutoTerminationPolicyInput, optFns ...func(*Options)) (*PutAutoTerminationPolicyOutput, error) {
	if params == nil {
		params = &PutAutoTerminationPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAutoTerminationPolicy", params, optFns, c.addOperationPutAutoTerminationPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAutoTerminationPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAutoTerminationPolicyInput struct {

	// Specifies the ID of the Amazon EMR cluster to which the auto-termination policy
	// will be attached.
	//
	// This member is required.
	ClusterId *string

	// Specifies the auto-termination policy to attach to the cluster.
	AutoTerminationPolicy *types.AutoTerminationPolicy

	noSmithyDocumentSerde
}

type PutAutoTerminationPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutAutoTerminationPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutAutoTerminationPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutAutoTerminationPolicy{}, middleware.After)
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
	if err = addOpPutAutoTerminationPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutAutoTerminationPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutAutoTerminationPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "PutAutoTerminationPolicy",
	}
}
