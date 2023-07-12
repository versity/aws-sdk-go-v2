// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a simulation application with a specific revision id.
func (c *Client) CreateSimulationApplicationVersion(ctx context.Context, params *CreateSimulationApplicationVersionInput, optFns ...func(*Options)) (*CreateSimulationApplicationVersionOutput, error) {
	if params == nil {
		params = &CreateSimulationApplicationVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSimulationApplicationVersion", params, optFns, c.addOperationCreateSimulationApplicationVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSimulationApplicationVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSimulationApplicationVersionInput struct {

	// The application information for the simulation application.
	//
	// This member is required.
	Application *string

	// The current revision id for the simulation application. If you provide a value
	// and it matches the latest revision ID, a new version will be created.
	CurrentRevisionId *string

	// The SHA256 digest used to identify the Docker image URI used to created the
	// simulation application.
	ImageDigest *string

	// The Amazon S3 eTag identifier for the zip file bundle that you use to create
	// the simulation application.
	S3Etags []string

	noSmithyDocumentSerde
}

type CreateSimulationApplicationVersionOutput struct {

	// The Amazon Resource Name (ARN) of the simulation application.
	Arn *string

	// The object that contains the Docker image URI used to create the simulation
	// application.
	Environment *types.Environment

	// The time, in milliseconds since the epoch, when the simulation application was
	// last updated.
	LastUpdatedAt *time.Time

	// The name of the simulation application.
	Name *string

	// The rendering engine for the simulation application.
	RenderingEngine *types.RenderingEngine

	// The revision ID of the simulation application.
	RevisionId *string

	// Information about the robot software suite (ROS distribution).
	RobotSoftwareSuite *types.RobotSoftwareSuite

	// The simulation software suite used by the simulation application.
	SimulationSoftwareSuite *types.SimulationSoftwareSuite

	// The sources of the simulation application.
	Sources []types.Source

	// The version of the simulation application.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSimulationApplicationVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSimulationApplicationVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSimulationApplicationVersion{}, middleware.After)
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
	if err = addOpCreateSimulationApplicationVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSimulationApplicationVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSimulationApplicationVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "CreateSimulationApplicationVersion",
	}
}
