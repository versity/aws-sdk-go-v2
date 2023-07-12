// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconvert

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieve a JSON array of up to twenty of your presets. This will return the
// presets themselves, not just a list of them. To retrieve the next twenty
// presets, use the nextToken string returned with the array.
func (c *Client) ListPresets(ctx context.Context, params *ListPresetsInput, optFns ...func(*Options)) (*ListPresetsOutput, error) {
	if params == nil {
		params = &ListPresetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPresets", params, optFns, c.addOperationListPresetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPresetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPresetsInput struct {

	// Optionally, specify a preset category to limit responses to only presets from
	// that category.
	Category *string

	// Optional. When you request a list of presets, you can choose to list them
	// alphabetically by NAME or chronologically by CREATION_DATE. If you don't
	// specify, the service will list them by name.
	ListBy types.PresetListBy

	// Optional. Number of presets, up to twenty, that will be returned at one time
	MaxResults int32

	// Use this string, provided with the response to a previous request, to request
	// the next batch of presets.
	NextToken *string

	// Optional. When you request lists of resources, you can specify whether they are
	// sorted in ASCENDING or DESCENDING order. Default varies by resource.
	Order types.Order

	noSmithyDocumentSerde
}

type ListPresetsOutput struct {

	// Use this string to request the next batch of presets.
	NextToken *string

	// List of presets
	Presets []types.Preset

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPresetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPresets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPresets{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPresets(options.Region), middleware.Before); err != nil {
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

// ListPresetsAPIClient is a client that implements the ListPresets operation.
type ListPresetsAPIClient interface {
	ListPresets(context.Context, *ListPresetsInput, ...func(*Options)) (*ListPresetsOutput, error)
}

var _ ListPresetsAPIClient = (*Client)(nil)

// ListPresetsPaginatorOptions is the paginator options for ListPresets
type ListPresetsPaginatorOptions struct {
	// Optional. Number of presets, up to twenty, that will be returned at one time
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPresetsPaginator is a paginator for ListPresets
type ListPresetsPaginator struct {
	options   ListPresetsPaginatorOptions
	client    ListPresetsAPIClient
	params    *ListPresetsInput
	nextToken *string
	firstPage bool
}

// NewListPresetsPaginator returns a new ListPresetsPaginator
func NewListPresetsPaginator(client ListPresetsAPIClient, params *ListPresetsInput, optFns ...func(*ListPresetsPaginatorOptions)) *ListPresetsPaginator {
	if params == nil {
		params = &ListPresetsInput{}
	}

	options := ListPresetsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPresetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPresetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPresets page.
func (p *ListPresetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPresetsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListPresets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPresets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconvert",
		OperationName: "ListPresets",
	}
}
