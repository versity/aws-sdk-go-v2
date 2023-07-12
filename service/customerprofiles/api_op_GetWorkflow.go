// Code generated by smithy-go-codegen DO NOT EDIT.

package customerprofiles

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/customerprofiles/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Get details of specified workflow.
func (c *Client) GetWorkflow(ctx context.Context, params *GetWorkflowInput, optFns ...func(*Options)) (*GetWorkflowOutput, error) {
	if params == nil {
		params = &GetWorkflowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetWorkflow", params, optFns, c.addOperationGetWorkflowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetWorkflowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWorkflowInput struct {

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// Unique identifier for the workflow.
	//
	// This member is required.
	WorkflowId *string

	noSmithyDocumentSerde
}

type GetWorkflowOutput struct {

	// Attributes provided for workflow execution.
	Attributes *types.WorkflowAttributes

	// Workflow error messages during execution (if any).
	ErrorDescription *string

	// The timestamp that represents when workflow execution last updated.
	LastUpdatedAt *time.Time

	// Workflow specific execution metrics.
	Metrics *types.WorkflowMetrics

	// The timestamp that represents when workflow execution started.
	StartDate *time.Time

	// Status of workflow execution.
	Status types.Status

	// Unique identifier for the workflow.
	WorkflowId *string

	// The type of workflow. The only supported value is APPFLOW_INTEGRATION.
	WorkflowType types.WorkflowType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetWorkflowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetWorkflow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetWorkflow{}, middleware.After)
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
	if err = addOpGetWorkflowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetWorkflow(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetWorkflow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "profile",
		OperationName: "GetWorkflow",
	}
}
