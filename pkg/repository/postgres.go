package repository

import (
	"database/sql"
	"fmt"
	"net/url"
	_ "tech_task/migrations"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose"
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

	conn, err := sql.Open(driver, connStr)
	if err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(25)
	conn.SetMaxIdleConns(25)
	conn.SetConnMaxLifetime(5 * time.Minute)

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	if err := goose.Up(conn, "/var"); err != nil {
		return nil, err
	}

	return conn, nil
}
