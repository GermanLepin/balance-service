package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upInit, downInit)
}

func upInit(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`
	CREATE TABLE IF NOT EXISTS users(
		id      INTEGER NOT NULL PRIMARY KEY,
		balance NUMERIC(10,2));

	CREATE TABLE IF NOT EXISTS descriptions(
		id_description       SERIAL PRIMARY KEY, 
		sender_receiver      VARCHAR(100), 
		amount               NUMERIC(10,2), 
		description          VARCHAR(255), 
		balance_at_moment    NUMERIC(10,2), 
		user_id              INTEGER, 
		FOREIGN KEY(user_id) REFERENCES users (id), 
		created_at TIMESTAMP NOT NULL, 
		refill VARCHAR(100)  NOT NULL);
	`)
	if err != nil {
		return err
	}
	return nil
}

func downInit(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`
	DROP TABLE IF EXISTS users;
	DROP TABLE IF EXISTS descriptions;`)
	if err != nil {
		return err
	}
	return nil
}
