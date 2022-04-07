package myinterface

import (
	"api2022/util"
	"fmt"
	"log"
	"net/http"
)

func Init() {
	http.HandleFunc("/api/myinterface/basic", basic)
	http.HandleFunc("/api/myinterface/ex1", ex1)
	http.HandleFunc("/api/myinterface/ex2", ex2)

}

func basic(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "interface...")

	var a interface{}
	a = "aa"
	log.Println(a.(string))

	var s interface{} = "aa"
	log.Println(s.(string))

	var b interface{} = 10

	log.Println(b.(int))

	var c interface{} = 10.5
	log.Println(c.(float64))

	d := [5]interface{}{1, "aa", 1.5, true, nil}
	showArray(d)

	e := []interface{}{2, "bb", 2.5, false, nil}
	showSlice(e)

}

func ex1(w http.ResponseWriter, r *http.Request) {

	p1 := PermanentEmployee{"1001", 5000, 20}
	p2 := PermanentEmployee{"1002", 6000, 30}
	p3 := ContractEmployee{"1003", 2000}

	emp := []SalaryCalculator{p1, p2, p3}

	result := struct {
		TotalExpense float64
	}{
		totalExpense(emp),
	}

	util.Render(w, result)

}

func ex2(w http.ResponseWriter, r *http.Request) {
	var d1 Describe
	p1 := Person{"aa", 20}
	d1 = p1
	//d1 = &p1 //Also allowed
	d1.Desc()

	//////////
	var d2 Describe
	a1 := Address{"Delhi", "India"}
	d2 = &a1
	//d2 = a1 //Not allowed - needed only pointer receiver
	d2.Desc()
}
