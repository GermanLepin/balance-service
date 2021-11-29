package handler

import (
	"math"
	"net/http"
	"tech_task/pkg/helpers/convert"
	"tech_task/pkg/helpers/jsonenc"
	"tech_task/pkg/helpers/parseform"
	"tech_task/pkg/helpers/validate"

	log "github.com/sirupsen/logrus"
)

var (
	currency = "currency"
	RUB      = "RUB"
	USD      = "USD"
	static   = 100.00
)

func (h *HttpService) BalanceInfo(w http.ResponseWriter, r *http.Request) {
	mapUser := parseform.ParsJSON(r)
	userIdString := string(mapUser[id])
	currency := parseform.Pars(r, currency)

	userId := validate.IdValidate(w, userIdString)
	if userId < 1 {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("Incorrect value id user")
		jsonenc.JSONError(w, "Incorrect value id user")
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
