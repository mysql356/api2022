package myinterface

import "log"

func (p PermanentEmployee) CalculateSalary() float64 {
	return p.BasicPay + p.Pf
}

func (c ContractEmployee) CalculateSalary() float64 {
	return c.BasicPay
}

func totalExpense(s []SalaryCalculator) float64 {
	exp := float64(0)
	for _, v := range s {
		exp = exp + v.CalculateSalary()
	}

	return exp
}

////////////

func showArray(a [5]interface{}) {
	for k, v := range a {
		log.Println(k, v)
	}
}

func showSlice(a []interface{}) {
	for k, v := range a {
		log.Println(k, v)
	}
}

////////////

func (p Person) Desc() {
	log.Println(p.Name, p.Age)
}

func (a *Address) Desc() {
	log.Println(a.State, a.Country)
}
