// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists your resource data sync configurations. Includes information about the
// last time a sync attempted to start, the last sync status, and the last time a
// sync successfully completed. The number of sync configurations might be too
// large to return using a single call to ListResourceDataSync. You can limit the
// number of sync configurations returned by using the MaxResults parameter. To
// determine whether there are more sync configurations to list, check the value of
// NextToken in the output. If there are more sync configurations to list, you can
// request them by specifying the NextToken returned in the call to the parameter
// of a subsequent call.
func (c *Client) ListResourceDataSync(ctx context.Context, params *ListResourceDataSyncInput, optFns ...func(*Options)) (*ListResourceDataSyncOutput, error) {
	if params == nil {
		params = &ListResourceDataSyncInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResourceDataSync", params, optFns, addOperationListResourceDataSyncMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResourceDataSyncOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResourceDataSyncInput struct {

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults int32

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string

	// View a list of resource data syncs according to the sync type. Specify
	// SyncToDestination to view resource data syncs that synchronize data to an Amazon
	// S3 bucket. Specify SyncFromSource to view resource data syncs from AWS
	// Organizations or from multiple AWS Regions.
	SyncType *string
}

type ListResourceDataSyncOutput struct {

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// A list of your current Resource Data Sync configurations and their statuses.
	ResourceDataSyncItems []types.ResourceDataSyncItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListResourceDataSyncMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListResourceDataSync{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResourceDataSync{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResourceDataSync(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListResourceDataSync(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "ListResourceDataSync",
	}
}
