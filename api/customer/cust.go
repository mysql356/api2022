package customer

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"
)

func Init() {

	http.HandleFunc("/api/cust", cust)
	http.HandleFunc("/api/cust/login", login)
}

func cust(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "cust...")
}

func login(w http.ResponseWriter, r *http.Request) {

	log.Print(r.URL)                         // /api/cust/login
	log.Print(html.EscapeString(r.URL.Path)) // /api/cust/login
	log.Print(r.Method)                      // GET, POST
	log.Print(r.Header.Get("User-Agent"))    // PostmanRuntime/7.26.8

	log.Print(r.Body) // &{0xc00018c0e0 <nil> <nil> false true {0 0} false false false 0x67e200}

	var resp response
	err := json.NewDecoder(r.Body).Decode(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp.Password = "xxx"
	resp.Username = strings.Replace(resp.Email, "@", "AT", -1)
	resp.Username = strings.Replace(resp.Username, ".", "DOT", -1)

	log.Print(resp.Email)
	log.Print(resp.Password)

	//output
	//fmt.Fprintf(w, "%+v\n", resp) // {Email:customer@gmail.co.im Password:xxx Username:customerATgmailDOTcoDOTim}
	//fmt.Fprintf(w, "%v\n", resp)  // {"email":"customer@gmail.co.in","password":"xxx","username":"customerATgmailDOTcoDOTin"}

	resp_byte, _ := json.Marshal(resp)
	fmt.Fprint(w, string(resp_byte)) // {"email":"customer@gmail.co.in","password":"xxx","username":"customerATgmailDOTcoDOTin"}

}
