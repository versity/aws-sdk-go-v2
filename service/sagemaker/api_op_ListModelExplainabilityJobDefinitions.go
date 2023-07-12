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

// Lists model explainability job definitions that satisfy various filters.
func (c *Client) ListModelExplainabilityJobDefinitions(ctx context.Context, params *ListModelExplainabilityJobDefinitionsInput, optFns ...func(*Options)) (*ListModelExplainabilityJobDefinitionsOutput, error) {
	if params == nil {
		params = &ListModelExplainabilityJobDefinitionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListModelExplainabilityJobDefinitions", params, optFns, c.addOperationListModelExplainabilityJobDefinitionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListModelExplainabilityJobDefinitionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListModelExplainabilityJobDefinitionsInput struct {

	// A filter that returns only model explainability jobs created after a specified
	// time.
	CreationTimeAfter *time.Time

	// A filter that returns only model explainability jobs created before a specified
	// time.
	CreationTimeBefore *time.Time

	// Name of the endpoint to monitor for model explainability.
	EndpointName *string

	// The maximum number of jobs to return in the response. The default value is 10.
	MaxResults *int32

	// Filter for model explainability jobs whose name contains a specified string.
	NameContains *string

	// The token returned if the response is truncated. To retrieve the next set of
	// job executions, use it in the next request.
	NextToken *string

	// Whether to sort results by the Name or CreationTime field. The default is
	// CreationTime .
	SortBy types.MonitoringJobDefinitionSortKey

	// Whether to sort the results in Ascending or Descending order. The default is
	// Descending .
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListModelExplainabilityJobDefinitionsOutput struct {

	// A JSON array in which each element is a summary for a explainability bias jobs.
	//
	// This member is required.
	JobDefinitionSummaries []types.MonitoringJobDefinitionSummary

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of jobs, use it in the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListModelExplainabilityJobDefinitionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListModelExplainabilityJobDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListModelExplainabilityJobDefinitions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListModelExplainabilityJobDefinitions(options.Region), middleware.Before); err != nil {
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

// ListModelExplainabilityJobDefinitionsAPIClient is a client that implements the
// ListModelExplainabilityJobDefinitions operation.
type ListModelExplainabilityJobDefinitionsAPIClient interface {
	ListModelExplainabilityJobDefinitions(context.Context, *ListModelExplainabilityJobDefinitionsInput, ...func(*Options)) (*ListModelExplainabilityJobDefinitionsOutput, error)
}

var _ ListModelExplainabilityJobDefinitionsAPIClient = (*Client)(nil)

// ListModelExplainabilityJobDefinitionsPaginatorOptions is the paginator options
// for ListModelExplainabilityJobDefinitions
type ListModelExplainabilityJobDefinitionsPaginatorOptions struct {
	// The maximum number of jobs to return in the response. The default value is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListModelExplainabilityJobDefinitionsPaginator is a paginator for
// ListModelExplainabilityJobDefinitions
type ListModelExplainabilityJobDefinitionsPaginator struct {
	options   ListModelExplainabilityJobDefinitionsPaginatorOptions
	client    ListModelExplainabilityJobDefinitionsAPIClient
	params    *ListModelExplainabilityJobDefinitionsInput
	nextToken *string
	firstPage bool
}

// NewListModelExplainabilityJobDefinitionsPaginator returns a new
// ListModelExplainabilityJobDefinitionsPaginator
func NewListModelExplainabilityJobDefinitionsPaginator(client ListModelExplainabilityJobDefinitionsAPIClient, params *ListModelExplainabilityJobDefinitionsInput, optFns ...func(*ListModelExplainabilityJobDefinitionsPaginatorOptions)) *ListModelExplainabilityJobDefinitionsPaginator {
	if params == nil {
		params = &ListModelExplainabilityJobDefinitionsInput{}
	}

	options := ListModelExplainabilityJobDefinitionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListModelExplainabilityJobDefinitionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListModelExplainabilityJobDefinitionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListModelExplainabilityJobDefinitions page.
func (p *ListModelExplainabilityJobDefinitionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListModelExplainabilityJobDefinitionsOutput, error) {
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

	result, err := p.client.ListModelExplainabilityJobDefinitions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListModelExplainabilityJobDefinitions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListModelExplainabilityJobDefinitions",
	}
}
