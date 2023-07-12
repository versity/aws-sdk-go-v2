// Code generated by smithy-go-codegen DO NOT EDIT.

package honeycode

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/honeycode/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The DescribeTableDataImportJob API allows you to retrieve the status and
// details of a table data import job.
func (c *Client) DescribeTableDataImportJob(ctx context.Context, params *DescribeTableDataImportJobInput, optFns ...func(*Options)) (*DescribeTableDataImportJobOutput, error) {
	if params == nil {
		params = &DescribeTableDataImportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTableDataImportJob", params, optFns, c.addOperationDescribeTableDataImportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTableDataImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTableDataImportJobInput struct {

	// The ID of the job that was returned by the StartTableDataImportJob request. If
	// a job with the specified id could not be found, this API throws
	// ResourceNotFoundException.
	//
	// This member is required.
	JobId *string

	// The ID of the table into which data was imported. If a table with the specified
	// id could not be found, this API throws ResourceNotFoundException.
	//
	// This member is required.
	TableId *string

	// The ID of the workbook into which data was imported. If a workbook with the
	// specified id could not be found, this API throws ResourceNotFoundException.
	//
	// This member is required.
	WorkbookId *string

	noSmithyDocumentSerde
}

type DescribeTableDataImportJobOutput struct {

	// The metadata about the job that was submitted for import.
	//
	// This member is required.
	JobMetadata *types.TableDataImportJobMetadata

	// The current status of the import job.
	//
	// This member is required.
	JobStatus types.TableDataImportJobStatus

	// A message providing more details about the current status of the import job.
	//
	// This member is required.
	Message *string

	// If job status is failed, error code to understand reason for the failure.
	ErrorCode types.ErrorCode

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTableDataImportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeTableDataImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeTableDataImportJob{}, middleware.After)
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
	if err = addOpDescribeTableDataImportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTableDataImportJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeTableDataImportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "honeycode",
		OperationName: "DescribeTableDataImportJob",
	}
}
