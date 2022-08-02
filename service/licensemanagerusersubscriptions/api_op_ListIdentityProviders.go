// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanagerusersubscriptions

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/licensemanagerusersubscriptions/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the identity providers for user-based subscriptions.
func (c *Client) ListIdentityProviders(ctx context.Context, params *ListIdentityProvidersInput, optFns ...func(*Options)) (*ListIdentityProvidersOutput, error) {
	if params == nil {
		params = &ListIdentityProvidersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListIdentityProviders", params, optFns, c.addOperationListIdentityProvidersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListIdentityProvidersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListIdentityProvidersInput struct {

	// Maximum number of results to return in a single call.
	MaxResults *int32

	// Token for the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListIdentityProvidersOutput struct {

	// Metadata that describes the list identity providers operation.
	//
	// This member is required.
	IdentityProviderSummaries []types.IdentityProviderSummary

	// Token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListIdentityProvidersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListIdentityProviders{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListIdentityProviders{}, middleware.After)
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
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListIdentityProviders(options.Region), middleware.Before); err != nil {
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

// ListIdentityProvidersAPIClient is a client that implements the
// ListIdentityProviders operation.
type ListIdentityProvidersAPIClient interface {
	ListIdentityProviders(context.Context, *ListIdentityProvidersInput, ...func(*Options)) (*ListIdentityProvidersOutput, error)
}

var _ ListIdentityProvidersAPIClient = (*Client)(nil)

// ListIdentityProvidersPaginatorOptions is the paginator options for
// ListIdentityProviders
type ListIdentityProvidersPaginatorOptions struct {
	// Maximum number of results to return in a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListIdentityProvidersPaginator is a paginator for ListIdentityProviders
type ListIdentityProvidersPaginator struct {
	options   ListIdentityProvidersPaginatorOptions
	client    ListIdentityProvidersAPIClient
	params    *ListIdentityProvidersInput
	nextToken *string
	firstPage bool
}

// NewListIdentityProvidersPaginator returns a new ListIdentityProvidersPaginator
func NewListIdentityProvidersPaginator(client ListIdentityProvidersAPIClient, params *ListIdentityProvidersInput, optFns ...func(*ListIdentityProvidersPaginatorOptions)) *ListIdentityProvidersPaginator {
	if params == nil {
		params = &ListIdentityProvidersInput{}
	}

	options := ListIdentityProvidersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListIdentityProvidersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListIdentityProvidersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListIdentityProviders page.
func (p *ListIdentityProvidersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListIdentityProvidersOutput, error) {
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

	result, err := p.client.ListIdentityProviders(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListIdentityProviders(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager-user-subscriptions",
		OperationName: "ListIdentityProviders",
	}
}
