// Code generated by smithy-go-codegen DO NOT EDIT.

package identitystore

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the attribute name and value of the group that you specified in the
// search. We only support DisplayName as a valid filter attribute path currently,
// and filter is required. This API returns minimum attributes, including GroupId
// and group DisplayName in the response.
func (c *Client) ListGroups(ctx context.Context, params *ListGroupsInput, optFns ...func(*Options)) (*ListGroupsOutput, error) {
	if params == nil {
		params = &ListGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGroups", params, optFns, addOperationListGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGroupsInput struct {

	// The globally unique identifier for the identity store, such as d-1234567890. In
	// this example, d- is a fixed prefix, and 1234567890 is a randomly generated
	// string which contains number and lower case letters. This value is generated at
	// the time that a new identity store is created.
	//
	// This member is required.
	IdentityStoreId *string

	// A list of Filter objects, which is used in the ListUsers and ListGroups request.
	Filters []types.Filter

	// The maximum number of results to be returned per request, which is used in the
	// ListUsers and ListGroups request to specify how many results to return in one
	// page. The length limit is 50 characters.
	MaxResults *int32

	// The pagination token used for the ListUsers and ListGroups APIs. This value is
	// generated by the identity store service and is returned in the API response if
	// the total results are more than the size of one page, and when this token is
	// used in the API request to search for the next page.
	NextToken *string
}

type ListGroupsOutput struct {

	// A list of Group objects in the identity store.
	//
	// This member is required.
	Groups []types.Group

	// The pagination token used for the ListUsers and ListGroups APIs. This value is
	// generated by the identity store service and is returned in the API response if
	// the total results are more than the size of one page, and when this token is
	// used in the API request to search for the next page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListGroups{}, middleware.After)
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
	if err = addOpListGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGroups(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "identitystore",
		OperationName: "ListGroups",
	}
}
