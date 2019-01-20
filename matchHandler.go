package main

import (
	"log"
)

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
		if whData.Origin == "scanner" && dbMatchData.IsShort != true {
			log.Println("Data from scanner and it's not posted yet")
			makeMatchImage(matchData, false)
			sendMatchToVk(matchID, matchText, false)
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
