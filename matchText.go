package main

import (
	"fmt"
)

func seriesTypeToText(id int) string {
	seriesTypes := []string{"bo1", "bo3", "bo5"}
	return seriesTypes[id]
}

func teamsToVS(radiantTeam, direTeam int) string {
	if radiantTeam > direTeam {
		return fmt.Sprintf("#t%dvs%d@rsltdtk", direTeam, radiantTeam)
	} else {
		return fmt.Sprintf("#t%dvs%d@rsltdtk", radiantTeam, direTeam)
	}
}

func seriesScoreToText(radiantID, direID, seriesID int, matchID int64) string {
	seriesData := getSeriesData(matchID, seriesID)
	radiantScore := 0
	direScore := 0
	for _, match := range seriesData.Rows {
		if (match.RadiantWin && match.RadiantTeamID == radiantID) || (!match.RadiantWin && match.DireTeamID == radiantID) {
			radiantScore++
		} else {
			direScore++
		}
	}
	return fmt.Sprintf("[%d:%d]", radiantScore, direScore)
}

func teamIDtoMention(teamID int, teamName string) string {
	var teamsData dotaSocialTeamsData
	teamsURL := "https://raw.githubusercontent.com/dotaspirit/dotasocial/master/teams.json"
	getJSON(teamsURL, &teamsData)
	for _, team := range teamsData {
		if team.TeamID == teamID {
			vk := team.Vk
			if vk != "" {
				return fmt.Sprintf("[%s|%s]", vk, teamName)
			}
		}
	}
	return teamName
}

func makeMatchText(matchData oDotaMatchData) string {
	radiantName := matchData.RadiantTeam.Name
	direName := matchData.DireTeam.Name
	radiantID := matchData.RadiantTeamID
	direID := matchData.DireTeamID
	radiantScore := matchData.RadiantScore
	direScore := matchData.DireScore
	hours, minutes, seconds := secondsToTime(matchData.Duration)
	leagueName := matchData.League.Name
	leagueID := matchData.Leagueid
	seriesID := matchData.SeriesID
	seriesType := matchData.SeriesType
	matchID := matchData.MatchID

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
	vsText := teamsToVS(radiantID, direID)
	direText := teamIDtoMention(direID, direName)
	radiantText := teamIDtoMention(radiantID, radiantName)
	seriesScoreText := "[0:0]"
	if seriesID != 0 {
		seriesScoreText = seriesScoreToText(radiantID, direID, seriesID, matchID)
	}

	if matchData.RadiantWin {
		if seriesID != 0 {
			return fmt.Sprintf("🏆 %s %d - %d %s\n%d:%02d:%02d @ %s, %s %s\n#dota2 #l%d@rsltdtk #s%d@rsltdtk %s",
				radiantText, radiantScore, direScore, direText, hours, minutes, seconds, leagueName, seriesScoreText, seriesText, leagueID, seriesID, vsText)
		} else {
			return fmt.Sprintf("🏆 %s %d - %d %s\n%d:%02d:%02d @ %s\n#dota2 #l%d@rsltdtk %s",
				radiantText, radiantScore, direScore, direText, hours, minutes, seconds, leagueName, leagueID, vsText)
		}
	} else {
		if seriesID != 0 {
			return fmt.Sprintf("%s %d - %d 🏆 %s\n%d:%02d:%02d @ %s, %s %s\n#dota2 #l%d@rsltdtk #s%d@rsltdtk %s",
				radiantText, radiantScore, direScore, direText, hours, minutes, seconds, leagueName, seriesScoreText, seriesText, leagueID, seriesID, vsText)
		} else {
			return fmt.Sprintf("%s %d - %d 🏆 %s\n%d:%02d:%02d @ %s\n#dota2 #l%d@rsltdtk %s",
				radiantText, radiantScore, direScore, direText, hours, minutes, seconds, leagueName, leagueID, vsText)
		}
	}
}
