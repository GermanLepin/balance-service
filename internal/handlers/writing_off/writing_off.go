package writingOff

import (
	"context"
	"net/http"
	"tech_task/pkg/helpers/jsonenc"
	"tech_task/pkg/helpers/parseform"
	"tech_task/pkg/helpers/pg"
	"tech_task/pkg/helpers/validate"

	log "github.com/sirupsen/logrus"
)

var (
	id       = "id"
	amount   = "amount"
	ctx      = context.Background()
	instance = pg.StartDB()
)

func WritingOff(w http.ResponseWriter, r *http.Request) {
	mapUser := parseform.ParsJSON(r)
	userIdString := string(mapUser[id])
	amountString := string(mapUser[amount])

	userId := validate.IdValidate(w, userIdString)
	if userId < 1 {
		return
	}

	correctAmount := validate.AmountValidate(w, amountString)
	if correctAmount < 0.01 {
		return
	}

	userIdBalance, balance := instance.BalanceInfoDB(ctx, w, userId)
	if userIdBalance == 0 {
		return
	}

	if correctAmount > balance {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("insufficient funds")
		jsonenc.JSONError(w, "insufficient funds")
		return
	}

	instance.WritingOffDB(ctx, userId, correctAmount)
	jsonenc.JSONWritingOff(w, userId, correctAmount)
}
