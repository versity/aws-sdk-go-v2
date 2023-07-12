// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an IAM entity to describe an identity provider (IdP) that supports
// OpenID Connect (OIDC) (http://openid.net/connect/) . The OIDC provider that you
// create with this operation can be used as a principal in a role's trust policy.
// Such a policy establishes a trust relationship between Amazon Web Services and
// the OIDC provider. If you are using an OIDC identity provider from Google,
// Facebook, or Amazon Cognito, you don't need to create a separate IAM identity
// provider. These OIDC identity providers are already built-in to Amazon Web
// Services and are available for your use. Instead, you can move directly to
// creating new roles using your identity provider. To learn more, see Creating a
// role for web identity or OpenID connect federation (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-idp_oidc.html)
// in the IAM User Guide. When you create the IAM OIDC provider, you specify the
// following:
//   - The URL of the OIDC identity provider (IdP) to trust
//   - A list of client IDs (also known as audiences) that identify the
//     application or applications allowed to authenticate using the OIDC provider
//   - A list of tags that are attached to the specified IAM OIDC provider
//   - A list of thumbprints of one or more server certificates that the IdP uses
//
// You get all of this information from the OIDC IdP you want to use to access
// Amazon Web Services. Amazon Web Services secures communication with some OIDC
// identity providers (IdPs) through our library of trusted certificate authorities
// (CAs) instead of using a certificate thumbprint to verify your IdP server
// certificate. These OIDC IdPs include Google, Auth0, and those that use an Amazon
// S3 bucket to host a JSON Web Key Set (JWKS) endpoint. In these cases, your
// legacy thumbprint remains in your configuration, but is no longer used for
// validation. The trust for the OIDC provider is derived from the IAM provider
// that this operation creates. Therefore, it is best to limit access to the
// CreateOpenIDConnectProvider operation to highly privileged users.
func (c *Client) CreateOpenIDConnectProvider(ctx context.Context, params *CreateOpenIDConnectProviderInput, optFns ...func(*Options)) (*CreateOpenIDConnectProviderOutput, error) {
	if params == nil {
		params = &CreateOpenIDConnectProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateOpenIDConnectProvider", params, optFns, c.addOperationCreateOpenIDConnectProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateOpenIDConnectProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateOpenIDConnectProviderInput struct {

	// A list of server certificate thumbprints for the OpenID Connect (OIDC) identity
	// provider's server certificates. Typically this list includes only one entry.
	// However, IAM lets you have up to five thumbprints for an OIDC provider. This
	// lets you maintain multiple thumbprints if the identity provider is rotating
	// certificates. The server certificate thumbprint is the hex-encoded SHA-1 hash
	// value of the X.509 certificate used by the domain where the OpenID Connect
	// provider makes its keys available. It is always a 40-character string. You must
	// provide at least one thumbprint when creating an IAM OIDC provider. For example,
	// assume that the OIDC provider is server.example.com and the provider stores its
	// keys at https://keys.server.example.com/openid-connect. In that case, the
	// thumbprint string would be the hex-encoded SHA-1 hash value of the certificate
	// used by https://keys.server.example.com. For more information about obtaining
	// the OIDC provider thumbprint, see Obtaining the thumbprint for an OpenID
	// Connect provider (https://docs.aws.amazon.com/IAM/latest/UserGuide/identity-providers-oidc-obtain-thumbprint.html)
	// in the IAM user Guide.
	//
	// This member is required.
	ThumbprintList []string

	// The URL of the identity provider. The URL must begin with https:// and should
	// correspond to the iss claim in the provider's OpenID Connect ID tokens. Per the
	// OIDC standard, path components are allowed but query parameters are not.
	// Typically the URL consists of only a hostname, like https://server.example.org
	// or https://example.com . The URL should not contain a port number. You cannot
	// register the same provider multiple times in a single Amazon Web Services
	// account. If you try to submit a URL that has already been used for an OpenID
	// Connect provider in the Amazon Web Services account, you will get an error.
	//
	// This member is required.
	Url *string

	// Provides a list of client IDs, also known as audiences. When a mobile or web
	// app registers with an OpenID Connect provider, they establish a value that
	// identifies the application. This is the value that's sent as the client_id
	// parameter on OAuth requests. You can register multiple client IDs with the same
	// provider. For example, you might have multiple applications that use the same
	// OIDC provider. You cannot register more than 100 client IDs with a single IAM
	// OIDC provider. There is no defined format for a client ID. The
	// CreateOpenIDConnectProviderRequest operation accepts client IDs up to 255
	// characters long.
	ClientIDList []string

	// A list of tags that you want to attach to the new IAM OpenID Connect (OIDC)
	// provider. Each tag consists of a key name and an associated value. For more
	// information about tagging, see Tagging IAM resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html)
	// in the IAM User Guide. If any one of the tags is invalid or if you exceed the
	// allowed maximum number of tags, then the entire request fails and the resource
	// is not created.
	Tags []types.Tag

	noSmithyDocumentSerde
}

// Contains the response to a successful CreateOpenIDConnectProvider request.
type CreateOpenIDConnectProviderOutput struct {

	// The Amazon Resource Name (ARN) of the new IAM OpenID Connect provider that is
	// created. For more information, see OpenIDConnectProviderListEntry .
	OpenIDConnectProviderArn *string

	// A list of tags that are attached to the new IAM OIDC provider. The returned
	// list of tags is sorted by tag key. For more information about tagging, see
	// Tagging IAM resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html)
	// in the IAM User Guide.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateOpenIDConnectProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateOpenIDConnectProvider{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateOpenIDConnectProvider{}, middleware.After)
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
	if err = addOpCreateOpenIDConnectProviderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateOpenIDConnectProvider(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateOpenIDConnectProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "CreateOpenIDConnectProvider",
	}
}
