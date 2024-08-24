package logging

import (
	"sync"

	"github.com/jsdbroughton/speckle-go/pkg/speckle/core"
)

var (
	track          = true
	hostApp        = "go"
	hostAppVersion = "go 1.x" // Update this with the appropriate version

	metricsTracker *MetricsTracker
	once           sync.Once
)

type MetricsTracker struct {
	lastUser   string
	lastServer string
	// Add other necessary fields
}

func Disable() {
	track = false
}

func Enable() {
	track = true
}

func SetHostApp(app string, version string) {
	hostApp = app
	if version != "" {
		hostAppVersion = version
	}
}

func Track(action string, account *core.Account, customProps map[string]interface{}) {
	if !track {
		return
	}

	once.Do(func() {
		metricsTracker = &MetricsTracker{}
	})

	// Implement tracking logic here
	// This could involve sending data to a service, writing to logs, etc.
	// For now, we'll just print the action
	println("Tracked action:", action)
}

func InitializeTracker(account *core.Account) {
	once.Do(func() {
		metricsTracker = &MetricsTracker{}
	})

	if account != nil && account.UserInfo.Email != "" {
		metricsTracker.lastUser = account.UserInfo.Email
	}
	if account != nil && account.ServerInfo.URL != "" {
		metricsTracker.lastServer = account.ServerInfo.URL
	}
}
