package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func webhoookHandler(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		rw.Write([]byte(`{"status":fail}`))
		panic(err)
	}
	var whData oDotaMatchData
	err = json.Unmarshal(body, &whData)
	if err != nil {
		rw.Write([]byte(`{"status":fail}`))
		panic(err)
	}
	rw.Write([]byte(`{"status":ok}`))
	handleMatch(whData)
}
