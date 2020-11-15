// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves utilization statistics for one or more fleets. These statistics
// provide insight into how available hosting resources are currently being used.
// To get statistics on available hosting resources, see DescribeFleetCapacity. You
// can request utilization data for all fleets, or specify a list of one or more
// fleet IDs. When requesting multiple fleets, use the pagination parameters to
// retrieve results as a set of sequential pages. If successful, a FleetUtilization
// object is returned for each requested fleet ID, unless the fleet identifier is
// not found. Some API operations may limit the number of fleet IDs allowed in one
// request. If a request exceeds this limit, the request fails and the error
// message includes the maximum allowed. Learn more Setting up GameLift Fleets
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)GameLift
// Metrics for Fleets
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/monitoring-cloudwatch.html#gamelift-metrics-fleet)
// Related operations
//
// * CreateFleet
//
// * ListFleets
//
// * DeleteFleet
//
// * Describe
// fleets:
//
// * DescribeFleetAttributes
//
// * DescribeFleetCapacity
//
// *
// DescribeFleetPortSettings
//
// * DescribeFleetUtilization
//
// *
// DescribeRuntimeConfiguration
//
// * DescribeEC2InstanceLimits
//
// *
// DescribeFleetEvents
//
// * UpdateFleetAttributes
//
// * StartFleetActions or
// StopFleetActions
func (c *Client) DescribeFleetUtilization(ctx context.Context, params *DescribeFleetUtilizationInput, optFns ...func(*Options)) (*DescribeFleetUtilizationOutput, error) {
	if params == nil {
		params = &DescribeFleetUtilizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFleetUtilization", params, optFns, addOperationDescribeFleetUtilizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFleetUtilizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type DescribeFleetUtilizationInput struct {

	// A unique identifier for a fleet(s) to retrieve utilization data for. You can use
	// either the fleet ID or ARN value. To retrieve attributes for all current fleets,
	// do not include this parameter. If the list of fleet identifiers includes fleets
	// that don't currently exist, the request succeeds but no attributes for that
	// fleet are returned.
	FleetIds []string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages. This parameter is ignored when the
	// request specifies one or a list of fleet IDs.
	Limit *int32

	// Token that indicates the start of the next sequential page of results. Use the
	// token that is returned with a previous call to this operation. To start at the
	// beginning of the result set, do not specify a value. This parameter is ignored
	// when the request specifies one or a list of fleet IDs.
	NextToken *string
}

// Represents the returned data in response to a request operation.
type DescribeFleetUtilizationOutput struct {

	// A collection of objects containing utilization information for each requested
	// fleet ID.
	FleetUtilization []types.FleetUtilization

	// Token that indicates where to resume retrieving results on the next call to this
	// operation. If no token is returned, these results represent the end of the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeFleetUtilizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeFleetUtilization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeFleetUtilization{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFleetUtilization(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeFleetUtilization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeFleetUtilization",
	}
}
