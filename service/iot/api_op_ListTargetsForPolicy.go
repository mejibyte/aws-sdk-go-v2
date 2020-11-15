// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List targets for the specified policy.
func (c *Client) ListTargetsForPolicy(ctx context.Context, params *ListTargetsForPolicyInput, optFns ...func(*Options)) (*ListTargetsForPolicyOutput, error) {
	if params == nil {
		params = &ListTargetsForPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTargetsForPolicy", params, optFns, addOperationListTargetsForPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTargetsForPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTargetsForPolicyInput struct {

	// The policy name.
	//
	// This member is required.
	PolicyName *string

	// A marker used to get the next set of results.
	Marker *string

	// The maximum number of results to return at one time.
	PageSize *int32
}

type ListTargetsForPolicyOutput struct {

	// A marker used to get the next set of results.
	NextMarker *string

	// The policy targets.
	Targets []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTargetsForPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTargetsForPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTargetsForPolicy{}, middleware.After)
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
	if err = addOpListTargetsForPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTargetsForPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListTargetsForPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListTargetsForPolicy",
	}
}
