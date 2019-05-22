package main

import (
	"fmt"
	"log"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
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

func forceScan(matchID int64) {
	matchURL := fmt.Sprintf("%s%s/%d", apiURL, "request", matchID)
	log.Printf("Force scan match %d data", matchID)
	r, _ := retryablehttp.Post(matchURL, "", nil)

	defer r.Body.Close()
}

func getLeaguesData() oDotaLeaguesData {
	var leaguesData oDotaLeaguesData
	leaguesURL := fmt.Sprintf("%s%s", apiURL, "leagues")
	getJSON(leaguesURL, &leaguesData)
	return leaguesData
}

func getPlayersData() oDotaPlayersData {
	var playersData oDotaPlayersData
	playersURL := fmt.Sprintf("%s%s", apiURL, "proPlayers")
	getJSON(playersURL, &playersData)
	return playersData
}
