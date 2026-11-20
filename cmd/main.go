package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"net/http"
	"ticket-auction-manager/openstub/internal/api"
	"ticket-auction-manager/openstub/internal/config"
	"ticket-auction-manager/openstub/internal/db"
	"ticket-auction-manager/openstub/internal/websrv"
)

//go:embed all:dist
var embedFS embed.FS

func main() {
	addr := flag.String("addr", "0.0.0.0", "Sets the address to listen to.")
	port := flag.String("port", "8000", "Sets the listening port.")

	flag.Parse()

	ls := fmt.Sprintf("%s:%s", *addr, *port)

	sysConf, err := config.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Read config for", sysConf.DeviceName+".")

	SysDB, err := db.NewSysDB()
	if err != nil {
		log.Fatalln(err)
	}

	mainSrv := http.NewServeMux()

	// Creates an API Server
	apiSrv, err := api.NewSrv(SysDB)
	if err != nil {
		log.Fatalln(err)
	}

	// Add API Server
	mainSrv.Handle("/api/", http.StripPrefix("/api", apiSrv))

	// Creates a webserver instance
	webSrv, err := websrv.NewSrv(embedFS)
	if err != nil {
		log.Fatalln(err)
	}

	// Add Webserver
	mainSrv.Handle("/web/", http.StripPrefix("/web", webSrv))

	mainSrv.Handle("GET /{$}", http.RedirectHandler("/web/", http.StatusPermanentRedirect))

	log.Println("Listening on http://" + ls + "/")
	log.Println("Web GUI should be available at http://" + ls + "/web/")
	http.ListenAndServe(ls, mainSrv)
}
