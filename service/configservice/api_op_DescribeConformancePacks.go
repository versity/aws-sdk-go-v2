// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of one or more conformance packs.
func (c *Client) DescribeConformancePacks(ctx context.Context, params *DescribeConformancePacksInput, optFns ...func(*Options)) (*DescribeConformancePacksOutput, error) {
	if params == nil {
		params = &DescribeConformancePacksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConformancePacks", params, optFns, c.addOperationDescribeConformancePacksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConformancePacksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConformancePacksInput struct {

	// Comma-separated list of conformance pack names for which you want details. If
	// you do not specify any names, Config returns details for all your conformance
	// packs.
	ConformancePackNames []string

	// The maximum number of conformance packs returned on each page.
	Limit int32

	// The nextToken string returned in a previous request that you use to request the
	// next page of results in a paginated response.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeConformancePacksOutput struct {

	// Returns a list of ConformancePackDetail objects.
	ConformancePackDetails []types.ConformancePackDetail

	// The nextToken string returned in a previous request that you use to request the
	// next page of results in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeConformancePacksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeConformancePacks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeConformancePacks{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConformancePacks(options.Region), middleware.Before); err != nil {
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

// DescribeConformancePacksAPIClient is a client that implements the
// DescribeConformancePacks operation.
type DescribeConformancePacksAPIClient interface {
	DescribeConformancePacks(context.Context, *DescribeConformancePacksInput, ...func(*Options)) (*DescribeConformancePacksOutput, error)
}

var _ DescribeConformancePacksAPIClient = (*Client)(nil)

// DescribeConformancePacksPaginatorOptions is the paginator options for
// DescribeConformancePacks
type DescribeConformancePacksPaginatorOptions struct {
	// The maximum number of conformance packs returned on each page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeConformancePacksPaginator is a paginator for DescribeConformancePacks
type DescribeConformancePacksPaginator struct {
	options   DescribeConformancePacksPaginatorOptions
	client    DescribeConformancePacksAPIClient
	params    *DescribeConformancePacksInput
	nextToken *string
	firstPage bool
}

// NewDescribeConformancePacksPaginator returns a new
// DescribeConformancePacksPaginator
func NewDescribeConformancePacksPaginator(client DescribeConformancePacksAPIClient, params *DescribeConformancePacksInput, optFns ...func(*DescribeConformancePacksPaginatorOptions)) *DescribeConformancePacksPaginator {
	if params == nil {
		params = &DescribeConformancePacksInput{}
	}

	options := DescribeConformancePacksPaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeConformancePacksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeConformancePacksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeConformancePacks page.
func (p *DescribeConformancePacksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeConformancePacksOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.DescribeConformancePacks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeConformancePacks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribeConformancePacks",
	}
}
