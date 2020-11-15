// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Assigns a tape to a tape pool for archiving. The tape assigned to a pool is
// archived in the S3 storage class that is associated with the pool. When you use
// your backup application to eject the tape, the tape is archived directly into
// the S3 storage class (S3 Glacier or S3 Glacier Deep Archive) that corresponds to
// the pool. Valid Values: GLACIER | DEEP_ARCHIVE
func (c *Client) AssignTapePool(ctx context.Context, params *AssignTapePoolInput, optFns ...func(*Options)) (*AssignTapePoolOutput, error) {
	if params == nil {
		params = &AssignTapePoolInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssignTapePool", params, optFns, addOperationAssignTapePoolMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssignTapePoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssignTapePoolInput struct {

	// The ID of the pool that you want to add your tape to for archiving. The tape in
	// this pool is archived in the S3 storage class that is associated with the pool.
	// When you use your backup application to eject the tape, the tape is archived
	// directly into the storage class (S3 Glacier or S3 Glacier Deep Archive) that
	// corresponds to the pool. Valid Values: GLACIER | DEEP_ARCHIVE
	//
	// This member is required.
	PoolId *string

	// The unique Amazon Resource Name (ARN) of the virtual tape that you want to add
	// to the tape pool.
	//
	// This member is required.
	TapeARN *string

	// Set permissions to bypass governance retention. If the lock type of the archived
	// tape is Governance, the tape's archived age is not older than
	// RetentionLockInDays, and the user does not already have
	// BypassGovernanceRetention, setting this to TRUE enables the user to bypass the
	// retention lock. This parameter is set to true by default for calls from the
	// console. Valid values: TRUE | FALSE
	BypassGovernanceRetention bool
}

type AssignTapePoolOutput struct {

	// The unique Amazon Resource Names (ARN) of the virtual tape that was added to the
	// tape pool.
	TapeARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAssignTapePoolMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssignTapePool{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssignTapePool{}, middleware.After)
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
	if err = addOpAssignTapePoolValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssignTapePool(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssignTapePool(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "AssignTapePool",
	}
}
