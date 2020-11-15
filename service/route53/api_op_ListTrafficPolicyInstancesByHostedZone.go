// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about the traffic policy instances that you created in a
// specified hosted zone. After you submit a CreateTrafficPolicyInstance or an
// UpdateTrafficPolicyInstance request, there's a brief delay while Amazon Route 53
// creates the resource record sets that are specified in the traffic policy
// definition. For more information, see the State response element. Route 53
// returns a maximum of 100 items in each response. If you have a lot of traffic
// policy instances, you can use the MaxItems parameter to list them in groups of
// up to 100.
func (c *Client) ListTrafficPolicyInstancesByHostedZone(ctx context.Context, params *ListTrafficPolicyInstancesByHostedZoneInput, optFns ...func(*Options)) (*ListTrafficPolicyInstancesByHostedZoneOutput, error) {
	if params == nil {
		params = &ListTrafficPolicyInstancesByHostedZoneInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTrafficPolicyInstancesByHostedZone", params, optFns, addOperationListTrafficPolicyInstancesByHostedZoneMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTrafficPolicyInstancesByHostedZoneOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request for the traffic policy instances that you created in a specified
// hosted zone.
type ListTrafficPolicyInstancesByHostedZoneInput struct {

	// The ID of the hosted zone that you want to list traffic policy instances for.
	//
	// This member is required.
	HostedZoneId *string

	// The maximum number of traffic policy instances to be included in the response
	// body for this request. If you have more than MaxItems traffic policy instances,
	// the value of the IsTruncated element in the response is true, and the values of
	// HostedZoneIdMarker, TrafficPolicyInstanceNameMarker, and
	// TrafficPolicyInstanceTypeMarker represent the first traffic policy instance that
	// Amazon Route 53 will return if you submit another request.
	MaxItems *string

	// If the value of IsTruncated in the previous response is true, you have more
	// traffic policy instances. To get more traffic policy instances, submit another
	// ListTrafficPolicyInstances request. For the value of trafficpolicyinstancename,
	// specify the value of TrafficPolicyInstanceNameMarker from the previous response,
	// which is the name of the first traffic policy instance in the next group of
	// traffic policy instances. If the value of IsTruncated in the previous response
	// was false, there are no more traffic policy instances to get.
	TrafficPolicyInstanceNameMarker *string

	// If the value of IsTruncated in the previous response is true, you have more
	// traffic policy instances. To get more traffic policy instances, submit another
	// ListTrafficPolicyInstances request. For the value of trafficpolicyinstancetype,
	// specify the value of TrafficPolicyInstanceTypeMarker from the previous response,
	// which is the type of the first traffic policy instance in the next group of
	// traffic policy instances. If the value of IsTruncated in the previous response
	// was false, there are no more traffic policy instances to get.
	TrafficPolicyInstanceTypeMarker types.RRType
}

// A complex type that contains the response information for the request.
type ListTrafficPolicyInstancesByHostedZoneOutput struct {

	// A flag that indicates whether there are more traffic policy instances to be
	// listed. If the response was truncated, you can get the next group of traffic
	// policy instances by submitting another ListTrafficPolicyInstancesByHostedZone
	// request and specifying the values of HostedZoneIdMarker,
	// TrafficPolicyInstanceNameMarker, and TrafficPolicyInstanceTypeMarker in the
	// corresponding request parameters.
	//
	// This member is required.
	IsTruncated bool

	// The value that you specified for the MaxItems parameter in the
	// ListTrafficPolicyInstancesByHostedZone request that produced the current
	// response.
	//
	// This member is required.
	MaxItems *string

	// A list that contains one TrafficPolicyInstance element for each traffic policy
	// instance that matches the elements in the request.
	//
	// This member is required.
	TrafficPolicyInstances []types.TrafficPolicyInstance

	// If IsTruncated is true, TrafficPolicyInstanceNameMarker is the name of the first
	// traffic policy instance in the next group of traffic policy instances.
	TrafficPolicyInstanceNameMarker *string

	// If IsTruncated is true, TrafficPolicyInstanceTypeMarker is the DNS type of the
	// resource record sets that are associated with the first traffic policy instance
	// in the next group of traffic policy instances.
	TrafficPolicyInstanceTypeMarker types.RRType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTrafficPolicyInstancesByHostedZoneMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListTrafficPolicyInstancesByHostedZone{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListTrafficPolicyInstancesByHostedZone{}, middleware.After)
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
	if err = addOpListTrafficPolicyInstancesByHostedZoneValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTrafficPolicyInstancesByHostedZone(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addSanitizeURLMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListTrafficPolicyInstancesByHostedZone(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "ListTrafficPolicyInstancesByHostedZone",
	}
}
