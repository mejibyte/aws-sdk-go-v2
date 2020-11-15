// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds the specified Availability Zones to the set of Availability Zones for the
// specified load balancer in EC2-Classic or a default VPC. For load balancers in a
// non-default VPC, use AttachLoadBalancerToSubnets. The load balancer evenly
// distributes requests across all its registered Availability Zones that contain
// instances. For more information, see Add or Remove Availability Zones
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-disable-az.html)
// in the Classic Load Balancers Guide.
func (c *Client) EnableAvailabilityZonesForLoadBalancer(ctx context.Context, params *EnableAvailabilityZonesForLoadBalancerInput, optFns ...func(*Options)) (*EnableAvailabilityZonesForLoadBalancerOutput, error) {
	if params == nil {
		params = &EnableAvailabilityZonesForLoadBalancerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableAvailabilityZonesForLoadBalancer", params, optFns, addOperationEnableAvailabilityZonesForLoadBalancerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableAvailabilityZonesForLoadBalancerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for EnableAvailabilityZonesForLoadBalancer.
type EnableAvailabilityZonesForLoadBalancerInput struct {

	// The Availability Zones. These must be in the same region as the load balancer.
	//
	// This member is required.
	AvailabilityZones []string

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string
}

// Contains the output of EnableAvailabilityZonesForLoadBalancer.
type EnableAvailabilityZonesForLoadBalancerOutput struct {

	// The updated list of Availability Zones for the load balancer.
	AvailabilityZones []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationEnableAvailabilityZonesForLoadBalancerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpEnableAvailabilityZonesForLoadBalancer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpEnableAvailabilityZonesForLoadBalancer{}, middleware.After)
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
	if err = addOpEnableAvailabilityZonesForLoadBalancerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableAvailabilityZonesForLoadBalancer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnableAvailabilityZonesForLoadBalancer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "EnableAvailabilityZonesForLoadBalancer",
	}
}
