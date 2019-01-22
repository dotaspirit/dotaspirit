package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"

	_ "net/http/pprof"
)

var (
	appconfig = appConfig{}
	dao       = dbDao{}
)

func init() {
	loadConfig("config/app.json", &appconfig)
	dao.Server = appconfig.DBServer
	dao.Database = appconfig.DBName
	dao.connect()
}

func main() {
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
	mux := http.NewServeMux()
	mux.HandleFunc("/", webhoookHandler)
	srv := &http.Server{
		Addr:    ":3682",
		Handler: cors.Default().Handler(mux),
	}
	log.Fatal(srv.ListenAndServe())
}
