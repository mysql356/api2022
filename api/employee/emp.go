package employee

import (
	"fmt"
	"net/http"
)

func Init() {
	http.HandleFunc("/api/emp", emp)
	http.HandleFunc("/api/emp/login", login)

}

func emp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "emp...")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "emp-login...")
}
