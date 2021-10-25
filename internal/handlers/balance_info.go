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
	id := validate.IdValidate(w, r, id)
	if id < 1 {
		return
	}

	userId, rubBalance := instance.BalanceInfoDB(ctx, w, id)
	if userId == 0 {
		return
	}

	jsonenc.JSONBalanceInfo(w, id, rubBalance)
}

func BalanceInfoConvert(w http.ResponseWriter, r *http.Request) {
	id := validate.IdValidate(w, r, id)
	if id < 1 {
		return
	}

	userId, rubBalance := instance.BalanceInfoDB(ctx, w, id)
	if userId == 0 {
		return
	}

	currency := parseform.Pars(w, r, currency)

	if currency == USD {
		rub := convert.GetConvertValue(w, RUB)
		usdToEur := convert.GetConvertValue(w, USD)
		usdAmount := convert.UsdAmount(usdToEur, rub, rubBalance)
		userBalanceUsd := math.Floor(usdAmount*static) / static
		jsonenc.JSONBalanceInfo(w, id, userBalanceUsd)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		jsonenc.JSONError(w, "Invalid currency type, only USD")
		log.Errorf("Invalid currency type, only USD")
	}
}
