package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func (h *Handler) UpBalance(w http.ResponseWriter, r *http.Request) {
	mapUser, err := ParsJSON(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userIDString := mapUser[id]
	amountString := mapUser[amount]

	userID, err := IdValidate(userIDString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		JSONError(w, err.Error())
		return
	}

	correctAmount, err := AmountValidate(amountString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		JSONError(w, err.Error())
		return
	}

	err = h.services.UpBalance.UpBalanceUser(ctx, userID, correctAmount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.WithError(err).Errorf("user not found")
		return
	}

	err = JSONUpBalance(w, userID, correctAmount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
