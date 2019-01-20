package main

import (
	"fmt"
	"log"
)

const (
	apiURL = "https://api.opendota.com/api/"
)

func getMatchData(matchID int64) oDotaMatchData {
	var matchData oDotaMatchData
	matchURL := fmt.Sprintf("%s%s/%d", apiURL, "matches", matchID)
	log.Println(matchID)
	getJSON(matchURL, &matchData)
	return matchData
}
