// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the name, ARN, NumberOfRules and effective dates of all Cost Categories
// defined in the account. You have the option to use EffectiveOn to return a list
// of Cost Categories that were active on a specific date. If there is no
// EffectiveOn specified, you’ll see Cost Categories that are effective on the
// current date. If Cost Category is still effective, EffectiveEnd is omitted in
// the response. ListCostCategoryDefinitions supports pagination. The request can
// have a MaxResults range up to 100.
func (c *Client) ListCostCategoryDefinitions(ctx context.Context, params *ListCostCategoryDefinitionsInput, optFns ...func(*Options)) (*ListCostCategoryDefinitionsOutput, error) {
	if params == nil {
		params = &ListCostCategoryDefinitionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCostCategoryDefinitions", params, optFns, addOperationListCostCategoryDefinitionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCostCategoryDefinitionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCostCategoryDefinitionsInput struct {

	// The date when the Cost Category was effective.
	EffectiveOn *string

	// The number of entries a paginated response contains.
	MaxResults int32

	// The token to retrieve the next set of results. Amazon Web Services provides the
	// token when the response from a previous call has more results than the maximum
	// page size.
	NextToken *string
}

type ListCostCategoryDefinitionsOutput struct {

	// A reference to a Cost Category containing enough information to identify the
	// Cost Category.
	CostCategoryReferences []types.CostCategoryReference

	// The token to retrieve the next set of results. Amazon Web Services provides the
	// token when the response from a previous call has more results than the maximum
	// page size.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListCostCategoryDefinitionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCostCategoryDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCostCategoryDefinitions{}, middleware.After)
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
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCostCategoryDefinitions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListCostCategoryDefinitions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "ListCostCategoryDefinitions",
	}
}
