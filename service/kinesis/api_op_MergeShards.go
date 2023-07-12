// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Merges two adjacent shards in a Kinesis data stream and combines them into a
// single shard to reduce the stream's capacity to ingest and transport data. This
// API is only supported for the data streams with the provisioned capacity mode.
// Two shards are considered adjacent if the union of the hash key ranges for the
// two shards form a contiguous set with no gaps. For example, if you have two
// shards, one with a hash key range of 276...381 and the other with a hash key
// range of 382...454, then you could merge these two shards into a single shard
// that would have a hash key range of 276...454. After the merge, the single child
// shard receives data for all hash key values covered by the two parent shards.
// When invoking this API, it is recommended you use the StreamARN input parameter
// rather than the StreamName input parameter. MergeShards is called when there is
// a need to reduce the overall capacity of a stream because of excess capacity
// that is not being used. You must specify the shard to be merged and the adjacent
// shard for a stream. For more information about merging shards, see Merge Two
// Shards (https://docs.aws.amazon.com/kinesis/latest/dev/kinesis-using-sdk-java-resharding-merge.html)
// in the Amazon Kinesis Data Streams Developer Guide. If the stream is in the
// ACTIVE state, you can call MergeShards . If a stream is in the CREATING ,
// UPDATING , or DELETING state, MergeShards returns a ResourceInUseException . If
// the specified stream does not exist, MergeShards returns a
// ResourceNotFoundException . You can use DescribeStreamSummary to check the
// state of the stream, which is returned in StreamStatus . MergeShards is an
// asynchronous operation. Upon receiving a MergeShards request, Amazon Kinesis
// Data Streams immediately returns a response and sets the StreamStatus to
// UPDATING . After the operation is completed, Kinesis Data Streams sets the
// StreamStatus to ACTIVE . Read and write operations continue to work while the
// stream is in the UPDATING state. You use DescribeStreamSummary and the
// ListShards APIs to determine the shard IDs that are specified in the MergeShards
// request. If you try to operate on too many streams in parallel using
// CreateStream , DeleteStream , MergeShards , or SplitShard , you receive a
// LimitExceededException . MergeShards has a limit of five transactions per
// second per account.
func (c *Client) MergeShards(ctx context.Context, params *MergeShardsInput, optFns ...func(*Options)) (*MergeShardsOutput, error) {
	if params == nil {
		params = &MergeShardsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "MergeShards", params, optFns, c.addOperationMergeShardsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*MergeShardsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for MergeShards .
type MergeShardsInput struct {

	// The shard ID of the adjacent shard for the merge.
	//
	// This member is required.
	AdjacentShardToMerge *string

	// The shard ID of the shard to combine with the adjacent shard for the merge.
	//
	// This member is required.
	ShardToMerge *string

	// The ARN of the stream.
	StreamARN *string

	// The name of the stream for the merge.
	StreamName *string

	noSmithyDocumentSerde
}

type MergeShardsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationMergeShardsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpMergeShards{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpMergeShards{}, middleware.After)
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
	if err = addOpMergeShardsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opMergeShards(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opMergeShards(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "MergeShards",
	}
}
