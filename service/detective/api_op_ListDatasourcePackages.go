// Code generated by smithy-go-codegen DO NOT EDIT.

package detective

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/detective/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists data source packages in the behavior graph.
func (c *Client) ListDatasourcePackages(ctx context.Context, params *ListDatasourcePackagesInput, optFns ...func(*Options)) (*ListDatasourcePackagesOutput, error) {
	if params == nil {
		params = &ListDatasourcePackagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDatasourcePackages", params, optFns, c.addOperationListDatasourcePackagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDatasourcePackagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDatasourcePackagesInput struct {

	// The ARN of the behavior graph.
	//
	// This member is required.
	GraphArn *string

	// The maximum number of results to return.
	MaxResults *int32

	// For requests to get the next page of results, the pagination token that was
	// returned with the previous set of results. The initial request does not include
	// a pagination token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDatasourcePackagesOutput struct {

	// Details on the data source packages active in the behavior graph.
	DatasourcePackages map[string]types.DatasourcePackageIngestDetail

	// For requests to get the next page of results, the pagination token that was
	// returned with the previous set of results. The initial request does not include
	// a pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDatasourcePackagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDatasourcePackages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDatasourcePackages{}, middleware.After)
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
	if err = addOpListDatasourcePackagesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDatasourcePackages(options.Region), middleware.Before); err != nil {
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

// ListDatasourcePackagesAPIClient is a client that implements the
// ListDatasourcePackages operation.
type ListDatasourcePackagesAPIClient interface {
	ListDatasourcePackages(context.Context, *ListDatasourcePackagesInput, ...func(*Options)) (*ListDatasourcePackagesOutput, error)
}

var _ ListDatasourcePackagesAPIClient = (*Client)(nil)

// ListDatasourcePackagesPaginatorOptions is the paginator options for
// ListDatasourcePackages
type ListDatasourcePackagesPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDatasourcePackagesPaginator is a paginator for ListDatasourcePackages
type ListDatasourcePackagesPaginator struct {
	options   ListDatasourcePackagesPaginatorOptions
	client    ListDatasourcePackagesAPIClient
	params    *ListDatasourcePackagesInput
	nextToken *string
	firstPage bool
}

// NewListDatasourcePackagesPaginator returns a new ListDatasourcePackagesPaginator
func NewListDatasourcePackagesPaginator(client ListDatasourcePackagesAPIClient, params *ListDatasourcePackagesInput, optFns ...func(*ListDatasourcePackagesPaginatorOptions)) *ListDatasourcePackagesPaginator {
	if params == nil {
		params = &ListDatasourcePackagesInput{}
	}

	options := ListDatasourcePackagesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDatasourcePackagesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDatasourcePackagesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDatasourcePackages page.
func (p *ListDatasourcePackagesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDatasourcePackagesOutput, error) {
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

	result, err := p.client.ListDatasourcePackages(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDatasourcePackages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "detective",
		OperationName: "ListDatasourcePackages",
	}
}
