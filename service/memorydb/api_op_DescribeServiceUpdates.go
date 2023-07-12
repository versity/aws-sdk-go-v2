// Code generated by smithy-go-codegen DO NOT EDIT.

package memorydb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/memorydb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns details of the service updates
func (c *Client) DescribeServiceUpdates(ctx context.Context, params *DescribeServiceUpdatesInput, optFns ...func(*Options)) (*DescribeServiceUpdatesOutput, error) {
	if params == nil {
		params = &DescribeServiceUpdatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeServiceUpdates", params, optFns, c.addOperationDescribeServiceUpdatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeServiceUpdatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeServiceUpdatesInput struct {

	// The list of cluster names to identify service updates to apply
	ClusterNames []string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	MaxResults *int32

	// An optional argument to pass in case the total number of records exceeds the
	// value of MaxResults. If nextToken is returned, there are more results available.
	// The value of nextToken is a unique pagination token for each page. Make the call
	// again using the returned token to retrieve the next page. Keep all other
	// arguments unchanged.
	NextToken *string

	// The unique ID of the service update to describe.
	ServiceUpdateName *string

	// The status(es) of the service updates to filter on
	Status []types.ServiceUpdateStatus

	noSmithyDocumentSerde
}

type DescribeServiceUpdatesOutput struct {

	// An optional argument to pass in case the total number of records exceeds the
	// value of MaxResults. If nextToken is returned, there are more results available.
	// The value of nextToken is a unique pagination token for each page. Make the call
	// again using the returned token to retrieve the next page. Keep all other
	// arguments unchanged.
	NextToken *string

	// A list of service updates
	ServiceUpdates []types.ServiceUpdate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeServiceUpdatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeServiceUpdates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeServiceUpdates{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeServiceUpdates(options.Region), middleware.Before); err != nil {
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

// DescribeServiceUpdatesAPIClient is a client that implements the
// DescribeServiceUpdates operation.
type DescribeServiceUpdatesAPIClient interface {
	DescribeServiceUpdates(context.Context, *DescribeServiceUpdatesInput, ...func(*Options)) (*DescribeServiceUpdatesOutput, error)
}

var _ DescribeServiceUpdatesAPIClient = (*Client)(nil)

// DescribeServiceUpdatesPaginatorOptions is the paginator options for
// DescribeServiceUpdates
type DescribeServiceUpdatesPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeServiceUpdatesPaginator is a paginator for DescribeServiceUpdates
type DescribeServiceUpdatesPaginator struct {
	options   DescribeServiceUpdatesPaginatorOptions
	client    DescribeServiceUpdatesAPIClient
	params    *DescribeServiceUpdatesInput
	nextToken *string
	firstPage bool
}

// NewDescribeServiceUpdatesPaginator returns a new DescribeServiceUpdatesPaginator
func NewDescribeServiceUpdatesPaginator(client DescribeServiceUpdatesAPIClient, params *DescribeServiceUpdatesInput, optFns ...func(*DescribeServiceUpdatesPaginatorOptions)) *DescribeServiceUpdatesPaginator {
	if params == nil {
		params = &DescribeServiceUpdatesInput{}
	}

	options := DescribeServiceUpdatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeServiceUpdatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeServiceUpdatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeServiceUpdates page.
func (p *DescribeServiceUpdatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeServiceUpdatesOutput, error) {
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

	result, err := p.client.DescribeServiceUpdates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeServiceUpdates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "memorydb",
		OperationName: "DescribeServiceUpdates",
	}
}
