package handler

import (
	"net/http"
	json "tech_task/pkg/helpers/json_responce"
	"tech_task/pkg/helpers/parse"
	"tech_task/pkg/helpers/validate"

	"github.com/sirupsen/logrus"
)

func (h *Handler) WritingOff(w http.ResponseWriter, r *http.Request) {
	mapUser, err := parse.ParsJSON(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userIdString := string(mapUser[id])
	amountString := string(mapUser[amount])

	userId, err := validate.IdValidate(userIdString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.JSONError(w, err.Error())
		return
	}

	correctAmount, err := validate.AmountValidate(amountString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.JSONError(w, err.Error())
		return
	}

	userId, balance, err := h.services.BalanceInfo.BalanceInfoUser(ctx, userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.JSONError(w, "User not found")
		return
	}

	if correctAmount > balance {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Errorf("insufficient funds")
		json.JSONError(w, "insufficient funds")
		return
	}

	h.services.WritingOff.WritingOffUser(ctx, userId, correctAmount)
	err = json.JSONWritingOff(w, userId, correctAmount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
