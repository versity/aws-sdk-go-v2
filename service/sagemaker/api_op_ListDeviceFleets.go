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

// Returns a list of devices in the fleet.
func (c *Client) ListDeviceFleets(ctx context.Context, params *ListDeviceFleetsInput, optFns ...func(*Options)) (*ListDeviceFleetsOutput, error) {
	if params == nil {
		params = &ListDeviceFleetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDeviceFleets", params, optFns, c.addOperationListDeviceFleetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDeviceFleetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDeviceFleetsInput struct {

	// Filter fleets where packaging job was created after specified time.
	CreationTimeAfter *time.Time

	// Filter fleets where the edge packaging job was created before specified time.
	CreationTimeBefore *time.Time

	// Select fleets where the job was updated after X
	LastModifiedTimeAfter *time.Time

	// Select fleets where the job was updated before X
	LastModifiedTimeBefore *time.Time

	// The maximum number of results to select.
	MaxResults *int32

	// Filter for fleets containing this name in their fleet device name.
	NameContains *string

	// The response from the last list when returning a list large enough to need
	// tokening.
	NextToken *string

	// The column to sort by.
	SortBy types.ListDeviceFleetsSortBy

	// What direction to sort in.
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListDeviceFleetsOutput struct {

	// Summary of the device fleet.
	//
	// This member is required.
	DeviceFleetSummaries []types.DeviceFleetSummary

	// The response from the last list when returning a list large enough to need
	// tokening.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDeviceFleetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDeviceFleets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDeviceFleets{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDeviceFleets(options.Region), middleware.Before); err != nil {
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

// ListDeviceFleetsAPIClient is a client that implements the ListDeviceFleets
// operation.
type ListDeviceFleetsAPIClient interface {
	ListDeviceFleets(context.Context, *ListDeviceFleetsInput, ...func(*Options)) (*ListDeviceFleetsOutput, error)
}

var _ ListDeviceFleetsAPIClient = (*Client)(nil)

// ListDeviceFleetsPaginatorOptions is the paginator options for ListDeviceFleets
type ListDeviceFleetsPaginatorOptions struct {
	// The maximum number of results to select.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDeviceFleetsPaginator is a paginator for ListDeviceFleets
type ListDeviceFleetsPaginator struct {
	options   ListDeviceFleetsPaginatorOptions
	client    ListDeviceFleetsAPIClient
	params    *ListDeviceFleetsInput
	nextToken *string
	firstPage bool
}

// NewListDeviceFleetsPaginator returns a new ListDeviceFleetsPaginator
func NewListDeviceFleetsPaginator(client ListDeviceFleetsAPIClient, params *ListDeviceFleetsInput, optFns ...func(*ListDeviceFleetsPaginatorOptions)) *ListDeviceFleetsPaginator {
	if params == nil {
		params = &ListDeviceFleetsInput{}
	}

	options := ListDeviceFleetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDeviceFleetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDeviceFleetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDeviceFleets page.
func (p *ListDeviceFleetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDeviceFleetsOutput, error) {
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

	result, err := p.client.ListDeviceFleets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDeviceFleets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListDeviceFleets",
	}
}
