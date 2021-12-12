package tech_task

import "time"

type Description struct {
	Id              int64     `json:"id"`
	SenderReceiver  string    `json:"sender receiver"`
	Amount          float64   `json:"amount"`
	Description     string    `json:"description"`
	BalanceAtMoment float64   `json:"balance at moment"`
	UserId          int64     `json:"user id"`
	CreatedAt       time.Time `json:"created at"`
	Refill          string    `json:"refill"`
}
