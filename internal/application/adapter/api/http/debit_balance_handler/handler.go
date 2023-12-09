package replenish_balance_handler

import (
	"balance-service/internal/application/dto"

	"context"
	"encoding/json"
	"net/http"
)

type (
	DebitBalanceService interface {
		DebitBalance(ctx context.Context, user dto.DebitBalanceRequest) (dto.User, error)
	}

	JsonService interface {
		ErrorJSON(w http.ResponseWriter, err error, statusCode int) error
	}
)

func (h *handler) DebitBalance(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var debitBalanceRequest dto.DebitBalanceRequest
	if err := json.NewDecoder(r.Body).Decode(&debitBalanceRequest); err != nil {
		h.jsonService.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	user, err := h.debitBalanceService.DebitBalance(ctx, debitBalanceRequest)
	if err != nil {
		h.jsonService.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	debitBalanceResponse := dto.DebitBalanceResponse{
		UserID:  user.ID,
		Name:    user.Name,
		Balance: user.Balance,
		Message: "successful withdrawal of funds",
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&debitBalanceResponse)
	if err != nil {
		h.jsonService.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
}

func New(
	debitBalanceService DebitBalanceService,
	jsonService JsonService,
) *handler {
	return &handler{
		debitBalanceService: debitBalanceService,
		jsonService:         jsonService,
	}
}

type handler struct {
	debitBalanceService DebitBalanceService
	jsonService         JsonService
}
