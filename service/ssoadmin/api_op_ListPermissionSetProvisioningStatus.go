// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the status of the permission set provisioning requests for a specified SSO
// instance.
func (c *Client) ListPermissionSetProvisioningStatus(ctx context.Context, params *ListPermissionSetProvisioningStatusInput, optFns ...func(*Options)) (*ListPermissionSetProvisioningStatusOutput, error) {
	if params == nil {
		params = &ListPermissionSetProvisioningStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPermissionSetProvisioningStatus", params, optFns, addOperationListPermissionSetProvisioningStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPermissionSetProvisioningStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPermissionSetProvisioningStatusInput struct {

	// The ARN of the SSO instance under which the operation will be executed. For more
	// information about ARNs, see Amazon Resource Names (ARNs) and AWS Service
	// Namespaces in the AWS General Reference.
	//
	// This member is required.
	InstanceArn *string

	// Filters results based on the passed attribute value.
	Filter *types.OperationStatusFilter

	// The maximum number of results to display for the assignment.
	MaxResults *int32

	// The pagination token for the list API. Initially the value is null. Use the
	// output of previous API calls to make subsequent calls.
	NextToken *string
}

type ListPermissionSetProvisioningStatusOutput struct {

	// The pagination token for the list API. Initially the value is null. Use the
	// output of previous API calls to make subsequent calls.
	NextToken *string

	// The status object for the permission set provisioning operation.
	PermissionSetsProvisioningStatus []types.PermissionSetProvisioningStatusMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListPermissionSetProvisioningStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPermissionSetProvisioningStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPermissionSetProvisioningStatus{}, middleware.After)
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
	if err = addOpListPermissionSetProvisioningStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPermissionSetProvisioningStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListPermissionSetProvisioningStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sso",
		OperationName: "ListPermissionSetProvisioningStatus",
	}
}
