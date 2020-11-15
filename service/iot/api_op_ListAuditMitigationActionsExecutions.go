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

// Gets the status of audit mitigation action tasks that were executed.
func (c *Client) ListAuditMitigationActionsExecutions(ctx context.Context, params *ListAuditMitigationActionsExecutionsInput, optFns ...func(*Options)) (*ListAuditMitigationActionsExecutionsOutput, error) {
	if params == nil {
		params = &ListAuditMitigationActionsExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAuditMitigationActionsExecutions", params, optFns, addOperationListAuditMitigationActionsExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAuditMitigationActionsExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAuditMitigationActionsExecutionsInput struct {

	// Specify this filter to limit results to those that were applied to a specific
	// audit finding.
	//
	// This member is required.
	FindingId *string

	// Specify this filter to limit results to actions for a specific audit mitigation
	// actions task.
	//
	// This member is required.
	TaskId *string

	// Specify this filter to limit results to those with a specific status.
	ActionStatus types.AuditMitigationActionsExecutionStatus

	// The maximum number of results to return at one time. The default is 25.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string
}

type ListAuditMitigationActionsExecutionsOutput struct {

	// A set of task execution results based on the input parameters. Details include
	// the mitigation action applied, start time, and task status.
	ActionsExecutions []types.AuditMitigationActionExecutionMetadata

	// The token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListAuditMitigationActionsExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAuditMitigationActionsExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAuditMitigationActionsExecutions{}, middleware.After)
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
	if err = addOpListAuditMitigationActionsExecutionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAuditMitigationActionsExecutions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListAuditMitigationActionsExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListAuditMitigationActionsExecutions",
	}
}
