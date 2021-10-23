package validate

import (
	"net/http"
	"strconv"
	"strings"
	"tech_task/pkg/helpers/jsonenc.go"
	"tech_task/pkg/helpers/parseform"

	log "github.com/sirupsen/logrus"
)

func AmountValidate(w http.ResponseWriter, r *http.Request, amount_s string) (correct_amount float64) {
	amount_string := parseform.Pars(w, r, amount_s)

	valid_amount := strings.Split(amount_string, ".")
	if len(valid_amount) > 1 {
		if len(valid_amount[1]) > 2 {
			w.WriteHeader(http.StatusBadRequest)
			log.Errorf("The amount have more then 2 decimal places")
			jsonenc.JSONError(w, "The amount have more then 2 decimal places")
			return
		}
	}

	amount, err := strconv.ParseFloat(amount_string, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(err).Errorf("Error with parcing amount")
		jsonenc.JSONError(w, "Error with parcing amount")
		return
	}

	if amount < 0.01 {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(err).Errorf("The amount is negative")
		jsonenc.JSONError(w, "The amount is negative")
		return
	}

	return amount
}
