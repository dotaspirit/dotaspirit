package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/color"
	"io"
	"os"
	"regexp"
	"strconv"
)

func secondsToTime(timer int) (hours, minutes, seconds int) {
	minutes = timer / 60
	seconds = timer % 60
	hours = minutes / 60
	minutes = minutes % 60
	return hours, minutes, seconds
}

func prettifyBigValues(value int) string {
	if value < 1000 {
		return strconv.Itoa(value)
	} else {
		result := fmt.Sprintf("%.1fk", float64(value)/1000)
		return result
	}
}

func getPlayerLane(role int) string {
	laneNames := [5]string{"Unknown", "Safe", "Mid", "Off", "Jun"}
	return laneNames[role]
}

func getLeagueName(leagueID int, leaguesData oDotaLeaguesData) string {
	for _, league := range leaguesData {
		if league.Leagueid == leagueID {
			return league.Name
		}
	}
	return "League"
}

func loadConfig(path string, target interface{}) error {
	buf := bytes.NewBuffer(nil)
	f, _ := os.Open(path)
	io.Copy(buf, f)
	f.Close()

	err := json.Unmarshal(buf.Bytes(), &target)

	if err != nil {
		return err
	}

	return nil
}

func stripPresentedBy(leagueName string) string {
	m := regexp.MustCompile(" – presented by.+$")
	poweredByRemoved := m.ReplaceAllString(string(leagueName), "")
	return poweredByRemoved
}

func parseHexColor(s string) (c color.RGBA) {
	switch len(s) {
	case 9:
		_, _ = fmt.Sscanf(s, "#%02x%02x%02x%02x", &c.R, &c.G, &c.B, &c.A)
	case 7:
		c.A = 0xff
		_, _ = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 5:
		c.A = 0xff
		_, _ = fmt.Sscanf(s, "#%1x%1x%1x%1x", &c.R, &c.G, &c.B, &c.A)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
		c.A *= 17
	case 4:
		c.A = 0xff
		_, _ = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	}
	return
}
