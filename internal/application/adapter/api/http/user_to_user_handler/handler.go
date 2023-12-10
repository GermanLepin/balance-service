package create_user_handler

import (
	"balance-service/internal/application/dto"

	"context"
	"net/http"
)

type (
	UserToUserService interface {
		UserToUser(ctx context.Context, user1 dto.User, user2 dto.User) error
	}

	JsonService interface {
		ErrorJSON(w http.ResponseWriter, err error, statusCode int) error
	}
)

func (h *handler) UserToUser(w http.ResponseWriter, r *http.Request) {

}

func New(
	userToUserService UserToUserService,
	jsonService JsonService,
) *handler {
	return &handler{
		userToUserService: userToUserService,
		jsonService:       jsonService,
	}
}

type handler struct {
	userToUserService UserToUserService
	jsonService       JsonService
}
