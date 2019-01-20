package main

func handleMatch(whData webhookData) {
	matchID := whData.MatchID
	dbMatchData, _ := dao.get(dbMatch{MatchID: matchID})
	if dbMatchData.MatchID != matchID {
		addMatch(matchID)
	}

	if dbMatchData.IsFull != true {
		matchData := getMatchData(matchID)
		matchText := makeMatchText(matchData)
		if whData.Origin == "scanner" && (dbMatchData.IsShort != true ||
			dbMatchData.IsFull != true) {
			makeMatchImage(matchData, false)
			sendMatchToVk(matchID, matchText, false)
		} else if whData.Origin == "" {
			makeMatchImage(matchData, true)
			if dbMatchData.IsShort == true {
				editMatchAtVk(matchID, dbMatchData.PostID, matchText)
			} else {
				sendMatchToVk(matchID, matchText, true)
			}
		}
	}

}
