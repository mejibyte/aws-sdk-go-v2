// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteTransitGatewayMulticastDomainInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the transit gateway multicast domain.
	//
	// TransitGatewayMulticastDomainId is a required field
	TransitGatewayMulticastDomainId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTransitGatewayMulticastDomainInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTransitGatewayMulticastDomainInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTransitGatewayMulticastDomainInput"}

	if s.TransitGatewayMulticastDomainId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransitGatewayMulticastDomainId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteTransitGatewayMulticastDomainOutput struct {
	_ struct{} `type:"structure"`

	// Information about the deleted transit gateway multicast domain.
	TransitGatewayMulticastDomain *TransitGatewayMulticastDomain `locationName:"transitGatewayMulticastDomain" type:"structure"`
}

// String returns the string representation
func (s DeleteTransitGatewayMulticastDomainOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteTransitGatewayMulticastDomain = "DeleteTransitGatewayMulticastDomain"

// DeleteTransitGatewayMulticastDomainRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified transit gateway multicast domain.
//
//    // Example sending a request using DeleteTransitGatewayMulticastDomainRequest.
//    req := client.DeleteTransitGatewayMulticastDomainRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteTransitGatewayMulticastDomain
func (c *Client) DeleteTransitGatewayMulticastDomainRequest(input *DeleteTransitGatewayMulticastDomainInput) DeleteTransitGatewayMulticastDomainRequest {
	op := &aws.Operation{
		Name:       opDeleteTransitGatewayMulticastDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTransitGatewayMulticastDomainInput{}
	}

	req := c.newRequest(op, input, &DeleteTransitGatewayMulticastDomainOutput{})
	return DeleteTransitGatewayMulticastDomainRequest{Request: req, Input: input, Copy: c.DeleteTransitGatewayMulticastDomainRequest}
}

// DeleteTransitGatewayMulticastDomainRequest is the request type for the
// DeleteTransitGatewayMulticastDomain API operation.
type DeleteTransitGatewayMulticastDomainRequest struct {
	*aws.Request
	Input *DeleteTransitGatewayMulticastDomainInput
	Copy  func(*DeleteTransitGatewayMulticastDomainInput) DeleteTransitGatewayMulticastDomainRequest
}

// Send marshals and sends the DeleteTransitGatewayMulticastDomain API request.
func (r DeleteTransitGatewayMulticastDomainRequest) Send(ctx context.Context) (*DeleteTransitGatewayMulticastDomainResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTransitGatewayMulticastDomainResponse{
		DeleteTransitGatewayMulticastDomainOutput: r.Request.Data.(*DeleteTransitGatewayMulticastDomainOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTransitGatewayMulticastDomainResponse is the response type for the
// DeleteTransitGatewayMulticastDomain API operation.
type DeleteTransitGatewayMulticastDomainResponse struct {
	*DeleteTransitGatewayMulticastDomainOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTransitGatewayMulticastDomain request.
func (r *DeleteTransitGatewayMulticastDomainResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}