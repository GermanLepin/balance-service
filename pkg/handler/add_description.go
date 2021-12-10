package handler

import (
	"context"
	"net/http"
	json "tech_task/pkg/helpers/json_responce"
	"tech_task/pkg/helpers/parse"
	"tech_task/pkg/helpers/validate"

	"github.com/sirupsen/logrus"
)

var (
	asc            = "asc"
	desc           = "desc"
	description    = "description"
	senderReceiver = "sender_receiver"
	refill         = "refill"
	FALSE          = "F"
	TRUE           = "T"
	nilValue       = ""
	data           = "created_at"
	amount         = "amount"
	sortBy         = "sort_by"
	orderBy        = "order_by"
	sqlOrderBy     = "ORDER BY"
	ctx            = context.Background()
	id             = "id"
	id1            = "id1"
	id2            = "id2"
	currency       = "currency"
	RUB            = "RUB"
	USD            = "USD"
	static         = 100.00
)

func (h *Handler) AddDescription(w http.ResponseWriter, r *http.Request) {
	mapUser, err := parse.ParsJSON(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	refill := string(mapUser[refill])
	if refill == nilValue {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Errorf("Refill is not null field")
		json.JSONError(w, "Refill is not null field")
		return
	}

	description := string(mapUser[description])
	senderReceiver := string(mapUser[senderReceiver])
	userIdString := string(mapUser[id])
	amountString := string(mapUser[amount])

	userId, err := validate.IdValidate(userIdString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.JSONError(w, err.Error())
		return
	}

	correctAmount, err := validate.AmountValidate(amountString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.JSONError(w, err.Error())
		return
	}

	switch refill {
	case TRUE:
		err := h.services.UpBalance.UpBalanceUser(ctx, userId, correctAmount)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.JSONError(w, "User not found")
			return
		}
	case FALSE:
		userId, balance, err := h.services.BalanceInfo.BalanceInfoUser(ctx, userId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.JSONError(w, "User not found")
			return
		}

		if correctAmount > balance {
			w.WriteHeader(http.StatusBadRequest)
			logrus.Errorf("Insufficient funds")
			json.JSONError(w, "Insufficient funds")
			return
		}

		h.services.WritingOff.WritingOffUser(ctx, userId, correctAmount)
	}

	userId, balanceAtMoment, err := h.services.BalanceInfo.BalanceInfoUser(ctx, userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.JSONError(w, "User not found")
		return
	}

	err = h.services.AddDescription.AddDescriptionUser(ctx, userId, balanceAtMoment, correctAmount, refill, description, senderReceiver)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.JSONError(w, "User not found")
		return
	}

	err = json.JSONUAddDescription(w, userId, balanceAtMoment, correctAmount, refill, description, senderReceiver)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
