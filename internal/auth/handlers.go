package auth

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type AH struct {
	SysDB *sql.DB
}

var nodetails string = "No basic auth details provided."
var notauthenticated string = "Credentials did not authenticate."
var createduser string = "Successfully created/updated user."

func (ah *AH) GetUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			http.Error(w, nodetails, http.StatusBadRequest)
			return
		}
		doesAuth := CheckUser(ah.SysDB, username, password)
		if !doesAuth {
			http.Error(w, notauthenticated, http.StatusUnauthorized)
			return
		}
		var rtnUsers []string
		rows, err := ah.SysDB.Query("SELECT username FROM users ORDER BY username")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for rows.Next() {
			var user string
			rows.Scan(&user)
			rtnUsers = append(rtnUsers, user)
		}
		if err = rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(rtnUsers)
	}
}

func (ah *AH) PostUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var inUser User
		err := json.NewDecoder(r.Body).Decode(&inUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		var AdminCount int
		err = ah.SysDB.QueryRow("SELECT COUNT(*) FROM users").Scan(&AdminCount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if AdminCount == 0 {
			err = NewUpdateUser(ah.SysDB, inUser.Username, inUser.Password)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			username, password, ok := r.BasicAuth()
			if !ok {
				http.Error(w, nodetails, http.StatusBadRequest)
				return
			}
			if !CheckUser(ah.SysDB, username, password) {
				http.Error(w, notauthenticated, http.StatusUnauthorized)
				return
			}
			err = NewUpdateUser(ah.SysDB, inUser.Username, inUser.Password)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(createduser))
	}
}
