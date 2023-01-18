package main

import (
	"encoding/json"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
)

func getJSON(url string, target interface{}) error {
	r, err := retryablehttp.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
