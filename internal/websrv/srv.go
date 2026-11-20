package websrv

import (
	"embed"
	"net/http"
)

func NewSrv(FS embed.FS) (*http.ServeMux, error) {
	WebHandler, err := Handler(FS, "dist", "index.html")
	if err != nil {
		return nil, err
	}

	webSrv := http.NewServeMux()

	webSrv.Handle("GET /", WebHandler)

	return webSrv, nil
}
