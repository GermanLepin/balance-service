package dto

import "github.com/google/uuid"

type ReplenishBalanceRequest struct {
	UserID uuid.UUID `json:"user_id"`
	Amount float64   `json:"amount"`
}

type DebitBalanceRequest struct {
	UserID uuid.UUID `json:"user_id"`
	Amount float64   `json:"amount"`
}
