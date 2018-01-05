package types

type LeaguesConfig struct {
	Whitelist   []int `json:"whitelist"`
	LeaguesData []struct {
		Leagueid int    `json:"leagueid"`
		Hashtags string `json:"hashtags"`
	} `json:"leaguesData"`
}

type AppConfig struct {
	VkGroupID int    `json:"vkGroupId"`
	VkAPIkey  string `json:"vkAPIkey"`
	TgChatID  int64  `json:"tgChatID"`
	TgAPIkey  string `json:"tgAPIkey"`
	IsDebug   bool   `json:"isDebug"`
}
