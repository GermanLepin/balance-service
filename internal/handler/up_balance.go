package handler

import (
	"net/http"
	"tech_task/pkg/helpers/jsonenc"
	"tech_task/pkg/helpers/parseform"
	"tech_task/pkg/helpers/validate"
)

func (h *HttpService) UpBalance(w http.ResponseWriter, r *http.Request) {
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

	instance.UpBalanceDB(ctx, userId, correctAmount)
	jsonenc.JSONUpBalance(w, userId, correctAmount)
}
