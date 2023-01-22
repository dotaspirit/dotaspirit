package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func webhoookHandler(rw http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	defer req.Body.Close()
	rw.Header().Add("Content-type", "application/json; charset=utf-8")
	if err != nil {
		rw.Write([]byte(`{"status":"fail"}`))
		panic(err)
	}
	if len(body) == 0 {
		rw.Write([]byte(`{"status":"fail"}`))
		panic(err)
	}
	var whData oDotaMatchData
	err = json.Unmarshal(body, &whData)
	if err != nil {
		rw.Write([]byte(`{"status":"fail"}`))
		panic(err)
	}
	rw.Write([]byte(`{"status":"ok"}`))
	handleMatch(whData)
}
