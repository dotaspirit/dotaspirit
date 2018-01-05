package messages

import (
	"fmt"
	"log"
	"strconv"

	"github.com/dotaspirit/dotaspirit/utils"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func makeTGText(matchID int64) string {
	matchData := loadMatchData(matchID)
	radiantName := matchData.RadiantTeam.Name
	direName := matchData.DireTeam.Name
	radiantScore := matchData.RadiantScore
	direScore := matchData.DireScore
	matchDuration := matchData.Duration
	minutes := matchDuration / 60
	seconds := matchDuration % 60
	hours := minutes / 60
	minutes = minutes % 60
	leagueName := matchData.League.Name
	seriesID := matchData.SeriesID
	seriesType := matchData.SeriesType
	leagueID := matchData.League.Leagueid

	if radiantName == "" {
		radiantName = "Radiant"
	}
	if direName == "" {
		direName = "Dire"
	}

	seriesText := seriesTypeToText(seriesType)
	leagueHashTags := getLeagueHashtag(leagueID)

	if matchData.RadiantWin == true {
		if seriesID != 0 {
			return fmt.Sprintf("🏆 %s %d - %d %s\n%d:%02d:%02d @ %s, %s\n#rsltdtk #dota2 %s #s%d",
				radiantName, radiantScore, direScore, direName, hours, minutes, seconds, leagueName, seriesText, leagueHashTags, seriesID)
		} else {
			return fmt.Sprintf("🏆 %s %d - %d %s\n%d:%02d:%02d @ %s\n#rsltdtk #dota2 %s",
				radiantName, radiantScore, direScore, direName, hours, minutes, seconds, leagueName, leagueHashTags)
		}
	} else {
		if seriesID != 0 {
			return fmt.Sprintf("%s %d - %d 🏆 %s\n%d:%02d:%02d @ %s, %s\n#rsltdtk #dota2 %s #s%d",
				radiantName, radiantScore, direScore, direName, hours, minutes, seconds, leagueName, seriesText, leagueHashTags, seriesID)
		} else {
			return fmt.Sprintf("%s %d - %d 🏆 %s\n%d:%02d:%02d @ %s\n#rsltdtk #dota2 %s",
				radiantName, radiantScore, direScore, direName, hours, minutes, seconds, leagueName, leagueHashTags)
		}
	}
}

func SendToTG(matchID int64) {
	appConfig := utils.LoadAppConfig()

	chatID := appConfig.TgChatID
	token := appConfig.TgAPIkey

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	moreInfo := []tgbotapi.InlineKeyboardButton{tgbotapi.NewInlineKeyboardButtonURL("More Info", fmt.Sprintf("https://www.opendota.com/matches/%d", matchID))}
	kbd := tgbotapi.NewInlineKeyboardMarkup(moreInfo)

	image := tgbotapi.NewPhotoUpload(chatID, "tmp/images/"+strconv.FormatInt(matchID, 10)+".png")
	image.Caption = makeTGText(matchID)
	image.ReplyMarkup = kbd
	bot.Send(image)
}
