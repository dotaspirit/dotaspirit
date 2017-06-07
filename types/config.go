package types

type LeaguesConfig struct {
	Whitelist   []int `json:"whitelist"`
	LeaguesData []struct {
		Leagueid int    `json:"leagueid"`
		Hashtags string `json:"hashtags"`
	} `json:"leaguesData"`
}
