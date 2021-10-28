package handlers

import (
	"net/http"
	"tech_task/pkg/helpers/jsonenc"
	"tech_task/pkg/helpers/parseform"
	"tech_task/pkg/helpers/validate"

	log "github.com/sirupsen/logrus"
)

func U2U(w http.ResponseWriter, r *http.Request) {
	mapUser := parseform.ParsJSON(w, r)
	userIdString1 := string(mapUser[id1])
	userIdString2 := string(mapUser[id2])
	amountString := string(mapUser[amount])

	userId1 := validate.IdValidate(w, r, userIdString1)
	if userId1 < 1 {
		return
	}

	userId2 := validate.IdValidate(w, r, userIdString2)
	if userId2 < 1 {
		return
	}

	corectAmount := validate.AmountValidate(w, r, amountString)
	if corectAmount < 0.01 {
		return
	}

	Id1, balance := instance.BalanceInfoDB(ctx, w, userId1)
	if Id1 == 0 {
		return
	}

	if corectAmount > balance {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("insufficient funds")
		jsonenc.JSONError(w, "insufficient funds")
		return
	}

	instance.WritingOffDB(ctx, userId1, corectAmount)
	instance.UpBalanceDB(ctx, w, userId2, corectAmount)
	jsonenc.JSONU2U(w, userId1, userId2, corectAmount)
}
