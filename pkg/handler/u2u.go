package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func (h *Handler) U2U(w http.ResponseWriter, r *http.Request) {
	mapUser, err := ParsJSON(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userIDString1 := string(mapUser[id1])
	userIDString2 := string(mapUser[id2])
	amountString := string(mapUser[amount])

	userID1, err := IdValidate(userIDString1)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		JSONError(w, err.Error())
		return
	}

	userID2, err := IdValidate(userIDString2)
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

	_, balance, err := h.services.BalanceInfo.BalanceInfoUser(ctx, userID1)
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

	h.services.WritingOff.WritingOffUser(ctx, userID1, correctAmount)
	h.services.UpBalance.UpBalanceUser(ctx, userID2, correctAmount)
	err = JSONU2U(w, userID1, userID2, correctAmount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
