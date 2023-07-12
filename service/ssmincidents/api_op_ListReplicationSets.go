// Code generated by smithy-go-codegen DO NOT EDIT.

package ssmincidents

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists details about the replication set configured in your account.
func (c *Client) ListReplicationSets(ctx context.Context, params *ListReplicationSetsInput, optFns ...func(*Options)) (*ListReplicationSetsOutput, error) {
	if params == nil {
		params = &ListReplicationSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListReplicationSets", params, optFns, c.addOperationListReplicationSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListReplicationSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListReplicationSetsInput struct {

	// The maximum number of results per page.
	MaxResults *int32

	// The pagination token to continue to the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListReplicationSetsOutput struct {

	// The Amazon Resource Name (ARN) of the list replication set.
	//
	// This member is required.
	ReplicationSetArns []string

	// The pagination token to continue to the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListReplicationSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListReplicationSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListReplicationSets{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListReplicationSets(options.Region), middleware.Before); err != nil {
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

// ListReplicationSetsAPIClient is a client that implements the
// ListReplicationSets operation.
type ListReplicationSetsAPIClient interface {
	ListReplicationSets(context.Context, *ListReplicationSetsInput, ...func(*Options)) (*ListReplicationSetsOutput, error)
}

var _ ListReplicationSetsAPIClient = (*Client)(nil)

// ListReplicationSetsPaginatorOptions is the paginator options for
// ListReplicationSets
type ListReplicationSetsPaginatorOptions struct {
	// The maximum number of results per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListReplicationSetsPaginator is a paginator for ListReplicationSets
type ListReplicationSetsPaginator struct {
	options   ListReplicationSetsPaginatorOptions
	client    ListReplicationSetsAPIClient
	params    *ListReplicationSetsInput
	nextToken *string
	firstPage bool
}

// NewListReplicationSetsPaginator returns a new ListReplicationSetsPaginator
func NewListReplicationSetsPaginator(client ListReplicationSetsAPIClient, params *ListReplicationSetsInput, optFns ...func(*ListReplicationSetsPaginatorOptions)) *ListReplicationSetsPaginator {
	if params == nil {
		params = &ListReplicationSetsInput{}
	}

	options := ListReplicationSetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListReplicationSetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListReplicationSetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListReplicationSets page.
func (p *ListReplicationSetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListReplicationSetsOutput, error) {
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

	result, err := p.client.ListReplicationSets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListReplicationSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm-incidents",
		OperationName: "ListReplicationSets",
	}
}
