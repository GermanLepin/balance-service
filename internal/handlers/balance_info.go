package handlers

import (
	"context"
	"math"
	"net/http"
	"tech_task/pkg/helpers/convert"
	"tech_task/pkg/helpers/jsonenc.go"
	"tech_task/pkg/helpers/parseform"
	"tech_task/pkg/helpers/pg"
	"tech_task/pkg/helpers/validate"

	log "github.com/sirupsen/logrus"
)

var (
	ctx      = context.Background()
	currency = "currency"
	RUB      = "RUB"
	USD      = "USD"
	static   = 100.00
	nilValue = ""
)

func BalanceInfo(w http.ResponseWriter, r *http.Request) {
	instance := pg.StartDB()
	currency := parseform.Pars(w, r, currency)

	id := validate.IdValidate(w, r, id)
	if id < 1 {
		return
	}

	userId, rubBalance := instance.BalanceInfoDB(ctx, w, id)
	if userId == 0 {
		return
	}

	if currency == USD {
		rub := convert.GetConvertValue(w, RUB)
		usdToEur := convert.GetConvertValue(w, USD)
		usdAmount := convert.UsdAmount(usdToEur, rub, rubBalance)
		userBalanceUsd := math.Floor(usdAmount*static) / static
		jsonenc.JSONEncoder(w, id, userBalanceUsd)
	} else if currency == RUB || currency == "" {
		jsonenc.JSONEncoder(w, id, rubBalance)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		jsonenc.JSONError(w, "Invalid currency type, only RUB or USD")
		log.Errorf("Invalid currency type, only RUB or USD")
	}
}
