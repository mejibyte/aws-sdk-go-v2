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

// Modifies the Availability Zone, instance count, instance type, or network
// platform (EC2-Classic or EC2-VPC) of your Reserved Instances. The Reserved
// Instances to be modified must be identical, except for Availability Zone,
// network platform, and instance type. For more information, see Modifying
// Reserved Instances
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ri-modifying.html) in the
// Amazon Elastic Compute Cloud User Guide.
func (c *Client) ModifyReservedInstances(ctx context.Context, params *ModifyReservedInstancesInput, optFns ...func(*Options)) (*ModifyReservedInstancesOutput, error) {
	if params == nil {
		params = &ModifyReservedInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyReservedInstances", params, optFns, addOperationModifyReservedInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyReservedInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for ModifyReservedInstances.
type ModifyReservedInstancesInput struct {

	// The IDs of the Reserved Instances to modify.
	//
	// This member is required.
	ReservedInstancesIds []string

	// The configuration settings for the Reserved Instances to modify.
	//
	// This member is required.
	TargetConfigurations []types.ReservedInstancesConfiguration

	// A unique, case-sensitive token you provide to ensure idempotency of your
	// modification request. For more information, see Ensuring Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string
}

// Contains the output of ModifyReservedInstances.
type ModifyReservedInstancesOutput struct {

	// The ID for the modification.
	ReservedInstancesModificationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyReservedInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyReservedInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyReservedInstances{}, middleware.After)
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
	if err = addOpModifyReservedInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyReservedInstances(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyReservedInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyReservedInstances",
	}
}
