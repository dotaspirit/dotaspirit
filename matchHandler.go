package main

import "log"

func handleMatch(whData webhookData) {
	matchID := whData.MatchID
	dbMatchData, _ := dao.get(dbMatch{MatchID: matchID})
	if dbMatchData.MatchID != matchID {
		addMatch(matchID)
	}

	if dbMatchData.IsFull != true {
		matchData := getMatchData(matchID)
		matchText := makeMatchText(matchData)
		log.Println("match data is not full")
		if whData.Origin == "scanner" && dbMatchData.IsShort != true {
			log.Println("data from scanner and it's not posted yet")
			makeMatchImage(matchData, false)
			sendMatchToVk(matchID, matchText, false)
		} else if whData.Origin == "" {
			log.Println("data from parser")
			makeMatchImage(matchData, true)
			if dbMatchData.IsShort == true {
				log.Println("data was posted, now editing")
				editMatchAtVk(matchID, dbMatchData.PostID, matchText)
			} else {
				log.Println("data was never posted, now posting")
				sendMatchToVk(matchID, matchText, true)
			}
		}
	}

}
