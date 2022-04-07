package routers

import (
	"fmt"
	"net/http"

	"api2022/api/customer"
	"api2022/api/employee"
	"api2022/api/myinterface"
	"api2022/api/structure"
)

func SetupRouter() {
	http.HandleFunc("/", index)
	http.HandleFunc("/home", home)
	employee.Init()
	customer.Init()
	structure.Init()
	myinterface.Init()
}

func index(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "index")
	http.ServeFile(w, r, "./static/index.htm")
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home")
}
