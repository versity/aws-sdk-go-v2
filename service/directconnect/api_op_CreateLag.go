// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a link aggregation group (LAG) with the specified number of bundled
// physical dedicated connections between the customer network and a specific
// Direct Connect location. A LAG is a logical interface that uses the Link
// Aggregation Control Protocol (LACP) to aggregate multiple interfaces, enabling
// you to treat them as a single interface. All connections in a LAG must use the
// same bandwidth (either 1Gbps or 10Gbps) and must terminate at the same Direct
// Connect endpoint. You can have up to 10 dedicated connections per LAG.
// Regardless of this limit, if you request more connections for the LAG than
// Direct Connect can allocate on a single endpoint, no LAG is created. You can
// specify an existing physical dedicated connection or interconnect to include in
// the LAG (which counts towards the total number of connections). Doing so
// interrupts the current physical dedicated connection, and re-establishes them as
// a member of the LAG. The LAG will be created on the same Direct Connect endpoint
// to which the dedicated connection terminates. Any virtual interfaces associated
// with the dedicated connection are automatically disassociated and re-associated
// with the LAG. The connection ID does not change. If the Amazon Web Services
// account used to create a LAG is a registered Direct Connect Partner, the LAG is
// automatically enabled to host sub-connections. For a LAG owned by a partner, any
// associated virtual interfaces cannot be directly configured.
func (c *Client) CreateLag(ctx context.Context, params *CreateLagInput, optFns ...func(*Options)) (*CreateLagOutput, error) {
	if params == nil {
		params = &CreateLagInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLag", params, optFns, c.addOperationCreateLagMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLagOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLagInput struct {

	// The bandwidth of the individual physical dedicated connections bundled by the
	// LAG. The possible values are 1Gbps and 10Gbps.
	//
	// This member is required.
	ConnectionsBandwidth *string

	// The name of the LAG.
	//
	// This member is required.
	LagName *string

	// The location for the LAG.
	//
	// This member is required.
	Location *string

	// The number of physical dedicated connections initially provisioned and bundled
	// by the LAG. You can have a maximum of four connections when the port speed is 1G
	// or 10G, or two when the port speed is 100G.
	//
	// This member is required.
	NumberOfConnections int32

	// The tags to associate with the automtically created LAGs.
	ChildConnectionTags []types.Tag

	// The ID of an existing dedicated connection to migrate to the LAG.
	ConnectionId *string

	// The name of the service provider associated with the LAG.
	ProviderName *string

	// Indicates whether the connection will support MAC Security (MACsec). All
	// connections in the LAG must be capable of supporting MAC Security (MACsec). For
	// information about MAC Security (MACsec) prerequisties, see MACsec prerequisties (https://docs.aws.amazon.com/directconnect/latest/UserGuide/direct-connect-mac-sec-getting-started.html#mac-sec-prerequisites)
	// in the Direct Connect User Guide.
	RequestMACSec *bool

	// The tags to associate with the LAG.
	Tags []types.Tag

	noSmithyDocumentSerde
}

// Information about a link aggregation group (LAG).
type CreateLagOutput struct {

	// Indicates whether the LAG can host other connections.
	AllowsHostedConnections bool

	// The Direct Connect endpoint that hosts the LAG.
	//
	// Deprecated: This member has been deprecated.
	AwsDevice *string

	// The Direct Connect endpoint that hosts the LAG.
	AwsDeviceV2 *string

	// The Direct Connect endpoint that terminates the logical connection. This device
	// might be different than the device that terminates the physical connection.
	AwsLogicalDeviceId *string

	// The connections bundled by the LAG.
	Connections []types.Connection

	// The individual bandwidth of the physical connections bundled by the LAG. The
	// possible values are 1Gbps and 10Gbps.
	ConnectionsBandwidth *string

	// The LAG MAC Security (MACsec) encryption mode. The valid values are no_encrypt ,
	// should_encrypt , and must_encrypt .
	EncryptionMode *string

	// Indicates whether the LAG supports a secondary BGP peer in the same address
	// family (IPv4/IPv6).
	HasLogicalRedundancy types.HasLogicalRedundancy

	// Indicates whether jumbo frames are supported.
	JumboFrameCapable *bool

	// The ID of the LAG.
	LagId *string

	// The name of the LAG.
	LagName *string

	// The state of the LAG. The following are the possible values:
	//   - requested : The initial state of a LAG. The LAG stays in the requested state
	//   until the Letter of Authorization (LOA) is available.
	//   - pending : The LAG has been approved and is being initialized.
	//   - available : The network link is established and the LAG is ready for use.
	//   - down : The network link is down.
	//   - deleting : The LAG is being deleted.
	//   - deleted : The LAG is deleted.
	//   - unknown : The state of the LAG is not available.
	LagState types.LagState

	// The location of the LAG.
	Location *string

	// Indicates whether the LAG supports MAC Security (MACsec).
	MacSecCapable *bool

	// The MAC Security (MACsec) security keys associated with the LAG.
	MacSecKeys []types.MacSecKey

	// The minimum number of physical dedicated connections that must be operational
	// for the LAG itself to be operational.
	MinimumLinks int32

	// The number of physical dedicated connections bundled by the LAG, up to a
	// maximum of 10.
	NumberOfConnections int32

	// The ID of the Amazon Web Services account that owns the LAG.
	OwnerAccount *string

	// The name of the service provider associated with the LAG.
	ProviderName *string

	// The Amazon Web Services Region where the connection is located.
	Region *string

	// The tags associated with the LAG.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLagMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLag{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLag{}, middleware.After)
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
	if err = addOpCreateLagValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLag(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLag(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "CreateLag",
	}
}
