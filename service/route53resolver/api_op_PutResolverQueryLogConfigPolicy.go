// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Specifies an Amazon Web Services account that you want to share a query logging
// configuration with, the query logging configuration that you want to share, and
// the operations that you want the account to be able to perform on the
// configuration.
func (c *Client) PutResolverQueryLogConfigPolicy(ctx context.Context, params *PutResolverQueryLogConfigPolicyInput, optFns ...func(*Options)) (*PutResolverQueryLogConfigPolicyOutput, error) {
	if params == nil {
		params = &PutResolverQueryLogConfigPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutResolverQueryLogConfigPolicy", params, optFns, c.addOperationPutResolverQueryLogConfigPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutResolverQueryLogConfigPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutResolverQueryLogConfigPolicyInput struct {

	// The Amazon Resource Name (ARN) of the account that you want to share rules with.
	//
	// This member is required.
	Arn *string

	// An Identity and Access Management policy statement that lists the query logging
	// configurations that you want to share with another Amazon Web Services account
	// and the operations that you want the account to be able to perform. You can
	// specify the following operations in the Actions section of the statement:
	//
	// *
	// route53resolver:AssociateResolverQueryLogConfig
	//
	// *
	// route53resolver:DisassociateResolverQueryLogConfig
	//
	// *
	// route53resolver:ListResolverQueryLogConfigAssociations
	//
	// *
	// route53resolver:ListResolverQueryLogConfigs
	//
	// In the Resource section of the
	// statement, you specify the ARNs for the query logging configurations that you
	// want to share with the account that you specified in Arn.
	//
	// This member is required.
	ResolverQueryLogConfigPolicy *string

	noSmithyDocumentSerde
}

// The response to a PutResolverQueryLogConfigPolicy request.
type PutResolverQueryLogConfigPolicyOutput struct {

	// Whether the PutResolverQueryLogConfigPolicy request was successful.
	ReturnValue bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutResolverQueryLogConfigPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutResolverQueryLogConfigPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutResolverQueryLogConfigPolicy{}, middleware.After)
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
	if err = addOpPutResolverQueryLogConfigPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutResolverQueryLogConfigPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutResolverQueryLogConfigPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "PutResolverQueryLogConfigPolicy",
	}
}
