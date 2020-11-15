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

// Lists endpoints.
func (c *Client) ListEndpoints(ctx context.Context, params *ListEndpointsInput, optFns ...func(*Options)) (*ListEndpointsOutput, error) {
	if params == nil {
		params = &ListEndpointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEndpoints", params, optFns, addOperationListEndpointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEndpointsInput struct {

	// A filter that returns only endpoints with a creation time greater than or equal
	// to the specified time (timestamp).
	CreationTimeAfter *time.Time

	// A filter that returns only endpoints that were created before the specified time
	// (timestamp).
	CreationTimeBefore *time.Time

	// A filter that returns only endpoints that were modified after the specified
	// timestamp.
	LastModifiedTimeAfter *time.Time

	// A filter that returns only endpoints that were modified before the specified
	// timestamp.
	LastModifiedTimeBefore *time.Time

	// The maximum number of endpoints to return in the response.
	MaxResults *int32

	// A string in endpoint names. This filter returns only endpoints whose name
	// contains the specified string.
	NameContains *string

	// If the result of a ListEndpoints request was truncated, the response includes a
	// NextToken. To retrieve the next set of endpoints, use the token in the next
	// request.
	NextToken *string

	// Sorts the list of results. The default is CreationTime.
	SortBy types.EndpointSortKey

	// The sort order for results. The default is Descending.
	SortOrder types.OrderKey

	// A filter that returns only endpoints with the specified status.
	StatusEquals types.EndpointStatus
}

type ListEndpointsOutput struct {

	// An array or endpoint objects.
	//
	// This member is required.
	Endpoints []types.EndpointSummary

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of training jobs, use it in the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListEndpointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEndpoints{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEndpoints(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListEndpoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListEndpoints",
	}
}
