// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakergeospatial

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemakergeospatial/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Use this operation to create an Earth observation job.
func (c *Client) StartEarthObservationJob(ctx context.Context, params *StartEarthObservationJobInput, optFns ...func(*Options)) (*StartEarthObservationJobOutput, error) {
	if params == nil {
		params = &StartEarthObservationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartEarthObservationJob", params, optFns, c.addOperationStartEarthObservationJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartEarthObservationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartEarthObservationJobInput struct {

	// The Amazon Resource Name (ARN) of the IAM role that you specified for the job.
	//
	// This member is required.
	ExecutionRoleArn *string

	// Input configuration information for the Earth Observation job.
	//
	// This member is required.
	InputConfig *types.InputConfigInput

	// An object containing information about the job configuration.
	//
	// This member is required.
	JobConfig types.JobConfigInput

	// The name of the Earth Observation job.
	//
	// This member is required.
	Name *string

	// A unique token that guarantees that the call to this API is idempotent.
	ClientToken *string

	// The Key Management Service key ID for server-side encryption.
	KmsKeyId *string

	// Each tag consists of a key and a value.
	Tags map[string]string

	noSmithyDocumentSerde
}

type StartEarthObservationJobOutput struct {

	// The Amazon Resource Name (ARN) of the Earth Observation job.
	//
	// This member is required.
	Arn *string

	// The creation time.
	//
	// This member is required.
	CreationTime *time.Time

	// The duration of the session, in seconds.
	//
	// This member is required.
	DurationInSeconds *int32

	// The Amazon Resource Name (ARN) of the IAM role that you specified for the job.
	//
	// This member is required.
	ExecutionRoleArn *string

	// An object containing information about the job configuration.
	//
	// This member is required.
	JobConfig types.JobConfigInput

	// The name of the Earth Observation job.
	//
	// This member is required.
	Name *string

	// The status of the Earth Observation job.
	//
	// This member is required.
	Status types.EarthObservationJobStatus

	// Input configuration information for the Earth Observation job.
	InputConfig *types.InputConfigOutput

	// The Key Management Service key ID for server-side encryption.
	KmsKeyId *string

	// Each tag consists of a key and a value.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartEarthObservationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartEarthObservationJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartEarthObservationJob{}, middleware.After)
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
	if err = addIdempotencyToken_opStartEarthObservationJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartEarthObservationJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartEarthObservationJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartEarthObservationJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartEarthObservationJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartEarthObservationJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartEarthObservationJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartEarthObservationJobInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartEarthObservationJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartEarthObservationJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartEarthObservationJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker-geospatial",
		OperationName: "StartEarthObservationJob",
	}
}
