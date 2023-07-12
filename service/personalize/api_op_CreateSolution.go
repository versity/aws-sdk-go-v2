// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates the configuration for training a model. A trained model is known as a
// solution version. After the configuration is created, you train the model
// (create a solution version) by calling the CreateSolutionVersion (https://docs.aws.amazon.com/personalize/latest/dg/API_CreateSolutionVersion.html)
// operation. Every time you call CreateSolutionVersion , a new version of the
// solution is created. After creating a solution version, you check its accuracy
// by calling GetSolutionMetrics (https://docs.aws.amazon.com/personalize/latest/dg/API_GetSolutionMetrics.html)
// . When you are satisfied with the version, you deploy it using CreateCampaign (https://docs.aws.amazon.com/personalize/latest/dg/API_CreateCampaign.html)
// . The campaign provides recommendations to a client through the
// GetRecommendations (https://docs.aws.amazon.com/personalize/latest/dg/API_RS_GetRecommendations.html)
// API. To train a model, Amazon Personalize requires training data and a recipe.
// The training data comes from the dataset group that you provide in the request.
// A recipe specifies the training algorithm and a feature transformation. You can
// specify one of the predefined recipes provided by Amazon Personalize. Amazon
// Personalize doesn't support configuring the hpoObjective for solution
// hyperparameter optimization at this time. Status A solution can be in one of the
// following states:
//   - CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//   - DELETE PENDING > DELETE IN_PROGRESS
//
// To get the status of the solution, call DescribeSolution (https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeSolution.html)
// . Wait until the status shows as ACTIVE before calling CreateSolutionVersion .
// Related APIs
//
//   - ListSolutions (https://docs.aws.amazon.com/personalize/latest/dg/API_ListSolutions.html)
//
//   - CreateSolutionVersion (https://docs.aws.amazon.com/personalize/latest/dg/API_CreateSolutionVersion.html)
//
//   - DescribeSolution (https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeSolution.html)
//
//   - DeleteSolution (https://docs.aws.amazon.com/personalize/latest/dg/API_DeleteSolution.html)
//
//   - ListSolutionVersions (https://docs.aws.amazon.com/personalize/latest/dg/API_ListSolutionVersions.html)
//
//   - DescribeSolutionVersion (https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeSolutionVersion.html)
func (c *Client) CreateSolution(ctx context.Context, params *CreateSolutionInput, optFns ...func(*Options)) (*CreateSolutionOutput, error) {
	if params == nil {
		params = &CreateSolutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSolution", params, optFns, c.addOperationCreateSolutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSolutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSolutionInput struct {

	// The Amazon Resource Name (ARN) of the dataset group that provides the training
	// data.
	//
	// This member is required.
	DatasetGroupArn *string

	// The name for the solution.
	//
	// This member is required.
	Name *string

	// When your have multiple event types (using an EVENT_TYPE schema field), this
	// parameter specifies which event type (for example, 'click' or 'like') is used
	// for training the model. If you do not provide an eventType , Amazon Personalize
	// will use all interactions for training with equal weight regardless of type.
	EventType *string

	// We don't recommend enabling automated machine learning. Instead, match your use
	// case to the available Amazon Personalize recipes. For more information, see
	// Determining your use case. (https://docs.aws.amazon.com/personalize/latest/dg/determining-use-case.html)
	// Whether to perform automated machine learning (AutoML). The default is false .
	// For this case, you must specify recipeArn . When set to true , Amazon
	// Personalize analyzes your training data and selects the optimal
	// USER_PERSONALIZATION recipe and hyperparameters. In this case, you must omit
	// recipeArn . Amazon Personalize determines the optimal recipe by running tests
	// with different values for the hyperparameters. AutoML lengthens the training
	// process as compared to selecting a specific recipe.
	PerformAutoML bool

	// Whether to perform hyperparameter optimization (HPO) on the specified or
	// selected recipe. The default is false . When performing AutoML, this parameter
	// is always true and you should not set it to false .
	PerformHPO *bool

	// The ARN of the recipe to use for model training. Only specified when
	// performAutoML is false.
	RecipeArn *string

	// The configuration to use with the solution. When performAutoML is set to true,
	// Amazon Personalize only evaluates the autoMLConfig section of the solution
	// configuration. Amazon Personalize doesn't support configuring the hpoObjective
	// at this time.
	SolutionConfig *types.SolutionConfig

	// A list of tags (https://docs.aws.amazon.com/personalize/latest/dg/tagging-resources.html)
	// to apply to the solution.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateSolutionOutput struct {

	// The ARN of the solution.
	SolutionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSolutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateSolution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateSolution{}, middleware.After)
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
	if err = addOpCreateSolutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSolution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSolution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "CreateSolution",
	}
}
