package service

import (
	"bank-system/account"
	"errors"
)

func Withdrawl(p *account.Account, ammount int) error {

	if ammount <= 0 {
		return errors.New("Invalid ammount")
	}
	if ammount > p.Balance {
		return errors.New("Insufficient balance")
	}
	p.Balance -= ammount
	return nil
}

func Deposit(P *account.Account, ammount int) error {
	if ammount <= 0 {
		return errors.New("Invalid ammount")
	}
	P.Balance += ammount
	return nil
}
