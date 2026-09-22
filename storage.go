package main

import (
	"bank-system/account"
	"encoding/json"
	"os"
)

const datafile = "bank_data.json"

func saveAccount(acc account.Account) error {
	data, err := json.Marshal(acc)
	if err != nil {
		return err
	}
	return os.WriteFile(datafile, data, 0644)
}
