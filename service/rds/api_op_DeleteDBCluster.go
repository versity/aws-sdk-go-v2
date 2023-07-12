// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The DeleteDBCluster action deletes a previously provisioned DB cluster. When
// you delete a DB cluster, all automated backups for that DB cluster are deleted
// and can't be recovered. Manual DB cluster snapshots of the specified DB cluster
// are not deleted. If you're deleting a Multi-AZ DB cluster with read replicas,
// all cluster members are terminated and read replicas are promoted to standalone
// instances. For more information on Amazon Aurora, see What is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide. For more information on Multi-AZ DB clusters,
// see Multi-AZ DB cluster deployments (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html)
// in the Amazon RDS User Guide.
func (c *Client) DeleteDBCluster(ctx context.Context, params *DeleteDBClusterInput, optFns ...func(*Options)) (*DeleteDBClusterOutput, error) {
	if params == nil {
		params = &DeleteDBClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDBCluster", params, optFns, c.addOperationDeleteDBClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDBClusterInput struct {

	// The DB cluster identifier for the DB cluster to be deleted. This parameter
	// isn't case-sensitive. Constraints:
	//   - Must match an existing DBClusterIdentifier.
	//
	// This member is required.
	DBClusterIdentifier *string

	// The DB cluster snapshot identifier of the new DB cluster snapshot created when
	// SkipFinalSnapshot is disabled. Specifying this parameter and also skipping the
	// creation of a final DB cluster snapshot with the SkipFinalShapshot parameter
	// results in an error. Constraints:
	//   - Must be 1 to 255 letters, numbers, or hyphens.
	//   - First character must be a letter
	//   - Can't end with a hyphen or contain two consecutive hyphens
	FinalDBSnapshotIdentifier *string

	// A value that indicates whether to skip the creation of a final DB cluster
	// snapshot before the DB cluster is deleted. If skip is specified, no DB cluster
	// snapshot is created. If skip isn't specified, a DB cluster snapshot is created
	// before the DB cluster is deleted. By default, skip isn't specified, and the DB
	// cluster snapshot is created. By default, this parameter is disabled. You must
	// specify a FinalDBSnapshotIdentifier parameter if SkipFinalSnapshot is disabled.
	SkipFinalSnapshot bool

	noSmithyDocumentSerde
}

type DeleteDBClusterOutput struct {

	// Contains the details of an Amazon Aurora DB cluster or Multi-AZ DB cluster. For
	// an Amazon Aurora DB cluster, this data type is used as a response element in the
	// operations CreateDBCluster , DeleteDBCluster , DescribeDBClusters ,
	// FailoverDBCluster , ModifyDBCluster , PromoteReadReplicaDBCluster ,
	// RestoreDBClusterFromS3 , RestoreDBClusterFromSnapshot ,
	// RestoreDBClusterToPointInTime , StartDBCluster , and StopDBCluster . For a
	// Multi-AZ DB cluster, this data type is used as a response element in the
	// operations CreateDBCluster , DeleteDBCluster , DescribeDBClusters ,
	// FailoverDBCluster , ModifyDBCluster , RebootDBCluster ,
	// RestoreDBClusterFromSnapshot , and RestoreDBClusterToPointInTime . For more
	// information on Amazon Aurora DB clusters, see What is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
	// in the Amazon Aurora User Guide. For more information on Multi-AZ DB clusters,
	// see Multi-AZ deployments with two readable standby DB instances (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html)
	// in the Amazon RDS User Guide.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteDBClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteDBCluster{}, middleware.After)
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
	if err = addOpDeleteDBClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDBCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteDBCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DeleteDBCluster",
	}
}
