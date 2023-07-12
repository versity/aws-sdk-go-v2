// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is documentation for AWS CloudHSM Classic. For more information, see AWS
// CloudHSM Classic FAQs (http://aws.amazon.com/cloudhsm/faqs-classic/) , the AWS
// CloudHSM Classic User Guide (https://docs.aws.amazon.com/cloudhsm/classic/userguide/)
// , and the AWS CloudHSM Classic API Reference (https://docs.aws.amazon.com/cloudhsm/classic/APIReference/)
// . For information about the current version of AWS CloudHSM, see AWS CloudHSM (http://aws.amazon.com/cloudhsm/)
// , the AWS CloudHSM User Guide (https://docs.aws.amazon.com/cloudhsm/latest/userguide/)
// , and the AWS CloudHSM API Reference (https://docs.aws.amazon.com/cloudhsm/latest/APIReference/)
// . Lists all of the clients. This operation supports pagination with the use of
// the NextToken member. If more results are available, the NextToken member of
// the response contains a token that you pass in the next call to ListLunaClients
// to retrieve the next set of items.
func (c *Client) ListLunaClients(ctx context.Context, params *ListLunaClientsInput, optFns ...func(*Options)) (*ListLunaClientsOutput, error) {
	if params == nil {
		params = &ListLunaClientsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLunaClients", params, optFns, c.addOperationListLunaClientsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLunaClientsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLunaClientsInput struct {

	// The NextToken value from a previous call to ListLunaClients . Pass null if this
	// is the first call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListLunaClientsOutput struct {

	// The list of clients.
	//
	// This member is required.
	ClientList []string

	// If not null, more results are available. Pass this to ListLunaClients to
	// retrieve the next set of items.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLunaClientsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListLunaClients{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListLunaClients{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLunaClients(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListLunaClients(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudhsm",
		OperationName: "ListLunaClients",
	}
}
