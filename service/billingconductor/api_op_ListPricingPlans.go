// Code generated by smithy-go-codegen DO NOT EDIT.

package billingconductor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/billingconductor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// A paginated call to get pricing plans for the given billing period. If you
// don't provide a billing period, the current billing period is used.
func (c *Client) ListPricingPlans(ctx context.Context, params *ListPricingPlansInput, optFns ...func(*Options)) (*ListPricingPlansOutput, error) {
	if params == nil {
		params = &ListPricingPlansInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPricingPlans", params, optFns, c.addOperationListPricingPlansMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPricingPlansOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPricingPlansInput struct {

	// The preferred billing period to get pricing plan.
	BillingPeriod *string

	// A ListPricingPlansFilter that specifies the Amazon Resource Name (ARNs) of
	// pricing plans to retrieve pricing plans information.
	Filters *types.ListPricingPlansFilter

	// The maximum number of pricing plans to retrieve.
	MaxResults *int32

	// The pagination token that's used on subsequent call to get pricing plans.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPricingPlansOutput struct {

	// The billing period for which the described pricing plans are applicable.
	BillingPeriod *string

	// The pagination token that's used on subsequent calls to get pricing plans.
	NextToken *string

	// A list of PricingPlanListElement retrieved.
	PricingPlans []types.PricingPlanListElement

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPricingPlansMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPricingPlans{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPricingPlans{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPricingPlans(options.Region), middleware.Before); err != nil {
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

// ListPricingPlansAPIClient is a client that implements the ListPricingPlans
// operation.
type ListPricingPlansAPIClient interface {
	ListPricingPlans(context.Context, *ListPricingPlansInput, ...func(*Options)) (*ListPricingPlansOutput, error)
}

var _ ListPricingPlansAPIClient = (*Client)(nil)

// ListPricingPlansPaginatorOptions is the paginator options for ListPricingPlans
type ListPricingPlansPaginatorOptions struct {
	// The maximum number of pricing plans to retrieve.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPricingPlansPaginator is a paginator for ListPricingPlans
type ListPricingPlansPaginator struct {
	options   ListPricingPlansPaginatorOptions
	client    ListPricingPlansAPIClient
	params    *ListPricingPlansInput
	nextToken *string
	firstPage bool
}

// NewListPricingPlansPaginator returns a new ListPricingPlansPaginator
func NewListPricingPlansPaginator(client ListPricingPlansAPIClient, params *ListPricingPlansInput, optFns ...func(*ListPricingPlansPaginatorOptions)) *ListPricingPlansPaginator {
	if params == nil {
		params = &ListPricingPlansInput{}
	}

	options := ListPricingPlansPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPricingPlansPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPricingPlansPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPricingPlans page.
func (p *ListPricingPlansPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPricingPlansOutput, error) {
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

	result, err := p.client.ListPricingPlans(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPricingPlans(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "billingconductor",
		OperationName: "ListPricingPlans",
	}
}
