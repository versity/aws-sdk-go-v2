// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AccessType string

// Enum values for AccessType
const (
	AccessTypeS3Sigv4                   AccessType = "S3_SIGV4"
	AccessTypeSecretsManagerAccessToken AccessType = "SECRETS_MANAGER_ACCESS_TOKEN"
)

// Values returns all known values for AccessType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (AccessType) Values() []AccessType {
	return []AccessType{
		"S3_SIGV4",
		"SECRETS_MANAGER_ACCESS_TOKEN",
	}
}

type AdMarkupType string

// Enum values for AdMarkupType
const (
	AdMarkupTypeDaterange      AdMarkupType = "DATERANGE"
	AdMarkupTypeScte35Enhanced AdMarkupType = "SCTE35_ENHANCED"
)

// Values returns all known values for AdMarkupType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AdMarkupType) Values() []AdMarkupType {
	return []AdMarkupType{
		"DATERANGE",
		"SCTE35_ENHANCED",
	}
}

type ChannelState string

// Enum values for ChannelState
const (
	ChannelStateRunning ChannelState = "RUNNING"
	ChannelStateStopped ChannelState = "STOPPED"
)

// Values returns all known values for ChannelState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ChannelState) Values() []ChannelState {
	return []ChannelState{
		"RUNNING",
		"STOPPED",
	}
}

type FillPolicy string

// Enum values for FillPolicy
const (
	FillPolicyFullAvailOnly FillPolicy = "FULL_AVAIL_ONLY"
	FillPolicyPartialAvail  FillPolicy = "PARTIAL_AVAIL"
)

// Values returns all known values for FillPolicy. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (FillPolicy) Values() []FillPolicy {
	return []FillPolicy{
		"FULL_AVAIL_ONLY",
		"PARTIAL_AVAIL",
	}
}

type LogType string

// Enum values for LogType
const (
	LogTypeAsRun LogType = "AS_RUN"
)

// Values returns all known values for LogType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (LogType) Values() []LogType {
	return []LogType{
		"AS_RUN",
	}
}

type MessageType string

// Enum values for MessageType
const (
	MessageTypeSpliceInsert MessageType = "SPLICE_INSERT"
	MessageTypeTimeSignal   MessageType = "TIME_SIGNAL"
)

// Values returns all known values for MessageType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (MessageType) Values() []MessageType {
	return []MessageType{
		"SPLICE_INSERT",
		"TIME_SIGNAL",
	}
}

type Mode string

// Enum values for Mode
const (
	ModeOff            Mode = "OFF"
	ModeBehindLiveEdge Mode = "BEHIND_LIVE_EDGE"
	ModeAfterLiveEdge  Mode = "AFTER_LIVE_EDGE"
)

// Values returns all known values for Mode. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Mode) Values() []Mode {
	return []Mode{
		"OFF",
		"BEHIND_LIVE_EDGE",
		"AFTER_LIVE_EDGE",
	}
}

type Operator string

// Enum values for Operator
const (
	OperatorEquals Operator = "EQUALS"
)

// Values returns all known values for Operator. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Operator) Values() []Operator {
	return []Operator{
		"EQUALS",
	}
}

type OriginManifestType string

// Enum values for OriginManifestType
const (
	OriginManifestTypeSinglePeriod OriginManifestType = "SINGLE_PERIOD"
	OriginManifestTypeMultiPeriod  OriginManifestType = "MULTI_PERIOD"
)

// Values returns all known values for OriginManifestType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OriginManifestType) Values() []OriginManifestType {
	return []OriginManifestType{
		"SINGLE_PERIOD",
		"MULTI_PERIOD",
	}
}

type PlaybackMode string

// Enum values for PlaybackMode
const (
	PlaybackModeLoop   PlaybackMode = "LOOP"
	PlaybackModeLinear PlaybackMode = "LINEAR"
)

// Values returns all known values for PlaybackMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PlaybackMode) Values() []PlaybackMode {
	return []PlaybackMode{
		"LOOP",
		"LINEAR",
	}
}

type RelativePosition string

// Enum values for RelativePosition
const (
	RelativePositionBeforeProgram RelativePosition = "BEFORE_PROGRAM"
	RelativePositionAfterProgram  RelativePosition = "AFTER_PROGRAM"
)

// Values returns all known values for RelativePosition. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RelativePosition) Values() []RelativePosition {
	return []RelativePosition{
		"BEFORE_PROGRAM",
		"AFTER_PROGRAM",
	}
}

type ScheduleEntryType string

// Enum values for ScheduleEntryType
const (
	ScheduleEntryTypeProgram     ScheduleEntryType = "PROGRAM"
	ScheduleEntryTypeFillerSlate ScheduleEntryType = "FILLER_SLATE"
)

// Values returns all known values for ScheduleEntryType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ScheduleEntryType) Values() []ScheduleEntryType {
	return []ScheduleEntryType{
		"PROGRAM",
		"FILLER_SLATE",
	}
}

type Tier string

// Enum values for Tier
const (
	TierBasic    Tier = "BASIC"
	TierStandard Tier = "STANDARD"
)

// Values returns all known values for Tier. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Tier) Values() []Tier {
	return []Tier{
		"BASIC",
		"STANDARD",
	}
}

type Type string

// Enum values for Type
const (
	TypeDash Type = "DASH"
	TypeHls  Type = "HLS"
)

// Values returns all known values for Type. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Type) Values() []Type {
	return []Type{
		"DASH",
		"HLS",
	}
}
