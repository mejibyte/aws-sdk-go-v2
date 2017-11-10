package iotdataplane_test

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awstesting/unit"
	"github.com/aws/aws-sdk-go-v2/service/iotdataplane"
)

func TestRequireEndpointIfRegionProvided(t *testing.T) {
	cfg := unit.Config()
	cfg.DisableParamValidation = true
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("")

	svc := iotdataplane.New(cfg)
	req, _ := svc.GetThingShadowRequest(nil)
	err := req.Build()

	if e, a := "", req.Metadata.Endpoint; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if err == nil {
		t.Errorf("expect error, got none")
	}
	if e, a := aws.ErrMissingEndpoint, err; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestRequireEndpointIfNoRegionProvided(t *testing.T) {
	cfg := unit.Config()
	cfg.DisableParamValidation = true
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("")

	svc := iotdataplane.New(cfg)

	req, _ := svc.GetThingShadowRequest(nil)
	err := req.Build()

	if e, a := "", req.Metadata.Endpoint; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if err == nil {
		t.Errorf("expect error, got none")
	}
	if e, a := aws.ErrMissingEndpoint, err; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestRequireEndpointUsed(t *testing.T) {
	cfg := unit.Config()
	cfg.DisableParamValidation = true
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://endpoint")

	svc := iotdataplane.New(cfg)
	req, _ := svc.GetThingShadowRequest(nil)
	err := req.Build()

	if e, a := "https://endpoint", req.Metadata.Endpoint; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}