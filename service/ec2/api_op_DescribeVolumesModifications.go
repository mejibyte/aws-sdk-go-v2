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

// Describes the most recent volume modification request for the specified EBS
// volumes. If a volume has never been modified, some information in the output
// will be null. If a volume has been modified more than once, the output includes
// only the most recent modification request. You can also use CloudWatch Events to
// check the status of a modification to an EBS volume. For information about
// CloudWatch Events, see the Amazon CloudWatch Events User Guide
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/). For more
// information, see Monitoring volume modifications
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-expand-volume.html#monitoring_mods)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) DescribeVolumesModifications(ctx context.Context, params *DescribeVolumesModificationsInput, optFns ...func(*Options)) (*DescribeVolumesModificationsOutput, error) {
	if params == nil {
		params = &DescribeVolumesModificationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVolumesModifications", params, optFns, addOperationDescribeVolumesModificationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVolumesModificationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVolumesModificationsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// The filters.
	//
	// * modification-state - The current modification state (modifying |
	// optimizing | completed | failed).
	//
	// * original-iops - The original IOPS rate of
	// the volume.
	//
	// * original-size - The original size of the volume, in GiB.
	//
	// *
	// original-volume-type - The original volume type of the volume (standard | io1 |
	// io2 | gp2 | sc1 | st1).
	//
	// * originalMultiAttachEnabled - Indicates whether
	// Multi-Attach support was enabled (true | false).
	//
	// * start-time - The
	// modification start time.
	//
	// * target-iops - The target IOPS rate of the volume.
	//
	// *
	// target-size - The target size of the volume, in GiB.
	//
	// * target-volume-type - The
	// target volume type of the volume (standard | io1 | io2 | gp2 | sc1 | st1).
	//
	// *
	// targetMultiAttachEnabled - Indicates whether Multi-Attach support is to be
	// enabled (true | false).
	//
	// * volume-id - The ID of the volume.
	Filters []types.Filter

	// The maximum number of results (up to a limit of 500) to be returned in a
	// paginated request.
	MaxResults int32

	// The nextToken value returned by a previous paginated request.
	NextToken *string

	// The IDs of the volumes.
	VolumeIds []string
}

type DescribeVolumesModificationsOutput struct {

	// Token for pagination, null if there are no more results
	NextToken *string

	// Information about the volume modifications.
	VolumesModifications []types.VolumeModification

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeVolumesModificationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeVolumesModifications{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVolumesModifications{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVolumesModifications(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeVolumesModifications(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVolumesModifications",
	}
}
