// Code generated by smithy-go-codegen DO NOT EDIT.

package ecrpublic

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Checks the availability of one or more image layers that are within a
// repository in a public registry. When an image is pushed to a repository, each
// image layer is checked to verify if it has been uploaded before. If it has been
// uploaded, then the image layer is skipped. This operation is used by the Amazon
// ECR proxy and is not generally used by customers for pulling and pushing images.
// In most cases, you should use the docker CLI to pull, tag, and push images.
func (c *Client) BatchCheckLayerAvailability(ctx context.Context, params *BatchCheckLayerAvailabilityInput, optFns ...func(*Options)) (*BatchCheckLayerAvailabilityOutput, error) {
	if params == nil {
		params = &BatchCheckLayerAvailabilityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchCheckLayerAvailability", params, optFns, c.addOperationBatchCheckLayerAvailabilityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchCheckLayerAvailabilityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchCheckLayerAvailabilityInput struct {

	// The digests of the image layers to check.
	//
	// This member is required.
	LayerDigests []string

	// The name of the repository that's associated with the image layers to check.
	//
	// This member is required.
	RepositoryName *string

	// The Amazon Web Services account ID, or registry alias, associated with the
	// public registry that contains the image layers to check. If you do not specify a
	// registry, the default public registry is assumed.
	RegistryId *string

	noSmithyDocumentSerde
}

type BatchCheckLayerAvailabilityOutput struct {

	// Any failures associated with the call.
	Failures []types.LayerFailure

	// A list of image layer objects that correspond to the image layer references in
	// the request.
	Layers []types.Layer

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchCheckLayerAvailabilityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchCheckLayerAvailability{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchCheckLayerAvailability{}, middleware.After)
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
	if err = addOpBatchCheckLayerAvailabilityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchCheckLayerAvailability(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchCheckLayerAvailability(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr-public",
		OperationName: "BatchCheckLayerAvailability",
	}
}
