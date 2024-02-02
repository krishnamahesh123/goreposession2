package pack7

import "fmt"

type Customer struct {
	Firstname      string
	Lastname       string
	Accountnumber  int
	Accountbalance int
	Email          string
	IsOpen         bool
}

func InitialiseCustomer() []Customer {
	cust1 := Customer{
		Firstname:      "Krishna",
		Lastname:       "Mahesh",
		Accountnumber:  1234,
		Accountbalance: 100,
		Email:          "krishnamaheshecea@gmail.com",
		IsOpen:         true,
	}
	cust2 := Customer{
		Firstname:      "Manikanta",
		Lastname:       "Venkata",
		Accountnumber:  5678,
		Accountbalance: 200,
		Email:          "manikantakothuri@gmail.com",
		IsOpen:         true,
	}
	cust3 := Customer{
		Firstname:      "Rambabu",
		Lastname:       "Simhadri",
		Accountnumber:  1486,
		Accountbalance: 300,
		Email:          "rambabusimhadri@gmail.com",
		IsOpen:         true,
	}
	var cus = make([]Customer, 4)
	cus = append(cus, cust1)
	cus = append(cus, cust2)
	cus = append(cus, cust3)
	return cus
}

func CreateNewAccount(cus []Customer, fname string, lname string, accno int, accbal int, email string) []Customer {
	cust4 := Customer{
		Firstname:      fname,
		Lastname:       lname,
		Accountnumber:  accno,
		Accountbalance: accbal,
		Email:          email,
		IsOpen:         true,
	}
	cus = append(cus, cust4)
	return cus
}
func AddMoney(cus []Customer, accno int, depomoney int) int {
	for i := 0; i < len(cus); i++ {
		if cus[i].Accountnumber == accno {
			balance := cus[i].Accountbalance
			fmt.Printf("The amount want to deposit is %d to the AccountNumber %d\n", depomoney, cus[i].Accountnumber)
			updated_balance := balance + depomoney
			cus[i].Accountbalance = updated_balance
			return updated_balance
		}
	}
	return -1
}
func WithdrawMoney(cus []Customer, accno int, withdrawmoney int) int {
	for i := 0; i < len(cus); i++ {
		if cus[i].Accountnumber == accno {
			balance := cus[i].Accountbalance
			fmt.Printf("The amount want to withdraw is %d from the AccountNumber %d\n", withdrawmoney, cus[i].Accountnumber)
			updated_balance := CheckBalance(balance, withdrawmoney)
			cus[i].Accountbalance = updated_balance
			return updated_balance
		}
	}
	return -1
}
func CheckBalance(balance int, withdraw int) int {
	if balance > withdraw {
		updated_balance := balance - withdraw
		return updated_balance
	} else {
		fmt.Println("Unable to withdraw/Transfer money because of low balance")
		return 0
	}

}

func GetAccountBalance(cus []Customer) {
	fmt.Println("The following is the AccountBalance of all Customers")
	for i := 0; i < len(cus); i++ {
		fmt.Printf("The Account balance is %d for the Account Number %d\n", cus[i].Accountbalance, cus[i].Accountnumber)
	}

}
func TransferMoney(cus []Customer, Fromaccno int, Toaccno int, TransferAmount int) {
	for i := 0; i < len(cus); i++ {
		if cus[i].Accountnumber == Fromaccno {
			fmt.Printf("Transfer Money of %d rupees from Account Number %d to the Account Number %d\n", TransferAmount, Fromaccno, Toaccno)
			balance := cus[i].Accountbalance
			updated_balance := CheckBalance(balance, TransferAmount)
			cus[i].Accountbalance = updated_balance
			fmt.Printf("The total balance in Account number %d after transfer of %d rupees to Account Number %d  is %d\n", Fromaccno, TransferAmount, Toaccno, updated_balance)
		}
		if cus[i].Accountnumber == Toaccno {
			bal := cus[i].Accountbalance
			cus[i].Accountbalance = bal + TransferAmount
			fmt.Printf("The total balance in Account number %d after transfer of %d rupees from Account Number %d is %d\n", Toaccno, TransferAmount, Fromaccno, cus[i].Accountbalance)

		}

	}

}
func CloseAccount(cus []Customer, accno int) {
	for i := 0; i < len(cus); i++ {
		if cus[i].Accountnumber == accno {
			check := cus[i].IsOpen
			if !check {
				fmt.Println("Account is already closed")
			} else {
				check = false
				cus[i].IsOpen = check
				fmt.Printf("Account Number %d is closed Successfully \n", cus[i].Accountnumber)
			}
		}

	}

}
func CheckAccountNumber(customer []Customer, accno int) bool {
	for i := 0; i < len(customer); i++ {
		if customer[i].Accountnumber == accno {
			return true
		}
	}
	return false
}

func LoginAsCustomer(customer []Customer, accountnumber int) error {
	for i := 0; i < len(customer); i++ {
		if customer[i].Accountnumber == accountnumber {
			return nil
		}
	}
	return fmt.Errorf("Customer not found")

}
