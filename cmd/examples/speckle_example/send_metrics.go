package speckle_example

import (
	"github.com/jsdbroughton/speckle-go/pkg/speckle/core"
	"github.com/jsdbroughton/speckle-go/pkg/speckle/logging"
)

// Example functions using the action constants
func TrackSDKAction(account *core.Account, name string) {
	logging.Track(logging.ActionSDK, account, map[string]interface{}{"name": name})
}

func TrackReceiveAction(account *core.Account, customProps map[string]interface{}) {
	logging.Track(logging.ActionReceive, account, customProps)
}

func TrackSendAction(account *core.Account, customProps map[string]interface{}) {
	logging.Track(logging.ActionSend, account, customProps)
}
