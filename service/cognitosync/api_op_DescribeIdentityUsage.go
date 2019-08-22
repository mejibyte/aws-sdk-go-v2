// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitosync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request for information about the usage of an identity pool.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-sync-2014-06-30/DescribeIdentityUsageRequest
type DescribeIdentityUsageInput struct {
	_ struct{} `type:"structure"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	//
	// IdentityId is a required field
	IdentityId *string `location:"uri" locationName:"IdentityId" min:"1" type:"string" required:"true"`

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	//
	// IdentityPoolId is a required field
	IdentityPoolId *string `location:"uri" locationName:"IdentityPoolId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeIdentityUsageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeIdentityUsageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeIdentityUsageInput"}

	if s.IdentityId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdentityId"))
	}
	if s.IdentityId != nil && len(*s.IdentityId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdentityId", 1))
	}

	if s.IdentityPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdentityPoolId"))
	}
	if s.IdentityPoolId != nil && len(*s.IdentityPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdentityPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeIdentityUsageInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.IdentityId != nil {
		v := *s.IdentityId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "IdentityId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IdentityPoolId != nil {
		v := *s.IdentityPoolId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "IdentityPoolId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The response to a successful DescribeIdentityUsage request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-sync-2014-06-30/DescribeIdentityUsageResponse
type DescribeIdentityUsageOutput struct {
	_ struct{} `type:"structure"`

	// Usage information for the identity.
	IdentityUsage *IdentityUsage `type:"structure"`
}

// String returns the string representation
func (s DescribeIdentityUsageOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeIdentityUsageOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.IdentityUsage != nil {
		v := s.IdentityUsage

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "IdentityUsage", v, metadata)
	}
	return nil
}

const opDescribeIdentityUsage = "DescribeIdentityUsage"

// DescribeIdentityUsageRequest returns a request value for making API operation for
// Amazon Cognito Sync.
//
// Gets usage information for an identity, including number of datasets and
// data usage.
//
// This API can be called with temporary user credentials provided by Cognito
// Identity or with developer credentials.
//
//    // Example sending a request using DescribeIdentityUsageRequest.
//    req := client.DescribeIdentityUsageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-sync-2014-06-30/DescribeIdentityUsage
func (c *Client) DescribeIdentityUsageRequest(input *DescribeIdentityUsageInput) DescribeIdentityUsageRequest {
	op := &aws.Operation{
		Name:       opDescribeIdentityUsage,
		HTTPMethod: "GET",
		HTTPPath:   "/identitypools/{IdentityPoolId}/identities/{IdentityId}",
	}

	if input == nil {
		input = &DescribeIdentityUsageInput{}
	}

	req := c.newRequest(op, input, &DescribeIdentityUsageOutput{})
	return DescribeIdentityUsageRequest{Request: req, Input: input, Copy: c.DescribeIdentityUsageRequest}
}

// DescribeIdentityUsageRequest is the request type for the
// DescribeIdentityUsage API operation.
type DescribeIdentityUsageRequest struct {
	*aws.Request
	Input *DescribeIdentityUsageInput
	Copy  func(*DescribeIdentityUsageInput) DescribeIdentityUsageRequest
}

// Send marshals and sends the DescribeIdentityUsage API request.
func (r DescribeIdentityUsageRequest) Send(ctx context.Context) (*DescribeIdentityUsageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeIdentityUsageResponse{
		DescribeIdentityUsageOutput: r.Request.Data.(*DescribeIdentityUsageOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeIdentityUsageResponse is the response type for the
// DescribeIdentityUsage API operation.
type DescribeIdentityUsageResponse struct {
	*DescribeIdentityUsageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeIdentityUsage request.
func (r *DescribeIdentityUsageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}