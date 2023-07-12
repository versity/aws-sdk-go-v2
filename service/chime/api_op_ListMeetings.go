// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists up to 100 active Amazon Chime SDK meetings. ListMeetings is not supported
// in the Amazon Chime SDK Meetings Namespace. Update your application to remove
// calls to this API. For more information about the Amazon Chime SDK, see Using
// the Amazon Chime SDK (https://docs.aws.amazon.com/chime-sdk/latest/dg/meetings-sdk.html)
// in the Amazon Chime SDK Developer Guide.
//
// Deprecated: ListMeetings is not supported in the Amazon Chime SDK Meetings
// Namespace. Update your application to remove calls to this API.
func (c *Client) ListMeetings(ctx context.Context, params *ListMeetingsInput, optFns ...func(*Options)) (*ListMeetingsOutput, error) {
	if params == nil {
		params = &ListMeetingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMeetings", params, optFns, c.addOperationListMeetingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMeetingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMeetingsInput struct {

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// The token to use to retrieve the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListMeetingsOutput struct {

	// The Amazon Chime SDK meeting information.
	Meetings []types.Meeting

	// The token to use to retrieve the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMeetingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListMeetings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListMeetings{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMeetings(options.Region), middleware.Before); err != nil {
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

// ListMeetingsAPIClient is a client that implements the ListMeetings operation.
type ListMeetingsAPIClient interface {
	ListMeetings(context.Context, *ListMeetingsInput, ...func(*Options)) (*ListMeetingsOutput, error)
}

var _ ListMeetingsAPIClient = (*Client)(nil)

// ListMeetingsPaginatorOptions is the paginator options for ListMeetings
type ListMeetingsPaginatorOptions struct {
	// The maximum number of results to return in a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMeetingsPaginator is a paginator for ListMeetings
type ListMeetingsPaginator struct {
	options   ListMeetingsPaginatorOptions
	client    ListMeetingsAPIClient
	params    *ListMeetingsInput
	nextToken *string
	firstPage bool
}

// NewListMeetingsPaginator returns a new ListMeetingsPaginator
func NewListMeetingsPaginator(client ListMeetingsAPIClient, params *ListMeetingsInput, optFns ...func(*ListMeetingsPaginatorOptions)) *ListMeetingsPaginator {
	if params == nil {
		params = &ListMeetingsInput{}
	}

	options := ListMeetingsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMeetingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMeetingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMeetings page.
func (p *ListMeetingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMeetingsOutput, error) {
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

	result, err := p.client.ListMeetings(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMeetings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "ListMeetings",
	}
}
