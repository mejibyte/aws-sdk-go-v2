// Code generated by smithy-go-codegen DO NOT EDIT.

package mq

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of all revisions for the specified configuration.
func (c *Client) ListConfigurationRevisions(ctx context.Context, params *ListConfigurationRevisionsInput, optFns ...func(*Options)) (*ListConfigurationRevisionsOutput, error) {
	if params == nil {
		params = &ListConfigurationRevisionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConfigurationRevisions", params, optFns, addOperationListConfigurationRevisionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConfigurationRevisionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConfigurationRevisionsInput struct {

	// The unique ID that Amazon MQ generates for the configuration.
	//
	// This member is required.
	ConfigurationId *string

	// The maximum number of configurations that Amazon MQ can return per page (20 by
	// default). This value must be an integer from 5 to 100.
	MaxResults int32

	// The token that specifies the next page of results Amazon MQ should return. To
	// request the first page, leave nextToken empty.
	NextToken *string
}

type ListConfigurationRevisionsOutput struct {

	// The unique ID that Amazon MQ generates for the configuration.
	ConfigurationId *string

	// The maximum number of configuration revisions that can be returned per page (20
	// by default). This value must be an integer from 5 to 100.
	MaxResults int32

	// The token that specifies the next page of results Amazon MQ should return. To
	// request the first page, leave nextToken empty.
	NextToken *string

	// The list of all revisions for the specified configuration.
	Revisions []types.ConfigurationRevision

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListConfigurationRevisionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListConfigurationRevisions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListConfigurationRevisions{}, middleware.After)
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
	if err = addOpListConfigurationRevisionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListConfigurationRevisions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListConfigurationRevisions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mq",
		OperationName: "ListConfigurationRevisions",
	}
}
