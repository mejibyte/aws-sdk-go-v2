// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of the properties of all entity recognizers that you created,
// including recognizers currently in training. Allows you to filter the list of
// recognizers based on criteria such as status and submission time. This call
// returns up to 500 entity recognizers in the list, with a default number of 100
// recognizers in the list. The results of this list are not in any particular
// order. Please get the list and sort locally if needed.
func (c *Client) ListEntityRecognizers(ctx context.Context, params *ListEntityRecognizersInput, optFns ...func(*Options)) (*ListEntityRecognizersOutput, error) {
	if params == nil {
		params = &ListEntityRecognizersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEntityRecognizers", params, optFns, addOperationListEntityRecognizersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEntityRecognizersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEntityRecognizersInput struct {

	// Filters the list of entities returned. You can filter on Status,
	// SubmitTimeBefore, or SubmitTimeAfter. You can only set one filter at a time.
	Filter *types.EntityRecognizerFilter

	// The maximum number of results to return on each page. The default is 100.
	MaxResults *int32

	// Identifies the next page of results to return.
	NextToken *string
}

type ListEntityRecognizersOutput struct {

	// The list of properties of an entity recognizer.
	EntityRecognizerPropertiesList []types.EntityRecognizerProperties

	// Identifies the next page of results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListEntityRecognizersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListEntityRecognizers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEntityRecognizers{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEntityRecognizers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListEntityRecognizers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "ListEntityRecognizers",
	}
}
