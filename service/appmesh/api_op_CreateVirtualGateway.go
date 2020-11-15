// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a virtual gateway. A virtual gateway allows resources outside your mesh
// to communicate to resources that are inside your mesh. The virtual gateway
// represents an Envoy proxy running in an Amazon ECS task, in a Kubernetes
// service, or on an Amazon EC2 instance. Unlike a virtual node, which represents
// an Envoy running with an application, a virtual gateway represents Envoy
// deployed by itself. For more information about virtual gateways, see Virtual
// gateways
// (https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_gateways.html).
func (c *Client) CreateVirtualGateway(ctx context.Context, params *CreateVirtualGatewayInput, optFns ...func(*Options)) (*CreateVirtualGatewayOutput, error) {
	if params == nil {
		params = &CreateVirtualGatewayInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVirtualGateway", params, optFns, addOperationCreateVirtualGatewayMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVirtualGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVirtualGatewayInput struct {

	// The name of the service mesh to create the virtual gateway in.
	//
	// This member is required.
	MeshName *string

	// The virtual gateway specification to apply.
	//
	// This member is required.
	Spec *types.VirtualGatewaySpec

	// The name to use for the virtual gateway.
	//
	// This member is required.
	VirtualGatewayName *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. Up to 36 letters, numbers, hyphens, and underscores are allowed.
	ClientToken *string

	// The AWS IAM account ID of the service mesh owner. If the account ID is not your
	// own, then the account that you specify must share the mesh with your account
	// before you can create the resource in the service mesh. For more information
	// about mesh sharing, see Working with shared meshes
	// (https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).
	MeshOwner *string

	// Optional metadata that you can apply to the virtual gateway to assist with
	// categorization and organization. Each tag consists of a key and an optional
	// value, both of which you define. Tag keys can have a maximum character length of
	// 128 characters, and tag values can have a maximum length of 256 characters.
	Tags []types.TagRef
}

type CreateVirtualGatewayOutput struct {

	// The full description of your virtual gateway following the create call.
	//
	// This member is required.
	VirtualGateway *types.VirtualGatewayData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateVirtualGatewayMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateVirtualGateway{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateVirtualGateway{}, middleware.After)
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
	if err = addOpCreateVirtualGatewayValidationMiddleware(stack); err != nil {
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
