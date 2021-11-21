package pg

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"tech_task/internal/godb"

	"github.com/jackc/pgx/v4/pgxpool"
)

//Структура конфига, которая включает в себя необходимые нам настройки соединения (сюда можно добавить любые другие поля для postgres типа ssl и т.д.)
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	Timeout  int
}

func NewPoolConfig(cfg *Config) (*pgxpool.Config, error) {
	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable&connect_timeout=%d",
		"postgres",
		url.QueryEscape(cfg.Username),
		url.QueryEscape(cfg.Password),
		cfg.Host,
		cfg.Port,
		cfg.DbName,
		cfg.Timeout)

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	return poolConfig, nil
}

func NewConnection(poolConfig *pgxpool.Config) (*pgxpool.Pool, error) {
	conn, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func StartDB() *godb.Instance {
	cfg := &Config{}

	if len(os.Getenv("POSTGRES_HOST")) == 0 {
		cfg.Host = "localhost"
	} else {
		cfg.Host = os.Getenv("POSTGRES_HOST")
	}

	if len(os.Getenv("POSTGRES_USER")) == 0 {
		cfg.Username = "postgres"
	} else {
		cfg.Username = os.Getenv("POSTGRES_USER")
	}

	if len(os.Getenv("POSTGRES_PASSWORD")) == 0 {
		cfg.Password = "1234"
	} else {
		cfg.Password = os.Getenv("POSTGRES_PASSWORD")
	}

	if len(os.Getenv("POSTGRES_DB")) == 0 {
		cfg.DbName = "avito_users_db"
	} else {
		cfg.DbName = os.Getenv("POSTGRES_DB")
	}

	cfg.Port = "54320"
	cfg.Timeout = 5

	poolConfig, err := NewPoolConfig(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Pool config error: %v\n", err)
		os.Exit(1)
	}

	poolConfig.MaxConns = 20

	c, err := NewConnection(poolConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connect to database failed: %v\n", err)
		os.Exit(1)
	}

	_, err = c.Exec(context.Background(), ";")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ping failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Ping OK!")

	ins := &godb.Instance{Db: c}
	return ins
}
