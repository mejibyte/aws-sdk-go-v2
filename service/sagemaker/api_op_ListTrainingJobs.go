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

// Lists training jobs.
func (c *Client) ListTrainingJobs(ctx context.Context, params *ListTrainingJobsInput, optFns ...func(*Options)) (*ListTrainingJobsOutput, error) {
	if params == nil {
		params = &ListTrainingJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTrainingJobs", params, optFns, addOperationListTrainingJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTrainingJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTrainingJobsInput struct {

	// A filter that returns only training jobs created after the specified time
	// (timestamp).
	CreationTimeAfter *time.Time

	// A filter that returns only training jobs created before the specified time
	// (timestamp).
	CreationTimeBefore *time.Time

	// A filter that returns only training jobs modified after the specified time
	// (timestamp).
	LastModifiedTimeAfter *time.Time

	// A filter that returns only training jobs modified before the specified time
	// (timestamp).
	LastModifiedTimeBefore *time.Time

	// The maximum number of training jobs to return in the response.
	MaxResults *int32

	// A string in the training job name. This filter returns only training jobs whose
	// name contains the specified string.
	NameContains *string

	// If the result of the previous ListTrainingJobs request was truncated, the
	// response includes a NextToken. To retrieve the next set of training jobs, use
	// the token in the next request.
	NextToken *string

	// The field to sort results by. The default is CreationTime.
	SortBy types.SortBy

	// The sort order for results. The default is Ascending.
	SortOrder types.SortOrder

	// A filter that retrieves only training jobs with a specific status.
	StatusEquals types.TrainingJobStatus
}

type ListTrainingJobsOutput struct {

	// An array of TrainingJobSummary objects, each listing a training job.
	//
	// This member is required.
	TrainingJobSummaries []types.TrainingJobSummary

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of training jobs, use it in the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTrainingJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTrainingJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTrainingJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTrainingJobs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListTrainingJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListTrainingJobs",
	}
}
