package main

import (
	"log"
	"time"
)

func handleGetFullMatchData(matchID int64, startTime time.Time) {
	currentTime := time.Now()
	hasSend := false
	for currentTime.Sub(startTime) < time.Hour*24 && hasSend == false {
		time.Sleep(5 * time.Minute)
		log.Printf("Retrying getting full match %d data", matchID)
		matchData := getMatchData(matchID)
		if len(matchData.RadiantGoldAdv) != 0 && len(matchData.PicksBans) != 0 {
			log.Printf("Found full match %d data", matchID)
			matchText := makeMatchText(matchData)
			makeMatchImage(matchData, false)
			dbMatchData, _ := dao.get(dbMatch{MatchID: matchID})
			editMatchAtVk(matchID, dbMatchData.PostID, matchText)
			hasSend = true
		}
		currentTime = time.Now()
	}
}

func handleMatch(whData webhookData) {
	matchID := whData.MatchID
	log.Printf("Received match %d from webhook", matchID)
	dbMatchData, _ := dao.get(dbMatch{MatchID: matchID})
	if dbMatchData.MatchID != matchID {
		addMatch(matchID)
	}

	if dbMatchData.IsFull != true {
		log.Println("Full match data wasn't posted")
		matchData := getMatchData(matchID)
		matchText := makeMatchText(matchData)
		if whData.Origin == "scanner" {
			if dbMatchData.IsShort != true {
				log.Println("Data from scanner and it's not posted yet")
				makeMatchImage(matchData, false)
				sendMatchToVk(matchID, matchText, false)
				startTime := time.Now()
				go handleGetFullMatchData(matchID, startTime)
			} else {
				log.Println("Data from scanner and it was posted")
			}
		} else if whData.Origin == "" {
			log.Println("Data from parser")
			makeMatchImage(matchData, true)
			if dbMatchData.IsShort == true {
				log.Println("Data was posted, now editing")
				editMatchAtVk(matchID, dbMatchData.PostID, matchText)
			} else {
				log.Println("Data never posted, now posting")
				sendMatchToVk(matchID, matchText, true)
			}
		}
	} else {
		log.Println("Full data was posted nothing to do here")
	}

}
