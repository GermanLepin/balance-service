package handlers

import (
	"net/http"
	"tech_task/pkg/helpers/jsonenc.go"
	"tech_task/pkg/helpers/parseform"
	"tech_task/pkg/helpers/pg"
	"tech_task/pkg/helpers/validate"

	log "github.com/sirupsen/logrus"
)

var (
	description    = "description"
	senderReceiver = "sender_receiver"
	refill         = "refill"
	FALSE          = "F"
	TRUE           = "T"
)

func AddDiscription(w http.ResponseWriter, r *http.Request) {
	instance := pg.StartDB()
	description := parseform.Pars(w, r, description)
	senderReceiver := parseform.Pars(w, r, senderReceiver)
	refill := parseform.Pars(w, r, refill)

	id := validate.IdValidate(w, r, id)
	if id < 1 {
		return
	}

	corectAmount := validate.AmountValidate(w, r, amount)
	if corectAmount < 0.01 {
		return
	}

	switch refill {
	case TRUE:
		instance.UpBalanceDB(ctx, w, id, corectAmount)
	case FALSE:
		userId, balance := instance.BalanceInfoDB(ctx, w, id)
		if userId == 0 {
			return
		}

		if corectAmount > balance {
			w.WriteHeader(http.StatusBadRequest)
			log.Errorf("insufficient funds")
			jsonenc.JSONError(w, "insufficient funds")
			return
		}

		instance.WritingOffDB(ctx, id, corectAmount)
	}

	_, balanceAtMoment := instance.BalanceInfoDB(ctx, w, id)

	instance.AddDiscriptionDB(ctx, w, id, balanceAtMoment, corectAmount, refill, description, senderReceiver)
	jsonenc.JSONUAddDiscription(w, id, balanceAtMoment, corectAmount, refill, description, senderReceiver)
}
