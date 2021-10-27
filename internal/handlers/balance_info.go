package handlers

import (
	"math"
	"net/http"
	"tech_task/pkg/helpers/convert"
	"tech_task/pkg/helpers/jsonenc"
	"tech_task/pkg/helpers/parseform"
	"tech_task/pkg/helpers/validate"

	log "github.com/sirupsen/logrus"
)

func BalanceInfo(w http.ResponseWriter, r *http.Request) {
	mapUser := parseform.ParsJSON(w, r)
	userIdString := string(mapUser[id])
	currency := parseform.Pars(w, r, currency)

	userId := validate.IdValidate(w, r, userIdString)
	if userId < 1 {
		return
	}

	Id, rubBalance := instance.BalanceInfoDB(ctx, w, userId)
	if Id == 0 {
		return
	}

	if currency == USD {
		rub := convert.GetConvertValue(w, RUB)
		usdToEur := convert.GetConvertValue(w, USD)
		usdAmount := convert.UsdAmount(usdToEur, rub, rubBalance)
		userBalanceUsd := math.Floor(usdAmount*static) / static
		jsonenc.JSONBalanceInfo(w, userId, userBalanceUsd)
	} else if currency == nilValue {
		jsonenc.JSONBalanceInfo(w, userId, rubBalance)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		jsonenc.JSONError(w, "Invalid currency type, only USD")
		log.Errorf("Invalid currency type, only USD")
	}
}
