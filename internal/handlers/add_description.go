package handlers

import (
	"net/http"
	"tech_task/pkg/helpers/jsonenc"
	"tech_task/pkg/helpers/parseform"
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

func AddDescription(w http.ResponseWriter, r *http.Request) {
	description := parseform.Pars(w, r, description)
	senderReceiver := parseform.Pars(w, r, senderReceiver)
	refill := parseform.Pars(w, r, refill)
	if refill == nilValue {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("Refill is not null field")
		jsonenc.JSONError(w, "Refill is not null field")
		return
	}

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
			log.Errorf("Insufficient funds")
			jsonenc.JSONError(w, "Insufficient funds")
			return
		}

		instance.WritingOffDB(ctx, id, corectAmount)
	}

	_, balanceAtMoment := instance.BalanceInfoDB(ctx, w, id)

	instance.AddDescriptionDB(ctx, w, id, balanceAtMoment, corectAmount, refill, description, senderReceiver)
	jsonenc.JSONUAddDescription(w, id, balanceAtMoment, corectAmount, refill, description, senderReceiver)
}
