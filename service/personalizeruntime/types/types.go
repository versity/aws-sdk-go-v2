// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// An object that identifies an action. The API returns a list of PredictedAction s.
type PredictedAction struct {

	// The ID of the recommended action.
	ActionId *string

	// The score of the recommended action. For information about action scores, see
	// How action recommendation scoring works (https://docs.aws.amazon.com/personalize/latest/dg/how-action-recommendation-scoring-works.html)
	// .
	Score *float64

	noSmithyDocumentSerde
}

// An object that identifies an item. The and APIs return a list of PredictedItem s.
type PredictedItem struct {

	// The recommended item ID.
	ItemId *string

	// Metadata about the item from your Items dataset.
	Metadata map[string]string

	// The name of the promotion that included the predicted item.
	PromotionName *string

	// A numeric representation of the model's certainty that the item will be the
	// next user selection. For more information on scoring logic, see how-scores-work .
	Score *float64

	noSmithyDocumentSerde
}

// Contains information on a promotion. A promotion defines additional business
// rules that apply to a configurable subset of recommended items.
type Promotion struct {

	// The Amazon Resource Name (ARN) of the filter used by the promotion. This filter
	// defines the criteria for promoted items. For more information, see Promotion
	// filters (https://docs.aws.amazon.com/personalize/latest/dg/promoting-items.html#promotion-filters)
	// .
	FilterArn *string

	// The values to use when promoting items. For each placeholder parameter in your
	// promotion's filter expression, provide the parameter name (in matching case) as
	// a key and the filter value(s) as the corresponding value. Separate multiple
	// values for one parameter with a comma. For filter expressions that use an
	// INCLUDE element to include items, you must provide values for all parameters
	// that are defined in the expression. For filters with expressions that use an
	// EXCLUDE element to exclude items, you can omit the filter-values . In this case,
	// Amazon Personalize doesn't use that portion of the expression to filter
	// recommendations. For more information on creating filters, see Filtering
	// recommendations and user segments (https://docs.aws.amazon.com/personalize/latest/dg/filter.html)
	// .
	FilterValues map[string]string

	// The name of the promotion.
	Name *string

	// The percentage of recommended items to apply the promotion to.
	PercentPromotedItems *int32

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
