// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about an on-premises storage system that you're using with
// DataSync Discovery.
func (c *Client) DescribeStorageSystem(ctx context.Context, params *DescribeStorageSystemInput, optFns ...func(*Options)) (*DescribeStorageSystemOutput, error) {
	if params == nil {
		params = &DescribeStorageSystemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStorageSystem", params, optFns, c.addOperationDescribeStorageSystemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStorageSystemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStorageSystemInput struct {

	// Specifies the Amazon Resource Name (ARN) of an on-premises storage system that
	// you're using with DataSync Discovery.
	//
	// This member is required.
	StorageSystemArn *string

	noSmithyDocumentSerde
}

type DescribeStorageSystemOutput struct {

	// The ARN of the DataSync agent that connects to and reads from your on-premises
	// storage system.
	AgentArns []string

	// The ARN of the Amazon CloudWatch log group that's used to monitor and log
	// discovery job events.
	CloudWatchLogGroupArn *string

	// Indicates whether your DataSync agent can connect to your on-premises storage
	// system.
	ConnectivityStatus types.StorageSystemConnectivityStatus

	// The time when you added the on-premises storage system to DataSync Discovery.
	CreationTime *time.Time

	// Describes the connectivity error that the DataSync agent is encountering with
	// your on-premises storage system.
	ErrorMessage *string

	// The name that you gave your on-premises storage system when adding it to
	// DataSync Discovery.
	Name *string

	// The ARN of the secret that stores your on-premises storage system's
	// credentials. DataSync Discovery stores these credentials in Secrets Manager (https://docs.aws.amazon.com/datasync/latest/userguide/discovery-configure-storage.html#discovery-add-storage)
	// .
	SecretsManagerArn *string

	// The server name and network port required to connect with your on-premises
	// storage system's management interface.
	ServerConfiguration *types.DiscoveryServerConfiguration

	// The ARN of the on-premises storage system that the discovery job looked at.
	StorageSystemArn *string

	// The type of on-premises storage system. DataSync Discovery currently only
	// supports NetApp Fabric-Attached Storage (FAS) and All Flash FAS (AFF) systems
	// running ONTAP 9.7 or later.
	SystemType types.DiscoverySystemType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeStorageSystemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeStorageSystem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeStorageSystem{}, middleware.After)
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
	if err = addEndpointPrefix_opDescribeStorageSystemMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeStorageSystemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStorageSystem(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opDescribeStorageSystemMiddleware struct {
}

func (*endpointPrefix_opDescribeStorageSystemMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDescribeStorageSystemMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "discovery-" + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opDescribeStorageSystemMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opDescribeStorageSystemMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeStorageSystem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "DescribeStorageSystem",
	}
}
