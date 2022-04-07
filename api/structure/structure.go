package structure

import (
	cm "api2022/common/struct/models"
	"api2022/util"
	"fmt"
	"log"
	"net/http"
)

func Init() {

	http.HandleFunc("/api/struct/basic", basic)
	http.HandleFunc("/api/struct/alias", alias)
	http.HandleFunc("/api/struct/pointer", pointer)
	http.HandleFunc("/api/struct/function", function)
	http.HandleFunc("/api/struct/method", method)
	http.HandleFunc("/api/struct/encapsulation", encapsulation)
	http.HandleFunc("/api/struct/anonymous", anonymous)
	http.HandleFunc("/api/struct/changes", changes)
	http.HandleFunc("/api/struct/structReturn", structReturn)
	http.HandleFunc("/api/struct/nested", nested)
	http.HandleFunc("/api/struct/promoted", promoted)
	http.HandleFunc("/api/struct/overloading", overloading)

}

func basic(w http.ResponseWriter, r *http.Request) {
	// V-R (without ref)
	p1 := cm.Person{FirstName: "Ak", LastName: "Singh"} //same type - so FirstName/LastName
	fmt.Fprintf(w, "%+v \n", p1.Get())

	// R-R (with ref)
	p2 := &cm.Person{FirstName: "Bk", LastName: "Singh"}
	fmt.Fprintf(w, "%+v \n", p2.Get())

	// V-R (without ref)
	p3 := cm.Person{}

	p3.FirstName = "Ck"
	p3.LastName = "Singh"
	fmt.Fprintf(w, "%+v \n", p3.Get())

	// R-R (with ref)
	var q cm.Person
	p4 := &q

	p4.FirstName = "Dk"
	p4.LastName = "Singh"
	fmt.Fprintf(w, "%+v \n", p4.Get())

	// R-R (with ref)
	p5 := new(cm.Person)

	p5.FirstName = "Ek"
	p5.LastName = "Singh"
	fmt.Fprintf(w, "%+v \n", p5.Get())

	util.Render(w, p5)
}

func alias(w http.ResponseWriter, r *http.Request) {
	n1 := cm.MyInt(5)
	n2 := cm.MyInt(10)
	sum := n1.Add(n2)

	fmt.Fprintf(w, "%+v \n", sum)
}

func pointer(w http.ResponseWriter, r *http.Request) {
	e1 := cm.Employee{Name: "aa", Age: 20} //field must be need, because type is not in same file
	e2 := &cm.Employee{Name: "bb", Age: 30}

	log.Println(e1.Name, e1.Age)
	log.Println((e2).Name, e2.Age)
	log.Println((*e2).Name, e2.Age)
	util.Render(w, e2)
}

func function(w http.ResponseWriter, rs *http.Request) {
	r := cm.Reactangle{
		Lenght: 10,
		Width:  5,
	}
	p := &r

	// V-V
	area1 := cm.Area(r) //cm.Area(*p) *************Attention*************
	fmt.Fprintf(w, "Area = %d \n", area1)

	// R-V - NOT ALLOWED
	//cm.Area(p) or cm.Area(&r)

	// R-R
	area2 := cm.Areap(p) // cm.Areap(&r)
	fmt.Fprintf(w, "Area = %d \n", area2)

	// V-R NOT ALLOWED
	//cm.Areap(r) or cm.Areap(*p)

	util.Render(w, r)
}

func method(w http.ResponseWriter, rs *http.Request) {
	r := cm.Reactangle{
		Lenght: 10,
		Width:  5,
	}
	p := &r

	// V-V
	area1 := r.Area() //cm.Area(*p) *************Attention*************
	fmt.Fprintf(w, "Area = %d \n", area1)

	// R-V
	area11 := p.Area()
	fmt.Fprintf(w, "Area = %d \n", area11)

	// R-R
	area2 := p.Areap() // cm.Areap(&r)
	fmt.Fprintf(w, "Area = %d \n", area2)

	// V-R
	area21 := r.Areap()
	fmt.Fprintf(w, "Area = %d \n", area21)

	util.Render(w, r)
}

func encapsulation(w http.ResponseWriter, r *http.Request) {
	s1 := new(cm.Student)
	s1.Set("Ak", "Sharma")
	fmt.Fprintf(w, "%+v \n", s1.Get())

	util.Render(w, s1)
}

func anonymous(w http.ResponseWriter, r *http.Request) {

	//single
	e1 := struct {
		Fn string //fn, ln => can be exported so not printed
		Ln string
	}{
		"ak",
		"singh",
	}
	str := fmt.Sprintf("%s %s", e1.Fn, e1.Ln)
	fmt.Fprintf(w, "%+v \n", str)
	util.Render(w, e1)
	fmt.Fprintf(w, "\n")
	//multiple
	e2 := []struct {
		Fn string
		Ln string
	}{
		{"ak", "singh"},
		{"bk", "singh"},
		{"ck", "singh"},
	}

	log.Println(e2)
	util.Render(w, e2)
}

func changes(w http.ResponseWriter, r *http.Request) {
	e := cm.Employee{
		Name: "Rahul",
		Age:  20,
	}

	//before changes
	log.Println(e)

	//temp changes
	e.Tch("tt")

	//after temp changes
	log.Println(e)

	//permanet changes
	e.Pch("pp") //(&e).Pch("PP")

	//after changes
	log.Println(e)

	util.Render(w, e)
}

func structReturn(w http.ResponseWriter, r *http.Request) {

	e1 := cm.Employee{
		Name: "aa",
		Age:  20,
	}

	util.Render(w, e1)

	obj1, _ := e1.ModifyByVal(w)
	util.Render(w, obj1)

	fmt.Fprintf(w, "\n")
	///////////////////////

	e2 := cm.Employee{
		Name: "xx",
		Age:  21,
	}

	util.Render(w, e2)

	obj2, _ := e2.ModifyByRef(w)
	util.Render(w, obj2)

	//{"Name":"aa","Age":20}{"Name":"bb","Age":30}{"Name":"aa","Age":60}
	//{"Name":"xx","Age":21}{"Name":"yy","Age":31}{"Name":"yy","Age":61}

}

func nested(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "nested")
	var e Employee
	e.Name = "aa"
	e.Age = 20
	e.Add.City = "New Delhi"
	e.Add.State = "Delhi"
	e.Cont = []Contact{{"aa@gmail.com", 1001},
		{"bb@gmail.com", 1002},
		{"cc@gmail.com", 1003},
	}
	log.Println(e)

	util.Render(w, e)
}

func promoted(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "promoted")
	var p Person
	p.Name = "xx"
	p.Age = 20
	p.City = "New Delhi"
	p.State = "Delhi"

	p.Cont = []Contact{{"xx@gmail.com", 2001},
		{"yy@gmail.com", 2002},
		{"zz@gmail.com", 2003},
	}
	log.Println(p)

	util.Render(w, p)
}

func overloading(w http.ResponseWriter, rs *http.Request) {
	//fmt.Fprintf(w, "promoted")

	r := cm.Reactangle{
		Lenght: 10,
		Width:  20,
	}

	c := cm.Circle{
		Radius: 12,
	}

	log.Println(r.CalculateArea(), c.CalculateArea())

	result := struct {
		Circle     float64
		Reactangle int
	}{
		c.CalculateArea(),
		r.CalculateArea(),
	}

	util.Render(w, result)

}
