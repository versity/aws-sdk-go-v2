// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Resets the password for a service-specific credential. The new password is
// Amazon Web Services generated and cryptographically strong. It cannot be
// configured by the user. Resetting the password immediately invalidates the
// previous password associated with this user.
func (c *Client) ResetServiceSpecificCredential(ctx context.Context, params *ResetServiceSpecificCredentialInput, optFns ...func(*Options)) (*ResetServiceSpecificCredentialOutput, error) {
	if params == nil {
		params = &ResetServiceSpecificCredentialInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResetServiceSpecificCredential", params, optFns, c.addOperationResetServiceSpecificCredentialMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResetServiceSpecificCredentialOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResetServiceSpecificCredentialInput struct {

	// The unique identifier of the service-specific credential. This parameter allows
	// (through its regex pattern (http://wikipedia.org/wiki/regex) ) a string of
	// characters that can consist of any upper or lowercased letter or digit.
	//
	// This member is required.
	ServiceSpecificCredentialId *string

	// The name of the IAM user associated with the service-specific credential. If
	// this value is not specified, then the operation assumes the user whose
	// credentials are used to call the operation. This parameter allows (through its
	// regex pattern (http://wikipedia.org/wiki/regex) ) a string of characters
	// consisting of upper and lowercase alphanumeric characters with no spaces. You
	// can also include any of the following characters: _+=,.@-
	UserName *string

	noSmithyDocumentSerde
}

type ResetServiceSpecificCredentialOutput struct {

	// A structure with details about the updated service-specific credential,
	// including the new password. This is the only time that you can access the
	// password. You cannot recover the password later, but you can reset it again.
	ServiceSpecificCredential *types.ServiceSpecificCredential

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationResetServiceSpecificCredentialMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpResetServiceSpecificCredential{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpResetServiceSpecificCredential{}, middleware.After)
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
	if err = addOpResetServiceSpecificCredentialValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opResetServiceSpecificCredential(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opResetServiceSpecificCredential(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "ResetServiceSpecificCredential",
	}
}
