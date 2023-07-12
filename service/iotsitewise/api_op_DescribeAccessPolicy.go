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

// Describes an access policy, which specifies an identity's access to an IoT
// SiteWise Monitor portal or project.
func (c *Client) DescribeAccessPolicy(ctx context.Context, params *DescribeAccessPolicyInput, optFns ...func(*Options)) (*DescribeAccessPolicyOutput, error) {
	if params == nil {
		params = &DescribeAccessPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAccessPolicy", params, optFns, c.addOperationDescribeAccessPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAccessPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAccessPolicyInput struct {

	// The ID of the access policy.
	//
	// This member is required.
	AccessPolicyId *string

	noSmithyDocumentSerde
}

type DescribeAccessPolicyOutput struct {

	// The ARN (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the access policy, which has the following format.
	// arn:${Partition}:iotsitewise:${Region}:${Account}:access-policy/${AccessPolicyId}
	//
	// This member is required.
	AccessPolicyArn *string

	// The date the access policy was created, in Unix epoch time.
	//
	// This member is required.
	AccessPolicyCreationDate *time.Time

	// The ID of the access policy.
	//
	// This member is required.
	AccessPolicyId *string

	// The identity (IAM Identity Center user, IAM Identity Center group, or IAM user)
	// to which this access policy applies.
	//
	// This member is required.
	AccessPolicyIdentity *types.Identity

	// The date the access policy was last updated, in Unix epoch time.
	//
	// This member is required.
	AccessPolicyLastUpdateDate *time.Time

	// The access policy permission. Note that a project ADMINISTRATOR is also known
	// as a project owner.
	//
	// This member is required.
	AccessPolicyPermission types.Permission

	// The IoT SiteWise Monitor resource (portal or project) to which this access
	// policy provides access.
	//
	// This member is required.
	AccessPolicyResource *types.Resource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAccessPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAccessPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAccessPolicy{}, middleware.After)
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
	if err = addEndpointPrefix_opDescribeAccessPolicyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeAccessPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAccessPolicy(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opDescribeAccessPolicyMiddleware struct {
}

func (*endpointPrefix_opDescribeAccessPolicyMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDescribeAccessPolicyMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "monitor." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opDescribeAccessPolicyMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opDescribeAccessPolicyMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAccessPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "DescribeAccessPolicy",
	}
}
