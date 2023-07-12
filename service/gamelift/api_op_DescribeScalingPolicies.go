// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves all scaling policies applied to a fleet. To get a fleet's scaling
// policies, specify the fleet ID. You can filter this request by policy status,
// such as to retrieve only active scaling policies. Use the pagination parameters
// to retrieve results as a set of sequential pages. If successful, set of
// ScalingPolicy objects is returned for the fleet. A fleet may have all of its
// scaling policies suspended. This operation does not affect the status of the
// scaling policies, which remains ACTIVE.
func (c *Client) DescribeScalingPolicies(ctx context.Context, params *DescribeScalingPoliciesInput, optFns ...func(*Options)) (*DescribeScalingPoliciesOutput, error) {
	if params == nil {
		params = &DescribeScalingPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeScalingPolicies", params, optFns, c.addOperationDescribeScalingPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeScalingPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeScalingPoliciesInput struct {

	// A unique identifier for the fleet for which to retrieve scaling policies. You
	// can use either the fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit *int32

	// The fleet location. If you don't specify this value, the response contains the
	// scaling policies of every location in the fleet.
	Location *string

	// A token that indicates the start of the next sequential page of results. Use
	// the token that is returned with a previous call to this operation. To start at
	// the beginning of the result set, do not specify a value.
	NextToken *string

	// Scaling policy status to filter results on. A scaling policy is only in force
	// when in an ACTIVE status.
	//   - ACTIVE -- The scaling policy is currently in force.
	//   - UPDATEREQUESTED -- A request to update the scaling policy has been
	//   received.
	//   - UPDATING -- A change is being made to the scaling policy.
	//   - DELETEREQUESTED -- A request to delete the scaling policy has been
	//   received.
	//   - DELETING -- The scaling policy is being deleted.
	//   - DELETED -- The scaling policy has been deleted.
	//   - ERROR -- An error occurred in creating the policy. It should be removed and
	//   recreated.
	StatusFilter types.ScalingStatusType

	noSmithyDocumentSerde
}

type DescribeScalingPoliciesOutput struct {

	// A token that indicates where to resume retrieving results on the next call to
	// this operation. If no token is returned, these results represent the end of the
	// list.
	NextToken *string

	// A collection of objects containing the scaling policies matching the request.
	ScalingPolicies []types.ScalingPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeScalingPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeScalingPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeScalingPolicies{}, middleware.After)
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
	if err = addOpDescribeScalingPoliciesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeScalingPolicies(options.Region), middleware.Before); err != nil {
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

// DescribeScalingPoliciesAPIClient is a client that implements the
// DescribeScalingPolicies operation.
type DescribeScalingPoliciesAPIClient interface {
	DescribeScalingPolicies(context.Context, *DescribeScalingPoliciesInput, ...func(*Options)) (*DescribeScalingPoliciesOutput, error)
}

var _ DescribeScalingPoliciesAPIClient = (*Client)(nil)

// DescribeScalingPoliciesPaginatorOptions is the paginator options for
// DescribeScalingPolicies
type DescribeScalingPoliciesPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeScalingPoliciesPaginator is a paginator for DescribeScalingPolicies
type DescribeScalingPoliciesPaginator struct {
	options   DescribeScalingPoliciesPaginatorOptions
	client    DescribeScalingPoliciesAPIClient
	params    *DescribeScalingPoliciesInput
	nextToken *string
	firstPage bool
}

// NewDescribeScalingPoliciesPaginator returns a new
// DescribeScalingPoliciesPaginator
func NewDescribeScalingPoliciesPaginator(client DescribeScalingPoliciesAPIClient, params *DescribeScalingPoliciesInput, optFns ...func(*DescribeScalingPoliciesPaginatorOptions)) *DescribeScalingPoliciesPaginator {
	if params == nil {
		params = &DescribeScalingPoliciesInput{}
	}

	options := DescribeScalingPoliciesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeScalingPoliciesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeScalingPoliciesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeScalingPolicies page.
func (p *DescribeScalingPoliciesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeScalingPoliciesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.DescribeScalingPolicies(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeScalingPolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeScalingPolicies",
	}
}
