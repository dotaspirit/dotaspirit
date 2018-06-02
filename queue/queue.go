package queue

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dotaspirit/dotaspirit/images"
	"github.com/dotaspirit/dotaspirit/matches"
	"github.com/dotaspirit/dotaspirit/messages"
	"github.com/dotaspirit/dotaspirit/types"
	"github.com/dotaspirit/dotaspirit/utils"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

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

func compareMatches(X, Y []int64) []int64 {
	m := make(map[int64]int64)

	for _, y := range Y {
		m[y]++
	}

	var ret []int64
	for _, x := range X {
		if m[x] > 0 {
			m[x]--
			continue
		}
		ret = append(ret, x)
	}

	return ret
}

func getOldMatches() []int64 {
	oldMatches := []int64{}

	file, err := os.Open("tmp/oldmatches")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matchID, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		oldMatches = append(oldMatches, matchID)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return oldMatches
}

func getMatches() []int64 {
	config := loadLeaguesConfig()

	odProMatches := types.ProMatches{}

	newMatches := []int64{}
	oldMatches := getOldMatches()

	getJson("https://api.opendota.com/api/proMatches", &odProMatches)

	for _, match := range odProMatches {
		for _, leagueID := range config.Whitelist {
			if leagueID == match.Leagueid {
				newMatches = append(newMatches, match.MatchID)
			}
		}
	}

	diffMatches := compareMatches(newMatches, oldMatches)

	return diffMatches
}

func saveDoneMatches(matchID int64, oldMatches []int64) {
	file, err := os.Create("tmp/oldmatches")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprintf(file, "%d\n", matchID)

	for i, match := range oldMatches {
		if i < 99 {
			fmt.Fprintf(file, "%d\n", match)
		}
	}
}

func QueueMatches(ret map[int64]int) {
	appConfig := utils.LoadAppConfig()
	matchIDs := getMatches()
	for i := len(matchIDs) - 1; i >= 0; i-- {
		matchID := matchIDs[i]
		if i < 100 {
			oldMatches := getOldMatches()
			err := matches.GetMatchData(matchID)
			if err != nil {
				log.Println(err)
				continue
			}
			time.Sleep(time.Second)
			err = images.MakeResultImage(matchID, ret)
			if err != nil {
				log.Println(err)
				ret[matchID]++
				continue
			}
			if !appConfig.IsDebug {
				err = messages.SendThoseMatches(matchID)
				if err != nil {
					log.Println(err)
					continue
				}
				//messages.SendToTG(matchID)
				saveDoneMatches(matchID, oldMatches)
			}
		}
	}
}
