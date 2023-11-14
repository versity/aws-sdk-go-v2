// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Initiates a flow to start a new chat for the customer. Response of this API
// provides a token required to obtain credentials from the
// CreateParticipantConnection (https://docs.aws.amazon.com/connect-participant/latest/APIReference/API_CreateParticipantConnection.html)
// API in the Amazon Connect Participant Service. When a new chat contact is
// successfully created, clients must subscribe to the participant’s connection for
// the created chat within 5 minutes. This is achieved by invoking
// CreateParticipantConnection (https://docs.aws.amazon.com/connect-participant/latest/APIReference/API_CreateParticipantConnection.html)
// with WEBSOCKET and CONNECTION_CREDENTIALS. A 429 error occurs in the following
// situations:
//   - API rate limit is exceeded. API TPS throttling returns a TooManyRequests
//     exception.
//   - The quota for concurrent active chats (https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html)
//     is exceeded. Active chat throttling returns a LimitExceededException .
//
// If you use the ChatDurationInMinutes parameter and receive a 400 error, your
// account may not support the ability to configure custom chat durations. For more
// information, contact Amazon Web Services Support. For more information about
// chat, see Chat (https://docs.aws.amazon.com/connect/latest/adminguide/chat.html)
// in the Amazon Connect Administrator Guide.
func (c *Client) StartChatContact(ctx context.Context, params *StartChatContactInput, optFns ...func(*Options)) (*StartChatContactOutput, error) {
	if params == nil {
		params = &StartChatContactInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartChatContact", params, optFns, c.addOperationStartChatContactMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartChatContactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartChatContactInput struct {

	// The identifier of the flow for initiating the chat. To see the ContactFlowId in
	// the Amazon Connect console user interface, on the navigation menu go to Routing,
	// Contact Flows. Choose the flow. On the flow page, under the name of the flow,
	// choose Show additional flow information. The ContactFlowId is the last part of
	// the ARN, shown here in bold:
	// arn:aws:connect:us-west-2:xxxxxxxxxxxx:instance/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/contact-flow/846ec553-a005-41c0-8341-xxxxxxxxxxxx
	//
	// This member is required.
	ContactFlowId *string

	// The identifier of the Amazon Connect instance. You can find the instance ID (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// Information identifying the participant.
	//
	// This member is required.
	ParticipantDetails *types.ParticipantDetails

	// A custom key-value pair using an attribute map. The attributes are standard
	// Amazon Connect attributes. They can be accessed in flows just like any other
	// contact attributes. There can be up to 32,768 UTF-8 bytes across all key-value
	// pairs per contact. Attribute keys can include only alphanumeric, dash, and
	// underscore characters.
	Attributes map[string]string

	// The total duration of the newly started chat session. If not specified, the
	// chat session duration defaults to 25 hour. The minimum configurable time is 60
	// minutes. The maximum configurable time is 10,080 minutes (7 days).
	ChatDurationInMinutes *int32

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. If not provided, the Amazon Web Services SDK populates this
	// field. For more information about idempotency, see Making retries safe with
	// idempotent APIs (https://aws.amazon.com/builders-library/making-retries-safe-with-idempotent-APIs/)
	// .
	ClientToken *string

	// The initial message to be sent to the newly created chat.
	InitialMessage *types.ChatMessage

	// Enable persistent chats. For more information about enabling persistent chat,
	// and for example use cases and how to configure for them, see Enable persistent
	// chat (https://docs.aws.amazon.com/connect/latest/adminguide/chat-persistence.html)
	// .
	PersistentChat *types.PersistentChat

	// The unique identifier for an Amazon Connect contact. This identifier is related
	// to the chat starting. You cannot provide data for both RelatedContactId and
	// PersistentChat.
	RelatedContactId *string

	// A set of system defined key-value pairs stored on individual contact segments
	// using an attribute map. The attributes are standard Amazon Connect attributes.
	// They can be accessed in flows. Attribute keys can include only alphanumeric, -,
	// and _. This field can be used to show channel subtype, such as connect:Guide .
	// The types application/vnd.amazonaws.connect.message.interactive and
	// application/vnd.amazonaws.connect.message.interactive.response must be present
	// in the SupportedMessagingContentTypes field of this API in order to set
	// SegmentAttributes as { "connect:Subtype": {"valueString" : "connect:Guide" }} .
	SegmentAttributes map[string]types.SegmentAttributeValue

	// The supported chat message content types. Supported types are text/plain ,
	// text/markdown , application/json ,
	// application/vnd.amazonaws.connect.message.interactive , and
	// application/vnd.amazonaws.connect.message.interactive.response . Content types
	// must always contain text/plain . You can then put any other supported type in
	// the list. For example, all the following lists are valid because they contain
	// text/plain : [text/plain, text/markdown, application/json] , [text/markdown,
	// text/plain] , [text/plain, application/json,
	// application/vnd.amazonaws.connect.message.interactive.response] . The type
	// application/vnd.amazonaws.connect.message.interactive is required to use the
	// Show view (https://docs.aws.amazon.com/connect/latest/adminguide/show-view-block.html)
	// flow block.
	SupportedMessagingContentTypes []string

	noSmithyDocumentSerde
}

type StartChatContactOutput struct {

	// The identifier of this contact within the Amazon Connect instance.
	ContactId *string

	// The contactId from which a persistent chat session is started. This field is
	// populated only for persistent chats.
	ContinuedFromContactId *string

	// The identifier for a chat participant. The participantId for a chat participant
	// is the same throughout the chat lifecycle.
	ParticipantId *string

	// The token used by the chat participant to call CreateParticipantConnection (https://docs.aws.amazon.com/connect-participant/latest/APIReference/API_CreateParticipantConnection.html)
	// . The participant token is valid for the lifetime of a chat participant.
	ParticipantToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartChatContactMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartChatContact{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartChatContact{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = addStartChatContactResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addIdempotencyToken_opStartChatContactMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartChatContactValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartChatContact(options.Region), middleware.Before); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpStartChatContact struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartChatContact) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartChatContact) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartChatContactInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartChatContactInput ")
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
func addIdempotencyToken_opStartChatContactMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartChatContact{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartChatContact(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "StartChatContact",
	}
}

type opStartChatContactResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opStartChatContactResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opStartChatContactResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "connect"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "connect"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("connect")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addStartChatContactResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opStartChatContactResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
