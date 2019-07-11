package main

import (
	"log"
	"time"
)

func handleGetFullMatchData(matchID int64, startTime time.Time) {
	currentTime := time.Now()
	hasSend := false
	retries := 0
	for currentTime.Sub(startTime) < time.Hour*24 && !hasSend {
		time.Sleep(5 * time.Minute)
		retries++
		log.Printf("Retrying getting full match %d data (retry %d)", matchID, retries)
		matchData := getMatchData(matchID)
		isNullMatch := matchData.DireScore == 0 && matchData.RadiantScore == 0
		if len(matchData.RadiantGoldAdv) != 0 &&
			len(matchData.PicksBans) != 0 &&
			!isNullMatch {
			log.Printf("Found full match %d data", matchID)
			matchText := makeMatchText(matchData)
			makeMatchImage(matchData, true)
			dbMatchData, _ := dao.get(dbMatch{MatchID: matchID})
			editMatchAtVk(matchID, dbMatchData.PostID, matchText)
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
	dbMatchData, _ := dao.get(dbMatch{MatchID: matchID})
	if dbMatchData.MatchID != matchID {
		addMatch(matchID)
	}
	if dbMatchData.IsFull != true {
		log.Println("Full match data wasn't posted")
		matchData := whData
		matchText := makeMatchText(matchData)
		if whData.Origin == "scanner" {
			if dbMatchData.IsShort != true {
				log.Println("Data from scanner and it's not posted yet")
				makeMatchImage(matchData, false)
				_, vkpost := sendMatchToVk(matchID, matchText, false)
				makeStoryImage(matchData)
				sendStoryToVK(matchID, vkpost)
				startTime := time.Now()
				if matchData.GameMode == 2 {
					go handleGetFullMatchData(matchID, startTime)
				}
			} else {
				log.Println("Data from scanner and it was posted")
			}
		} else if whData.Origin == "" {
			log.Println("Data from parser")
			matchData := getMatchData(matchID)
			makeMatchImage(matchData, true)
			if dbMatchData.IsShort == true {
				log.Println("Data was posted, now editing")
				editMatchAtVk(matchID, dbMatchData.PostID, matchText)
			} else {
				log.Println("Data never posted, now posting")
				_, vkpost := sendMatchToVk(matchID, matchText, true)
				makeStoryImage(matchData)
				sendStoryToVK(matchID, vkpost)
			}
		}
	} else {
		log.Println("Full data was posted nothing to do here")
	}

}
