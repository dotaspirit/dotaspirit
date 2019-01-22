package main

import (
	"log"
	"net/http"
	"time"

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
	server := &http.Server{
		Addr:         ":8080",
		Handler:      http.DefaultServeMux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
	log.Fatal(server.ListenAndServe())

	mux := http.NewServeMux()
	mux.HandleFunc("/", webhoookHandler)
	srv := &http.Server{
		Addr:    ":3682",
		Handler: cors.Default().Handler(mux),
	}
	log.Fatal(srv.ListenAndServe())
}
