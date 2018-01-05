package types

type ProMatches []struct {
	MatchID       int64  `json:"match_id"`
	Duration      int    `json:"duration"`
	StartTime     int    `json:"start_time"`
	RadiantTeamID int    `json:"radiant_team_id"`
	RadiantName   string `json:"radiant_name"`
	DireTeamID    int    `json:"dire_team_id"`
	DireName      string `json:"dire_name"`
	Leagueid      int    `json:"leagueid"`
	LeagueName    string `json:"league_name"`
	SeriesID      int    `json:"series_id"`
	SeriesType    int    `json:"series_type"`
	RadiantWin    bool   `json:"radiant_win"`
}

type Match struct {
	MatchID               int64 `json:"match_id"`
	BarracksStatusDire    int   `json:"barracks_status_dire"`
	BarracksStatusRadiant int   `json:"barracks_status_radiant"`
	Chat                  []struct {
		Time       int    `json:"time"`
		Type       string `json:"type"`
		Unit       string `json:"unit,omitempty"`
		Key        string `json:"key"`
		Slot       int    `json:"slot"`
		PlayerSlot int    `json:"player_slot"`
	} `json:"chat"`
	Cluster   int `json:"cluster"`
	Cosmetics struct {
	} `json:"-"`
	DireScore    int `json:"dire_score"`
	DraftTimings []struct {
		Order          int         `json:"order"`
		Pick           bool        `json:"pick"`
		ActiveTeam     int         `json:"active_team"`
		HeroID         int         `json:"hero_id"`
		PlayerSlot     interface{} `json:"player_slot"`
		ExtraTime      int         `json:"extra_time"`
		TotalTimeTaken int         `json:"total_time_taken"`
	} `json:"draft_timings"`
	Duration       int   `json:"duration"`
	Engine         int   `json:"engine"`
	FirstBloodTime int   `json:"first_blood_time"`
	GameMode       int   `json:"game_mode"`
	HumanPlayers   int   `json:"human_players"`
	Leagueid       int   `json:"leagueid"`
	LobbyType      int   `json:"lobby_type"`
	MatchSeqNum    int64 `json:"match_seq_num"`
	NegativeVotes  int   `json:"negative_votes"`
	Objectives     []struct {
		Time       int    `json:"time"`
		Type       string `json:"type"`
		Slot       int    `json:"slot,omitempty"`
		Key        int    `json:"key,string,omitempty"`
		PlayerSlot int    `json:"player_slot,omitempty"`
		Unit       string `json:"unit,omitempty"`
		Team       int    `json:"team,omitempty"`
	} `json:"objectives"`
	PicksBans []struct {
		IsPick  bool  `json:"is_pick"`
		HeroID  int   `json:"hero_id"`
		Team    int   `json:"team"`
		Order   int   `json:"order"`
		Ord     int   `json:"ord"`
		MatchID int64 `json:"match_id"`
	} `json:"picks_bans"`
	PositiveVotes  int         `json:"positive_votes"`
	RadiantGoldAdv []int       `json:"radiant_gold_adv"`
	RadiantScore   int         `json:"radiant_score"`
	RadiantWin     bool        `json:"radiant_win"`
	RadiantXpAdv   []int       `json:"radiant_xp_adv"`
	Skill          interface{} `json:"skill"`
	StartTime      int         `json:"start_time"`
	Teamfights     []struct {
		Start     int `json:"start"`
		End       int `json:"end"`
		LastDeath int `json:"last_death"`
		Deaths    int `json:"deaths"`
		Players   []struct {
			DeathsPos struct {
				Num122 struct {
					Num118 int `json:"118"`
				} `json:"122"`
			} `json:"deaths_pos"`
			AbilityUses struct {
				SkeletonKingHellfireBlast int `json:"skeleton_king_hellfire_blast"`
			} `json:"ability_uses"`
			AbilityTargets struct {
			} `json:"ability_targets"`
			ItemUses struct {
				Tpscroll  int `json:"tpscroll"`
				MagicWand int `json:"magic_wand"`
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
	TowerStatusDire    int `json:"tower_status_dire"`
	TowerStatusRadiant int `json:"tower_status_radiant"`
	Version            int `json:"version"`
	ReplaySalt         int `json:"replay_salt"`
	SeriesID           int `json:"series_id"`
	SeriesType         int `json:"series_type"`
	League             struct {
		Leagueid int    `json:"leagueid"`
		Ticket   string `json:"ticket"`
		Banner   string `json:"banner"`
		Tier     string `json:"tier"`
		Name     string `json:"name"`
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
		MatchID        int64 `json:"match_id"`
		PlayerSlot     int   `json:"player_slot"`
		AbilityTargets struct {
			SkeletonKingHellfireBlast struct {
				NpcDotaHeroEarthSpirit  int `json:"npc_dota_hero_earth_spirit"`
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroShadowDemon  int `json:"npc_dota_hero_shadow_demon"`
				NpcDotaHeroMorphling    int `json:"npc_dota_hero_morphling"`
				NpcDotaHeroBroodmother  int `json:"npc_dota_hero_broodmother"`
			} `json:"skeleton_king_hellfire_blast"`
		} `json:"ability_targets"`
		AbilityUpgradesArr []int `json:"ability_upgrades_arr"`
		AbilityUses        struct {
			SkeletonKingHellfireBlast int `json:"skeleton_king_hellfire_blast"`
			SkeletonKingVampiricAura  int `json:"skeleton_king_vampiric_aura"`
			SkeletonKingMortalStrike  int `json:"skeleton_king_mortal_strike"`
		} `json:"ability_uses"`
		AccountID int `json:"account_id"`
		Actions   struct {
			Num1  int `json:"1"`
			Num2  int `json:"2"`
			Num3  int `json:"3"`
			Num4  int `json:"4"`
			Num5  int `json:"5"`
			Num6  int `json:"6"`
			Num7  int `json:"7"`
			Num8  int `json:"8"`
			Num9  int `json:"9"`
			Num10 int `json:"10"`
			Num11 int `json:"11"`
			Num16 int `json:"16"`
			Num17 int `json:"17"`
			Num19 int `json:"19"`
			Num23 int `json:"23"`
			Num27 int `json:"27"`
			Num31 int `json:"31"`
			Num33 int `json:"33"`
		} `json:"actions"`
		AdditionalUnits interface{} `json:"additional_units"`
		Assists         int         `json:"assists"`
		Backpack0       int         `json:"backpack_0"`
		Backpack1       int         `json:"backpack_1"`
		Backpack2       int         `json:"backpack_2"`
		BuybackLog      []struct {
			Time       int    `json:"time"`
			Slot       int    `json:"slot"`
			Type       string `json:"type"`
			PlayerSlot int    `json:"player_slot"`
		} `json:"buyback_log"`
		CampsStacked  int `json:"camps_stacked"`
		CreepsStacked int `json:"creeps_stacked"`
		Damage        struct {
		} `json:"-"`
		DamageInflictor struct {
		} `json:"-"`
		DamageInflictorReceived struct {
		} `json:"-"`
		DamageTaken struct {
		} `json:"-"`
		Deaths            int   `json:"deaths"`
		Denies            int   `json:"denies"`
		DnT               []int `json:"dn_t"`
		FirstbloodClaimed int   `json:"firstblood_claimed"`
		Gold              int   `json:"gold"`
		GoldPerMin        int   `json:"gold_per_min"`
		GoldReasons       struct {
			Num0  int `json:"0"`
			Num1  int `json:"1"`
			Num2  int `json:"2"`
			Num6  int `json:"6"`
			Num11 int `json:"11"`
			Num12 int `json:"12"`
			Num13 int `json:"13"`
		} `json:"gold_reasons"`
		GoldSpent   int   `json:"gold_spent"`
		GoldT       []int `json:"gold_t"`
		HeroDamage  int   `json:"hero_damage"`
		HeroHealing int   `json:"hero_healing"`
		HeroHits    struct {
			SkeletonKingHellfireBlast int `json:"skeleton_king_hellfire_blast"`
			Null                      int `json:"null"`
			BladeMail                 int `json:"blade_mail"`
		} `json:"hero_hits"`
		HeroID   int `json:"hero_id"`
		Item0    int `json:"item_0"`
		Item1    int `json:"item_1"`
		Item2    int `json:"item_2"`
		Item3    int `json:"item_3"`
		Item4    int `json:"item_4"`
		Item5    int `json:"item_5"`
		ItemUses struct {
		} `json:"-"`
		KillStreaks struct {
		} `json:"kill_streaks"`
		Killed struct {
		} `json:"-"`
		KilledBy struct {
			NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
			NpcDotaHeroBroodmother  int `json:"npc_dota_hero_broodmother"`
			NpcDotaHeroMorphling    int `json:"npc_dota_hero_morphling"`
		} `json:"-"`
		Kills    int `json:"kills"`
		KillsLog []struct {
			Time int    `json:"time"`
			Key  string `json:"key"`
		} `json:"kills_log"`
		LanePos struct {
		} `json:"-"`
		LastHits     int   `json:"last_hits"`
		LeaverStatus int   `json:"leaver_status"`
		Level        int   `json:"level"`
		LhT          []int `json:"lh_t"`
		LifeState    struct {
			Num0 int `json:"0"`
			Num1 int `json:"1"`
			Num2 int `json:"2"`
		} `json:"life_state"`
		MaxHeroHit struct {
			Type       string      `json:"type"`
			Time       int         `json:"time"`
			Max        bool        `json:"max"`
			Inflictor  interface{} `json:"inflictor"`
			Unit       string      `json:"unit"`
			Key        string      `json:"key"`
			Value      int         `json:"value"`
			Slot       int         `json:"slot"`
			PlayerSlot int         `json:"player_slot"`
		} `json:"max_hero_hit"`
		MultiKills struct {
		} `json:"multi_kills"`
		Obs struct {
		} `json:"obs"`
		ObsLeftLog        []interface{} `json:"obs_left_log"`
		ObsLog            []interface{} `json:"obs_log"`
		ObsPlaced         int           `json:"obs_placed"`
		PartyID           int           `json:"party_id"`
		PartySize         int           `json:"party_size"`
		PerformanceOthers interface{}   `json:"performance_others"`
		PermanentBuffs    interface{}   `json:"permanent_buffs"`
		Pings             int           `json:"pings"`
		PredVict          bool          `json:"pred_vict"`
		Purchase          struct {
		} `json:"-"`
		PurchaseLog []struct {
			Time int    `json:"time"`
			Key  string `json:"key"`
		} `json:"-"`
		Randomed      bool        `json:"randomed"`
		Repicked      interface{} `json:"repicked"`
		RoshansKilled int         `json:"roshans_killed"`
		RunePickups   int         `json:"rune_pickups"`
		Runes         struct {
		} `json:"runes"`
		RunesLog []interface{} `json:"runes_log"`
		Sen      struct {
		} `json:"sen"`
		SenLeftLog             []interface{} `json:"sen_left_log"`
		SenLog                 []interface{} `json:"sen_log"`
		SenPlaced              int           `json:"sen_placed"`
		Stuns                  float64       `json:"stuns"`
		TeamfightParticipation float64       `json:"teamfight_participation"`
		Times                  []int         `json:"times"`
		TowerDamage            int           `json:"tower_damage"`
		TowersKilled           int           `json:"towers_killed"`
		XpPerMin               int           `json:"xp_per_min"`
		XpReasons              struct {
			Num0 int `json:"0"`
			Num1 int `json:"1"`
			Num2 int `json:"2"`
		} `json:"xp_reasons"`
		XpT               []int       `json:"xp_t"`
		Personaname       string      `json:"personaname"`
		Name              string      `json:"name"`
		LastLogin         interface{} `json:"last_login"`
		RadiantWin        bool        `json:"radiant_win"`
		StartTime         int         `json:"start_time"`
		Duration          int         `json:"duration"`
		Cluster           int         `json:"cluster"`
		LobbyType         int         `json:"lobby_type"`
		GameMode          int         `json:"game_mode"`
		Patch             int         `json:"patch"`
		Region            int         `json:"region"`
		IsRadiant         bool        `json:"isRadiant"`
		Win               int         `json:"win"`
		Lose              int         `json:"lose"`
		TotalGold         int         `json:"total_gold"`
		TotalXp           int         `json:"total_xp"`
		KillsPerMin       float64     `json:"kills_per_min"`
		Kda               int         `json:"kda"`
		Abandons          int         `json:"abandons"`
		NeutralKills      int         `json:"neutral_kills"`
		TowerKills        int         `json:"tower_kills"`
		CourierKills      int         `json:"courier_kills"`
		LaneKills         int         `json:"lane_kills"`
		HeroKills         int         `json:"hero_kills"`
		ObserverKills     int         `json:"observer_kills"`
		SentryKills       int         `json:"sentry_kills"`
		RoshanKills       int         `json:"roshan_kills"`
		NecronomiconKills int         `json:"necronomicon_kills"`
		AncientKills      int         `json:"ancient_kills"`
		BuybackCount      int         `json:"buyback_count"`
		ObserverUses      int         `json:"observer_uses"`
		SentryUses        int         `json:"sentry_uses"`
		LaneEfficiency    float64     `json:"lane_efficiency"`
		LaneEfficiencyPct int         `json:"lane_efficiency_pct"`
		Lane              int         `json:"lane"`
		LaneRole          int         `json:"lane_role"`
		IsRoaming         bool        `json:"is_roaming"`
		PurchaseTime      struct {
			Tango          int `json:"tango"`
			EnchantedMango int `json:"enchanted_mango"`
			Branches       int `json:"branches"`
			StoutShield    int `json:"stout_shield"`
			Tpscroll       int `json:"tpscroll"`
			MagicWand      int `json:"magic_wand"`
			MagicStick     int `json:"magic_stick"`
			QuellingBlade  int `json:"quelling_blade"`
			Gauntlets      int `json:"gauntlets"`
			RingOfRegen    int `json:"ring_of_regen"`
			SoulRing       int `json:"soul_ring"`
			Boots          int `json:"boots"`
			BladesOfAttack int `json:"blades_of_attack"`
			PhaseBoots     int `json:"phase_boots"`
			Claymore       int `json:"claymore"`
			ShadowAmulet   int `json:"shadow_amulet"`
			InvisSword     int `json:"invis_sword"`
			Broadsword     int `json:"broadsword"`
			Chainmail      int `json:"chainmail"`
			BladeMail      int `json:"blade_mail"`
			Robe           int `json:"robe"`
		} `json:"purchase_time"`
		FirstPurchaseTime struct {
			Tango          int `json:"tango"`
			EnchantedMango int `json:"enchanted_mango"`
			Branches       int `json:"branches"`
			StoutShield    int `json:"stout_shield"`
			Tpscroll       int `json:"tpscroll"`
			MagicWand      int `json:"magic_wand"`
			MagicStick     int `json:"magic_stick"`
			QuellingBlade  int `json:"quelling_blade"`
			Gauntlets      int `json:"gauntlets"`
			RingOfRegen    int `json:"ring_of_regen"`
			SoulRing       int `json:"soul_ring"`
			Boots          int `json:"boots"`
			BladesOfAttack int `json:"blades_of_attack"`
			PhaseBoots     int `json:"phase_boots"`
			Claymore       int `json:"claymore"`
			ShadowAmulet   int `json:"shadow_amulet"`
			InvisSword     int `json:"invis_sword"`
			Broadsword     int `json:"broadsword"`
			Chainmail      int `json:"chainmail"`
			BladeMail      int `json:"blade_mail"`
			Robe           int `json:"robe"`
		} `json:"first_purchase_time"`
		ItemWin struct {
			Tango          int `json:"tango"`
			EnchantedMango int `json:"enchanted_mango"`
			Branches       int `json:"branches"`
			StoutShield    int `json:"stout_shield"`
			Tpscroll       int `json:"tpscroll"`
			MagicWand      int `json:"magic_wand"`
			MagicStick     int `json:"magic_stick"`
			QuellingBlade  int `json:"quelling_blade"`
			Gauntlets      int `json:"gauntlets"`
			RingOfRegen    int `json:"ring_of_regen"`
			SoulRing       int `json:"soul_ring"`
			Boots          int `json:"boots"`
			BladesOfAttack int `json:"blades_of_attack"`
			PhaseBoots     int `json:"phase_boots"`
			Claymore       int `json:"claymore"`
			ShadowAmulet   int `json:"shadow_amulet"`
			InvisSword     int `json:"invis_sword"`
			Broadsword     int `json:"broadsword"`
			Chainmail      int `json:"chainmail"`
			BladeMail      int `json:"blade_mail"`
			Robe           int `json:"robe"`
		} `json:"item_win"`
		ItemUsage struct {
			Tango          int `json:"tango"`
			EnchantedMango int `json:"enchanted_mango"`
			Branches       int `json:"branches"`
			StoutShield    int `json:"stout_shield"`
			Tpscroll       int `json:"tpscroll"`
			MagicWand      int `json:"magic_wand"`
			MagicStick     int `json:"magic_stick"`
			QuellingBlade  int `json:"quelling_blade"`
			Gauntlets      int `json:"gauntlets"`
			RingOfRegen    int `json:"ring_of_regen"`
			SoulRing       int `json:"soul_ring"`
			Boots          int `json:"boots"`
			BladesOfAttack int `json:"blades_of_attack"`
			PhaseBoots     int `json:"phase_boots"`
			Claymore       int `json:"claymore"`
			ShadowAmulet   int `json:"shadow_amulet"`
			InvisSword     int `json:"invis_sword"`
			Broadsword     int `json:"broadsword"`
			Chainmail      int `json:"chainmail"`
			BladeMail      int `json:"blade_mail"`
			Robe           int `json:"robe"`
		} `json:"item_usage"`
		PurchaseTpscroll int           `json:"purchase_tpscroll"`
		ActionsPerMin    int           `json:"actions_per_min"`
		LifeStateDead    int           `json:"life_state_dead"`
		RankTier         int           `json:"rank_tier"`
		Cosmetics        []interface{} `json:"cosmetics"`
		Benchmarks       struct {
			GoldPerMin struct {
				Raw int     `json:"raw"`
				Pct float64 `json:"pct"`
			} `json:"gold_per_min"`
			XpPerMin struct {
				Raw int     `json:"raw"`
				Pct float64 `json:"pct"`
			} `json:"xp_per_min"`
			KillsPerMin struct {
				Raw float64 `json:"raw"`
				Pct float64 `json:"pct"`
			} `json:"kills_per_min"`
			LastHitsPerMin struct {
				Raw float64 `json:"raw"`
				Pct float64 `json:"pct"`
			} `json:"last_hits_per_min"`
			HeroDamagePerMin struct {
				Raw float64 `json:"raw"`
				Pct float64 `json:"pct"`
			} `json:"hero_damage_per_min"`
			HeroHealingPerMin struct {
				Raw float64 `json:"raw"`
				Pct float64 `json:"pct"`
			} `json:"hero_healing_per_min"`
			TowerDamage struct {
				Raw int     `json:"raw"`
				Pct float64 `json:"pct"`
			} `json:"tower_damage"`
		} `json:"benchmarks"`
		PurchaseWardObserver int `json:"purchase_ward_observer,omitempty"`
		PurchaseWardSentry   int `json:"purchase_ward_sentry,omitempty"`
	} `json:"players"`
	Patch         int `json:"patch"`
	Region        int `json:"region"`
	AllWordCounts struct {
	} `json:"-"`
	MyWordCounts struct {
	} `json:"-"`
	Comeback  int    `json:"comeback"`
	Stomp     int    `json:"stomp"`
	ReplayURL string `json:"replay_url"`
}
