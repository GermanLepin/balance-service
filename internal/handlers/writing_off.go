package handlers

import (
	"net/http"
	"tech_task/pkg/helpers/jsonenc"
	"tech_task/pkg/helpers/parseform"
	"tech_task/pkg/helpers/validate"

	log "github.com/sirupsen/logrus"
)

func WritingOff(w http.ResponseWriter, r *http.Request) {
	mapUser := parseform.ParsJSON(w, r)
	userIdString := string(mapUser[id])
	amountString := string(mapUser[amount])

	userId := validate.IdValidate(w, r, userIdString)
	if userId < 1 {
		return
	}

	corectAmount := validate.AmountValidate(w, r, amountString)
	if corectAmount < 0.01 {
		return
	}

	userIdBalance, balance := instance.BalanceInfoDB(ctx, w, userId)
	if userIdBalance == 0 {
		return
	}

	if corectAmount > balance {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("insufficient funds")
		jsonenc.JSONError(w, "insufficient funds")
		return
	}

	instance.WritingOffDB(ctx, userId, corectAmount)
	jsonenc.JSONWritingOff(w, userId, corectAmount)
}
