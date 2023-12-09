package connection

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "balance-service/db/postgres/changelog"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose"
)

var (
	driver = "pgx"
)

func StartDB() *sql.DB {
	conn := connectToDB()
	if conn == nil {
		log.Panic("cannot connect to Postgres")
	}

	if err := goose.Up(conn, "/var"); err != nil {
		log.Panic("cannot run the migrations")
	}

	// if smth goes wrong we always can run down Migrations goose.Down()

	return conn
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	var counts int64

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("postgres is not ready yet")
			counts++
		} else {
			log.Println("connected to Postgres")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("backing off for 2 seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}

func openDB(dsn string) (*sql.DB, error) {
	conn, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(25)
	conn.SetMaxIdleConns(25)
	conn.SetConnMaxLifetime(5 * time.Minute)

	if err = conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}
