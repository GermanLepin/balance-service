package handlers

import (
	"net/http"
	"tech_task/pkg/helpers/jsonenc"
	"tech_task/pkg/helpers/parseform"
	"tech_task/pkg/helpers/validate"
)

func UpBalance(w http.ResponseWriter, r *http.Request) {
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

	instance.UpBalanceDB(ctx, w, userId, corectAmount)
	jsonenc.JSONUpBalance(w, userId, corectAmount)
}
