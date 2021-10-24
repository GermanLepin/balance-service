package handlers

import (
	"net/http"
	"tech_task/pkg/helpers/jsonenc.go"
	"tech_task/pkg/helpers/pg"
	"tech_task/pkg/helpers/validate"
)

var (
	id     = "id"
	amount = "amount"
)

func UpBalance(w http.ResponseWriter, r *http.Request) {
	instance := pg.StartDB()

	id := validate.IdValidate(w, r, id)
	if id < 1 {
		return
	}

	corectAmount := validate.AmountValidate(w, r, amount)
	if corectAmount < 0.01 {
		return
	}

	instance.UpBalanceDB(ctx, w, id, corectAmount)
	jsonenc.JSONUpBalance(w, id, corectAmount)
}
