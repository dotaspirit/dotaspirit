package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgraph-io/badger"
	"github.com/rs/cors"
)

var (
	appconfig appConfig
	cConfig   colorConfig
	db        *badger.DB
)

func init() {
	loadConfig("./config/app.json", &appconfig)
	db, _ = badger.Open(badger.DefaultOptions("./config/badger"))
	fmt.Println(db)
	err := loadConfig("./config/colors.json", &cConfig)
	if err != nil {
		log.Println(err.Error())
	}
}

func main() {
	defer db.Close()
	mux := http.NewServeMux()
	mux.HandleFunc("/", webhoookHandler)
	srv := &http.Server{
		Addr:    ":3682",
		Handler: cors.Default().Handler(mux),
	}
	log.Fatal(srv.ListenAndServe())
}
