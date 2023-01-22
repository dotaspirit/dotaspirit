package main

type oDotaLeaguesData []struct {
	Leagueid int    `json:"leagueid"`
	Name     string `json:"name"`
}

type oDotaSeriesData struct {
	Rows []struct {
		RadiantWin    bool `json:"radiant_win"`
		RadiantTeamID int  `json:"radiant_team_id"`
		DireTeamID    int  `json:"dire_team_id"`
	} `json:"rows"`
}

type dotaSocialTeamsData []struct {
	TeamID    int    `json:"team_id"`
	Facebook  string `json:"facebook,omitempty"`
	Instagram string `json:"instagram,omitempty"`
	Twitch    string `json:"twitch,omitempty"`
	Twitter   string `json:"twitter,omitempty"`
	Youtube   string `json:"youtube,omitempty"`
	Weibo     string `json:"weibo,omitempty"`
	Vk        string `json:"vk,omitempty"`
}

type dotaSocialLeaguesData []struct {
	LeagueID  int    `json:"leagueid"`
	Facebook  string `json:"facebook,omitempty"`
	Instagram string `json:"instagram,omitempty"`
	Twitch    string `json:"twitch,omitempty"`
	Twitter   string `json:"twitter,omitempty"`
	Youtube   string `json:"youtube,omitempty"`
	Weibo     string `json:"weibo,omitempty"`
	Vk        string `json:"vk,omitempty"`
}

type oDotaMatchData struct {
	MatchID     int64  `json:"match_id"`
	RadiantName string `json:"radiant_name"`
	DireName    string `json:"dire_name"`
	DireScore   int    `json:"dire_score"`
	DireTeamID  int    `json:"dire_team_id"`
	Duration    int    `json:"duration"`
	GameMode    int    `json:"game_mode"`
	Leagueid    int    `json:"leagueid"`
	PicksBans   []struct {
		IsPick  bool  `json:"is_pick"`
		HeroID  int   `json:"hero_id"`
		Team    int   `json:"team"`
		Order   int   `json:"order"`
		Ord     int   `json:"ord"`
		MatchID int64 `json:"match_id"`
	} `json:"picks_bans"`
	RadiantGoldAdv []int `json:"radiant_gold_adv"`
	RadiantScore   int   `json:"radiant_score"`
	RadiantTeamID  int   `json:"radiant_team_id"`
	RadiantWin     bool  `json:"radiant_win"`
	RadiantXpAdv   []int `json:"radiant_xp_adv"`
	SeriesID       int   `json:"series_id"`
	SeriesType     int   `json:"series_type"`
	League         struct {
		Leagueid int         `json:"leagueid"`
		Ticket   interface{} `json:"ticket"`
		Banner   interface{} `json:"banner"`
		Tier     string      `json:"tier"`
		Name     string      `json:"name"`
	} `json:"league"`
	RadiantTeam struct {
		TeamID  int    `json:"team_id"`
		Name    string `json:"name"`
		Tag     string `json:"tag"`
		LogoURL string `json:"logo_url"`
	} `json:"radiant_team"`
	DireTeam struct {
		TeamID  int    `json:"team_id"`
		Name    string `json:"name"`
		Tag     string `json:"tag"`
		LogoURL string `json:"logo_url"`
	} `json:"dire_team"`
	Players []struct {
		HeroID      int    `json:"hero_id"`
		LaneRole    int    `json:"lane_role"`
		Name        string `json:"name"`
		Level       int    `json:"level"`
		HeroDamage  int    `json:"hero_damage"`
		HeroHealing int    `json:"hero_healing"`
		NetWorth    int    `json:"net_worth"`
		Kills       int    `json:"kills"`
		Deaths      int    `json:"deaths"`
		Assists     int    `json:"assists"`
	}
}

type colorConfig struct {
	Bg          string   `json:"bg"`
	Text        string   `json:"text"`
	PlayerColor []string `json:"playerColor"`
	TextDire    string   `json:"textDire"`
	TextRadiant string   `json:"textRadiant"`
	TextGold    string   `json:"textGold"`
	TextDamage  string   `json:"textDamage"`
	TextHealing string   `json:"textHealing"`
	TextXP      string   `json:"textXP"`
	PlayerBack  string   `json:"playerBack"`
	Overlay     string   `json:"overlay"`
	GraphLine   string   `json:"graphLine"`
	GraphXP     string   `json:"graphXP"`
	GraphGold   string   `json:"graphGold"`
}

type appConfig struct {
	VkGroupID int    `json:"vkGroupId"`
	VkAPIkey  string `json:"vkAPIkey"`
	IsDebug   bool   `json:"isDebug"`
}

type getWallUploadServer struct {
	Response struct {
		UploadURL string `json:"upload_url"`
		AlbumID   int    `json:"album_id"`
		UserID    int    `json:"user_id"`
	} `json:"response"`
}

type uploadResponse struct {
	Server int    `json:"server"`
	Photo  string `json:"photo"`
	Hash   string `json:"hash"`
}

type savedPhoto struct {
	Response []struct {
		ID      int    `json:"id"`
		AlbumID int    `json:"album_id"`
		OwnerID int    `json:"owner_id"`
		UserID  int    `json:"user_id"`
		Text    string `json:"text"`
		Date    int    `json:"date"`
	} `json:"response"`
}

type postID struct {
	Response struct {
		PostID int `json:"post_id"`
	} `json:"response"`
}
