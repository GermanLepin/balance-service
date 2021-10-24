package parseform

import (
	"net/http"
	"strings"
)

func Pars(w http.ResponseWriter, r *http.Request, value string) (correctVal string) {
	r.ParseForm()
	paramsRequest := r.Form
	valueSlice := paramsRequest[value]
	correctValue := strings.Join(valueSlice, " ")
	return correctValue
}
