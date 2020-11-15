// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from the
// prior release, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
// Updates the specified WebACL. A Web ACL defines a collection of rules to use to
// inspect and control web requests. Each rule has an action defined (allow, block,
// or count) for requests that match the statement of the rule. In the Web ACL, you
// assign a default action to take (allow, block) for any request that does not
// match any of the rules. The rules in a Web ACL can be a combination of the types
// Rule, RuleGroup, and managed rule group. You can associate a Web ACL with one or
// more AWS resources to protect. The resources can be Amazon CloudFront, an Amazon
// API Gateway REST API, an Application Load Balancer, or an AWS AppSync GraphQL
// API.
func (c *Client) UpdateWebACL(ctx context.Context, params *UpdateWebACLInput, optFns ...func(*Options)) (*UpdateWebACLOutput, error) {
	if params == nil {
		params = &UpdateWebACLInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateWebACL", params, optFns, addOperationUpdateWebACLMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateWebACLOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateWebACLInput struct {

	// The action to perform if none of the Rules contained in the WebACL match.
	//
	// This member is required.
	DefaultAction *types.DefaultAction

	// The unique identifier for the Web ACL. This ID is returned in the responses to
	// create and list commands. You provide it to operations like update and delete.
	//
	// This member is required.
	Id *string

	// A token used for optimistic locking. AWS WAF returns a token to your get and
	// list requests, to mark the state of the entity at the time of the request. To
	// make changes to the entity associated with the token, you provide the token to
	// operations like update and delete. AWS WAF uses the token to ensure that no
	// changes have been made to the entity since you last retrieved it. If a change
	// has been made, the update fails with a WAFOptimisticLockException. If this
	// happens, perform another get, and use the new token returned by that operation.
	//
	// This member is required.
	LockToken *string

	// The name of the Web ACL. You cannot change the name of a Web ACL after you
	// create it.
	//
	// This member is required.
	Name *string

	// Specifies whether this is for an AWS CloudFront distribution or for a regional
	// application. A regional application can be an Application Load Balancer (ALB),
	// an API Gateway REST API, or an AppSync GraphQL API. To work with CloudFront, you
	// must also specify the Region US East (N. Virginia) as follows:
	//
	// * CLI - Specify
	// the Region when you use the CloudFront scope: --scope=CLOUDFRONT
	// --region=us-east-1.
	//
	// * API and SDKs - For all calls, use the Region endpoint
	// us-east-1.
	//
	// This member is required.
	Scope types.Scope

	// Defines and enables Amazon CloudWatch metrics and web request sample collection.
	//
	// This member is required.
	VisibilityConfig *types.VisibilityConfig

	// A description of the Web ACL that helps with identification. You cannot change
	// the description of a Web ACL after you create it.
	Description *string

	// The Rule statements used to identify the web requests that you want to allow,
	// block, or count. Each rule includes one top-level statement that AWS WAF uses to
	// identify matching web requests, and parameters that govern how AWS WAF handles
	// them.
	Rules []types.Rule
}

type UpdateWebACLOutput struct {

	// A token used for optimistic locking. AWS WAF returns this token to your update
	// requests. You use NextLockToken in the same manner as you use LockToken.
	NextLockToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateWebACLMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateWebACL{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateWebACL{}, middleware.After)
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
	if err = addOpUpdateWebACLValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateWebACL(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateWebACL(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "UpdateWebACL",
	}
}
