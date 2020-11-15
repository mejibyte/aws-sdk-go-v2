// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Information about the thing registration tasks.
func (c *Client) ListThingRegistrationTaskReports(ctx context.Context, params *ListThingRegistrationTaskReportsInput, optFns ...func(*Options)) (*ListThingRegistrationTaskReportsOutput, error) {
	if params == nil {
		params = &ListThingRegistrationTaskReportsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListThingRegistrationTaskReports", params, optFns, addOperationListThingRegistrationTaskReportsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListThingRegistrationTaskReportsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListThingRegistrationTaskReportsInput struct {

	// The type of task report.
	//
	// This member is required.
	ReportType types.ReportType

	// The id of the task.
	//
	// This member is required.
	TaskId *string

	// The maximum number of results to return per request.
	MaxResults *int32

	// The token to retrieve the next set of results.
	NextToken *string
}

type ListThingRegistrationTaskReportsOutput struct {

	// The token used to get the next set of results, or null if there are no
	// additional results.
	NextToken *string

	// The type of task report.
	ReportType types.ReportType

	// Links to the task resources.
	ResourceLinks []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListThingRegistrationTaskReportsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListThingRegistrationTaskReports{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListThingRegistrationTaskReports{}, middleware.After)
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
	if err = addOpListThingRegistrationTaskReportsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListThingRegistrationTaskReports(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListThingRegistrationTaskReports(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListThingRegistrationTaskReports",
	}
}
