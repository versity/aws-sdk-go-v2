// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type BlockerStatus string

// Enum values for BlockerStatus
const (
	BlockerStatusActive   BlockerStatus = "ACTIVE"
	BlockerStatusResolved BlockerStatus = "RESOLVED"
)

// Values returns all known values for BlockerStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BlockerStatus) Values() []BlockerStatus {
	return []BlockerStatus{
		"ACTIVE",
		"RESOLVED",
	}
}

type BlockerType string

// Enum values for BlockerType
const (
	BlockerTypeAutomated BlockerType = "AUTOMATED"
)

// Values returns all known values for BlockerType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (BlockerType) Values() []BlockerType {
	return []BlockerType{
		"AUTOMATED",
	}
}

type ConnectionStatus string

// Enum values for ConnectionStatus
const (
	ConnectionStatusPending   ConnectionStatus = "PENDING"
	ConnectionStatusAvailable ConnectionStatus = "AVAILABLE"
	ConnectionStatusError     ConnectionStatus = "ERROR"
)

// Values returns all known values for ConnectionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConnectionStatus) Values() []ConnectionStatus {
	return []ConnectionStatus{
		"PENDING",
		"AVAILABLE",
		"ERROR",
	}
}

type ProviderType string

// Enum values for ProviderType
const (
	ProviderTypeBitbucket              ProviderType = "Bitbucket"
	ProviderTypeGithub                 ProviderType = "GitHub"
	ProviderTypeGithubEnterpriseServer ProviderType = "GitHubEnterpriseServer"
	ProviderTypeGitlab                 ProviderType = "GitLab"
)

// Values returns all known values for ProviderType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ProviderType) Values() []ProviderType {
	return []ProviderType{
		"Bitbucket",
		"GitHub",
		"GitHubEnterpriseServer",
		"GitLab",
	}
}

type RepositorySyncStatus string

// Enum values for RepositorySyncStatus
const (
	RepositorySyncStatusFailed     RepositorySyncStatus = "FAILED"
	RepositorySyncStatusInitiated  RepositorySyncStatus = "INITIATED"
	RepositorySyncStatusInProgress RepositorySyncStatus = "IN_PROGRESS"
	RepositorySyncStatusSucceeded  RepositorySyncStatus = "SUCCEEDED"
	RepositorySyncStatusQueued     RepositorySyncStatus = "QUEUED"
)

// Values returns all known values for RepositorySyncStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RepositorySyncStatus) Values() []RepositorySyncStatus {
	return []RepositorySyncStatus{
		"FAILED",
		"INITIATED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"QUEUED",
	}
}

type ResourceSyncStatus string

// Enum values for ResourceSyncStatus
const (
	ResourceSyncStatusFailed     ResourceSyncStatus = "FAILED"
	ResourceSyncStatusInitiated  ResourceSyncStatus = "INITIATED"
	ResourceSyncStatusInProgress ResourceSyncStatus = "IN_PROGRESS"
	ResourceSyncStatusSucceeded  ResourceSyncStatus = "SUCCEEDED"
)

// Values returns all known values for ResourceSyncStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceSyncStatus) Values() []ResourceSyncStatus {
	return []ResourceSyncStatus{
		"FAILED",
		"INITIATED",
		"IN_PROGRESS",
		"SUCCEEDED",
	}
}

type SyncConfigurationType string

// Enum values for SyncConfigurationType
const (
	SyncConfigurationTypeCfnStackSync SyncConfigurationType = "CFN_STACK_SYNC"
)

// Values returns all known values for SyncConfigurationType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SyncConfigurationType) Values() []SyncConfigurationType {
	return []SyncConfigurationType{
		"CFN_STACK_SYNC",
	}
}
