// Code generated by smithy-go-codegen DO NOT EDIT.

package backupstorage

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backupstorage/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
)

// Upload object that can store object metadata String and data blob in single API
// call using inline chunk field.
func (c *Client) PutObject(ctx context.Context, params *PutObjectInput, optFns ...func(*Options)) (*PutObjectOutput, error) {
	if params == nil {
		params = &PutObjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutObject", params, optFns, c.addOperationPutObjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutObjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutObjectInput struct {

	// Backup job Id for the in-progress backup.
	//
	// This member is required.
	BackupJobId *string

	// The name of the Object to be uploaded.
	//
	// This member is required.
	ObjectName *string

	// Inline chunk data to be uploaded.
	InlineChunk io.Reader

	// Inline chunk checksum
	InlineChunkChecksum *string

	// Inline chunk checksum algorithm
	InlineChunkChecksumAlgorithm *string

	// Length of the inline chunk data.
	InlineChunkLength int64

	// Store user defined metadata like backup checksum, disk ids, restore metadata
	// etc.
	MetadataString *string

	// object checksum
	ObjectChecksum *string

	// object checksum algorithm
	ObjectChecksumAlgorithm types.SummaryChecksumAlgorithm

	// Throw an exception if Object name is already exist.
	ThrowOnDuplicate bool

	noSmithyDocumentSerde
}

type PutObjectOutput struct {

	// Inline chunk checksum
	//
	// This member is required.
	InlineChunkChecksum *string

	// Inline chunk checksum algorithm
	//
	// This member is required.
	InlineChunkChecksumAlgorithm types.DataChecksumAlgorithm

	// object checksum
	//
	// This member is required.
	ObjectChecksum *string

	// object checksum algorithm
	//
	// This member is required.
	ObjectChecksumAlgorithm types.SummaryChecksumAlgorithm

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutObjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutObject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutObject{}, middleware.After)
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
	if err = v4.AddUnsignedPayloadMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
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
	if err = addOpPutObjectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutObject(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutObject(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup-storage",
		OperationName: "PutObject",
	}
}
