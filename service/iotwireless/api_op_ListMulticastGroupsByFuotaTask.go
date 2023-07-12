// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List all multicast groups associated with a fuota task.
func (c *Client) ListMulticastGroupsByFuotaTask(ctx context.Context, params *ListMulticastGroupsByFuotaTaskInput, optFns ...func(*Options)) (*ListMulticastGroupsByFuotaTaskOutput, error) {
	if params == nil {
		params = &ListMulticastGroupsByFuotaTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMulticastGroupsByFuotaTask", params, optFns, c.addOperationListMulticastGroupsByFuotaTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMulticastGroupsByFuotaTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMulticastGroupsByFuotaTaskInput struct {

	// The ID of a FUOTA task.
	//
	// This member is required.
	Id *string

	// The maximum number of results to return in this operation.
	MaxResults int32

	// To retrieve the next set of results, the nextToken value from a previous
	// response; otherwise null to receive the first set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListMulticastGroupsByFuotaTaskOutput struct {

	// List of multicast groups associated with a FUOTA task.
	MulticastGroupList []types.MulticastGroupByFuotaTask

	// To retrieve the next set of results, the nextToken value from a previous
	// response; otherwise null to receive the first set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMulticastGroupsByFuotaTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListMulticastGroupsByFuotaTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListMulticastGroupsByFuotaTask{}, middleware.After)
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
	if err = addOpListMulticastGroupsByFuotaTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMulticastGroupsByFuotaTask(options.Region), middleware.Before); err != nil {
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

// ListMulticastGroupsByFuotaTaskAPIClient is a client that implements the
// ListMulticastGroupsByFuotaTask operation.
type ListMulticastGroupsByFuotaTaskAPIClient interface {
	ListMulticastGroupsByFuotaTask(context.Context, *ListMulticastGroupsByFuotaTaskInput, ...func(*Options)) (*ListMulticastGroupsByFuotaTaskOutput, error)
}

var _ ListMulticastGroupsByFuotaTaskAPIClient = (*Client)(nil)

// ListMulticastGroupsByFuotaTaskPaginatorOptions is the paginator options for
// ListMulticastGroupsByFuotaTask
type ListMulticastGroupsByFuotaTaskPaginatorOptions struct {
	// The maximum number of results to return in this operation.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMulticastGroupsByFuotaTaskPaginator is a paginator for
// ListMulticastGroupsByFuotaTask
type ListMulticastGroupsByFuotaTaskPaginator struct {
	options   ListMulticastGroupsByFuotaTaskPaginatorOptions
	client    ListMulticastGroupsByFuotaTaskAPIClient
	params    *ListMulticastGroupsByFuotaTaskInput
	nextToken *string
	firstPage bool
}

// NewListMulticastGroupsByFuotaTaskPaginator returns a new
// ListMulticastGroupsByFuotaTaskPaginator
func NewListMulticastGroupsByFuotaTaskPaginator(client ListMulticastGroupsByFuotaTaskAPIClient, params *ListMulticastGroupsByFuotaTaskInput, optFns ...func(*ListMulticastGroupsByFuotaTaskPaginatorOptions)) *ListMulticastGroupsByFuotaTaskPaginator {
	if params == nil {
		params = &ListMulticastGroupsByFuotaTaskInput{}
	}

	options := ListMulticastGroupsByFuotaTaskPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMulticastGroupsByFuotaTaskPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMulticastGroupsByFuotaTaskPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMulticastGroupsByFuotaTask page.
func (p *ListMulticastGroupsByFuotaTaskPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMulticastGroupsByFuotaTaskOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListMulticastGroupsByFuotaTask(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMulticastGroupsByFuotaTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "ListMulticastGroupsByFuotaTask",
	}
}
