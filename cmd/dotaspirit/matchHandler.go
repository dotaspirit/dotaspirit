package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func handleGetFullMatchData(matchID int64, startTime time.Time) {
	currentTime := time.Now()
	hasSend := false
	retries := 0
	for currentTime.Sub(startTime) < time.Hour*24 && !hasSend {
		expBackoff := time.Duration(5 * retries)
		time.Sleep(expBackoff * time.Minute)
		retries++
		log.Printf("Retrying getting full match %d data (retry %d)", matchID, retries)
		matchData := getMatchData(matchID)
		isNullMatch := matchData.DireScore == 0 && matchData.RadiantScore == 0
		if len(matchData.RadiantGoldAdv) != 0 &&
			len(matchData.PicksBans) != 0 &&
			!isNullMatch {
			log.Printf("Found full match %d data", matchID)
			matchText := makeMatchText(matchData)
			makeMatchImage(matchData)
			sPostID := getData(matchID)
			postID, _ := strconv.Atoi(sPostID)
			editMatchAtVk(matchID, postID, matchText)
			hasSend = true
		}
		if retries%12 == 0 {
			forceScan(matchID)
		}
		currentTime = time.Now()
	}
}

func handleMatch(whData oDotaMatchData) {
	isNullMatch := whData.DireScore == 0 && whData.RadiantScore == 0
	fmt.Println(isNullMatch)
	matchID := whData.MatchID
	log.Printf("Received match %d from webhook", matchID)

	// makeMatchImage(whData, false)

	if getData(matchID) == "" && !isNullMatch {
		log.Println("Match wasn't posted yet")
		matchData := whData
		matchText := makeMatchText(matchData)
		makeMatchImage(matchData)
		_, postID := sendMatchToVk(matchID, matchText, false)
		markSent(matchID, postID)
		startTime := time.Now()
		if matchData.GameMode == 2 {
			go handleGetFullMatchData(matchID, startTime)
		}
	} else {
		log.Println("Data was posted nothing to do here")
	}

}
