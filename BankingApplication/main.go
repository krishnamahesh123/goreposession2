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
	var FirstName, LastName, Email string
	var AccountNumber, AccountBalance, accno, depositmoney, withdrawmoney, Fromaccno, Toaccno, TransferAmount, empid int

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
			fmt.Println("Please Enter Account number to login")
			fmt.Scanf("%d\n", &accno)
			err := pack7.LoginAsCustomer(updatedCus, accno)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Logged as a Customer")
				for i := 0; i < 6; i++ {
					fmt.Println("Choose an option:")
					fmt.Println("1. Create New Customer Account")
					fmt.Println("2. Deposit Money")
					fmt.Println("3. Withdraw Money ")
					fmt.Println("4. View AccountBalance of all Customers")
					fmt.Println("5. Transfer Money ")
					fmt.Println("6. Close Account")
					fmt.Print("Enter option (1 or 2 or 3 or 4 or 5 or 6): ")
					scanner.Scan()
					option := scanner.Text()

					switch option {
					case "1":
						fmt.Println("Enter Firstname")
						fmt.Scanf("%s\n", &FirstName)
						fmt.Println("Enter Lastname")
						fmt.Scanf("%s\n", &LastName)
						fmt.Println("Enter Account Number")
						fmt.Scanf("%d\n", &AccountNumber)
						fmt.Println("Enter Account Balance")
						fmt.Scanf("%d\n", &AccountBalance)
						fmt.Println("Enter Email")
						fmt.Scanf("%s\n", &Email)
						updatedCus = pack7.CreateNewAccount(updatedCus, FirstName, LastName, AccountNumber, AccountBalance, Email)
						for i := 0; i < len(updatedCus); i++ {
							if updatedCus[i].Accountnumber == AccountNumber {
								fmt.Printf("New Account created successfully to %s\n", updatedCus[i].Firstname)
							}
						}

					case "2":
						fmt.Println("Enter AccountNumber to deposit")
						fmt.Scanf("%d\n", &accno)
						fmt.Println("Enter Money to deposit")
						fmt.Scanf("%d\n", &depositmoney)
						totalbal := pack7.AddMoney(updatedCus, accno, depositmoney)
						if totalbal == -1 {
							fmt.Println("The given Account Number does not exist in the Bank")
						} else {
							fmt.Printf("The total amount after deposit is %d to the Account Number %d\n", totalbal, accno)
						}
					case "3":
						fmt.Println("Enter AccountNumber to withdraw")
						fmt.Scanf("%d\n", &accno)
						fmt.Println("Enter Money to withdraw")
						fmt.Scanf("%d\n", &withdrawmoney)
						updatedbalance := pack7.WithdrawMoney(updatedCus, accno, withdrawmoney)
						if updatedbalance == -1 {
							fmt.Println("The given Account Number does not exist in the Bank")
						} else {
							fmt.Printf("The total amount after withdraw is %d from the Account Number %d\n", updatedbalance, accno)
						}
					case "4":
						pack7.GetAccountBalance(updatedCus)
					case "5":
						fmt.Println("Enter From Account Number")
						fmt.Println("Enter To Account Number")
						fmt.Scanf("%d%d\n", &Fromaccno, &Toaccno)
						fmt.Println("Enter the amount to transfer")
						fmt.Scanf("%d\n", &TransferAmount)
						pack7.TransferMoney(updatedCus, Fromaccno, Toaccno, TransferAmount)
					case "6":
						fmt.Println("Enter Account number to close")
						fmt.Scanf("%d\n", &accno)
						pack7.CloseAccount(updatedCus, accno)
					default:
						fmt.Println("Enter valid number")
					}
				}
			}
		case "2":
			fmt.Println("Please Enter Employee id to login")
			fmt.Scanf("%d\n", &empid)
			err := pack8.LoginAsEmployee(emp, empid)
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
