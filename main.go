package main

import (
	"bank-system/account"
	"bank-system/service"
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("\n--Welcome to Bank of Cheaters--")
	var user account.Account
	for {

		fmt.Println("\nOpen account: 1")
		fmt.Println("Deposit: 2")
		fmt.Println("Withdrawl: 3")
		fmt.Println("Close account: 4")
		fmt.Println("Close Application: Q")
		fmt.Print("What you want to do: ")

		var option string
		fmt.Scanln(&option)
		switch option {
		case "1":
			var name string
			fmt.Print("\nEnter name:")
			fmt.Scanln(&name)
			_, err := strconv.Atoi(name)
			if err == nil {
				fmt.Println("<<<Error name must be number>>>")
				continue
			}
			var age string
			fmt.Print("\nEnter age:")
			fmt.Scanln(&age)
			agenum, err := strconv.Atoi(age)
			if err != nil {
				fmt.Println("<<<Age must be number>>>")
				continue
			}
			if agenum < 18 {
				fmt.Println("<<<must be 18+>>>")
				continue
			}
			var balance string
			fmt.Print("\nInitial deposit ammount:")
			fmt.Scanln(&balance)
			balancenum, err := strconv.Atoi(balance)
			if err != nil {
				fmt.Println("<<<Error balance must be number>>>")
				continue
			}
			if balancenum <= 0 {
				fmt.Println("<<<Invalid ammount>>>")
				continue
			}
			user = account.Account{
				Name:    name,
				Age:     agenum,
				Balance: balancenum,
			}
			fmt.Println("account created ")
		case "2":
			if user.Name == "" {
				fmt.Println("\nOpen account first")
				continue
			}
			var ammount string
			fmt.Print("\nEnter ammount for Deposit:")
			fmt.Scanln(&ammount)
			ammountnum, err := strconv.Atoi(ammount)
			if err != nil {
				fmt.Println("<<<Error Number only>>>")
				continue
			}
			err = service.Deposit(&user, ammountnum)
			if err != nil {
				fmt.Println("<<<Error>>>", err)
				continue
			}
			fmt.Printf("Success %s your balance is: %d\n", user.Name, user.Balance)
		case "3":
			if user.Name == "" {
				fmt.Println("\nOpen account first")
				continue
			}
			fmt.Print("\nEnter ammount for Withdrawl:")
			var ammount string
			fmt.Scanln(&ammount)
			ammountnum, err := strconv.Atoi(ammount)
			if err != nil {
				fmt.Println("<<<Ammount must be number >>>")
				continue
			}
			err = service.Withdrawl(&user, ammountnum)
			if err != nil {
				fmt.Println("<<<Error>>>", err)
				continue
			}
			fmt.Printf("Success %s your balance is %d\n", user.Name, user.Balance)
		case "4":
			if user.Name == "" {
				fmt.Println("\nYou dont have account")
				continue
			}
			fmt.Println("\nYour account will be removed...!")
			var value string
			fmt.Print("\n YES NO:")
			fmt.Scanln(&value)
			switch value {
			case "Yes", "yes":
				user = account.Account{}
				fmt.Println("----Account deleted----")
				continue
			case "No", "no":
				continue
			}

		case "Q", "q":

			fmt.Printf("\n Good bye %s", user.Name)
			return
		default:
			fmt.Println("\n<<<Wrong option >>>")

		}

	}
}
