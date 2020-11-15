// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns a description of a notebook instance lifecycle configuration. For
// information about notebook instance lifestyle configurations, see Step 2.1:
// (Optional) Customize a Notebook Instance
// (https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html).
func (c *Client) DescribeNotebookInstanceLifecycleConfig(ctx context.Context, params *DescribeNotebookInstanceLifecycleConfigInput, optFns ...func(*Options)) (*DescribeNotebookInstanceLifecycleConfigOutput, error) {
	if params == nil {
		params = &DescribeNotebookInstanceLifecycleConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeNotebookInstanceLifecycleConfig", params, optFns, addOperationDescribeNotebookInstanceLifecycleConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeNotebookInstanceLifecycleConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeNotebookInstanceLifecycleConfigInput struct {

	// The name of the lifecycle configuration to describe.
	//
	// This member is required.
	NotebookInstanceLifecycleConfigName *string
}

type DescribeNotebookInstanceLifecycleConfigOutput struct {

	// A timestamp that tells when the lifecycle configuration was created.
	CreationTime *time.Time

	// A timestamp that tells when the lifecycle configuration was last modified.
	LastModifiedTime *time.Time

	// The Amazon Resource Name (ARN) of the lifecycle configuration.
	NotebookInstanceLifecycleConfigArn *string

	// The name of the lifecycle configuration.
	NotebookInstanceLifecycleConfigName *string

	// The shell script that runs only once, when you create a notebook instance.
	OnCreate []types.NotebookInstanceLifecycleHook

	// The shell script that runs every time you start a notebook instance, including
	// when you create the notebook instance.
	OnStart []types.NotebookInstanceLifecycleHook

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeNotebookInstanceLifecycleConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeNotebookInstanceLifecycleConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeNotebookInstanceLifecycleConfig{}, middleware.After)
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
	if err = addOpDescribeNotebookInstanceLifecycleConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNotebookInstanceLifecycleConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeNotebookInstanceLifecycleConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeNotebookInstanceLifecycleConfig",
	}
}
