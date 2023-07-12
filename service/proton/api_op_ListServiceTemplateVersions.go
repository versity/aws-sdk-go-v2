// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List major or minor versions of a service template with detail data.
func (c *Client) ListServiceTemplateVersions(ctx context.Context, params *ListServiceTemplateVersionsInput, optFns ...func(*Options)) (*ListServiceTemplateVersionsOutput, error) {
	if params == nil {
		params = &ListServiceTemplateVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListServiceTemplateVersions", params, optFns, c.addOperationListServiceTemplateVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListServiceTemplateVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListServiceTemplateVersionsInput struct {

	// The name of the service template.
	//
	// This member is required.
	TemplateName *string

	// To view a list of minor of versions under a major version of a service
	// template, include major Version . To view a list of major versions of a service
	// template, exclude major Version .
	MajorVersion *string

	// The maximum number of major or minor versions of a service template to list.
	MaxResults *int32

	// A token that indicates the location of the next major or minor version in the
	// array of major or minor versions of a service template, after the list of major
	// or minor versions that was previously requested.
	NextToken *string

	noSmithyDocumentSerde
}

type ListServiceTemplateVersionsOutput struct {

	// An array of major or minor versions of a service template with detail data.
	//
	// This member is required.
	TemplateVersions []types.ServiceTemplateVersionSummary

	// A token that indicates the location of the next major or minor version in the
	// array of major or minor versions of a service template, after the current
	// requested list of service major or minor versions.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListServiceTemplateVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListServiceTemplateVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListServiceTemplateVersions{}, middleware.After)
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
	if err = addOpListServiceTemplateVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListServiceTemplateVersions(options.Region), middleware.Before); err != nil {
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

// ListServiceTemplateVersionsAPIClient is a client that implements the
// ListServiceTemplateVersions operation.
type ListServiceTemplateVersionsAPIClient interface {
	ListServiceTemplateVersions(context.Context, *ListServiceTemplateVersionsInput, ...func(*Options)) (*ListServiceTemplateVersionsOutput, error)
}

var _ ListServiceTemplateVersionsAPIClient = (*Client)(nil)

// ListServiceTemplateVersionsPaginatorOptions is the paginator options for
// ListServiceTemplateVersions
type ListServiceTemplateVersionsPaginatorOptions struct {
	// The maximum number of major or minor versions of a service template to list.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListServiceTemplateVersionsPaginator is a paginator for
// ListServiceTemplateVersions
type ListServiceTemplateVersionsPaginator struct {
	options   ListServiceTemplateVersionsPaginatorOptions
	client    ListServiceTemplateVersionsAPIClient
	params    *ListServiceTemplateVersionsInput
	nextToken *string
	firstPage bool
}

// NewListServiceTemplateVersionsPaginator returns a new
// ListServiceTemplateVersionsPaginator
func NewListServiceTemplateVersionsPaginator(client ListServiceTemplateVersionsAPIClient, params *ListServiceTemplateVersionsInput, optFns ...func(*ListServiceTemplateVersionsPaginatorOptions)) *ListServiceTemplateVersionsPaginator {
	if params == nil {
		params = &ListServiceTemplateVersionsInput{}
	}

	options := ListServiceTemplateVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListServiceTemplateVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListServiceTemplateVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListServiceTemplateVersions page.
func (p *ListServiceTemplateVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListServiceTemplateVersionsOutput, error) {
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

	result, err := p.client.ListServiceTemplateVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListServiceTemplateVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "proton",
		OperationName: "ListServiceTemplateVersions",
	}
}
