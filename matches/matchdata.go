package matches

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func GetMatchData(matchID int64) (err error) {
	var sMatchID = strconv.FormatInt(matchID, 10)
	var filepath = "./tmp/matches/" + sMatchID
	var url = "https://api.opendota.com/api/matches/" + sMatchID

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("[DATA] Done for match #%d\n", matchID)
	return nil
}
