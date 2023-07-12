// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes one or more of your job queues.
func (c *Client) DescribeJobQueues(ctx context.Context, params *DescribeJobQueuesInput, optFns ...func(*Options)) (*DescribeJobQueuesOutput, error) {
	if params == nil {
		params = &DescribeJobQueuesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeJobQueues", params, optFns, c.addOperationDescribeJobQueuesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeJobQueuesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeJobQueues .
type DescribeJobQueuesInput struct {

	// A list of up to 100 queue names or full queue Amazon Resource Name (ARN)
	// entries.
	JobQueues []string

	// The maximum number of results returned by DescribeJobQueues in paginated
	// output. When this parameter is used, DescribeJobQueues only returns maxResults
	// results in a single page and a nextToken response element. The remaining
	// results of the initial request can be seen by sending another DescribeJobQueues
	// request with the returned nextToken value. This value can be between 1 and 100.
	// If this parameter isn't used, then DescribeJobQueues returns up to 100 results
	// and a nextToken value if applicable.
	MaxResults *int32

	// The nextToken value returned from a previous paginated DescribeJobQueues
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. This value is null when there are no more results
	// to return. Treat this token as an opaque identifier that's only used to retrieve
	// the next items in a list and not for other programmatic purposes.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeJobQueuesOutput struct {

	// The list of job queues.
	JobQueues []types.JobQueueDetail

	// The nextToken value to include in a future DescribeJobQueues request. When the
	// results of a DescribeJobQueues request exceed maxResults , this value can be
	// used to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeJobQueuesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeJobQueues{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeJobQueues{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeJobQueues(options.Region), middleware.Before); err != nil {
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

// DescribeJobQueuesAPIClient is a client that implements the DescribeJobQueues
// operation.
type DescribeJobQueuesAPIClient interface {
	DescribeJobQueues(context.Context, *DescribeJobQueuesInput, ...func(*Options)) (*DescribeJobQueuesOutput, error)
}

var _ DescribeJobQueuesAPIClient = (*Client)(nil)

// DescribeJobQueuesPaginatorOptions is the paginator options for DescribeJobQueues
type DescribeJobQueuesPaginatorOptions struct {
	// The maximum number of results returned by DescribeJobQueues in paginated
	// output. When this parameter is used, DescribeJobQueues only returns maxResults
	// results in a single page and a nextToken response element. The remaining
	// results of the initial request can be seen by sending another DescribeJobQueues
	// request with the returned nextToken value. This value can be between 1 and 100.
	// If this parameter isn't used, then DescribeJobQueues returns up to 100 results
	// and a nextToken value if applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeJobQueuesPaginator is a paginator for DescribeJobQueues
type DescribeJobQueuesPaginator struct {
	options   DescribeJobQueuesPaginatorOptions
	client    DescribeJobQueuesAPIClient
	params    *DescribeJobQueuesInput
	nextToken *string
	firstPage bool
}

// NewDescribeJobQueuesPaginator returns a new DescribeJobQueuesPaginator
func NewDescribeJobQueuesPaginator(client DescribeJobQueuesAPIClient, params *DescribeJobQueuesInput, optFns ...func(*DescribeJobQueuesPaginatorOptions)) *DescribeJobQueuesPaginator {
	if params == nil {
		params = &DescribeJobQueuesInput{}
	}

	options := DescribeJobQueuesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeJobQueuesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeJobQueuesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeJobQueues page.
func (p *DescribeJobQueuesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeJobQueuesOutput, error) {
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

	result, err := p.client.DescribeJobQueues(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeJobQueues(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "batch",
		OperationName: "DescribeJobQueues",
	}
}
