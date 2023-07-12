// Code generated by smithy-go-codegen DO NOT EDIT.

package panorama

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/panorama/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of package import jobs.
func (c *Client) ListPackageImportJobs(ctx context.Context, params *ListPackageImportJobsInput, optFns ...func(*Options)) (*ListPackageImportJobsOutput, error) {
	if params == nil {
		params = &ListPackageImportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPackageImportJobs", params, optFns, c.addOperationListPackageImportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPackageImportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPackageImportJobsInput struct {

	// The maximum number of package import jobs to return in one page of results.
	MaxResults int32

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPackageImportJobsOutput struct {

	// A list of package import jobs.
	//
	// This member is required.
	PackageImportJobs []types.PackageImportJob

	// A pagination token that's included if more results are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPackageImportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPackageImportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPackageImportJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPackageImportJobs(options.Region), middleware.Before); err != nil {
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

// ListPackageImportJobsAPIClient is a client that implements the
// ListPackageImportJobs operation.
type ListPackageImportJobsAPIClient interface {
	ListPackageImportJobs(context.Context, *ListPackageImportJobsInput, ...func(*Options)) (*ListPackageImportJobsOutput, error)
}

var _ ListPackageImportJobsAPIClient = (*Client)(nil)

// ListPackageImportJobsPaginatorOptions is the paginator options for
// ListPackageImportJobs
type ListPackageImportJobsPaginatorOptions struct {
	// The maximum number of package import jobs to return in one page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPackageImportJobsPaginator is a paginator for ListPackageImportJobs
type ListPackageImportJobsPaginator struct {
	options   ListPackageImportJobsPaginatorOptions
	client    ListPackageImportJobsAPIClient
	params    *ListPackageImportJobsInput
	nextToken *string
	firstPage bool
}

// NewListPackageImportJobsPaginator returns a new ListPackageImportJobsPaginator
func NewListPackageImportJobsPaginator(client ListPackageImportJobsAPIClient, params *ListPackageImportJobsInput, optFns ...func(*ListPackageImportJobsPaginatorOptions)) *ListPackageImportJobsPaginator {
	if params == nil {
		params = &ListPackageImportJobsInput{}
	}

	options := ListPackageImportJobsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPackageImportJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPackageImportJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPackageImportJobs page.
func (p *ListPackageImportJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPackageImportJobsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListPackageImportJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPackageImportJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "panorama",
		OperationName: "ListPackageImportJobs",
	}
}
