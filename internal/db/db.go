package db

import (
	"database/sql"
	"ticket-auction-manager/openstub/internal/env"

	_ "modernc.org/sqlite"
)

func NewSysDB() (*sql.DB, error) {
	dbPath := env.GetSysDBPath()
	SysDB, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}
	err = InitDB(SysDB)
	if err != nil {
		return nil, err
	}
	return SysDB, nil
}
