package handlers

import (
	"net/http"
	"tech_task/pkg/helpers/jsonenc"
	"tech_task/pkg/helpers/validate"

	log "github.com/sirupsen/logrus"
)

var (
	id1 = "id1"
	id2 = "id2"
)

func U2U(w http.ResponseWriter, r *http.Request) {
	id1 := validate.IdValidate(w, r, id1)
	if id1 < 1 {
		return
	}

	id2 := validate.IdValidate(w, r, id2)
	if id2 < 1 {
		return
	}

	amount := validate.AmountValidate(w, r, amount)
	if amount < 0.01 {
		return
	}

	userId, balance := instance.BalanceInfoDB(ctx, w, id1)
	if userId == 0 {
		return
	}

	if amount > balance {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("insufficient funds")
		jsonenc.JSONError(w, "insufficient funds")
		return
	}

	instance.WritingOffDB(ctx, id1, amount)
	instance.UpBalanceDB(ctx, w, id2, amount)
	jsonenc.JSONU2U(w, id1, id2, amount)
}
