package messages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/dotaspirit/dotaspirit/types"
)

// https://oauth.vk.com/authorize?client_id=APP_ID&redirect_uri=https://oauth.vk.com/blank.html&display=page&scope=photos,wall,offline&v=5.63&revoke=1&response_type=token

func loadLeaguesConfig() types.LeaguesConfig {
	buf := bytes.NewBuffer(nil)
	f, _ := os.Open("config/leagues.json") // Error handling elided for brevity.
	io.Copy(buf, f)                        // Error handling elided for brevity.
	f.Close()

	var jsonobject types.LeaguesConfig

	err := json.Unmarshal(buf.Bytes(), &jsonobject)

	if err != nil {
		fmt.Println("error:", err)
	}

	return jsonobject
}

func loadMatchData(matchID int64) types.Match {
	var sMatchID = strconv.FormatInt(matchID, 10)
	file, e := ioutil.ReadFile("./tmp/matches/" + sMatchID)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	var jsonobject types.Match
	err := json.Unmarshal(file, &jsonobject)

	if err != nil {
		fmt.Println("error:", err)
	}

	return jsonobject
}

func seriesTypeToText(id int) string {
	seriesTypes := []string{"bo1", "bo3", "bo5"}
	return seriesTypes[id]
}

func getLeagueHashtag(id int) string {
	leaguesData := loadLeaguesConfig()

	for _, league := range leaguesData.LeaguesData {
		if league.Leagueid == id {
			return league.Hashtags
		}
	}

	return ""
}

func makeVkText(matchID int64) string {
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
			return fmt.Sprintf("🏆 [rsltdtk|%s] %d - %d %s\n%d:%02d:%02d @ %s, %s\n#rsltdtk #dota2 %s #results@rsltdtk #s%d@rsltdtk",
				radiantName, radiantScore, direScore, direName, hours, minutes, seconds, leagueName, seriesText, leagueHashTags, seriesID)
		} else {
			return fmt.Sprintf("🏆 [rsltdtk|%s] %d - %d %s\n%d:%02d:%02d @ %s\n#rsltdtk #dota2 %s #results@rsltdtk",
				radiantName, radiantScore, direScore, direName, hours, minutes, seconds, leagueName, leagueHashTags)
		}
	} else {
		if seriesID != 0 {
			return fmt.Sprintf("%s %d - %d 🏆 [rsltdtk|%s]\n%d:%02d:%02d @ %s, %s\n#rsltdtk #dota2 %s #results@rsltdtk #s%d@rsltdtk",
				radiantName, radiantScore, direScore, direName, hours, minutes, seconds, leagueName, seriesText, leagueHashTags, seriesID)
		} else {
			return fmt.Sprintf("%s %d - %d 🏆 [rsltdtk|%s]\n%d:%02d:%02d @ %s\n#rsltdtk #dota2 %s #results@rsltdtk",
				radiantName, radiantScore, direScore, direName, hours, minutes, seconds, leagueName, leagueHashTags)
		}
	}
}

func vkGetWallUploadServer(groupID int, accessToken string) types.GetWallUploadServer {
	unResp := types.GetWallUploadServer{}
	resp, err := http.Get("https://api.vk.com/method/" + "photos.getWallUploadServer?" +
		url.Values{
			"access_token": {accessToken},
			"v":            {"5.63"},
			"group_id":     {strconv.Itoa(groupID)}}.Encode())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&unResp)
	if err != nil {
		log.Fatal(err)
	}
	return unResp
}

func postFile(filename string, targetUrl string) types.UploadResponse {
	unResp := types.UploadResponse{}
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("photo", filename)
	if err != nil {
		log.Fatal(err)
	}

	// open file handle
	fh, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		log.Fatal(err)
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&unResp)
	if err != nil {
		log.Fatal(err)
	}
	return unResp
}

func vkSavePhoto(upResp types.UploadResponse, groupID int, accessToken string) types.SavedPhoto {
	unResp := types.SavedPhoto{}
	resp, err := http.Get("https://api.vk.com/method/" + "photos.saveWallPhoto?" +
		url.Values{
			"group_id":     {strconv.Itoa(groupID)},
			"access_token": {accessToken},
			"v":            {"5.63"},
			"server":       {strconv.Itoa(upResp.Server)},
			"hash":         {upResp.Hash},
			"photo":        {upResp.Photo}}.Encode())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&unResp)
	if err != nil {
		log.Fatal(err)
	}
	return unResp
}

func SendThoseMatches(matchID int64) (err error) {
	groupID := 123
	accessToken := "token"

	upServer := vkGetWallUploadServer(groupID, accessToken)
	upLink := postFile("tmp/images/"+strconv.FormatInt(matchID, 10)+".png", upServer.Response.UploadURL)
	upPhoto := vkSavePhoto(upLink, groupID, accessToken)

	_, err = http.Get("https://api.vk.com/method/" + "wall.post?" +
		url.Values{"owner_id": {strconv.Itoa(-groupID)},
			"access_token": {accessToken},
			"v":            {"5.63"},
			"message":      {makeVkText(matchID)},
			"attachments":  {"photo" + strconv.Itoa(upPhoto.Response[0].OwnerID) + "_" + strconv.Itoa(upPhoto.Response[0].ID) + ",https://www.opendota.com/matches/" + strconv.FormatInt(matchID, 10) + "/"},
			"from_group":   {"1"}}.Encode())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[VK] Done for match #%d\n", matchID)
	return nil
}
