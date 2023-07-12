// Code generated by smithy-go-codegen DO NOT EDIT.

package healthlake

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/healthlake/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Displays the properties of a FHIR import job, including the ID, ARN, name, and
// the status of the job.
func (c *Client) DescribeFHIRImportJob(ctx context.Context, params *DescribeFHIRImportJobInput, optFns ...func(*Options)) (*DescribeFHIRImportJobOutput, error) {
	if params == nil {
		params = &DescribeFHIRImportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFHIRImportJob", params, optFns, c.addOperationDescribeFHIRImportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFHIRImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFHIRImportJobInput struct {

	// The AWS-generated ID of the Data Store.
	//
	// This member is required.
	DatastoreId *string

	// The AWS-generated job ID.
	//
	// This member is required.
	JobId *string

	noSmithyDocumentSerde
}

type DescribeFHIRImportJobOutput struct {

	// The properties of the Import job request, including the ID, ARN, name, and the
	// status of the job.
	//
	// This member is required.
	ImportJobProperties *types.ImportJobProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFHIRImportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeFHIRImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeFHIRImportJob{}, middleware.After)
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
	if err = addOpDescribeFHIRImportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFHIRImportJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeFHIRImportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "healthlake",
		OperationName: "DescribeFHIRImportJob",
	}
}
