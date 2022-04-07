package myinterface

type SalaryCalculator interface {
	CalculateSalary() float64
}

type PermanentEmployee struct {
	EmployeeID string
	BasicPay   float64
	Pf         float64
}

type ContractEmployee struct {
	EmployeeID string
	BasicPay   float64
}

/////////////////
type Describe interface {
	Desc()
}

type Person struct {
	Name string
	Age  int
}

type Address struct {
	State   string
	Country string
}
