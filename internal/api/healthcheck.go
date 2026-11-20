package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type HC struct {
	Healthy         bool   `json:"healthy"`
	IsSetup         bool   `json:"isSetup"`
	OpenStubVersion string `json:"openStubVersion"`
}

func HealthCheck(SysDB *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var AdminCount int
		err := SysDB.QueryRow("SELECT COUNT(*) FROM users").Scan(&AdminCount)
		if err != nil {
			http.Error(w, "Can't get admin count.", http.StatusInternalServerError)
			return
		}
		rtnCheck := new(HC)
		if AdminCount > 0 {
			rtnCheck.IsSetup = true
		} else {
			rtnCheck.IsSetup = false
		}
		rtnCheck.Healthy = true
		rtnCheck.OpenStubVersion = "dev"

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(rtnCheck)
	}
}
