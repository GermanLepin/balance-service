package dto

import "github.com/google/uuid"

type ReplenishBalanceRequest struct {
	UserID uuid.UUID `json:"user_id"`
	Amount float32   `json:"amount"`
}

type DepleteBalanceRequest struct {
	UserID uuid.UUID `json:"user_id"`
	Amount float32   `json:"amount"`
}
