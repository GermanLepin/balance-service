package fetch_balance_info_handler

import (
	"balance-service/internal/application/dto"

	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type (
	FetchBalanceInfoService interface {
		FetchBalanceInfo(ctx context.Context, userID uuid.UUID) (dto.User, error)
	}

	JsonService interface {
		ErrorJSON(w http.ResponseWriter, err error, statusCode int) error
	}
)

func (h *handler) FetchBalanceInfo(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var user dto.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.jsonService.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	user, err := h.fetchBalanceInfoService.FetchBalanceInfo(ctx, user.ID)
	if err != nil {
		h.jsonService.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&user)
	if err != nil {
		h.jsonService.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
}

func New(
	fetchBalanceInfoService FetchBalanceInfoService,
	jsonService JsonService,
) *handler {
	return &handler{
		fetchBalanceInfoService: fetchBalanceInfoService,
		jsonService:             jsonService,
	}
}

type handler struct {
	fetchBalanceInfoService FetchBalanceInfoService
	jsonService             JsonService
}
