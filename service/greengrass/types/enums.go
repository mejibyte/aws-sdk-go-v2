// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type BulkDeploymentStatus string

// Enum values for BulkDeploymentStatus
const (
	BulkDeploymentStatusInitializing BulkDeploymentStatus = "Initializing"
	BulkDeploymentStatusRunning      BulkDeploymentStatus = "Running"
	BulkDeploymentStatusCompleted    BulkDeploymentStatus = "Completed"
	BulkDeploymentStatusStopping     BulkDeploymentStatus = "Stopping"
	BulkDeploymentStatusStopped      BulkDeploymentStatus = "Stopped"
	BulkDeploymentStatusFailed       BulkDeploymentStatus = "Failed"
)

type DeploymentType string

// Enum values for DeploymentType
const (
	DeploymentTypeNewdeployment        DeploymentType = "NewDeployment"
	DeploymentTypeRedeployment         DeploymentType = "Redeployment"
	DeploymentTypeResetdeployment      DeploymentType = "ResetDeployment"
	DeploymentTypeForceresetdeployment DeploymentType = "ForceResetDeployment"
)

type EncodingType string

// Enum values for EncodingType
const (
	EncodingTypeBinary EncodingType = "binary"
	EncodingTypeJson   EncodingType = "json"
)

type FunctionIsolationMode string

// Enum values for FunctionIsolationMode
const (
	FunctionIsolationModeGreengrasscontainer FunctionIsolationMode = "GreengrassContainer"
	FunctionIsolationModeNocontainer         FunctionIsolationMode = "NoContainer"
)

type LoggerComponent string

// Enum values for LoggerComponent
const (
	LoggerComponentGreengrasssystem LoggerComponent = "GreengrassSystem"
	LoggerComponentLambda           LoggerComponent = "Lambda"
)

type LoggerLevel string

// Enum values for LoggerLevel
const (
	LoggerLevelDebug LoggerLevel = "DEBUG"
	LoggerLevelInfo  LoggerLevel = "INFO"
	LoggerLevelWarn  LoggerLevel = "WARN"
	LoggerLevelError LoggerLevel = "ERROR"
	LoggerLevelFatal LoggerLevel = "FATAL"
)

type LoggerType string

// Enum values for LoggerType
const (
	LoggerTypeFilesystem    LoggerType = "FileSystem"
	LoggerTypeAwscloudwatch LoggerType = "AWSCloudWatch"
)

type Permission string

// Enum values for Permission
const (
	PermissionRo Permission = "ro"
	PermissionRw Permission = "rw"
)

type SoftwareToUpdate string

// Enum values for SoftwareToUpdate
const (
	SoftwareToUpdateCore      SoftwareToUpdate = "core"
	SoftwareToUpdateOta_agent SoftwareToUpdate = "ota_agent"
)

type UpdateAgentLogLevel string

// Enum values for UpdateAgentLogLevel
const (
	UpdateAgentLogLevelNone    UpdateAgentLogLevel = "NONE"
	UpdateAgentLogLevelTrace   UpdateAgentLogLevel = "TRACE"
	UpdateAgentLogLevelDebug   UpdateAgentLogLevel = "DEBUG"
	UpdateAgentLogLevelVerbose UpdateAgentLogLevel = "VERBOSE"
	UpdateAgentLogLevelInfo    UpdateAgentLogLevel = "INFO"
	UpdateAgentLogLevelWarn    UpdateAgentLogLevel = "WARN"
	UpdateAgentLogLevelError   UpdateAgentLogLevel = "ERROR"
	UpdateAgentLogLevelFatal   UpdateAgentLogLevel = "FATAL"
)

type UpdateTargetsArchitecture string

// Enum values for UpdateTargetsArchitecture
const (
	UpdateTargetsArchitectureArmv6l  UpdateTargetsArchitecture = "armv6l"
	UpdateTargetsArchitectureArmv7l  UpdateTargetsArchitecture = "armv7l"
	UpdateTargetsArchitectureX86_64  UpdateTargetsArchitecture = "x86_64"
	UpdateTargetsArchitectureAarch64 UpdateTargetsArchitecture = "aarch64"
)

type UpdateTargetsOperatingSystem string

// Enum values for UpdateTargetsOperatingSystem
const (
	UpdateTargetsOperatingSystemUbuntu       UpdateTargetsOperatingSystem = "ubuntu"
	UpdateTargetsOperatingSystemRaspbian     UpdateTargetsOperatingSystem = "raspbian"
	UpdateTargetsOperatingSystemAmazon_linux UpdateTargetsOperatingSystem = "amazon_linux"
	UpdateTargetsOperatingSystemOpenwrt      UpdateTargetsOperatingSystem = "openwrt"
)