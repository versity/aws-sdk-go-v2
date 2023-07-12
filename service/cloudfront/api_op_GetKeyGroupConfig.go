// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a key group configuration. To get a key group configuration, you must
// provide the key group's identifier. If the key group is referenced in a
// distribution's cache behavior, you can get the key group's identifier using
// ListDistributions or GetDistribution . If the key group is not referenced in a
// cache behavior, you can get the identifier using ListKeyGroups .
func (c *Client) GetKeyGroupConfig(ctx context.Context, params *GetKeyGroupConfigInput, optFns ...func(*Options)) (*GetKeyGroupConfigOutput, error) {
	if params == nil {
		params = &GetKeyGroupConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetKeyGroupConfig", params, optFns, c.addOperationGetKeyGroupConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetKeyGroupConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetKeyGroupConfigInput struct {

	// The identifier of the key group whose configuration you are getting. To get the
	// identifier, use ListKeyGroups .
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type GetKeyGroupConfigOutput struct {

	// The identifier for this version of the key group.
	ETag *string

	// The key group configuration.
	KeyGroupConfig *types.KeyGroupConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetKeyGroupConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetKeyGroupConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetKeyGroupConfig{}, middleware.After)
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
	if err = addOpGetKeyGroupConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetKeyGroupConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetKeyGroupConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "GetKeyGroupConfig",
	}
}
