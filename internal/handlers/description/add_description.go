package description

import (
	"context"
	"net/http"
	"tech_task/pkg/helpers/jsonenc"
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
	RUB            = "RUB"
	USD            = "USD"
	nilValue       = ""
	sortByDesc     = "sort_by_desc"
	data           = "created_at"
	id             = "id"
	amount         = "amount"
	sort_by        = "sort_by"
	ctx            = context.Background()
	instance       = pg.StartDB()
)

func AddDescription(w http.ResponseWriter, r *http.Request) {
	mapUser := parseform.ParsJSON(w, r)
	refill := string(mapUser[refill])
	if refill == nilValue {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("Refill is not null field")
		jsonenc.JSONError(w, "Refill is not null field")
		return
	}

	description := string(mapUser[description])
	senderReceiver := string(mapUser[senderReceiver])

	userIdString := string(mapUser[id])
	userId := validate.IdValidate(w, r, userIdString)
	if userId < 1 {
		return
	}

	amountString := string(mapUser[amount])
	corectAmount := validate.AmountValidate(w, r, amountString)
	if corectAmount < 0.01 {
		return
	}

	switch refill {
	case TRUE:
		instance.UpBalanceDB(ctx, w, userId, corectAmount)
	case FALSE:
		userId, balance := instance.BalanceInfoDB(ctx, w, userId)
		if userId == 0 {
			return
		}

		if corectAmount > balance {
			w.WriteHeader(http.StatusBadRequest)
			log.Errorf("Insufficient funds")
			jsonenc.JSONError(w, "Insufficient funds")
			return
		}

		instance.WritingOffDB(ctx, userId, corectAmount)
	}

	_, balanceAtMoment := instance.BalanceInfoDB(ctx, w, userId)

	instance.AddDescriptionDB(ctx, w, userId, balanceAtMoment, corectAmount, refill, description, senderReceiver)
	jsonenc.JSONUAddDescription(w, userId, balanceAtMoment, corectAmount, refill, description, senderReceiver)
}
