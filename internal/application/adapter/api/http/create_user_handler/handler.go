package create_user_handler

import (
	"balance-service/internal/application/dto"
	"errors"

	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type CretaeUserService interface {
	CreateUser(ctx context.Context, user dto.User) error
}

func (h *handler) CretaeUser(w http.ResponseWriter, r *http.Request) {
	var user dto.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error: unmarshal body %s\n", err)
	}
	user.ID = uuid.New()

	if err := validateBalance(user.Balance); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error: %s\n", err)
	}

	ctx := context.Background()
	if err := h.cretaeUserService.CreateUser(ctx, user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error: cannot save a user to db %s\n", err)
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error: cannot sand a request %s\n", err)
	}
}

func validateBalance(balance float64) error {
	if balance < 0 {
		log.Printf("balance cannot be negative")
		return errors.New("balance cannot be negative")
	}

	return nil
}

func New(cretaeUserService CretaeUserService) *handler {
	return &handler{
		cretaeUserService: cretaeUserService,
	}
}

type handler struct {
	cretaeUserService CretaeUserService
}
