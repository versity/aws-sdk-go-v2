// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates all resource shares that use a managed permission to a different
// managed permission. This operation always applies the default version of the
// target managed permission. You can optionally specify that the update applies to
// only resource shares that currently use a specified version. This enables you to
// update to the latest version, without changing the which managed permission is
// used. You can use this operation to update all of your resource shares to use
// the current default version of the permission by specifying the same value for
// the fromPermissionArn and toPermissionArn parameters. You can use the optional
// fromPermissionVersion parameter to update only those resources that use a
// specified version of the managed permission to the new managed permission. To
// successfully perform this operation, you must have permission to update the
// resource-based policy on all affected resource types.
func (c *Client) ReplacePermissionAssociations(ctx context.Context, params *ReplacePermissionAssociationsInput, optFns ...func(*Options)) (*ReplacePermissionAssociationsOutput, error) {
	if params == nil {
		params = &ReplacePermissionAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReplacePermissionAssociations", params, optFns, c.addOperationReplacePermissionAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ReplacePermissionAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReplacePermissionAssociationsInput struct {

	// Specifies the Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the managed permission that you want to replace.
	//
	// This member is required.
	FromPermissionArn *string

	// Specifies the ARN of the managed permission that you want to associate with
	// resource shares in place of the one specified by fromPerssionArn and
	// fromPermissionVersion . The operation always associates the version that is
	// currently the default for the specified managed permission.
	//
	// This member is required.
	ToPermissionArn *string

	// Specifies a unique, case-sensitive identifier that you provide to ensure the
	// idempotency of the request. This lets you safely retry the request without
	// accidentally performing the same operation a second time. Passing the same value
	// to a later call to an operation requires that you also pass the same value for
	// all other parameters. We recommend that you use a UUID type of value. (https://wikipedia.org/wiki/Universally_unique_identifier)
	// . If you don't provide this value, then Amazon Web Services generates a random
	// one for you. If you retry the operation with the same ClientToken , but with
	// different parameters, the retry fails with an IdempotentParameterMismatch error.
	ClientToken *string

	// Specifies that you want to updated the permissions for only those resource
	// shares that use the specified version of the managed permission.
	FromPermissionVersion *int32

	noSmithyDocumentSerde
}

type ReplacePermissionAssociationsOutput struct {

	// The idempotency identifier associated with this request. If you want to repeat
	// the same operation in an idempotent manner then you must include this value in
	// the clientToken request parameter of that later call. All other parameters must
	// also have the same values that you used in the first call.
	ClientToken *string

	// Specifies a data structure that you can use to track the asynchronous tasks
	// that RAM performs to complete this operation. You can use the
	// ListReplacePermissionAssociationsWork operation and pass the id value returned
	// in this structure.
	ReplacePermissionAssociationsWork *types.ReplacePermissionAssociationsWork

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationReplacePermissionAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpReplacePermissionAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpReplacePermissionAssociations{}, middleware.After)
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
	if err = addOpReplacePermissionAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opReplacePermissionAssociations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opReplacePermissionAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "ReplacePermissionAssociations",
	}
}
