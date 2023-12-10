package deplete_balance_handler

import (
	"balance-service/internal/application/dto"
	"context"
	"encoding/json"
	"net/http"
)

type (
	DepleteBalanceService interface {
		DepleteBalance(ctx context.Context, user dto.DepleteBalanceRequest) (dto.User, error)
	}

	JsonService interface {
		ErrorJSON(w http.ResponseWriter, err error, statusCode int) error
	}
)

func (h *handler) DepleteBalance(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var depleteBalanceRequest dto.DepleteBalanceRequest
	if err := json.NewDecoder(r.Body).Decode(&depleteBalanceRequest); err != nil {
		h.jsonService.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	user, err := h.depleteBalanceService.DepleteBalance(ctx, depleteBalanceRequest)
	if err != nil {
		h.jsonService.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	depleteBalanceResponse := dto.DepleteBalanceResponse{
		UserID:  user.ID,
		Name:    user.Name,
		Balance: user.Balance,
		Message: "successful withdrawal of funds",
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&depleteBalanceResponse)
	if err != nil {
		h.jsonService.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
}

func New(
	depleteBalanceService DepleteBalanceService,
	jsonService JsonService,
) *handler {
	return &handler{
		depleteBalanceService: depleteBalanceService,
		jsonService:           jsonService,
	}
}

type handler struct {
	depleteBalanceService DepleteBalanceService
	jsonService           JsonService
}
