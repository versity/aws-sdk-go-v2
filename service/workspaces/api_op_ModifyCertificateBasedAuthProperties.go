// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the properties of the certificate-based authentication you want to use
// with your WorkSpaces.
func (c *Client) ModifyCertificateBasedAuthProperties(ctx context.Context, params *ModifyCertificateBasedAuthPropertiesInput, optFns ...func(*Options)) (*ModifyCertificateBasedAuthPropertiesOutput, error) {
	if params == nil {
		params = &ModifyCertificateBasedAuthPropertiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyCertificateBasedAuthProperties", params, optFns, c.addOperationModifyCertificateBasedAuthPropertiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyCertificateBasedAuthPropertiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyCertificateBasedAuthPropertiesInput struct {

	// The resource identifiers, in the form of directory IDs.
	//
	// This member is required.
	ResourceId *string

	// The properties of the certificate-based authentication.
	CertificateBasedAuthProperties *types.CertificateBasedAuthProperties

	// The properties of the certificate-based authentication you want to delete.
	PropertiesToDelete []types.DeletableCertificateBasedAuthProperty

	noSmithyDocumentSerde
}

type ModifyCertificateBasedAuthPropertiesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyCertificateBasedAuthPropertiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpModifyCertificateBasedAuthProperties{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyCertificateBasedAuthProperties{}, middleware.After)
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
	if err = addOpModifyCertificateBasedAuthPropertiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyCertificateBasedAuthProperties(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyCertificateBasedAuthProperties(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "ModifyCertificateBasedAuthProperties",
	}
}
