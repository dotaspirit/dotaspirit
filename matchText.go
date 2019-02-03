package main

import "fmt"

func seriesTypeToText(id int) string {
	seriesTypes := []string{"bo1", "bo3", "bo5"}
	return seriesTypes[id]
}

func makeMatchText(matchData oDotaMatchData) string {
	radiantName := matchData.RadiantTeam.Name
	direName := matchData.DireTeam.Name
	radiantScore := matchData.RadiantScore
	direScore := matchData.DireScore
	hours, minutes, seconds := secondsToTime(matchData.Duration)
	leagueName := matchData.League.Name
	seriesID := matchData.SeriesID
	seriesType := matchData.SeriesType

	if radiantName == "" {
		radiantName = matchData.RadiantName
	}
	if direName == "" {
		direName = matchData.DireName
	}

	if radiantName == "" {
		radiantName = "Radiant"
	}
	if direName == "" {
		direName = "Dire"
	}

	if leagueName == "" {
		leaguesData := getLeaguesData()
		leagueName = getLeagueName(matchData.Leagueid, leaguesData)
	}

	seriesText := seriesTypeToText(seriesType)

	if matchData.RadiantWin == true {
		if seriesID != 0 {
			return fmt.Sprintf("🏆 [rsltdtk|%s] %d - %d %s\n%d:%02d:%02d @ %s, %s\n#rsltdtk #dota2 #results@rsltdtk #s%d@rsltdtk",
				radiantName, radiantScore, direScore, direName, hours, minutes, seconds, leagueName, seriesText, seriesID)
		} else {
			return fmt.Sprintf("🏆 [rsltdtk|%s] %d - %d %s\n%d:%02d:%02d @ %s\n#rsltdtk #dota2 #results@rsltdtk",
				radiantName, radiantScore, direScore, direName, hours, minutes, seconds, leagueName)
		}
	} else {
		if seriesID != 0 {
			return fmt.Sprintf("%s %d - %d 🏆 [rsltdtk|%s]\n%d:%02d:%02d @ %s, %s\n#rsltdtk #dota2 #results@rsltdtk #s%d@rsltdtk",
				radiantName, radiantScore, direScore, direName, hours, minutes, seconds, leagueName, seriesText, seriesID)
		} else {
			return fmt.Sprintf("%s %d - %d 🏆 [rsltdtk|%s]\n%d:%02d:%02d @ %s\n#rsltdtk #dota2 #results@rsltdtk",
				radiantName, radiantScore, direScore, direName, hours, minutes, seconds, leagueName)
		}
	}
}
