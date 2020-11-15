// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation returns information about a job you previously initiated,
// including the job initiation date, the user who initiated the job, the job
// status code/message and the Amazon SNS topic to notify after Amazon S3 Glacier
// (Glacier) completes the job. For more information about initiating a job, see
// InitiateJob. This operation enables you to check the status of your job.
// However, it is strongly recommended that you set up an Amazon SNS topic and
// specify it in your initiate job request so that Glacier can notify the topic
// after it completes the job. A job ID will not expire for at least 24 hours after
// Glacier completes the job. An AWS account has full permission to perform all
// operations (actions). However, AWS Identity and Access Management (IAM) users
// don't have any permissions by default. You must grant them explicit permission
// to perform specific actions. For more information, see Access Control Using AWS
// Identity and Access Management (IAM)
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html).
// For more information about using this operation, see the documentation for the
// underlying REST API Describe Job
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-describe-job-get.html)
// in the Amazon Glacier Developer Guide.
func (c *Client) DescribeJob(ctx context.Context, params *DescribeJobInput, optFns ...func(*Options)) (*DescribeJobOutput, error) {
	if params == nil {
		params = &DescribeJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeJob", params, optFns, addOperationDescribeJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options for retrieving a job description.
type DescribeJobInput struct {

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen), in
	// which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// The ID of the job to describe.
	//
	// This member is required.
	JobId *string

	// The name of the vault.
	//
	// This member is required.
	VaultName *string
}

// Contains the description of an Amazon S3 Glacier job.
type DescribeJobOutput struct {

	// The job type. This value is either ArchiveRetrieval, InventoryRetrieval, or
	// Select.
	Action types.ActionCode

	// The archive ID requested for a select job or archive retrieval. Otherwise, this
	// field is null.
	ArchiveId *string

	// The SHA256 tree hash of the entire archive for an archive retrieval. For
	// inventory retrieval or select jobs, this field is null.
	ArchiveSHA256TreeHash *string

	// For an archive retrieval job, this value is the size in bytes of the archive
	// being requested for download. For an inventory retrieval or select job, this
	// value is null.
	ArchiveSizeInBytes *int64

	// The job status. When a job is completed, you get the job's output using Get Job
	// Output (GET output).
	Completed bool

	// The UTC time that the job request completed. While the job is in progress, the
	// value is null.
	CompletionDate *string

	// The UTC date when the job was created. This value is a string representation of
	// ISO 8601 date format, for example "2012-03-20T17:03:43.221Z".
	CreationDate *string

	// Parameters used for range inventory retrieval.
	InventoryRetrievalParameters *types.InventoryRetrievalJobDescription

	// For an inventory retrieval job, this value is the size in bytes of the inventory
	// requested for download. For an archive retrieval or select job, this value is
	// null.
	InventorySizeInBytes *int64

	// The job description provided when initiating the job.
	JobDescription *string

	// An opaque string that identifies an Amazon S3 Glacier job.
	JobId *string

	// Contains the job output location.
	JobOutputPath *string

	// Contains the location where the data from the select job is stored.
	OutputLocation *types.OutputLocation

	// The retrieved byte range for archive retrieval jobs in the form
	// StartByteValue-EndByteValue. If no range was specified in the archive retrieval,
	// then the whole archive is retrieved. In this case, StartByteValue equals 0 and
	// EndByteValue equals the size of the archive minus 1. For inventory retrieval or
	// select jobs, this field is null.
	RetrievalByteRange *string

	// For an archive retrieval job, this value is the checksum of the archive.
	// Otherwise, this value is null. The SHA256 tree hash value for the requested
	// range of an archive. If the InitiateJob request for an archive specified a
	// tree-hash aligned range, then this field returns a value. If the whole archive
	// is retrieved, this value is the same as the ArchiveSHA256TreeHash value. This
	// field is null for the following:
	//
	// * Archive retrieval jobs that specify a range
	// that is not tree-hash aligned
	//
	// * Archival jobs that specify a range that is
	// equal to the whole archive, when the job status is InProgress
	//
	// * Inventory
	// jobs
	//
	// * Select jobs
	SHA256TreeHash *string

	// An Amazon SNS topic that receives notification.
	SNSTopic *string

	// Contains the parameters used for a select.
	SelectParameters *types.SelectParameters

	// The status code can be InProgress, Succeeded, or Failed, and indicates the
	// status of the job.
	StatusCode types.StatusCode

	// A friendly message that describes the job status.
	StatusMessage *string

	// The tier to use for a select or an archive retrieval. Valid values are
	// Expedited, Standard, or Bulk. Standard is the default.
	Tier *string

	// The Amazon Resource Name (ARN) of the vault from which an archive retrieval was
	// requested.
	VaultARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeJob{}, middleware.After)
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
	if err = addOpDescribeJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeJob(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddTreeHashMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion); err != nil {
		return err
	}
	if err = glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "DescribeJob",
	}
}
