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
	FALSE    = "F"
	TRUE     = "T"
	RUB      = "RUB"
	USD      = "USD"
	id       = "id"
	amount   = "amount"
	ctx      = context.Background()
	instance = pg.StartDB()
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
