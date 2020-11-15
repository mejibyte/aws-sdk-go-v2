// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the settings for a replication group.
//
// * Scaling for Amazon ElastiCache
// for Redis (cluster mode enabled)
// (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/scaling-redis-cluster-mode-enabled.html)
// in the ElastiCache User Guide
//
// * ModifyReplicationGroupShardConfiguration
// (https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ModifyReplicationGroupShardConfiguration.html)
// in the ElastiCache API Reference
//
// This operation is valid for Redis only.
func (c *Client) ModifyReplicationGroup(ctx context.Context, params *ModifyReplicationGroupInput, optFns ...func(*Options)) (*ModifyReplicationGroupOutput, error) {
	if params == nil {
		params = &ModifyReplicationGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyReplicationGroup", params, optFns, addOperationModifyReplicationGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyReplicationGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a ModifyReplicationGroups operation.
type ModifyReplicationGroupInput struct {

	// The identifier of the replication group to modify.
	//
	// This member is required.
	ReplicationGroupId *string

	// If true, this parameter causes the modifications in this request and any pending
	// modifications to be applied, asynchronously and as soon as possible, regardless
	// of the PreferredMaintenanceWindow setting for the replication group. If false,
	// changes to the nodes in the replication group are applied on the next
	// maintenance reboot, or the next failure reboot, whichever occurs first. Valid
	// values: true | false Default: false
	ApplyImmediately bool

	// Reserved parameter. The password used to access a password protected server.
	// This parameter must be specified with the auth-token-update-strategy  parameter.
	// Password constraints:
	//
	// * Must be only printable ASCII characters
	//
	// * Must be at
	// least 16 characters and no more than 128 characters in length
	//
	// * Cannot contain
	// any of the following characters: '/', '"', or '@', '%'
	//
	// For more information,
	// see AUTH password at AUTH (http://redis.io/commands/AUTH).
	AuthToken *string

	// Specifies the strategy to use to update the AUTH token. This parameter must be
	// specified with the auth-token parameter. Possible values:
	//
	// * Rotate
	//
	// * Set
	//
	// For
	// more information, see Authenticating Users with Redis AUTH
	// (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/auth.html)
	AuthTokenUpdateStrategy types.AuthTokenUpdateStrategyType

	// This parameter is currently disabled.
	AutoMinorVersionUpgrade *bool

	// Determines whether a read replica is automatically promoted to read/write
	// primary if the existing primary encounters a failure. Valid values: true | false
	AutomaticFailoverEnabled *bool

	// A valid cache node type that you want to scale this replication group to.
	CacheNodeType *string

	// The name of the cache parameter group to apply to all of the clusters in this
	// replication group. This change is asynchronously applied as soon as possible for
	// parameters when the ApplyImmediately parameter is specified as true for this
	// request.
	CacheParameterGroupName *string

	// A list of cache security group names to authorize for the clusters in this
	// replication group. This change is asynchronously applied as soon as possible.
	// This parameter can be used only with replication group containing clusters
	// running outside of an Amazon Virtual Private Cloud (Amazon VPC). Constraints:
	// Must contain no more than 255 alphanumeric characters. Must not be Default.
	CacheSecurityGroupNames []string

	// The upgraded version of the cache engine to be run on the clusters in the
	// replication group. Important: You can upgrade to a newer engine version (see
	// Selecting a Cache Engine and Version
	// (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/SelectEngine.html#VersionManagement)),
	// but you cannot downgrade to an earlier engine version. If you want to use an
	// earlier engine version, you must delete the existing replication group and
	// create it anew with the earlier engine version.
	EngineVersion *string

	// A flag indicating if you have Multi-AZ enabled to enhance fault tolerance. For
	// more information, see Minimizing Downtime: Multi-AZ
	// (http://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/AutoFailover.html).
	MultiAZEnabled *bool

	// Deprecated. This parameter is not used.
	NodeGroupId *string

	// The Amazon Resource Name (ARN) of the Amazon SNS topic to which notifications
	// are sent. The Amazon SNS topic owner must be same as the replication group
	// owner.
	NotificationTopicArn *string

	// The status of the Amazon SNS notification topic for the replication group.
	// Notifications are sent only if the status is active. Valid values: active |
	// inactive
	NotificationTopicStatus *string

	// Specifies the weekly time range during which maintenance on the cluster is
	// performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H
	// Clock UTC). The minimum maintenance window is a 60 minute period. Valid values
	// for ddd are:
	//
	// * sun
	//
	// * mon
	//
	// * tue
	//
	// * wed
	//
	// * thu
	//
	// * fri
	//
	// * sat
	//
	// Example:
	// sun:23:00-mon:01:30
	PreferredMaintenanceWindow *string

	// For replication groups with a single primary, if this parameter is specified,
	// ElastiCache promotes the specified cluster in the specified replication group to
	// the primary role. The nodes of all other clusters in the replication group are
	// read replicas.
	PrimaryClusterId *string

	// Removes the user groups that can access this replication group.
	RemoveUserGroups *bool

	// A description for the replication group. Maximum length is 255 characters.
	ReplicationGroupDescription *string

	// Specifies the VPC Security Groups associated with the clusters in the
	// replication group. This parameter can be used only with replication group
	// containing clusters running in an Amazon Virtual Private Cloud (Amazon VPC).
	SecurityGroupIds []string

	// The number of days for which ElastiCache retains automatic node group (shard)
	// snapshots before deleting them. For example, if you set SnapshotRetentionLimit
	// to 5, a snapshot that was taken today is retained for 5 days before being
	// deleted. Important If the value of SnapshotRetentionLimit is set to zero (0),
	// backups are turned off.
	SnapshotRetentionLimit *int32

	// The daily time range (in UTC) during which ElastiCache begins taking a daily
	// snapshot of the node group (shard) specified by SnapshottingClusterId. Example:
	// 05:00-09:00 If you do not specify this parameter, ElastiCache automatically
	// chooses an appropriate time range.
	SnapshotWindow *string

	// The cluster ID that is used as the daily snapshot source for the replication
	// group. This parameter cannot be set for Redis (cluster mode enabled) replication
	// groups.
	SnapshottingClusterId *string

	// A list of user group IDs.
	UserGroupIdsToAdd []string

	// A list of users groups to remove, meaning the users in the group no longer can
	// access thereplication group.
	UserGroupIdsToRemove []string
}

type ModifyReplicationGroupOutput struct {

	// Contains all of the attributes of a specific Redis replication group.
	ReplicationGroup *types.ReplicationGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyReplicationGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyReplicationGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyReplicationGroup{}, middleware.After)
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
	if err = addOpModifyReplicationGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyReplicationGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyReplicationGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "ModifyReplicationGroup",
	}
}
