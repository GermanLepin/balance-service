package tech_task

import "time"

type Description struct {
	Id              int64
	SenderReceiver  string
	Amount          float64
	Description     string
	BalanceAtMoment float64
	UserId          int64
	CreatedAt       time.Time
	Refil           string
}
