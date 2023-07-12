// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about a query logging policy. A query logging policy specifies
// the Resolver query logging operations and resources that you want to allow
// another Amazon Web Services account to be able to use.
func (c *Client) GetResolverQueryLogConfigPolicy(ctx context.Context, params *GetResolverQueryLogConfigPolicyInput, optFns ...func(*Options)) (*GetResolverQueryLogConfigPolicyOutput, error) {
	if params == nil {
		params = &GetResolverQueryLogConfigPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetResolverQueryLogConfigPolicy", params, optFns, c.addOperationGetResolverQueryLogConfigPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetResolverQueryLogConfigPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResolverQueryLogConfigPolicyInput struct {

	// The ARN of the query logging configuration that you want to get the query
	// logging policy for.
	//
	// This member is required.
	Arn *string

	noSmithyDocumentSerde
}

type GetResolverQueryLogConfigPolicyOutput struct {

	// Information about the query logging policy for the query logging configuration
	// that you specified in a GetResolverQueryLogConfigPolicy request.
	ResolverQueryLogConfigPolicy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetResolverQueryLogConfigPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetResolverQueryLogConfigPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetResolverQueryLogConfigPolicy{}, middleware.After)
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
	if err = addOpGetResolverQueryLogConfigPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetResolverQueryLogConfigPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetResolverQueryLogConfigPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "GetResolverQueryLogConfigPolicy",
	}
}
