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

func loadAccount() account.Account {
	data, err := os.ReadFile(datafile)
	if err != nil {
		return account.Account{}
	}
	var acc account.Account
	err = json.Unmarshal(data, &acc)
	if err != nil {
		return account.Account{}
	}
	return acc
}
