// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets aggregated values for an asset property. For more information, see
// Querying aggregates (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/query-industrial-data.html#aggregates)
// in the IoT SiteWise User Guide. To identify an asset property, you must specify
// one of the following:
//   - The assetId and propertyId of an asset property.
//   - A propertyAlias , which is a data stream alias (for example,
//     /company/windfarm/3/turbine/7/temperature ). To define an asset property's
//     alias, see UpdateAssetProperty (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_UpdateAssetProperty.html)
//     .
func (c *Client) GetAssetPropertyAggregates(ctx context.Context, params *GetAssetPropertyAggregatesInput, optFns ...func(*Options)) (*GetAssetPropertyAggregatesOutput, error) {
	if params == nil {
		params = &GetAssetPropertyAggregatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAssetPropertyAggregates", params, optFns, c.addOperationGetAssetPropertyAggregatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAssetPropertyAggregatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAssetPropertyAggregatesInput struct {

	// The data aggregating function.
	//
	// This member is required.
	AggregateTypes []types.AggregateType

	// The inclusive end of the range from which to query historical data, expressed
	// in seconds in Unix epoch time.
	//
	// This member is required.
	EndDate *time.Time

	// The time interval over which to aggregate data.
	//
	// This member is required.
	Resolution *string

	// The exclusive start of the range from which to query historical data, expressed
	// in seconds in Unix epoch time.
	//
	// This member is required.
	StartDate *time.Time

	// The ID of the asset.
	AssetId *string

	// The maximum number of results to return for each paginated request. A result
	// set is returned in the two cases, whichever occurs first.
	//   - The size of the result set is equal to 1 MB.
	//   - The number of data points in the result set is equal to the value of
	//   maxResults . The maximum value of maxResults is 250.
	MaxResults *int32

	// The token to be used for the next set of paginated results.
	NextToken *string

	// The alias that identifies the property, such as an OPC-UA server data stream
	// path (for example, /company/windfarm/3/turbine/7/temperature ). For more
	// information, see Mapping industrial data streams to asset properties (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/connect-data-streams.html)
	// in the IoT SiteWise User Guide.
	PropertyAlias *string

	// The ID of the asset property.
	PropertyId *string

	// The quality by which to filter asset data.
	Qualities []types.Quality

	// The chronological sorting order of the requested information. Default: ASCENDING
	TimeOrdering types.TimeOrdering

	noSmithyDocumentSerde
}

type GetAssetPropertyAggregatesOutput struct {

	// The requested aggregated values.
	//
	// This member is required.
	AggregatedValues []types.AggregatedValue

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAssetPropertyAggregatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAssetPropertyAggregates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAssetPropertyAggregates{}, middleware.After)
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
	if err = addEndpointPrefix_opGetAssetPropertyAggregatesMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetAssetPropertyAggregatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAssetPropertyAggregates(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetAssetPropertyAggregatesMiddleware struct {
}

func (*endpointPrefix_opGetAssetPropertyAggregatesMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetAssetPropertyAggregatesMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "data." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opGetAssetPropertyAggregatesMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opGetAssetPropertyAggregatesMiddleware{}, `OperationSerializer`, middleware.After)
}

// GetAssetPropertyAggregatesAPIClient is a client that implements the
// GetAssetPropertyAggregates operation.
type GetAssetPropertyAggregatesAPIClient interface {
	GetAssetPropertyAggregates(context.Context, *GetAssetPropertyAggregatesInput, ...func(*Options)) (*GetAssetPropertyAggregatesOutput, error)
}

var _ GetAssetPropertyAggregatesAPIClient = (*Client)(nil)

// GetAssetPropertyAggregatesPaginatorOptions is the paginator options for
// GetAssetPropertyAggregates
type GetAssetPropertyAggregatesPaginatorOptions struct {
	// The maximum number of results to return for each paginated request. A result
	// set is returned in the two cases, whichever occurs first.
	//   - The size of the result set is equal to 1 MB.
	//   - The number of data points in the result set is equal to the value of
	//   maxResults . The maximum value of maxResults is 250.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetAssetPropertyAggregatesPaginator is a paginator for
// GetAssetPropertyAggregates
type GetAssetPropertyAggregatesPaginator struct {
	options   GetAssetPropertyAggregatesPaginatorOptions
	client    GetAssetPropertyAggregatesAPIClient
	params    *GetAssetPropertyAggregatesInput
	nextToken *string
	firstPage bool
}

// NewGetAssetPropertyAggregatesPaginator returns a new
// GetAssetPropertyAggregatesPaginator
func NewGetAssetPropertyAggregatesPaginator(client GetAssetPropertyAggregatesAPIClient, params *GetAssetPropertyAggregatesInput, optFns ...func(*GetAssetPropertyAggregatesPaginatorOptions)) *GetAssetPropertyAggregatesPaginator {
	if params == nil {
		params = &GetAssetPropertyAggregatesInput{}
	}

	options := GetAssetPropertyAggregatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetAssetPropertyAggregatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetAssetPropertyAggregatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetAssetPropertyAggregates page.
func (p *GetAssetPropertyAggregatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetAssetPropertyAggregatesOutput, error) {
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

	result, err := p.client.GetAssetPropertyAggregates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetAssetPropertyAggregates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "GetAssetPropertyAggregates",
	}
}
