package dto

import "time"

type Description struct {
	Id              int64     `json:"id"`
	SenderReceiver  string    `json:"sender_receiver"`
	Amount          float64   `json:"amount"`
	Description     string    `json:"description"`
	BalanceAtMoment float64   `json:"balance_at_moment"`
	UserID          int64     `json:"user_id"`
	CreatedAt       time.Time `json:"created_at"`
	Refill          string    `json:"refill"`
}
