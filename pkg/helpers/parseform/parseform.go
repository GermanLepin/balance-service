package parseform

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	mapUser map[string]string
)

func Pars(w http.ResponseWriter, r *http.Request, value string) (correctVal string) {
	r.ParseForm()
	paramsRequest := r.Form
	valueSlice := paramsRequest[value]
	correctValue := strings.Join(valueSlice, " ")
	return correctValue
}

func ParsJSON(w http.ResponseWriter, r *http.Request) map[string]string {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error parcing request")
	}

	if err := json.Unmarshal(body, &mapUser); err != nil {
		log.Printf("Error parcing JSON")
	}

	return mapUser
}
