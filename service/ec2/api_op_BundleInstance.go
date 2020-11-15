// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Bundles an Amazon instance store-backed Windows instance. During bundling, only
// the root device volume (C:\) is bundled. Data on other instance store volumes is
// not preserved. This action is not applicable for Linux/Unix instances or Windows
// instances that are backed by Amazon EBS.
func (c *Client) BundleInstance(ctx context.Context, params *BundleInstanceInput, optFns ...func(*Options)) (*BundleInstanceOutput, error) {
	if params == nil {
		params = &BundleInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BundleInstance", params, optFns, addOperationBundleInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BundleInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for BundleInstance.
type BundleInstanceInput struct {

	// The ID of the instance to bundle. Type: String Default: None Required: Yes
	//
	// This member is required.
	InstanceId *string

	// The bucket in which to store the AMI. You can specify a bucket that you already
	// own or a new bucket that Amazon EC2 creates on your behalf. If you specify a
	// bucket that belongs to someone else, Amazon EC2 returns an error.
	//
	// This member is required.
	Storage *types.Storage

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool
}

// Contains the output of BundleInstance.
type BundleInstanceOutput struct {

	// Information about the bundle task.
	BundleTask *types.BundleTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBundleInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpBundleInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpBundleInstance{}, middleware.After)
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
	if err = addOpBundleInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBundleInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBundleInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "BundleInstance",
	}
}
