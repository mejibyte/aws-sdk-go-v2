// Code generated by smithy-go-codegen DO NOT EDIT.

package iotjobsdataplane

import (
	"bytes"
	"context"
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/httpbinding"
	smithyjson "github.com/awslabs/smithy-go/json"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

type awsRestjson1_serializeOpDescribeJobExecution struct {
}

func (*awsRestjson1_serializeOpDescribeJobExecution) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDescribeJobExecution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeJobExecutionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/things/{thingName}/jobs/{jobId}")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsDescribeJobExecutionInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsDescribeJobExecutionInput(v *DescribeJobExecutionInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ExecutionNumber != nil {
		encoder.SetQuery("executionNumber").Long(*v.ExecutionNumber)
	}

	if v.IncludeJobDocument != nil {
		encoder.SetQuery("includeJobDocument").Boolean(*v.IncludeJobDocument)
	}

	if v.JobId == nil || len(*v.JobId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member jobId must not be empty")}
	}
	if v.JobId != nil {
		if err := encoder.SetURI("jobId").String(*v.JobId); err != nil {
			return err
		}
	}

	if v.ThingName == nil || len(*v.ThingName) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member thingName must not be empty")}
	}
	if v.ThingName != nil {
		if err := encoder.SetURI("thingName").String(*v.ThingName); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpGetPendingJobExecutions struct {
}

func (*awsRestjson1_serializeOpGetPendingJobExecutions) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetPendingJobExecutions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetPendingJobExecutionsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/things/{thingName}/jobs")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsGetPendingJobExecutionsInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetPendingJobExecutionsInput(v *GetPendingJobExecutionsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ThingName == nil || len(*v.ThingName) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member thingName must not be empty")}
	}
	if v.ThingName != nil {
		if err := encoder.SetURI("thingName").String(*v.ThingName); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpStartNextPendingJobExecution struct {
}

func (*awsRestjson1_serializeOpStartNextPendingJobExecution) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpStartNextPendingJobExecution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*StartNextPendingJobExecutionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/things/{thingName}/jobs/$next")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "PUT"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsStartNextPendingJobExecutionInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentStartNextPendingJobExecutionInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsStartNextPendingJobExecutionInput(v *StartNextPendingJobExecutionInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ThingName == nil || len(*v.ThingName) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member thingName must not be empty")}
	}
	if v.ThingName != nil {
		if err := encoder.SetURI("thingName").String(*v.ThingName); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentStartNextPendingJobExecutionInput(v *StartNextPendingJobExecutionInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.StatusDetails != nil {
		ok := object.Key("statusDetails")
		if err := awsRestjson1_serializeDocumentDetailsMap(v.StatusDetails, ok); err != nil {
			return err
		}
	}

	if v.StepTimeoutInMinutes != nil {
		ok := object.Key("stepTimeoutInMinutes")
		ok.Long(*v.StepTimeoutInMinutes)
	}

	return nil
}

type awsRestjson1_serializeOpUpdateJobExecution struct {
}

func (*awsRestjson1_serializeOpUpdateJobExecution) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpUpdateJobExecution) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UpdateJobExecutionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/things/{thingName}/jobs/{jobId}")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsUpdateJobExecutionInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentUpdateJobExecutionInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsUpdateJobExecutionInput(v *UpdateJobExecutionInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.JobId == nil || len(*v.JobId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member jobId must not be empty")}
	}
	if v.JobId != nil {
		if err := encoder.SetURI("jobId").String(*v.JobId); err != nil {
			return err
		}
	}

	if v.ThingName == nil || len(*v.ThingName) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member thingName must not be empty")}
	}
	if v.ThingName != nil {
		if err := encoder.SetURI("thingName").String(*v.ThingName); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentUpdateJobExecutionInput(v *UpdateJobExecutionInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ExecutionNumber != nil {
		ok := object.Key("executionNumber")
		ok.Long(*v.ExecutionNumber)
	}

	if v.ExpectedVersion != nil {
		ok := object.Key("expectedVersion")
		ok.Long(*v.ExpectedVersion)
	}

	if v.IncludeJobDocument != nil {
		ok := object.Key("includeJobDocument")
		ok.Boolean(*v.IncludeJobDocument)
	}

	if v.IncludeJobExecutionState != nil {
		ok := object.Key("includeJobExecutionState")
		ok.Boolean(*v.IncludeJobExecutionState)
	}

	if len(v.Status) > 0 {
		ok := object.Key("status")
		ok.String(string(v.Status))
	}

	if v.StatusDetails != nil {
		ok := object.Key("statusDetails")
		if err := awsRestjson1_serializeDocumentDetailsMap(v.StatusDetails, ok); err != nil {
			return err
		}
	}

	if v.StepTimeoutInMinutes != nil {
		ok := object.Key("stepTimeoutInMinutes")
		ok.Long(*v.StepTimeoutInMinutes)
	}

	return nil
}

func awsRestjson1_serializeDocumentDetailsMap(v map[string]string, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		om.String(v[key])
	}
	return nil
}
