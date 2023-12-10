package replenish_balance_handler

import (
	"balance-service/internal/application/dto"

	"context"
	"encoding/json"
	"net/http"
)

type (
	ReplenishBalanceService interface {
		ReplenishBalance(ctx context.Context, user dto.ReplenishBalanceRequest) (dto.User, error)
	}

	JsonService interface {
		ErrorJSON(w http.ResponseWriter, err error, statusCode int) error
	}
)

func (h *handler) ReplenishBalance(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var replenishBalanceRequest dto.ReplenishBalanceRequest
	if err := json.NewDecoder(r.Body).Decode(&replenishBalanceRequest); err != nil {
		h.jsonService.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	user, err := h.replenishBalanceService.ReplenishBalance(ctx, replenishBalanceRequest)
	if err != nil {
		h.jsonService.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	replenishBalanceResponse := dto.ReplenishBalanceResponse{
		UserID:  user.ID,
		Name:    user.Name,
		Balance: user.Balance,
		Message: "successful balance replenishment",
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&replenishBalanceResponse)
	if err != nil {
		h.jsonService.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
}

func New(
	replenishBalanceService ReplenishBalanceService,
	jsonService JsonService,
) *handler {
	return &handler{
		replenishBalanceService: replenishBalanceService,
		jsonService:             jsonService,
	}
}

type handler struct {
	replenishBalanceService ReplenishBalanceService
	jsonService             JsonService
}
