// Code generated by smithy-go-codegen DO NOT EDIT.

package braket

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/braket/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/httpbinding"
	smithyjson "github.com/awslabs/smithy-go/json"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

type awsRestjson1_serializeOpCancelQuantumTask struct {
}

func (*awsRestjson1_serializeOpCancelQuantumTask) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpCancelQuantumTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CancelQuantumTaskInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/quantum-task/{quantumTaskArn}/cancel")
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

	if err := awsRestjson1_serializeOpHttpBindingsCancelQuantumTaskInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentCancelQuantumTaskInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsCancelQuantumTaskInput(v *CancelQuantumTaskInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.QuantumTaskArn == nil || len(*v.QuantumTaskArn) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member quantumTaskArn must not be empty")}
	}
	if v.QuantumTaskArn != nil {
		if err := encoder.SetURI("quantumTaskArn").String(*v.QuantumTaskArn); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentCancelQuantumTaskInput(v *CancelQuantumTaskInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ClientToken != nil {
		ok := object.Key("clientToken")
		ok.String(*v.ClientToken)
	}

	return nil
}

type awsRestjson1_serializeOpCreateQuantumTask struct {
}

func (*awsRestjson1_serializeOpCreateQuantumTask) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpCreateQuantumTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreateQuantumTaskInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/quantum-task")
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

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentCreateQuantumTaskInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsCreateQuantumTaskInput(v *CreateQuantumTaskInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentCreateQuantumTaskInput(v *CreateQuantumTaskInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Action != nil {
		ok := object.Key("action")
		ok.String(*v.Action)
	}

	if v.ClientToken != nil {
		ok := object.Key("clientToken")
		ok.String(*v.ClientToken)
	}

	if v.DeviceArn != nil {
		ok := object.Key("deviceArn")
		ok.String(*v.DeviceArn)
	}

	if v.DeviceParameters != nil {
		ok := object.Key("deviceParameters")
		ok.String(*v.DeviceParameters)
	}

	if v.OutputS3Bucket != nil {
		ok := object.Key("outputS3Bucket")
		ok.String(*v.OutputS3Bucket)
	}

	if v.OutputS3KeyPrefix != nil {
		ok := object.Key("outputS3KeyPrefix")
		ok.String(*v.OutputS3KeyPrefix)
	}

	if v.Shots != nil {
		ok := object.Key("shots")
		ok.Long(*v.Shots)
	}

	return nil
}

type awsRestjson1_serializeOpGetDevice struct {
}

func (*awsRestjson1_serializeOpGetDevice) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetDevice) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetDeviceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/device/{deviceArn}")
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

	if err := awsRestjson1_serializeOpHttpBindingsGetDeviceInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetDeviceInput(v *GetDeviceInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.DeviceArn == nil || len(*v.DeviceArn) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member deviceArn must not be empty")}
	}
	if v.DeviceArn != nil {
		if err := encoder.SetURI("deviceArn").String(*v.DeviceArn); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpGetQuantumTask struct {
}

func (*awsRestjson1_serializeOpGetQuantumTask) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetQuantumTask) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetQuantumTaskInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/quantum-task/{quantumTaskArn}")
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

	if err := awsRestjson1_serializeOpHttpBindingsGetQuantumTaskInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetQuantumTaskInput(v *GetQuantumTaskInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.QuantumTaskArn == nil || len(*v.QuantumTaskArn) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member quantumTaskArn must not be empty")}
	}
	if v.QuantumTaskArn != nil {
		if err := encoder.SetURI("quantumTaskArn").String(*v.QuantumTaskArn); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpSearchDevices struct {
}

func (*awsRestjson1_serializeOpSearchDevices) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpSearchDevices) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*SearchDevicesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/devices")
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

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentSearchDevicesInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsSearchDevicesInput(v *SearchDevicesInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentSearchDevicesInput(v *SearchDevicesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Filters != nil {
		ok := object.Key("filters")
		if err := awsRestjson1_serializeDocumentSearchDevicesFilterList(v.Filters, ok); err != nil {
			return err
		}
	}

	if v.MaxResults != nil {
		ok := object.Key("maxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("nextToken")
		ok.String(*v.NextToken)
	}

	return nil
}

type awsRestjson1_serializeOpSearchQuantumTasks struct {
}

func (*awsRestjson1_serializeOpSearchQuantumTasks) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpSearchQuantumTasks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*SearchQuantumTasksInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/quantum-tasks")
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

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentSearchQuantumTasksInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsSearchQuantumTasksInput(v *SearchQuantumTasksInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentSearchQuantumTasksInput(v *SearchQuantumTasksInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Filters != nil {
		ok := object.Key("filters")
		if err := awsRestjson1_serializeDocumentSearchQuantumTasksFilterList(v.Filters, ok); err != nil {
			return err
		}
	}

	if v.MaxResults != nil {
		ok := object.Key("maxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("nextToken")
		ok.String(*v.NextToken)
	}

	return nil
}

func awsRestjson1_serializeDocumentSearchDevicesFilter(v *types.SearchDevicesFilter, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Name != nil {
		ok := object.Key("name")
		ok.String(*v.Name)
	}

	if v.Values != nil {
		ok := object.Key("values")
		if err := awsRestjson1_serializeDocumentString256List(v.Values, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentSearchDevicesFilterList(v []types.SearchDevicesFilter, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentSearchDevicesFilter(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentSearchQuantumTasksFilter(v *types.SearchQuantumTasksFilter, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Name != nil {
		ok := object.Key("name")
		ok.String(*v.Name)
	}

	if len(v.Operator) > 0 {
		ok := object.Key("operator")
		ok.String(string(v.Operator))
	}

	if v.Values != nil {
		ok := object.Key("values")
		if err := awsRestjson1_serializeDocumentString256List(v.Values, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentSearchQuantumTasksFilterList(v []types.SearchQuantumTasksFilter, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentSearchQuantumTasksFilter(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentString256List(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}
