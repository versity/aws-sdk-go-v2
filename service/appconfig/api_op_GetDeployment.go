// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about a configuration deployment.
func (c *Client) GetDeployment(ctx context.Context, params *GetDeploymentInput, optFns ...func(*Options)) (*GetDeploymentOutput, error) {
	if params == nil {
		params = &GetDeploymentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDeployment", params, optFns, c.addOperationGetDeploymentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDeploymentInput struct {

	// The ID of the application that includes the deployment you want to get.
	//
	// This member is required.
	ApplicationId *string

	// The sequence number of the deployment.
	//
	// This member is required.
	DeploymentNumber *int32

	// The ID of the environment that includes the deployment you want to get.
	//
	// This member is required.
	EnvironmentId *string

	noSmithyDocumentSerde
}

type GetDeploymentOutput struct {

	// The ID of the application that was deployed.
	ApplicationId *string

	// A list of extensions that were processed as part of the deployment. The
	// extensions that were previously associated to the configuration profile,
	// environment, or the application when StartDeployment was called.
	AppliedExtensions []types.AppliedExtension

	// The time the deployment completed.
	CompletedAt *time.Time

	// Information about the source location of the configuration.
	ConfigurationLocationUri *string

	// The name of the configuration.
	ConfigurationName *string

	// The ID of the configuration profile that was deployed.
	ConfigurationProfileId *string

	// The configuration version that was deployed.
	ConfigurationVersion *string

	// Total amount of time the deployment lasted.
	DeploymentDurationInMinutes int32

	// The sequence number of the deployment.
	DeploymentNumber int32

	// The ID of the deployment strategy that was deployed.
	DeploymentStrategyId *string

	// The description of the deployment.
	Description *string

	// The ID of the environment that was deployed.
	EnvironmentId *string

	// A list containing all events related to a deployment. The most recent events
	// are displayed first.
	EventLog []types.DeploymentEvent

	// The amount of time that AppConfig monitored for alarms before considering the
	// deployment to be complete and no longer eligible for automatic rollback.
	FinalBakeTimeInMinutes int32

	// The percentage of targets to receive a deployed configuration during each
	// interval.
	GrowthFactor float32

	// The algorithm used to define how percentage grew over time.
	GrowthType types.GrowthType

	// The Amazon Resource Name of the Key Management Service key used to encrypt
	// configuration data. You can encrypt secrets stored in Secrets Manager, Amazon
	// Simple Storage Service (Amazon S3) objects encrypted with SSE-KMS, or secure
	// string parameters stored in Amazon Web Services Systems Manager Parameter Store.
	KmsKeyArn *string

	// The KMS key identifier (key ID, key alias, or key ARN). AppConfig uses this ID
	// to encrypt the configuration data using a customer managed key.
	KmsKeyIdentifier *string

	// The percentage of targets for which the deployment is available.
	PercentageComplete float32

	// The time the deployment started.
	StartedAt *time.Time

	// The state of the deployment.
	State types.DeploymentState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDeploymentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDeployment{}, middleware.After)
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
	if err = addOpGetDeploymentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDeployment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDeployment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "GetDeployment",
	}
}
