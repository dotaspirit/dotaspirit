package main

import (
	"log"
	"strconv"
	"time"
)

func handleGetFullMatchData(matchID int64, startTime time.Time, instantUpdate bool) {
	currentTime := time.Now()
	hasSend := false
	retries := 0
	if !instantUpdate {
		time.Sleep(15 * time.Minute)
	}
	for currentTime.Sub(startTime) < time.Hour*24 && !hasSend {
		expBackoff := time.Duration(5 * retries)
		time.Sleep(expBackoff * time.Minute)
		retries++
		log.Printf("Retrying getting full match %d data (retry %d)", matchID, retries)
		matchData := getMatchData(matchID)
		isNullMatch := matchData.DireScore == 0 && matchData.RadiantScore == 0
		if matchData.Version != 0 && !isNullMatch {
			log.Printf("Found full match %d data", matchID)
			matchText := makeMatchText(matchData)
			makeMatchImage(matchData, true)
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
	matchID := whData.MatchID
	log.Printf("Received match %d from webhook", matchID)

	dbData := ""

	if appconfig.IsDebug {
		matchData := getMatchData(matchID)
		makeMatchImage(matchData, true)
	} else {
		dbData = getData(matchID)
	}

	if dbData == "" && !appconfig.IsDebug {
		markHandling(matchID)
		log.Println("Match wasn't posted yet")
		matchData := getMatchData(matchID)
		isNullMatch := matchData.DireScore == 0 && matchData.RadiantScore == 0
		if !isNullMatch {
			matchText := makeMatchText(matchData)
			makeMatchImage(matchData, false)
			postID, _ := sendMatchToVk(matchID, matchText, false)
			markSent(matchID, postID)
			log.Printf("Match posted to vk https://vk.com/wall-%d_%d", appconfig.VkGroupID, postID)
			startTime := time.Now()
			if matchData.GameMode == 2 {
				go handleGetFullMatchData(matchID, startTime, false)
			}
		}
	} else if appconfig.IsDebug {
		log.Println("Debug enabled not posting")
	} else if dbData != "done" && dbData != "" {
		log.Println("Match posted trying to edit")
		startTime := time.Now()
		go handleGetFullMatchData(matchID, startTime, true)
	} else {
		log.Println("Data was posted nothing to do here")
	}

}
