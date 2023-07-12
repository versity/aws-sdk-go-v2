// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an origin request policy. You cannot delete an origin request policy if
// it's attached to any cache behaviors. First update your distributions to remove
// the origin request policy from all cache behaviors, then delete the origin
// request policy. To delete an origin request policy, you must provide the
// policy's identifier and version. To get the identifier, you can use
// ListOriginRequestPolicies or GetOriginRequestPolicy .
func (c *Client) DeleteOriginRequestPolicy(ctx context.Context, params *DeleteOriginRequestPolicyInput, optFns ...func(*Options)) (*DeleteOriginRequestPolicyOutput, error) {
	if params == nil {
		params = &DeleteOriginRequestPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteOriginRequestPolicy", params, optFns, c.addOperationDeleteOriginRequestPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteOriginRequestPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteOriginRequestPolicyInput struct {

	// The unique identifier for the origin request policy that you are deleting. To
	// get the identifier, you can use ListOriginRequestPolicies .
	//
	// This member is required.
	Id *string

	// The version of the origin request policy that you are deleting. The version is
	// the origin request policy's ETag value, which you can get using
	// ListOriginRequestPolicies , GetOriginRequestPolicy , or
	// GetOriginRequestPolicyConfig .
	IfMatch *string

	noSmithyDocumentSerde
}

type DeleteOriginRequestPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteOriginRequestPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteOriginRequestPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteOriginRequestPolicy{}, middleware.After)
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
	if err = addOpDeleteOriginRequestPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteOriginRequestPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteOriginRequestPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "DeleteOriginRequestPolicy",
	}
}
