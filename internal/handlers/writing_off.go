package handlers

import (
	"net/http"
	"tech_task/pkg/helpers/jsonenc.go"
	"tech_task/pkg/helpers/pg"
	"tech_task/pkg/helpers/validate"

	log "github.com/sirupsen/logrus"
)

func WritingOff(w http.ResponseWriter, r *http.Request) {
	instance := pg.StartDB()

	id := validate.IdValidate(w, r, id)
	if id < 1 {
		return
	}

	amount := validate.AmountValidate(w, r, amount)
	if amount < 0.01 {
		return
	}

	user_id, balance := instance.BalanceInfoDB(ctx, w, id)
	if user_id == 0 {
		return
	}

	if amount > balance {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("insufficient funds")
		jsonenc.JSONError(w, "insufficient funds")
		return
	}

	instance.WritingOffDB(ctx, id, amount)
	jsonenc.JSONWritingOff(w, id, amount)
}
