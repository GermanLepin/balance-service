package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type (
	JsonService interface {
		ParsJSON(r *http.Request) (map[string]string, error)

		WriteJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error
		ReadJSON(w http.ResponseWriter, r *http.Request, data any) error
	}
)

func (s *service) AddDescription(w http.ResponseWriter, r *http.Request) {
	mapUser, err := s.jsonService.ParsJSON(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userIDString := mapUser[id]
	userID, err := IdValidate(userIDString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		JSONError(w, err.Error())
		return
	}

	amountString := mapUser[amount]
	correctAmount, err := AmountValidate(amountString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		JSONError(w, err.Error())
		return
	}

	description := mapUser[description]
	if description == nilValue {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Errorf("Description is not null field")
		JSONError(w, "Description is not null field")
		return
	}

	senderReceiver := mapUser[senderReceiver]
	if senderReceiver == nilValue {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Errorf("Sender receiver is not null field")
		JSONError(w, "Sender receiver is not null field")
		return
	}

	refill := mapUser[refill]
	if refill == nilValue {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Errorf("Refill is not null field")
		JSONError(w, "Refill is not null field")
		return
	}

	switch refill {
	case TRUE:
		err := h.services.UpBalance.UpBalanceUser(ctx, userID, correctAmount)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			JSONError(w, "User not found")
			return
		}
	case FALSE:
		userID, balance, err := h.services.BalanceInfo.BalanceInfoUser(ctx, userID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			JSONError(w, "User not found")
			return
		}

		if correctAmount > balance {
			w.WriteHeader(http.StatusBadRequest)
			logrus.Errorf("Insufficient funds")
			JSONError(w, "Insufficient funds")
			return
		}

		h.services.WritingOff.WritingOffUser(ctx, userID, correctAmount)
	}

	userID, balanceAtMoment, err := h.services.BalanceInfo.BalanceInfoUser(ctx, userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		JSONError(w, "User not found")
		return
	}

	err = h.services.AddDescription.AddDescriptionUser(ctx, userID, balanceAtMoment, correctAmount, refill, description, senderReceiver)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		JSONError(w, "User not found")
		return
	}

	err = JSONUAddDescription(w, userID, balanceAtMoment, correctAmount, refill, description, senderReceiver)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func New(
	jsonService JsonService,
) *service {
	return &service{
		jsonService: jsonService,
	}
}

type service struct {
	jsonService JsonService
}
