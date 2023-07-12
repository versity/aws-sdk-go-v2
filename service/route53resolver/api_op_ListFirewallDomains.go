// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the domains that you have defined for the specified firewall domain
// list. A single call might return only a partial list of the domains. For
// information, see MaxResults .
func (c *Client) ListFirewallDomains(ctx context.Context, params *ListFirewallDomainsInput, optFns ...func(*Options)) (*ListFirewallDomainsOutput, error) {
	if params == nil {
		params = &ListFirewallDomainsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFirewallDomains", params, optFns, c.addOperationListFirewallDomainsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFirewallDomainsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFirewallDomainsInput struct {

	// The ID of the domain list whose domains you want to retrieve.
	//
	// This member is required.
	FirewallDomainListId *string

	// The maximum number of objects that you want Resolver to return for this
	// request. If more objects are available, in the response, Resolver provides a
	// NextToken value that you can use in a subsequent call to get the next batch of
	// objects. If you don't specify a value for MaxResults , Resolver returns up to
	// 100 objects.
	MaxResults *int32

	// For the first call to this list request, omit this value. When you request a
	// list of objects, Resolver returns at most the number of objects specified in
	// MaxResults . If more objects are available for retrieval, Resolver returns a
	// NextToken value in the response. To retrieve the next batch of objects, use the
	// token that was returned for the prior request in your next request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListFirewallDomainsOutput struct {

	// A list of the domains in the firewall domain list. This might be a partial list
	// of the domains that you've defined in the domain list. For information, see
	// MaxResults .
	Domains []string

	// If objects are still available for retrieval, Resolver returns this token in
	// the response. To retrieve the next batch of objects, provide this token in your
	// next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFirewallDomainsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListFirewallDomains{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFirewallDomains{}, middleware.After)
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
	if err = addOpListFirewallDomainsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFirewallDomains(options.Region), middleware.Before); err != nil {
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

// ListFirewallDomainsAPIClient is a client that implements the
// ListFirewallDomains operation.
type ListFirewallDomainsAPIClient interface {
	ListFirewallDomains(context.Context, *ListFirewallDomainsInput, ...func(*Options)) (*ListFirewallDomainsOutput, error)
}

var _ ListFirewallDomainsAPIClient = (*Client)(nil)

// ListFirewallDomainsPaginatorOptions is the paginator options for
// ListFirewallDomains
type ListFirewallDomainsPaginatorOptions struct {
	// The maximum number of objects that you want Resolver to return for this
	// request. If more objects are available, in the response, Resolver provides a
	// NextToken value that you can use in a subsequent call to get the next batch of
	// objects. If you don't specify a value for MaxResults , Resolver returns up to
	// 100 objects.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFirewallDomainsPaginator is a paginator for ListFirewallDomains
type ListFirewallDomainsPaginator struct {
	options   ListFirewallDomainsPaginatorOptions
	client    ListFirewallDomainsAPIClient
	params    *ListFirewallDomainsInput
	nextToken *string
	firstPage bool
}

// NewListFirewallDomainsPaginator returns a new ListFirewallDomainsPaginator
func NewListFirewallDomainsPaginator(client ListFirewallDomainsAPIClient, params *ListFirewallDomainsInput, optFns ...func(*ListFirewallDomainsPaginatorOptions)) *ListFirewallDomainsPaginator {
	if params == nil {
		params = &ListFirewallDomainsInput{}
	}

	options := ListFirewallDomainsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFirewallDomainsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFirewallDomainsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFirewallDomains page.
func (p *ListFirewallDomainsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFirewallDomainsOutput, error) {
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

	result, err := p.client.ListFirewallDomains(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFirewallDomains(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "ListFirewallDomains",
	}
}
