// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the organizational units (OUs) in a parent organizational unit or root.
// Always check the NextToken response parameter for a null value when calling a
// List* operation. These operations can occasionally return an empty set of
// results even when there are more results available. The NextToken response
// parameter value is null only when there are no more results to display. This
// operation can be called only from the organization's management account or by a
// member account that is a delegated administrator for an AWS service.
func (c *Client) ListOrganizationalUnitsForParent(ctx context.Context, params *ListOrganizationalUnitsForParentInput, optFns ...func(*Options)) (*ListOrganizationalUnitsForParentOutput, error) {
	if params == nil {
		params = &ListOrganizationalUnitsForParentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOrganizationalUnitsForParent", params, optFns, addOperationListOrganizationalUnitsForParentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOrganizationalUnitsForParentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOrganizationalUnitsForParentInput struct {

	// The unique identifier (ID) of the root or OU whose child OUs you want to list.
	// The regex pattern (http://wikipedia.org/wiki/regex) for a parent ID string
	// requires one of the following:
	//
	// * Root - A string that begins with "r-" followed
	// by from 4 to 32 lowercase letters or digits.
	//
	// * Organizational unit (OU) - A
	// string that begins with "ou-" followed by from 4 to 32 lowercase letters or
	// digits (the ID of the root that the OU is in). This string is followed by a
	// second "-" dash and from 8 to 32 additional lowercase letters or digits.
	//
	// This member is required.
	ParentId *string

	// The total number of results that you want included on each page of the response.
	// If you do not include this parameter, it defaults to a value that is specific to
	// the operation. If additional items exist beyond the maximum you specify, the
	// NextToken response element is present and has a value (is not null). Include
	// that value as the NextToken request parameter in the next call to the operation
	// to get the next part of the results. Note that Organizations might return fewer
	// results than the maximum even when there are more results available. You should
	// check NextToken after every operation to ensure that you receive all of the
	// results.
	MaxResults *int32

	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more output
	// is available. Set this parameter to the value of the previous call's NextToken
	// response to indicate where the output should continue from.
	NextToken *string
}

type ListOrganizationalUnitsForParentOutput struct {

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null.
	NextToken *string

	// A list of the OUs in the specified root or parent OU.
	OrganizationalUnits []types.OrganizationalUnit

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListOrganizationalUnitsForParentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListOrganizationalUnitsForParent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListOrganizationalUnitsForParent{}, middleware.After)
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
	if err = addOpListOrganizationalUnitsForParentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOrganizationalUnitsForParent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListOrganizationalUnitsForParent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "ListOrganizationalUnitsForParent",
	}
}
