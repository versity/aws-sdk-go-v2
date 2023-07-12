// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Allows you to either upgrade your domain or perform an Upgrade eligibility
// check to a compatible Elasticsearch version.
func (c *Client) UpgradeElasticsearchDomain(ctx context.Context, params *UpgradeElasticsearchDomainInput, optFns ...func(*Options)) (*UpgradeElasticsearchDomainOutput, error) {
	if params == nil {
		params = &UpgradeElasticsearchDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpgradeElasticsearchDomain", params, optFns, c.addOperationUpgradeElasticsearchDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpgradeElasticsearchDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for request parameters to UpgradeElasticsearchDomain operation.
type UpgradeElasticsearchDomainInput struct {

	// The name of an Elasticsearch domain. Domain names are unique across the domains
	// owned by an account within an AWS region. Domain names start with a letter or
	// number and can contain the following characters: a-z (lowercase), 0-9, and -
	// (hyphen).
	//
	// This member is required.
	DomainName *string

	// The version of Elasticsearch that you intend to upgrade the domain to.
	//
	// This member is required.
	TargetVersion *string

	// This flag, when set to True, indicates that an Upgrade Eligibility Check needs
	// to be performed. This will not actually perform the Upgrade.
	PerformCheckOnly *bool

	noSmithyDocumentSerde
}

// Container for response returned by UpgradeElasticsearchDomain operation.
type UpgradeElasticsearchDomainOutput struct {

	// Specifies change details of the domain configuration change.
	ChangeProgressDetails *types.ChangeProgressDetails

	// The name of an Elasticsearch domain. Domain names are unique across the domains
	// owned by an account within an AWS region. Domain names start with a letter or
	// number and can contain the following characters: a-z (lowercase), 0-9, and -
	// (hyphen).
	DomainName *string

	// This flag, when set to True, indicates that an Upgrade Eligibility Check needs
	// to be performed. This will not actually perform the Upgrade.
	PerformCheckOnly *bool

	// The version of Elasticsearch that you intend to upgrade the domain to.
	TargetVersion *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpgradeElasticsearchDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpgradeElasticsearchDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpgradeElasticsearchDomain{}, middleware.After)
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
	if err = addOpUpgradeElasticsearchDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpgradeElasticsearchDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpgradeElasticsearchDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "UpgradeElasticsearchDomain",
	}
}
