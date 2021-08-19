// Code generated by smithy-go-codegen DO NOT EDIT.

package appflow

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appflow/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of connector-profile details matching the provided
// connector-profile names and connector-types. Both input lists are optional, and
// you can use them to filter the result. If no names or connector-types are
// provided, returns all connector profiles in a paginated form. If there is no
// match, this operation returns an empty list.
func (c *Client) DescribeConnectorProfiles(ctx context.Context, params *DescribeConnectorProfilesInput, optFns ...func(*Options)) (*DescribeConnectorProfilesOutput, error) {
	if params == nil {
		params = &DescribeConnectorProfilesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConnectorProfiles", params, optFns, c.addOperationDescribeConnectorProfilesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConnectorProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConnectorProfilesInput struct {

	// The name of the connector profile. The name is unique for each ConnectorProfile
	// in the Amazon Web Services account.
	ConnectorProfileNames []string

	// The type of connector, such as Salesforce, Amplitude, and so on.
	ConnectorType types.ConnectorType

	// Specifies the maximum number of items that should be returned in the result set.
	// The default for maxResults is 20 (for all paginated API operations).
	MaxResults *int32

	// The pagination token for the next page of data.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeConnectorProfilesOutput struct {

	// Returns information about the connector profiles associated with the flow.
	ConnectorProfileDetails []types.ConnectorProfile

	// The pagination token for the next page of data. If nextToken=null, this means
	// that all records have been fetched.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeConnectorProfilesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeConnectorProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeConnectorProfiles{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConnectorProfiles(options.Region), middleware.Before); err != nil {
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

// DescribeConnectorProfilesAPIClient is a client that implements the
// DescribeConnectorProfiles operation.
type DescribeConnectorProfilesAPIClient interface {
	DescribeConnectorProfiles(context.Context, *DescribeConnectorProfilesInput, ...func(*Options)) (*DescribeConnectorProfilesOutput, error)
}

var _ DescribeConnectorProfilesAPIClient = (*Client)(nil)

// DescribeConnectorProfilesPaginatorOptions is the paginator options for
// DescribeConnectorProfiles
type DescribeConnectorProfilesPaginatorOptions struct {
	// Specifies the maximum number of items that should be returned in the result set.
	// The default for maxResults is 20 (for all paginated API operations).
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeConnectorProfilesPaginator is a paginator for DescribeConnectorProfiles
type DescribeConnectorProfilesPaginator struct {
	options   DescribeConnectorProfilesPaginatorOptions
	client    DescribeConnectorProfilesAPIClient
	params    *DescribeConnectorProfilesInput
	nextToken *string
	firstPage bool
}

// NewDescribeConnectorProfilesPaginator returns a new
// DescribeConnectorProfilesPaginator
func NewDescribeConnectorProfilesPaginator(client DescribeConnectorProfilesAPIClient, params *DescribeConnectorProfilesInput, optFns ...func(*DescribeConnectorProfilesPaginatorOptions)) *DescribeConnectorProfilesPaginator {
	if params == nil {
		params = &DescribeConnectorProfilesInput{}
	}

	options := DescribeConnectorProfilesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeConnectorProfilesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeConnectorProfilesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeConnectorProfiles page.
func (p *DescribeConnectorProfilesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeConnectorProfilesOutput, error) {
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

	result, err := p.client.DescribeConnectorProfiles(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeConnectorProfiles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appflow",
		OperationName: "DescribeConnectorProfiles",
	}
}
