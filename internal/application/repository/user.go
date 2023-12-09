package repository

import (
	"balance-service/internal/application/dto"
	"context"
	"database/sql"

	"github.com/google/uuid"
)

func (u *userRepository) CreateUser(ctx context.Context, user dto.User) error {
	err := u.db.QueryRow("insert into service.users (id, name, balance) values ($1,$2,$3);", user.ID, user.Name, user.Balance)
	if err != nil {
		return err.Err()
	}

	return nil
}

func (u *userRepository) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	err := u.db.QueryRow("delete from service.users where id = $1;", userID)
	if err != nil {
		return err.Err()
	}

	return nil
}

func (u *userRepository) FetchUser(ctx context.Context, userID uuid.UUID) (dto.User, error) {
	var user dto.User

	err := u.db.QueryRow("select * from service.users where id = $1;", userID).Scan(&user.ID, &user.Name, &user.Balance)
	if err != nil {
		return user, err
	}

	return user, nil
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}
