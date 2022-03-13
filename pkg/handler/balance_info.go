package handler

import (
	"math"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (h *Handler) BalanceInfo(w http.ResponseWriter, r *http.Request) {
	mapUser, err := ParsJSON(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userIDString := mapUser[id]
	currency := Pars(r, currency)

	userID, err := IdValidate(userIDString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		JSONError(w, err.Error())
		return
	}

	_, rubBalance, err := h.services.BalanceInfo.BalanceInfoUser(ctx, userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		JSONError(w, "User not found")
		return
	}

	if currency == USD {
		rub := GetConvertValue(w, RUB)
		usdToEur := GetConvertValue(w, USD)
		usdAmount := UsdAmount(usdToEur, rub, rubBalance)
		userBalanceUsd := math.Floor(usdAmount*static) / static
		err = JSONBalanceInfo(w, userID, userBalanceUsd)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else if currency == nilValue {
		err = JSONBalanceInfo(w, userID, rubBalance)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		JSONError(w, "Invalid currency type, only USD")
		logrus.Errorf("invalid currency type, only USD")
		return
	}
}
