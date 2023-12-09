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

func (u *userRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	err := u.db.QueryRow("delete from service.users where id = $1;", id)
	if err != nil {
		return err.Err()
	}

	return nil
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}
