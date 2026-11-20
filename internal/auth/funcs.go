package auth

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"ticket-auction-manager/openstub/internal/env"
)

func saltAndHash(password, salt string) string {
	hBytes := sha256.Sum256([]byte(password + salt))
	return hex.EncodeToString(hBytes[:])
}

func CheckUser(SysDB *sql.DB, username, password string) bool {
	var pwHash, salt string
	err := SysDB.QueryRow("SELECT password, salt FROM users WHERE username = ?", username).Scan(&pwHash, &salt)
	if err != nil {
		return false
	}
	outHash := sha256.Sum256([]byte(password + salt))
	outStr := hex.EncodeToString(outHash[:])
	if pwHash == outStr {
		return true
	} else {
		return false
	}
}

func NewUpdateUser(SysDB *sql.DB, username, password string) error {
	Tx, err := SysDB.Begin()
	if err != nil {
		return err
	}
	defer Tx.Rollback()
	salt := env.GenStrBytes(32)
	dbPw := saltAndHash(password, salt)
	_, err = Tx.Exec("INSERT INTO users VALUES (?, ?, ?) ON CONFLICT (username) DO UPDATE SET password = EXCLUDED.password, salt = EXCLUDED.salt", username, dbPw, salt)
	if err != nil {
		return err
	}
	err = Tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
