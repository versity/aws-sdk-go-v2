// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a recommendation run that is used to generate rules when you don't know
// what rules to write. Glue Data Quality analyzes the data and comes up with
// recommendations for a potential ruleset. You can then triage the ruleset and
// modify the generated ruleset to your liking.
func (c *Client) StartDataQualityRuleRecommendationRun(ctx context.Context, params *StartDataQualityRuleRecommendationRunInput, optFns ...func(*Options)) (*StartDataQualityRuleRecommendationRunOutput, error) {
	if params == nil {
		params = &StartDataQualityRuleRecommendationRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartDataQualityRuleRecommendationRun", params, optFns, c.addOperationStartDataQualityRuleRecommendationRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartDataQualityRuleRecommendationRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartDataQualityRuleRecommendationRunInput struct {

	// The data source (Glue table) associated with this run.
	//
	// This member is required.
	DataSource *types.DataSource

	// An IAM role supplied to encrypt the results of the run.
	//
	// This member is required.
	Role *string

	// Used for idempotency and is recommended to be set to a random ID (such as a
	// UUID) to avoid creating or starting multiple instances of the same resource.
	ClientToken *string

	// A name for the ruleset.
	CreatedRulesetName *string

	// The number of G.1X workers to be used in the run. The default is 5.
	NumberOfWorkers *int32

	// The timeout for a run in minutes. This is the maximum time that a run can
	// consume resources before it is terminated and enters TIMEOUT status. The
	// default is 2,880 minutes (48 hours).
	Timeout *int32

	noSmithyDocumentSerde
}

type StartDataQualityRuleRecommendationRunOutput struct {

	// The unique run identifier associated with this run.
	RunId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartDataQualityRuleRecommendationRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartDataQualityRuleRecommendationRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartDataQualityRuleRecommendationRun{}, middleware.After)
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
	if err = addOpStartDataQualityRuleRecommendationRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartDataQualityRuleRecommendationRun(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartDataQualityRuleRecommendationRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "StartDataQualityRuleRecommendationRun",
	}
}
