package delete_user_handler

import (
	"balance-service/internal/application/dto"

	"context"
	"encoding/json"
	"net/http"

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

	var user dto.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.jsonService.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	user, err := h.deleteUserService.DeleteUser(ctx, user.ID)
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
