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
	cus[0] = cust1
	cus[1] = cust2
	cus[2] = cust3
	return cus
}

func CreateNewAccount(cus []Customer) []Customer {
	cust4 := Customer{
		Firstname:     "Pallavi",
		Lastname:      "Kastala",
		Accountnumber: 1789,
		Email:         "pallavikastala@gmail.com",
		IsOpen:        true,
	}
	cus[3] = cust4
	return cus

}
func AddMoney(cus []Customer) int {
	balance := cus[0].Accountbalance
	deposit := 100
	fmt.Printf("The amount want to deposit is %d to the AccountNumber %d\n", deposit, cus[0].Accountnumber)
	updated_balance := balance + deposit
	cus[0].Accountbalance = updated_balance
	return updated_balance
}
func WithdrawMoney(cus []Customer) int {
	balance := cus[1].Accountbalance
	withdraw := 50
	fmt.Printf("The amount want to withdraw is %d from the AccountNumber %d\n", withdraw, cus[1].Accountnumber)
	updated_balance := CheckBalance(balance, withdraw)
	cus[1].Accountbalance = updated_balance
	return updated_balance

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
func TransferMoney(cus []Customer) {
	TransferAmount := 50
	fmt.Printf("Transfer Money of %d rupees from Account Number %d to the Account Number %d\n", TransferAmount, cus[2].Accountnumber, cus[3].Accountnumber)
	balance := cus[2].Accountbalance
	updated_balance := CheckBalance(balance, TransferAmount)
	cus[2].Accountbalance = updated_balance
	fmt.Printf("The total balance in Account number %d after transfer of %d rupees to Account Number %d  is %d\n", cus[2].Accountnumber, TransferAmount, cus[3].Accountnumber, updated_balance)
	bal := cus[3].Accountbalance
	cus[3].Accountbalance = bal + TransferAmount
	fmt.Printf("The total balance in Account number %d after transfer of %d rupees from Account Number %d is %d\n", cus[3].Accountnumber, TransferAmount, cus[2].Accountnumber, cus[3].Accountbalance)
}
func CloseAccount(cus []Customer) {
	check := cus[3].IsOpen
	if !check {
		fmt.Println("Account is already closed")
	} else {
		check = false
		cus[3].IsOpen = check
		fmt.Printf("Account Number %d is closed Successfully \n", cus[3].Accountnumber)
	}

}

func LoginAsCustomer(customer []Customer, accountnumber int) error {
	for i := 0; i < len(customer); i++ {
		if customer[i].Accountnumber == accountnumber {
			return nil
		}
	}
	return fmt.Errorf("Customer not found")

}
