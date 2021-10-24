package parseform

import (
	"net/http"
	"strings"
	"tech_task/pkg/helpers/jsonenc"

	log "github.com/sirupsen/logrus"
)

func Pars(w http.ResponseWriter, r *http.Request, value string) (correctVal string) {
	r.ParseForm()
	paramsRequest := r.Form
	valueSlice := paramsRequest[value]
	correctValue := strings.Join(valueSlice, " ")
	if len(correctValue) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("Errors parameters were not passed")
		jsonenc.JSONError(w, "Errors parameters were not passed")
		return
	}
	return correctValue
}
