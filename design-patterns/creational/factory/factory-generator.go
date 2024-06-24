package main

import "fmt"

type EmployeeG struct {
	Name, Position string
	AnnualIncome   int
}

// what if we want factories for specific roles?

// functional approach
func NewEmployeeFactory(position string,
	annualIncome int) func(name string) *EmployeeG {
	return func(name string) *EmployeeG {
		return &EmployeeG{name, position, annualIncome}
	}
}

// structural approach
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func NewEmployeeFactory2(position string,
	annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

func (f *EmployeeFactory) Create(name string) *EmployeeG {
	return &EmployeeG{name, f.Position, f.AnnualIncome}
}

func main2() {
	developerFactory := NewEmployeeFactory("Developer", 60000)
	managerFactory := NewEmployeeFactory("Manager", 80000)

	developer := developerFactory("Adam")
	fmt.Println(developer)

	manager := managerFactory("Jane")
	fmt.Println(manager)

	bossFactory := NewEmployeeFactory2("CEO", 100000)
	// can modify post-hoc
	bossFactory.AnnualIncome = 110000
	boss := bossFactory.Create("Sam")
	fmt.Println(boss)
}
