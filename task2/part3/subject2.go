package part3

import "fmt"

type Person struct{
	Name string
	Age int

}

type Employee struct {
	Person
	EmployeeId string
}

func (e *Employee) PrintInfo() {
	fmt.Printf("Name: %s, Age: %d, EmployeeId: %s\n", e.Name, e.Age, e.EmployeeId)
}
