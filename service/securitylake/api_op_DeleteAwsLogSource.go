// Code generated by smithy-go-codegen DO NOT EDIT.

package securitylake

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securitylake/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes a natively supported Amazon Web Service as an Amazon Security Lake
// source. You can remove a source for one or more Regions. When you remove the
// source, Security Lake stops collecting data from that source in the specified
// Regions and accounts, and subscribers can no longer consume new data from the
// source. However, subscribers can still consume data that Security Lake collected
// from the source before removal. You can choose any source type in any Amazon Web
// Services Region for either accounts that are part of a trusted organization or
// standalone accounts.
func (c *Client) DeleteAwsLogSource(ctx context.Context, params *DeleteAwsLogSourceInput, optFns ...func(*Options)) (*DeleteAwsLogSourceOutput, error) {
	if params == nil {
		params = &DeleteAwsLogSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAwsLogSource", params, optFns, c.addOperationDeleteAwsLogSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAwsLogSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAwsLogSourceInput struct {

	// Specify the natively-supported Amazon Web Services service to remove as a
	// source in Security Lake.
	//
	// This member is required.
	Sources []types.AwsLogSourceConfiguration

	noSmithyDocumentSerde
}

type DeleteAwsLogSourceOutput struct {

	// Deletion of the Amazon Web Services sources failed as the account is not a part
	// of the organization.
	Failed []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteAwsLogSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteAwsLogSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteAwsLogSource{}, middleware.After)
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
	if err = addOpDeleteAwsLogSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAwsLogSource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteAwsLogSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securitylake",
		OperationName: "DeleteAwsLogSource",
	}
}
