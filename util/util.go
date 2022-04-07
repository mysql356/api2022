package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Render(w http.ResponseWriter, p interface{}) {

	resp_byte, _ := json.Marshal(p)
	fmt.Fprint(w, string(resp_byte))
}
