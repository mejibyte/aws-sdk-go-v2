// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all Amazon ES domains associated with the package.
func (c *Client) ListDomainsForPackage(ctx context.Context, params *ListDomainsForPackageInput, optFns ...func(*Options)) (*ListDomainsForPackageOutput, error) {
	if params == nil {
		params = &ListDomainsForPackageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDomainsForPackage", params, optFns, addOperationListDomainsForPackageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDomainsForPackageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for request parameters to ListDomainsForPackage operation.
type ListDomainsForPackageInput struct {

	// The package for which to list domains.
	//
	// This member is required.
	PackageID *string

	// Limits results to a maximum number of domains.
	MaxResults int32

	// Used for pagination. Only necessary if a previous API call includes a non-null
	// NextToken value. If provided, returns results for the next page.
	NextToken *string
}

// Container for response parameters to ListDomainsForPackage operation.
type ListDomainsForPackageOutput struct {

	// List of DomainPackageDetails objects.
	DomainPackageDetailsList []types.DomainPackageDetails

	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDomainsForPackageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDomainsForPackage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDomainsForPackage{}, middleware.After)
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
	if err = addOpListDomainsForPackageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDomainsForPackage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListDomainsForPackage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "ListDomainsForPackage",
	}
}
