// Code generated by smithy-go-codegen DO NOT EDIT.

package ssooidc

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

func (c *Client) StartDeviceAuthorization(ctx context.Context, params *StartDeviceAuthorizationInput, optFns ...func(*Options)) (*StartDeviceAuthorizationOutput, error) {
	if params == nil {
		params = &StartDeviceAuthorizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartDeviceAuthorization", params, optFns, addOperationStartDeviceAuthorizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartDeviceAuthorizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartDeviceAuthorizationInput struct {
	ClientId *string

	ClientSecret *string

	StartUrl *string
}

type StartDeviceAuthorizationOutput struct {
	DeviceCode *string

	ExpiresIn int32

	Interval int32

	UserCode *string

	VerificationUri *string

	VerificationUriComplete *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartDeviceAuthorizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartDeviceAuthorization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartDeviceAuthorization{}, middleware.After)
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
	if err = addOpStartDeviceAuthorizationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartDeviceAuthorization(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartDeviceAuthorization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "awsssooidc",
		OperationName: "StartDeviceAuthorization",
	}
}
