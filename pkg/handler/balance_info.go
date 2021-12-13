package handler

import (
	"math"
	"net/http"
	"tech_task/pkg/convert"
	json "tech_task/pkg/helpers/json_responce"
	"tech_task/pkg/helpers/parse"
	"tech_task/pkg/helpers/validate"

	"github.com/sirupsen/logrus"
)

func (h *Handler) BalanceInfo(w http.ResponseWriter, r *http.Request) {
	mapUser, err := parse.ParsJSON(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userIDString := string(mapUser[id])
	currency := parse.Pars(r, currency)

	userID, err := validate.IdValidate(userIDString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.JSONError(w, err.Error())
		return
	}

	_, rubBalance, err := h.services.BalanceInfo.BalanceInfoUser(ctx, userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.JSONError(w, "User not found")
		return
	}

	if currency == USD {
		rub := convert.GetConvertValue(w, RUB)
		usdToEur := convert.GetConvertValue(w, USD)
		usdAmount := convert.UsdAmount(usdToEur, rub, rubBalance)
		userBalanceUsd := math.Floor(usdAmount*static) / static
		err = json.JSONBalanceInfo(w, userID, userBalanceUsd)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else if currency == nilValue {
		err = json.JSONBalanceInfo(w, userID, rubBalance)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.JSONError(w, "Invalid currency type, only USD")
		logrus.Errorf("invalid currency type, only USD")
		return
	}
}
