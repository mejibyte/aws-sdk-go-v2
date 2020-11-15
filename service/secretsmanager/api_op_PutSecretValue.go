// Code generated by smithy-go-codegen DO NOT EDIT.

package secretsmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stores a new encrypted secret value in the specified secret. To do this, the
// operation creates a new version and attaches it to the secret. The version can
// contain a new SecretString value or a new SecretBinary value. You can also
// specify the staging labels that are initially attached to the new version. The
// Secrets Manager console uses only the SecretString field. To add binary data to
// a secret with the SecretBinary field you must use the AWS CLI or one of the AWS
// SDKs.
//
// * If this operation creates the first version for the secret then Secrets
// Manager automatically attaches the staging label AWSCURRENT to the new
// version.
//
// * If another version of this secret already exists, then this
// operation does not automatically move any staging labels other than those that
// you explicitly specify in the VersionStages parameter.
//
// * If this operation
// moves the staging label AWSCURRENT from another version to this version (because
// you included it in the StagingLabels parameter) then Secrets Manager also
// automatically moves the staging label AWSPREVIOUS to the version that AWSCURRENT
// was removed from.
//
// * This operation is idempotent. If a version with a VersionId
// with the same value as the ClientRequestToken parameter already exists and you
// specify the same secret data, the operation succeeds but does nothing. However,
// if the secret data is different, then the operation fails because you cannot
// modify an existing version; you can only create new ones.
//
// * If you call an
// operation to encrypt or decrypt the SecretString or SecretBinary for a secret in
// the same account as the calling user and that secret doesn't specify a AWS KMS
// encryption key, Secrets Manager uses the account's default AWS managed customer
// master key (CMK) with the alias aws/secretsmanager. If this key doesn't already
// exist in your account then Secrets Manager creates it for you automatically. All
// users and roles in the same AWS account automatically have access to use the
// default CMK. Note that if an Secrets Manager API call results in AWS creating
// the account's AWS-managed CMK, it can result in a one-time significant delay in
// returning the result.
//
// * If the secret resides in a different AWS account from
// the credentials calling an API that requires encryption or decryption of the
// secret value then you must create and use a custom AWS KMS CMK because you can't
// access the default CMK for the account using credentials from a different AWS
// account. Store the ARN of the CMK in the secret when you create the secret or
// when you update it by including it in the KMSKeyId. If you call an API that must
// encrypt or decrypt SecretString or SecretBinary using credentials from a
// different account then the AWS KMS key policy must grant cross-account access to
// that other account's user or role for both the kms:GenerateDataKey and
// kms:Decrypt operations.
//
// Minimum permissions To run this command, you must have
// the following permissions:
//
// * secretsmanager:PutSecretValue
//
// *
// kms:GenerateDataKey - needed only if you use a customer-managed AWS KMS key to
// encrypt the secret. You do not need this permission to use the account's default
// AWS managed CMK for Secrets Manager.
//
// Related operations
//
// * To retrieve the
// encrypted value you store in the version of a secret, use GetSecretValue.
//
// * To
// create a secret, use CreateSecret.
//
// * To get the details for a secret, use
// DescribeSecret.
//
// * To list the versions attached to a secret, use
// ListSecretVersionIds.
func (c *Client) PutSecretValue(ctx context.Context, params *PutSecretValueInput, optFns ...func(*Options)) (*PutSecretValueOutput, error) {
	if params == nil {
		params = &PutSecretValueInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutSecretValue", params, optFns, addOperationPutSecretValueMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutSecretValueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutSecretValueInput struct {

	// Specifies the secret to which you want to add a new version. You can specify
	// either the Amazon Resource Name (ARN) or the friendly name of the secret. The
	// secret must already exist. If you specify an ARN, we generally recommend that
	// you specify a complete ARN. You can specify a partial ARN too—for example, if
	// you don’t include the final hyphen and six random characters that Secrets
	// Manager adds at the end of the ARN when you created the secret. A partial ARN
	// match can work as long as it uniquely matches only one secret. However, if your
	// secret has a name that ends in a hyphen followed by six characters (before
	// Secrets Manager adds the hyphen and six characters to the ARN) and you try to
	// use that as a partial ARN, then those characters cause Secrets Manager to assume
	// that you’re specifying a complete ARN. This confusion can cause unexpected
	// results. To avoid this situation, we recommend that you don’t create secret
	// names ending with a hyphen followed by six characters. If you specify an
	// incomplete ARN without the random suffix, and instead provide the 'friendly
	// name', you must not include the random suffix. If you do include the random
	// suffix added by Secrets Manager, you receive either a ResourceNotFoundException
	// or an AccessDeniedException error, depending on your permissions.
	//
	// This member is required.
	SecretId *string

	// (Optional) Specifies a unique identifier for the new version of the secret. If
	// you use the AWS CLI or one of the AWS SDK to call this operation, then you can
	// leave this parameter empty. The CLI or SDK generates a random UUID for you and
	// includes that in the request. If you don't use the SDK and instead generate a
	// raw HTTP request to the Secrets Manager service endpoint, then you must generate
	// a ClientRequestToken yourself for new versions and include that value in the
	// request. This value helps ensure idempotency. Secrets Manager uses this value to
	// prevent the accidental creation of duplicate versions if there are failures and
	// retries during the Lambda rotation function's processing. We recommend that you
	// generate a UUID-type (https://wikipedia.org/wiki/Universally_unique_identifier)
	// value to ensure uniqueness within the specified secret.
	//
	// * If the
	// ClientRequestToken value isn't already associated with a version of the secret
	// then a new version of the secret is created.
	//
	// * If a version with this value
	// already exists and that version's SecretString or SecretBinary values are the
	// same as those in the request then the request is ignored (the operation is
	// idempotent).
	//
	// * If a version with this value already exists and the version of
	// the SecretString and SecretBinary values are different from those in the request
	// then the request fails because you cannot modify an existing secret version. You
	// can only create new versions to store new secret values.
	//
	// This value becomes the
	// VersionId of the new version.
	ClientRequestToken *string

	// (Optional) Specifies binary data that you want to encrypt and store in the new
	// version of the secret. To use this parameter in the command-line tools, we
	// recommend that you store your binary data in a file and then use the appropriate
	// technique for your tool to pass the contents of the file as a parameter. Either
	// SecretBinary or SecretString must have a value, but not both. They cannot both
	// be empty. This parameter is not accessible if the secret using the Secrets
	// Manager console.
	SecretBinary []byte

	// (Optional) Specifies text data that you want to encrypt and store in this new
	// version of the secret. Either SecretString or SecretBinary must have a value,
	// but not both. They cannot both be empty. If you create this secret by using the
	// Secrets Manager console then Secrets Manager puts the protected secret text in
	// only the SecretString parameter. The Secrets Manager console stores the
	// information as a JSON structure of key/value pairs that the default Lambda
	// rotation function knows how to parse. For storing multiple values, we recommend
	// that you use a JSON text string argument and specify key/value pairs. For
	// information on how to format a JSON parameter for the various command line tool
	// environments, see Using JSON for Parameters
	// (https://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#cli-using-param-json)
	// in the AWS CLI User Guide. For example:
	// [{"username":"bob"},{"password":"abc123xyz456"}] If your command-line tool or
	// SDK requires quotation marks around the parameter, you should use single quotes
	// to avoid confusion with the double quotes required in the JSON text.
	SecretString *string

	// (Optional) Specifies a list of staging labels that are attached to this version
	// of the secret. These staging labels are used to track the versions through the
	// rotation process by the Lambda rotation function. A staging label must be unique
	// to a single version of the secret. If you specify a staging label that's already
	// associated with a different version of the same secret then that staging label
	// is automatically removed from the other version and attached to this version. If
	// you do not specify a value for VersionStages then Secrets Manager automatically
	// moves the staging label AWSCURRENT to this new version.
	VersionStages []string
}

type PutSecretValueOutput struct {

	// The Amazon Resource Name (ARN) for the secret for which you just created a
	// version.
	ARN *string

	// The friendly name of the secret for which you just created or updated a version.
	Name *string

	// The unique identifier of the version of the secret you just created or updated.
	VersionId *string

	// The list of staging labels that are currently attached to this version of the
	// secret. Staging labels are used to track a version as it progresses through the
	// secret rotation process.
	VersionStages []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutSecretValueMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutSecretValue{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutSecretValue{}, middleware.After)
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
	if err = addIdempotencyToken_opPutSecretValueMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpPutSecretValueValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutSecretValue(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpPutSecretValue struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpPutSecretValue) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpPutSecretValue) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*PutSecretValueInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *PutSecretValueInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opPutSecretValueMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpPutSecretValue{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opPutSecretValue(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "secretsmanager",
		OperationName: "PutSecretValue",
	}
}
