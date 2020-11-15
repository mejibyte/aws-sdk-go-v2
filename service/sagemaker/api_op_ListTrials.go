// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Lists the trials in your account. Specify an experiment name to limit the list
// to the trials that are part of that experiment. Specify a trial component name
// to limit the list to the trials that associated with that trial component. The
// list can be filtered to show only trials that were created in a specific time
// range. The list can be sorted by trial name or creation time.
func (c *Client) ListTrials(ctx context.Context, params *ListTrialsInput, optFns ...func(*Options)) (*ListTrialsOutput, error) {
	if params == nil {
		params = &ListTrialsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTrials", params, optFns, addOperationListTrialsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTrialsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTrialsInput struct {

	// A filter that returns only trials created after the specified time.
	CreatedAfter *time.Time

	// A filter that returns only trials created before the specified time.
	CreatedBefore *time.Time

	// A filter that returns only trials that are part of the specified experiment.
	ExperimentName *string

	// The maximum number of trials to return in the response. The default value is 10.
	MaxResults *int32

	// If the previous call to ListTrials didn't return the full set of trials, the
	// call returns a token for getting the next set of trials.
	NextToken *string

	// The property used to sort results. The default value is CreationTime.
	SortBy types.SortTrialsBy

	// The sort order. The default value is Descending.
	SortOrder types.SortOrder

	// A filter that returns only trials that are associated with the specified trial
	// component.
	TrialComponentName *string
}

type ListTrialsOutput struct {

	// A token for getting the next set of trials, if there are any.
	NextToken *string

	// A list of the summaries of your trials.
	TrialSummaries []types.TrialSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTrialsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTrials{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTrials{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTrials(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListTrials(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListTrials",
	}
}
