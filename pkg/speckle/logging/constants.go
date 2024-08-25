package logging

import (
	"fmt"
	"runtime"
)

const (
	// TrackDefault is the default tracking state
	TrackDefault = true

	// HostAppDefault is the default host application
	HostAppDefault = "go"

	AnalyticsUrl   = "https://analytics.speckle.systems/track?ip=1"
	AnalyticsToken = "acd87c5a50b56df91a795e999812a3a4"
	MaxRetries     = 3
	RetryDelay     = 1

	// Actions
	ActionSDK       = "SDK Action"
	ActionConnector = "Connector Action"
	ActionReceive   = "Receive"
	ActionSend      = "Send"

	// Legacy actions (not in use since 2.15)
	ActionAccounts      = "Get Local Accounts"
	ActionBranch        = "Branch Action"
	ActionClient        = "Speckle Client"
	ActionCommit        = "Commit Action"
	ActionDeserialize   = "serialization/deserialize"
	ActionInvite        = "Invite Action"
	ActionOtherUser     = "Other User Action"
	ActionPermission    = "Permission Action"
	ActionSerialize     = "serialization/serialize"
	ActionServer        = "Server Action"
	ActionStream        = "Stream Action"
	ActionStreamWrapper = "Stream Wrapper"
	ActionUser          = "User Action"
)

var (
	// HostAppVersionDefault is the default host application version
	HostAppVersionDefault = fmt.Sprintf("go %s", runtime.Version())

	// Platforms maps GOOS values to human-readable platform names
	Platforms = map[string]string{
		"windows": "Windows",
		"darwin":  "Mac OS X",
		"linux":   "Linux",
	}
)
