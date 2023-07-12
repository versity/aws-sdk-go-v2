// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalogappregistry

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the details of all attribute groups associated with a specific
// application. The results display in pages.
func (c *Client) ListAttributeGroupsForApplication(ctx context.Context, params *ListAttributeGroupsForApplicationInput, optFns ...func(*Options)) (*ListAttributeGroupsForApplicationOutput, error) {
	if params == nil {
		params = &ListAttributeGroupsForApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAttributeGroupsForApplication", params, optFns, c.addOperationListAttributeGroupsForApplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAttributeGroupsForApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAttributeGroupsForApplicationInput struct {

	// The name or ID of the application.
	//
	// This member is required.
	Application *string

	// The upper bound of the number of results to return. The value cannot exceed 25.
	// If you omit this parameter, it defaults to 25. This value is optional.
	MaxResults *int32

	// This token retrieves the next page of results after a previous API call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAttributeGroupsForApplicationOutput struct {

	// The details related to a specific attribute group.
	AttributeGroupsDetails []types.AttributeGroupDetails

	// The token to use to get the next page of results after a previous API call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAttributeGroupsForApplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAttributeGroupsForApplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAttributeGroupsForApplication{}, middleware.After)
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
	if err = addOpListAttributeGroupsForApplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAttributeGroupsForApplication(options.Region), middleware.Before); err != nil {
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

// ListAttributeGroupsForApplicationAPIClient is a client that implements the
// ListAttributeGroupsForApplication operation.
type ListAttributeGroupsForApplicationAPIClient interface {
	ListAttributeGroupsForApplication(context.Context, *ListAttributeGroupsForApplicationInput, ...func(*Options)) (*ListAttributeGroupsForApplicationOutput, error)
}

var _ ListAttributeGroupsForApplicationAPIClient = (*Client)(nil)

// ListAttributeGroupsForApplicationPaginatorOptions is the paginator options for
// ListAttributeGroupsForApplication
type ListAttributeGroupsForApplicationPaginatorOptions struct {
	// The upper bound of the number of results to return. The value cannot exceed 25.
	// If you omit this parameter, it defaults to 25. This value is optional.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAttributeGroupsForApplicationPaginator is a paginator for
// ListAttributeGroupsForApplication
type ListAttributeGroupsForApplicationPaginator struct {
	options   ListAttributeGroupsForApplicationPaginatorOptions
	client    ListAttributeGroupsForApplicationAPIClient
	params    *ListAttributeGroupsForApplicationInput
	nextToken *string
	firstPage bool
}

// NewListAttributeGroupsForApplicationPaginator returns a new
// ListAttributeGroupsForApplicationPaginator
func NewListAttributeGroupsForApplicationPaginator(client ListAttributeGroupsForApplicationAPIClient, params *ListAttributeGroupsForApplicationInput, optFns ...func(*ListAttributeGroupsForApplicationPaginatorOptions)) *ListAttributeGroupsForApplicationPaginator {
	if params == nil {
		params = &ListAttributeGroupsForApplicationInput{}
	}

	options := ListAttributeGroupsForApplicationPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAttributeGroupsForApplicationPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAttributeGroupsForApplicationPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAttributeGroupsForApplication page.
func (p *ListAttributeGroupsForApplicationPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAttributeGroupsForApplicationOutput, error) {
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

	result, err := p.client.ListAttributeGroupsForApplication(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAttributeGroupsForApplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "ListAttributeGroupsForApplication",
	}
}
