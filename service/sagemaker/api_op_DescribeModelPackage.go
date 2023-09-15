// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a description of the specified model package, which is used to create
// SageMaker models or list them on Amazon Web Services Marketplace. To create
// models in SageMaker, buyers can subscribe to model packages listed on Amazon Web
// Services Marketplace.
func (c *Client) DescribeModelPackage(ctx context.Context, params *DescribeModelPackageInput, optFns ...func(*Options)) (*DescribeModelPackageOutput, error) {
	if params == nil {
		params = &DescribeModelPackageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeModelPackage", params, optFns, c.addOperationDescribeModelPackageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeModelPackageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeModelPackageInput struct {

	// The name or Amazon Resource Name (ARN) of the model package to describe. When
	// you specify a name, the name must have 1 to 63 characters. Valid characters are
	// a-z, A-Z, 0-9, and - (hyphen).
	//
	// This member is required.
	ModelPackageName *string

	noSmithyDocumentSerde
}

type DescribeModelPackageOutput struct {

	// A timestamp specifying when the model package was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the model package.
	//
	// This member is required.
	ModelPackageArn *string

	// The name of the model package being described.
	//
	// This member is required.
	ModelPackageName *string

	// The current status of the model package.
	//
	// This member is required.
	ModelPackageStatus types.ModelPackageStatus

	// Details about the current status of the model package.
	//
	// This member is required.
	ModelPackageStatusDetails *types.ModelPackageStatusDetails

	// An array of additional Inference Specification objects. Each additional
	// Inference Specification specifies artifacts based on this model package that can
	// be used on inference endpoints. Generally used with SageMaker Neo to store the
	// compiled artifacts.
	AdditionalInferenceSpecifications []types.AdditionalInferenceSpecificationDefinition

	// A description provided for the model approval.
	ApprovalDescription *string

	// Whether the model package is certified for listing on Amazon Web Services
	// Marketplace.
	CertifyForMarketplace bool

	// Information about the user who created or modified an experiment, trial, trial
	// component, lineage group, project, or model card.
	CreatedBy *types.UserContext

	// The metadata properties associated with the model package versions.
	CustomerMetadataProperties map[string]string

	// The machine learning domain of the model package you specified. Common machine
	// learning domains include computer vision and natural language processing.
	Domain *string

	// Represents the drift check baselines that can be used when the model monitor is
	// set using the model package. For more information, see the topic on Drift
	// Detection against Previous Baselines in SageMaker Pipelines (https://docs.aws.amazon.com/sagemaker/latest/dg/pipelines-quality-clarify-baseline-lifecycle.html#pipelines-quality-clarify-baseline-drift-detection)
	// in the Amazon SageMaker Developer Guide.
	DriftCheckBaselines *types.DriftCheckBaselines

	// Details about inference jobs that can be run with models based on this model
	// package.
	InferenceSpecification *types.InferenceSpecification

	// Information about the user who created or modified an experiment, trial, trial
	// component, lineage group, project, or model card.
	LastModifiedBy *types.UserContext

	// The last time that the model package was modified.
	LastModifiedTime *time.Time

	// Metadata properties of the tracking entity, trial, or trial component.
	MetadataProperties *types.MetadataProperties

	// The approval status of the model package.
	ModelApprovalStatus types.ModelApprovalStatus

	// Metrics for the model.
	ModelMetrics *types.ModelMetrics

	// A brief summary of the model package.
	ModelPackageDescription *string

	// If the model is a versioned model, the name of the model group that the
	// versioned model belongs to.
	ModelPackageGroupName *string

	// The version of the model package.
	ModelPackageVersion *int32

	// The Amazon Simple Storage Service (Amazon S3) path where the sample payload are
	// stored. This path points to a single gzip compressed tar archive (.tar.gz
	// suffix).
	SamplePayloadUrl *string

	// Indicates if you want to skip model validation.
	SkipModelValidation types.SkipModelValidation

	// Details about the algorithm that was used to create the model package.
	SourceAlgorithmSpecification *types.SourceAlgorithmSpecification

	// The machine learning task you specified that your model package accomplishes.
	// Common machine learning tasks include object detection and image classification.
	Task *string

	// Configurations for one or more transform jobs that SageMaker runs to test the
	// model package.
	ValidationSpecification *types.ModelPackageValidationSpecification

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeModelPackageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeModelPackage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeModelPackage{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = addDescribeModelPackageResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDescribeModelPackageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeModelPackage(options.Region), middleware.Before); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeModelPackage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeModelPackage",
	}
}

type opDescribeModelPackageResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeModelPackageResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeModelPackageResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "sagemaker"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "sagemaker"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("sagemaker")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addDescribeModelPackageResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeModelPackageResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
