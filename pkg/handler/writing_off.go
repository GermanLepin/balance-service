package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func (h *Handler) WritingOff(w http.ResponseWriter, r *http.Request) {
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

	userID, balance, err := h.services.BalanceInfo.BalanceInfoUser(ctx, userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		JSONError(w, "User not found")
		return
	}

	if correctAmount > balance {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Errorf("insufficient funds")
		JSONError(w, "insufficient funds")
		return
	}

	h.services.WritingOff.WritingOffUser(ctx, userID, correctAmount)
	err = JSONWritingOff(w, userID, correctAmount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
