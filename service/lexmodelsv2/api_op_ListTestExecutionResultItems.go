// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of test execution result items.
func (c *Client) ListTestExecutionResultItems(ctx context.Context, params *ListTestExecutionResultItemsInput, optFns ...func(*Options)) (*ListTestExecutionResultItemsOutput, error) {
	if params == nil {
		params = &ListTestExecutionResultItemsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTestExecutionResultItems", params, optFns, c.addOperationListTestExecutionResultItemsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTestExecutionResultItemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTestExecutionResultItemsInput struct {

	// The filter for the list of results from the test set execution.
	//
	// This member is required.
	ResultFilterBy *types.TestExecutionResultFilterBy

	// The unique identifier of the test execution to list the result items.
	//
	// This member is required.
	TestExecutionId *string

	// The maximum number of test execution result items to return in each page. If
	// there are fewer results than the max page size, only the actual number of
	// results are returned.
	MaxResults *int32

	// If the response from the ListTestExecutionResultItems operation contains more
	// results than specified in the maxResults parameter, a token is returned in the
	// response. Use that token in the nextToken parameter to return the next page of
	// results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListTestExecutionResultItemsOutput struct {

	// A token that indicates whether there are more results to return in a response
	// to the ListTestExecutionResultItems operation. If the nextToken field is
	// present, you send the contents as the nextToken parameter of a
	// ListTestExecutionResultItems operation request to get the next page of results.
	NextToken *string

	// The list of results from the test execution.
	TestExecutionResults *types.TestExecutionResultItems

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTestExecutionResultItemsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTestExecutionResultItems{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTestExecutionResultItems{}, middleware.After)
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
	if err = addOpListTestExecutionResultItemsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTestExecutionResultItems(options.Region), middleware.Before); err != nil {
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

// ListTestExecutionResultItemsAPIClient is a client that implements the
// ListTestExecutionResultItems operation.
type ListTestExecutionResultItemsAPIClient interface {
	ListTestExecutionResultItems(context.Context, *ListTestExecutionResultItemsInput, ...func(*Options)) (*ListTestExecutionResultItemsOutput, error)
}

var _ ListTestExecutionResultItemsAPIClient = (*Client)(nil)

// ListTestExecutionResultItemsPaginatorOptions is the paginator options for
// ListTestExecutionResultItems
type ListTestExecutionResultItemsPaginatorOptions struct {
	// The maximum number of test execution result items to return in each page. If
	// there are fewer results than the max page size, only the actual number of
	// results are returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTestExecutionResultItemsPaginator is a paginator for
// ListTestExecutionResultItems
type ListTestExecutionResultItemsPaginator struct {
	options   ListTestExecutionResultItemsPaginatorOptions
	client    ListTestExecutionResultItemsAPIClient
	params    *ListTestExecutionResultItemsInput
	nextToken *string
	firstPage bool
}

// NewListTestExecutionResultItemsPaginator returns a new
// ListTestExecutionResultItemsPaginator
func NewListTestExecutionResultItemsPaginator(client ListTestExecutionResultItemsAPIClient, params *ListTestExecutionResultItemsInput, optFns ...func(*ListTestExecutionResultItemsPaginatorOptions)) *ListTestExecutionResultItemsPaginator {
	if params == nil {
		params = &ListTestExecutionResultItemsInput{}
	}

	options := ListTestExecutionResultItemsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTestExecutionResultItemsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTestExecutionResultItemsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTestExecutionResultItems page.
func (p *ListTestExecutionResultItemsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTestExecutionResultItemsOutput, error) {
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

	result, err := p.client.ListTestExecutionResultItems(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTestExecutionResultItems(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "ListTestExecutionResultItems",
	}
}
