package main

type oDotaLeaguesData []struct {
	Leagueid int    `json:"leagueid"`
	Name     string `json:"name"`
}

type oDotaPlayersData []struct {
	AccountID   int    `json:"account_id"`
	Personaname string `json:"personaname"`
	Name        string `json:"name"`
}

type oDotaTeamsData []struct {
	TeamID  int    `json:"team_id"`
	LogoURL string `json:"logo_url"`
}

type oDotaSeriesData struct {
	Rows []struct {
		RadiantWin    bool `json:"radiant_win"`
		RadiantTeamID int  `json:"radiant_team_id"`
		DireTeamID    int  `json:"dire_team_id"`
	} `json:"rows"`
}

type oDotaMatchData struct {
	MatchID  int64 `json:"match_id"`
	Duration int   `json:"duration"`
	GameMode int   `json:"game_mode"`
	Leagueid int   `json:"leagueid"`
	League   struct {
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
	RadiantName    string `json:"radiant_name"`
	RadiantTeamID  int    `json:"radiant_team_id"`
	RadiantScore   int    `json:"radiant_score"`
	RadiantWin     bool   `json:"radiant_win"`
	RadiantGoldAdv []int  `json:"radiant_gold_adv"`
	RadiantXpAdv   []int  `json:"radiant_xp_adv"`
	DireTeam       struct {
		TeamID  int    `json:"team_id"`
		Name    string `json:"name"`
		Tag     string `json:"tag"`
		LogoURL string `json:"logo_url"`
	} `json:"dire_team"`
	DireName   string `json:"dire_name"`
	DireTeamID int    `json:"dire_team_id"`
	DireScore  int    `json:"dire_score"`
	Players    []struct {
		AccountID   int    `json:"account_id"`
		HeroDamage  int    `json:"hero_damage"`
		HeroHealing int    `json:"hero_healing"`
		Level       int    `json:"level"`
		HeroID      int    `json:"hero_id"`
		TotalGold   int    `json:"total_gold"`
		Gold        int    `json:"gold"`
		GoldSpent   int    `json:"gold_spent"`
		Kills       int    `json:"kills"`
		Deaths      int    `json:"deaths"`
		Assists     int    `json:"assists"`
		Name        string `json:"name"`
		Personaname string `json:"personaname"`
		LaneRole    int    `json:"lane_role"`
	} `json:"players"`
	PicksBans []struct {
		IsPick  bool  `json:"is_pick"`
		HeroID  int   `json:"hero_id"`
		Team    int   `json:"team"`
		Order   int   `json:"order"`
		Ord     int   `json:"ord"`
		MatchID int64 `json:"match_id"`
	} `json:"picks_bans"`
	SeriesID   int `json:"series_id"`
	SeriesType int `json:"series_type"`
	Teamfights []struct {
		Start     int `json:"start"`
		End       int `json:"end"`
		LastDeath int `json:"last_death"`
		Deaths    int `json:"deaths"`
		Players   []struct {
			DeathsPos struct {
			} `json:"deaths_pos"`
			AbilityUses struct {
				TinyCraggyExterior int `json:"tiny_craggy_exterior"`
				TinyAvalanche      int `json:"tiny_avalanche"`
				TinyToss           int `json:"tiny_toss"`
				TinyTossTree       int `json:"tiny_toss_tree"`
			} `json:"ability_uses"`
			AbilityTargets struct {
			} `json:"ability_targets"`
			ItemUses struct {
				PhaseBoots int `json:"phase_boots"`
				Blink      int `json:"blink"`
				Bottle     int `json:"bottle"`
			} `json:"item_uses"`
			Killed struct {
			} `json:"killed"`
			Deaths    int `json:"deaths"`
			Buybacks  int `json:"buybacks"`
			Damage    int `json:"damage"`
			Healing   int `json:"healing"`
			GoldDelta int `json:"gold_delta"`
			XpDelta   int `json:"xp_delta"`
			XpStart   int `json:"xp_start"`
			XpEnd     int `json:"xp_end"`
		} `json:"players"`
	} `json:"teamfights"`
	Objectives []struct {
		Time       int    `json:"time"`
		Type       string `json:"type"`
		Slot       int    `json:"slot,omitempty"`
		Key        int    `json:"key,omitempty"`
		PlayerSlot int    `json:"player_slot,omitempty"`
		Unit       string `json:"unit,omitempty"`
		Team       int    `json:"team,omitempty"`
	} `json:"objectives"`
	Origin string `json:"origin"`
}

type colorConfig struct {
	Bg            string   `json:"bg"`
	PlayerColor   []string `json:"playerColor"`
	KDAColor      []string `json:"kdaColor"`
	Text          string   `json:"text"`
	TextDire      string   `json:"textDire"`
	TextRadiant   string   `json:"textRadiant"`
	TextDark      string   `json:"textDark"`
	TextGold      string   `json:"textGold"`
	TextDamage    string   `json:"textDamage"`
	TextHealing   string   `json:"textHealing"`
	TextXP        string   `json:"textXP"`
	TextTF        string   `json:"textTF"`
	TextTFDire    string   `json:"textTFDire"`
	TextTFRadiant string   `json:"textTFRadiant"`
	BackDire      string   `json:"backDire"`
	BackRadiant   string   `json:"backRadiant"`
	BackLost      string   `json:"backLost"`
	TextSize8     float64  `json:"textSize8"`
	TextSize12    float64  `json:"textSize12"`
	TextSize15    float64  `json:"textSize15"`
	TextSize20    float64  `json:"textSize20"`
	TextSize28    float64  `json:"textSize28"`
	TextSize35    float64  `json:"textSize35"`
	TextSize48    float64  `json:"textSize48"`
}

type appConfig struct {
	VkGroupID int    `json:"vkGroupId"`
	VkAPIkey  string `json:"vkAPIkey"`
	TgChatID  int64  `json:"tgChatID"`
	TgAPIkey  string `json:"tgAPIkey"`
	DBServer  string `json:"dbServer"`
	DBName    string `json:"dbName"`
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
