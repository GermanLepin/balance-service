package handler

import (
	"net/http"

	json "tech_task/pkg/helpers/json_responce"
	"tech_task/pkg/helpers/parse"
	"tech_task/pkg/helpers/validate"

	"github.com/sirupsen/logrus"
)

func (h *Handler) U2U(w http.ResponseWriter, r *http.Request) {
	mapUser, err := parse.ParsJSON(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userIdString1 := string(mapUser[id1])
	userIdString2 := string(mapUser[id2])
	amountString := string(mapUser[amount])

	userId1, err := validate.IdValidate(userIdString1)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.JSONError(w, err.Error())
		return
	}

	userId2, err := validate.IdValidate(userIdString2)
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

	_, balance, err := h.services.BalanceInfo.BalanceInfoUser(ctx, userId1)
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

	h.services.WritingOff.WritingOffUser(ctx, userId1, correctAmount)
	h.services.UpBalance.UpBalanceUser(ctx, userId2, correctAmount)
	err = json.JSONU2U(w, userId1, userId2, correctAmount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
