// Code generated by smithy-go-codegen DO NOT EDIT.

package synthetics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/synthetics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of Synthetics canary runtime versions. For more information, see
// Canary Runtime Versions (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library.html)
// .
func (c *Client) DescribeRuntimeVersions(ctx context.Context, params *DescribeRuntimeVersionsInput, optFns ...func(*Options)) (*DescribeRuntimeVersionsOutput, error) {
	if params == nil {
		params = &DescribeRuntimeVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRuntimeVersions", params, optFns, c.addOperationDescribeRuntimeVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRuntimeVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRuntimeVersionsInput struct {

	// Specify this parameter to limit how many runs are returned each time you use
	// the DescribeRuntimeVersions operation. If you omit this parameter, the default
	// of 100 is used.
	MaxResults *int32

	// A token that indicates that there is more data available. You can use this
	// token in a subsequent DescribeRuntimeVersions operation to retrieve the next
	// set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeRuntimeVersionsOutput struct {

	// A token that indicates that there is more data available. You can use this
	// token in a subsequent DescribeRuntimeVersions operation to retrieve the next
	// set of results.
	NextToken *string

	// An array of objects that display the details about each Synthetics canary
	// runtime version.
	RuntimeVersions []types.RuntimeVersion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRuntimeVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeRuntimeVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeRuntimeVersions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRuntimeVersions(options.Region), middleware.Before); err != nil {
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

// DescribeRuntimeVersionsAPIClient is a client that implements the
// DescribeRuntimeVersions operation.
type DescribeRuntimeVersionsAPIClient interface {
	DescribeRuntimeVersions(context.Context, *DescribeRuntimeVersionsInput, ...func(*Options)) (*DescribeRuntimeVersionsOutput, error)
}

var _ DescribeRuntimeVersionsAPIClient = (*Client)(nil)

// DescribeRuntimeVersionsPaginatorOptions is the paginator options for
// DescribeRuntimeVersions
type DescribeRuntimeVersionsPaginatorOptions struct {
	// Specify this parameter to limit how many runs are returned each time you use
	// the DescribeRuntimeVersions operation. If you omit this parameter, the default
	// of 100 is used.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeRuntimeVersionsPaginator is a paginator for DescribeRuntimeVersions
type DescribeRuntimeVersionsPaginator struct {
	options   DescribeRuntimeVersionsPaginatorOptions
	client    DescribeRuntimeVersionsAPIClient
	params    *DescribeRuntimeVersionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeRuntimeVersionsPaginator returns a new
// DescribeRuntimeVersionsPaginator
func NewDescribeRuntimeVersionsPaginator(client DescribeRuntimeVersionsAPIClient, params *DescribeRuntimeVersionsInput, optFns ...func(*DescribeRuntimeVersionsPaginatorOptions)) *DescribeRuntimeVersionsPaginator {
	if params == nil {
		params = &DescribeRuntimeVersionsInput{}
	}

	options := DescribeRuntimeVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeRuntimeVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeRuntimeVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeRuntimeVersions page.
func (p *DescribeRuntimeVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeRuntimeVersionsOutput, error) {
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

	result, err := p.client.DescribeRuntimeVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeRuntimeVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "synthetics",
		OperationName: "DescribeRuntimeVersions",
	}
}
