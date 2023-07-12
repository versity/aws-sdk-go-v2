// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets details about a read set.
func (c *Client) GetReadSetMetadata(ctx context.Context, params *GetReadSetMetadataInput, optFns ...func(*Options)) (*GetReadSetMetadataOutput, error) {
	if params == nil {
		params = &GetReadSetMetadataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetReadSetMetadata", params, optFns, c.addOperationGetReadSetMetadataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetReadSetMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetReadSetMetadataInput struct {

	// The read set's ID.
	//
	// This member is required.
	Id *string

	// The read set's sequence store ID.
	//
	// This member is required.
	SequenceStoreId *string

	noSmithyDocumentSerde
}

type GetReadSetMetadataOutput struct {

	// The read set's ARN.
	//
	// This member is required.
	Arn *string

	// When the read set was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The read set's file type.
	//
	// This member is required.
	FileType types.FileType

	// The read set's ID.
	//
	// This member is required.
	Id *string

	// The read set's sequence store ID.
	//
	// This member is required.
	SequenceStoreId *string

	// The read set's status.
	//
	// This member is required.
	Status types.ReadSetStatus

	// The read set's description.
	Description *string

	// The read set's files.
	Files *types.ReadSetFiles

	// The read set's name.
	Name *string

	// The read set's genome reference ARN.
	ReferenceArn *string

	// The read set's sample ID.
	SampleId *string

	// The read set's sequence information.
	SequenceInformation *types.SequenceInformation

	// The status message for a read set. It provides more detail as to why the read
	// set has a status.
	StatusMessage *string

	// The read set's subject ID.
	SubjectId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetReadSetMetadataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetReadSetMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetReadSetMetadata{}, middleware.After)
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
	if err = addEndpointPrefix_opGetReadSetMetadataMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetReadSetMetadataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetReadSetMetadata(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetReadSetMetadataMiddleware struct {
}

func (*endpointPrefix_opGetReadSetMetadataMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetReadSetMetadataMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "control-storage-" + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opGetReadSetMetadataMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opGetReadSetMetadataMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opGetReadSetMetadata(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "omics",
		OperationName: "GetReadSetMetadata",
	}
}
