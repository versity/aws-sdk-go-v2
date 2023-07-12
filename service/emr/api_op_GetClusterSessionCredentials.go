// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides temporary, HTTP basic credentials that are associated with a given
// runtime IAM role and used by a cluster with fine-grained access control
// activated. You can use these credentials to connect to cluster endpoints that
// support username and password authentication.
func (c *Client) GetClusterSessionCredentials(ctx context.Context, params *GetClusterSessionCredentialsInput, optFns ...func(*Options)) (*GetClusterSessionCredentialsOutput, error) {
	if params == nil {
		params = &GetClusterSessionCredentialsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetClusterSessionCredentials", params, optFns, c.addOperationGetClusterSessionCredentialsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetClusterSessionCredentialsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetClusterSessionCredentialsInput struct {

	// The unique identifier of the cluster.
	//
	// This member is required.
	ClusterId *string

	// The Amazon Resource Name (ARN) of the runtime role for interactive workload
	// submission on the cluster. The runtime role can be a cross-account IAM role. The
	// runtime role ARN is a combination of account ID, role name, and role type using
	// the following format: arn:partition:service:region:account:resource .
	//
	// This member is required.
	ExecutionRoleArn *string

	noSmithyDocumentSerde
}

type GetClusterSessionCredentialsOutput struct {

	// The credentials that you can use to connect to cluster endpoints that support
	// username and password authentication.
	Credentials types.Credentials

	// The time when the credentials that are returned by the
	// GetClusterSessionCredentials API expire.
	ExpiresAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetClusterSessionCredentialsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetClusterSessionCredentials{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetClusterSessionCredentials{}, middleware.After)
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
	if err = addOpGetClusterSessionCredentialsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetClusterSessionCredentials(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetClusterSessionCredentials(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "GetClusterSessionCredentials",
	}
}
