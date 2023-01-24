package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func webhoookHandler(rw http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	defer req.Body.Close()
	rw.Header().Add("Content-type", "application/json; charset=utf-8")
	if err != nil {
		rw.WriteHeader(http.StatusTeapot)
		rw.Write([]byte(`{"status":"fail"}`))
		log.Println(err)
		return
	}
	if req.Method != "POST" {
		rw.WriteHeader(http.StatusTeapot)
		rw.Write([]byte(`{"status":"fail"}`))
		return
	}
	var whData oDotaMatchData
	err = json.Unmarshal(body, &whData)
	if err != nil {
		rw.WriteHeader(http.StatusTeapot)
		rw.Write([]byte(`{"status":"fail"}`))
		log.Println(err)
		return
	}
	rw.Write([]byte(`{"status":"ok"}`))
	handleMatch(whData)
}
