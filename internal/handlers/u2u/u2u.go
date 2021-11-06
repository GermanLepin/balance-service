package u2u

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
	id1      = "id1"
	id2      = "id2"
	amount   = "amount"
	ctx      = context.Background()
	instance = pg.StartDB()
)

func U2U(w http.ResponseWriter, r *http.Request) {
	mapUser := parseform.ParsJSON(r)
	userIdString1 := string(mapUser[id1])
	userIdString2 := string(mapUser[id2])
	amountString := string(mapUser[amount])

	userId1 := validate.IdValidate(w, userIdString1)
	if userId1 < 1 {
		return
	}

	userId2 := validate.IdValidate(w, userIdString2)
	if userId2 < 1 {
		return
	}

	correctAmount := validate.AmountValidate(w, amountString)
	if correctAmount < 0.01 {
		return
	}

	Id1, balance := instance.BalanceInfoDB(ctx, w, userId1)
	if Id1 == 0 {
		return
	}

	if correctAmount > balance {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("insufficient funds")
		jsonenc.JSONError(w, "insufficient funds")
		return
	}

	instance.WritingOffDB(ctx, userId1, correctAmount)
	instance.UpBalanceDB(ctx, userId2, correctAmount)
	jsonenc.JSONU2U(w, userId1, userId2, correctAmount)
}
