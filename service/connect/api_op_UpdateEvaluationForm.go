// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates details about a specific evaluation form version in the specified
// Amazon Connect instance. Question and section identifiers cannot be duplicated
// within the same evaluation form. This operation does not support partial
// updates. Instead it does a full update of evaluation form content.
func (c *Client) UpdateEvaluationForm(ctx context.Context, params *UpdateEvaluationFormInput, optFns ...func(*Options)) (*UpdateEvaluationFormOutput, error) {
	if params == nil {
		params = &UpdateEvaluationFormInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEvaluationForm", params, optFns, c.addOperationUpdateEvaluationFormMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEvaluationFormOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEvaluationFormInput struct {

	// The unique identifier for the evaluation form.
	//
	// This member is required.
	EvaluationFormId *string

	// A version of the evaluation form to update.
	//
	// This member is required.
	EvaluationFormVersion *int32

	// The identifier of the Amazon Connect instance. You can find the instance ID (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// Items that are part of the evaluation form. The total number of sections and
	// questions must not exceed 100 each. Questions must be contained in a section.
	//
	// This member is required.
	Items []types.EvaluationFormItem

	// A title of the evaluation form.
	//
	// This member is required.
	Title *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. If not provided, the Amazon Web Services SDK populates this
	// field. For more information about idempotency, see Making retries safe with
	// idempotent APIs (https://aws.amazon.com/builders-library/making-retries-safe-with-idempotent-APIs/)
	// .
	ClientToken *string

	// A flag indicating whether the operation must create a new version.
	CreateNewVersion *bool

	// The description of the evaluation form.
	Description *string

	// A scoring strategy of the evaluation form.
	ScoringStrategy *types.EvaluationFormScoringStrategy

	noSmithyDocumentSerde
}

type UpdateEvaluationFormOutput struct {

	// The Amazon Resource Name (ARN) for the contact evaluation resource.
	//
	// This member is required.
	EvaluationFormArn *string

	// The unique identifier for the evaluation form.
	//
	// This member is required.
	EvaluationFormId *string

	// The version of the updated evaluation form resource.
	//
	// This member is required.
	EvaluationFormVersion *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateEvaluationFormMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateEvaluationForm{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateEvaluationForm{}, middleware.After)
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
	if err = addIdempotencyToken_opUpdateEvaluationFormMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateEvaluationFormValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEvaluationForm(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateEvaluationForm struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateEvaluationForm) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateEvaluationForm) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateEvaluationFormInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateEvaluationFormInput ")
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
func addIdempotencyToken_opUpdateEvaluationFormMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateEvaluationForm{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateEvaluationForm(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "UpdateEvaluationForm",
	}
}
