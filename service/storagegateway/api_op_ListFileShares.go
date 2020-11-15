// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of the file shares for a specific file gateway, or the list of file
// shares that belong to the calling user account. This operation is only supported
// for file gateways.
func (c *Client) ListFileShares(ctx context.Context, params *ListFileSharesInput, optFns ...func(*Options)) (*ListFileSharesOutput, error) {
	if params == nil {
		params = &ListFileSharesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFileShares", params, optFns, addOperationListFileSharesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFileSharesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// ListFileShareInput
type ListFileSharesInput struct {

	// The Amazon Resource Name (ARN) of the gateway whose file shares you want to
	// list. If this field is not present, all file shares under your account are
	// listed.
	GatewayARN *string

	// The maximum number of file shares to return in the response. The value must be
	// an integer with a value greater than zero. Optional.
	Limit *int32

	// Opaque pagination token returned from a previous ListFileShares operation. If
	// present, Marker specifies where to continue the list from after a previous call
	// to ListFileShares. Optional.
	Marker *string
}

// ListFileShareOutput
type ListFileSharesOutput struct {

	// An array of information about the file gateway's file shares.
	FileShareInfoList []types.FileShareInfo

	// If the request includes Marker, the response returns that value in this field.
	Marker *string

	// If a value is present, there are more file shares to return. In a subsequent
	// request, use NextMarker as the value for Marker to retrieve the next set of file
	// shares.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListFileSharesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListFileShares{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFileShares{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFileShares(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListFileShares(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "ListFileShares",
	}
}
