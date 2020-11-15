// Code generated by smithy-go-codegen DO NOT EDIT.

package timestreamwrite

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// A list of tables, along with the name, status and retention properties of each
// table.
func (c *Client) ListTables(ctx context.Context, params *ListTablesInput, optFns ...func(*Options)) (*ListTablesOutput, error) {
	if params == nil {
		params = &ListTablesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTables", params, optFns, addOperationListTablesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTablesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTablesInput struct {

	// The name of the Timestream database.
	DatabaseName *string

	// The total number of items to return in the output. If the total number of items
	// available is more than the value specified, a NextToken is provided in the
	// output. To resume pagination, provide the NextToken value as argument of a
	// subsequent API invocation.
	MaxResults *int32

	// The pagination token. To resume pagination, provide the NextToken value as
	// argument of a subsequent API invocation.
	NextToken *string
}

type ListTablesOutput struct {

	// A token to specify where to start paginating. This is the NextToken from a
	// previously truncated response.
	NextToken *string

	// A list of tables.
	Tables []types.Table

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTablesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListTables{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListTables{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTables(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListTables(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "timestream",
		OperationName: "ListTables",
	}
}
