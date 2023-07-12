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

// Returns a list of the SageMaker notebook instances in the requester's account
// in an Amazon Web Services Region.
func (c *Client) ListNotebookInstances(ctx context.Context, params *ListNotebookInstancesInput, optFns ...func(*Options)) (*ListNotebookInstancesOutput, error) {
	if params == nil {
		params = &ListNotebookInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListNotebookInstances", params, optFns, c.addOperationListNotebookInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListNotebookInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListNotebookInstancesInput struct {

	// A filter that returns only notebook instances with associated with the
	// specified git repository.
	AdditionalCodeRepositoryEquals *string

	// A filter that returns only notebook instances that were created after the
	// specified time (timestamp).
	CreationTimeAfter *time.Time

	// A filter that returns only notebook instances that were created before the
	// specified time (timestamp).
	CreationTimeBefore *time.Time

	// A string in the name or URL of a Git repository associated with this notebook
	// instance. This filter returns only notebook instances associated with a git
	// repository with a name that contains the specified string.
	DefaultCodeRepositoryContains *string

	// A filter that returns only notebook instances that were modified after the
	// specified time (timestamp).
	LastModifiedTimeAfter *time.Time

	// A filter that returns only notebook instances that were modified before the
	// specified time (timestamp).
	LastModifiedTimeBefore *time.Time

	// The maximum number of notebook instances to return.
	MaxResults *int32

	// A string in the notebook instances' name. This filter returns only notebook
	// instances whose name contains the specified string.
	NameContains *string

	// If the previous call to the ListNotebookInstances is truncated, the response
	// includes a NextToken . You can use this token in your subsequent
	// ListNotebookInstances request to fetch the next set of notebook instances. You
	// might specify a filter or a sort order in your request. When response is
	// truncated, you must use the same values for the filer and sort order in the next
	// request.
	NextToken *string

	// A string in the name of a notebook instances lifecycle configuration associated
	// with this notebook instance. This filter returns only notebook instances
	// associated with a lifecycle configuration with a name that contains the
	// specified string.
	NotebookInstanceLifecycleConfigNameContains *string

	// The field to sort results by. The default is Name .
	SortBy types.NotebookInstanceSortKey

	// The sort order for results.
	SortOrder types.NotebookInstanceSortOrder

	// A filter that returns only notebook instances with the specified status.
	StatusEquals types.NotebookInstanceStatus

	noSmithyDocumentSerde
}

type ListNotebookInstancesOutput struct {

	// If the response to the previous ListNotebookInstances request was truncated,
	// SageMaker returns this token. To retrieve the next set of notebook instances,
	// use the token in the next request.
	NextToken *string

	// An array of NotebookInstanceSummary objects, one for each notebook instance.
	NotebookInstances []types.NotebookInstanceSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListNotebookInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListNotebookInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListNotebookInstances{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListNotebookInstances(options.Region), middleware.Before); err != nil {
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

// ListNotebookInstancesAPIClient is a client that implements the
// ListNotebookInstances operation.
type ListNotebookInstancesAPIClient interface {
	ListNotebookInstances(context.Context, *ListNotebookInstancesInput, ...func(*Options)) (*ListNotebookInstancesOutput, error)
}

var _ ListNotebookInstancesAPIClient = (*Client)(nil)

// ListNotebookInstancesPaginatorOptions is the paginator options for
// ListNotebookInstances
type ListNotebookInstancesPaginatorOptions struct {
	// The maximum number of notebook instances to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListNotebookInstancesPaginator is a paginator for ListNotebookInstances
type ListNotebookInstancesPaginator struct {
	options   ListNotebookInstancesPaginatorOptions
	client    ListNotebookInstancesAPIClient
	params    *ListNotebookInstancesInput
	nextToken *string
	firstPage bool
}

// NewListNotebookInstancesPaginator returns a new ListNotebookInstancesPaginator
func NewListNotebookInstancesPaginator(client ListNotebookInstancesAPIClient, params *ListNotebookInstancesInput, optFns ...func(*ListNotebookInstancesPaginatorOptions)) *ListNotebookInstancesPaginator {
	if params == nil {
		params = &ListNotebookInstancesInput{}
	}

	options := ListNotebookInstancesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListNotebookInstancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListNotebookInstancesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListNotebookInstances page.
func (p *ListNotebookInstancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListNotebookInstancesOutput, error) {
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

	result, err := p.client.ListNotebookInstances(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListNotebookInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListNotebookInstances",
	}
}
