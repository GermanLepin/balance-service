package entities

import (
	"context"
	"net/http"
	"time"
)

type User struct {
	Id      int64
	Balance float64
}

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

type UserService interface {
	BalanceInfoDB(ctx context.Context, w http.ResponseWriter, id int64) (userId int64, balance float64)
}
