// Code generated by smithy-go-codegen DO NOT EDIT.

package customerprofiles

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/customerprofiles/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a new calculated attribute definition. After creation, new object data
// ingested into Customer Profiles will be included in the calculated attribute,
// which can be retrieved for a profile using the GetCalculatedAttributeForProfile (https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_GetCalculatedAttributeForProfile.html)
// API. Defining a calculated attribute makes it available for all profiles within
// a domain. Each calculated attribute can only reference one ObjectType and at
// most, two fields from that ObjectType .
func (c *Client) CreateCalculatedAttributeDefinition(ctx context.Context, params *CreateCalculatedAttributeDefinitionInput, optFns ...func(*Options)) (*CreateCalculatedAttributeDefinitionOutput, error) {
	if params == nil {
		params = &CreateCalculatedAttributeDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCalculatedAttributeDefinition", params, optFns, c.addOperationCreateCalculatedAttributeDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCalculatedAttributeDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCalculatedAttributeDefinitionInput struct {

	// Mathematical expression and a list of attribute items specified in that
	// expression.
	//
	// This member is required.
	AttributeDetails *types.AttributeDetails

	// The unique name of the calculated attribute.
	//
	// This member is required.
	CalculatedAttributeName *string

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// The aggregation operation to perform for the calculated attribute.
	//
	// This member is required.
	Statistic types.Statistic

	// The conditions including range, object count, and threshold for the calculated
	// attribute.
	Conditions *types.Conditions

	// The description of the calculated attribute.
	Description *string

	// The display name of the calculated attribute.
	DisplayName *string

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateCalculatedAttributeDefinitionOutput struct {

	// Mathematical expression and a list of attribute items specified in that
	// expression.
	AttributeDetails *types.AttributeDetails

	// The unique name of the calculated attribute.
	CalculatedAttributeName *string

	// The conditions including range, object count, and threshold for the calculated
	// attribute.
	Conditions *types.Conditions

	// The timestamp of when the calculated attribute definition was created.
	CreatedAt *time.Time

	// The description of the calculated attribute.
	Description *string

	// The display name of the calculated attribute.
	DisplayName *string

	// The timestamp of when the calculated attribute definition was most recently
	// edited.
	LastUpdatedAt *time.Time

	// The aggregation operation to perform for the calculated attribute.
	Statistic types.Statistic

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCalculatedAttributeDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateCalculatedAttributeDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateCalculatedAttributeDefinition{}, middleware.After)
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
	if err = addOpCreateCalculatedAttributeDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCalculatedAttributeDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCalculatedAttributeDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "profile",
		OperationName: "CreateCalculatedAttributeDefinition",
	}
}
