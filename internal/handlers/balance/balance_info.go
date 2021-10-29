package balance

import (
	"context"
	"math"
	"net/http"
	"tech_task/pkg/helpers/convert"
	"tech_task/pkg/helpers/jsonenc"
	"tech_task/pkg/helpers/parseform"
	"tech_task/pkg/helpers/pg"
	"tech_task/pkg/helpers/validate"

	log "github.com/sirupsen/logrus"
)

var (
	FALSE    = "F"
	TRUE     = "T"
	currency = "currency"
	RUB      = "RUB"
	USD      = "USD"
	static   = 100.00
	nilValue = ""
	id       = "id"
	ctx      = context.Background()
	instance = pg.StartDB()
)

func BalanceInfo(w http.ResponseWriter, r *http.Request) {
	mapUser := parseform.ParsJSON(r)
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
