package entities

type User struct {
	Id      int64
	Balance float64
}

type InfoOperation struct {
	Id              int64
	CreatedAt       int64
	Discription     string
	SenderReceiver  string
	BalanceAtMoment string
	Amount          int64
	Refil           string
	UserId          int64
}
