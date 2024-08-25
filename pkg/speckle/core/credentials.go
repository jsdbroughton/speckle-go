package core

import (
	"encoding/json"
	"os"
	"path/filepath"
)

import (
	"github.com/jsdbroughton/speckle-go/internal/storage"
	"github.com/jsdbroughton/speckle-go/internal/utils"
)

func GetLocalAccounts(basePath string) ([]Account, error) {
	var accounts []Account

	// Try to get accounts from SQLite storage
	sqliteAccounts, err := getAccountsFromSQLite(basePath)
	if err == nil {
		accounts = append(accounts, sqliteAccounts...)
	}

	// Try to get accounts from JSON files
	jsonAccounts, err := getAccountsFromJSON(basePath)
	if err == nil {
		accounts = append(accounts, jsonAccounts...)
	}

	return accounts, nil
}

func getAccountsFromSQLite(basePath string) ([]Account, error) {
	accountStorage, err := storage.NewSQLiteTransport("Accounts", basePath)
	if err != nil {
		return nil, err
	}
	defer func(accountStorage *storage.SQLiteTransport) {
		err := accountStorage.Close()
		if err != nil {

		}
	}(accountStorage)

	objects, err := accountStorage.GetAllObjects()
	if err != nil {
		return nil, err
	}

	var accounts []Account
	for _, obj := range objects {
		var account Account
		if err := json.Unmarshal([]byte(obj.Data), &account); err != nil {
			continue
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func getAccountsFromJSON(basePath string) ([]Account, error) {
	jsonPath, err := utils.GetAccountsFolderPath(basePath)
	if err != nil {
		return nil, err
	}

	if err := os.MkdirAll(jsonPath, os.ModePerm); err != nil {
		return nil, err
	}

	files, err := os.ReadDir(jsonPath)
	if err != nil {
		return nil, err
	}

	var accounts []Account
	for _, file := range files {
		if filepath.Ext(file.Name()) != ".json" {
			continue
		}

		data, err := os.ReadFile(filepath.Join(jsonPath, file.Name()))
		if err != nil {
			continue
		}

		var account Account
		if err := json.Unmarshal(data, &account); err != nil {
			continue
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func GetDefaultAccount(basePath string) (*Account, error) {
	// Implement the core functionality to get the default account
	// This should not include any metrics tracking
	// ...
	return nil, nil
}

func GetAccountFromToken(token, serverURL string) (*Account, error) {
	// Implement the core functionality to get an account from a token
	// This should not include any metrics tracking
	// ...
	return nil, nil
}
