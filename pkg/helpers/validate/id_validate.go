package validate

import (
	"net/http"
	"strconv"
	"tech_task/pkg/helpers/jsonenc.go"
	"tech_task/pkg/helpers/parseform"

	log "github.com/sirupsen/logrus"
)

func IdValidate(w http.ResponseWriter, r *http.Request, id_account string) (id int64) {
	user_id := parseform.Pars(w, r, id_account)

	id, err := strconv.ParseInt(user_id, 0, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.WithError(err).Errorf("Error with parcing id")
		jsonenc.JSONError(w, "Error with parcing id")
		return
	}

	if id < 1 {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("Incorrect value id user")
		jsonenc.JSONError(w, "Incorrect value id user")
		return
	}

	return
}
