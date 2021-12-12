package repository

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	driver = "pgx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	Timeout  int
}

func NewConnection(cfg *Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable&connect_timeout=%d",
		"postgres",
		url.QueryEscape(cfg.Username),
		url.QueryEscape(cfg.Password),
		cfg.Host,
		cfg.Port,
		cfg.DbName,
		cfg.Timeout)

	db, err := sql.Open(driver, connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
