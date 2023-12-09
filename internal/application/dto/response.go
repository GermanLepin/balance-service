package dto

import "github.com/google/uuid"

type CretaeUserResponse struct {
	UserID  uuid.UUID `json:"user_id"`
	Name    string    `json:"name"`
	Balance float64   `json:"balance"`
	Message string    `json:"message"`
}

type DeleteUserResponse struct {
	UserID  uuid.UUID `json:"user_id"`
	Name    string    `json:"name"`
	Balance float64   `json:"balance"`
	Message string    `json:"message"`
}

type BalanceInfoResponse struct {
	UserID  uuid.UUID `json:"user_id"`
	Name    string    `json:"name"`
	Balance float64   `json:"balance"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
