// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudsearch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the analysis schemes configured for a domain. An analysis scheme defines
// language-specific text processing options for a text field. Can be limited to
// specific analysis schemes by name. By default, shows all analysis schemes and
// includes any pending changes to the configuration. Set the Deployed option to
// true to show the active configuration and exclude pending changes. For more
// information, see Configuring Analysis Schemes (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-analysis-schemes.html)
// in the Amazon CloudSearch Developer Guide.
func (c *Client) DescribeAnalysisSchemes(ctx context.Context, params *DescribeAnalysisSchemesInput, optFns ...func(*Options)) (*DescribeAnalysisSchemesOutput, error) {
	if params == nil {
		params = &DescribeAnalysisSchemesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAnalysisSchemes", params, optFns, c.addOperationDescribeAnalysisSchemesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAnalysisSchemesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DescribeAnalysisSchemes operation.
// Specifies the name of the domain you want to describe. To limit the response to
// particular analysis schemes, specify the names of the analysis schemes you want
// to describe. To show the active configuration and exclude any pending changes,
// set the Deployed option to true .
type DescribeAnalysisSchemesInput struct {

	// The name of the domain you want to describe.
	//
	// This member is required.
	DomainName *string

	// The analysis schemes you want to describe.
	AnalysisSchemeNames []string

	// Whether to display the deployed configuration ( true ) or include any pending
	// changes ( false ). Defaults to false .
	Deployed *bool

	noSmithyDocumentSerde
}

// The result of a DescribeAnalysisSchemes request. Contains the analysis schemes
// configured for the domain specified in the request.
type DescribeAnalysisSchemesOutput struct {

	// The analysis scheme descriptions.
	//
	// This member is required.
	AnalysisSchemes []types.AnalysisSchemeStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAnalysisSchemesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeAnalysisSchemes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeAnalysisSchemes{}, middleware.After)
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
	if err = addOpDescribeAnalysisSchemesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAnalysisSchemes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAnalysisSchemes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudsearch",
		OperationName: "DescribeAnalysisSchemes",
	}
}
