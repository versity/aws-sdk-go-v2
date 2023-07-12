// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists all the experiments in your account. The list can be filtered to show
// only experiments that were created in a specific time range. The list can be
// sorted by experiment name or creation time.
func (c *Client) ListExperiments(ctx context.Context, params *ListExperimentsInput, optFns ...func(*Options)) (*ListExperimentsOutput, error) {
	if params == nil {
		params = &ListExperimentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListExperiments", params, optFns, c.addOperationListExperimentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListExperimentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListExperimentsInput struct {

	// A filter that returns only experiments created after the specified time.
	CreatedAfter *time.Time

	// A filter that returns only experiments created before the specified time.
	CreatedBefore *time.Time

	// The maximum number of experiments to return in the response. The default value
	// is 10.
	MaxResults *int32

	// If the previous call to ListExperiments didn't return the full set of
	// experiments, the call returns a token for getting the next set of experiments.
	NextToken *string

	// The property used to sort results. The default value is CreationTime .
	SortBy types.SortExperimentsBy

	// The sort order. The default value is Descending .
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListExperimentsOutput struct {

	// A list of the summaries of your experiments.
	ExperimentSummaries []types.ExperimentSummary

	// A token for getting the next set of experiments, if there are any.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListExperimentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListExperiments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListExperiments{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListExperiments(options.Region), middleware.Before); err != nil {
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

// ListExperimentsAPIClient is a client that implements the ListExperiments
// operation.
type ListExperimentsAPIClient interface {
	ListExperiments(context.Context, *ListExperimentsInput, ...func(*Options)) (*ListExperimentsOutput, error)
}

var _ ListExperimentsAPIClient = (*Client)(nil)

// ListExperimentsPaginatorOptions is the paginator options for ListExperiments
type ListExperimentsPaginatorOptions struct {
	// The maximum number of experiments to return in the response. The default value
	// is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListExperimentsPaginator is a paginator for ListExperiments
type ListExperimentsPaginator struct {
	options   ListExperimentsPaginatorOptions
	client    ListExperimentsAPIClient
	params    *ListExperimentsInput
	nextToken *string
	firstPage bool
}

// NewListExperimentsPaginator returns a new ListExperimentsPaginator
func NewListExperimentsPaginator(client ListExperimentsAPIClient, params *ListExperimentsInput, optFns ...func(*ListExperimentsPaginatorOptions)) *ListExperimentsPaginator {
	if params == nil {
		params = &ListExperimentsInput{}
	}

	options := ListExperimentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListExperimentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListExperimentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListExperiments page.
func (p *ListExperimentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListExperimentsOutput, error) {
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

	result, err := p.client.ListExperiments(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListExperiments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListExperiments",
	}
}
