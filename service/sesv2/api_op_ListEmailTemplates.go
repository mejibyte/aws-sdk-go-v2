// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the email templates present in your Amazon SES account in the current AWS
// Region. You can execute this operation no more than once per second.
func (c *Client) ListEmailTemplates(ctx context.Context, params *ListEmailTemplatesInput, optFns ...func(*Options)) (*ListEmailTemplatesOutput, error) {
	if params == nil {
		params = &ListEmailTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEmailTemplates", params, optFns, addOperationListEmailTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEmailTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to list the email templates present in your Amazon SES
// account in the current AWS Region. For more information, see the Amazon SES
// Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/send-personalized-email-api.html).
type ListEmailTemplatesInput struct {

	// A token returned from a previous call to ListEmailTemplates to indicate the
	// position in the list of email templates.
	NextToken *string

	// The number of results to show in a single call to ListEmailTemplates. If the
	// number of results is larger than the number you specified in this parameter,
	// then the response includes a NextToken element, which you can use to obtain
	// additional results. The value you specify has to be at least 1, and can be no
	// more than 10.
	PageSize *int32
}

// The following elements are returned by the service.
type ListEmailTemplatesOutput struct {

	// A token indicating that there are additional email templates available to be
	// listed. Pass this token to a subsequent ListEmailTemplates call to retrieve the
	// next 10 email templates.
	NextToken *string

	// An array the contains the name and creation time stamp for each template in your
	// Amazon SES account.
	TemplatesMetadata []types.EmailTemplateMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListEmailTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEmailTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEmailTemplates{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEmailTemplates(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListEmailTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListEmailTemplates",
	}
}
