package db

import (
	"database/sql"
	"log"
)

func InitDB(SysDB *sql.DB) error {
	Tx, err := SysDB.Begin()
	if err != nil {
		return err
	}
	defer Tx.Rollback()
	_, err = Tx.Exec(`CREATE TABLE IF NOT EXISTS users (
		username TEXT NOT NULL PRIMARY KEY,
		password TEXT,
		salt TEXT
		)`)
	if err != nil {
		return err
	}

	Tx.Commit()
	log.Println("SysDB initialized successfully.")
	return nil
}
