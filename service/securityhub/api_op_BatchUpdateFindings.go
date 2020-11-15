// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Used by Security Hub customers to update information about their investigation
// into a finding. Requested by master accounts or member accounts. Master accounts
// can update findings for their account and their member accounts. Member accounts
// can update findings for their account. Updates from BatchUpdateFindings do not
// affect the value of UpdatedAt for a finding. Master and member accounts can use
// BatchUpdateFindings to update the following finding fields and objects.
//
// *
// Confidence
//
// * Criticality
//
// * Note
//
// * RelatedFindings
//
// * Severity
//
// * Types
//
// *
// UserDefinedFields
//
// * VerificationState
//
// * Workflow
//
// You can configure IAM
// policies to restrict access to fields and field values. For example, you might
// not want member accounts to be able to suppress findings or change the finding
// severity. See Configuring access to BatchUpdateFindings
// (https://docs.aws.amazon.com/securityhub/latest/userguide/finding-update-batchupdatefindings.html#batchupdatefindings-configure-access)
// in the AWS Security Hub User Guide.
func (c *Client) BatchUpdateFindings(ctx context.Context, params *BatchUpdateFindingsInput, optFns ...func(*Options)) (*BatchUpdateFindingsOutput, error) {
	if params == nil {
		params = &BatchUpdateFindingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchUpdateFindings", params, optFns, addOperationBatchUpdateFindingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchUpdateFindingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchUpdateFindingsInput struct {

	// The list of findings to update. BatchUpdateFindings can be used to update up to
	// 100 findings at a time. For each finding, the list provides the finding
	// identifier and the ARN of the finding provider.
	//
	// This member is required.
	FindingIdentifiers []types.AwsSecurityFindingIdentifier

	// The updated value for the finding confidence. Confidence is defined as the
	// likelihood that a finding accurately identifies the behavior or issue that it
	// was intended to identify. Confidence is scored on a 0-100 basis using a ratio
	// scale, where 0 means zero percent confidence and 100 means 100 percent
	// confidence.
	Confidence int32

	// The updated value for the level of importance assigned to the resources
	// associated with the findings. A score of 0 means that the underlying resources
	// have no criticality, and a score of 100 is reserved for the most critical
	// resources.
	Criticality int32

	// The updated note.
	Note *types.NoteUpdate

	// A list of findings that are related to the updated findings.
	RelatedFindings []types.RelatedFinding

	// Used to update the finding severity.
	Severity *types.SeverityUpdate

	// One or more finding types in the format of namespace/category/classifier that
	// classify a finding. Valid namespace values are as follows.
	//
	// * Software and
	// Configuration Checks
	//
	// * TTPs
	//
	// * Effects
	//
	// * Unusual Behaviors
	//
	// * Sensitive Data
	// Identifications
	Types []string

	// A list of name/value string pairs associated with the finding. These are custom,
	// user-defined fields added to a finding.
	UserDefinedFields map[string]string

	// Indicates the veracity of a finding. The available values for VerificationState
	// are as follows.
	//
	// * UNKNOWN – The default disposition of a security finding
	//
	// *
	// TRUE_POSITIVE – The security finding is confirmed
	//
	// * FALSE_POSITIVE – The
	// security finding was determined to be a false alarm
	//
	// * BENIGN_POSITIVE – A
	// special case of TRUE_POSITIVE where the finding doesn't pose any threat, is
	// expected, or both
	VerificationState types.VerificationState

	// Used to update the workflow status of a finding. The workflow status indicates
	// the progress of the investigation into the finding.
	Workflow *types.WorkflowUpdate
}

type BatchUpdateFindingsOutput struct {

	// The list of findings that were updated successfully.
	//
	// This member is required.
	ProcessedFindings []types.AwsSecurityFindingIdentifier

	// The list of findings that were not updated.
	//
	// This member is required.
	UnprocessedFindings []types.BatchUpdateFindingsUnprocessedFinding

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchUpdateFindingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchUpdateFindings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchUpdateFindings{}, middleware.After)
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
	if err = addOpBatchUpdateFindingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchUpdateFindings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchUpdateFindings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "BatchUpdateFindings",
	}
}
