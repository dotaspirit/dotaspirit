package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func webhoookHandler(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	var whData webhookData
	err = json.Unmarshal(body, &whData)
	if err != nil {
		panic(err)
	}
	handleMatch(whData)
}
