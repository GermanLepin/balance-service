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
	ctx       = context.Background()
	currency  = "currency"
	RUB       = "RUB"
	USD       = "USD"
	static    = 100.00
	nil_value = ""
)

func BalanceInfo(w http.ResponseWriter, r *http.Request) {
	instance := pg.StartDB()
	currency := parseform.Pars(w, r, currency)

	id := validate.IdValidate(w, r, id)
	if id < 1 {
		return
	}

	user_id, rub_balance := instance.BalanceInfoDB(ctx, w, id)
	if user_id == 0 {
		return
	}

	if currency == USD {
		rub := convert.GetConvertValue(w, RUB)
		usd_to_eur := convert.GetConvertValue(w, USD)
		usd_amount := convert.UsdAmount(usd_to_eur, rub, rub_balance)
		user_balance_usd := math.Floor(usd_amount*static) / static
		jsonenc.JSONEncoder(w, id, user_balance_usd)
	} else if currency == RUB || currency == "" {
		jsonenc.JSONEncoder(w, id, rub_balance)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		jsonenc.JSONError(w, "Invalid currency type, only RUB or USD")
		log.Errorf("Invalid currency type, only RUB or USD")
	}
}
