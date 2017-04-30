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
		Unit       string `json:"unit"`
		Key        string `json:"key"`
		Slot       int    `json:"slot"`
		PlayerSlot int    `json:"player_slot"`
	} `json:"-"`
	Cluster   int `json:"cluster"`
	Cosmetics struct {
		Num4317  int `json:"4317"`
		Num4584  int `json:"4584"`
		Num4627  int `json:"4627"`
		Num4634  int `json:"4634"`
		Num4639  int `json:"4639"`
		Num4866  int `json:"4866"`
		Num4905  int `json:"4905"`
		Num5345  int `json:"5345"`
		Num5771  int `json:"5771"`
		Num5970  int `json:"5970"`
		Num6273  int `json:"6273"`
		Num6275  int `json:"6275"`
		Num6431  int `json:"6431"`
		Num6947  int `json:"6947"`
		Num6996  int `json:"6996"`
		Num7247  int `json:"7247"`
		Num8021  int `json:"8021"`
		Num8061  int `json:"8061"`
		Num8072  int `json:"8072"`
		Num8259  int `json:"8259"`
		Num8722  int `json:"8722"`
		Num8723  int `json:"8723"`
		Num8724  int `json:"8724"`
		Num8725  int `json:"8725"`
		Num8829  int `json:"8829"`
		Num8831  int `json:"8831"`
		Num8915  int `json:"8915"`
		Num9020  int `json:"9020"`
		Num9028  int `json:"9028"`
		Num9050  int `json:"9050"`
		Num9062  int `json:"9062"`
		Num9063  int `json:"9063"`
		Num9064  int `json:"9064"`
		Num9065  int `json:"9065"`
		Num9118  int `json:"9118"`
		Num9320  int `json:"9320"`
		Num9321  int `json:"9321"`
		Num9322  int `json:"9322"`
		Num9323  int `json:"9323"`
		Num9800  int `json:"9800"`
		Num10757 int `json:"10757"`
		Num11092 int `json:"11092"`
		Num11384 int `json:"11384"`
		Num11965 int `json:"11965"`
		Num11968 int `json:"11968"`
		Num11993 int `json:"11993"`
		Num17239 int `json:"17239"`
	} `json:"-"`
	DireScore      int   `json:"dire_score"`
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
		Key        int    `json:"key,omitempty"`
		PlayerSlot int    `json:"player_slot,omitempty"`
		Team       int    `json:"team,omitempty"`
	} `json:"-"`
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
			} `json:"deaths_pos"`
			AbilityUses struct {
			} `json:"ability_uses"`
			ItemUses struct {
				Tpscroll int `json:"tpscroll"`
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
	RadiantTeam        struct {
		TeamID int    `json:"team_id"`
		Name   string `json:"name"`
		Tag    string `json:"tag"`
	} `json:"radiant_team"`
	DireTeam struct {
		TeamID int    `json:"team_id"`
		Name   string `json:"name"`
		Tag    string `json:"tag"`
	} `json:"dire_team"`
	League struct {
		Leagueid int    `json:"leagueid"`
		Ticket   string `json:"ticket"`
		Banner   string `json:"banner"`
		Tier     string `json:"tier"`
		Name     string `json:"name"`
	} `json:"league"`
	Players []struct {
		MatchID            int64 `json:"match_id"`
		PlayerSlot         int   `json:"player_slot"`
		AbilityUpgradesArr []int `json:"ability_upgrades_arr"`
		AbilityUses        struct {
			OmniknightPurification  int `json:"omniknight_purification"`
			OmniknightRepel         int `json:"omniknight_repel"`
			OmniknightGuardianAngel int `json:"omniknight_guardian_angel"`
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
			Num10 int `json:"10"`
			Num11 int `json:"11"`
			Num13 int `json:"13"`
			Num15 int `json:"15"`
			Num16 int `json:"16"`
			Num17 int `json:"17"`
			Num19 int `json:"19"`
			Num25 int `json:"25"`
			Num27 int `json:"27"`
			Num33 int `json:"33"`
		} `json:"-"`
		AdditionalUnits interface{}   `json:"additional_units"`
		Assists         int           `json:"assists"`
		Backpack0       int           `json:"backpack_0"`
		Backpack1       int           `json:"backpack_1"`
		Backpack2       int           `json:"backpack_2"`
		BuybackLog      []interface{} `json:"buyback_log"`
		CampsStacked    int           `json:"camps_stacked"`
		CreepsStacked   int           `json:"creeps_stacked"`
		Damage          struct {
			NpcDotaCreepBadguysMelee       int `json:"npc_dota_creep_badguys_melee"`
			NpcDotaCreepGoodguysRanged     int `json:"npc_dota_creep_goodguys_ranged"`
			NpcDotaCreepGoodguysMelee      int `json:"npc_dota_creep_goodguys_melee"`
			NpcDotaHeroDoomBringer         int `json:"npc_dota_hero_doom_bringer"`
			NpcDotaHeroMonkeyKing          int `json:"npc_dota_hero_monkey_king"`
			NpcDotaCreepBadguysRanged      int `json:"npc_dota_creep_badguys_ranged"`
			NpcDotaBadguysSiege            int `json:"npc_dota_badguys_siege"`
			NpcDotaHeroNevermore           int `json:"npc_dota_hero_nevermore"`
			NpcDotaNeutralDarkTrollWarlord int `json:"npc_dota_neutral_dark_troll_warlord"`
			NpcDotaNeutralDarkTroll        int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaBadguysTower1Bot        int `json:"npc_dota_badguys_tower1_bot"`
			NpcDotaNeutralCentaurOutrunner int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaNeutralCentaurKhan      int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaNeutralSatyrHellcaller  int `json:"npc_dota_neutral_satyr_hellcaller"`
			NpcDotaNeutralSatyrSoulstealer int `json:"npc_dota_neutral_satyr_soulstealer"`
			NpcDotaNeutralSatyrTrickster   int `json:"npc_dota_neutral_satyr_trickster"`
			NpcDotaNeutralMudGolem         int `json:"npc_dota_neutral_mud_golem"`
			NpcDotaNeutralOgreMagi         int `json:"npc_dota_neutral_ogre_magi"`
			NpcDotaNeutralOgreMauler       int `json:"npc_dota_neutral_ogre_mauler"`
			NpcDotaNeutralMudGolemSplit    int `json:"npc_dota_neutral_mud_golem_split"`
			NpcDotaHeroSpectre             int `json:"npc_dota_hero_spectre"`
			NpcDotaHeroRubick              int `json:"npc_dota_hero_rubick"`
			NpcDotaRoshan                  int `json:"npc_dota_roshan"`
			NpcDotaBadguysTower2Mid        int `json:"npc_dota_badguys_tower2_mid"`
			NpcDotaBadguysTower3Mid        int `json:"npc_dota_badguys_tower3_mid"`
			NpcDotaBadguysTower2Top        int `json:"npc_dota_badguys_tower2_top"`
			NpcDotaBadguysTower3Top        int `json:"npc_dota_badguys_tower3_top"`
			NpcDotaBadguysMeleeRaxTop      int `json:"npc_dota_badguys_melee_rax_top"`
		} `json:"-"`
		DamageInflictor struct {
			Null                   int `json:"null"`
			OrbOfVenom             int `json:"orb_of_venom"`
			OmniknightPurification int `json:"omniknight_purification"`
		} `json:"damage_inflictor"`
		DamageInflictorReceived struct {
			Null                     int `json:"null"`
			OrbOfVenom               int `json:"orb_of_venom"`
			DoomBringerInfernalBlade int `json:"doom_bringer_infernal_blade"`
			CentaurKhanWarStomp      int `json:"centaur_khan_war_stomp"`
			SatyrSoulstealerManaBurn int `json:"satyr_soulstealer_mana_burn"`
			DoomBringerScorchedEarth int `json:"doom_bringer_scorched_earth"`
			NevermoreShadowraze2     int `json:"nevermore_shadowraze2"`
			RubickFadeBolt           int `json:"rubick_fade_bolt"`
			SpectreDesolate          int `json:"spectre_desolate"`
			SpectreSpectralDagger    int `json:"spectre_spectral_dagger"`
			DoomBringerDoom          int `json:"doom_bringer_doom"`
			SpectreDispersion        int `json:"spectre_dispersion"`
			NevermoreRequiem         int `json:"nevermore_requiem"`
			Radiance                 int `json:"radiance"`
		} `json:"-"`
		DamageTaken struct {
			NpcDotaHeroDoomBringer         int `json:"npc_dota_hero_doom_bringer"`
			NpcDotaCreepBadguysRanged      int `json:"npc_dota_creep_badguys_ranged"`
			NpcDotaCreepBadguysMelee       int `json:"npc_dota_creep_badguys_melee"`
			NpcDotaHeroMonkeyKing          int `json:"npc_dota_hero_monkey_king"`
			NpcDotaBadguysTower1Bot        int `json:"npc_dota_badguys_tower1_bot"`
			NpcDotaHeroNevermore           int `json:"npc_dota_hero_nevermore"`
			NpcDotaHeroRubick              int `json:"npc_dota_hero_rubick"`
			NpcDotaHeroSpectre             int `json:"npc_dota_hero_spectre"`
			NpcDotaNeutralDarkTrollWarlord int `json:"npc_dota_neutral_dark_troll_warlord"`
			NpcDotaNeutralDarkTroll        int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaNeutralCentaurKhan      int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaNeutralCentaurOutrunner int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaNeutralSatyrSoulstealer int `json:"npc_dota_neutral_satyr_soulstealer"`
			NpcDotaNeutralSatyrHellcaller  int `json:"npc_dota_neutral_satyr_hellcaller"`
			NpcDotaNeutralSatyrTrickster   int `json:"npc_dota_neutral_satyr_trickster"`
			NpcDotaNeutralMudGolem         int `json:"npc_dota_neutral_mud_golem"`
			NpcDotaNeutralMudGolemSplit    int `json:"npc_dota_neutral_mud_golem_split"`
			NpcDotaNeutralOgreMagi         int `json:"npc_dota_neutral_ogre_magi"`
			NpcDotaNeutralOgreMauler       int `json:"npc_dota_neutral_ogre_mauler"`
			NpcDotaBadguysTower1Mid        int `json:"npc_dota_badguys_tower1_mid"`
			NpcDotaRoshan                  int `json:"npc_dota_roshan"`
			NpcDotaBadguysTower2Mid        int `json:"npc_dota_badguys_tower2_mid"`
			NpcDotaBadguysTower2Top        int `json:"npc_dota_badguys_tower2_top"`
			NpcDotaBadguysTower3Top        int `json:"npc_dota_badguys_tower3_top"`
		} `json:"-"`
		Deaths      int   `json:"deaths"`
		Denies      int   `json:"denies"`
		DnT         []int `json:"dn_t"`
		Gold        int   `json:"gold"`
		GoldPerMin  int   `json:"gold_per_min"`
		GoldReasons struct {
			Num0  int `json:"0"`
			Num1  int `json:"1"`
			Num6  int `json:"6"`
			Num11 int `json:"11"`
			Num12 int `json:"12"`
			Num13 int `json:"13"`
			Num14 int `json:"14"`
		} `json:"-"`
		GoldSpent   int   `json:"gold_spent"`
		GoldT       []int `json:"gold_t"`
		HeroDamage  int   `json:"hero_damage"`
		HeroHealing int   `json:"hero_healing"`
		HeroHits    struct {
			Null                   int `json:"null"`
			OrbOfVenom             int `json:"orb_of_venom"`
			OmniknightPurification int `json:"omniknight_purification"`
		} `json:"-"`
		HeroID   int `json:"hero_id"`
		Item0    int `json:"item_0"`
		Item1    int `json:"item_1"`
		Item2    int `json:"item_2"`
		Item3    int `json:"item_3"`
		Item4    int `json:"item_4"`
		Item5    int `json:"item_5"`
		ItemUses struct {
			Tango           int `json:"tango"`
			ArcaneBoots     int `json:"arcane_boots"`
			Tpscroll        int `json:"tpscroll"`
			MagicStick      int `json:"magic_stick"`
			HandOfMidas     int `json:"hand_of_midas"`
			Dust            int `json:"dust"`
			TangoSingle     int `json:"tango_single"`
			Clarity         int `json:"clarity"`
			Mekansm         int `json:"mekansm"`
			GuardianGreaves int `json:"guardian_greaves"`
		} `json:"-"`
		KillStreaks struct {
		} `json:"kill_streaks"`
		Killed struct {
			NpcDotaCreepBadguysMelee       int `json:"npc_dota_creep_badguys_melee"`
			NpcDotaCreepGoodguysRanged     int `json:"npc_dota_creep_goodguys_ranged"`
			NpcDotaCreepGoodguysMelee      int `json:"npc_dota_creep_goodguys_melee"`
			NpcDotaCreepBadguysRanged      int `json:"npc_dota_creep_badguys_ranged"`
			NpcDotaHeroMonkeyKing          int `json:"npc_dota_hero_monkey_king"`
			NpcDotaBadguysSiege            int `json:"npc_dota_badguys_siege"`
			NpcDotaNeutralDarkTrollWarlord int `json:"npc_dota_neutral_dark_troll_warlord"`
			NpcDotaNeutralDarkTroll        int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaNeutralCentaurOutrunner int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaNeutralCentaurKhan      int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaNeutralSatyrTrickster   int `json:"npc_dota_neutral_satyr_trickster"`
			NpcDotaNeutralSatyrHellcaller  int `json:"npc_dota_neutral_satyr_hellcaller"`
			NpcDotaNeutralSatyrSoulstealer int `json:"npc_dota_neutral_satyr_soulstealer"`
			NpcDotaNeutralMudGolem         int `json:"npc_dota_neutral_mud_golem"`
			NpcDotaNeutralMudGolemSplit    int `json:"npc_dota_neutral_mud_golem_split"`
			NpcDotaNeutralOgreMauler       int `json:"npc_dota_neutral_ogre_mauler"`
			NpcDotaNeutralOgreMagi         int `json:"npc_dota_neutral_ogre_magi"`
			NpcDotaHeroRubick              int `json:"npc_dota_hero_rubick"`
			NpcDotaRoshan                  int `json:"npc_dota_roshan"`
		} `json:"-"`
		KilledBy struct {
			NpcDotaHeroSpectre     int `json:"npc_dota_hero_spectre"`
			NpcDotaHeroDoomBringer int `json:"npc_dota_hero_doom_bringer"`
			NpcDotaHeroNevermore   int `json:"npc_dota_hero_nevermore"`
		} `json:"-"`
		Kills    int `json:"kills"`
		KillsLog []struct {
			Time int    `json:"time"`
			Key  string `json:"key"`
		} `json:"-"`
		LanePos struct {
			Num72 struct {
				Num78 int `json:"78"`
				Num80 int `json:"80"`
			} `json:"72"`
			Num74 struct {
				Num74 int `json:"74"`
				Num76 int `json:"76"`
				Num78 int `json:"78"`
				Num80 int `json:"80"`
			} `json:"74"`
			Num76 struct {
				Num76 int `json:"76"`
				Num78 int `json:"78"`
				Num80 int `json:"80"`
			} `json:"76"`
			Num78 struct {
				Num78 int `json:"78"`
				Num80 int `json:"80"`
			} `json:"78"`
			Num80 struct {
				Num82 int `json:"82"`
			} `json:"80"`
			Num82 struct {
				Num78 int `json:"78"`
				Num82 int `json:"82"`
			} `json:"82"`
			Num84 struct {
				Num78 int `json:"78"`
				Num82 int `json:"82"`
			} `json:"84"`
			Num88 struct {
				Num78 int `json:"78"`
				Num80 int `json:"80"`
				Num82 int `json:"82"`
			} `json:"88"`
			Num90 struct {
				Num80 int `json:"80"`
				Num82 int `json:"82"`
			} `json:"90"`
			Num92 struct {
				Num82 int `json:"82"`
			} `json:"92"`
			Num94 struct {
				Num80 int `json:"80"`
				Num82 int `json:"82"`
			} `json:"94"`
			Num96 struct {
				Num80 int `json:"80"`
				Num82 int `json:"82"`
			} `json:"96"`
			Num98 struct {
				Num80 int `json:"80"`
			} `json:"98"`
			Num100 struct {
				Num82 int `json:"82"`
			} `json:"100"`
			Num102 struct {
				Num80 int `json:"80"`
				Num82 int `json:"82"`
			} `json:"102"`
			Num104 struct {
				Num80 int `json:"80"`
				Num82 int `json:"82"`
				Num96 int `json:"96"`
			} `json:"104"`
			Num106 struct {
				Num80 int `json:"80"`
				Num82 int `json:"82"`
				Num90 int `json:"90"`
				Num94 int `json:"94"`
				Num96 int `json:"96"`
			} `json:"106"`
			Num108 struct {
				Num82 int `json:"82"`
				Num86 int `json:"86"`
				Num88 int `json:"88"`
				Num98 int `json:"98"`
			} `json:"108"`
			Num110 struct {
				Num80 int `json:"80"`
				Num82 int `json:"82"`
				Num84 int `json:"84"`
				Num98 int `json:"98"`
			} `json:"110"`
			Num112 struct {
				Num80  int `json:"80"`
				Num100 int `json:"100"`
			} `json:"112"`
			Num114 struct {
				Num80  int `json:"80"`
				Num102 int `json:"102"`
			} `json:"114"`
			Num116 struct {
				Num80  int `json:"80"`
				Num104 int `json:"104"`
			} `json:"116"`
			Num118 struct {
				Num80  int `json:"80"`
				Num104 int `json:"104"`
			} `json:"118"`
			Num120 struct {
				Num80  int `json:"80"`
				Num106 int `json:"106"`
			} `json:"120"`
			Num122 struct {
				Num106 int `json:"106"`
			} `json:"122"`
			Num124 struct {
				Num80  int `json:"80"`
				Num106 int `json:"106"`
			} `json:"124"`
			Num126 struct {
				Num80  int `json:"80"`
				Num90  int `json:"90"`
				Num92  int `json:"92"`
				Num94  int `json:"94"`
				Num106 int `json:"106"`
			} `json:"126"`
			Num128 struct {
				Num80  int `json:"80"`
				Num94  int `json:"94"`
				Num108 int `json:"108"`
			} `json:"128"`
			Num130 struct {
				Num94  int `json:"94"`
				Num96  int `json:"96"`
				Num108 int `json:"108"`
			} `json:"130"`
			Num132 struct {
				Num78  int `json:"78"`
				Num94  int `json:"94"`
				Num96  int `json:"96"`
				Num108 int `json:"108"`
			} `json:"132"`
			Num134 struct {
				Num76  int `json:"76"`
				Num78  int `json:"78"`
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num102 int `json:"102"`
				Num104 int `json:"104"`
				Num106 int `json:"106"`
			} `json:"134"`
			Num136 struct {
				Num76 int `json:"76"`
				Num78 int `json:"78"`
				Num80 int `json:"80"`
				Num88 int `json:"88"`
				Num92 int `json:"92"`
				Num94 int `json:"94"`
				Num98 int `json:"98"`
			} `json:"136"`
			Num138 struct {
				Num76 int `json:"76"`
				Num82 int `json:"82"`
				Num84 int `json:"84"`
				Num86 int `json:"86"`
				Num94 int `json:"94"`
			} `json:"138"`
			Num140 struct {
				Num76 int `json:"76"`
				Num78 int `json:"78"`
			} `json:"140"`
			Num142 struct {
				Num78 int `json:"78"`
				Num80 int `json:"80"`
			} `json:"142"`
			Num144 struct {
				Num78 int `json:"78"`
			} `json:"144"`
			Num146 struct {
				Num78 int `json:"78"`
				Num80 int `json:"80"`
			} `json:"146"`
			Num148 struct {
				Num78 int `json:"78"`
				Num80 int `json:"80"`
				Num98 int `json:"98"`
			} `json:"148"`
			Num150 struct {
				Num78 int `json:"78"`
				Num80 int `json:"80"`
				Num96 int `json:"96"`
			} `json:"150"`
			Num152 struct {
				Num78 int `json:"78"`
				Num98 int `json:"98"`
			} `json:"152"`
			Num154 struct {
				Num80 int `json:"80"`
				Num96 int `json:"96"`
				Num98 int `json:"98"`
			} `json:"154"`
			Num156 struct {
				Num80 int `json:"80"`
				Num96 int `json:"96"`
			} `json:"156"`
			Num158 struct {
				Num80 int `json:"80"`
				Num82 int `json:"82"`
				Num92 int `json:"92"`
				Num94 int `json:"94"`
			} `json:"158"`
			Num160 struct {
				Num80 int `json:"80"`
				Num90 int `json:"90"`
				Num92 int `json:"92"`
				Num94 int `json:"94"`
				Num96 int `json:"96"`
			} `json:"160"`
			Num162 struct {
				Num80 int `json:"80"`
				Num82 int `json:"82"`
				Num88 int `json:"88"`
				Num90 int `json:"90"`
				Num96 int `json:"96"`
			} `json:"162"`
			Num164 struct {
				Num78 int `json:"78"`
				Num80 int `json:"80"`
				Num82 int `json:"82"`
				Num88 int `json:"88"`
				Num92 int `json:"92"`
				Num94 int `json:"94"`
				Num96 int `json:"96"`
			} `json:"164"`
			Num166 struct {
				Num74 int `json:"74"`
				Num78 int `json:"78"`
				Num80 int `json:"80"`
				Num82 int `json:"82"`
				Num84 int `json:"84"`
				Num88 int `json:"88"`
				Num96 int `json:"96"`
				Num98 int `json:"98"`
			} `json:"166"`
			Num168 struct {
				Num76  int `json:"76"`
				Num78  int `json:"78"`
				Num80  int `json:"80"`
				Num82  int `json:"82"`
				Num84  int `json:"84"`
				Num86  int `json:"86"`
				Num88  int `json:"88"`
				Num90  int `json:"90"`
				Num94  int `json:"94"`
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
			} `json:"168"`
			Num170 struct {
				Num78  int `json:"78"`
				Num80  int `json:"80"`
				Num82  int `json:"82"`
				Num84  int `json:"84"`
				Num86  int `json:"86"`
				Num88  int `json:"88"`
				Num90  int `json:"90"`
				Num94  int `json:"94"`
				Num96  int `json:"96"`
				Num100 int `json:"100"`
			} `json:"170"`
			Num172 struct {
				Num78  int `json:"78"`
				Num80  int `json:"80"`
				Num82  int `json:"82"`
				Num84  int `json:"84"`
				Num86  int `json:"86"`
				Num88  int `json:"88"`
				Num90  int `json:"90"`
				Num92  int `json:"92"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num102 int `json:"102"`
			} `json:"172"`
			Num174 struct {
				Num82  int `json:"82"`
				Num84  int `json:"84"`
				Num86  int `json:"86"`
				Num88  int `json:"88"`
				Num90  int `json:"90"`
				Num92  int `json:"92"`
				Num94  int `json:"94"`
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num102 int `json:"102"`
				Num104 int `json:"104"`
				Num106 int `json:"106"`
				Num108 int `json:"108"`
				Num110 int `json:"110"`
				Num112 int `json:"112"`
			} `json:"174"`
			Num176 struct {
				Num82  int `json:"82"`
				Num84  int `json:"84"`
				Num86  int `json:"86"`
				Num88  int `json:"88"`
				Num90  int `json:"90"`
				Num92  int `json:"92"`
				Num94  int `json:"94"`
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num102 int `json:"102"`
				Num104 int `json:"104"`
				Num106 int `json:"106"`
				Num108 int `json:"108"`
				Num110 int `json:"110"`
			} `json:"176"`
			Num178 struct {
				Num84  int `json:"84"`
				Num86  int `json:"86"`
				Num88  int `json:"88"`
				Num92  int `json:"92"`
				Num94  int `json:"94"`
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num102 int `json:"102"`
				Num104 int `json:"104"`
				Num106 int `json:"106"`
				Num108 int `json:"108"`
				Num110 int `json:"110"`
			} `json:"178"`
			Num180 struct {
				Num86 int `json:"86"`
				Num88 int `json:"88"`
				Num90 int `json:"90"`
				Num92 int `json:"92"`
				Num94 int `json:"94"`
				Num96 int `json:"96"`
				Num98 int `json:"98"`
			} `json:"180"`
			Num182 struct {
				Num86 int `json:"86"`
				Num88 int `json:"88"`
				Num90 int `json:"90"`
				Num92 int `json:"92"`
				Num94 int `json:"94"`
				Num96 int `json:"96"`
				Num98 int `json:"98"`
			} `json:"182"`
			Num184 struct {
				Num98 int `json:"98"`
			} `json:"184"`
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
			Type       string `json:"type"`
			Time       int    `json:"time"`
			Max        bool   `json:"max"`
			Inflictor  string `json:"inflictor"`
			Unit       string `json:"unit"`
			Key        string `json:"key"`
			Value      int    `json:"value"`
			Slot       int    `json:"slot"`
			PlayerSlot int    `json:"player_slot"`
		} `json:"max_hero_hit"`
		MultiKills struct {
		} `json:"multi_kills"`
		Obs struct {
		} `json:"obs"`
		ObsLeftLog     []interface{} `json:"obs_left_log"`
		ObsLog         []interface{} `json:"obs_log"`
		ObsPlaced      int           `json:"obs_placed"`
		PartyID        int           `json:"party_id"`
		PermanentBuffs interface{}   `json:"permanent_buffs"`
		Pings          int           `json:"pings"`
		PredVict       bool          `json:"pred_vict"`
		Purchase       struct {
			OrbOfVenom            int `json:"orb_of_venom"`
			Boots                 int `json:"boots"`
			ArcaneBoots           int `json:"arcane_boots"`
			EnergyBooster         int `json:"energy_booster"`
			Tpscroll              int `json:"tpscroll"`
			InfusedRaindrop       int `json:"infused_raindrop"`
			MagicStick            int `json:"magic_stick"`
			Gloves                int `json:"gloves"`
			RecipeHandOfMidas     int `json:"recipe_hand_of_midas"`
			HandOfMidas           int `json:"hand_of_midas"`
			RingOfRegen           int `json:"ring_of_regen"`
			Branches              int `json:"branches"`
			Headdress             int `json:"headdress"`
			RecipeHeaddress       int `json:"recipe_headdress"`
			Dust                  int `json:"dust"`
			Chainmail             int `json:"chainmail"`
			Buckler               int `json:"buckler"`
			RecipeBuckler         int `json:"recipe_buckler"`
			RecipeMekansm         int `json:"recipe_mekansm"`
			Clarity               int `json:"clarity"`
			Mekansm               int `json:"mekansm"`
			Gem                   int `json:"gem"`
			RecipeGuardianGreaves int `json:"recipe_guardian_greaves"`
			GuardianGreaves       int `json:"guardian_greaves"`
		} `json:"-"`
		PurchaseLog []struct {
			Time int    `json:"time"`
			Key  string `json:"key"`
		} `json:"purchase_log"`
		Randomed    bool        `json:"randomed"`
		Repicked    interface{} `json:"repicked"`
		RunePickups int         `json:"rune_pickups"`
		Runes       struct {
			Num5 int `json:"5"`
		} `json:"runes"`
		RunesLog []struct {
			Time int `json:"time"`
			Key  int `json:"key"`
		} `json:"runes_log"`
		Sen struct {
		} `json:"sen"`
		SenLeftLog  []interface{} `json:"sen_left_log"`
		SenLog      []interface{} `json:"sen_log"`
		SenPlaced   int           `json:"sen_placed"`
		Stuns       int           `json:"-"`
		Times       []int         `json:"times"`
		TowerDamage int           `json:"tower_damage"`
		XpPerMin    int           `json:"xp_per_min"`
		XpReasons   struct {
			Num0 int `json:"0"`
			Num1 int `json:"1"`
			Num2 int `json:"2"`
			Num3 int `json:"3"`
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
			OrbOfVenom      int `json:"orb_of_venom"`
			Boots           int `json:"boots"`
			ArcaneBoots     int `json:"arcane_boots"`
			EnergyBooster   int `json:"energy_booster"`
			Tpscroll        int `json:"tpscroll"`
			InfusedRaindrop int `json:"infused_raindrop"`
			MagicStick      int `json:"magic_stick"`
			Gloves          int `json:"gloves"`
			HandOfMidas     int `json:"hand_of_midas"`
			RingOfRegen     int `json:"ring_of_regen"`
			Branches        int `json:"branches"`
			Headdress       int `json:"headdress"`
			Dust            int `json:"dust"`
			Chainmail       int `json:"chainmail"`
			Buckler         int `json:"buckler"`
			Clarity         int `json:"clarity"`
			Mekansm         int `json:"mekansm"`
			Gem             int `json:"gem"`
			GuardianGreaves int `json:"guardian_greaves"`
		} `json:"-"`
		FirstPurchaseTime struct {
			OrbOfVenom      int `json:"orb_of_venom"`
			Boots           int `json:"boots"`
			ArcaneBoots     int `json:"arcane_boots"`
			EnergyBooster   int `json:"energy_booster"`
			Tpscroll        int `json:"tpscroll"`
			InfusedRaindrop int `json:"infused_raindrop"`
			MagicStick      int `json:"magic_stick"`
			Gloves          int `json:"gloves"`
			HandOfMidas     int `json:"hand_of_midas"`
			RingOfRegen     int `json:"ring_of_regen"`
			Branches        int `json:"branches"`
			Headdress       int `json:"headdress"`
			Dust            int `json:"dust"`
			Chainmail       int `json:"chainmail"`
			Buckler         int `json:"buckler"`
			Clarity         int `json:"clarity"`
			Mekansm         int `json:"mekansm"`
			Gem             int `json:"gem"`
			GuardianGreaves int `json:"guardian_greaves"`
		} `json:"-"`
		ItemWin struct {
			OrbOfVenom      int `json:"orb_of_venom"`
			Boots           int `json:"boots"`
			ArcaneBoots     int `json:"arcane_boots"`
			EnergyBooster   int `json:"energy_booster"`
			Tpscroll        int `json:"tpscroll"`
			InfusedRaindrop int `json:"infused_raindrop"`
			MagicStick      int `json:"magic_stick"`
			Gloves          int `json:"gloves"`
			HandOfMidas     int `json:"hand_of_midas"`
			RingOfRegen     int `json:"ring_of_regen"`
			Branches        int `json:"branches"`
			Headdress       int `json:"headdress"`
			Dust            int `json:"dust"`
			Chainmail       int `json:"chainmail"`
			Buckler         int `json:"buckler"`
			Clarity         int `json:"clarity"`
			Mekansm         int `json:"mekansm"`
			Gem             int `json:"gem"`
			GuardianGreaves int `json:"guardian_greaves"`
		} `json:"-"`
		ItemUsage struct {
			OrbOfVenom      int `json:"orb_of_venom"`
			Boots           int `json:"boots"`
			ArcaneBoots     int `json:"arcane_boots"`
			EnergyBooster   int `json:"energy_booster"`
			Tpscroll        int `json:"tpscroll"`
			InfusedRaindrop int `json:"infused_raindrop"`
			MagicStick      int `json:"magic_stick"`
			Gloves          int `json:"gloves"`
			HandOfMidas     int `json:"hand_of_midas"`
			RingOfRegen     int `json:"ring_of_regen"`
			Branches        int `json:"branches"`
			Headdress       int `json:"headdress"`
			Dust            int `json:"dust"`
			Chainmail       int `json:"chainmail"`
			Buckler         int `json:"buckler"`
			Clarity         int `json:"clarity"`
			Mekansm         int `json:"mekansm"`
			Gem             int `json:"gem"`
			GuardianGreaves int `json:"guardian_greaves"`
		} `json:"-"`
		PurchaseTpscroll    int         `json:"purchase_tpscroll"`
		PurchaseGem         int         `json:"purchase_gem,omitempty"`
		ActionsPerMin       int         `json:"actions_per_min"`
		LifeStateDead       int         `json:"life_state_dead"`
		SoloCompetitiveRank interface{} `json:"solo_competitive_rank"`
		Cosmetics           []struct {
			ItemID          int         `json:"item_id"`
			Name            string      `json:"name"`
			Prefab          string      `json:"prefab"`
			CreationDate    interface{} `json:"creation_date"`
			ImageInventory  string      `json:"image_inventory"`
			ImagePath       string      `json:"image_path"`
			ItemDescription string      `json:"item_description"`
			ItemName        string      `json:"item_name"`
			ItemRarity      interface{} `json:"item_rarity"`
			ItemTypeName    string      `json:"item_type_name"`
			UsedByHeroes    string      `json:"used_by_heroes"`
		} `json:"-"`
		Benchmarks struct {
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
		} `json:"-"`
		PurchaseWardObserver int `json:"purchase_ward_observer,omitempty"`
		PurchaseWardSentry   int `json:"purchase_ward_sentry,omitempty"`
	} `json:"players"`
	Patch         int `json:"patch"`
	Region        int `json:"region"`
	AllWordCounts struct {
		Kale int `json:"kale"`
		Go   int `json:"go"`
		G    int `json:"g"`
		Gg   int `json:"gg"`
	} `json:"-"`
	MyWordCounts struct {
	} `json:"-"`
	Throw     int    `json:"throw"`
	Loss      int    `json:"loss"`
	ReplayURL string `json:"replay_url"`
}
