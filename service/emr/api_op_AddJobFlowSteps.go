// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input argument to the AddJobFlowSteps operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/AddJobFlowStepsInput
type AddJobFlowStepsInput struct {
	_ struct{} `type:"structure"`

	// A string that uniquely identifies the job flow. This identifier is returned
	// by RunJobFlow and can also be obtained from ListClusters.
	//
	// JobFlowId is a required field
	JobFlowId *string `type:"string" required:"true"`

	// A list of StepConfig to be executed by the job flow.
	//
	// Steps is a required field
	Steps []StepConfig `type:"list" required:"true"`
}

// String returns the string representation
func (s AddJobFlowStepsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddJobFlowStepsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddJobFlowStepsInput"}

	if s.JobFlowId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobFlowId"))
	}

	if s.Steps == nil {
		invalidParams.Add(aws.NewErrParamRequired("Steps"))
	}
	if s.Steps != nil {
		for i, v := range s.Steps {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Steps", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The output for the AddJobFlowSteps operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/AddJobFlowStepsOutput
type AddJobFlowStepsOutput struct {
	_ struct{} `type:"structure"`

	// The identifiers of the list of steps added to the job flow.
	StepIds []string `type:"list"`
}

// String returns the string representation
func (s AddJobFlowStepsOutput) String() string {
	return awsutil.Prettify(s)
}

const opAddJobFlowSteps = "AddJobFlowSteps"

// AddJobFlowStepsRequest returns a request value for making API operation for
// Amazon Elastic MapReduce.
//
// AddJobFlowSteps adds new steps to a running cluster. A maximum of 256 steps
// are allowed in each job flow.
//
// If your cluster is long-running (such as a Hive data warehouse) or complex,
// you may require more than 256 steps to process your data. You can bypass
// the 256-step limitation in various ways, including using SSH to connect to
// the master node and submitting queries directly to the software running on
// the master node, such as Hive and Hadoop. For more information on how to
// do this, see Add More than 256 Steps to a Cluster (https://docs.aws.amazon.com/emr/latest/ManagementGuide/AddMoreThan256Steps.html)
// in the Amazon EMR Management Guide.
//
// A step specifies the location of a JAR file stored either on the master node
// of the cluster or in Amazon S3. Each step is performed by the main function
// of the main class of the JAR file. The main class can be specified either
// in the manifest of the JAR or by using the MainFunction parameter of the
// step.
//
// Amazon EMR executes each step in the order listed. For a step to be considered
// complete, the main function must exit with a zero exit code and all Hadoop
// jobs started while the step was running must have completed and run successfully.
//
// You can only add steps to a cluster that is in one of the following states:
// STARTING, BOOTSTRAPPING, RUNNING, or WAITING.
//
//    // Example sending a request using AddJobFlowStepsRequest.
//    req := client.AddJobFlowStepsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/AddJobFlowSteps
func (c *Client) AddJobFlowStepsRequest(input *AddJobFlowStepsInput) AddJobFlowStepsRequest {
	op := &aws.Operation{
		Name:       opAddJobFlowSteps,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddJobFlowStepsInput{}
	}

	req := c.newRequest(op, input, &AddJobFlowStepsOutput{})
	return AddJobFlowStepsRequest{Request: req, Input: input, Copy: c.AddJobFlowStepsRequest}
}

// AddJobFlowStepsRequest is the request type for the
// AddJobFlowSteps API operation.
type AddJobFlowStepsRequest struct {
	*aws.Request
	Input *AddJobFlowStepsInput
	Copy  func(*AddJobFlowStepsInput) AddJobFlowStepsRequest
}

// Send marshals and sends the AddJobFlowSteps API request.
func (r AddJobFlowStepsRequest) Send(ctx context.Context) (*AddJobFlowStepsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddJobFlowStepsResponse{
		AddJobFlowStepsOutput: r.Request.Data.(*AddJobFlowStepsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddJobFlowStepsResponse is the response type for the
// AddJobFlowSteps API operation.
type AddJobFlowStepsResponse struct {
	*AddJobFlowStepsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddJobFlowSteps request.
func (r *AddJobFlowStepsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}