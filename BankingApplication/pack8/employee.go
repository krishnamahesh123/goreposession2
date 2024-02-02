package pack8

import (
	"BankingApplication/pack7"
	"fmt"
)

type Employee struct {
	Firstname     string
	Lastname      string
	Age           int
	Experience    string
	Qualification string
	Empid         int
}

func InitialiseEmployee() []Employee {
	emp1 := Employee{
		Firstname:     "Anurag",
		Lastname:      "Dullur",
		Age:           24,
		Experience:    "2years",
		Qualification: "MS",
		Empid:         1,
	}
	emp2 := Employee{
		Firstname:     "Haris",
		Lastname:      "Kolagani",
		Age:           24,
		Experience:    "2years",
		Qualification: "MS",
		Empid:         2,
	}
	emp3 := Employee{
		Firstname:     "Durga Mounika",
		Lastname:      "Jagarlamudi",
		Age:           25,
		Experience:    "3years",
		Qualification: "MS",
		Empid:         3,
	}
	var emp = make([]Employee, 0)
	emp = append(emp, emp1)
	emp = append(emp, emp2)
	emp = append(emp, emp3)
	return emp
}
func CustomerInfo(cus []pack7.Customer) {
	count := 0
	inactive := 0
	for i := 0; i < len(cus); i++ {
		fmt.Println(cus[i])
		if cus[i].IsOpen {
			count++
		} else {
			inactive++
		}
	}
	fmt.Printf("The number of open accounts is %d\n", count)
	fmt.Printf("The number of closed accounts is %d\n", inactive)

}
func LoginAsEmployee(employee []Employee, empid int) error {
	for i := 0; i < len(employee); i++ {
		if employee[i].Empid == empid {
			return nil
		}
	}
	return fmt.Errorf("Employee not found")

}
