package repository

// import (
// 	"context"
// 	"database/sql"
// 	"tech_task"
// )

// type UpBalance interface {
// 	UpBalanceDB(ctx context.Context, uid int64, amount float64) error
// }

// type BalanceInfo interface {
// 	BalanceInfoDB(ctx context.Context, uid int64) (userID int64, balance float64, err error)
// }

// type WritingOff interface {
// 	WritingOffDB(ctx context.Context, uid int64, amount float64) (userID int64, amountWritingOff float64, err error)
// }

// type AddDescription interface {
// 	AddDescriptionDB(ctx context.Context, uid int64, balanceAtMoment float64, correctAmount float64, refill string, description string, senderReceiver string) error
// }

// type GetDescriptions interface {
// 	GetDescriptionsDB(ctx context.Context, uid int64, sortBy string, orderBy string) (descriptionsList []tech_task.Description, err error)
// }

// type Repository struct {
// 	UpBalance
// 	BalanceInfo
// 	WritingOff
// 	AddDescription
// 	GetDescriptions
// }

// func NewRepository(db *sql.DB) *Repository {
// 	return &Repository{
// 		UpBalance:       NewUpBalancePostgres(db),
// 		BalanceInfo:     NewBalanceInfoPostgres(db),
// 		WritingOff:      NewWritingOffPostgres(db),
// 		AddDescription:  NewAddDescriptionPostgres(db),
// 		GetDescriptions: NewGetDescriptionsPostgres(db),
// 	}
// }
