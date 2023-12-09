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
			id integer not null primary key, 
			name varchar(100),
			balance numeric(10,2)
		);

		create table service.descriptions (
			id_description serial primary key, 
			sender_receiver varchar(100), 
			amount numeric(10,2), 
			description varchar(255), 
			balance_at_moment numeric(10,2), 
			user_id integer, 
			foreign key(user_id) references service.users(id), 
			created_at timestamp not null, 
			refill varchar(100) not null
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
