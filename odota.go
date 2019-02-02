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
