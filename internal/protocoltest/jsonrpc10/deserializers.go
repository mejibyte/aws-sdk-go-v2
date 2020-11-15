// Code generated by smithy-go-codegen DO NOT EDIT.

package jsonrpc10

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/jsonrpc10/types"
	smithy "github.com/awslabs/smithy-go"
	smithyio "github.com/awslabs/smithy-go/io"
	"github.com/awslabs/smithy-go/middleware"
	"github.com/awslabs/smithy-go/ptr"
	smithytime "github.com/awslabs/smithy-go/time"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"io"
	"strings"
	"time"
)

type awsAwsjson10_deserializeOpEmptyInputAndEmptyOutput struct {
}

func (*awsAwsjson10_deserializeOpEmptyInputAndEmptyOutput) ID() string {
	return "OperationDeserializer"
}

func (m *awsAwsjson10_deserializeOpEmptyInputAndEmptyOutput) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsAwsjson10_deserializeOpErrorEmptyInputAndEmptyOutput(response, &metadata)
	}
	output := &EmptyInputAndEmptyOutputOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	err = awsAwsjson10_deserializeOpDocumentEmptyInputAndEmptyOutputOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	return out, metadata, err
}

func awsAwsjson10_deserializeOpErrorEmptyInputAndEmptyOutput(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

type awsAwsjson10_deserializeOpGreetingWithErrors struct {
}

func (*awsAwsjson10_deserializeOpGreetingWithErrors) ID() string {
	return "OperationDeserializer"
}

func (m *awsAwsjson10_deserializeOpGreetingWithErrors) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsAwsjson10_deserializeOpErrorGreetingWithErrors(response, &metadata)
	}
	output := &GreetingWithErrorsOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	err = awsAwsjson10_deserializeOpDocumentGreetingWithErrorsOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	return out, metadata, err
}

func awsAwsjson10_deserializeOpErrorGreetingWithErrors(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("ComplexError", errorCode):
		return awsAwsjson10_deserializeErrorComplexError(response, errorBody)

	case strings.EqualFold("FooError", errorCode):
		return awsAwsjson10_deserializeErrorFooError(response, errorBody)

	case strings.EqualFold("InvalidGreeting", errorCode):
		return awsAwsjson10_deserializeErrorInvalidGreeting(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

type awsAwsjson10_deserializeOpJsonUnions struct {
}

func (*awsAwsjson10_deserializeOpJsonUnions) ID() string {
	return "OperationDeserializer"
}

func (m *awsAwsjson10_deserializeOpJsonUnions) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsAwsjson10_deserializeOpErrorJsonUnions(response, &metadata)
	}
	output := &JsonUnionsOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	err = awsAwsjson10_deserializeOpDocumentJsonUnionsOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	return out, metadata, err
}

func awsAwsjson10_deserializeOpErrorJsonUnions(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

type awsAwsjson10_deserializeOpNoInputAndNoOutput struct {
}

func (*awsAwsjson10_deserializeOpNoInputAndNoOutput) ID() string {
	return "OperationDeserializer"
}

func (m *awsAwsjson10_deserializeOpNoInputAndNoOutput) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsAwsjson10_deserializeOpErrorNoInputAndNoOutput(response, &metadata)
	}
	output := &NoInputAndNoOutputOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	err = awsAwsjson10_deserializeDocumentNoInputAndNoOutputOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	return out, metadata, err
}

func awsAwsjson10_deserializeOpErrorNoInputAndNoOutput(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

type awsAwsjson10_deserializeOpNoInputAndOutput struct {
}

func (*awsAwsjson10_deserializeOpNoInputAndOutput) ID() string {
	return "OperationDeserializer"
}

func (m *awsAwsjson10_deserializeOpNoInputAndOutput) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsAwsjson10_deserializeOpErrorNoInputAndOutput(response, &metadata)
	}
	output := &NoInputAndOutputOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	err = awsAwsjson10_deserializeOpDocumentNoInputAndOutputOutput(&output, shape)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return out, metadata, err
	}

	return out, metadata, err
}

func awsAwsjson10_deserializeOpErrorNoInputAndOutput(response *smithyhttp.Response, metadata *middleware.Metadata) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsAwsjson10_deserializeErrorComplexError(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	output := &types.ComplexError{}
	err := awsAwsjson10_deserializeDocumentComplexError(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	return output
}

func awsAwsjson10_deserializeErrorFooError(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	output := &types.FooError{}
	err := awsAwsjson10_deserializeDocumentFooError(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	return output
}

func awsAwsjson10_deserializeErrorInvalidGreeting(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()
	var shape interface{}
	if err := decoder.Decode(&shape); err != nil && err != io.EOF {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	output := &types.InvalidGreeting{}
	err := awsAwsjson10_deserializeDocumentInvalidGreeting(&output, shape)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		err = &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
		return err
	}

	errorBody.Seek(0, io.SeekStart)
	return output
}

func awsAwsjson10_deserializeDocumentComplexError(v **types.ComplexError, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.ComplexError
	if *v == nil {
		sv = &types.ComplexError{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Nested":
			if err := awsAwsjson10_deserializeDocumentComplexNestedErrorData(&sv.Nested, value); err != nil {
				return err
			}

		case "TopLevel":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.TopLevel = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentComplexNestedErrorData(v **types.ComplexNestedErrorData, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.ComplexNestedErrorData
	if *v == nil {
		sv = &types.ComplexNestedErrorData{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Fooooo":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Foo = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentFooError(v **types.FooError, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.FooError
	if *v == nil {
		sv = &types.FooError{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentInvalidGreeting(v **types.InvalidGreeting, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.InvalidGreeting
	if *v == nil {
		sv = &types.InvalidGreeting{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "Message":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Message = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentMyUnion(v *types.MyUnion, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var uv types.MyUnion
loop:
	for key, value := range shape {
		switch key {
		case "blobValue":
			var mv []byte
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected Blob to be []byte, got %T instead", value)
				}
				dv, err := base64.StdEncoding.DecodeString(jtv)
				if err != nil {
					return fmt.Errorf("failed to base64 decode Blob, %w", err)
				}
				mv = dv
			}
			uv = &types.MyUnionMemberBlobValue{Value: mv}
			break loop

		case "booleanValue":
			var mv bool
			if value != nil {
				jtv, ok := value.(bool)
				if !ok {
					return fmt.Errorf("expected Boolean to be of type *bool, got %T instead", value)
				}
				mv = jtv
			}
			uv = &types.MyUnionMemberBooleanValue{Value: mv}
			break loop

		case "enumValue":
			var mv types.FooEnum
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected FooEnum to be of type string, got %T instead", value)
				}
				mv = types.FooEnum(jtv)
			}
			uv = &types.MyUnionMemberEnumValue{Value: mv}
			break loop

		case "listValue":
			var mv []string
			if err := awsAwsjson10_deserializeDocumentStringList(&mv, value); err != nil {
				return err
			}
			uv = &types.MyUnionMemberListValue{Value: mv}
			break loop

		case "mapValue":
			var mv map[string]string
			if err := awsAwsjson10_deserializeDocumentStringMap(&mv, value); err != nil {
				return err
			}
			uv = &types.MyUnionMemberMapValue{Value: mv}
			break loop

		case "numberValue":
			var mv int32
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected Integer to be json.Number, got %T instead", value)
				}
				i64, err := jtv.Int64()
				if err != nil {
					return err
				}
				mv = int32(i64)
			}
			uv = &types.MyUnionMemberNumberValue{Value: mv}
			break loop

		case "stringValue":
			var mv string
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				mv = jtv
			}
			uv = &types.MyUnionMemberStringValue{Value: mv}
			break loop

		case "structureValue":
			var mv types.GreetingStruct
			destAddr := &mv
			if err := awsAwsjson10_deserializeDocumentGreetingStruct(&destAddr, value); err != nil {
				return err
			}
			mv = *destAddr
			uv = &types.MyUnionMemberStructureValue{Value: mv}
			break loop

		case "timestampValue":
			var mv time.Time
			if value != nil {
				jtv, ok := value.(json.Number)
				if !ok {
					return fmt.Errorf("expected Timestamp to be json.Number, got %T instead", value)
				}
				f64, err := jtv.Float64()
				if err != nil {
					return err
				}
				mv = smithytime.ParseEpochSeconds(f64)
			}
			uv = &types.MyUnionMemberTimestampValue{Value: mv}
			break loop

		default:
			uv = &types.UnknownUnionMember{Tag: key}
			// TODO: FIX ME
			_ = value
			break loop

		}
	}
	*v = uv
	return nil
}

func awsAwsjson10_deserializeDocumentGreetingStruct(v **types.GreetingStruct, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *types.GreetingStruct
	if *v == nil {
		sv = &types.GreetingStruct{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "hi":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Hi = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentStringList(v *[]string, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.([]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var cv []string
	if *v == nil {
		cv = []string{}
	} else {
		cv = *v
	}

	for _, value := range shape {
		var col string
		if value != nil {
			jtv, ok := value.(string)
			if !ok {
				return fmt.Errorf("expected String to be of type string, got %T instead", value)
			}
			col = jtv
		}
		cv = append(cv, col)

	}
	*v = cv
	return nil
}

func awsAwsjson10_deserializeDocumentStringMap(v *map[string]string, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var mv map[string]string
	if *v == nil {
		mv = map[string]string{}
	} else {
		mv = *v
	}

	for key, value := range shape {
		var parsedVal string
		if value != nil {
			jtv, ok := value.(string)
			if !ok {
				return fmt.Errorf("expected String to be of type string, got %T instead", value)
			}
			parsedVal = jtv
		}
		mv[key] = parsedVal

	}
	*v = mv
	return nil
}

func awsAwsjson10_deserializeOpDocumentEmptyInputAndEmptyOutputOutput(v **EmptyInputAndEmptyOutputOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *EmptyInputAndEmptyOutputOutput
	if *v == nil {
		sv = &EmptyInputAndEmptyOutputOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeOpDocumentGreetingWithErrorsOutput(v **GreetingWithErrorsOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *GreetingWithErrorsOutput
	if *v == nil {
		sv = &GreetingWithErrorsOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "greeting":
			if value != nil {
				jtv, ok := value.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", value)
				}
				sv.Greeting = ptr.String(jtv)
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeOpDocumentJsonUnionsOutput(v **JsonUnionsOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *JsonUnionsOutput
	if *v == nil {
		sv = &JsonUnionsOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		case "contents":
			if err := awsAwsjson10_deserializeDocumentMyUnion(&sv.Contents, value); err != nil {
				return err
			}

		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentNoInputAndNoOutputOutput(v **NoInputAndNoOutputOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *NoInputAndNoOutputOutput
	if *v == nil {
		sv = &NoInputAndNoOutputOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}

func awsAwsjson10_deserializeOpDocumentNoInputAndOutputOutput(v **NoInputAndOutputOutput, value interface{}) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	if value == nil {
		return nil
	}

	shape, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected JSON type %v", value)
	}

	var sv *NoInputAndOutputOutput
	if *v == nil {
		sv = &NoInputAndOutputOutput{}
	} else {
		sv = *v
	}

	for key, value := range shape {
		switch key {
		default:
			_, _ = key, value

		}
	}
	*v = sv
	return nil
}
