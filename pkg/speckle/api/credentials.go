package api

import (
	"github.com/jsdbroughton/speckle-go/pkg/speckle/core"
	"github.com/jsdbroughton/speckle-go/pkg/speckle/logging"
)

func GetLocalAccounts(basePath string) ([]core.Account, error) {
	accounts, err := core.GetLocalAccounts(basePath)
	if err != nil {
		return nil, err
	}

	var defaultAccount *core.Account
	if len(accounts) > 0 {
		for _, acc := range accounts {
			if acc.IsDefault {
				defaultAccount = &acc
				break
			}
		}
		if defaultAccount == nil {
			defaultAccount = &accounts[0]
		}
	}

	logging.Track("SDK Action", defaultAccount, map[string]interface{}{"name": "Get Local Accounts"})
	return accounts, nil
}

func GetDefaultAccount(basePath string) (*core.Account, error) {
	account, err := core.GetDefaultAccount(basePath)
	if err != nil {
		return nil, err
	}

	if account != nil {
		logging.InitializeTracker(account)
	}
	return account, nil
}

func GetAccountFromToken(token, serverURL string) (*core.Account, error) {
	account, err := core.GetAccountFromToken(token, serverURL)
	if err != nil {
		return nil, err
	}

	logging.Track("SDK Action", account, map[string]interface{}{"name": "Get Account From Token"})
	return account, nil
}
