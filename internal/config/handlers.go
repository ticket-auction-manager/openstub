package config

import (
	"encoding/json"
	"net/http"
)

type CH struct{}

func (ch *CH) GetConfig() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conf, err := GetConfig()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(conf)
	}
}
