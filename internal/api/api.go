package api

import (
	"database/sql"
	"net/http"
	"ticket-auction-manager/openstub/internal/auth"
	"ticket-auction-manager/openstub/internal/config"
)

func NewSrv(SysDB *sql.DB) (*http.ServeMux, error) {
	apiSrv := http.NewServeMux()

	apiSrv.HandleFunc("GET /health", HealthCheck(SysDB))

	confHandler := &config.CH{}
	apiSrv.HandleFunc("GET /config", confHandler.GetConfig())

	authHandler := &auth.AH{SysDB: SysDB}
	apiSrv.HandleFunc("GET /auth", authHandler.GetUsers())
	apiSrv.HandleFunc("POST /auth", authHandler.PostUser())

	return apiSrv, nil
}
