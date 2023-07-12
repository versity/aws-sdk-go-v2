// Code generated by smithy-go-codegen DO NOT EDIT.

package mq

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a new configuration for the specified configuration name. Amazon MQ
// uses the default configuration (the engine type and version).
func (c *Client) CreateConfiguration(ctx context.Context, params *CreateConfigurationInput, optFns ...func(*Options)) (*CreateConfigurationOutput, error) {
	if params == nil {
		params = &CreateConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateConfiguration", params, optFns, c.addOperationCreateConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates a new configuration for the specified configuration name. Amazon MQ
// uses the default configuration (the engine type and version).
type CreateConfigurationInput struct {

	// Required. The type of broker engine. Currently, Amazon MQ supports ACTIVEMQ and
	// RABBITMQ.
	//
	// This member is required.
	EngineType types.EngineType

	// Required. The broker engine's version. For a list of supported engine versions,
	// see Supported engines (https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/broker-engine.html)
	// .
	//
	// This member is required.
	EngineVersion *string

	// Required. The name of the configuration. This value can contain only
	// alphanumeric characters, dashes, periods, underscores, and tildes (- . _ ~).
	// This value must be 1-150 characters long.
	//
	// This member is required.
	Name *string

	// Optional. The authentication strategy associated with the configuration. The
	// default is SIMPLE.
	AuthenticationStrategy types.AuthenticationStrategy

	// Create tags when creating the configuration.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateConfigurationOutput struct {

	// Required. The Amazon Resource Name (ARN) of the configuration.
	Arn *string

	// Optional. The authentication strategy associated with the configuration. The
	// default is SIMPLE.
	AuthenticationStrategy types.AuthenticationStrategy

	// Required. The date and time of the configuration.
	Created *time.Time

	// Required. The unique ID that Amazon MQ generates for the configuration.
	Id *string

	// The latest revision of the configuration.
	LatestRevision *types.ConfigurationRevision

	// Required. The name of the configuration. This value can contain only
	// alphanumeric characters, dashes, periods, underscores, and tildes (- . _ ~).
	// This value must be 1-150 characters long.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateConfiguration{}, middleware.After)
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
	if err = addOpCreateConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mq",
		OperationName: "CreateConfiguration",
	}
}
