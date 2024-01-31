package main

import (
	"BankingApplication/pack7"
	"BankingApplication/pack8"
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 5; i++ {
		fmt.Println("Choose an option:")
		fmt.Println("1. Customer")
		fmt.Println("2. Employee")
		fmt.Print("Enter option (1 or 2): ")
		scanner.Scan()
		choice := scanner.Text()
		updatedCus := pack7.InitialiseCustomer()
		emp := pack8.InitialiseEmployee()

		switch choice {
		case "1":
			err := pack7.LoginAsCustomer(updatedCus, 1234)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Logged as a Customer")
				updatedCus := pack7.CreateNewAccount(updatedCus)
				fmt.Printf("New Account created successfully to %s\n", updatedCus[3].Firstname)
				total_balance := pack7.AddMoney(updatedCus)
				fmt.Printf("The total amount after deposit is %d to the Account Number %d\n", total_balance, updatedCus[0].Accountnumber)
				total_bal := pack7.WithdrawMoney(updatedCus)
				fmt.Printf("The total amount after withdraw is %d from the Account Number %d\n", total_bal, updatedCus[1].Accountnumber)
				pack7.GetAccountBalance(updatedCus)
				pack7.TransferMoney(updatedCus)
				pack7.CloseAccount(updatedCus)
			}
		case "2":
			err := pack8.LoginAsEmployee(emp, 1)
			if err != nil {
				fmt.Println(err)
			} else {
				pack8.CustomerInfo(updatedCus)
				fmt.Println("The following are the employee details")
				for _, details := range emp {
					fmt.Println(details)
				}
			}
		default:
			fmt.Println("Please enter a valid input ")
		}
	}

}
