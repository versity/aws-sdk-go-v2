// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies an existing DMS Serverless replication configuration that you can use
// to start a replication. This command includes input validation and logic to
// check the state of any replication that uses this configuration. You can only
// modify a replication configuration before any replication that uses it has
// started. As soon as you have initially started a replication with a given
// configuiration, you can't modify that configuration, even if you stop it. Other
// run statuses that allow you to run this command include FAILED and CREATED. A
// provisioning state that allows you to run this command is FAILED_PROVISION.
func (c *Client) ModifyReplicationConfig(ctx context.Context, params *ModifyReplicationConfigInput, optFns ...func(*Options)) (*ModifyReplicationConfigOutput, error) {
	if params == nil {
		params = &ModifyReplicationConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyReplicationConfig", params, optFns, c.addOperationModifyReplicationConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyReplicationConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyReplicationConfigInput struct {

	// The Amazon Resource Name of the replication to modify.
	//
	// This member is required.
	ReplicationConfigArn *string

	// Configuration parameters for provisioning an DMS Serverless replication.
	ComputeConfig *types.ComputeConfig

	// The new replication config to apply to the replication.
	ReplicationConfigIdentifier *string

	// The settings for the replication.
	ReplicationSettings *string

	// The type of replication.
	ReplicationType types.MigrationTypeValue

	// The Amazon Resource Name (ARN) of the source endpoint for this DMS serverless
	// replication configuration.
	SourceEndpointArn *string

	// Additional settings for the replication.
	SupplementalSettings *string

	// Table mappings specified in the replication.
	TableMappings *string

	// The Amazon Resource Name (ARN) of the target endpoint for this DMS serverless
	// replication configuration.
	TargetEndpointArn *string

	noSmithyDocumentSerde
}

type ModifyReplicationConfigOutput struct {

	// Information about the serverless replication config that was modified.
	ReplicationConfig *types.ReplicationConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyReplicationConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpModifyReplicationConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyReplicationConfig{}, middleware.After)
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
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpModifyReplicationConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyReplicationConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyReplicationConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "ModifyReplicationConfig",
	}
}
