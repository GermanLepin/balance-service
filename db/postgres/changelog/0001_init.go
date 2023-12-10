package changelog

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upInit, downInit)
}

func upInit(tx *sql.Tx) error {
	_, err := tx.Exec(`
		create schema service;

		create table service.users (
			id uuid not null primary key, 
			name varchar(100),
			balance numeric(10,2)
		);
	`)
	if err != nil {
		return err
	}

	return nil
}

func downInit(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
