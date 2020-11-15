// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of all historical purchases, renewals, and system renewal
// transactions for an AWS account. The list is paginated and ordered by a
// descending timestamp (most recent transactions are first). The API returns a
// NotEligible error if the user is not permitted to invoke the operation. If you
// must be able to invoke this operation, contact aws-devicefarm-support@amazon.com
// (mailto:aws-devicefarm-support@amazon.com).
func (c *Client) ListOfferingTransactions(ctx context.Context, params *ListOfferingTransactionsInput, optFns ...func(*Options)) (*ListOfferingTransactionsOutput, error) {
	if params == nil {
		params = &ListOfferingTransactionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOfferingTransactions", params, optFns, addOperationListOfferingTransactionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOfferingTransactionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to list the offering transaction history.
type ListOfferingTransactionsInput struct {

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string
}

// Returns the transaction log of the specified offerings.
type ListOfferingTransactionsOutput struct {

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	// The audit log of subscriptions you have purchased and modified through AWS
	// Device Farm.
	OfferingTransactions []types.OfferingTransaction

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListOfferingTransactionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListOfferingTransactions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListOfferingTransactions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOfferingTransactions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListOfferingTransactions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "ListOfferingTransactions",
	}
}
