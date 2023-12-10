package delete_user_handler

import (
	"balance-service/internal/application/dto"

	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type (
	DeleteUserService interface {
		DeleteUser(ctx context.Context, userID uuid.UUID) (dto.User, error)
	}

	JsonService interface {
		ErrorJSON(w http.ResponseWriter, err error, statusCode int) error
	}
)

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	userID := chi.URLParam(r, "uuid")
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		h.jsonService.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	user, err := h.deleteUserService.DeleteUser(ctx, userUUID)
	if err != nil {
		h.jsonService.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	deleteUserResponse := dto.DeleteUserResponse{
		UserID:  user.ID,
		Name:    user.Name,
		Balance: user.Balance,
		Message: "successful user deletion",
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&deleteUserResponse)
	if err != nil {
		h.jsonService.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
}

func New(
	deleteUserService DeleteUserService,
	jsonService JsonService,
) *handler {
	return &handler{
		deleteUserService: deleteUserService,
		jsonService:       jsonService,
	}
}

type handler struct {
	deleteUserService DeleteUserService
	jsonService       JsonService
}
