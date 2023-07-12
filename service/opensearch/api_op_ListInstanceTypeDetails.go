// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all instance types and available features for a given OpenSearch or
// Elasticsearch version.
func (c *Client) ListInstanceTypeDetails(ctx context.Context, params *ListInstanceTypeDetailsInput, optFns ...func(*Options)) (*ListInstanceTypeDetailsOutput, error) {
	if params == nil {
		params = &ListInstanceTypeDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInstanceTypeDetails", params, optFns, c.addOperationListInstanceTypeDetailsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInstanceTypeDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListInstanceTypeDetailsInput struct {

	// The version of OpenSearch or Elasticsearch, in the format Elasticsearch_X.Y or
	// OpenSearch_X.Y. Defaults to the latest version of OpenSearch.
	//
	// This member is required.
	EngineVersion *string

	// The name of the domain.
	DomainName *string

	// An optional parameter that lists information for a given instance type.
	InstanceType *string

	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results.
	MaxResults int32

	// If your initial ListInstanceTypeDetails operation returns a nextToken , you can
	// include the returned nextToken in subsequent ListInstanceTypeDetails
	// operations, which returns results in the next page.
	NextToken *string

	// An optional parameter that specifies the Availability Zones for the domain.
	RetrieveAZs *bool

	noSmithyDocumentSerde
}

type ListInstanceTypeDetailsOutput struct {

	// Lists all supported instance types and features for the given OpenSearch or
	// Elasticsearch version.
	InstanceTypeDetails []types.InstanceTypeDetails

	// When nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListInstanceTypeDetailsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListInstanceTypeDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListInstanceTypeDetails{}, middleware.After)
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
	if err = addOpListInstanceTypeDetailsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListInstanceTypeDetails(options.Region), middleware.Before); err != nil {
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

// ListInstanceTypeDetailsAPIClient is a client that implements the
// ListInstanceTypeDetails operation.
type ListInstanceTypeDetailsAPIClient interface {
	ListInstanceTypeDetails(context.Context, *ListInstanceTypeDetailsInput, ...func(*Options)) (*ListInstanceTypeDetailsOutput, error)
}

var _ ListInstanceTypeDetailsAPIClient = (*Client)(nil)

// ListInstanceTypeDetailsPaginatorOptions is the paginator options for
// ListInstanceTypeDetails
type ListInstanceTypeDetailsPaginatorOptions struct {
	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListInstanceTypeDetailsPaginator is a paginator for ListInstanceTypeDetails
type ListInstanceTypeDetailsPaginator struct {
	options   ListInstanceTypeDetailsPaginatorOptions
	client    ListInstanceTypeDetailsAPIClient
	params    *ListInstanceTypeDetailsInput
	nextToken *string
	firstPage bool
}

// NewListInstanceTypeDetailsPaginator returns a new
// ListInstanceTypeDetailsPaginator
func NewListInstanceTypeDetailsPaginator(client ListInstanceTypeDetailsAPIClient, params *ListInstanceTypeDetailsInput, optFns ...func(*ListInstanceTypeDetailsPaginatorOptions)) *ListInstanceTypeDetailsPaginator {
	if params == nil {
		params = &ListInstanceTypeDetailsInput{}
	}

	options := ListInstanceTypeDetailsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListInstanceTypeDetailsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListInstanceTypeDetailsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListInstanceTypeDetails page.
func (p *ListInstanceTypeDetailsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListInstanceTypeDetailsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListInstanceTypeDetails(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListInstanceTypeDetails(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "ListInstanceTypeDetails",
	}
}
