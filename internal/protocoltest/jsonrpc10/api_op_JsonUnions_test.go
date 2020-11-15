// Code generated by smithy-go-codegen DO NOT EDIT.

package jsonrpc10

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/jsonrpc10/types"
	"github.com/awslabs/smithy-go/middleware"
	"github.com/awslabs/smithy-go/ptr"
	smithytesting "github.com/awslabs/smithy-go/testing"
	smithytime "github.com/awslabs/smithy-go/time"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestClient_JsonUnions_awsAwsjson10Serialize(t *testing.T) {
	cases := map[string]struct {
		Params        *JsonUnionsInput
		ExpectMethod  string
		ExpectURIPath string
		ExpectQuery   []smithytesting.QueryItem
		RequireQuery  []string
		ForbidQuery   []string
		ExpectHeader  http.Header
		RequireHeader []string
		ForbidHeader  []string
		BodyMediaType string
		BodyAssert    func(io.Reader) error
	}{
		// Serializes a string union value
		"AwsJson10SerializeStringUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberStringValue{Value: "foo"},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
				"X-Amz-Target": []string{"JsonRpc10.JsonUnions"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "stringValue": "foo"
			    }
			}`))
			},
		},
		// Serializes a boolean union value
		"AwsJson10SerializeBooleanUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberBooleanValue{Value: true},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
				"X-Amz-Target": []string{"JsonRpc10.JsonUnions"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "booleanValue": true
			    }
			}`))
			},
		},
		// Serializes a number union value
		"AwsJson10SerializeNumberUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberNumberValue{Value: 1},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
				"X-Amz-Target": []string{"JsonRpc10.JsonUnions"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "numberValue": 1
			    }
			}`))
			},
		},
		// Serializes a blob union value
		"AwsJson10SerializeBlobUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberBlobValue{Value: []byte("foo")},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
				"X-Amz-Target": []string{"JsonRpc10.JsonUnions"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "blobValue": "Zm9v"
			    }
			}`))
			},
		},
		// Serializes a timestamp union value
		"AwsJson10SerializeTimestampUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberTimestampValue{Value: smithytime.ParseEpochSeconds(1398796238)},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
				"X-Amz-Target": []string{"JsonRpc10.JsonUnions"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "timestampValue": 1398796238
			    }
			}`))
			},
		},
		// Serializes an enum union value
		"AwsJson10SerializeEnumUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberEnumValue{Value: types.FooEnum("Foo")},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
				"X-Amz-Target": []string{"JsonRpc10.JsonUnions"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "enumValue": "Foo"
			    }
			}`))
			},
		},
		// Serializes a list union value
		"AwsJson10SerializeListUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberListValue{Value: []string{
					"foo",
					"bar",
				}},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
				"X-Amz-Target": []string{"JsonRpc10.JsonUnions"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "listValue": ["foo", "bar"]
			    }
			}`))
			},
		},
		// Serializes a map union value
		"AwsJson10SerializeMapUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberMapValue{Value: map[string]string{
					"foo":  "bar",
					"spam": "eggs",
				}},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
				"X-Amz-Target": []string{"JsonRpc10.JsonUnions"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "mapValue": {
			            "foo": "bar",
			            "spam": "eggs"
			        }
			    }
			}`))
			},
		},
		// Serializes a structure union value
		"AwsJson10SerializeStructureUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberStructureValue{Value: types.GreetingStruct{
					Hi: ptr.String("hello"),
				}},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
				"X-Amz-Target": []string{"JsonRpc10.JsonUnions"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "structureValue": {
			            "hi": "hello"
			        }
			    }
			}`))
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			var actualReq *http.Request
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				actualReq = r.Clone(r.Context())
				if len(actualReq.URL.RawPath) == 0 {
					actualReq.URL.RawPath = actualReq.URL.Path
				}
				if v := actualReq.ContentLength; v != 0 {
					actualReq.Header.Set("Content-Length", strconv.FormatInt(v, 10))
				}
				var buf bytes.Buffer
				if _, err := io.Copy(&buf, r.Body); err != nil {
					t.Errorf("failed to read request body, %v", err)
				}
				actualReq.Body = ioutil.NopCloser(&buf)

				w.WriteHeader(200)
			}))
			defer server.Close()
			url := server.URL
			client := New(Options{
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = url
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient: awshttp.NewBuildableClient(),
				Region:     "us-west-2",
			})
			result, err := client.JsonUnions(context.Background(), c.Params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if e, a := c.ExpectMethod, actualReq.Method; e != a {
				t.Errorf("expect %v method, got %v", e, a)
			}
			if e, a := c.ExpectURIPath, actualReq.URL.RawPath; e != a {
				t.Errorf("expect %v path, got %v", e, a)
			}
			queryItems := smithytesting.ParseRawQuery(actualReq.URL.RawQuery)
			smithytesting.AssertHasQuery(t, c.ExpectQuery, queryItems)
			smithytesting.AssertHasQueryKeys(t, c.RequireQuery, queryItems)
			smithytesting.AssertNotHaveQueryKeys(t, c.ForbidQuery, queryItems)
			smithytesting.AssertHasHeader(t, c.ExpectHeader, actualReq.Header)
			smithytesting.AssertHasHeaderKeys(t, c.RequireHeader, actualReq.Header)
			smithytesting.AssertNotHaveHeaderKeys(t, c.ForbidHeader, actualReq.Header)
			if c.BodyAssert != nil {
				if err := c.BodyAssert(actualReq.Body); err != nil {
					t.Errorf("expect body equal, got %v", err)
				}
			}
		})
	}
}

func TestClient_JsonUnions_awsAwsjson10Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *JsonUnionsOutput
	}{
		// Deserializes a string union value
		"AwsJson10DeserializeStringUnionValue": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "contents": {
			        "stringValue": "foo"
			    }
			}`),
			ExpectResult: &JsonUnionsOutput{
				Contents: &types.MyUnionMemberStringValue{Value: "foo"},
			},
		},
		// Deserializes a boolean union value
		"AwsJson10DeserializeBooleanUnionValue": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "contents": {
			        "booleanValue": true
			    }
			}`),
			ExpectResult: &JsonUnionsOutput{
				Contents: &types.MyUnionMemberBooleanValue{Value: true},
			},
		},
		// Deserializes a number union value
		"AwsJson10DeserializeNumberUnionValue": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "contents": {
			        "numberValue": 1
			    }
			}`),
			ExpectResult: &JsonUnionsOutput{
				Contents: &types.MyUnionMemberNumberValue{Value: 1},
			},
		},
		// Deserializes a blob union value
		"AwsJson10DeserializeBlobUnionValue": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "contents": {
			        "blobValue": "Zm9v"
			    }
			}`),
			ExpectResult: &JsonUnionsOutput{
				Contents: &types.MyUnionMemberBlobValue{Value: []byte("foo")},
			},
		},
		// Deserializes a timestamp union value
		"AwsJson10DeserializeTimestampUnionValue": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "contents": {
			        "timestampValue": 1398796238
			    }
			}`),
			ExpectResult: &JsonUnionsOutput{
				Contents: &types.MyUnionMemberTimestampValue{Value: smithytime.ParseEpochSeconds(1398796238)},
			},
		},
		// Deserializes an enum union value
		"AwsJson10DeserializeEnumUnionValue": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "contents": {
			        "enumValue": "Foo"
			    }
			}`),
			ExpectResult: &JsonUnionsOutput{
				Contents: &types.MyUnionMemberEnumValue{Value: types.FooEnum("Foo")},
			},
		},
		// Deserializes a list union value
		"AwsJson10DeserializeListUnionValue": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "contents": {
			        "listValue": ["foo", "bar"]
			    }
			}`),
			ExpectResult: &JsonUnionsOutput{
				Contents: &types.MyUnionMemberListValue{Value: []string{
					"foo",
					"bar",
				}},
			},
		},
		// Deserializes a map union value
		"AwsJson10DeserializeMapUnionValue": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "contents": {
			        "mapValue": {
			            "foo": "bar",
			            "spam": "eggs"
			        }
			    }
			}`),
			ExpectResult: &JsonUnionsOutput{
				Contents: &types.MyUnionMemberMapValue{Value: map[string]string{
					"foo":  "bar",
					"spam": "eggs",
				}},
			},
		},
		// Deserializes a structure union value
		"AwsJson10DeserializeStructureUnionValue": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "contents": {
			        "structureValue": {
			            "hi": "hello"
			        }
			    }
			}`),
			ExpectResult: &JsonUnionsOutput{
				Contents: &types.MyUnionMemberStructureValue{Value: types.GreetingStruct{
					Hi: ptr.String("hello"),
				}},
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			url := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = url
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				Region: "us-west-2",
			})
			var params JsonUnionsInput
			result, err := client.JsonUnions(context.Background(), &params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if err := smithytesting.CompareValues(c.ExpectResult, result, cmpopts.IgnoreUnexported(middleware.Metadata{})); err != nil {
				t.Errorf("expect c.ExpectResult value match:\n%v", err)
			}
		})
	}
}
