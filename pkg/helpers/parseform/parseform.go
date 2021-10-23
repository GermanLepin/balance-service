package parseform

import (
	"net/http"
	"strings"
)

func Pars(w http.ResponseWriter, r *http.Request, value string) (correctVal string) {
	r.ParseForm()
	params_request := r.Form
	valie_slice := params_request[value]
	correctValue := strings.Join(valie_slice, " ")
	return correctValue
}
