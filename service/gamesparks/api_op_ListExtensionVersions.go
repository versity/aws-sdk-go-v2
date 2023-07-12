// Code generated by smithy-go-codegen DO NOT EDIT.

package gamesparks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamesparks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a paginated list of available versions for the extension. Each time an API
// change is made to an extension, the version is incremented. The list retrieved
// by this operation shows the versions that are currently available.
func (c *Client) ListExtensionVersions(ctx context.Context, params *ListExtensionVersionsInput, optFns ...func(*Options)) (*ListExtensionVersionsOutput, error) {
	if params == nil {
		params = &ListExtensionVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListExtensionVersions", params, optFns, c.addOperationListExtensionVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListExtensionVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListExtensionVersionsInput struct {

	// The name of the extension.
	//
	// This member is required.
	Name *string

	// The namespace (qualifier) of the extension.
	//
	// This member is required.
	Namespace *string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	MaxResults *int32

	// The token that indicates the start of the next sequential page of results. Use
	// the token that is returned with a previous call to this operation. To start at
	// the beginning of the result set, do not specify a value.
	NextToken *string

	noSmithyDocumentSerde
}

type ListExtensionVersionsOutput struct {

	// The list of extension versions.
	ExtensionVersions []types.ExtensionVersionDetails

	// The token that indicates the start of the next sequential page of results. Use
	// this value when making the next call to this operation to continue where the
	// last one finished.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListExtensionVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListExtensionVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListExtensionVersions{}, middleware.After)
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
	if err = addOpListExtensionVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListExtensionVersions(options.Region), middleware.Before); err != nil {
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

// ListExtensionVersionsAPIClient is a client that implements the
// ListExtensionVersions operation.
type ListExtensionVersionsAPIClient interface {
	ListExtensionVersions(context.Context, *ListExtensionVersionsInput, ...func(*Options)) (*ListExtensionVersionsOutput, error)
}

var _ ListExtensionVersionsAPIClient = (*Client)(nil)

// ListExtensionVersionsPaginatorOptions is the paginator options for
// ListExtensionVersions
type ListExtensionVersionsPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListExtensionVersionsPaginator is a paginator for ListExtensionVersions
type ListExtensionVersionsPaginator struct {
	options   ListExtensionVersionsPaginatorOptions
	client    ListExtensionVersionsAPIClient
	params    *ListExtensionVersionsInput
	nextToken *string
	firstPage bool
}

// NewListExtensionVersionsPaginator returns a new ListExtensionVersionsPaginator
func NewListExtensionVersionsPaginator(client ListExtensionVersionsAPIClient, params *ListExtensionVersionsInput, optFns ...func(*ListExtensionVersionsPaginatorOptions)) *ListExtensionVersionsPaginator {
	if params == nil {
		params = &ListExtensionVersionsInput{}
	}

	options := ListExtensionVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListExtensionVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListExtensionVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListExtensionVersions page.
func (p *ListExtensionVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListExtensionVersionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListExtensionVersions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListExtensionVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamesparks",
		OperationName: "ListExtensionVersions",
	}
}
