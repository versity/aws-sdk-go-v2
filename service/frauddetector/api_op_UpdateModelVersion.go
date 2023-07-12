// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a model version. Updating a model version retrains an existing model
// version using updated training data and produces a new minor version of the
// model. You can update the training data set location and data access role
// attributes using this action. This action creates and trains a new minor version
// of the model, for example version 1.01, 1.02, 1.03.
func (c *Client) UpdateModelVersion(ctx context.Context, params *UpdateModelVersionInput, optFns ...func(*Options)) (*UpdateModelVersionOutput, error) {
	if params == nil {
		params = &UpdateModelVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateModelVersion", params, optFns, c.addOperationUpdateModelVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateModelVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateModelVersionInput struct {

	// The major version number.
	//
	// This member is required.
	MajorVersionNumber *string

	// The model ID.
	//
	// This member is required.
	ModelId *string

	// The model type.
	//
	// This member is required.
	ModelType types.ModelTypeEnum

	// The details of the external events data used for training the model version.
	// Required if trainingDataSource is EXTERNAL_EVENTS .
	ExternalEventsDetail *types.ExternalEventsDetail

	// The details of the ingested event used for training the model version. Required
	// if your trainingDataSource is INGESTED_EVENTS .
	IngestedEventsDetail *types.IngestedEventsDetail

	// A collection of key and value pairs.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type UpdateModelVersionOutput struct {

	// The model ID.
	ModelId *string

	// The model type.
	ModelType types.ModelTypeEnum

	// The model version number of the model version updated.
	ModelVersionNumber *string

	// The status of the updated model version.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateModelVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateModelVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateModelVersion{}, middleware.After)
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
	if err = addOpUpdateModelVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateModelVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateModelVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "UpdateModelVersion",
	}
}
