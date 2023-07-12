// Code generated by smithy-go-codegen DO NOT EDIT.

package kendraranking

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendraranking/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists your rescore execution plans. A rescore execution plan is an Amazon
// Kendra Intelligent Ranking resource used for provisioning the Rescore API.
func (c *Client) ListRescoreExecutionPlans(ctx context.Context, params *ListRescoreExecutionPlansInput, optFns ...func(*Options)) (*ListRescoreExecutionPlansOutput, error) {
	if params == nil {
		params = &ListRescoreExecutionPlansInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRescoreExecutionPlans", params, optFns, c.addOperationListRescoreExecutionPlansMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRescoreExecutionPlansOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRescoreExecutionPlansInput struct {

	// The maximum number of rescore execution plans to return.
	MaxResults *int32

	// If the response is truncated, Amazon Kendra Intelligent Ranking returns a
	// pagination token in the response. You can use this pagination token to retrieve
	// the next set of rescore execution plans.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRescoreExecutionPlansOutput struct {

	// If the response is truncated, Amazon Kendra Intelligent Ranking returns a
	// pagination token in the response.
	NextToken *string

	// An array of summary information for one or more rescore execution plans.
	SummaryItems []types.RescoreExecutionPlanSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRescoreExecutionPlansMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListRescoreExecutionPlans{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListRescoreExecutionPlans{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRescoreExecutionPlans(options.Region), middleware.Before); err != nil {
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

// ListRescoreExecutionPlansAPIClient is a client that implements the
// ListRescoreExecutionPlans operation.
type ListRescoreExecutionPlansAPIClient interface {
	ListRescoreExecutionPlans(context.Context, *ListRescoreExecutionPlansInput, ...func(*Options)) (*ListRescoreExecutionPlansOutput, error)
}

var _ ListRescoreExecutionPlansAPIClient = (*Client)(nil)

// ListRescoreExecutionPlansPaginatorOptions is the paginator options for
// ListRescoreExecutionPlans
type ListRescoreExecutionPlansPaginatorOptions struct {
	// The maximum number of rescore execution plans to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRescoreExecutionPlansPaginator is a paginator for ListRescoreExecutionPlans
type ListRescoreExecutionPlansPaginator struct {
	options   ListRescoreExecutionPlansPaginatorOptions
	client    ListRescoreExecutionPlansAPIClient
	params    *ListRescoreExecutionPlansInput
	nextToken *string
	firstPage bool
}

// NewListRescoreExecutionPlansPaginator returns a new
// ListRescoreExecutionPlansPaginator
func NewListRescoreExecutionPlansPaginator(client ListRescoreExecutionPlansAPIClient, params *ListRescoreExecutionPlansInput, optFns ...func(*ListRescoreExecutionPlansPaginatorOptions)) *ListRescoreExecutionPlansPaginator {
	if params == nil {
		params = &ListRescoreExecutionPlansInput{}
	}

	options := ListRescoreExecutionPlansPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRescoreExecutionPlansPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRescoreExecutionPlansPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRescoreExecutionPlans page.
func (p *ListRescoreExecutionPlansPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRescoreExecutionPlansOutput, error) {
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

	result, err := p.client.ListRescoreExecutionPlans(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRescoreExecutionPlans(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra-ranking",
		OperationName: "ListRescoreExecutionPlans",
	}
}
