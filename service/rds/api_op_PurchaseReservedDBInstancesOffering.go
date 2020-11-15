// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Purchases a reserved DB instance offering.
func (c *Client) PurchaseReservedDBInstancesOffering(ctx context.Context, params *PurchaseReservedDBInstancesOfferingInput, optFns ...func(*Options)) (*PurchaseReservedDBInstancesOfferingOutput, error) {
	if params == nil {
		params = &PurchaseReservedDBInstancesOfferingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PurchaseReservedDBInstancesOffering", params, optFns, addOperationPurchaseReservedDBInstancesOfferingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PurchaseReservedDBInstancesOfferingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type PurchaseReservedDBInstancesOfferingInput struct {

	// The ID of the Reserved DB instance offering to purchase. Example:
	// 438012d3-4052-4cc7-b2e3-8d3372e0e706
	//
	// This member is required.
	ReservedDBInstancesOfferingId *string

	// The number of instances to reserve. Default: 1
	DBInstanceCount *int32

	// Customer-specified identifier to track this reservation. Example:
	// myreservationID
	ReservedDBInstanceId *string

	// A list of tags. For more information, see Tagging Amazon RDS Resources
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in
	// the Amazon RDS User Guide.
	Tags []types.Tag
}

type PurchaseReservedDBInstancesOfferingOutput struct {

	// This data type is used as a response element in the DescribeReservedDBInstances
	// and PurchaseReservedDBInstancesOffering actions.
	ReservedDBInstance *types.ReservedDBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPurchaseReservedDBInstancesOfferingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPurchaseReservedDBInstancesOffering{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPurchaseReservedDBInstancesOffering{}, middleware.After)
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
	if err = addOpPurchaseReservedDBInstancesOfferingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPurchaseReservedDBInstancesOffering(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPurchaseReservedDBInstancesOffering(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "PurchaseReservedDBInstancesOffering",
	}
}
