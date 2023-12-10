package fetch_balance_info_handler

import (
	"balance-service/internal/application/dto"

	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
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

	userID := chi.URLParam(r, "uuid")
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		h.jsonService.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	user, err := h.fetchBalanceInfoService.FetchBalanceInfo(ctx, userUUID)
	if err != nil {
		h.jsonService.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	balanceInfoResponse := dto.BalanceInfoResponse{
		UserID:  user.ID,
		Name:    user.Name,
		Balance: user.Balance,
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&balanceInfoResponse)
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
