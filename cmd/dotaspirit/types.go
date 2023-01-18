package main

import "time"

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
	MatchID               int64  `json:"match_id"`
	RadiantName           string `json:"radiant_name"`
	DireName              string `json:"dire_name"`
	BarracksStatusDire    int    `json:"barracks_status_dire"`
	BarracksStatusRadiant int    `json:"barracks_status_radiant"`
	Chat                  []struct {
		Time       int    `json:"time"`
		Type       string `json:"type"`
		Key        string `json:"key"`
		Slot       int    `json:"slot"`
		PlayerSlot int    `json:"player_slot"`
	} `json:"chat"`
	Cluster   int `json:"cluster"`
	Cosmetics struct {
		Num647   int `json:"647"`
		Num5995  int `json:"5995"`
		Num6694  int `json:"6694"`
		Num6934  int `json:"6934"`
		Num6935  int `json:"6935"`
		Num6936  int `json:"6936"`
		Num6937  int `json:"6937"`
		Num6938  int `json:"6938"`
		Num6941  int `json:"6941"`
		Num7212  int `json:"7212"`
		Num7459  int `json:"7459"`
		Num7818  int `json:"7818"`
		Num8073  int `json:"8073"`
		Num8737  int `json:"8737"`
		Num9318  int `json:"9318"`
		Num9667  int `json:"9667"`
		Num9732  int `json:"9732"`
		Num9734  int `json:"9734"`
		Num9736  int `json:"9736"`
		Num9772  int `json:"9772"`
		Num9996  int `json:"9996"`
		Num10070 int `json:"10070"`
		Num10091 int `json:"10091"`
		Num10147 int `json:"10147"`
		Num10193 int `json:"10193"`
		Num10433 int `json:"10433"`
		Num11903 int `json:"11903"`
		Num11998 int `json:"11998"`
		Num12322 int `json:"12322"`
		Num12465 int `json:"12465"`
		Num12540 int `json:"12540"`
		Num12687 int `json:"12687"`
		Num12688 int `json:"12688"`
		Num12690 int `json:"12690"`
		Num12931 int `json:"12931"`
		Num13248 int `json:"13248"`
		Num13249 int `json:"13249"`
		Num13250 int `json:"13250"`
		Num13251 int `json:"13251"`
		Num13252 int `json:"13252"`
		Num13539 int `json:"13539"`
		Num13767 int `json:"13767"`
		Num13768 int `json:"13768"`
		Num13769 int `json:"13769"`
		Num13772 int `json:"13772"`
		Num14296 int `json:"14296"`
		Num17257 int `json:"17257"`
		Num17661 int `json:"17661"`
		Num17844 int `json:"17844"`
		Num17967 int `json:"17967"`
		Num17968 int `json:"17968"`
		Num17969 int `json:"17969"`
		Num17998 int `json:"17998"`
		Num18000 int `json:"18000"`
		Num18001 int `json:"18001"`
		Num18110 int `json:"18110"`
		Num18408 int `json:"18408"`
		Num18492 int `json:"18492"`
		Num18494 int `json:"18494"`
		Num18567 int `json:"18567"`
		Num18850 int `json:"18850"`
		Num19004 int `json:"19004"`
		Num19266 int `json:"19266"`
		Num22266 int `json:"22266"`
		Num22293 int `json:"22293"`
		Num22317 int `json:"22317"`
		Num22319 int `json:"22319"`
		Num23237 int `json:"23237"`
		Num23836 int `json:"23836"`
	} `json:"cosmetics"`
	DireScore    int `json:"dire_score"`
	DireTeamID   int `json:"dire_team_id"`
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
	// Objectives     []struct {
	// 	Time       int    `json:"time"`
	// 	Type       string `json:"type"`
	// 	Slot       int    `json:"slot,omitempty"`
	// 	Key        int    `json:"key,omitempty"`
	// 	PlayerSlot int    `json:"player_slot,omitempty"`
	// 	Team       int    `json:"team,omitempty"`
	// 	Unit       string `json:"unit,omitempty"`
	// } `json:"objectives"`
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
	RadiantTeamID  int         `json:"radiant_team_id"`
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
			} `json:"deaths_pos,omitempty"`
			AbilityUses struct {
				PrimalBeastOnslaught        int `json:"primal_beast_onslaught"`
				PrimalBeastOnslaughtRelease int `json:"primal_beast_onslaught_release"`
				PrimalBeastTrample          int `json:"primal_beast_trample"`
				PrimalBeastPulverize        int `json:"primal_beast_pulverize"`
				PrimalBeastUproar           int `json:"primal_beast_uproar"`
				PlusHighFive                int `json:"plus_high_five"`
			} `json:"ability_uses,omitempty"`
			AbilityTargets struct {
			} `json:"ability_targets"`
			ItemUses struct {
				ArcaneBoots   int `json:"arcane_boots"`
				Clarity       int `json:"clarity"`
				WardDispenser int `json:"ward_dispenser"`
			} `json:"item_uses,omitempty"`
			Killed struct {
			} `json:"killed,omitempty"`
			Deaths       int `json:"deaths"`
			Buybacks     int `json:"buybacks"`
			Damage       int `json:"damage"`
			Healing      int `json:"healing"`
			GoldDelta    int `json:"gold_delta"`
			XpDelta      int `json:"xp_delta"`
			XpStart      int `json:"xp_start"`
			XpEnd        int `json:"xp_end"`
			AbilityUses0 struct {
				PuckWaningRift    int `json:"puck_waning_rift"`
				PuckIllusoryOrb   int `json:"puck_illusory_orb"`
				PuckEtherealJaunt int `json:"puck_ethereal_jaunt"`
				PuckDreamCoil     int `json:"puck_dream_coil"`
			} `json:"ability_uses,omitempty"`
			ItemUses0 struct {
				Bottle      int `json:"bottle"`
				PowerTreads int `json:"power_treads"`
				Tpscroll    int `json:"tpscroll"`
			} `json:"item_uses,omitempty"`
			Killed0 struct {
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
			} `json:"killed,omitempty"`
			AbilityUses1 struct {
				EnchantressEnchant int `json:"enchantress_enchant"`
			} `json:"ability_uses,omitempty"`
			ItemUses1 struct {
				WardDispenser int `json:"ward_dispenser"`
				WardObserver  int `json:"ward_observer"`
				TrustyShovel  int `json:"trusty_shovel"`
			} `json:"item_uses,omitempty"`
			AbilityUses2 struct {
				TemplarAssassinRefraction  int `json:"templar_assassin_refraction"`
				TemplarAssassinPsionicTrap int `json:"templar_assassin_psionic_trap"`
			} `json:"ability_uses,omitempty"`
			ItemUses2 struct {
				PowerTreads int `json:"power_treads"`
			} `json:"item_uses,omitempty"`
			DeathsPos0 struct {
				Num134 struct {
					Num104 int `json:"104"`
				} `json:"134"`
			} `json:"deaths_pos,omitempty"`
			AbilityUses3 struct {
				DeathProphetSilence      int `json:"death_prophet_silence"`
				DeathProphetSpiritSiphon int `json:"death_prophet_spirit_siphon"`
				DeathProphetCarrionSwarm int `json:"death_prophet_carrion_swarm"`
			} `json:"ability_uses,omitempty"`
			ItemUses3 struct {
				PhaseBoots   int `json:"phase_boots"`
				PogoStick    int `json:"pogo_stick"`
				SpiritVessel int `json:"spirit_vessel"`
			} `json:"item_uses,omitempty"`
			AbilityUses4 struct {
				MiranaLeap     int `json:"mirana_leap"`
				MiranaArrow    int `json:"mirana_arrow"`
				MiranaStarfall int `json:"mirana_starfall"`
			} `json:"ability_uses,omitempty"`
			ItemUses4 struct {
				UnstableWand int `json:"unstable_wand"`
				WardSentry   int `json:"ward_sentry"`
			} `json:"item_uses,omitempty"`
			AbilityUses5 struct {
				QueenofpainBlink        int `json:"queenofpain_blink"`
				QueenofpainScreamOfPain int `json:"queenofpain_scream_of_pain"`
				QueenofpainSonicWave    int `json:"queenofpain_sonic_wave"`
			} `json:"ability_uses,omitempty"`
			ItemUses5 struct {
				PowerTreads int `json:"power_treads"`
				ArcaneRing  int `json:"arcane_ring"`
			} `json:"item_uses,omitempty"`
			Killed1 struct {
				NpcDotaHeroDeathProphet int `json:"npc_dota_hero_death_prophet"`
			} `json:"killed,omitempty"`
			DeathsPos1 struct {
				Num126 struct {
					Num100 int `json:"100"`
				} `json:"126"`
			} `json:"deaths_pos,omitempty"`
			AbilityUses6 struct {
				DragonKnightDragonTail  int `json:"dragon_knight_dragon_tail"`
				DragonKnightBreatheFire int `json:"dragon_knight_breathe_fire"`
				PlusHighFive            int `json:"plus_high_five"`
			} `json:"ability_uses,omitempty"`
			ItemUses6 struct {
				SmokeOfDeceit int `json:"smoke_of_deceit"`
				PowerTreads   int `json:"power_treads"`
				Blink         int `json:"blink"`
				Tpscroll      int `json:"tpscroll"`
			} `json:"item_uses,omitempty"`
			AbilityUses7 struct {
				ChenHolyPersuasion int `json:"chen_holy_persuasion"`
			} `json:"ability_uses,omitempty"`
			ItemUses7 struct {
			} `json:"item_uses,omitempty"`
			DeathsPos2 struct {
				Num128 struct {
					Num126 int `json:"126"`
				} `json:"128"`
			} `json:"deaths_pos,omitempty"`
			AbilityUses8 struct {
				TrollWarlordWhirlingAxesMelee int `json:"troll_warlord_whirling_axes_melee"`
			} `json:"ability_uses,omitempty"`
			ItemUses8 struct {
				PowerTreads   int `json:"power_treads"`
				QuellingBlade int `json:"quelling_blade"`
			} `json:"item_uses,omitempty"`
		} `json:"players"`
	} `json:"teamfights"`
	TowerStatusDire    int `json:"tower_status_dire"`
	TowerStatusRadiant int `json:"tower_status_radiant"`
	Version            int `json:"version"`
	ReplaySalt         int `json:"replay_salt"`
	SeriesID           int `json:"series_id"`
	SeriesType         int `json:"series_type"`
	League             struct {
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
		MatchID        int64 `json:"match_id"`
		PlayerSlot     int   `json:"player_slot"`
		AbilityTargets struct {
			PrimalBeastPulverize struct {
				NpcDotaHeroChen         int `json:"npc_dota_hero_chen"`
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
			} `json:"primal_beast_pulverize"`
		} `json:"ability_targets,omitempty"`
		AbilityUpgradesArr []int `json:"ability_upgrades_arr"`
		AbilityUses        struct {
			PrimalBeastUproar           int `json:"primal_beast_uproar"`
			PrimalBeastOnslaught        int `json:"primal_beast_onslaught"`
			PrimalBeastOnslaughtRelease int `json:"primal_beast_onslaught_release"`
			PrimalBeastTrample          int `json:"primal_beast_trample"`
			PrimalBeastPulverize        int `json:"primal_beast_pulverize"`
			PlusHighFive                int `json:"plus_high_five"`
			AbilityCapture              int `json:"ability_capture"`
			PrimalBeastRockThrow        int `json:"primal_beast_rock_throw"`
		} `json:"ability_uses,omitempty"`
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
			Num11 int `json:"11"`
			Num13 int `json:"13"`
			Num14 int `json:"14"`
			Num15 int `json:"15"`
			Num16 int `json:"16"`
			Num17 int `json:"17"`
			Num19 int `json:"19"`
			Num21 int `json:"21"`
			Num23 int `json:"23"`
			Num24 int `json:"24"`
			Num27 int `json:"27"`
			Num28 int `json:"28"`
			Num31 int `json:"31"`
			Num33 int `json:"33"`
			Num37 int `json:"37"`
			Num38 int `json:"38"`
		} `json:"actions,omitempty"`
		AdditionalUnits interface{} `json:"additional_units"`
		Assists         int         `json:"assists"`
		Backpack0       int         `json:"backpack_0"`
		Backpack1       int         `json:"backpack_1"`
		Backpack2       int         `json:"backpack_2"`
		Backpack3       interface{} `json:"backpack_3"`
		BuybackLog      []struct {
			Time       int    `json:"time"`
			Slot       int    `json:"slot"`
			Type       string `json:"type"`
			PlayerSlot int    `json:"player_slot"`
		} `json:"buyback_log"`
		CampsStacked  int           `json:"camps_stacked"`
		ConnectionLog []interface{} `json:"connection_log"`
		CreepsStacked int           `json:"creeps_stacked"`
		Damage        struct {
			NpcDotaHeroChen                        int `json:"npc_dota_hero_chen"`
			NpcDotaCreepBadguysMelee               int `json:"npc_dota_creep_badguys_melee"`
			NpcDotaCreepGoodguysMelee              int `json:"npc_dota_creep_goodguys_melee"`
			NpcDotaObserverWards                   int `json:"npc_dota_observer_wards"`
			NpcDotaCreepGoodguysRanged             int `json:"npc_dota_creep_goodguys_ranged"`
			NpcDotaCreepGoodguysFlagbearer         int `json:"npc_dota_creep_goodguys_flagbearer"`
			NpcDotaHeroTrollWarlord                int `json:"npc_dota_hero_troll_warlord"`
			NpcDotaCreepBadguysRanged              int `json:"npc_dota_creep_badguys_ranged"`
			NpcDotaNeutralForestTrollBerserker     int `json:"npc_dota_neutral_forest_troll_berserker"`
			NpcDotaSentryWards                     int `json:"npc_dota_sentry_wards"`
			NpcDotaNeutralSatyrSoulstealer         int `json:"npc_dota_neutral_satyr_soulstealer"`
			NpcDotaNeutralOgreMagi                 int `json:"npc_dota_neutral_ogre_magi"`
			NpcDotaBadguysTower1Top                int `json:"npc_dota_badguys_tower1_top"`
			NpcDotaHeroDragonKnight                int `json:"npc_dota_hero_dragon_knight"`
			NpcDotaNeutralGiantWolf                int `json:"npc_dota_neutral_giant_wolf"`
			NpcDotaCreepBadguysFlagbearer          int `json:"npc_dota_creep_badguys_flagbearer"`
			NpcDotaBadguysTower1Mid                int `json:"npc_dota_badguys_tower1_mid"`
			NpcDotaNeutralWarpineRaider            int `json:"npc_dota_neutral_warpine_raider"`
			NpcDotaDarkTrollWarlordSkeletonWarrior int `json:"npc_dota_dark_troll_warlord_skeleton_warrior"`
			NpcDotaNeutralEnragedWildkin           int `json:"npc_dota_neutral_enraged_wildkin"`
			NpcDotaNeutralBlackDragon              int `json:"npc_dota_neutral_black_dragon"`
			NpcDotaNeutralBigThunderLizard         int `json:"npc_dota_neutral_big_thunder_lizard"`
			NpcDotaNeutralPolarFurbolgUrsaWarrior  int `json:"npc_dota_neutral_polar_furbolg_ursa_warrior"`
			NpcDotaNeutralPolarFurbolgChampion     int `json:"npc_dota_neutral_polar_furbolg_champion"`
			NpcDotaHeroMirana                      int `json:"npc_dota_hero_mirana"`
			NpcDotaBadguysSiege                    int `json:"npc_dota_badguys_siege"`
			NpcDotaHeroQueenofpain                 int `json:"npc_dota_hero_queenofpain"`
			NpcDotaNeutralCentaurKhan              int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaNeutralDarkTrollWarlord         int `json:"npc_dota_neutral_dark_troll_warlord"`
			NpcDotaNeutralCentaurOutrunner         int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaCreepBadguysMeleeUpgraded       int `json:"npc_dota_creep_badguys_melee_upgraded"`
			NpcDotaCreepBadguysRangedUpgraded      int `json:"npc_dota_creep_badguys_ranged_upgraded"`
			NpcDotaNeutralSatyrTrickster           int `json:"npc_dota_neutral_satyr_trickster"`
			NpcDotaNeutralMudGolem                 int `json:"npc_dota_neutral_mud_golem"`
			NpcDotaNeutralOgreMauler               int `json:"npc_dota_neutral_ogre_mauler"`
		} `json:"damage,omitempty"`
		DamageInflictor struct {
			Null                 int `json:"null"`
			PrimalBeastTrample   int `json:"primal_beast_trample"`
			PrimalBeastOnslaught int `json:"primal_beast_onslaught"`
			PrimalBeastPulverize int `json:"primal_beast_pulverize"`
			PrimalBeastRockThrow int `json:"primal_beast_rock_throw"`
			Stormcrafter         int `json:"stormcrafter"`
		} `json:"damage_inflictor,omitempty"`
		DamageInflictorReceived struct {
			Null                           int `json:"null"`
			TrollWarlordWhirlingAxesRanged int `json:"troll_warlord_whirling_axes_ranged"`
			TrollWarlordWhirlingAxesMelee  int `json:"troll_warlord_whirling_axes_melee"`
			QueenofpainSonicWave           int `json:"queenofpain_sonic_wave"`
			WarpineRaiderSeedShot          int `json:"warpine_raider_seed_shot"`
			MiranaArrow                    int `json:"mirana_arrow"`
			QueenofpainScreamOfPain        int `json:"queenofpain_scream_of_pain"`
			UrnOfShadows                   int `json:"urn_of_shadows"`
			MiranaStarfall                 int `json:"mirana_starfall"`
			DragonKnightElderDragonForm    int `json:"dragon_knight_elder_dragon_form"`
			SpiritVessel                   int `json:"spirit_vessel"`
			QueenofpainShadowStrike        int `json:"queenofpain_shadow_strike"`
			Orchid                         int `json:"orchid"`
			QueenofpainBlink               int `json:"queenofpain_blink"`
			DragonKnightDragonTail         int `json:"dragon_knight_dragon_tail"`
			DragonKnightBreatheFire        int `json:"dragon_knight_breathe_fire"`
			DragonKnightFireball           int `json:"dragon_knight_fireball"`
			WraithPact                     int `json:"wraith_pact"`
			Stormcrafter                   int `json:"stormcrafter"`
		} `json:"damage_inflictor_received,omitempty"`
		DamageTaken struct {
			NpcDotaCreepBadguysMelee           int `json:"npc_dota_creep_badguys_melee"`
			NpcDotaHeroChen                    int `json:"npc_dota_hero_chen"`
			NpcDotaHeroTrollWarlord            int `json:"npc_dota_hero_troll_warlord"`
			NpcDotaCreepBadguysFlagbearer      int `json:"npc_dota_creep_badguys_flagbearer"`
			NpcDotaCreepBadguysRanged          int `json:"npc_dota_creep_badguys_ranged"`
			NpcDotaBadguysTower1Top            int `json:"npc_dota_badguys_tower1_top"`
			NpcDotaHeroQueenofpain             int `json:"npc_dota_hero_queenofpain"`
			NpcDotaNeutralKoboldTaskmaster     int `json:"npc_dota_neutral_kobold_taskmaster"`
			NpcDotaNeutralForestTrollBerserker int `json:"npc_dota_neutral_forest_troll_berserker"`
			NpcDotaBadguysTower1Mid            int `json:"npc_dota_badguys_tower1_mid"`
			NpcDotaNeutralBigThunderLizard     int `json:"npc_dota_neutral_big_thunder_lizard"`
			NpcDotaNeutralBlackDragon          int `json:"npc_dota_neutral_black_dragon"`
			NpcDotaHeroDragonKnight            int `json:"npc_dota_hero_dragon_knight"`
			NpcDotaBadguysSiege                int `json:"npc_dota_badguys_siege"`
			NpcDotaHeroMirana                  int `json:"npc_dota_hero_mirana"`
			NpcDotaBadguysTower2Top            int `json:"npc_dota_badguys_tower2_top"`
			NpcDotaNeutralMudGolem             int `json:"npc_dota_neutral_mud_golem"`
		} `json:"damage_taken,omitempty"`
		DamageTargets struct {
			Null struct {
				NpcDotaHeroChen         int `json:"npc_dota_hero_chen"`
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
				NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
			} `json:"null"`
			PrimalBeastTrample struct {
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
				NpcDotaHeroChen         int `json:"npc_dota_hero_chen"`
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
			} `json:"primal_beast_trample"`
			PrimalBeastOnslaught struct {
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
				NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
				NpcDotaHeroChen         int `json:"npc_dota_hero_chen"`
			} `json:"primal_beast_onslaught"`
			PrimalBeastPulverize struct {
				NpcDotaHeroChen         int `json:"npc_dota_hero_chen"`
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
			} `json:"primal_beast_pulverize"`
			PrimalBeastRockThrow struct {
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
				NpcDotaHeroChen         int `json:"npc_dota_hero_chen"`
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
			} `json:"primal_beast_rock_throw"`
			Stormcrafter struct {
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
				NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
			} `json:"stormcrafter"`
		} `json:"damage_targets,omitempty"`
		Deaths            int   `json:"deaths"`
		Denies            int   `json:"denies"`
		DnT               []int `json:"dn_t"`
		FirstbloodClaimed int   `json:"firstblood_claimed"`
		Gold              int   `json:"gold"`
		GoldPerMin        int   `json:"gold_per_min"`
		GoldReasons       struct {
			Num0  int `json:"0"`
			Num1  int `json:"1"`
			Num6  int `json:"6"`
			Num11 int `json:"11"`
			Num12 int `json:"12"`
			Num13 int `json:"13"`
			Num14 int `json:"14"`
			Num15 int `json:"15"`
			Num16 int `json:"16"`
			Num17 int `json:"17"`
			Num19 int `json:"19"`
			Num20 int `json:"20"`
		} `json:"gold_reasons,omitempty"`
		GoldSpent   int   `json:"gold_spent"`
		GoldT       []int `json:"gold_t"`
		HeroDamage  int   `json:"hero_damage"`
		HeroHealing int   `json:"hero_healing"`
		HeroHits    struct {
			Null                 int `json:"null"`
			PrimalBeastTrample   int `json:"primal_beast_trample"`
			PrimalBeastOnslaught int `json:"primal_beast_onslaught"`
			PrimalBeastPulverize int `json:"primal_beast_pulverize"`
			PrimalBeastRockThrow int `json:"primal_beast_rock_throw"`
			Stormcrafter         int `json:"stormcrafter"`
		} `json:"hero_hits,omitempty"`
		HeroID      int `json:"hero_id"`
		Item0       int `json:"item_0"`
		Item1       int `json:"item_1"`
		Item2       int `json:"item_2"`
		Item3       int `json:"item_3"`
		Item4       int `json:"item_4"`
		Item5       int `json:"item_5"`
		ItemNeutral int `json:"item_neutral"`
		ItemUses    struct {
			WardDispenser   int `json:"ward_dispenser"`
			WardSentry      int `json:"ward_sentry"`
			Tango           int `json:"tango"`
			MagicStick      int `json:"magic_stick"`
			Clarity         int `json:"clarity"`
			Flask           int `json:"flask"`
			WardObserver    int `json:"ward_observer"`
			Tpscroll        int `json:"tpscroll"`
			TomeOfKnowledge int `json:"tome_of_knowledge"`
			SmokeOfDeceit   int `json:"smoke_of_deceit"`
			ArcaneBoots     int `json:"arcane_boots"`
			MagicWand       int `json:"magic_wand"`
			Blink           int `json:"blink"`
			Dust            int `json:"dust"`
			Cyclone         int `json:"cyclone"`
			Stormcrafter    int `json:"stormcrafter"`
		} `json:"item_uses,omitempty"`
		KillStreaks struct {
		} `json:"kill_streaks,omitempty"`
		Killed struct {
			NpcDotaCreepGoodguysMelee              int `json:"npc_dota_creep_goodguys_melee"`
			NpcDotaObserverWards                   int `json:"npc_dota_observer_wards"`
			NpcDotaNeutralForestTrollBerserker     int `json:"npc_dota_neutral_forest_troll_berserker"`
			NpcDotaSentryWards                     int `json:"npc_dota_sentry_wards"`
			NpcDotaCreepBadguysMelee               int `json:"npc_dota_creep_badguys_melee"`
			NpcDotaNeutralGiantWolf                int `json:"npc_dota_neutral_giant_wolf"`
			NpcDotaDarkTrollWarlordSkeletonWarrior int `json:"npc_dota_dark_troll_warlord_skeleton_warrior"`
			NpcDotaCreepBadguysRanged              int `json:"npc_dota_creep_badguys_ranged"`
			NpcDotaCreepBadguysFlagbearer          int `json:"npc_dota_creep_badguys_flagbearer"`
			NpcDotaHeroDragonKnight                int `json:"npc_dota_hero_dragon_knight"`
			NpcDotaBadguysSiege                    int `json:"npc_dota_badguys_siege"`
			NpcDotaNeutralCentaurOutrunner         int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaNeutralCentaurKhan              int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaHeroMirana                      int `json:"npc_dota_hero_mirana"`
			NpcDotaNeutralSatyrSoulstealer         int `json:"npc_dota_neutral_satyr_soulstealer"`
			NpcDotaNeutralDarkTrollWarlord         int `json:"npc_dota_neutral_dark_troll_warlord"`
			NpcDotaCreepBadguysMeleeUpgraded       int `json:"npc_dota_creep_badguys_melee_upgraded"`
			NpcDotaCreepBadguysRangedUpgraded      int `json:"npc_dota_creep_badguys_ranged_upgraded"`
			NpcDotaNeutralSatyrTrickster           int `json:"npc_dota_neutral_satyr_trickster"`
		} `json:"killed,omitempty"`
		KilledBy struct {
			NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
			NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
			NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
			NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
		} `json:"killed_by,omitempty"`
		Kills    int `json:"kills"`
		KillsLog []struct {
			Time int    `json:"time"`
			Key  string `json:"key"`
		} `json:"kills_log"`
		LanePos struct {
			Num70 struct {
				Num74 int `json:"74"`
				Num76 int `json:"76"`
			} `json:"70"`
			Num72 struct {
				Num78 int `json:"78"`
			} `json:"72"`
			Num74 struct {
				Num74  int `json:"74"`
				Num80  int `json:"80"`
				Num84  int `json:"84"`
				Num86  int `json:"86"`
				Num90  int `json:"90"`
				Num92  int `json:"92"`
				Num94  int `json:"94"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num104 int `json:"104"`
				Num106 int `json:"106"`
				Num110 int `json:"110"`
			} `json:"74"`
			Num76 struct {
				Num76  int `json:"76"`
				Num112 int `json:"112"`
				Num114 int `json:"114"`
				Num118 int `json:"118"`
				Num120 int `json:"120"`
				Num124 int `json:"124"`
				Num126 int `json:"126"`
				Num130 int `json:"130"`
				Num164 int `json:"164"`
			} `json:"76"`
			Num78 struct {
				Num78  int `json:"78"`
				Num132 int `json:"132"`
				Num134 int `json:"134"`
				Num138 int `json:"138"`
				Num140 int `json:"140"`
				Num150 int `json:"150"`
				Num152 int `json:"152"`
				Num154 int `json:"154"`
				Num156 int `json:"156"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
			} `json:"78"`
			Num80 struct {
				Num80  int `json:"80"`
				Num140 int `json:"140"`
				Num142 int `json:"142"`
				Num144 int `json:"144"`
				Num146 int `json:"146"`
				Num150 int `json:"150"`
				Num152 int `json:"152"`
				Num154 int `json:"154"`
				Num156 int `json:"156"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
			} `json:"80"`
			Num82 struct {
				Num82  int `json:"82"`
				Num144 int `json:"144"`
				Num146 int `json:"146"`
				Num148 int `json:"148"`
				Num150 int `json:"150"`
				Num152 int `json:"152"`
				Num154 int `json:"154"`
				Num156 int `json:"156"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
				Num172 int `json:"172"`
			} `json:"82"`
			Num84 struct {
				Num84  int `json:"84"`
				Num144 int `json:"144"`
				Num146 int `json:"146"`
				Num148 int `json:"148"`
				Num154 int `json:"154"`
				Num156 int `json:"156"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
			} `json:"84"`
			Num86 struct {
				Num86  int `json:"86"`
				Num150 int `json:"150"`
				Num152 int `json:"152"`
				Num154 int `json:"154"`
				Num156 int `json:"156"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
				Num172 int `json:"172"`
				Num186 int `json:"186"`
			} `json:"86"`
			Num88 struct {
				Num88  int `json:"88"`
				Num150 int `json:"150"`
				Num154 int `json:"154"`
				Num156 int `json:"156"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
				Num174 int `json:"174"`
				Num184 int `json:"184"`
			} `json:"88"`
			Num90 struct {
				Num90  int `json:"90"`
				Num148 int `json:"148"`
				Num158 int `json:"158"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
				Num176 int `json:"176"`
				Num182 int `json:"182"`
			} `json:"90"`
			Num92 struct {
				Num92  int `json:"92"`
				Num146 int `json:"146"`
				Num156 int `json:"156"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
				Num172 int `json:"172"`
				Num176 int `json:"176"`
				Num180 int `json:"180"`
			} `json:"92"`
			Num94 struct {
				Num94  int `json:"94"`
				Num146 int `json:"146"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
				Num170 int `json:"170"`
				Num176 int `json:"176"`
				Num178 int `json:"178"`
			} `json:"94"`
			Num96 struct {
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num144 int `json:"144"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
				Num176 int `json:"176"`
				Num178 int `json:"178"`
			} `json:"96"`
			Num98 struct {
				Num100 int `json:"100"`
				Num102 int `json:"102"`
				Num142 int `json:"142"`
				Num144 int `json:"144"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
				Num172 int `json:"172"`
				Num174 int `json:"174"`
				Num176 int `json:"176"`
			} `json:"98"`
			Num100 struct {
				Num104 int `json:"104"`
				Num106 int `json:"106"`
				Num140 int `json:"140"`
				Num142 int `json:"142"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
				Num176 int `json:"176"`
			} `json:"100"`
			Num102 struct {
				Num108 int `json:"108"`
				Num140 int `json:"140"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
				Num174 int `json:"174"`
				Num176 int `json:"176"`
			} `json:"102"`
			Num104 struct {
				Num140 int `json:"140"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num166 int `json:"166"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
			} `json:"104"`
			Num106 struct {
				Num110 int `json:"110"`
				Num138 int `json:"138"`
				Num140 int `json:"140"`
				Num156 int `json:"156"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
				Num170 int `json:"170"`
			} `json:"106"`
			Num108 struct {
				Num112 int `json:"112"`
				Num136 int `json:"136"`
				Num138 int `json:"138"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
			} `json:"108"`
			Num110 struct {
				Num114 int `json:"114"`
				Num116 int `json:"116"`
				Num136 int `json:"136"`
				Num158 int `json:"158"`
			} `json:"110"`
			Num112 struct {
				Num118 int `json:"118"`
				Num124 int `json:"124"`
				Num126 int `json:"126"`
				Num128 int `json:"128"`
				Num130 int `json:"130"`
				Num132 int `json:"132"`
				Num134 int `json:"134"`
				Num136 int `json:"136"`
				Num138 int `json:"138"`
				Num142 int `json:"142"`
				Num158 int `json:"158"`
			} `json:"112"`
			Num114 struct {
				Num120 int `json:"120"`
				Num122 int `json:"122"`
				Num124 int `json:"124"`
				Num126 int `json:"126"`
				Num128 int `json:"128"`
				Num136 int `json:"136"`
				Num138 int `json:"138"`
				Num142 int `json:"142"`
				Num156 int `json:"156"`
				Num158 int `json:"158"`
			} `json:"114"`
			Num116 struct {
				Num120 int `json:"120"`
				Num122 int `json:"122"`
				Num124 int `json:"124"`
				Num140 int `json:"140"`
				Num156 int `json:"156"`
			} `json:"116"`
			Num118 struct {
				Num118 int `json:"118"`
				Num134 int `json:"134"`
				Num138 int `json:"138"`
				Num140 int `json:"140"`
				Num142 int `json:"142"`
				Num144 int `json:"144"`
				Num146 int `json:"146"`
				Num156 int `json:"156"`
			} `json:"118"`
			Num120 struct {
				Num118 int `json:"118"`
				Num128 int `json:"128"`
				Num130 int `json:"130"`
				Num132 int `json:"132"`
				Num148 int `json:"148"`
				Num150 int `json:"150"`
				Num154 int `json:"154"`
			} `json:"120"`
			Num122 struct {
				Num116 int `json:"116"`
				Num126 int `json:"126"`
				Num146 int `json:"146"`
				Num152 int `json:"152"`
				Num156 int `json:"156"`
			} `json:"122"`
			Num124 struct {
				Num124 int `json:"124"`
				Num146 int `json:"146"`
				Num150 int `json:"150"`
				Num152 int `json:"152"`
				Num154 int `json:"154"`
			} `json:"124"`
			Num126 struct {
				Num114 int `json:"114"`
				Num122 int `json:"122"`
				Num134 int `json:"134"`
				Num138 int `json:"138"`
				Num144 int `json:"144"`
				Num148 int `json:"148"`
			} `json:"126"`
			Num128 struct {
				Num114 int `json:"114"`
				Num122 int `json:"122"`
				Num130 int `json:"130"`
				Num132 int `json:"132"`
				Num134 int `json:"134"`
				Num136 int `json:"136"`
				Num140 int `json:"140"`
				Num142 int `json:"142"`
				Num144 int `json:"144"`
				Num146 int `json:"146"`
			} `json:"128"`
			Num130 struct {
				Num114 int `json:"114"`
				Num120 int `json:"120"`
				Num132 int `json:"132"`
				Num142 int `json:"142"`
			} `json:"130"`
			Num132 struct {
				Num114 int `json:"114"`
				Num116 int `json:"116"`
				Num120 int `json:"120"`
				Num128 int `json:"128"`
				Num130 int `json:"130"`
				Num132 int `json:"132"`
			} `json:"132"`
			Num134 struct {
				Num118 int `json:"118"`
				Num130 int `json:"130"`
				Num132 int `json:"132"`
			} `json:"134"`
			Num136 struct {
				Num118 int `json:"118"`
				Num128 int `json:"128"`
				Num132 int `json:"132"`
				Num136 int `json:"136"`
			} `json:"136"`
			Num138 struct {
				Num128 int `json:"128"`
				Num132 int `json:"132"`
				Num134 int `json:"134"`
			} `json:"138"`
			Num140 struct {
				Num126 int `json:"126"`
			} `json:"140"`
			Num142 struct {
				Num128 int `json:"128"`
				Num130 int `json:"130"`
			} `json:"142"`
		} `json:"lane_pos,omitempty"`
		LastHits     int   `json:"last_hits"`
		LeaverStatus int   `json:"leaver_status"`
		Level        int   `json:"level"`
		LhT          []int `json:"lh_t"`
		LifeState    struct {
			Num0 int `json:"0"`
			Num1 int `json:"1"`
			Num2 int `json:"2"`
		} `json:"life_state,omitempty"`
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
		} `json:"multi_kills,omitempty"`
		NetWorth int `json:"net_worth"`
		Obs      struct {
			Num76 struct {
				Num154 int `json:"154"`
			} `json:"76"`
			Num92 struct {
				Num80 int `json:"80"`
			} `json:"92"`
			Num94 struct {
				Num138 int `json:"138"`
			} `json:"94"`
			Num108 struct {
				Num140 int `json:"140"`
			} `json:"108"`
			Num110 struct {
				Num184 int `json:"184"`
			} `json:"110"`
			Num120 struct {
				Num142 int `json:"142"`
			} `json:"120"`
			Num132 struct {
				Num158 int `json:"158"`
			} `json:"132"`
			Num144 struct {
				Num122 int `json:"122"`
			} `json:"144"`
			Num154 struct {
				Num80 int `json:"80"`
			} `json:"154"`
		} `json:"obs,omitempty"`
		ObsLeftLog []struct {
			Time         int    `json:"time"`
			Type         string `json:"type"`
			Key          string `json:"key"`
			Slot         int    `json:"slot"`
			Attackername string `json:"attackername"`
			X            int    `json:"x"`
			Y            int    `json:"y"`
			Z            int    `json:"z"`
			Entityleft   bool   `json:"entityleft"`
			Ehandle      int    `json:"ehandle"`
			PlayerSlot   int    `json:"player_slot"`
		} `json:"obs_left_log"`
		ObsLog []struct {
			Time       int    `json:"time"`
			Type       string `json:"type"`
			Key        string `json:"key"`
			Slot       int    `json:"slot"`
			X          int    `json:"x"`
			Y          int    `json:"y"`
			Z          int    `json:"z"`
			Entityleft bool   `json:"entityleft"`
			Ehandle    int    `json:"ehandle"`
			PlayerSlot int    `json:"player_slot"`
		} `json:"obs_log"`
		ObsPlaced         int         `json:"obs_placed"`
		PartyID           int         `json:"party_id"`
		PartySize         int         `json:"party_size"`
		PerformanceOthers interface{} `json:"performance_others"`
		PermanentBuffs    []struct {
			PermanentBuff int `json:"permanent_buff"`
			StackCount    int `json:"stack_count"`
			GrantTime     int `json:"grant_time"`
		} `json:"permanent_buffs"`
		Pings    int  `json:"pings"`
		PredVict bool `json:"pred_vict"`
		Purchase struct {
			WardDispenser   int `json:"ward_dispenser"`
			MagicStick      int `json:"magic_stick"`
			Tango           int `json:"tango"`
			Branches        int `json:"branches"`
			Flask           int `json:"flask"`
			WardSentry      int `json:"ward_sentry"`
			WardObserver    int `json:"ward_observer"`
			WindLace        int `json:"wind_lace"`
			Clarity         int `json:"clarity"`
			InfusedRaindrop int `json:"infused_raindrop"`
			Boots           int `json:"boots"`
			TomeOfKnowledge int `json:"tome_of_knowledge"`
			Tpscroll        int `json:"tpscroll"`
			EnergyBooster   int `json:"energy_booster"`
			SmokeOfDeceit   int `json:"smoke_of_deceit"`
			ArcaneBoots     int `json:"arcane_boots"`
			Dust            int `json:"dust"`
			RecipeMagicWand int `json:"recipe_magic_wand"`
			MagicWand       int `json:"magic_wand"`
			Blink           int `json:"blink"`
			AghanimsShard   int `json:"aghanims_shard"`
			StaffOfWizardry int `json:"staff_of_wizardry"`
			RecipeCyclone   int `json:"recipe_cyclone"`
			VoidStone       int `json:"void_stone"`
			Cyclone         int `json:"cyclone"`
		} `json:"purchase,omitempty"`
		PurchaseLog []struct {
			Time int    `json:"time"`
			Key  string `json:"key"`
		} `json:"purchase_log"`
		Randomed      bool        `json:"randomed"`
		Repicked      interface{} `json:"repicked"`
		RoshansKilled int         `json:"roshans_killed"`
		RunePickups   int         `json:"rune_pickups"`
		Runes         struct {
			Num5 int `json:"5"`
		} `json:"runes,omitempty"`
		RunesLog []struct {
			Time int `json:"time"`
			Key  int `json:"key"`
		} `json:"runes_log"`
		Sen struct {
			Num76 struct {
				Num154 int `json:"154"`
			} `json:"76"`
			Num90 struct {
				Num154 int `json:"154"`
				Num160 int `json:"160"`
				Num164 int `json:"164"`
			} `json:"90"`
			Num94 struct {
				Num76  int `json:"76"`
				Num132 int `json:"132"`
			} `json:"94"`
			Num106 struct {
				Num156 int `json:"156"`
			} `json:"106"`
			Num108 struct {
				Num158 int `json:"158"`
			} `json:"108"`
			Num110 struct {
				Num140 int `json:"140"`
			} `json:"110"`
			Num114 struct {
				Num146 int `json:"146"`
			} `json:"114"`
			Num122 struct {
				Num102 int `json:"102"`
			} `json:"122"`
			Num126 struct {
				Num162 int `json:"162"`
			} `json:"126"`
			Num130 struct {
				Num100 int `json:"100"`
			} `json:"130"`
			Num136 struct {
				Num146 int `json:"146"`
			} `json:"136"`
			Num144 struct {
				Num120 int `json:"120"`
			} `json:"144"`
			Num146 struct {
				Num122 int `json:"122"`
			} `json:"146"`
		} `json:"sen,omitempty"`
		SenLeftLog []struct {
			Time         int    `json:"time"`
			Type         string `json:"type"`
			Key          string `json:"key"`
			Slot         int    `json:"slot"`
			Attackername string `json:"attackername"`
			X            int    `json:"x"`
			Y            int    `json:"y"`
			Z            int    `json:"z"`
			Entityleft   bool   `json:"entityleft"`
			Ehandle      int    `json:"ehandle"`
			PlayerSlot   int    `json:"player_slot"`
		} `json:"sen_left_log"`
		SenLog []struct {
			Time       int    `json:"time"`
			Type       string `json:"type"`
			Key        string `json:"key"`
			Slot       int    `json:"slot"`
			X          int    `json:"x"`
			Y          int    `json:"y"`
			Z          int    `json:"z"`
			Entityleft bool   `json:"entityleft"`
			Ehandle    int    `json:"ehandle"`
			PlayerSlot int    `json:"player_slot"`
		} `json:"sen_log"`
		SenPlaced              int     `json:"sen_placed"`
		Stuns                  float64 `json:"stuns"`
		TeamfightParticipation float64 `json:"teamfight_participation"`
		Times                  []int   `json:"times"`
		TowerDamage            int     `json:"tower_damage"`
		TowersKilled           int     `json:"towers_killed"`
		XpPerMin               int     `json:"xp_per_min"`
		XpReasons              struct {
			Num0 int `json:"0"`
			Num1 int `json:"1"`
			Num2 int `json:"2"`
			Num4 int `json:"4"`
			Num5 int `json:"5"`
		} `json:"xp_reasons,omitempty"`
		XpT               []int     `json:"xp_t"`
		Personaname       string    `json:"personaname"`
		Name              string    `json:"name"`
		LastLogin         time.Time `json:"last_login"`
		RadiantWin        bool      `json:"radiant_win"`
		StartTime         int       `json:"start_time"`
		Duration          int       `json:"duration"`
		Cluster           int       `json:"cluster"`
		LobbyType         int       `json:"lobby_type"`
		GameMode          int       `json:"game_mode"`
		IsContributor     bool      `json:"is_contributor"`
		Patch             int       `json:"patch"`
		Region            int       `json:"region"`
		IsRadiant         bool      `json:"isRadiant"`
		Win               int       `json:"win"`
		Lose              int       `json:"lose"`
		TotalGold         int       `json:"total_gold"`
		TotalXp           int       `json:"total_xp"`
		KillsPerMin       float64   `json:"kills_per_min,omitempty"`
		Kda               int       `json:"kda"`
		Abandons          int       `json:"abandons"`
		NeutralKills      int       `json:"neutral_kills"`
		TowerKills        int       `json:"tower_kills"`
		CourierKills      int       `json:"courier_kills"`
		LaneKills         int       `json:"lane_kills"`
		HeroKills         int       `json:"hero_kills"`
		ObserverKills     int       `json:"observer_kills"`
		SentryKills       int       `json:"sentry_kills"`
		RoshanKills       int       `json:"roshan_kills"`
		NecronomiconKills int       `json:"necronomicon_kills"`
		AncientKills      int       `json:"ancient_kills"`
		BuybackCount      int       `json:"buyback_count"`
		ObserverUses      int       `json:"observer_uses"`
		SentryUses        int       `json:"sentry_uses"`
		LaneEfficiency    float64   `json:"lane_efficiency"`
		LaneEfficiencyPct int       `json:"lane_efficiency_pct"`
		Lane              int       `json:"lane"`
		LaneRole          int       `json:"lane_role"`
		IsRoaming         bool      `json:"is_roaming"`
		PurchaseTime      struct {
			MagicStick      int `json:"magic_stick"`
			Tango           int `json:"tango"`
			Branches        int `json:"branches"`
			Flask           int `json:"flask"`
			WardSentry      int `json:"ward_sentry"`
			WardObserver    int `json:"ward_observer"`
			WindLace        int `json:"wind_lace"`
			Clarity         int `json:"clarity"`
			InfusedRaindrop int `json:"infused_raindrop"`
			Boots           int `json:"boots"`
			TomeOfKnowledge int `json:"tome_of_knowledge"`
			Tpscroll        int `json:"tpscroll"`
			EnergyBooster   int `json:"energy_booster"`
			SmokeOfDeceit   int `json:"smoke_of_deceit"`
			ArcaneBoots     int `json:"arcane_boots"`
			Dust            int `json:"dust"`
			MagicWand       int `json:"magic_wand"`
			Blink           int `json:"blink"`
			AghanimsShard   int `json:"aghanims_shard"`
			StaffOfWizardry int `json:"staff_of_wizardry"`
			VoidStone       int `json:"void_stone"`
			Cyclone         int `json:"cyclone"`
		} `json:"purchase_time,omitempty"`
		FirstPurchaseTime struct {
			MagicStick      int `json:"magic_stick"`
			Tango           int `json:"tango"`
			Branches        int `json:"branches"`
			Flask           int `json:"flask"`
			WardSentry      int `json:"ward_sentry"`
			WardObserver    int `json:"ward_observer"`
			WindLace        int `json:"wind_lace"`
			Clarity         int `json:"clarity"`
			InfusedRaindrop int `json:"infused_raindrop"`
			Boots           int `json:"boots"`
			TomeOfKnowledge int `json:"tome_of_knowledge"`
			Tpscroll        int `json:"tpscroll"`
			EnergyBooster   int `json:"energy_booster"`
			SmokeOfDeceit   int `json:"smoke_of_deceit"`
			ArcaneBoots     int `json:"arcane_boots"`
			Dust            int `json:"dust"`
			MagicWand       int `json:"magic_wand"`
			Blink           int `json:"blink"`
			AghanimsShard   int `json:"aghanims_shard"`
			StaffOfWizardry int `json:"staff_of_wizardry"`
			VoidStone       int `json:"void_stone"`
			Cyclone         int `json:"cyclone"`
		} `json:"first_purchase_time,omitempty"`
		ItemWin struct {
			MagicStick      int `json:"magic_stick"`
			Tango           int `json:"tango"`
			Branches        int `json:"branches"`
			Flask           int `json:"flask"`
			WardSentry      int `json:"ward_sentry"`
			WardObserver    int `json:"ward_observer"`
			WindLace        int `json:"wind_lace"`
			Clarity         int `json:"clarity"`
			InfusedRaindrop int `json:"infused_raindrop"`
			Boots           int `json:"boots"`
			TomeOfKnowledge int `json:"tome_of_knowledge"`
			Tpscroll        int `json:"tpscroll"`
			EnergyBooster   int `json:"energy_booster"`
			SmokeOfDeceit   int `json:"smoke_of_deceit"`
			ArcaneBoots     int `json:"arcane_boots"`
			Dust            int `json:"dust"`
			MagicWand       int `json:"magic_wand"`
			Blink           int `json:"blink"`
			AghanimsShard   int `json:"aghanims_shard"`
			StaffOfWizardry int `json:"staff_of_wizardry"`
			VoidStone       int `json:"void_stone"`
			Cyclone         int `json:"cyclone"`
		} `json:"item_win,omitempty"`
		ItemUsage struct {
			MagicStick      int `json:"magic_stick"`
			Tango           int `json:"tango"`
			Branches        int `json:"branches"`
			Flask           int `json:"flask"`
			WardSentry      int `json:"ward_sentry"`
			WardObserver    int `json:"ward_observer"`
			WindLace        int `json:"wind_lace"`
			Clarity         int `json:"clarity"`
			InfusedRaindrop int `json:"infused_raindrop"`
			Boots           int `json:"boots"`
			TomeOfKnowledge int `json:"tome_of_knowledge"`
			Tpscroll        int `json:"tpscroll"`
			EnergyBooster   int `json:"energy_booster"`
			SmokeOfDeceit   int `json:"smoke_of_deceit"`
			ArcaneBoots     int `json:"arcane_boots"`
			Dust            int `json:"dust"`
			MagicWand       int `json:"magic_wand"`
			Blink           int `json:"blink"`
			AghanimsShard   int `json:"aghanims_shard"`
			StaffOfWizardry int `json:"staff_of_wizardry"`
			VoidStone       int `json:"void_stone"`
			Cyclone         int `json:"cyclone"`
		} `json:"item_usage,omitempty"`
		PurchaseWardObserver int  `json:"purchase_ward_observer,omitempty"`
		PurchaseWardSentry   int  `json:"purchase_ward_sentry,omitempty"`
		PurchaseTpscroll     int  `json:"purchase_tpscroll,omitempty"`
		ActionsPerMin        int  `json:"actions_per_min"`
		LifeStateDead        int  `json:"life_state_dead"`
		RankTier             int  `json:"rank_tier"`
		IsSubscriber         bool `json:"is_subscriber"`
		Cosmetics            []struct {
			ItemID          int         `json:"item_id"`
			Name            string      `json:"name"`
			Prefab          string      `json:"prefab"`
			CreationDate    time.Time   `json:"creation_date"`
			ImageInventory  string      `json:"image_inventory"`
			ImagePath       string      `json:"image_path"`
			ItemDescription string      `json:"item_description"`
			ItemName        string      `json:"item_name"`
			ItemRarity      string      `json:"item_rarity"`
			ItemTypeName    string      `json:"item_type_name"`
			UsedByHeroes    interface{} `json:"used_by_heroes"`
		} `json:"cosmetics"`
		// Benchmarks struct {
		// 	GoldPerMin struct {
		// 		Raw int     `json:"raw"`
		// 		Pct float64 `json:"pct"`
		// 	} `json:"gold_per_min"`
		// 	XpPerMin struct {
		// 		Raw int     `json:"raw"`
		// 		Pct float64 `json:"pct"`
		// 	} `json:"xp_per_min"`
		// 	KillsPerMin struct {
		// 		Raw float64 `json:"raw"`
		// 		Pct float64 `json:"pct"`
		// 	} `json:"kills_per_min"`
		// 	LastHitsPerMin struct {
		// 		Raw float64 `json:"raw"`
		// 		Pct float64 `json:"pct"`
		// 	} `json:"last_hits_per_min"`
		// 	HeroDamagePerMin struct {
		// 		Raw float64 `json:"raw"`
		// 		Pct float64 `json:"pct"`
		// 	} `json:"hero_damage_per_min"`
		// 	HeroHealingPerMin struct {
		// 		Raw int     `json:"raw"`
		// 		Pct float64 `json:"pct"`
		// 	} `json:"hero_healing_per_min"`
		// 	TowerDamage struct {
		// 		Raw int     `json:"raw"`
		// 		Pct float64 `json:"pct"`
		// 	} `json:"tower_damage"`
		// 	StunsPerMin struct {
		// 		Raw float64 `json:"raw"`
		// 		Pct float64 `json:"pct"`
		// 	} `json:"stuns_per_min"`
		// 	Lhten struct {
		// 		Raw int     `json:"raw"`
		// 		Pct float64 `json:"pct"`
		// 	} `json:"lhten"`
		// } `json:"benchmarks"`
		AbilityTargets0 struct {
		} `json:"ability_targets,omitempty"`
		AbilityUses0 struct {
			PuckIllusoryOrb     int `json:"puck_illusory_orb"`
			PuckEtherealJaunt   int `json:"puck_ethereal_jaunt"`
			PuckPhaseShift      int `json:"puck_phase_shift"`
			SeasonalTi11Balloon int `json:"seasonal_ti11_balloon"`
			PuckWaningRift      int `json:"puck_waning_rift"`
			PuckDreamCoil       int `json:"puck_dream_coil"`
		} `json:"ability_uses,omitempty"`
		Actions0 struct {
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
			Num12 int `json:"12"`
			Num13 int `json:"13"`
			Num14 int `json:"14"`
			Num15 int `json:"15"`
			Num16 int `json:"16"`
			Num17 int `json:"17"`
			Num19 int `json:"19"`
			Num24 int `json:"24"`
			Num28 int `json:"28"`
			Num32 int `json:"32"`
			Num37 int `json:"37"`
			Num38 int `json:"38"`
		} `json:"actions,omitempty"`
		Damage0 struct {
			NpcDotaCreepBadguysMelee               int `json:"npc_dota_creep_badguys_melee"`
			NpcDotaCreepGoodguysRanged             int `json:"npc_dota_creep_goodguys_ranged"`
			NpcDotaCreepBadguysRanged              int `json:"npc_dota_creep_badguys_ranged"`
			NpcDotaHeroQueenofpain                 int `json:"npc_dota_hero_queenofpain"`
			NpcDotaCreepGoodguysMelee              int `json:"npc_dota_creep_goodguys_melee"`
			NpcDotaCreepBadguysFlagbearer          int `json:"npc_dota_creep_badguys_flagbearer"`
			NpcDotaCreepGoodguysFlagbearer         int `json:"npc_dota_creep_goodguys_flagbearer"`
			NpcDotaBadguysSiege                    int `json:"npc_dota_badguys_siege"`
			NpcDotaBadguysTower1Mid                int `json:"npc_dota_badguys_tower1_mid"`
			NpcDotaHeroDragonKnight                int `json:"npc_dota_hero_dragon_knight"`
			NpcDotaNeutralDarkTroll                int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaCourier                         int `json:"npc_dota_courier"`
			NpcDotaNeutralGiantWolf                int `json:"npc_dota_neutral_giant_wolf"`
			NpcDotaNeutralMudGolemSplit            int `json:"npc_dota_neutral_mud_golem_split"`
			NpcDotaHeroMirana                      int `json:"npc_dota_hero_mirana"`
			NpcDotaHeroChen                        int `json:"npc_dota_hero_chen"`
			NpcDotaBadguysTower2Mid                int `json:"npc_dota_badguys_tower2_mid"`
			NpcDotaNeutralSatyrTrickster           int `json:"npc_dota_neutral_satyr_trickster"`
			NpcDotaNeutralSatyrSoulstealer         int `json:"npc_dota_neutral_satyr_soulstealer"`
			NpcDotaNeutralSatyrHellcaller          int `json:"npc_dota_neutral_satyr_hellcaller"`
			NpcDotaSentryWards                     int `json:"npc_dota_sentry_wards"`
			NpcDotaNeutralRockGolem                int `json:"npc_dota_neutral_rock_golem"`
			NpcDotaNeutralSmallThunderLizard       int `json:"npc_dota_neutral_small_thunder_lizard"`
			NpcDotaNeutralBlackDragon              int `json:"npc_dota_neutral_black_dragon"`
			NpcDotaNeutralGraniteGolem             int `json:"npc_dota_neutral_granite_golem"`
			NpcDotaNeutralBlackDrake               int `json:"npc_dota_neutral_black_drake"`
			NpcDotaNeutralBigThunderLizard         int `json:"npc_dota_neutral_big_thunder_lizard"`
			NpcDotaNeutralEnragedWildkin           int `json:"npc_dota_neutral_enraged_wildkin"`
			NpcDotaNeutralWildkin                  int `json:"npc_dota_neutral_wildkin"`
			NpcDotaNeutralOgreMagi                 int `json:"npc_dota_neutral_ogre_magi"`
			NpcDotaNeutralOgreMauler               int `json:"npc_dota_neutral_ogre_mauler"`
			NpcDotaNeutralWarpineRaider            int `json:"npc_dota_neutral_warpine_raider"`
			NpcDotaHeroTrollWarlord                int `json:"npc_dota_hero_troll_warlord"`
			NpcDotaNeutralKobold                   int `json:"npc_dota_neutral_kobold"`
			NpcDotaNeutralKoboldTunneler           int `json:"npc_dota_neutral_kobold_tunneler"`
			NpcDotaNeutralKoboldTaskmaster         int `json:"npc_dota_neutral_kobold_taskmaster"`
			NpcDotaNeutralMudGolem                 int `json:"npc_dota_neutral_mud_golem"`
			NpcDotaNeutralCentaurKhan              int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaNeutralCentaurOutrunner         int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaNeutralPolarFurbolgUrsaWarrior  int `json:"npc_dota_neutral_polar_furbolg_ursa_warrior"`
			NpcDotaNeutralPolarFurbolgChampion     int `json:"npc_dota_neutral_polar_furbolg_champion"`
			NpcDotaRoshan                          int `json:"npc_dota_roshan"`
			NpcDotaObserverWards                   int `json:"npc_dota_observer_wards"`
			NpcDotaNeutralAlphaWolf                int `json:"npc_dota_neutral_alpha_wolf"`
			NpcDotaBadguysTower2Top                int `json:"npc_dota_badguys_tower2_top"`
			NpcDotaDarkTrollWarlordSkeletonWarrior int `json:"npc_dota_dark_troll_warlord_skeleton_warrior"`
			NpcDotaNeutralForestTrollHighPriest    int `json:"npc_dota_neutral_forest_troll_high_priest"`
			NpcDotaNeutralForestTrollBerserker     int `json:"npc_dota_neutral_forest_troll_berserker"`
			NpcDotaBadguysTower2Bot                int `json:"npc_dota_badguys_tower2_bot"`
			NpcDotaItemWraithPactTotem             int `json:"npc_dota_item_wraith_pact_totem"`
			NpcDotaNeutralGhost                    int `json:"npc_dota_neutral_ghost"`
			NpcDotaNeutralFelBeast                 int `json:"npc_dota_neutral_fel_beast"`
			NpcDotaCreepBadguysRangedUpgraded      int `json:"npc_dota_creep_badguys_ranged_upgraded"`
			NpcDotaCreepBadguysMeleeUpgraded       int `json:"npc_dota_creep_badguys_melee_upgraded"`
			NpcDotaCreepBadguysFlagbearerUpgraded  int `json:"npc_dota_creep_badguys_flagbearer_upgraded"`
			NpcDotaBadguysSiegeUpgraded            int `json:"npc_dota_badguys_siege_upgraded"`
		} `json:"damage,omitempty"`
		DamageInflictor0 struct {
			PuckIllusoryOrb   int `json:"puck_illusory_orb"`
			Null              int `json:"null"`
			PuckDreamCoil     int `json:"puck_dream_coil"`
			PuckWaningRift    int `json:"puck_waning_rift"`
			WitchBlade        int `json:"witch_blade"`
			OverwhelmingBlink int `json:"overwhelming_blink"`
		} `json:"damage_inflictor,omitempty"`
		DamageInflictorReceived0 struct {
			Null                          int `json:"null"`
			QueenofpainScreamOfPain       int `json:"queenofpain_scream_of_pain"`
			MiranaArrow                   int `json:"mirana_arrow"`
			DragonKnightDragonTail        int `json:"dragon_knight_dragon_tail"`
			DragonKnightBreatheFire       int `json:"dragon_knight_breathe_fire"`
			CentaurKhanWarStomp           int `json:"centaur_khan_war_stomp"`
			Orchid                        int `json:"orchid"`
			DragonKnightElderDragonForm   int `json:"dragon_knight_elder_dragon_form"`
			QueenofpainSonicWave          int `json:"queenofpain_sonic_wave"`
			SpiritVessel                  int `json:"spirit_vessel"`
			Bfury                         int `json:"bfury"`
			TrollWarlordWhirlingAxesMelee int `json:"troll_warlord_whirling_axes_melee"`
			QueenofpainBlink              int `json:"queenofpain_blink"`
			DragonKnightFireball          int `json:"dragon_knight_fireball"`
			QueenofpainShadowStrike       int `json:"queenofpain_shadow_strike"`
			Stormcrafter                  int `json:"stormcrafter"`
			WraithPact                    int `json:"wraith_pact"`
			ShivasGuard                   int `json:"shivas_guard"`
			MiranaStarfall                int `json:"mirana_starfall"`
		} `json:"damage_inflictor_received,omitempty"`
		DamageTaken0 struct {
			NpcDotaHeroQueenofpain                int `json:"npc_dota_hero_queenofpain"`
			NpcDotaCreepBadguysMelee              int `json:"npc_dota_creep_badguys_melee"`
			NpcDotaCreepBadguysFlagbearer         int `json:"npc_dota_creep_badguys_flagbearer"`
			NpcDotaCreepBadguysRanged             int `json:"npc_dota_creep_badguys_ranged"`
			NpcDotaHeroMirana                     int `json:"npc_dota_hero_mirana"`
			NpcDotaNeutralDarkTroll               int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaBadguysTower1Mid               int `json:"npc_dota_badguys_tower1_mid"`
			NpcDotaBadguysSiege                   int `json:"npc_dota_badguys_siege"`
			NpcDotaNeutralSatyrTrickster          int `json:"npc_dota_neutral_satyr_trickster"`
			NpcDotaNeutralSatyrHellcaller         int `json:"npc_dota_neutral_satyr_hellcaller"`
			NpcDotaNeutralSatyrSoulstealer        int `json:"npc_dota_neutral_satyr_soulstealer"`
			NpcDotaNeutralRockGolem               int `json:"npc_dota_neutral_rock_golem"`
			NpcDotaNeutralBlackDragon             int `json:"npc_dota_neutral_black_dragon"`
			NpcDotaNeutralSmallThunderLizard      int `json:"npc_dota_neutral_small_thunder_lizard"`
			NpcDotaNeutralBigThunderLizard        int `json:"npc_dota_neutral_big_thunder_lizard"`
			NpcDotaNeutralGraniteGolem            int `json:"npc_dota_neutral_granite_golem"`
			NpcDotaNeutralBlackDrake              int `json:"npc_dota_neutral_black_drake"`
			NpcDotaNeutralEnragedWildkin          int `json:"npc_dota_neutral_enraged_wildkin"`
			NpcDotaNeutralWildkin                 int `json:"npc_dota_neutral_wildkin"`
			NpcDotaNeutralOgreMauler              int `json:"npc_dota_neutral_ogre_mauler"`
			NpcDotaNeutralWarpineRaider           int `json:"npc_dota_neutral_warpine_raider"`
			NpcDotaNeutralCentaurOutrunner        int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaNeutralCentaurKhan             int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaHeroDragonKnight               int `json:"npc_dota_hero_dragon_knight"`
			NpcDotaHeroChen                       int `json:"npc_dota_hero_chen"`
			NpcDotaNeutralPolarFurbolgUrsaWarrior int `json:"npc_dota_neutral_polar_furbolg_ursa_warrior"`
			NpcDotaNeutralPolarFurbolgChampion    int `json:"npc_dota_neutral_polar_furbolg_champion"`
			NpcDotaBadguysTower2Top               int `json:"npc_dota_badguys_tower2_top"`
			NpcDotaHeroTrollWarlord               int `json:"npc_dota_hero_troll_warlord"`
			NpcDotaNeutralAlphaWolf               int `json:"npc_dota_neutral_alpha_wolf"`
			NpcDotaNeutralGiantWolf               int `json:"npc_dota_neutral_giant_wolf"`
			NpcDotaNeutralForestTrollBerserker    int `json:"npc_dota_neutral_forest_troll_berserker"`
			NpcDotaNeutralForestTrollHighPriest   int `json:"npc_dota_neutral_forest_troll_high_priest"`
			NpcDotaBadguysTower2Bot               int `json:"npc_dota_badguys_tower2_bot"`
			NpcDotaCreepBadguysRangedUpgraded     int `json:"npc_dota_creep_badguys_ranged_upgraded"`
			NpcDotaCreepBadguysMeleeUpgraded      int `json:"npc_dota_creep_badguys_melee_upgraded"`
			NpcDotaCreepBadguysFlagbearerUpgraded int `json:"npc_dota_creep_badguys_flagbearer_upgraded"`
		} `json:"damage_taken,omitempty"`
		DamageTargets0 struct {
			PuckIllusoryOrb struct {
				NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
				NpcDotaHeroChen         int `json:"npc_dota_hero_chen"`
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
			} `json:"puck_illusory_orb"`
			Null struct {
				NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
				NpcDotaHeroChen         int `json:"npc_dota_hero_chen"`
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
			} `json:"null"`
			PuckDreamCoil struct {
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
				NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
				NpcDotaHeroChen         int `json:"npc_dota_hero_chen"`
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
			} `json:"puck_dream_coil"`
			PuckWaningRift struct {
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
				NpcDotaHeroChen         int `json:"npc_dota_hero_chen"`
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
				NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
			} `json:"puck_waning_rift"`
			WitchBlade struct {
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
				NpcDotaHeroChen         int `json:"npc_dota_hero_chen"`
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
			} `json:"witch_blade"`
			OverwhelmingBlink struct {
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
				NpcDotaHeroChen         int `json:"npc_dota_hero_chen"`
			} `json:"overwhelming_blink"`
		} `json:"damage_targets,omitempty"`
		GoldReasons0 struct {
			Num0  int `json:"0"`
			Num1  int `json:"1"`
			Num6  int `json:"6"`
			Num11 int `json:"11"`
			Num12 int `json:"12"`
			Num13 int `json:"13"`
			Num14 int `json:"14"`
			Num15 int `json:"15"`
			Num16 int `json:"16"`
			Num17 int `json:"17"`
			Num19 int `json:"19"`
			Num21 int `json:"21"`
		} `json:"gold_reasons,omitempty"`
		HeroHits0 struct {
			PuckIllusoryOrb   int `json:"puck_illusory_orb"`
			Null              int `json:"null"`
			PuckDreamCoil     int `json:"puck_dream_coil"`
			PuckWaningRift    int `json:"puck_waning_rift"`
			WitchBlade        int `json:"witch_blade"`
			OverwhelmingBlink int `json:"overwhelming_blink"`
		} `json:"hero_hits,omitempty"`
		ItemUses0 struct {
			WardObserver      int `json:"ward_observer"`
			Branches          int `json:"branches"`
			Tango             int `json:"tango"`
			Bottle            int `json:"bottle"`
			Tpscroll          int `json:"tpscroll"`
			WardDispenser     int `json:"ward_dispenser"`
			WardSentry        int `json:"ward_sentry"`
			PowerTreads       int `json:"power_treads"`
			Clarity           int `json:"clarity"`
			MagicWand         int `json:"magic_wand"`
			FaerieFire        int `json:"faerie_fire"`
			Blink             int `json:"blink"`
			Vambrace          int `json:"vambrace"`
			Gem               int `json:"gem"`
			OverwhelmingBlink int `json:"overwhelming_blink"`
		} `json:"item_uses,omitempty"`
		KillStreaks0 struct {
			Num3 int `json:"3"`
			Num4 int `json:"4"`
			Num5 int `json:"5"`
		} `json:"kill_streaks,omitempty"`
		Killed0 struct {
			NpcDotaCreepGoodguysRanged             int `json:"npc_dota_creep_goodguys_ranged"`
			NpcDotaCreepBadguysMelee               int `json:"npc_dota_creep_badguys_melee"`
			NpcDotaCreepBadguysRanged              int `json:"npc_dota_creep_badguys_ranged"`
			NpcDotaCreepGoodguysMelee              int `json:"npc_dota_creep_goodguys_melee"`
			NpcDotaCreepBadguysFlagbearer          int `json:"npc_dota_creep_badguys_flagbearer"`
			NpcDotaCourier                         int `json:"npc_dota_courier"`
			NpcDotaNeutralMudGolemSplit            int `json:"npc_dota_neutral_mud_golem_split"`
			NpcDotaHeroMirana                      int `json:"npc_dota_hero_mirana"`
			NpcDotaHeroChen                        int `json:"npc_dota_hero_chen"`
			NpcDotaSentryWards                     int `json:"npc_dota_sentry_wards"`
			NpcDotaNeutralSatyrTrickster           int `json:"npc_dota_neutral_satyr_trickster"`
			NpcDotaNeutralSatyrSoulstealer         int `json:"npc_dota_neutral_satyr_soulstealer"`
			NpcDotaNeutralSatyrHellcaller          int `json:"npc_dota_neutral_satyr_hellcaller"`
			NpcDotaNeutralBlackDrake               int `json:"npc_dota_neutral_black_drake"`
			NpcDotaNeutralSmallThunderLizard       int `json:"npc_dota_neutral_small_thunder_lizard"`
			NpcDotaNeutralEnragedWildkin           int `json:"npc_dota_neutral_enraged_wildkin"`
			NpcDotaNeutralWildkin                  int `json:"npc_dota_neutral_wildkin"`
			NpcDotaNeutralOgreMauler               int `json:"npc_dota_neutral_ogre_mauler"`
			NpcDotaNeutralOgreMagi                 int `json:"npc_dota_neutral_ogre_magi"`
			NpcDotaNeutralWarpineRaider            int `json:"npc_dota_neutral_warpine_raider"`
			NpcDotaHeroTrollWarlord                int `json:"npc_dota_hero_troll_warlord"`
			NpcDotaNeutralKobold                   int `json:"npc_dota_neutral_kobold"`
			NpcDotaHeroDragonKnight                int `json:"npc_dota_hero_dragon_knight"`
			NpcDotaNeutralCentaurOutrunner         int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaNeutralPolarFurbolgUrsaWarrior  int `json:"npc_dota_neutral_polar_furbolg_ursa_warrior"`
			NpcDotaNeutralPolarFurbolgChampion     int `json:"npc_dota_neutral_polar_furbolg_champion"`
			NpcDotaObserverWards                   int `json:"npc_dota_observer_wards"`
			NpcDotaNeutralGiantWolf                int `json:"npc_dota_neutral_giant_wolf"`
			NpcDotaNeutralAlphaWolf                int `json:"npc_dota_neutral_alpha_wolf"`
			NpcDotaNeutralCentaurKhan              int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaBadguysSiege                    int `json:"npc_dota_badguys_siege"`
			NpcDotaDarkTrollWarlordSkeletonWarrior int `json:"npc_dota_dark_troll_warlord_skeleton_warrior"`
			NpcDotaNeutralForestTrollBerserker     int `json:"npc_dota_neutral_forest_troll_berserker"`
			NpcDotaItemWraithPactTotem             int `json:"npc_dota_item_wraith_pact_totem"`
			NpcDotaNeutralFelBeast                 int `json:"npc_dota_neutral_fel_beast"`
			NpcDotaNeutralGhost                    int `json:"npc_dota_neutral_ghost"`
			NpcDotaCreepBadguysRangedUpgraded      int `json:"npc_dota_creep_badguys_ranged_upgraded"`
			NpcDotaCreepBadguysMeleeUpgraded       int `json:"npc_dota_creep_badguys_melee_upgraded"`
			NpcDotaCreepBadguysFlagbearerUpgraded  int `json:"npc_dota_creep_badguys_flagbearer_upgraded"`
		} `json:"killed,omitempty"`
		KilledBy0 struct {
			NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
		} `json:"killed_by,omitempty"`
		LanePos0 struct {
			Num74 struct {
				Num76 int `json:"76"`
			} `json:"74"`
			Num76 struct {
				Num78 int `json:"78"`
			} `json:"76"`
			Num84 struct {
				Num86 int `json:"86"`
				Num88 int `json:"88"`
			} `json:"84"`
			Num86 struct {
				Num88 int `json:"88"`
			} `json:"86"`
			Num88 struct {
				Num90 int `json:"90"`
			} `json:"88"`
			Num90 struct {
				Num92 int `json:"92"`
			} `json:"90"`
			Num92 struct {
				Num94 int `json:"94"`
			} `json:"92"`
			Num94 struct {
				Num94 int `json:"94"`
			} `json:"94"`
			Num96 struct {
				Num96 int `json:"96"`
			} `json:"96"`
			Num98 struct {
				Num98 int `json:"98"`
			} `json:"98"`
			Num100 struct {
				Num98 int `json:"98"`
			} `json:"100"`
			Num102 struct {
				Num98 int `json:"98"`
			} `json:"102"`
			Num106 struct {
				Num108 int `json:"108"`
			} `json:"106"`
			Num108 struct {
				Num110 int `json:"110"`
			} `json:"108"`
			Num110 struct {
				Num110 int `json:"110"`
				Num112 int `json:"112"`
				Num114 int `json:"114"`
			} `json:"110"`
			Num112 struct {
				Num110 int `json:"110"`
				Num114 int `json:"114"`
				Num116 int `json:"116"`
				Num118 int `json:"118"`
			} `json:"112"`
			Num114 struct {
				Num104 int `json:"104"`
				Num110 int `json:"110"`
				Num116 int `json:"116"`
				Num118 int `json:"118"`
				Num120 int `json:"120"`
				Num122 int `json:"122"`
				Num126 int `json:"126"`
				Num128 int `json:"128"`
				Num130 int `json:"130"`
				Num132 int `json:"132"`
				Num134 int `json:"134"`
				Num136 int `json:"136"`
			} `json:"114"`
			Num116 struct {
				Num104 int `json:"104"`
				Num110 int `json:"110"`
				Num118 int `json:"118"`
				Num120 int `json:"120"`
				Num122 int `json:"122"`
				Num124 int `json:"124"`
				Num132 int `json:"132"`
				Num134 int `json:"134"`
			} `json:"116"`
			Num118 struct {
				Num110 int `json:"110"`
				Num118 int `json:"118"`
				Num120 int `json:"120"`
				Num122 int `json:"122"`
				Num130 int `json:"130"`
				Num132 int `json:"132"`
				Num136 int `json:"136"`
			} `json:"118"`
			Num120 struct {
				Num106 int `json:"106"`
				Num110 int `json:"110"`
				Num118 int `json:"118"`
				Num120 int `json:"120"`
				Num122 int `json:"122"`
				Num124 int `json:"124"`
				Num126 int `json:"126"`
				Num128 int `json:"128"`
				Num138 int `json:"138"`
			} `json:"120"`
			Num122 struct {
				Num106 int `json:"106"`
				Num112 int `json:"112"`
				Num118 int `json:"118"`
				Num120 int `json:"120"`
				Num122 int `json:"122"`
				Num124 int `json:"124"`
				Num126 int `json:"126"`
				Num138 int `json:"138"`
			} `json:"122"`
			Num124 struct {
				Num108 int `json:"108"`
				Num114 int `json:"114"`
				Num116 int `json:"116"`
				Num118 int `json:"118"`
				Num122 int `json:"122"`
				Num124 int `json:"124"`
				Num126 int `json:"126"`
				Num128 int `json:"128"`
			} `json:"124"`
			Num126 struct {
				Num110 int `json:"110"`
				Num112 int `json:"112"`
				Num114 int `json:"114"`
				Num116 int `json:"116"`
				Num120 int `json:"120"`
				Num122 int `json:"122"`
				Num124 int `json:"124"`
				Num126 int `json:"126"`
				Num128 int `json:"128"`
				Num130 int `json:"130"`
				Num136 int `json:"136"`
				Num140 int `json:"140"`
			} `json:"126"`
			Num128 struct {
				Num108 int `json:"108"`
				Num110 int `json:"110"`
				Num114 int `json:"114"`
				Num116 int `json:"116"`
				Num120 int `json:"120"`
				Num122 int `json:"122"`
				Num126 int `json:"126"`
				Num128 int `json:"128"`
				Num130 int `json:"130"`
				Num138 int `json:"138"`
				Num140 int `json:"140"`
				Num142 int `json:"142"`
				Num144 int `json:"144"`
				Num146 int `json:"146"`
				Num148 int `json:"148"`
			} `json:"128"`
			Num130 struct {
				Num106 int `json:"106"`
				Num112 int `json:"112"`
				Num114 int `json:"114"`
				Num120 int `json:"120"`
				Num122 int `json:"122"`
				Num124 int `json:"124"`
				Num126 int `json:"126"`
				Num128 int `json:"128"`
				Num130 int `json:"130"`
				Num148 int `json:"148"`
				Num150 int `json:"150"`
			} `json:"130"`
			Num132 struct {
				Num104 int `json:"104"`
				Num114 int `json:"114"`
				Num116 int `json:"116"`
				Num118 int `json:"118"`
				Num120 int `json:"120"`
				Num124 int `json:"124"`
				Num126 int `json:"126"`
				Num128 int `json:"128"`
			} `json:"132"`
			Num134 struct {
				Num104 int `json:"104"`
				Num114 int `json:"114"`
				Num116 int `json:"116"`
				Num118 int `json:"118"`
				Num120 int `json:"120"`
				Num122 int `json:"122"`
				Num124 int `json:"124"`
				Num126 int `json:"126"`
				Num130 int `json:"130"`
				Num132 int `json:"132"`
			} `json:"134"`
			Num136 struct {
				Num102 int `json:"102"`
				Num116 int `json:"116"`
				Num118 int `json:"118"`
				Num122 int `json:"122"`
				Num130 int `json:"130"`
				Num132 int `json:"132"`
				Num134 int `json:"134"`
			} `json:"136"`
			Num138 struct {
				Num102 int `json:"102"`
				Num114 int `json:"114"`
				Num116 int `json:"116"`
				Num132 int `json:"132"`
				Num134 int `json:"134"`
			} `json:"138"`
			Num140 struct {
				Num116 int `json:"116"`
				Num132 int `json:"132"`
				Num134 int `json:"134"`
			} `json:"140"`
			Num142 struct {
				Num100 int `json:"100"`
				Num118 int `json:"118"`
				Num132 int `json:"132"`
			} `json:"142"`
			Num144 struct {
				Num100 int `json:"100"`
				Num132 int `json:"132"`
			} `json:"144"`
			Num146 struct {
				Num100 int `json:"100"`
				Num130 int `json:"130"`
				Num132 int `json:"132"`
			} `json:"146"`
			Num150 struct {
				Num100 int `json:"100"`
			} `json:"150"`
			Num152 struct {
				Num100 int `json:"100"`
			} `json:"152"`
			Num164 struct {
				Num98 int `json:"98"`
			} `json:"164"`
			Num166 struct {
				Num86 int `json:"86"`
				Num96 int `json:"96"`
			} `json:"166"`
			Num168 struct {
				Num88 int `json:"88"`
				Num90 int `json:"90"`
				Num92 int `json:"92"`
				Num94 int `json:"94"`
			} `json:"168"`
			Num170 struct {
				Num88 int `json:"88"`
				Num94 int `json:"94"`
			} `json:"170"`
			Num172 struct {
				Num88 int `json:"88"`
				Num90 int `json:"90"`
				Num94 int `json:"94"`
			} `json:"172"`
			Num174 struct {
				Num90 int `json:"90"`
				Num92 int `json:"92"`
			} `json:"174"`
		} `json:"lane_pos,omitempty"`
		MultiKills0 struct {
			Num2 int `json:"2"`
		} `json:"multi_kills,omitempty"`
		Obs0 struct {
			Num124 struct {
				Num118 int `json:"118"`
			} `json:"124"`
			Num144 struct {
				Num120 int `json:"120"`
			} `json:"144"`
		} `json:"obs,omitempty"`
		Purchase0 struct {
			WardObserver            int `json:"ward_observer"`
			Branches                int `json:"branches"`
			FaerieFire              int `json:"faerie_fire"`
			Tango                   int `json:"tango"`
			Bottle                  int `json:"bottle"`
			MagicStick              int `json:"magic_stick"`
			RecipeMagicWand         int `json:"recipe_magic_wand"`
			MagicWand               int `json:"magic_wand"`
			Robe                    int `json:"robe"`
			InfusedRaindrop         int `json:"infused_raindrop"`
			WardSentry              int `json:"ward_sentry"`
			WardDispenser           int `json:"ward_dispenser"`
			Boots                   int `json:"boots"`
			Gloves                  int `json:"gloves"`
			PowerTreads             int `json:"power_treads"`
			Tpscroll                int `json:"tpscroll"`
			Clarity                 int `json:"clarity"`
			BlitzKnuckles           int `json:"blitz_knuckles"`
			Chainmail               int `json:"chainmail"`
			RecipeWitchBlade        int `json:"recipe_witch_blade"`
			WitchBlade              int `json:"witch_blade"`
			Blink                   int `json:"blink"`
			RecipeSphere            int `json:"recipe_sphere"`
			UltimateOrb             int `json:"ultimate_orb"`
			RingOfHealth            int `json:"ring_of_health"`
			Pers                    int `json:"pers"`
			VoidStone               int `json:"void_stone"`
			Sphere                  int `json:"sphere"`
			StaffOfWizardry         int `json:"staff_of_wizardry"`
			Kaya                    int `json:"kaya"`
			RecipeKaya              int `json:"recipe_kaya"`
			OgreAxe                 int `json:"ogre_axe"`
			BeltOfStrength          int `json:"belt_of_strength"`
			Sange                   int `json:"sange"`
			RecipeSange             int `json:"recipe_sange"`
			KayaAndSange            int `json:"kaya_and_sange"`
			RecipeTravelBoots       int `json:"recipe_travel_boots"`
			RecipeAetherLens        int `json:"recipe_aether_lens"`
			VitalityBooster         int `json:"vitality_booster"`
			EnergyBooster           int `json:"energy_booster"`
			SoulBooster             int `json:"soul_booster"`
			PointBooster            int `json:"point_booster"`
			AetherLens              int `json:"aether_lens"`
			OctarineCore            int `json:"octarine_core"`
			RecipeOverwhelmingBlink int `json:"recipe_overwhelming_blink"`
			Reaver                  int `json:"reaver"`
			OverwhelmingBlink       int `json:"overwhelming_blink"`
		} `json:"purchase,omitempty"`
		Runes0 struct {
			Num0 int `json:"0"`
			Num1 int `json:"1"`
			Num3 int `json:"3"`
			Num4 int `json:"4"`
			Num5 int `json:"5"`
			Num6 int `json:"6"`
			Num7 int `json:"7"`
		} `json:"runes,omitempty"`
		Sen0 struct {
			Num122 struct {
				Num124 int `json:"124"`
			} `json:"122"`
		} `json:"sen,omitempty"`
		XpReasons0 struct {
			Num0 int `json:"0"`
			Num1 int `json:"1"`
			Num2 int `json:"2"`
			Num3 int `json:"3"`
			Num5 int `json:"5"`
		} `json:"xp_reasons,omitempty"`
		PurchaseTime0 struct {
			WardObserver      int `json:"ward_observer"`
			Branches          int `json:"branches"`
			FaerieFire        int `json:"faerie_fire"`
			Tango             int `json:"tango"`
			Bottle            int `json:"bottle"`
			MagicStick        int `json:"magic_stick"`
			MagicWand         int `json:"magic_wand"`
			Robe              int `json:"robe"`
			InfusedRaindrop   int `json:"infused_raindrop"`
			WardSentry        int `json:"ward_sentry"`
			Boots             int `json:"boots"`
			Gloves            int `json:"gloves"`
			PowerTreads       int `json:"power_treads"`
			Tpscroll          int `json:"tpscroll"`
			Clarity           int `json:"clarity"`
			BlitzKnuckles     int `json:"blitz_knuckles"`
			Chainmail         int `json:"chainmail"`
			WitchBlade        int `json:"witch_blade"`
			Blink             int `json:"blink"`
			UltimateOrb       int `json:"ultimate_orb"`
			RingOfHealth      int `json:"ring_of_health"`
			Pers              int `json:"pers"`
			VoidStone         int `json:"void_stone"`
			Sphere            int `json:"sphere"`
			StaffOfWizardry   int `json:"staff_of_wizardry"`
			Kaya              int `json:"kaya"`
			OgreAxe           int `json:"ogre_axe"`
			BeltOfStrength    int `json:"belt_of_strength"`
			Sange             int `json:"sange"`
			KayaAndSange      int `json:"kaya_and_sange"`
			VitalityBooster   int `json:"vitality_booster"`
			EnergyBooster     int `json:"energy_booster"`
			SoulBooster       int `json:"soul_booster"`
			PointBooster      int `json:"point_booster"`
			AetherLens        int `json:"aether_lens"`
			OctarineCore      int `json:"octarine_core"`
			Reaver            int `json:"reaver"`
			OverwhelmingBlink int `json:"overwhelming_blink"`
		} `json:"purchase_time,omitempty"`
		FirstPurchaseTime0 struct {
			WardObserver      int `json:"ward_observer"`
			Branches          int `json:"branches"`
			FaerieFire        int `json:"faerie_fire"`
			Tango             int `json:"tango"`
			Bottle            int `json:"bottle"`
			MagicStick        int `json:"magic_stick"`
			MagicWand         int `json:"magic_wand"`
			Robe              int `json:"robe"`
			InfusedRaindrop   int `json:"infused_raindrop"`
			WardSentry        int `json:"ward_sentry"`
			Boots             int `json:"boots"`
			Gloves            int `json:"gloves"`
			PowerTreads       int `json:"power_treads"`
			Tpscroll          int `json:"tpscroll"`
			Clarity           int `json:"clarity"`
			BlitzKnuckles     int `json:"blitz_knuckles"`
			Chainmail         int `json:"chainmail"`
			WitchBlade        int `json:"witch_blade"`
			Blink             int `json:"blink"`
			UltimateOrb       int `json:"ultimate_orb"`
			RingOfHealth      int `json:"ring_of_health"`
			Pers              int `json:"pers"`
			VoidStone         int `json:"void_stone"`
			Sphere            int `json:"sphere"`
			StaffOfWizardry   int `json:"staff_of_wizardry"`
			Kaya              int `json:"kaya"`
			OgreAxe           int `json:"ogre_axe"`
			BeltOfStrength    int `json:"belt_of_strength"`
			Sange             int `json:"sange"`
			KayaAndSange      int `json:"kaya_and_sange"`
			VitalityBooster   int `json:"vitality_booster"`
			EnergyBooster     int `json:"energy_booster"`
			SoulBooster       int `json:"soul_booster"`
			PointBooster      int `json:"point_booster"`
			AetherLens        int `json:"aether_lens"`
			OctarineCore      int `json:"octarine_core"`
			Reaver            int `json:"reaver"`
			OverwhelmingBlink int `json:"overwhelming_blink"`
		} `json:"first_purchase_time,omitempty"`
		ItemWin0 struct {
			WardObserver      int `json:"ward_observer"`
			Branches          int `json:"branches"`
			FaerieFire        int `json:"faerie_fire"`
			Tango             int `json:"tango"`
			Bottle            int `json:"bottle"`
			MagicStick        int `json:"magic_stick"`
			MagicWand         int `json:"magic_wand"`
			Robe              int `json:"robe"`
			InfusedRaindrop   int `json:"infused_raindrop"`
			WardSentry        int `json:"ward_sentry"`
			Boots             int `json:"boots"`
			Gloves            int `json:"gloves"`
			PowerTreads       int `json:"power_treads"`
			Tpscroll          int `json:"tpscroll"`
			Clarity           int `json:"clarity"`
			BlitzKnuckles     int `json:"blitz_knuckles"`
			Chainmail         int `json:"chainmail"`
			WitchBlade        int `json:"witch_blade"`
			Blink             int `json:"blink"`
			UltimateOrb       int `json:"ultimate_orb"`
			RingOfHealth      int `json:"ring_of_health"`
			Pers              int `json:"pers"`
			VoidStone         int `json:"void_stone"`
			Sphere            int `json:"sphere"`
			StaffOfWizardry   int `json:"staff_of_wizardry"`
			Kaya              int `json:"kaya"`
			OgreAxe           int `json:"ogre_axe"`
			BeltOfStrength    int `json:"belt_of_strength"`
			Sange             int `json:"sange"`
			KayaAndSange      int `json:"kaya_and_sange"`
			VitalityBooster   int `json:"vitality_booster"`
			EnergyBooster     int `json:"energy_booster"`
			SoulBooster       int `json:"soul_booster"`
			PointBooster      int `json:"point_booster"`
			AetherLens        int `json:"aether_lens"`
			OctarineCore      int `json:"octarine_core"`
			Reaver            int `json:"reaver"`
			OverwhelmingBlink int `json:"overwhelming_blink"`
		} `json:"item_win,omitempty"`
		ItemUsage0 struct {
			WardObserver      int `json:"ward_observer"`
			Branches          int `json:"branches"`
			FaerieFire        int `json:"faerie_fire"`
			Tango             int `json:"tango"`
			Bottle            int `json:"bottle"`
			MagicStick        int `json:"magic_stick"`
			MagicWand         int `json:"magic_wand"`
			Robe              int `json:"robe"`
			InfusedRaindrop   int `json:"infused_raindrop"`
			WardSentry        int `json:"ward_sentry"`
			Boots             int `json:"boots"`
			Gloves            int `json:"gloves"`
			PowerTreads       int `json:"power_treads"`
			Tpscroll          int `json:"tpscroll"`
			Clarity           int `json:"clarity"`
			BlitzKnuckles     int `json:"blitz_knuckles"`
			Chainmail         int `json:"chainmail"`
			WitchBlade        int `json:"witch_blade"`
			Blink             int `json:"blink"`
			UltimateOrb       int `json:"ultimate_orb"`
			RingOfHealth      int `json:"ring_of_health"`
			Pers              int `json:"pers"`
			VoidStone         int `json:"void_stone"`
			Sphere            int `json:"sphere"`
			StaffOfWizardry   int `json:"staff_of_wizardry"`
			Kaya              int `json:"kaya"`
			OgreAxe           int `json:"ogre_axe"`
			BeltOfStrength    int `json:"belt_of_strength"`
			Sange             int `json:"sange"`
			KayaAndSange      int `json:"kaya_and_sange"`
			VitalityBooster   int `json:"vitality_booster"`
			EnergyBooster     int `json:"energy_booster"`
			SoulBooster       int `json:"soul_booster"`
			PointBooster      int `json:"point_booster"`
			AetherLens        int `json:"aether_lens"`
			OctarineCore      int `json:"octarine_core"`
			Reaver            int `json:"reaver"`
			OverwhelmingBlink int `json:"overwhelming_blink"`
		} `json:"item_usage,omitempty"`
		AbilityTargets1 struct {
			EnchantressEnchant struct {
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
				NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
				NpcDotaHeroChen         int `json:"npc_dota_hero_chen"`
			} `json:"enchantress_enchant"`
		} `json:"ability_targets,omitempty"`
		AbilityUses1 struct {
			EnchantressEnchant           int `json:"enchantress_enchant"`
			EnchantressNaturesAttendants int `json:"enchantress_natures_attendants"`
			AbilityCapture               int `json:"ability_capture"`
		} `json:"ability_uses,omitempty"`
		Actions1 struct {
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
			Num12 int `json:"12"`
			Num13 int `json:"13"`
			Num14 int `json:"14"`
			Num15 int `json:"15"`
			Num16 int `json:"16"`
			Num17 int `json:"17"`
			Num19 int `json:"19"`
			Num23 int `json:"23"`
			Num24 int `json:"24"`
			Num25 int `json:"25"`
			Num30 int `json:"30"`
			Num33 int `json:"33"`
			Num37 int `json:"37"`
			Num38 int `json:"38"`
		} `json:"actions,omitempty"`
		Damage1 struct {
			NpcDotaHeroMirana                     int `json:"npc_dota_hero_mirana"`
			NpcDotaHeroDragonKnight               int `json:"npc_dota_hero_dragon_knight"`
			NpcDotaCreepGoodguysMelee             int `json:"npc_dota_creep_goodguys_melee"`
			NpcDotaCreepGoodguysRanged            int `json:"npc_dota_creep_goodguys_ranged"`
			NpcDotaSentryWards                    int `json:"npc_dota_sentry_wards"`
			NpcDotaObserverWards                  int `json:"npc_dota_observer_wards"`
			NpcDotaNeutralCentaurOutrunner        int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaCreepGoodguysFlagbearer        int `json:"npc_dota_creep_goodguys_flagbearer"`
			NpcDotaNeutralDarkTroll               int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaNeutralGnollAssassin           int `json:"npc_dota_neutral_gnoll_assassin"`
			NpcDotaHeroQueenofpain                int `json:"npc_dota_hero_queenofpain"`
			NpcDotaBadguysTower1Mid               int `json:"npc_dota_badguys_tower1_mid"`
			NpcDotaNeutralGiantWolf               int `json:"npc_dota_neutral_giant_wolf"`
			NpcDotaCourier                        int `json:"npc_dota_courier"`
			NpcDotaNeutralOgreMagi                int `json:"npc_dota_neutral_ogre_magi"`
			NpcDotaBadguysTower1Top               int `json:"npc_dota_badguys_tower1_top"`
			NpcDotaNeutralSatyrSoulstealer        int `json:"npc_dota_neutral_satyr_soulstealer"`
			NpcDotaNeutralCentaurKhan             int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaCreepBadguysRanged             int `json:"npc_dota_creep_badguys_ranged"`
			NpcDotaNeutralPolarFurbolgChampion    int `json:"npc_dota_neutral_polar_furbolg_champion"`
			NpcDotaNeutralIceShaman               int `json:"npc_dota_neutral_ice_shaman"`
			NpcDotaNeutralWildkin                 int `json:"npc_dota_neutral_wildkin"`
			NpcDotaNeutralAlphaWolf               int `json:"npc_dota_neutral_alpha_wolf"`
			NpcDotaNeutralDarkTrollWarlord        int `json:"npc_dota_neutral_dark_troll_warlord"`
			NpcDotaNeutralWarpineRaider           int `json:"npc_dota_neutral_warpine_raider"`
			NpcDotaHeroChen                       int `json:"npc_dota_hero_chen"`
			NpcDotaNeutralEnragedWildkin          int `json:"npc_dota_neutral_enraged_wildkin"`
			NpcDotaCreepBadguysMelee              int `json:"npc_dota_creep_badguys_melee"`
			NpcDotaNeutralKobold                  int `json:"npc_dota_neutral_kobold"`
			NpcDotaNeutralPolarFurbolgUrsaWarrior int `json:"npc_dota_neutral_polar_furbolg_ursa_warrior"`
			NpcDotaCreepBadguysFlagbearer         int `json:"npc_dota_creep_badguys_flagbearer"`
			NpcDotaRoshan                         int `json:"npc_dota_roshan"`
			NpcDotaBadguysTower2Mid               int `json:"npc_dota_badguys_tower2_mid"`
			NpcDotaHeroTrollWarlord               int `json:"npc_dota_hero_troll_warlord"`
			NpcDotaNeutralSatyrTrickster          int `json:"npc_dota_neutral_satyr_trickster"`
			NpcDotaNeutralMudGolem                int `json:"npc_dota_neutral_mud_golem"`
			NpcDotaBadguysTower2Top               int `json:"npc_dota_badguys_tower2_top"`
			NpcDotaBadguysTower3Top               int `json:"npc_dota_badguys_tower3_top"`
			NpcDotaCreepBadguysMeleeUpgraded      int `json:"npc_dota_creep_badguys_melee_upgraded"`
			NpcDotaCreepBadguysRangedUpgraded     int `json:"npc_dota_creep_badguys_ranged_upgraded"`
			NpcDotaCreepBadguysFlagbearerUpgraded int `json:"npc_dota_creep_badguys_flagbearer_upgraded"`
		} `json:"damage,omitempty"`
		DamageInflictor1 struct {
			Null                               int `json:"null"`
			CentaurKhanWarStomp                int `json:"centaur_khan_war_stomp"`
			EnchantressImpetus                 int `json:"enchantress_impetus"`
			PolarFurbolgUrsaWarriorThunderClap int `json:"polar_furbolg_ursa_warrior_thunder_clap"`
			WarpineRaiderSeedShot              int `json:"warpine_raider_seed_shot"`
			Dust                               int `json:"dust"`
			SatyrHellcallerShockwave           int `json:"satyr_hellcaller_shockwave"`
		} `json:"damage_inflictor,omitempty"`
		DamageInflictorReceived1 struct {
			Null                           int `json:"null"`
			DragonKnightBreatheFire        int `json:"dragon_knight_breathe_fire"`
			WarpineRaiderSeedShot          int `json:"warpine_raider_seed_shot"`
			QueenofpainScreamOfPain        int `json:"queenofpain_scream_of_pain"`
			QueenofpainSonicWave           int `json:"queenofpain_sonic_wave"`
			DragonKnightDragonTail         int `json:"dragon_knight_dragon_tail"`
			MiranaArrow                    int `json:"mirana_arrow"`
			MiranaStarfall                 int `json:"mirana_starfall"`
			QueenofpainShadowStrike        int `json:"queenofpain_shadow_strike"`
			UrnOfShadows                   int `json:"urn_of_shadows"`
			Bfury                          int `json:"bfury"`
			TrollWarlordWhirlingAxesMelee  int `json:"troll_warlord_whirling_axes_melee"`
			QueenofpainBlink               int `json:"queenofpain_blink"`
			Orchid                         int `json:"orchid"`
			TrollWarlordWhirlingAxesRanged int `json:"troll_warlord_whirling_axes_ranged"`
			WraithPact                     int `json:"wraith_pact"`
			DragonKnightFireball           int `json:"dragon_knight_fireball"`
			DragonKnightElderDragonForm    int `json:"dragon_knight_elder_dragon_form"`
			Stormcrafter                   int `json:"stormcrafter"`
			ShivasGuard                    int `json:"shivas_guard"`
			SpiritVessel                   int `json:"spirit_vessel"`
			Bloodthorn                     int `json:"bloodthorn"`
		} `json:"damage_inflictor_received,omitempty"`
		DamageTaken1 struct {
			NpcDotaHeroMirana                  int `json:"npc_dota_hero_mirana"`
			NpcDotaCreepBadguysMelee           int `json:"npc_dota_creep_badguys_melee"`
			NpcDotaCreepBadguysRanged          int `json:"npc_dota_creep_badguys_ranged"`
			NpcDotaNeutralDarkTroll            int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaNeutralCentaurOutrunner     int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaBadguysTower1Top            int `json:"npc_dota_badguys_tower1_top"`
			NpcDotaHeroDragonKnight            int `json:"npc_dota_hero_dragon_knight"`
			NpcDotaHeroChen                    int `json:"npc_dota_hero_chen"`
			NpcDotaNeutralIceShaman            int `json:"npc_dota_neutral_ice_shaman"`
			NpcDotaNeutralAlphaWolf            int `json:"npc_dota_neutral_alpha_wolf"`
			NpcDotaNeutralGiantWolf            int `json:"npc_dota_neutral_giant_wolf"`
			NpcDotaHeroQueenofpain             int `json:"npc_dota_hero_queenofpain"`
			NpcDotaNeutralKobold               int `json:"npc_dota_neutral_kobold"`
			NpcDotaNeutralPolarFurbolgChampion int `json:"npc_dota_neutral_polar_furbolg_champion"`
			NpcDotaNeutralCentaurKhan          int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaRoshan                      int `json:"npc_dota_roshan"`
			NpcDotaHeroTrollWarlord            int `json:"npc_dota_hero_troll_warlord"`
			NpcDotaCreepBadguysRangedUpgraded  int `json:"npc_dota_creep_badguys_ranged_upgraded"`
			NpcDotaCreepBadguysMeleeUpgraded   int `json:"npc_dota_creep_badguys_melee_upgraded"`
		} `json:"damage_taken,omitempty"`
		DamageTargets1 struct {
			Null struct {
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
				NpcDotaHeroChen         int `json:"npc_dota_hero_chen"`
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
			} `json:"null"`
			CentaurKhanWarStomp struct {
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
			} `json:"centaur_khan_war_stomp"`
			EnchantressImpetus struct {
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
				NpcDotaHeroChen         int `json:"npc_dota_hero_chen"`
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
			} `json:"enchantress_impetus"`
			PolarFurbolgUrsaWarriorThunderClap struct {
				NpcDotaHeroQueenofpain int `json:"npc_dota_hero_queenofpain"`
			} `json:"polar_furbolg_ursa_warrior_thunder_clap"`
			WarpineRaiderSeedShot struct {
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
			} `json:"warpine_raider_seed_shot"`
			Dust struct {
				NpcDotaHeroMirana int `json:"npc_dota_hero_mirana"`
				NpcDotaHeroChen   int `json:"npc_dota_hero_chen"`
			} `json:"dust"`
			SatyrHellcallerShockwave struct {
				NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
				NpcDotaHeroChen         int `json:"npc_dota_hero_chen"`
			} `json:"satyr_hellcaller_shockwave"`
		} `json:"damage_targets,omitempty"`
		GoldReasons1 struct {
			Num0  int `json:"0"`
			Num1  int `json:"1"`
			Num6  int `json:"6"`
			Num11 int `json:"11"`
			Num12 int `json:"12"`
			Num13 int `json:"13"`
			Num14 int `json:"14"`
			Num15 int `json:"15"`
			Num16 int `json:"16"`
			Num17 int `json:"17"`
			Num19 int `json:"19"`
			Num20 int `json:"20"`
			Num21 int `json:"21"`
		} `json:"gold_reasons,omitempty"`
		HeroHits1 struct {
			Null                               int `json:"null"`
			CentaurKhanWarStomp                int `json:"centaur_khan_war_stomp"`
			EnchantressImpetus                 int `json:"enchantress_impetus"`
			PolarFurbolgUrsaWarriorThunderClap int `json:"polar_furbolg_ursa_warrior_thunder_clap"`
			WarpineRaiderSeedShot              int `json:"warpine_raider_seed_shot"`
			Dust                               int `json:"dust"`
			SatyrHellcallerShockwave           int `json:"satyr_hellcaller_shockwave"`
		} `json:"hero_hits,omitempty"`
		ItemUses1 struct {
			SmokeOfDeceit      int `json:"smoke_of_deceit"`
			Tango              int `json:"tango"`
			WardSentry         int `json:"ward_sentry"`
			WardDispenser      int `json:"ward_dispenser"`
			WardObserver       int `json:"ward_observer"`
			MagicWand          int `json:"magic_wand"`
			Tpscroll           int `json:"tpscroll"`
			TrustyShovel       int `json:"trusty_shovel"`
			EnchantedMango     int `json:"enchanted_mango"`
			Dust               int `json:"dust"`
			ForceStaff         int `json:"force_staff"`
			Flask              int `json:"flask"`
			FaerieFire         int `json:"faerie_fire"`
			Bullwhip           int `json:"bullwhip"`
			MedallionOfCourage int `json:"medallion_of_courage"`
			HeavensHalberd     int `json:"heavens_halberd"`
			TomeOfKnowledge    int `json:"tome_of_knowledge"`
		} `json:"item_uses,omitempty"`
		Killed1 struct {
			NpcDotaCreepGoodguysMelee             int `json:"npc_dota_creep_goodguys_melee"`
			NpcDotaSentryWards                    int `json:"npc_dota_sentry_wards"`
			NpcDotaObserverWards                  int `json:"npc_dota_observer_wards"`
			NpcDotaNeutralGnollAssassin           int `json:"npc_dota_neutral_gnoll_assassin"`
			NpcDotaNeutralGiantWolf               int `json:"npc_dota_neutral_giant_wolf"`
			NpcDotaNeutralDarkTroll               int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaCourier                        int `json:"npc_dota_courier"`
			NpcDotaNeutralCentaurOutrunner        int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaNeutralSatyrSoulstealer        int `json:"npc_dota_neutral_satyr_soulstealer"`
			NpcDotaNeutralWildkin                 int `json:"npc_dota_neutral_wildkin"`
			NpcDotaNeutralAlphaWolf               int `json:"npc_dota_neutral_alpha_wolf"`
			NpcDotaNeutralPolarFurbolgChampion    int `json:"npc_dota_neutral_polar_furbolg_champion"`
			NpcDotaCreepBadguysRanged             int `json:"npc_dota_creep_badguys_ranged"`
			NpcDotaCreepBadguysMelee              int `json:"npc_dota_creep_badguys_melee"`
			NpcDotaNeutralKobold                  int `json:"npc_dota_neutral_kobold"`
			NpcDotaNeutralWarpineRaider           int `json:"npc_dota_neutral_warpine_raider"`
			NpcDotaCreepBadguysFlagbearer         int `json:"npc_dota_creep_badguys_flagbearer"`
			NpcDotaCreepGoodguysRanged            int `json:"npc_dota_creep_goodguys_ranged"`
			NpcDotaNeutralSatyrTrickster          int `json:"npc_dota_neutral_satyr_trickster"`
			NpcDotaCreepBadguysMeleeUpgraded      int `json:"npc_dota_creep_badguys_melee_upgraded"`
			NpcDotaCreepBadguysRangedUpgraded     int `json:"npc_dota_creep_badguys_ranged_upgraded"`
			NpcDotaCreepBadguysFlagbearerUpgraded int `json:"npc_dota_creep_badguys_flagbearer_upgraded"`
		} `json:"killed,omitempty"`
		LanePos1 struct {
			Num74 struct {
				Num76 int `json:"76"`
				Num78 int `json:"78"`
			} `json:"74"`
			Num76 struct {
				Num78 int `json:"78"`
			} `json:"76"`
			Num78 struct {
				Num80 int `json:"80"`
			} `json:"78"`
			Num80 struct {
				Num82 int `json:"82"`
			} `json:"80"`
			Num82 struct {
				Num84 int `json:"84"`
			} `json:"82"`
			Num84 struct {
				Num86 int `json:"86"`
				Num88 int `json:"88"`
			} `json:"84"`
			Num86 struct {
				Num90 int `json:"90"`
			} `json:"86"`
			Num88 struct {
				Num92 int `json:"92"`
			} `json:"88"`
			Num90 struct {
				Num94  int `json:"94"`
				Num132 int `json:"132"`
				Num136 int `json:"136"`
				Num138 int `json:"138"`
			} `json:"90"`
			Num92 struct {
				Num94  int `json:"94"`
				Num128 int `json:"128"`
				Num130 int `json:"130"`
				Num140 int `json:"140"`
				Num182 int `json:"182"`
			} `json:"92"`
			Num94 struct {
				Num96  int `json:"96"`
				Num130 int `json:"130"`
				Num176 int `json:"176"`
				Num178 int `json:"178"`
				Num180 int `json:"180"`
				Num182 int `json:"182"`
			} `json:"94"`
			Num96 struct {
				Num98  int `json:"98"`
				Num130 int `json:"130"`
				Num166 int `json:"166"`
				Num178 int `json:"178"`
			} `json:"96"`
			Num98 struct {
				Num100 int `json:"100"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
				Num176 int `json:"176"`
			} `json:"98"`
			Num100 struct {
				Num132 int `json:"132"`
				Num164 int `json:"164"`
				Num168 int `json:"168"`
				Num174 int `json:"174"`
			} `json:"100"`
			Num102 struct {
				Num102 int `json:"102"`
				Num134 int `json:"134"`
				Num162 int `json:"162"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
				Num174 int `json:"174"`
			} `json:"102"`
			Num104 struct {
				Num104 int `json:"104"`
				Num136 int `json:"136"`
				Num160 int `json:"160"`
			} `json:"104"`
			Num106 struct {
				Num106 int `json:"106"`
				Num136 int `json:"136"`
				Num160 int `json:"160"`
				Num166 int `json:"166"`
			} `json:"106"`
			Num108 struct {
				Num108 int `json:"108"`
				Num138 int `json:"138"`
				Num158 int `json:"158"`
				Num168 int `json:"168"`
			} `json:"108"`
			Num110 struct {
				Num106 int `json:"106"`
				Num140 int `json:"140"`
				Num158 int `json:"158"`
				Num168 int `json:"168"`
			} `json:"110"`
			Num112 struct {
				Num108 int `json:"108"`
				Num142 int `json:"142"`
				Num158 int `json:"158"`
			} `json:"112"`
			Num114 struct {
				Num108 int `json:"108"`
				Num114 int `json:"114"`
				Num118 int `json:"118"`
				Num144 int `json:"144"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
			} `json:"114"`
			Num116 struct {
				Num114 int `json:"114"`
				Num116 int `json:"116"`
				Num118 int `json:"118"`
				Num134 int `json:"134"`
				Num138 int `json:"138"`
				Num140 int `json:"140"`
				Num144 int `json:"144"`
				Num146 int `json:"146"`
				Num148 int `json:"148"`
				Num156 int `json:"156"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
			} `json:"116"`
			Num118 struct {
				Num110 int `json:"110"`
				Num112 int `json:"112"`
				Num114 int `json:"114"`
				Num116 int `json:"116"`
				Num130 int `json:"130"`
				Num132 int `json:"132"`
				Num148 int `json:"148"`
				Num150 int `json:"150"`
				Num154 int `json:"154"`
				Num156 int `json:"156"`
			} `json:"118"`
			Num120 struct {
				Num114 int `json:"114"`
				Num116 int `json:"116"`
				Num128 int `json:"128"`
				Num134 int `json:"134"`
				Num136 int `json:"136"`
				Num150 int `json:"150"`
				Num152 int `json:"152"`
				Num154 int `json:"154"`
			} `json:"120"`
			Num122 struct {
				Num116 int `json:"116"`
				Num134 int `json:"134"`
				Num136 int `json:"136"`
				Num138 int `json:"138"`
			} `json:"122"`
			Num124 struct {
				Num112 int `json:"112"`
				Num114 int `json:"114"`
				Num116 int `json:"116"`
				Num126 int `json:"126"`
				Num134 int `json:"134"`
				Num136 int `json:"136"`
				Num138 int `json:"138"`
				Num152 int `json:"152"`
			} `json:"124"`
			Num126 struct {
				Num112 int `json:"112"`
				Num114 int `json:"114"`
				Num126 int `json:"126"`
				Num134 int `json:"134"`
				Num136 int `json:"136"`
				Num140 int `json:"140"`
				Num148 int `json:"148"`
				Num150 int `json:"150"`
			} `json:"126"`
			Num128 struct {
				Num108 int `json:"108"`
				Num110 int `json:"110"`
				Num112 int `json:"112"`
				Num114 int `json:"114"`
				Num132 int `json:"132"`
				Num134 int `json:"134"`
				Num136 int `json:"136"`
				Num138 int `json:"138"`
				Num140 int `json:"140"`
				Num142 int `json:"142"`
				Num144 int `json:"144"`
				Num146 int `json:"146"`
				Num148 int `json:"148"`
			} `json:"128"`
			Num130 struct {
				Num106 int `json:"106"`
				Num112 int `json:"112"`
				Num114 int `json:"114"`
				Num126 int `json:"126"`
				Num138 int `json:"138"`
			} `json:"130"`
			Num132 struct {
				Num104 int `json:"104"`
				Num114 int `json:"114"`
				Num116 int `json:"116"`
				Num126 int `json:"126"`
				Num134 int `json:"134"`
				Num136 int `json:"136"`
			} `json:"132"`
			Num134 struct {
				Num104 int `json:"104"`
				Num114 int `json:"114"`
				Num116 int `json:"116"`
				Num118 int `json:"118"`
				Num124 int `json:"124"`
				Num134 int `json:"134"`
				Num138 int `json:"138"`
			} `json:"134"`
			Num136 struct {
				Num102 int `json:"102"`
				Num116 int `json:"116"`
				Num118 int `json:"118"`
				Num124 int `json:"124"`
				Num134 int `json:"134"`
				Num136 int `json:"136"`
			} `json:"136"`
			Num138 struct {
				Num100 int `json:"100"`
				Num102 int `json:"102"`
				Num114 int `json:"114"`
				Num116 int `json:"116"`
				Num126 int `json:"126"`
				Num134 int `json:"134"`
				Num136 int `json:"136"`
			} `json:"138"`
			Num140 struct {
				Num98  int `json:"98"`
				Num102 int `json:"102"`
				Num108 int `json:"108"`
				Num110 int `json:"110"`
				Num112 int `json:"112"`
				Num114 int `json:"114"`
				Num128 int `json:"128"`
				Num130 int `json:"130"`
				Num132 int `json:"132"`
				Num134 int `json:"134"`
			} `json:"140"`
			Num142 struct {
				Num90  int `json:"90"`
				Num92  int `json:"92"`
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num102 int `json:"102"`
				Num104 int `json:"104"`
				Num106 int `json:"106"`
				Num108 int `json:"108"`
				Num110 int `json:"110"`
				Num112 int `json:"112"`
				Num132 int `json:"132"`
			} `json:"142"`
			Num144 struct {
				Num90  int `json:"90"`
				Num94  int `json:"94"`
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num102 int `json:"102"`
				Num104 int `json:"104"`
				Num106 int `json:"106"`
				Num110 int `json:"110"`
				Num112 int `json:"112"`
				Num114 int `json:"114"`
				Num128 int `json:"128"`
			} `json:"144"`
			Num146 struct {
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num112 int `json:"112"`
				Num114 int `json:"114"`
				Num116 int `json:"116"`
				Num118 int `json:"118"`
				Num128 int `json:"128"`
			} `json:"146"`
			Num148 struct {
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num116 int `json:"116"`
				Num126 int `json:"126"`
			} `json:"148"`
			Num150 struct {
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num118 int `json:"118"`
				Num120 int `json:"120"`
				Num122 int `json:"122"`
				Num124 int `json:"124"`
			} `json:"150"`
			Num152 struct {
				Num90  int `json:"90"`
				Num92  int `json:"92"`
				Num94  int `json:"94"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num118 int `json:"118"`
			} `json:"152"`
			Num154 struct {
				Num90  int `json:"90"`
				Num94  int `json:"94"`
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
			} `json:"154"`
			Num156 struct {
				Num92 int `json:"92"`
				Num94 int `json:"94"`
				Num96 int `json:"96"`
				Num98 int `json:"98"`
			} `json:"156"`
			Num158 struct {
				Num94 int `json:"94"`
				Num96 int `json:"96"`
				Num98 int `json:"98"`
			} `json:"158"`
			Num160 struct {
				Num82 int `json:"82"`
				Num84 int `json:"84"`
				Num86 int `json:"86"`
				Num92 int `json:"92"`
				Num94 int `json:"94"`
				Num96 int `json:"96"`
				Num98 int `json:"98"`
			} `json:"160"`
			Num162 struct {
				Num80 int `json:"80"`
				Num86 int `json:"86"`
				Num88 int `json:"88"`
				Num90 int `json:"90"`
				Num94 int `json:"94"`
				Num96 int `json:"96"`
			} `json:"162"`
			Num164 struct {
				Num78 int `json:"78"`
				Num82 int `json:"82"`
				Num90 int `json:"90"`
				Num94 int `json:"94"`
				Num96 int `json:"96"`
				Num98 int `json:"98"`
			} `json:"164"`
			Num166 struct {
				Num82 int `json:"82"`
				Num84 int `json:"84"`
				Num86 int `json:"86"`
				Num88 int `json:"88"`
				Num90 int `json:"90"`
				Num94 int `json:"94"`
				Num96 int `json:"96"`
				Num98 int `json:"98"`
			} `json:"166"`
			Num168 struct {
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
			} `json:"168"`
			Num170 struct {
				Num84  int `json:"84"`
				Num86  int `json:"86"`
				Num88  int `json:"88"`
				Num90  int `json:"90"`
				Num92  int `json:"92"`
				Num94  int `json:"94"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num104 int `json:"104"`
			} `json:"170"`
			Num172 struct {
				Num86  int `json:"86"`
				Num88  int `json:"88"`
				Num90  int `json:"90"`
				Num92  int `json:"92"`
				Num94  int `json:"94"`
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num102 int `json:"102"`
			} `json:"172"`
			Num174 struct {
				Num88  int `json:"88"`
				Num90  int `json:"90"`
				Num92  int `json:"92"`
				Num94  int `json:"94"`
				Num96  int `json:"96"`
				Num100 int `json:"100"`
			} `json:"174"`
			Num176 struct {
				Num90  int `json:"90"`
				Num92  int `json:"92"`
				Num94  int `json:"94"`
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num102 int `json:"102"`
			} `json:"176"`
			Num178 struct {
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num102 int `json:"102"`
			} `json:"178"`
		} `json:"lane_pos,omitempty"`
		Obs1 struct {
			Num84 struct {
				Num110 int `json:"110"`
			} `json:"84"`
			Num94 struct {
				Num102 int `json:"102"`
				Num138 int `json:"138"`
			} `json:"94"`
			Num118 struct {
				Num166 int `json:"166"`
			} `json:"118"`
			Num120 struct {
				Num144 int `json:"144"`
			} `json:"120"`
			Num134 struct {
				Num134 int `json:"134"`
			} `json:"134"`
			Num136 struct {
				Num146 int `json:"146"`
				Num172 int `json:"172"`
			} `json:"136"`
			Num140 struct {
				Num102 int `json:"102"`
			} `json:"140"`
			Num142 struct {
				Num122 int `json:"122"`
			} `json:"142"`
			Num162 struct {
				Num114 int `json:"114"`
			} `json:"162"`
			Num174 struct {
				Num140 int `json:"140"`
			} `json:"174"`
		} `json:"obs,omitempty"`
		Purchase1 struct {
			Tango                int `json:"tango"`
			Circlet              int `json:"circlet"`
			FaerieFire           int `json:"faerie_fire"`
			WardSentry           int `json:"ward_sentry"`
			SmokeOfDeceit        int `json:"smoke_of_deceit"`
			Branches             int `json:"branches"`
			MagicStick           int `json:"magic_stick"`
			WardObserver         int `json:"ward_observer"`
			WardDispenser        int `json:"ward_dispenser"`
			Boots                int `json:"boots"`
			RecipeMagicWand      int `json:"recipe_magic_wand"`
			MagicWand            int `json:"magic_wand"`
			FluffyHat            int `json:"fluffy_hat"`
			StaffOfWizardry      int `json:"staff_of_wizardry"`
			RecipeForceStaff     int `json:"recipe_force_staff"`
			Dust                 int `json:"dust"`
			ForceStaff           int `json:"force_staff"`
			Gem                  int `json:"gem"`
			Chainmail            int `json:"chainmail"`
			SobiMask             int `json:"sobi_mask"`
			MedallionOfCourage   int `json:"medallion_of_courage"`
			BlightStone          int `json:"blight_stone"`
			OgreAxe              int `json:"ogre_axe"`
			BeltOfStrength       int `json:"belt_of_strength"`
			Sange                int `json:"sange"`
			RecipeSange          int `json:"recipe_sange"`
			RecipeHeavensHalberd int `json:"recipe_heavens_halberd"`
			TalismanOfEvasion    int `json:"talisman_of_evasion"`
			HeavensHalberd       int `json:"heavens_halberd"`
			TomeOfKnowledge      int `json:"tome_of_knowledge"`
		} `json:"purchase,omitempty"`
		Sen1 struct {
			Num70 struct {
				Num112 int `json:"112"`
			} `json:"70"`
			Num78 struct {
				Num156 int `json:"156"`
			} `json:"78"`
			Num80 struct {
				Num136 int `json:"136"`
			} `json:"80"`
			Num84 struct {
				Num110 int `json:"110"`
			} `json:"84"`
			Num92 struct {
				Num118 int `json:"118"`
			} `json:"92"`
			Num96 struct {
				Num88  int `json:"88"`
				Num140 int `json:"140"`
			} `json:"96"`
			Num116 struct {
				Num160 int `json:"160"`
				Num168 int `json:"168"`
			} `json:"116"`
			Num122 struct {
				Num138 int `json:"138"`
			} `json:"122"`
			Num124 struct {
				Num138 int `json:"138"`
			} `json:"124"`
			Num126 struct {
				Num164 int `json:"164"`
			} `json:"126"`
			Num130 struct {
				Num148 int `json:"148"`
				Num158 int `json:"158"`
			} `json:"130"`
			Num132 struct {
				Num160 int `json:"160"`
			} `json:"132"`
			Num144 struct {
				Num94  int `json:"94"`
				Num122 int `json:"122"`
			} `json:"144"`
			Num148 struct {
				Num144 int `json:"144"`
			} `json:"148"`
			Num152 struct {
				Num90 int `json:"90"`
			} `json:"152"`
			Num158 struct {
				Num120 int `json:"120"`
			} `json:"158"`
			Num162 struct {
				Num132 int `json:"132"`
			} `json:"162"`
		} `json:"sen,omitempty"`
		XpReasons1 struct {
			Num0 int `json:"0"`
			Num1 int `json:"1"`
			Num2 int `json:"2"`
			Num3 int `json:"3"`
			Num4 int `json:"4"`
			Num5 int `json:"5"`
		} `json:"xp_reasons,omitempty"`
		PurchaseTime1 struct {
			Tango              int `json:"tango"`
			Circlet            int `json:"circlet"`
			FaerieFire         int `json:"faerie_fire"`
			WardSentry         int `json:"ward_sentry"`
			SmokeOfDeceit      int `json:"smoke_of_deceit"`
			Branches           int `json:"branches"`
			MagicStick         int `json:"magic_stick"`
			WardObserver       int `json:"ward_observer"`
			Boots              int `json:"boots"`
			MagicWand          int `json:"magic_wand"`
			FluffyHat          int `json:"fluffy_hat"`
			StaffOfWizardry    int `json:"staff_of_wizardry"`
			Dust               int `json:"dust"`
			ForceStaff         int `json:"force_staff"`
			Gem                int `json:"gem"`
			Chainmail          int `json:"chainmail"`
			SobiMask           int `json:"sobi_mask"`
			MedallionOfCourage int `json:"medallion_of_courage"`
			BlightStone        int `json:"blight_stone"`
			OgreAxe            int `json:"ogre_axe"`
			BeltOfStrength     int `json:"belt_of_strength"`
			Sange              int `json:"sange"`
			TalismanOfEvasion  int `json:"talisman_of_evasion"`
			HeavensHalberd     int `json:"heavens_halberd"`
			TomeOfKnowledge    int `json:"tome_of_knowledge"`
		} `json:"purchase_time,omitempty"`
		FirstPurchaseTime1 struct {
			Tango              int `json:"tango"`
			Circlet            int `json:"circlet"`
			FaerieFire         int `json:"faerie_fire"`
			WardSentry         int `json:"ward_sentry"`
			SmokeOfDeceit      int `json:"smoke_of_deceit"`
			Branches           int `json:"branches"`
			MagicStick         int `json:"magic_stick"`
			WardObserver       int `json:"ward_observer"`
			Boots              int `json:"boots"`
			MagicWand          int `json:"magic_wand"`
			FluffyHat          int `json:"fluffy_hat"`
			StaffOfWizardry    int `json:"staff_of_wizardry"`
			Dust               int `json:"dust"`
			ForceStaff         int `json:"force_staff"`
			Gem                int `json:"gem"`
			Chainmail          int `json:"chainmail"`
			SobiMask           int `json:"sobi_mask"`
			MedallionOfCourage int `json:"medallion_of_courage"`
			BlightStone        int `json:"blight_stone"`
			OgreAxe            int `json:"ogre_axe"`
			BeltOfStrength     int `json:"belt_of_strength"`
			Sange              int `json:"sange"`
			TalismanOfEvasion  int `json:"talisman_of_evasion"`
			HeavensHalberd     int `json:"heavens_halberd"`
			TomeOfKnowledge    int `json:"tome_of_knowledge"`
		} `json:"first_purchase_time,omitempty"`
		ItemWin1 struct {
			Tango              int `json:"tango"`
			Circlet            int `json:"circlet"`
			FaerieFire         int `json:"faerie_fire"`
			WardSentry         int `json:"ward_sentry"`
			SmokeOfDeceit      int `json:"smoke_of_deceit"`
			Branches           int `json:"branches"`
			MagicStick         int `json:"magic_stick"`
			WardObserver       int `json:"ward_observer"`
			Boots              int `json:"boots"`
			MagicWand          int `json:"magic_wand"`
			FluffyHat          int `json:"fluffy_hat"`
			StaffOfWizardry    int `json:"staff_of_wizardry"`
			Dust               int `json:"dust"`
			ForceStaff         int `json:"force_staff"`
			Gem                int `json:"gem"`
			Chainmail          int `json:"chainmail"`
			SobiMask           int `json:"sobi_mask"`
			MedallionOfCourage int `json:"medallion_of_courage"`
			BlightStone        int `json:"blight_stone"`
			OgreAxe            int `json:"ogre_axe"`
			BeltOfStrength     int `json:"belt_of_strength"`
			Sange              int `json:"sange"`
			TalismanOfEvasion  int `json:"talisman_of_evasion"`
			HeavensHalberd     int `json:"heavens_halberd"`
			TomeOfKnowledge    int `json:"tome_of_knowledge"`
		} `json:"item_win,omitempty"`
		ItemUsage1 struct {
			Tango              int `json:"tango"`
			Circlet            int `json:"circlet"`
			FaerieFire         int `json:"faerie_fire"`
			WardSentry         int `json:"ward_sentry"`
			SmokeOfDeceit      int `json:"smoke_of_deceit"`
			Branches           int `json:"branches"`
			MagicStick         int `json:"magic_stick"`
			WardObserver       int `json:"ward_observer"`
			Boots              int `json:"boots"`
			MagicWand          int `json:"magic_wand"`
			FluffyHat          int `json:"fluffy_hat"`
			StaffOfWizardry    int `json:"staff_of_wizardry"`
			Dust               int `json:"dust"`
			ForceStaff         int `json:"force_staff"`
			Gem                int `json:"gem"`
			Chainmail          int `json:"chainmail"`
			SobiMask           int `json:"sobi_mask"`
			MedallionOfCourage int `json:"medallion_of_courage"`
			BlightStone        int `json:"blight_stone"`
			OgreAxe            int `json:"ogre_axe"`
			BeltOfStrength     int `json:"belt_of_strength"`
			Sange              int `json:"sange"`
			TalismanOfEvasion  int `json:"talisman_of_evasion"`
			HeavensHalberd     int `json:"heavens_halberd"`
			TomeOfKnowledge    int `json:"tome_of_knowledge"`
		} `json:"item_usage,omitempty"`
		PurchaseGem     int `json:"purchase_gem,omitempty"`
		AbilityTargets2 struct {
		} `json:"ability_targets,omitempty"`
		AbilityUses2 struct {
			TemplarAssassinMeld        int `json:"templar_assassin_meld"`
			TemplarAssassinRefraction  int `json:"templar_assassin_refraction"`
			TemplarAssassinPsionicTrap int `json:"templar_assassin_psionic_trap"`
		} `json:"ability_uses,omitempty"`
		Actions2 struct {
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
			Num12 int `json:"12"`
			Num13 int `json:"13"`
			Num14 int `json:"14"`
			Num15 int `json:"15"`
			Num16 int `json:"16"`
			Num17 int `json:"17"`
			Num19 int `json:"19"`
			Num23 int `json:"23"`
			Num28 int `json:"28"`
			Num37 int `json:"37"`
		} `json:"actions,omitempty"`
		Damage2 struct {
			NpcDotaHeroMirana                      int `json:"npc_dota_hero_mirana"`
			NpcDotaCreepGoodguysMelee              int `json:"npc_dota_creep_goodguys_melee"`
			NpcDotaHeroDragonKnight                int `json:"npc_dota_hero_dragon_knight"`
			NpcDotaCreepBadguysMelee               int `json:"npc_dota_creep_badguys_melee"`
			NpcDotaCreepGoodguysRanged             int `json:"npc_dota_creep_goodguys_ranged"`
			NpcDotaCreepBadguysRanged              int `json:"npc_dota_creep_badguys_ranged"`
			NpcDotaCreepBadguysFlagbearer          int `json:"npc_dota_creep_badguys_flagbearer"`
			NpcDotaBadguysSiege                    int `json:"npc_dota_badguys_siege"`
			NpcDotaNeutralDarkTroll                int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaNeutralCentaurKhan              int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaNeutralCentaurOutrunner         int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaNeutralMudGolem                 int `json:"npc_dota_neutral_mud_golem"`
			NpcDotaNeutralMudGolemSplit            int `json:"npc_dota_neutral_mud_golem_split"`
			NpcDotaNeutralKobold                   int `json:"npc_dota_neutral_kobold"`
			NpcDotaObserverWards                   int `json:"npc_dota_observer_wards"`
			NpcDotaNeutralDarkTrollWarlord         int `json:"npc_dota_neutral_dark_troll_warlord"`
			NpcDotaDarkTrollWarlordSkeletonWarrior int `json:"npc_dota_dark_troll_warlord_skeleton_warrior"`
			NpcDotaHeroQueenofpain                 int `json:"npc_dota_hero_queenofpain"`
			NpcDotaNeutralAlphaWolf                int `json:"npc_dota_neutral_alpha_wolf"`
			NpcDotaNeutralGiantWolf                int `json:"npc_dota_neutral_giant_wolf"`
			NpcDotaNeutralOgreMauler               int `json:"npc_dota_neutral_ogre_mauler"`
			NpcDotaNeutralOgreMagi                 int `json:"npc_dota_neutral_ogre_magi"`
			NpcDotaNeutralWarpineRaider            int `json:"npc_dota_neutral_warpine_raider"`
			NpcDotaNeutralSatyrTrickster           int `json:"npc_dota_neutral_satyr_trickster"`
			NpcDotaNeutralSatyrSoulstealer         int `json:"npc_dota_neutral_satyr_soulstealer"`
			NpcDotaNeutralSatyrHellcaller          int `json:"npc_dota_neutral_satyr_hellcaller"`
			NpcDotaNeutralFrostbittenGolem         int `json:"npc_dota_neutral_frostbitten_golem"`
			NpcDotaNeutralSmallThunderLizard       int `json:"npc_dota_neutral_small_thunder_lizard"`
			NpcDotaNeutralBigThunderLizard         int `json:"npc_dota_neutral_big_thunder_lizard"`
			NpcDotaNeutralIceShaman                int `json:"npc_dota_neutral_ice_shaman"`
			NpcDotaNeutralPolarFurbolgChampion     int `json:"npc_dota_neutral_polar_furbolg_champion"`
			NpcDotaNeutralWildkin                  int `json:"npc_dota_neutral_wildkin"`
			NpcDotaNeutralPolarFurbolgUrsaWarrior  int `json:"npc_dota_neutral_polar_furbolg_ursa_warrior"`
			NpcDotaNeutralEnragedWildkin           int `json:"npc_dota_neutral_enraged_wildkin"`
			NpcDotaNeutralBlackDrake               int `json:"npc_dota_neutral_black_drake"`
			NpcDotaNeutralRockGolem                int `json:"npc_dota_neutral_rock_golem"`
			NpcDotaNeutralBlackDragon              int `json:"npc_dota_neutral_black_dragon"`
			NpcDotaNeutralGraniteGolem             int `json:"npc_dota_neutral_granite_golem"`
			NpcDotaNeutralFelBeast                 int `json:"npc_dota_neutral_fel_beast"`
			NpcDotaNeutralGhost                    int `json:"npc_dota_neutral_ghost"`
			NpcDotaRoshan                          int `json:"npc_dota_roshan"`
			NpcDotaHeroChen                        int `json:"npc_dota_hero_chen"`
			NpcDotaNeutralGnollAssassin            int `json:"npc_dota_neutral_gnoll_assassin"`
			NpcDotaBadguysTower1Bot                int `json:"npc_dota_badguys_tower1_bot"`
			NpcDotaBadguysTower2Mid                int `json:"npc_dota_badguys_tower2_mid"`
			NpcDotaHeroTrollWarlord                int `json:"npc_dota_hero_troll_warlord"`
			NpcDotaNeutralForestTrollBerserker     int `json:"npc_dota_neutral_forest_troll_berserker"`
			NpcDotaBadguysTower2Top                int `json:"npc_dota_badguys_tower2_top"`
			NpcDotaNeutralKoboldTaskmaster         int `json:"npc_dota_neutral_kobold_taskmaster"`
			NpcDotaNeutralKoboldTunneler           int `json:"npc_dota_neutral_kobold_tunneler"`
			NpcDotaNeutralForestTrollHighPriest    int `json:"npc_dota_neutral_forest_troll_high_priest"`
			NpcDotaBadguysTower2Bot                int `json:"npc_dota_badguys_tower2_bot"`
			NpcDotaCreepBadguysMeleeUpgraded       int `json:"npc_dota_creep_badguys_melee_upgraded"`
			NpcDotaCreepBadguysRangedUpgraded      int `json:"npc_dota_creep_badguys_ranged_upgraded"`
			NpcDotaCreepBadguysFlagbearerUpgraded  int `json:"npc_dota_creep_badguys_flagbearer_upgraded"`
			NpcDotaBadguysSiegeUpgraded            int `json:"npc_dota_badguys_siege_upgraded"`
		} `json:"damage,omitempty"`
		DamageInflictor2 struct {
			TemplarAssassinMeld      int `json:"templar_assassin_meld"`
			Null                     int `json:"null"`
			TemplarAssassinPsiBlades int `json:"templar_assassin_psi_blades"`
			TemplarAssassinTrap      int `json:"templar_assassin_trap"`
		} `json:"damage_inflictor,omitempty"`
		DamageInflictorReceived2 struct {
			DragonKnightDragonTail         int `json:"dragon_knight_dragon_tail"`
			Null                           int `json:"null"`
			Bfury                          int `json:"bfury"`
			QueenofpainSonicWave           int `json:"queenofpain_sonic_wave"`
			TrollWarlordWhirlingAxesMelee  int `json:"troll_warlord_whirling_axes_melee"`
			QueenofpainShadowStrike        int `json:"queenofpain_shadow_strike"`
			Orchid                         int `json:"orchid"`
			DragonKnightFireball           int `json:"dragon_knight_fireball"`
			DragonKnightElderDragonForm    int `json:"dragon_knight_elder_dragon_form"`
			MiranaArrow                    int `json:"mirana_arrow"`
			ShivasGuard                    int `json:"shivas_guard"`
			DragonKnightBreatheFire        int `json:"dragon_knight_breathe_fire"`
			WraithPact                     int `json:"wraith_pact"`
			QueenofpainBlink               int `json:"queenofpain_blink"`
			TrollWarlordWhirlingAxesRanged int `json:"troll_warlord_whirling_axes_ranged"`
		} `json:"damage_inflictor_received,omitempty"`
		DamageTaken2 struct {
			NpcDotaHeroDragonKnight                int `json:"npc_dota_hero_dragon_knight"`
			NpcDotaCreepBadguysRanged              int `json:"npc_dota_creep_badguys_ranged"`
			NpcDotaCreepBadguysMelee               int `json:"npc_dota_creep_badguys_melee"`
			NpcDotaHeroMirana                      int `json:"npc_dota_hero_mirana"`
			NpcDotaCreepBadguysFlagbearer          int `json:"npc_dota_creep_badguys_flagbearer"`
			NpcDotaNeutralDarkTroll                int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaNeutralCentaurOutrunner         int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaNeutralMudGolemSplit            int `json:"npc_dota_neutral_mud_golem_split"`
			NpcDotaNeutralKobold                   int `json:"npc_dota_neutral_kobold"`
			NpcDotaNeutralDarkTrollWarlord         int `json:"npc_dota_neutral_dark_troll_warlord"`
			NpcDotaDarkTrollWarlordSkeletonWarrior int `json:"npc_dota_dark_troll_warlord_skeleton_warrior"`
			NpcDotaNeutralOgreMagi                 int `json:"npc_dota_neutral_ogre_magi"`
			NpcDotaNeutralOgreMauler               int `json:"npc_dota_neutral_ogre_mauler"`
			NpcDotaNeutralIceShaman                int `json:"npc_dota_neutral_ice_shaman"`
			NpcDotaNeutralSmallThunderLizard       int `json:"npc_dota_neutral_small_thunder_lizard"`
			NpcDotaNeutralFrostbittenGolem         int `json:"npc_dota_neutral_frostbitten_golem"`
			NpcDotaNeutralBigThunderLizard         int `json:"npc_dota_neutral_big_thunder_lizard"`
			NpcDotaNeutralGraniteGolem             int `json:"npc_dota_neutral_granite_golem"`
			NpcDotaNeutralBlackDragon              int `json:"npc_dota_neutral_black_dragon"`
			NpcDotaNeutralBlackDrake               int `json:"npc_dota_neutral_black_drake"`
			NpcDotaNeutralRockGolem                int `json:"npc_dota_neutral_rock_golem"`
			NpcDotaNeutralSatyrHellcaller          int `json:"npc_dota_neutral_satyr_hellcaller"`
			NpcDotaNeutralWarpineRaider            int `json:"npc_dota_neutral_warpine_raider"`
			NpcDotaNeutralEnragedWildkin           int `json:"npc_dota_neutral_enraged_wildkin"`
			NpcDotaNeutralWildkin                  int `json:"npc_dota_neutral_wildkin"`
			NpcDotaNeutralCentaurKhan              int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaNeutralGiantWolf                int `json:"npc_dota_neutral_giant_wolf"`
			NpcDotaHeroTrollWarlord                int `json:"npc_dota_hero_troll_warlord"`
			NpcDotaHeroQueenofpain                 int `json:"npc_dota_hero_queenofpain"`
			NpcDotaHeroChen                        int `json:"npc_dota_hero_chen"`
			NpcDotaNeutralPolarFurbolgChampion     int `json:"npc_dota_neutral_polar_furbolg_champion"`
			NpcDotaCreepBadguysRangedUpgraded      int `json:"npc_dota_creep_badguys_ranged_upgraded"`
			NpcDotaCreepBadguysMeleeUpgraded       int `json:"npc_dota_creep_badguys_melee_upgraded"`
			NpcDotaEnragedWildkinTornado           int `json:"npc_dota_enraged_wildkin_tornado"`
		} `json:"damage_taken,omitempty"`
		DamageTargets2 struct {
			TemplarAssassinMeld struct {
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroChen         int `json:"npc_dota_hero_chen"`
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
				NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
			} `json:"templar_assassin_meld"`
			Null struct {
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
				NpcDotaHeroChen         int `json:"npc_dota_hero_chen"`
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
			} `json:"null"`
			TemplarAssassinPsiBlades struct {
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
			} `json:"templar_assassin_psi_blades"`
			TemplarAssassinTrap struct {
				NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
				NpcDotaHeroChen         int `json:"npc_dota_hero_chen"`
			} `json:"templar_assassin_trap"`
		} `json:"damage_targets,omitempty"`
		GoldReasons2 struct {
			Num0  int `json:"0"`
			Num1  int `json:"1"`
			Num6  int `json:"6"`
			Num11 int `json:"11"`
			Num12 int `json:"12"`
			Num13 int `json:"13"`
			Num14 int `json:"14"`
			Num15 int `json:"15"`
			Num16 int `json:"16"`
			Num17 int `json:"17"`
			Num19 int `json:"19"`
		} `json:"gold_reasons,omitempty"`
		HeroHits2 struct {
			TemplarAssassinMeld      int `json:"templar_assassin_meld"`
			Null                     int `json:"null"`
			TemplarAssassinPsiBlades int `json:"templar_assassin_psi_blades"`
			TemplarAssassinTrap      int `json:"templar_assassin_trap"`
		} `json:"hero_hits,omitempty"`
		ItemUses2 struct {
			Tango          int `json:"tango"`
			MagicStick     int `json:"magic_stick"`
			PowerTreads    int `json:"power_treads"`
			TrustyShovel   int `json:"trusty_shovel"`
			Flask          int `json:"flask"`
			EnchantedMango int `json:"enchanted_mango"`
			Clarity        int `json:"clarity"`
			Tpscroll       int `json:"tpscroll"`
			MagicWand      int `json:"magic_wand"`
			Bullwhip       int `json:"bullwhip"`
			Blink          int `json:"blink"`
			DaggerOfRistul int `json:"dagger_of_ristul"`
			HurricanePike  int `json:"hurricane_pike"`
			TricksterCloak int `json:"trickster_cloak"`
		} `json:"item_uses,omitempty"`
		KillStreaks1 struct {
			Num3 int `json:"3"`
		} `json:"kill_streaks,omitempty"`
		Killed2 struct {
			NpcDotaCreepGoodguysMelee              int `json:"npc_dota_creep_goodguys_melee"`
			NpcDotaCreepBadguysMelee               int `json:"npc_dota_creep_badguys_melee"`
			NpcDotaCreepGoodguysRanged             int `json:"npc_dota_creep_goodguys_ranged"`
			NpcDotaCreepBadguysRanged              int `json:"npc_dota_creep_badguys_ranged"`
			NpcDotaCreepBadguysFlagbearer          int `json:"npc_dota_creep_badguys_flagbearer"`
			NpcDotaHeroDragonKnight                int `json:"npc_dota_hero_dragon_knight"`
			NpcDotaBadguysSiege                    int `json:"npc_dota_badguys_siege"`
			NpcDotaNeutralDarkTroll                int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaNeutralCentaurOutrunner         int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaNeutralCentaurKhan              int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaNeutralMudGolem                 int `json:"npc_dota_neutral_mud_golem"`
			NpcDotaNeutralMudGolemSplit            int `json:"npc_dota_neutral_mud_golem_split"`
			NpcDotaNeutralKobold                   int `json:"npc_dota_neutral_kobold"`
			NpcDotaObserverWards                   int `json:"npc_dota_observer_wards"`
			NpcDotaDarkTrollWarlordSkeletonWarrior int `json:"npc_dota_dark_troll_warlord_skeleton_warrior"`
			NpcDotaNeutralDarkTrollWarlord         int `json:"npc_dota_neutral_dark_troll_warlord"`
			NpcDotaNeutralGiantWolf                int `json:"npc_dota_neutral_giant_wolf"`
			NpcDotaNeutralAlphaWolf                int `json:"npc_dota_neutral_alpha_wolf"`
			NpcDotaNeutralOgreMagi                 int `json:"npc_dota_neutral_ogre_magi"`
			NpcDotaNeutralOgreMauler               int `json:"npc_dota_neutral_ogre_mauler"`
			NpcDotaNeutralWarpineRaider            int `json:"npc_dota_neutral_warpine_raider"`
			NpcDotaNeutralSatyrTrickster           int `json:"npc_dota_neutral_satyr_trickster"`
			NpcDotaNeutralSatyrSoulstealer         int `json:"npc_dota_neutral_satyr_soulstealer"`
			NpcDotaNeutralSatyrHellcaller          int `json:"npc_dota_neutral_satyr_hellcaller"`
			NpcDotaNeutralSmallThunderLizard       int `json:"npc_dota_neutral_small_thunder_lizard"`
			NpcDotaNeutralFrostbittenGolem         int `json:"npc_dota_neutral_frostbitten_golem"`
			NpcDotaNeutralIceShaman                int `json:"npc_dota_neutral_ice_shaman"`
			NpcDotaNeutralBigThunderLizard         int `json:"npc_dota_neutral_big_thunder_lizard"`
			NpcDotaNeutralWildkin                  int `json:"npc_dota_neutral_wildkin"`
			NpcDotaNeutralEnragedWildkin           int `json:"npc_dota_neutral_enraged_wildkin"`
			NpcDotaNeutralPolarFurbolgChampion     int `json:"npc_dota_neutral_polar_furbolg_champion"`
			NpcDotaNeutralPolarFurbolgUrsaWarrior  int `json:"npc_dota_neutral_polar_furbolg_ursa_warrior"`
			NpcDotaNeutralRockGolem                int `json:"npc_dota_neutral_rock_golem"`
			NpcDotaNeutralBlackDrake               int `json:"npc_dota_neutral_black_drake"`
			NpcDotaNeutralGraniteGolem             int `json:"npc_dota_neutral_granite_golem"`
			NpcDotaNeutralBlackDragon              int `json:"npc_dota_neutral_black_dragon"`
			NpcDotaNeutralFelBeast                 int `json:"npc_dota_neutral_fel_beast"`
			NpcDotaNeutralGhost                    int `json:"npc_dota_neutral_ghost"`
			NpcDotaRoshan                          int `json:"npc_dota_roshan"`
			NpcDotaHeroChen                        int `json:"npc_dota_hero_chen"`
			NpcDotaNeutralGnollAssassin            int `json:"npc_dota_neutral_gnoll_assassin"`
			NpcDotaBadguysTower1Bot                int `json:"npc_dota_badguys_tower1_bot"`
			NpcDotaNeutralForestTrollBerserker     int `json:"npc_dota_neutral_forest_troll_berserker"`
			NpcDotaNeutralKoboldTaskmaster         int `json:"npc_dota_neutral_kobold_taskmaster"`
			NpcDotaNeutralKoboldTunneler           int `json:"npc_dota_neutral_kobold_tunneler"`
			NpcDotaNeutralForestTrollHighPriest    int `json:"npc_dota_neutral_forest_troll_high_priest"`
			NpcDotaBadguysTower2Bot                int `json:"npc_dota_badguys_tower2_bot"`
			NpcDotaCreepBadguysMeleeUpgraded       int `json:"npc_dota_creep_badguys_melee_upgraded"`
			NpcDotaCreepBadguysRangedUpgraded      int `json:"npc_dota_creep_badguys_ranged_upgraded"`
			NpcDotaCreepBadguysFlagbearerUpgraded  int `json:"npc_dota_creep_badguys_flagbearer_upgraded"`
		} `json:"killed,omitempty"`
		KilledBy1 struct {
			NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
			NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
			NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
		} `json:"killed_by,omitempty"`
		LanePos2 struct {
			Num72 struct {
				Num78 int `json:"78"`
			} `json:"72"`
			Num74 struct {
				Num78 int `json:"78"`
			} `json:"74"`
			Num76 struct {
				Num78 int `json:"78"`
			} `json:"76"`
			Num78 struct {
				Num80 int `json:"80"`
			} `json:"78"`
			Num80 struct {
				Num80 int `json:"80"`
			} `json:"80"`
			Num82 struct {
				Num82 int `json:"82"`
			} `json:"82"`
			Num84 struct {
				Num82 int `json:"82"`
			} `json:"84"`
			Num86 struct {
				Num84 int `json:"84"`
			} `json:"86"`
			Num88 struct {
				Num86 int `json:"86"`
			} `json:"88"`
			Num90 struct {
				Num88 int `json:"88"`
			} `json:"90"`
			Num92 struct {
				Num90 int `json:"90"`
				Num92 int `json:"92"`
			} `json:"92"`
			Num94 struct {
				Num94 int `json:"94"`
			} `json:"94"`
			Num96 struct {
				Num96 int `json:"96"`
			} `json:"96"`
			Num98 struct {
				Num98 int `json:"98"`
			} `json:"98"`
			Num102 struct {
				Num98 int `json:"98"`
			} `json:"102"`
			Num104 struct {
				Num98 int `json:"98"`
			} `json:"104"`
			Num108 struct {
				Num100 int `json:"100"`
			} `json:"108"`
			Num110 struct {
				Num102 int `json:"102"`
			} `json:"110"`
			Num112 struct {
				Num102 int `json:"102"`
			} `json:"112"`
			Num114 struct {
				Num104 int `json:"104"`
			} `json:"114"`
			Num116 struct {
				Num106 int `json:"106"`
			} `json:"116"`
			Num118 struct {
				Num106 int `json:"106"`
			} `json:"118"`
			Num120 struct {
				Num116 int `json:"116"`
			} `json:"120"`
			Num122 struct {
				Num106 int `json:"106"`
				Num108 int `json:"108"`
				Num112 int `json:"112"`
				Num116 int `json:"116"`
			} `json:"122"`
			Num124 struct {
				Num90  int `json:"90"`
				Num92  int `json:"92"`
				Num94  int `json:"94"`
				Num108 int `json:"108"`
				Num110 int `json:"110"`
				Num114 int `json:"114"`
			} `json:"124"`
			Num126 struct {
				Num78  int `json:"78"`
				Num88  int `json:"88"`
				Num96  int `json:"96"`
				Num108 int `json:"108"`
				Num114 int `json:"114"`
			} `json:"126"`
			Num128 struct {
				Num78  int `json:"78"`
				Num80  int `json:"80"`
				Num84  int `json:"84"`
				Num88  int `json:"88"`
				Num94  int `json:"94"`
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num102 int `json:"102"`
				Num104 int `json:"104"`
				Num106 int `json:"106"`
				Num114 int `json:"114"`
			} `json:"128"`
			Num130 struct {
				Num80  int `json:"80"`
				Num86  int `json:"86"`
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num104 int `json:"104"`
				Num106 int `json:"106"`
			} `json:"130"`
			Num132 struct {
				Num78  int `json:"78"`
				Num82  int `json:"82"`
				Num84  int `json:"84"`
				Num86  int `json:"86"`
				Num96  int `json:"96"`
				Num104 int `json:"104"`
				Num106 int `json:"106"`
			} `json:"132"`
			Num134 struct {
				Num78  int `json:"78"`
				Num80  int `json:"80"`
				Num88  int `json:"88"`
				Num94  int `json:"94"`
				Num100 int `json:"100"`
				Num102 int `json:"102"`
				Num104 int `json:"104"`
			} `json:"134"`
			Num136 struct {
				Num78  int `json:"78"`
				Num88  int `json:"88"`
				Num90  int `json:"90"`
				Num92  int `json:"92"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num102 int `json:"102"`
				Num118 int `json:"118"`
			} `json:"136"`
			Num138 struct {
				Num78  int `json:"78"`
				Num90  int `json:"90"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num102 int `json:"102"`
				Num114 int `json:"114"`
				Num116 int `json:"116"`
				Num118 int `json:"118"`
			} `json:"138"`
			Num140 struct {
				Num78  int `json:"78"`
				Num92  int `json:"92"`
				Num96  int `json:"96"`
				Num100 int `json:"100"`
				Num102 int `json:"102"`
				Num110 int `json:"110"`
				Num112 int `json:"112"`
				Num114 int `json:"114"`
			} `json:"140"`
			Num142 struct {
				Num78  int `json:"78"`
				Num90  int `json:"90"`
				Num92  int `json:"92"`
				Num94  int `json:"94"`
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num102 int `json:"102"`
				Num106 int `json:"106"`
				Num108 int `json:"108"`
				Num110 int `json:"110"`
			} `json:"142"`
			Num144 struct {
				Num82  int `json:"82"`
				Num84  int `json:"84"`
				Num86  int `json:"86"`
				Num88  int `json:"88"`
				Num90  int `json:"90"`
				Num94  int `json:"94"`
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num102 int `json:"102"`
				Num104 int `json:"104"`
			} `json:"144"`
			Num146 struct {
				Num80  int `json:"80"`
				Num88  int `json:"88"`
				Num100 int `json:"100"`
				Num102 int `json:"102"`
			} `json:"146"`
			Num148 struct {
				Num78  int `json:"78"`
				Num80  int `json:"80"`
				Num88  int `json:"88"`
				Num100 int `json:"100"`
				Num102 int `json:"102"`
			} `json:"148"`
			Num150 struct {
				Num80  int `json:"80"`
				Num86  int `json:"86"`
				Num88  int `json:"88"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
			} `json:"150"`
			Num152 struct {
				Num78 int `json:"78"`
				Num80 int `json:"80"`
				Num82 int `json:"82"`
				Num84 int `json:"84"`
				Num86 int `json:"86"`
				Num98 int `json:"98"`
			} `json:"152"`
			Num154 struct {
				Num78 int `json:"78"`
				Num80 int `json:"80"`
				Num82 int `json:"82"`
				Num84 int `json:"84"`
				Num96 int `json:"96"`
				Num98 int `json:"98"`
			} `json:"154"`
			Num156 struct {
				Num78 int `json:"78"`
				Num80 int `json:"80"`
				Num82 int `json:"82"`
				Num84 int `json:"84"`
				Num96 int `json:"96"`
			} `json:"156"`
			Num158 struct {
				Num80 int `json:"80"`
				Num94 int `json:"94"`
				Num96 int `json:"96"`
			} `json:"158"`
			Num160 struct {
				Num80 int `json:"80"`
				Num90 int `json:"90"`
				Num92 int `json:"92"`
				Num96 int `json:"96"`
			} `json:"160"`
			Num162 struct {
				Num78 int `json:"78"`
				Num80 int `json:"80"`
				Num88 int `json:"88"`
				Num96 int `json:"96"`
			} `json:"162"`
			Num164 struct {
				Num80 int `json:"80"`
				Num88 int `json:"88"`
				Num96 int `json:"96"`
			} `json:"164"`
			Num166 struct {
				Num78 int `json:"78"`
				Num80 int `json:"80"`
				Num82 int `json:"82"`
				Num86 int `json:"86"`
				Num96 int `json:"96"`
			} `json:"166"`
			Num168 struct {
				Num80 int `json:"80"`
				Num82 int `json:"82"`
				Num86 int `json:"86"`
				Num94 int `json:"94"`
			} `json:"168"`
			Num170 struct {
				Num80 int `json:"80"`
				Num82 int `json:"82"`
				Num84 int `json:"84"`
				Num86 int `json:"86"`
			} `json:"170"`
			Num172 struct {
				Num80  int `json:"80"`
				Num82  int `json:"82"`
				Num84  int `json:"84"`
				Num86  int `json:"86"`
				Num88  int `json:"88"`
				Num90  int `json:"90"`
				Num94  int `json:"94"`
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
			} `json:"174"`
			Num176 struct {
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
			} `json:"176"`
			Num178 struct {
				Num90  int `json:"90"`
				Num92  int `json:"92"`
				Num100 int `json:"100"`
				Num102 int `json:"102"`
				Num104 int `json:"104"`
			} `json:"178"`
		} `json:"lane_pos,omitempty"`
		Obs2 struct {
		} `json:"obs,omitempty"`
		Purchase2 struct {
			Circlet              int `json:"circlet"`
			MagicStick           int `json:"magic_stick"`
			Tango                int `json:"tango"`
			Branches             int `json:"branches"`
			BootsOfElves         int `json:"boots_of_elves"`
			Boots                int `json:"boots"`
			Gloves               int `json:"gloves"`
			PowerTreads          int `json:"power_treads"`
			BladeOfAlacrity      int `json:"blade_of_alacrity"`
			BeltOfStrength       int `json:"belt_of_strength"`
			RecipeDragonLance    int `json:"recipe_dragon_lance"`
			Clarity              int `json:"clarity"`
			DragonLance          int `json:"dragon_lance"`
			SobiMask             int `json:"sobi_mask"`
			RingOfBasilius       int `json:"ring_of_basilius"`
			RecipeRingOfBasilius int `json:"recipe_ring_of_basilius"`
			RecipeMagicWand      int `json:"recipe_magic_wand"`
			MagicWand            int `json:"magic_wand"`
			BlightStone          int `json:"blight_stone"`
			Tpscroll             int `json:"tpscroll"`
			MithrilHammer        int `json:"mithril_hammer"`
			Desolator            int `json:"desolator"`
			AghanimsShard        int `json:"aghanims_shard"`
			Blink                int `json:"blink"`
			RecipeSphere         int `json:"recipe_sphere"`
			UltimateOrb          int `json:"ultimate_orb"`
			RingOfHealth         int `json:"ring_of_health"`
			Pers                 int `json:"pers"`
			VoidStone            int `json:"void_stone"`
			Sphere               int `json:"sphere"`
			StaffOfWizardry      int `json:"staff_of_wizardry"`
			FluffyHat            int `json:"fluffy_hat"`
			ForceStaff           int `json:"force_staff"`
			RecipeForceStaff     int `json:"recipe_force_staff"`
			RecipeHurricanePike  int `json:"recipe_hurricane_pike"`
			HurricanePike        int `json:"hurricane_pike"`
			Broadsword           int `json:"broadsword"`
			BladesOfAttack       int `json:"blades_of_attack"`
			LesserCrit           int `json:"lesser_crit"`
			RecipeLesserCrit     int `json:"recipe_lesser_crit"`
			RecipeGreaterCrit    int `json:"recipe_greater_crit"`
			DemonEdge            int `json:"demon_edge"`
			GreaterCrit          int `json:"greater_crit"`
		} `json:"purchase,omitempty"`
		Runes1 struct {
			Num0 int `json:"0"`
		} `json:"runes,omitempty"`
		Sen2 struct {
		} `json:"sen,omitempty"`
		XpReasons2 struct {
			Num0 int `json:"0"`
			Num1 int `json:"1"`
			Num2 int `json:"2"`
			Num3 int `json:"3"`
			Num5 int `json:"5"`
		} `json:"xp_reasons,omitempty"`
		PurchaseTime2 struct {
			Circlet         int `json:"circlet"`
			MagicStick      int `json:"magic_stick"`
			Tango           int `json:"tango"`
			Branches        int `json:"branches"`
			BootsOfElves    int `json:"boots_of_elves"`
			Boots           int `json:"boots"`
			Gloves          int `json:"gloves"`
			PowerTreads     int `json:"power_treads"`
			BladeOfAlacrity int `json:"blade_of_alacrity"`
			BeltOfStrength  int `json:"belt_of_strength"`
			Clarity         int `json:"clarity"`
			DragonLance     int `json:"dragon_lance"`
			SobiMask        int `json:"sobi_mask"`
			RingOfBasilius  int `json:"ring_of_basilius"`
			MagicWand       int `json:"magic_wand"`
			BlightStone     int `json:"blight_stone"`
			Tpscroll        int `json:"tpscroll"`
			MithrilHammer   int `json:"mithril_hammer"`
			Desolator       int `json:"desolator"`
			AghanimsShard   int `json:"aghanims_shard"`
			Blink           int `json:"blink"`
			UltimateOrb     int `json:"ultimate_orb"`
			RingOfHealth    int `json:"ring_of_health"`
			Pers            int `json:"pers"`
			VoidStone       int `json:"void_stone"`
			Sphere          int `json:"sphere"`
			StaffOfWizardry int `json:"staff_of_wizardry"`
			FluffyHat       int `json:"fluffy_hat"`
			ForceStaff      int `json:"force_staff"`
			HurricanePike   int `json:"hurricane_pike"`
			Broadsword      int `json:"broadsword"`
			BladesOfAttack  int `json:"blades_of_attack"`
			LesserCrit      int `json:"lesser_crit"`
			DemonEdge       int `json:"demon_edge"`
			GreaterCrit     int `json:"greater_crit"`
		} `json:"purchase_time,omitempty"`
		FirstPurchaseTime2 struct {
			Circlet         int `json:"circlet"`
			MagicStick      int `json:"magic_stick"`
			Tango           int `json:"tango"`
			Branches        int `json:"branches"`
			BootsOfElves    int `json:"boots_of_elves"`
			Boots           int `json:"boots"`
			Gloves          int `json:"gloves"`
			PowerTreads     int `json:"power_treads"`
			BladeOfAlacrity int `json:"blade_of_alacrity"`
			BeltOfStrength  int `json:"belt_of_strength"`
			Clarity         int `json:"clarity"`
			DragonLance     int `json:"dragon_lance"`
			SobiMask        int `json:"sobi_mask"`
			RingOfBasilius  int `json:"ring_of_basilius"`
			MagicWand       int `json:"magic_wand"`
			BlightStone     int `json:"blight_stone"`
			Tpscroll        int `json:"tpscroll"`
			MithrilHammer   int `json:"mithril_hammer"`
			Desolator       int `json:"desolator"`
			AghanimsShard   int `json:"aghanims_shard"`
			Blink           int `json:"blink"`
			UltimateOrb     int `json:"ultimate_orb"`
			RingOfHealth    int `json:"ring_of_health"`
			Pers            int `json:"pers"`
			VoidStone       int `json:"void_stone"`
			Sphere          int `json:"sphere"`
			StaffOfWizardry int `json:"staff_of_wizardry"`
			FluffyHat       int `json:"fluffy_hat"`
			ForceStaff      int `json:"force_staff"`
			HurricanePike   int `json:"hurricane_pike"`
			Broadsword      int `json:"broadsword"`
			BladesOfAttack  int `json:"blades_of_attack"`
			LesserCrit      int `json:"lesser_crit"`
			DemonEdge       int `json:"demon_edge"`
			GreaterCrit     int `json:"greater_crit"`
		} `json:"first_purchase_time,omitempty"`
		ItemWin2 struct {
			Circlet         int `json:"circlet"`
			MagicStick      int `json:"magic_stick"`
			Tango           int `json:"tango"`
			Branches        int `json:"branches"`
			BootsOfElves    int `json:"boots_of_elves"`
			Boots           int `json:"boots"`
			Gloves          int `json:"gloves"`
			PowerTreads     int `json:"power_treads"`
			BladeOfAlacrity int `json:"blade_of_alacrity"`
			BeltOfStrength  int `json:"belt_of_strength"`
			Clarity         int `json:"clarity"`
			DragonLance     int `json:"dragon_lance"`
			SobiMask        int `json:"sobi_mask"`
			RingOfBasilius  int `json:"ring_of_basilius"`
			MagicWand       int `json:"magic_wand"`
			BlightStone     int `json:"blight_stone"`
			Tpscroll        int `json:"tpscroll"`
			MithrilHammer   int `json:"mithril_hammer"`
			Desolator       int `json:"desolator"`
			AghanimsShard   int `json:"aghanims_shard"`
			Blink           int `json:"blink"`
			UltimateOrb     int `json:"ultimate_orb"`
			RingOfHealth    int `json:"ring_of_health"`
			Pers            int `json:"pers"`
			VoidStone       int `json:"void_stone"`
			Sphere          int `json:"sphere"`
			StaffOfWizardry int `json:"staff_of_wizardry"`
			FluffyHat       int `json:"fluffy_hat"`
			ForceStaff      int `json:"force_staff"`
			HurricanePike   int `json:"hurricane_pike"`
			Broadsword      int `json:"broadsword"`
			BladesOfAttack  int `json:"blades_of_attack"`
			LesserCrit      int `json:"lesser_crit"`
			DemonEdge       int `json:"demon_edge"`
			GreaterCrit     int `json:"greater_crit"`
		} `json:"item_win,omitempty"`
		ItemUsage2 struct {
			Circlet         int `json:"circlet"`
			MagicStick      int `json:"magic_stick"`
			Tango           int `json:"tango"`
			Branches        int `json:"branches"`
			BootsOfElves    int `json:"boots_of_elves"`
			Boots           int `json:"boots"`
			Gloves          int `json:"gloves"`
			PowerTreads     int `json:"power_treads"`
			BladeOfAlacrity int `json:"blade_of_alacrity"`
			BeltOfStrength  int `json:"belt_of_strength"`
			Clarity         int `json:"clarity"`
			DragonLance     int `json:"dragon_lance"`
			SobiMask        int `json:"sobi_mask"`
			RingOfBasilius  int `json:"ring_of_basilius"`
			MagicWand       int `json:"magic_wand"`
			BlightStone     int `json:"blight_stone"`
			Tpscroll        int `json:"tpscroll"`
			MithrilHammer   int `json:"mithril_hammer"`
			Desolator       int `json:"desolator"`
			AghanimsShard   int `json:"aghanims_shard"`
			Blink           int `json:"blink"`
			UltimateOrb     int `json:"ultimate_orb"`
			RingOfHealth    int `json:"ring_of_health"`
			Pers            int `json:"pers"`
			VoidStone       int `json:"void_stone"`
			Sphere          int `json:"sphere"`
			StaffOfWizardry int `json:"staff_of_wizardry"`
			FluffyHat       int `json:"fluffy_hat"`
			ForceStaff      int `json:"force_staff"`
			HurricanePike   int `json:"hurricane_pike"`
			Broadsword      int `json:"broadsword"`
			BladesOfAttack  int `json:"blades_of_attack"`
			LesserCrit      int `json:"lesser_crit"`
			DemonEdge       int `json:"demon_edge"`
			GreaterCrit     int `json:"greater_crit"`
		} `json:"item_usage,omitempty"`
		AbilityTargets3 struct {
			DeathProphetSpiritSiphon struct {
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroChen         int `json:"npc_dota_hero_chen"`
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
				NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
			} `json:"death_prophet_spirit_siphon"`
			DeathProphetCarrionSwarm struct {
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
				NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
			} `json:"death_prophet_carrion_swarm"`
		} `json:"ability_targets,omitempty"`
		AbilityUses3 struct {
			DeathProphetCarrionSwarm int `json:"death_prophet_carrion_swarm"`
			DeathProphetSpiritSiphon int `json:"death_prophet_spirit_siphon"`
			DeathProphetExorcism     int `json:"death_prophet_exorcism"`
			DeathProphetSilence      int `json:"death_prophet_silence"`
		} `json:"ability_uses,omitempty"`
		Actions3 struct {
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
			Num14 int `json:"14"`
			Num15 int `json:"15"`
			Num16 int `json:"16"`
			Num19 int `json:"19"`
			Num23 int `json:"23"`
			Num27 int `json:"27"`
			Num33 int `json:"33"`
			Num37 int `json:"37"`
			Num38 int `json:"38"`
		} `json:"actions,omitempty"`
		Damage3 struct {
			NpcDotaHeroTrollWarlord                int `json:"npc_dota_hero_troll_warlord"`
			NpcDotaCreepBadguysRanged              int `json:"npc_dota_creep_badguys_ranged"`
			NpcDotaCreepBadguysMelee               int `json:"npc_dota_creep_badguys_melee"`
			NpcDotaCreepGoodguysRanged             int `json:"npc_dota_creep_goodguys_ranged"`
			NpcDotaCreepGoodguysMelee              int `json:"npc_dota_creep_goodguys_melee"`
			NpcDotaHeroChen                        int `json:"npc_dota_hero_chen"`
			NpcDotaCreepBadguysFlagbearer          int `json:"npc_dota_creep_badguys_flagbearer"`
			NpcDotaCreepGoodguysFlagbearer         int `json:"npc_dota_creep_goodguys_flagbearer"`
			NpcDotaBadguysTower1Top                int `json:"npc_dota_badguys_tower1_top"`
			NpcDotaBadguysSiege                    int `json:"npc_dota_badguys_siege"`
			NpcDotaNeutralSatyrSoulstealer         int `json:"npc_dota_neutral_satyr_soulstealer"`
			NpcDotaNeutralOgreMagi                 int `json:"npc_dota_neutral_ogre_magi"`
			NpcDotaHeroDragonKnight                int `json:"npc_dota_hero_dragon_knight"`
			NpcDotaBadguysTower1Mid                int `json:"npc_dota_badguys_tower1_mid"`
			NpcDotaHeroMirana                      int `json:"npc_dota_hero_mirana"`
			NpcDotaNeutralGiantWolf                int `json:"npc_dota_neutral_giant_wolf"`
			NpcDotaNeutralAlphaWolf                int `json:"npc_dota_neutral_alpha_wolf"`
			NpcDotaNeutralWildkin                  int `json:"npc_dota_neutral_wildkin"`
			NpcDotaNeutralEnragedWildkin           int `json:"npc_dota_neutral_enraged_wildkin"`
			NpcDotaHeroQueenofpain                 int `json:"npc_dota_hero_queenofpain"`
			NpcDotaNeutralWarpineRaider            int `json:"npc_dota_neutral_warpine_raider"`
			NpcDotaNeutralBlackDragon              int `json:"npc_dota_neutral_black_dragon"`
			NpcDotaNeutralRockGolem                int `json:"npc_dota_neutral_rock_golem"`
			NpcDotaNeutralSmallThunderLizard       int `json:"npc_dota_neutral_small_thunder_lizard"`
			NpcDotaNeutralBlackDrake               int `json:"npc_dota_neutral_black_drake"`
			NpcDotaNeutralGraniteGolem             int `json:"npc_dota_neutral_granite_golem"`
			NpcDotaNeutralBigThunderLizard         int `json:"npc_dota_neutral_big_thunder_lizard"`
			NpcDotaNeutralOgreMauler               int `json:"npc_dota_neutral_ogre_mauler"`
			NpcDotaNeutralPolarFurbolgChampion     int `json:"npc_dota_neutral_polar_furbolg_champion"`
			NpcDotaNeutralPolarFurbolgUrsaWarrior  int `json:"npc_dota_neutral_polar_furbolg_ursa_warrior"`
			NpcDotaDarkTrollWarlordSkeletonWarrior int `json:"npc_dota_dark_troll_warlord_skeleton_warrior"`
			NpcDotaNeutralCentaurKhan              int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaNeutralDarkTrollWarlord         int `json:"npc_dota_neutral_dark_troll_warlord"`
			NpcDotaNeutralKobold                   int `json:"npc_dota_neutral_kobold"`
			NpcDotaNeutralCentaurOutrunner         int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaRoshan                          int `json:"npc_dota_roshan"`
			NpcDotaBadguysTower1Bot                int `json:"npc_dota_badguys_tower1_bot"`
			NpcDotaNeutralSatyrHellcaller          int `json:"npc_dota_neutral_satyr_hellcaller"`
			NpcDotaNeutralSatyrTrickster           int `json:"npc_dota_neutral_satyr_trickster"`
			NpcDotaBadguysTower2Mid                int `json:"npc_dota_badguys_tower2_mid"`
			NpcDotaBadguysTower2Top                int `json:"npc_dota_badguys_tower2_top"`
			NpcDotaNeutralMudGolem                 int `json:"npc_dota_neutral_mud_golem"`
			NpcDotaNeutralMudGolemSplit            int `json:"npc_dota_neutral_mud_golem_split"`
			NpcDotaNeutralGhost                    int `json:"npc_dota_neutral_ghost"`
			NpcDotaSentryWards                     int `json:"npc_dota_sentry_wards"`
			NpcDotaNeutralDarkTroll                int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaBadguysTower2Bot                int `json:"npc_dota_badguys_tower2_bot"`
			NpcDotaNeutralKoboldTaskmaster         int `json:"npc_dota_neutral_kobold_taskmaster"`
			NpcDotaNeutralForestTrollBerserker     int `json:"npc_dota_neutral_forest_troll_berserker"`
			NpcDotaCreepBadguysMeleeUpgraded       int `json:"npc_dota_creep_badguys_melee_upgraded"`
			NpcDotaCreepBadguysFlagbearerUpgraded  int `json:"npc_dota_creep_badguys_flagbearer_upgraded"`
			NpcDotaBadguysSiegeUpgraded            int `json:"npc_dota_badguys_siege_upgraded"`
			NpcDotaCreepBadguysRangedUpgraded      int `json:"npc_dota_creep_badguys_ranged_upgraded"`
		} `json:"damage,omitempty"`
		DamageInflictor3 struct {
			Null                     int `json:"null"`
			DeathProphetCarrionSwarm int `json:"death_prophet_carrion_swarm"`
			DeathProphetSpiritSiphon int `json:"death_prophet_spirit_siphon"`
			UrnOfShadows             int `json:"urn_of_shadows"`
			DeathProphetExorcism     int `json:"death_prophet_exorcism"`
			SpiritVessel             int `json:"spirit_vessel"`
			DragonScale              int `json:"dragon_scale"`
		} `json:"damage_inflictor,omitempty"`
		DamageInflictorReceived3 struct {
			Null                           int `json:"null"`
			TrollWarlordWhirlingAxesRanged int `json:"troll_warlord_whirling_axes_ranged"`
			TrollWarlordWhirlingAxesMelee  int `json:"troll_warlord_whirling_axes_melee"`
			SatyrSoulstealerManaBurn       int `json:"satyr_soulstealer_mana_burn"`
			QueenofpainSonicWave           int `json:"queenofpain_sonic_wave"`
			DragonKnightBreatheFire        int `json:"dragon_knight_breathe_fire"`
			MudGolemHurlBoulder            int `json:"mud_golem_hurl_boulder"`
			DragonKnightDragonTail         int `json:"dragon_knight_dragon_tail"`
			MiranaArrow                    int `json:"mirana_arrow"`
			QueenofpainScreamOfPain        int `json:"queenofpain_scream_of_pain"`
			MiranaStarfall                 int `json:"mirana_starfall"`
			QueenofpainShadowStrike        int `json:"queenofpain_shadow_strike"`
			QueenofpainBlink               int `json:"queenofpain_blink"`
			Bfury                          int `json:"bfury"`
			DragonKnightFireball           int `json:"dragon_knight_fireball"`
			WraithPact                     int `json:"wraith_pact"`
			Stormcrafter                   int `json:"stormcrafter"`
			SpiritVessel                   int `json:"spirit_vessel"`
			DragonKnightElderDragonForm    int `json:"dragon_knight_elder_dragon_form"`
		} `json:"damage_inflictor_received,omitempty"`
		DamageTaken3 struct {
			NpcDotaHeroChen                       int `json:"npc_dota_hero_chen"`
			NpcDotaCreepBadguysMelee              int `json:"npc_dota_creep_badguys_melee"`
			NpcDotaHeroTrollWarlord               int `json:"npc_dota_hero_troll_warlord"`
			NpcDotaCreepBadguysRanged             int `json:"npc_dota_creep_badguys_ranged"`
			NpcDotaCreepBadguysFlagbearer         int `json:"npc_dota_creep_badguys_flagbearer"`
			NpcDotaHeroQueenofpain                int `json:"npc_dota_hero_queenofpain"`
			NpcDotaBadguysTower1Top               int `json:"npc_dota_badguys_tower1_top"`
			NpcDotaHeroDragonKnight               int `json:"npc_dota_hero_dragon_knight"`
			NpcDotaBadguysSiege                   int `json:"npc_dota_badguys_siege"`
			NpcDotaNeutralSmallThunderLizard      int `json:"npc_dota_neutral_small_thunder_lizard"`
			NpcDotaNeutralBlackDrake              int `json:"npc_dota_neutral_black_drake"`
			NpcDotaNeutralBlackDragon             int `json:"npc_dota_neutral_black_dragon"`
			NpcDotaNeutralGraniteGolem            int `json:"npc_dota_neutral_granite_golem"`
			NpcDotaNeutralRockGolem               int `json:"npc_dota_neutral_rock_golem"`
			NpcDotaNeutralBigThunderLizard        int `json:"npc_dota_neutral_big_thunder_lizard"`
			NpcDotaNeutralWarpineRaider           int `json:"npc_dota_neutral_warpine_raider"`
			NpcDotaHeroMirana                     int `json:"npc_dota_hero_mirana"`
			NpcDotaNeutralPolarFurbolgChampion    int `json:"npc_dota_neutral_polar_furbolg_champion"`
			NpcDotaNeutralOgreMauler              int `json:"npc_dota_neutral_ogre_mauler"`
			NpcDotaNeutralCentaurOutrunner        int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaRoshan                         int `json:"npc_dota_roshan"`
			NpcDotaBadguysTower1Bot               int `json:"npc_dota_badguys_tower1_bot"`
			NpcDotaNeutralSatyrSoulstealer        int `json:"npc_dota_neutral_satyr_soulstealer"`
			NpcDotaNeutralSatyrTrickster          int `json:"npc_dota_neutral_satyr_trickster"`
			NpcDotaNeutralSatyrHellcaller         int `json:"npc_dota_neutral_satyr_hellcaller"`
			NpcDotaNeutralCentaurKhan             int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaNeutralGiantWolf               int `json:"npc_dota_neutral_giant_wolf"`
			NpcDotaNeutralMudGolem                int `json:"npc_dota_neutral_mud_golem"`
			NpcDotaNeutralDarkTroll               int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaNeutralDarkTrollWarlord        int `json:"npc_dota_neutral_dark_troll_warlord"`
			NpcDotaNeutralForestTrollBerserker    int `json:"npc_dota_neutral_forest_troll_berserker"`
			NpcDotaNeutralKoboldTaskmaster        int `json:"npc_dota_neutral_kobold_taskmaster"`
			NpcDotaNeutralPolarFurbolgUrsaWarrior int `json:"npc_dota_neutral_polar_furbolg_ursa_warrior"`
		} `json:"damage_taken,omitempty"`
		DamageTargets3 struct {
			Null struct {
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
				NpcDotaHeroChen         int `json:"npc_dota_hero_chen"`
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
			} `json:"null"`
			DeathProphetCarrionSwarm struct {
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
				NpcDotaHeroChen         int `json:"npc_dota_hero_chen"`
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
				NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
			} `json:"death_prophet_carrion_swarm"`
			DeathProphetSpiritSiphon struct {
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroChen         int `json:"npc_dota_hero_chen"`
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
				NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
			} `json:"death_prophet_spirit_siphon"`
			UrnOfShadows struct {
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
			} `json:"urn_of_shadows"`
			DeathProphetExorcism struct {
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroChen         int `json:"npc_dota_hero_chen"`
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
				NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
			} `json:"death_prophet_exorcism"`
			SpiritVessel struct {
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
			} `json:"spirit_vessel"`
			DragonScale struct {
				NpcDotaHeroDragonKnight int `json:"npc_dota_hero_dragon_knight"`
				NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
				NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
				NpcDotaHeroMirana       int `json:"npc_dota_hero_mirana"`
			} `json:"dragon_scale"`
		} `json:"damage_targets,omitempty"`
		GoldReasons3 struct {
			Num0  int `json:"0"`
			Num1  int `json:"1"`
			Num6  int `json:"6"`
			Num11 int `json:"11"`
			Num12 int `json:"12"`
			Num13 int `json:"13"`
			Num14 int `json:"14"`
			Num15 int `json:"15"`
			Num16 int `json:"16"`
			Num17 int `json:"17"`
			Num19 int `json:"19"`
		} `json:"gold_reasons,omitempty"`
		HeroHits3 struct {
			Null                     int `json:"null"`
			DeathProphetCarrionSwarm int `json:"death_prophet_carrion_swarm"`
			DeathProphetSpiritSiphon int `json:"death_prophet_spirit_siphon"`
			UrnOfShadows             int `json:"urn_of_shadows"`
			DeathProphetExorcism     int `json:"death_prophet_exorcism"`
			SpiritVessel             int `json:"spirit_vessel"`
			DragonScale              int `json:"dragon_scale"`
		} `json:"hero_hits,omitempty"`
		ItemUses3 struct {
			Tango               int `json:"tango"`
			UrnOfShadows        int `json:"urn_of_shadows"`
			MagicWand           int `json:"magic_wand"`
			Tpscroll            int `json:"tpscroll"`
			PhaseBoots          int `json:"phase_boots"`
			PogoStick           int `json:"pogo_stick"`
			Clarity             int `json:"clarity"`
			SpiritVessel        int `json:"spirit_vessel"`
			BlackKingBar        int `json:"black_king_bar"`
			Dust                int `json:"dust"`
			AghanimsShardRoshan int `json:"aghanims_shard_roshan"`
			HeavensHalberd      int `json:"heavens_halberd"`
			HeavyBlade          int `json:"heavy_blade"`
		} `json:"item_uses,omitempty"`
		Killed3 struct {
			NpcDotaCreepBadguysMelee               int `json:"npc_dota_creep_badguys_melee"`
			NpcDotaCreepGoodguysRanged             int `json:"npc_dota_creep_goodguys_ranged"`
			NpcDotaCreepGoodguysMelee              int `json:"npc_dota_creep_goodguys_melee"`
			NpcDotaCreepBadguysRanged              int `json:"npc_dota_creep_badguys_ranged"`
			NpcDotaCreepBadguysFlagbearer          int `json:"npc_dota_creep_badguys_flagbearer"`
			NpcDotaCreepGoodguysFlagbearer         int `json:"npc_dota_creep_goodguys_flagbearer"`
			NpcDotaHeroTrollWarlord                int `json:"npc_dota_hero_troll_warlord"`
			NpcDotaBadguysTower1Top                int `json:"npc_dota_badguys_tower1_top"`
			NpcDotaHeroDragonKnight                int `json:"npc_dota_hero_dragon_knight"`
			NpcDotaBadguysTower1Mid                int `json:"npc_dota_badguys_tower1_mid"`
			NpcDotaBadguysSiege                    int `json:"npc_dota_badguys_siege"`
			NpcDotaNeutralWarpineRaider            int `json:"npc_dota_neutral_warpine_raider"`
			NpcDotaNeutralRockGolem                int `json:"npc_dota_neutral_rock_golem"`
			NpcDotaNeutralSmallThunderLizard       int `json:"npc_dota_neutral_small_thunder_lizard"`
			NpcDotaNeutralPolarFurbolgChampion     int `json:"npc_dota_neutral_polar_furbolg_champion"`
			NpcDotaDarkTrollWarlordSkeletonWarrior int `json:"npc_dota_dark_troll_warlord_skeleton_warrior"`
			NpcDotaNeutralCentaurKhan              int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaNeutralDarkTrollWarlord         int `json:"npc_dota_neutral_dark_troll_warlord"`
			NpcDotaNeutralKobold                   int `json:"npc_dota_neutral_kobold"`
			NpcDotaNeutralOgreMauler               int `json:"npc_dota_neutral_ogre_mauler"`
			NpcDotaNeutralCentaurOutrunner         int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaNeutralSatyrSoulstealer         int `json:"npc_dota_neutral_satyr_soulstealer"`
			NpcDotaNeutralSatyrTrickster           int `json:"npc_dota_neutral_satyr_trickster"`
			NpcDotaNeutralSatyrHellcaller          int `json:"npc_dota_neutral_satyr_hellcaller"`
			NpcDotaBadguysTower2Top                int `json:"npc_dota_badguys_tower2_top"`
			NpcDotaBadguysTower2Mid                int `json:"npc_dota_badguys_tower2_mid"`
			NpcDotaNeutralGiantWolf                int `json:"npc_dota_neutral_giant_wolf"`
			NpcDotaNeutralAlphaWolf                int `json:"npc_dota_neutral_alpha_wolf"`
			NpcDotaNeutralMudGolem                 int `json:"npc_dota_neutral_mud_golem"`
			NpcDotaNeutralMudGolemSplit            int `json:"npc_dota_neutral_mud_golem_split"`
			NpcDotaSentryWards                     int `json:"npc_dota_sentry_wards"`
			NpcDotaNeutralDarkTroll                int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaNeutralKoboldTaskmaster         int `json:"npc_dota_neutral_kobold_taskmaster"`
			NpcDotaNeutralForestTrollBerserker     int `json:"npc_dota_neutral_forest_troll_berserker"`
			NpcDotaNeutralPolarFurbolgUrsaWarrior  int `json:"npc_dota_neutral_polar_furbolg_ursa_warrior"`
			NpcDotaNeutralEnragedWildkin           int `json:"npc_dota_neutral_enraged_wildkin"`
		} `json:"killed,omitempty"`
		KilledBy2 struct {
			NpcDotaHeroQueenofpain  int `json:"npc_dota_hero_queenofpain"`
			NpcDotaHeroTrollWarlord int `json:"npc_dota_hero_troll_warlord"`
		} `json:"killed_by,omitempty"`
		LanePos3 struct {
			Num72 struct {
				Num78 int `json:"78"`
				Num80 int `json:"80"`
			} `json:"72"`
			Num74 struct {
				Num82  int `json:"82"`
				Num84  int `json:"84"`
				Num138 int `json:"138"`
			} `json:"74"`
			Num76 struct {
				Num86  int `json:"86"`
				Num138 int `json:"138"`
				Num140 int `json:"140"`
				Num142 int `json:"142"`
				Num144 int `json:"144"`
				Num146 int `json:"146"`
				Num148 int `json:"148"`
				Num150 int `json:"150"`
				Num152 int `json:"152"`
				Num154 int `json:"154"`
				Num156 int `json:"156"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
			} `json:"76"`
			Num78 struct {
				Num86  int `json:"86"`
				Num138 int `json:"138"`
				Num140 int `json:"140"`
				Num148 int `json:"148"`
				Num150 int `json:"150"`
				Num152 int `json:"152"`
				Num154 int `json:"154"`
				Num156 int `json:"156"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
			} `json:"78"`
			Num80 struct {
				Num88  int `json:"88"`
				Num150 int `json:"150"`
				Num152 int `json:"152"`
				Num154 int `json:"154"`
				Num156 int `json:"156"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
			} `json:"80"`
			Num82 struct {
				Num90  int `json:"90"`
				Num92  int `json:"92"`
				Num150 int `json:"150"`
				Num152 int `json:"152"`
				Num154 int `json:"154"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
			} `json:"82"`
			Num84 struct {
				Num92  int `json:"92"`
				Num146 int `json:"146"`
				Num148 int `json:"148"`
				Num150 int `json:"150"`
				Num152 int `json:"152"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
				Num174 int `json:"174"`
			} `json:"84"`
			Num86 struct {
				Num92  int `json:"92"`
				Num146 int `json:"146"`
				Num152 int `json:"152"`
				Num162 int `json:"162"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
				Num176 int `json:"176"`
			} `json:"86"`
			Num88 struct {
				Num92  int `json:"92"`
				Num146 int `json:"146"`
				Num154 int `json:"154"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
				Num178 int `json:"178"`
				Num182 int `json:"182"`
			} `json:"88"`
			Num90 struct {
				Num96  int `json:"96"`
				Num144 int `json:"144"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
				Num176 int `json:"176"`
				Num178 int `json:"178"`
				Num180 int `json:"180"`
				Num182 int `json:"182"`
			} `json:"90"`
			Num92 struct {
				Num96  int `json:"96"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
				Num180 int `json:"180"`
				Num182 int `json:"182"`
			} `json:"92"`
			Num94 struct {
				Num98  int `json:"98"`
				Num144 int `json:"144"`
				Num172 int `json:"172"`
				Num174 int `json:"174"`
			} `json:"94"`
			Num96 struct {
				Num100 int `json:"100"`
				Num144 int `json:"144"`
				Num172 int `json:"172"`
				Num174 int `json:"174"`
				Num176 int `json:"176"`
			} `json:"96"`
			Num98 struct {
				Num100 int `json:"100"`
				Num142 int `json:"142"`
				Num168 int `json:"168"`
				Num174 int `json:"174"`
				Num176 int `json:"176"`
			} `json:"98"`
			Num100 struct {
				Num102 int `json:"102"`
				Num140 int `json:"140"`
				Num142 int `json:"142"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
				Num176 int `json:"176"`
			} `json:"100"`
			Num102 struct {
				Num104 int `json:"104"`
				Num138 int `json:"138"`
				Num170 int `json:"170"`
			} `json:"102"`
			Num104 struct {
				Num106 int `json:"106"`
				Num138 int `json:"138"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
				Num176 int `json:"176"`
			} `json:"104"`
			Num106 struct {
				Num108 int `json:"108"`
				Num110 int `json:"110"`
				Num136 int `json:"136"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
				Num174 int `json:"174"`
				Num176 int `json:"176"`
			} `json:"106"`
			Num108 struct {
				Num112 int `json:"112"`
				Num134 int `json:"134"`
				Num158 int `json:"158"`
				Num174 int `json:"174"`
			} `json:"108"`
			Num110 struct {
				Num114 int `json:"114"`
				Num128 int `json:"128"`
				Num132 int `json:"132"`
			} `json:"110"`
			Num112 struct {
				Num116 int `json:"116"`
				Num128 int `json:"128"`
				Num130 int `json:"130"`
				Num132 int `json:"132"`
				Num134 int `json:"134"`
				Num158 int `json:"158"`
			} `json:"112"`
			Num114 struct {
				Num118 int `json:"118"`
				Num120 int `json:"120"`
				Num122 int `json:"122"`
				Num124 int `json:"124"`
				Num126 int `json:"126"`
				Num128 int `json:"128"`
				Num132 int `json:"132"`
				Num134 int `json:"134"`
				Num156 int `json:"156"`
			} `json:"114"`
			Num116 struct {
				Num120 int `json:"120"`
				Num124 int `json:"124"`
			} `json:"116"`
			Num118 struct {
				Num118 int `json:"118"`
				Num120 int `json:"120"`
				Num122 int `json:"122"`
				Num156 int `json:"156"`
			} `json:"118"`
			Num120 struct {
				Num118 int `json:"118"`
				Num154 int `json:"154"`
			} `json:"120"`
			Num122 struct {
				Num152 int `json:"152"`
			} `json:"122"`
			Num124 struct {
				Num152 int `json:"152"`
			} `json:"124"`
			Num126 struct {
				Num146 int `json:"146"`
				Num150 int `json:"150"`
			} `json:"126"`
			Num128 struct {
				Num130 int `json:"130"`
				Num132 int `json:"132"`
				Num134 int `json:"134"`
				Num138 int `json:"138"`
				Num140 int `json:"140"`
				Num144 int `json:"144"`
				Num146 int `json:"146"`
			} `json:"128"`
			Num130 struct {
				Num130 int `json:"130"`
				Num132 int `json:"132"`
				Num142 int `json:"142"`
			} `json:"130"`
			Num132 struct {
				Num130 int `json:"130"`
				Num140 int `json:"140"`
			} `json:"132"`
			Num134 struct {
				Num130 int `json:"130"`
				Num138 int `json:"138"`
			} `json:"134"`
			Num136 struct {
				Num134 int `json:"134"`
			} `json:"136"`
			Num138 struct {
				Num128 int `json:"128"`
			} `json:"138"`
			Num140 struct {
				Num126 int `json:"126"`
				Num132 int `json:"132"`
			} `json:"140"`
			Num142 struct {
				Num130 int `json:"130"`
			} `json:"142"`
			Num144 struct {
				Num126 int `json:"126"`
			} `json:"144"`
			Num146 struct {
				Num128 int `json:"128"`
			} `json:"146"`
		} `json:"lane_pos,omitempty"`
		Obs3 struct {
		} `json:"obs,omitempty"`
		Purchase3 struct {
			Circlet              int `json:"circlet"`
			Branches             int `json:"branches"`
			Tango                int `json:"tango"`
			MagicStick           int `json:"magic_stick"`
			RecipeMagicWand      int `json:"recipe_magic_wand"`
			MagicWand            int `json:"magic_wand"`
			SobiMask             int `json:"sobi_mask"`
			RingOfProtection     int `json:"ring_of_protection"`
			RecipeUrnOfShadows   int `json:"recipe_urn_of_shadows"`
			UrnOfShadows         int `json:"urn_of_shadows"`
			Boots                int `json:"boots"`
			Chainmail            int `json:"chainmail"`
			BladesOfAttack       int `json:"blades_of_attack"`
			PhaseBoots           int `json:"phase_boots"`
			Tpscroll             int `json:"tpscroll"`
			VitalityBooster      int `json:"vitality_booster"`
			Clarity              int `json:"clarity"`
			RecipeSpiritVessel   int `json:"recipe_spirit_vessel"`
			InfusedRaindrop      int `json:"infused_raindrop"`
			SpiritVessel         int `json:"spirit_vessel"`
			OgreAxe              int `json:"ogre_axe"`
			Dust                 int `json:"dust"`
			MithrilHammer        int `json:"mithril_hammer"`
			RecipeBlackKingBar   int `json:"recipe_black_king_bar"`
			BlackKingBar         int `json:"black_king_bar"`
			BeltOfStrength       int `json:"belt_of_strength"`
			Sange                int `json:"sange"`
			RecipeSange          int `json:"recipe_sange"`
			RecipeHeavensHalberd int `json:"recipe_heavens_halberd"`
			TalismanOfEvasion    int `json:"talisman_of_evasion"`
			HeavensHalberd       int `json:"heavens_halberd"`
			Platemail            int `json:"platemail"`
		} `json:"purchase,omitempty"`
		Sen3 struct {
		} `json:"sen,omitempty"`
		XpReasons3 struct {
			Num0 int `json:"0"`
			Num1 int `json:"1"`
			Num2 int `json:"2"`
			Num3 int `json:"3"`
			Num5 int `json:"5"`
		} `json:"xp_reasons,omitempty"`
		PurchaseTime3 struct {
			Circlet           int `json:"circlet"`
			Branches          int `json:"branches"`
			Tango             int `json:"tango"`
			MagicStick        int `json:"magic_stick"`
			MagicWand         int `json:"magic_wand"`
			SobiMask          int `json:"sobi_mask"`
			RingOfProtection  int `json:"ring_of_protection"`
			UrnOfShadows      int `json:"urn_of_shadows"`
			Boots             int `json:"boots"`
			Chainmail         int `json:"chainmail"`
			BladesOfAttack    int `json:"blades_of_attack"`
			PhaseBoots        int `json:"phase_boots"`
			Tpscroll          int `json:"tpscroll"`
			VitalityBooster   int `json:"vitality_booster"`
			Clarity           int `json:"clarity"`
			InfusedRaindrop   int `json:"infused_raindrop"`
			SpiritVessel      int `json:"spirit_vessel"`
			OgreAxe           int `json:"ogre_axe"`
			Dust              int `json:"dust"`
			MithrilHammer     int `json:"mithril_hammer"`
			BlackKingBar      int `json:"black_king_bar"`
			BeltOfStrength    int `json:"belt_of_strength"`
			Sange             int `json:"sange"`
			TalismanOfEvasion int `json:"talisman_of_evasion"`
			HeavensHalberd    int `json:"heavens_halberd"`
			Platemail         int `json:"platemail"`
		} `json:"purchase_time,omitempty"`
		FirstPurchaseTime3 struct {
			Circlet           int `json:"circlet"`
			Branches          int `json:"branches"`
			Tango             int `json:"tango"`
			MagicStick        int `json:"magic_stick"`
			MagicWand         int `json:"magic_wand"`
			SobiMask          int `json:"sobi_mask"`
			RingOfProtection  int `json:"ring_of_protection"`
			UrnOfShadows      int `json:"urn_of_shadows"`
			Boots             int `json:"boots"`
			Chainmail         int `json:"chainmail"`
			BladesOfAttack    int `json:"blades_of_attack"`
			PhaseBoots        int `json:"phase_boots"`
			Tpscroll          int `json:"tpscroll"`
			VitalityBooster   int `json:"vitality_booster"`
			Clarity           int `json:"clarity"`
			InfusedRaindrop   int `json:"infused_raindrop"`
			SpiritVessel      int `json:"spirit_vessel"`
			OgreAxe           int `json:"ogre_axe"`
			Dust              int `json:"dust"`
			MithrilHammer     int `json:"mithril_hammer"`
			BlackKingBar      int `json:"black_king_bar"`
			BeltOfStrength    int `json:"belt_of_strength"`
			Sange             int `json:"sange"`
			TalismanOfEvasion int `json:"talisman_of_evasion"`
			HeavensHalberd    int `json:"heavens_halberd"`
			Platemail         int `json:"platemail"`
		} `json:"first_purchase_time,omitempty"`
		ItemWin3 struct {
			Circlet           int `json:"circlet"`
			Branches          int `json:"branches"`
			Tango             int `json:"tango"`
			MagicStick        int `json:"magic_stick"`
			MagicWand         int `json:"magic_wand"`
			SobiMask          int `json:"sobi_mask"`
			RingOfProtection  int `json:"ring_of_protection"`
			UrnOfShadows      int `json:"urn_of_shadows"`
			Boots             int `json:"boots"`
			Chainmail         int `json:"chainmail"`
			BladesOfAttack    int `json:"blades_of_attack"`
			PhaseBoots        int `json:"phase_boots"`
			Tpscroll          int `json:"tpscroll"`
			VitalityBooster   int `json:"vitality_booster"`
			Clarity           int `json:"clarity"`
			InfusedRaindrop   int `json:"infused_raindrop"`
			SpiritVessel      int `json:"spirit_vessel"`
			OgreAxe           int `json:"ogre_axe"`
			Dust              int `json:"dust"`
			MithrilHammer     int `json:"mithril_hammer"`
			BlackKingBar      int `json:"black_king_bar"`
			BeltOfStrength    int `json:"belt_of_strength"`
			Sange             int `json:"sange"`
			TalismanOfEvasion int `json:"talisman_of_evasion"`
			HeavensHalberd    int `json:"heavens_halberd"`
			Platemail         int `json:"platemail"`
		} `json:"item_win,omitempty"`
		ItemUsage3 struct {
			Circlet           int `json:"circlet"`
			Branches          int `json:"branches"`
			Tango             int `json:"tango"`
			MagicStick        int `json:"magic_stick"`
			MagicWand         int `json:"magic_wand"`
			SobiMask          int `json:"sobi_mask"`
			RingOfProtection  int `json:"ring_of_protection"`
			UrnOfShadows      int `json:"urn_of_shadows"`
			Boots             int `json:"boots"`
			Chainmail         int `json:"chainmail"`
			BladesOfAttack    int `json:"blades_of_attack"`
			PhaseBoots        int `json:"phase_boots"`
			Tpscroll          int `json:"tpscroll"`
			VitalityBooster   int `json:"vitality_booster"`
			Clarity           int `json:"clarity"`
			InfusedRaindrop   int `json:"infused_raindrop"`
			SpiritVessel      int `json:"spirit_vessel"`
			OgreAxe           int `json:"ogre_axe"`
			Dust              int `json:"dust"`
			MithrilHammer     int `json:"mithril_hammer"`
			BlackKingBar      int `json:"black_king_bar"`
			BeltOfStrength    int `json:"belt_of_strength"`
			Sange             int `json:"sange"`
			TalismanOfEvasion int `json:"talisman_of_evasion"`
			HeavensHalberd    int `json:"heavens_halberd"`
			Platemail         int `json:"platemail"`
		} `json:"item_usage,omitempty"`
		AbilityTargets4 struct {
		} `json:"ability_targets,omitempty"`
		AbilityUses4 struct {
			MiranaLeap     int `json:"mirana_leap"`
			MiranaArrow    int `json:"mirana_arrow"`
			MiranaStarfall int `json:"mirana_starfall"`
			MiranaInvis    int `json:"mirana_invis"`
			AbilityCapture int `json:"ability_capture"`
		} `json:"ability_uses,omitempty"`
		Actions4 struct {
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
			Num13 int `json:"13"`
			Num14 int `json:"14"`
			Num15 int `json:"15"`
			Num16 int `json:"16"`
			Num17 int `json:"17"`
			Num18 int `json:"18"`
			Num19 int `json:"19"`
			Num31 int `json:"31"`
			Num32 int `json:"32"`
			Num33 int `json:"33"`
			Num37 int `json:"37"`
			Num38 int `json:"38"`
		} `json:"actions,omitempty"`
		Damage4 struct {
			NpcDotaHeroEnchantress                int `json:"npc_dota_hero_enchantress"`
			NpcDotaHeroTemplarAssassin            int `json:"npc_dota_hero_templar_assassin"`
			NpcDotaCourier                        int `json:"npc_dota_courier"`
			NpcDotaCreepBadguysMelee              int `json:"npc_dota_creep_badguys_melee"`
			NpcDotaNeutralPolarFurbolgUrsaWarrior int `json:"npc_dota_neutral_polar_furbolg_ursa_warrior"`
			NpcDotaNeutralPolarFurbolgChampion    int `json:"npc_dota_neutral_polar_furbolg_champion"`
			NpcDotaNeutralDarkTrollWarlord        int `json:"npc_dota_neutral_dark_troll_warlord"`
			NpcDotaNeutralDarkTroll               int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaCreepBadguysFlagbearer         int `json:"npc_dota_creep_badguys_flagbearer"`
			NpcDotaNeutralWarpineRaider           int `json:"npc_dota_neutral_warpine_raider"`
			NpcDotaCreepGoodguysMelee             int `json:"npc_dota_creep_goodguys_melee"`
			NpcDotaNeutralRockGolem               int `json:"npc_dota_neutral_rock_golem"`
			NpcDotaSentryWards                    int `json:"npc_dota_sentry_wards"`
			NpcDotaGoodguysSiege                  int `json:"npc_dota_goodguys_siege"`
			NpcDotaHeroPuck                       int `json:"npc_dota_hero_puck"`
			NpcDotaCreepGoodguysFlagbearer        int `json:"npc_dota_creep_goodguys_flagbearer"`
			NpcDotaCreepGoodguysRanged            int `json:"npc_dota_creep_goodguys_ranged"`
			NpcDotaNeutralBlackDragon             int `json:"npc_dota_neutral_black_dragon"`
			NpcDotaHeroDeathProphet               int `json:"npc_dota_hero_death_prophet"`
			NpcDotaTemplarAssassinPsionicTrap     int `json:"npc_dota_templar_assassin_psionic_trap"`
			NpcDotaNeutralEnragedWildkin          int `json:"npc_dota_neutral_enraged_wildkin"`
			NpcDotaNeutralSatyrHellcaller         int `json:"npc_dota_neutral_satyr_hellcaller"`
			NpcDotaNeutralHarpyScout              int `json:"npc_dota_neutral_harpy_scout"`
			NpcDotaNeutralHarpyStorm              int `json:"npc_dota_neutral_harpy_storm"`
			NpcDotaHeroPrimalBeast                int `json:"npc_dota_hero_primal_beast"`
			NpcDotaNeutralWildkin                 int `json:"npc_dota_neutral_wildkin"`
			NpcDotaNeutralCentaurKhan             int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaNeutralMudGolemSplit           int `json:"npc_dota_neutral_mud_golem_split"`
			NpcDotaNeutralCentaurOutrunner        int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaNeutralKoboldTaskmaster        int `json:"npc_dota_neutral_kobold_taskmaster"`
			NpcDotaNeutralKobold                  int `json:"npc_dota_neutral_kobold"`
			NpcDotaNeutralKoboldTunneler          int `json:"npc_dota_neutral_kobold_tunneler"`
			NpcDotaObserverWards                  int `json:"npc_dota_observer_wards"`
			NpcDotaRoshan                         int `json:"npc_dota_roshan"`
			NpcDotaNeutralSatyrSoulstealer        int `json:"npc_dota_neutral_satyr_soulstealer"`
			NpcDotaNeutralSatyrTrickster          int `json:"npc_dota_neutral_satyr_trickster"`
			NpcDotaNeutralOgreMagi                int `json:"npc_dota_neutral_ogre_magi"`
			NpcDotaNeutralOgreMauler              int `json:"npc_dota_neutral_ogre_mauler"`
			NpcDotaGoodguysRangeRaxBot            int `json:"npc_dota_goodguys_range_rax_bot"`
			NpcDotaGoodguysRangeRaxTop            int `json:"npc_dota_goodguys_range_rax_top"`
			NpcDotaGoodguysMeleeRaxTop            int `json:"npc_dota_goodguys_melee_rax_top"`
			NpcDotaGoodguysFillers                int `json:"npc_dota_goodguys_fillers"`
		} `json:"damage,omitempty"`
		DamageInflictor4 struct {
			Null           int `json:"null"`
			MiranaArrow    int `json:"mirana_arrow"`
			MiranaStarfall int `json:"mirana_starfall"`
			UrnOfShadows   int `json:"urn_of_shadows"`
			SpiritVessel   int `json:"spirit_vessel"`
			Stormcrafter   int `json:"stormcrafter"`
		} `json:"damage_inflictor,omitempty"`
		DamageInflictorReceived4 struct {
			TemplarAssassinMeld      int `json:"templar_assassin_meld"`
			Null                     int `json:"null"`
			DeathProphetCarrionSwarm int `json:"death_prophet_carrion_swarm"`
			PuckDreamCoil            int `json:"puck_dream_coil"`
			PuckWaningRift           int `json:"puck_waning_rift"`
			EnchantressImpetus       int `json:"enchantress_impetus"`
			PuckIllusoryOrb          int `json:"puck_illusory_orb"`
			WitchBlade               int `json:"witch_blade"`
			DeathProphetSpiritSiphon int `json:"death_prophet_spirit_siphon"`
			SpiritVessel             int `json:"spirit_vessel"`
			PrimalBeastOnslaught     int `json:"primal_beast_onslaught"`
			PrimalBeastPulverize     int `json:"primal_beast_pulverize"`
			Dust                     int `json:"dust"`
			SatyrHellcallerShockwave int `json:"satyr_hellcaller_shockwave"`
			PrimalBeastTrample       int `json:"primal_beast_trample"`
			DragonScale              int `json:"dragon_scale"`
			PrimalBeastRockThrow     int `json:"primal_beast_rock_throw"`
			DeathProphetExorcism     int `json:"death_prophet_exorcism"`
			TemplarAssassinTrap      int `json:"templar_assassin_trap"`
			OverwhelmingBlink        int `json:"overwhelming_blink"`
			TemplarAssassinPsiBlades int `json:"templar_assassin_psi_blades"`
		} `json:"damage_inflictor_received,omitempty"`
		DamageTaken4 struct {
			NpcDotaHeroTemplarAssassin     int `json:"npc_dota_hero_templar_assassin"`
			NpcDotaHeroEnchantress         int `json:"npc_dota_hero_enchantress"`
			NpcDotaNeutralDarkTroll        int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaNeutralDarkTrollWarlord int `json:"npc_dota_neutral_dark_troll_warlord"`
			NpcDotaCreepGoodguysMelee      int `json:"npc_dota_creep_goodguys_melee"`
			NpcDotaCreepGoodguysRanged     int `json:"npc_dota_creep_goodguys_ranged"`
			NpcDotaHeroDeathProphet        int `json:"npc_dota_hero_death_prophet"`
			NpcDotaHeroPuck                int `json:"npc_dota_hero_puck"`
			NpcDotaCreepGoodguysFlagbearer int `json:"npc_dota_creep_goodguys_flagbearer"`
			NpcDotaHeroPrimalBeast         int `json:"npc_dota_hero_primal_beast"`
			NpcDotaEnragedWildkinTornado   int `json:"npc_dota_enraged_wildkin_tornado"`
			NpcDotaNeutralCentaurOutrunner int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaNeutralCentaurKhan      int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaNeutralKoboldTaskmaster int `json:"npc_dota_neutral_kobold_taskmaster"`
			NpcDotaNeutralKoboldTunneler   int `json:"npc_dota_neutral_kobold_tunneler"`
			NpcDotaNeutralKobold           int `json:"npc_dota_neutral_kobold"`
			NpcDotaNeutralOgreMauler       int `json:"npc_dota_neutral_ogre_mauler"`
			NpcDotaNeutralSatyrTrickster   int `json:"npc_dota_neutral_satyr_trickster"`
			NpcDotaNeutralGraniteGolem     int `json:"npc_dota_neutral_granite_golem"`
			NpcDotaNeutralRockGolem        int `json:"npc_dota_neutral_rock_golem"`
			NpcDotaNeutralOgreMagi         int `json:"npc_dota_neutral_ogre_magi"`
		} `json:"damage_taken,omitempty"`
		DamageTargets4 struct {
			Null struct {
				NpcDotaHeroEnchantress     int `json:"npc_dota_hero_enchantress"`
				NpcDotaHeroTemplarAssassin int `json:"npc_dota_hero_templar_assassin"`
				NpcDotaHeroPuck            int `json:"npc_dota_hero_puck"`
				NpcDotaHeroDeathProphet    int `json:"npc_dota_hero_death_prophet"`
				NpcDotaHeroPrimalBeast     int `json:"npc_dota_hero_primal_beast"`
			} `json:"null"`
			MiranaArrow struct {
				NpcDotaHeroPuck            int `json:"npc_dota_hero_puck"`
				NpcDotaHeroEnchantress     int `json:"npc_dota_hero_enchantress"`
				NpcDotaHeroDeathProphet    int `json:"npc_dota_hero_death_prophet"`
				NpcDotaHeroPrimalBeast     int `json:"npc_dota_hero_primal_beast"`
				NpcDotaHeroTemplarAssassin int `json:"npc_dota_hero_templar_assassin"`
			} `json:"mirana_arrow"`
			MiranaStarfall struct {
				NpcDotaHeroEnchantress  int `json:"npc_dota_hero_enchantress"`
				NpcDotaHeroDeathProphet int `json:"npc_dota_hero_death_prophet"`
				NpcDotaHeroPrimalBeast  int `json:"npc_dota_hero_primal_beast"`
				NpcDotaHeroPuck         int `json:"npc_dota_hero_puck"`
			} `json:"mirana_starfall"`
			UrnOfShadows struct {
				NpcDotaHeroPrimalBeast int `json:"npc_dota_hero_primal_beast"`
				NpcDotaHeroEnchantress int `json:"npc_dota_hero_enchantress"`
			} `json:"urn_of_shadows"`
			SpiritVessel struct {
				NpcDotaHeroPrimalBeast  int `json:"npc_dota_hero_primal_beast"`
				NpcDotaHeroPuck         int `json:"npc_dota_hero_puck"`
				NpcDotaHeroEnchantress  int `json:"npc_dota_hero_enchantress"`
				NpcDotaHeroDeathProphet int `json:"npc_dota_hero_death_prophet"`
			} `json:"spirit_vessel"`
			Stormcrafter struct {
				NpcDotaHeroPuck         int `json:"npc_dota_hero_puck"`
				NpcDotaHeroPrimalBeast  int `json:"npc_dota_hero_primal_beast"`
				NpcDotaHeroEnchantress  int `json:"npc_dota_hero_enchantress"`
				NpcDotaHeroDeathProphet int `json:"npc_dota_hero_death_prophet"`
			} `json:"stormcrafter"`
		} `json:"damage_targets,omitempty"`
		HeroHits4 struct {
			Null           int `json:"null"`
			MiranaArrow    int `json:"mirana_arrow"`
			MiranaStarfall int `json:"mirana_starfall"`
			UrnOfShadows   int `json:"urn_of_shadows"`
			SpiritVessel   int `json:"spirit_vessel"`
			Stormcrafter   int `json:"stormcrafter"`
		} `json:"hero_hits,omitempty"`
		ItemUses4 struct {
			WardDispenser   int `json:"ward_dispenser"`
			Branches        int `json:"branches"`
			Tango           int `json:"tango"`
			WardObserver    int `json:"ward_observer"`
			Clarity         int `json:"clarity"`
			EnchantedMango  int `json:"enchanted_mango"`
			WardSentry      int `json:"ward_sentry"`
			Tpscroll        int `json:"tpscroll"`
			SmokeOfDeceit   int `json:"smoke_of_deceit"`
			UnstableWand    int `json:"unstable_wand"`
			UrnOfShadows    int `json:"urn_of_shadows"`
			TomeOfKnowledge int `json:"tome_of_knowledge"`
			SpiritVessel    int `json:"spirit_vessel"`
			Gem             int `json:"gem"`
			ArcaneBoots     int `json:"arcane_boots"`
			MagicWand       int `json:"magic_wand"`
			LotusOrb        int `json:"lotus_orb"`
			Dust            int `json:"dust"`
			Cyclone         int `json:"cyclone"`
			Stormcrafter    int `json:"stormcrafter"`
			WindWaker       int `json:"wind_waker"`
		} `json:"item_uses,omitempty"`
		Killed4 struct {
			NpcDotaNeutralPolarFurbolgUrsaWarrior int `json:"npc_dota_neutral_polar_furbolg_ursa_warrior"`
			NpcDotaNeutralPolarFurbolgChampion    int `json:"npc_dota_neutral_polar_furbolg_champion"`
			NpcDotaNeutralDarkTrollWarlord        int `json:"npc_dota_neutral_dark_troll_warlord"`
			NpcDotaNeutralDarkTroll               int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaCreepBadguysFlagbearer         int `json:"npc_dota_creep_badguys_flagbearer"`
			NpcDotaNeutralWarpineRaider           int `json:"npc_dota_neutral_warpine_raider"`
			NpcDotaSentryWards                    int `json:"npc_dota_sentry_wards"`
			NpcDotaGoodguysSiege                  int `json:"npc_dota_goodguys_siege"`
			NpcDotaCreepGoodguysMelee             int `json:"npc_dota_creep_goodguys_melee"`
			NpcDotaCreepGoodguysFlagbearer        int `json:"npc_dota_creep_goodguys_flagbearer"`
			NpcDotaCreepGoodguysRanged            int `json:"npc_dota_creep_goodguys_ranged"`
			NpcDotaTemplarAssassinPsionicTrap     int `json:"npc_dota_templar_assassin_psionic_trap"`
			NpcDotaNeutralEnragedWildkin          int `json:"npc_dota_neutral_enraged_wildkin"`
			NpcDotaNeutralSatyrHellcaller         int `json:"npc_dota_neutral_satyr_hellcaller"`
			NpcDotaNeutralHarpyScout              int `json:"npc_dota_neutral_harpy_scout"`
			NpcDotaNeutralHarpyStorm              int `json:"npc_dota_neutral_harpy_storm"`
			NpcDotaHeroEnchantress                int `json:"npc_dota_hero_enchantress"`
			NpcDotaNeutralCentaurOutrunner        int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaNeutralCentaurKhan             int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaNeutralKoboldTaskmaster        int `json:"npc_dota_neutral_kobold_taskmaster"`
			NpcDotaNeutralKobold                  int `json:"npc_dota_neutral_kobold"`
			NpcDotaObserverWards                  int `json:"npc_dota_observer_wards"`
			NpcDotaHeroPrimalBeast                int `json:"npc_dota_hero_primal_beast"`
			NpcDotaGoodguysRangeRaxTop            int `json:"npc_dota_goodguys_range_rax_top"`
		} `json:"killed,omitempty"`
		KilledBy3 struct {
			NpcDotaHeroPuck        int `json:"npc_dota_hero_puck"`
			NpcDotaHeroPrimalBeast int `json:"npc_dota_hero_primal_beast"`
		} `json:"killed_by,omitempty"`
		LanePos4 struct {
			Num112 struct {
				Num154 int `json:"154"`
			} `json:"112"`
			Num114 struct {
				Num154 int `json:"154"`
				Num156 int `json:"156"`
			} `json:"114"`
			Num116 struct {
				Num156 int `json:"156"`
			} `json:"116"`
			Num118 struct {
				Num154 int `json:"154"`
				Num156 int `json:"156"`
			} `json:"118"`
			Num120 struct {
				Num154 int `json:"154"`
				Num156 int `json:"156"`
			} `json:"120"`
			Num122 struct {
				Num154 int `json:"154"`
				Num156 int `json:"156"`
				Num158 int `json:"158"`
			} `json:"122"`
			Num124 struct {
				Num152 int `json:"152"`
				Num158 int `json:"158"`
			} `json:"124"`
			Num126 struct {
				Num150 int `json:"150"`
				Num160 int `json:"160"`
			} `json:"126"`
			Num128 struct {
				Num148 int `json:"148"`
				Num162 int `json:"162"`
			} `json:"128"`
			Num130 struct {
				Num138 int `json:"138"`
				Num140 int `json:"140"`
				Num142 int `json:"142"`
				Num144 int `json:"144"`
				Num146 int `json:"146"`
				Num162 int `json:"162"`
			} `json:"130"`
			Num132 struct {
				Num116 int `json:"116"`
				Num118 int `json:"118"`
				Num120 int `json:"120"`
				Num124 int `json:"124"`
				Num134 int `json:"134"`
				Num136 int `json:"136"`
				Num164 int `json:"164"`
			} `json:"132"`
			Num134 struct {
				Num118 int `json:"118"`
				Num120 int `json:"120"`
				Num124 int `json:"124"`
				Num132 int `json:"132"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
			} `json:"134"`
			Num136 struct {
				Num120 int `json:"120"`
				Num122 int `json:"122"`
				Num124 int `json:"124"`
				Num128 int `json:"128"`
				Num130 int `json:"130"`
				Num168 int `json:"168"`
			} `json:"136"`
			Num138 struct {
				Num118 int `json:"118"`
				Num122 int `json:"122"`
				Num124 int `json:"124"`
				Num126 int `json:"126"`
				Num128 int `json:"128"`
				Num134 int `json:"134"`
				Num170 int `json:"170"`
			} `json:"138"`
			Num140 struct {
				Num116 int `json:"116"`
				Num124 int `json:"124"`
				Num126 int `json:"126"`
				Num128 int `json:"128"`
				Num132 int `json:"132"`
				Num134 int `json:"134"`
				Num136 int `json:"136"`
				Num138 int `json:"138"`
				Num170 int `json:"170"`
			} `json:"140"`
			Num142 struct {
				Num102 int `json:"102"`
				Num104 int `json:"104"`
				Num106 int `json:"106"`
				Num108 int `json:"108"`
				Num110 int `json:"110"`
				Num112 int `json:"112"`
				Num114 int `json:"114"`
				Num116 int `json:"116"`
				Num118 int `json:"118"`
				Num120 int `json:"120"`
				Num122 int `json:"122"`
				Num124 int `json:"124"`
				Num126 int `json:"126"`
				Num132 int `json:"132"`
				Num136 int `json:"136"`
				Num138 int `json:"138"`
				Num140 int `json:"140"`
				Num170 int `json:"170"`
			} `json:"142"`
			Num144 struct {
				Num102 int `json:"102"`
				Num104 int `json:"104"`
				Num106 int `json:"106"`
				Num108 int `json:"108"`
				Num114 int `json:"114"`
				Num118 int `json:"118"`
				Num120 int `json:"120"`
				Num122 int `json:"122"`
				Num124 int `json:"124"`
				Num128 int `json:"128"`
				Num130 int `json:"130"`
				Num134 int `json:"134"`
				Num136 int `json:"136"`
				Num138 int `json:"138"`
				Num140 int `json:"140"`
				Num170 int `json:"170"`
			} `json:"144"`
			Num146 struct {
				Num102 int `json:"102"`
				Num110 int `json:"110"`
				Num112 int `json:"112"`
				Num114 int `json:"114"`
				Num116 int `json:"116"`
				Num118 int `json:"118"`
				Num122 int `json:"122"`
				Num128 int `json:"128"`
				Num130 int `json:"130"`
				Num132 int `json:"132"`
				Num138 int `json:"138"`
				Num170 int `json:"170"`
			} `json:"146"`
			Num148 struct {
				Num90  int `json:"90"`
				Num100 int `json:"100"`
				Num114 int `json:"114"`
				Num116 int `json:"116"`
				Num118 int `json:"118"`
				Num122 int `json:"122"`
				Num126 int `json:"126"`
				Num128 int `json:"128"`
				Num172 int `json:"172"`
			} `json:"148"`
			Num150 struct {
				Num92  int `json:"92"`
				Num100 int `json:"100"`
				Num116 int `json:"116"`
				Num118 int `json:"118"`
				Num120 int `json:"120"`
				Num122 int `json:"122"`
				Num124 int `json:"124"`
				Num126 int `json:"126"`
				Num128 int `json:"128"`
				Num172 int `json:"172"`
			} `json:"150"`
			Num152 struct {
				Num86  int `json:"86"`
				Num92  int `json:"92"`
				Num94  int `json:"94"`
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num108 int `json:"108"`
				Num110 int `json:"110"`
				Num118 int `json:"118"`
				Num120 int `json:"120"`
				Num122 int `json:"122"`
				Num124 int `json:"124"`
				Num126 int `json:"126"`
				Num172 int `json:"172"`
			} `json:"152"`
			Num154 struct {
				Num82  int `json:"82"`
				Num84  int `json:"84"`
				Num88  int `json:"88"`
				Num90  int `json:"90"`
				Num94  int `json:"94"`
				Num96  int `json:"96"`
				Num106 int `json:"106"`
				Num108 int `json:"108"`
				Num118 int `json:"118"`
				Num122 int `json:"122"`
				Num126 int `json:"126"`
				Num174 int `json:"174"`
			} `json:"154"`
			Num156 struct {
				Num80  int `json:"80"`
				Num82  int `json:"82"`
				Num84  int `json:"84"`
				Num90  int `json:"90"`
				Num92  int `json:"92"`
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num102 int `json:"102"`
				Num104 int `json:"104"`
				Num106 int `json:"106"`
				Num120 int `json:"120"`
				Num122 int `json:"122"`
				Num124 int `json:"124"`
				Num126 int `json:"126"`
				Num176 int `json:"176"`
			} `json:"156"`
			Num158 struct {
				Num86  int `json:"86"`
				Num92  int `json:"92"`
				Num94  int `json:"94"`
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num118 int `json:"118"`
				Num120 int `json:"120"`
				Num122 int `json:"122"`
				Num124 int `json:"124"`
				Num176 int `json:"176"`
			} `json:"158"`
			Num160 struct {
				Num86  int `json:"86"`
				Num88  int `json:"88"`
				Num90  int `json:"90"`
				Num92  int `json:"92"`
				Num94  int `json:"94"`
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num106 int `json:"106"`
				Num116 int `json:"116"`
				Num118 int `json:"118"`
				Num176 int `json:"176"`
			} `json:"160"`
			Num162 struct {
				Num88  int `json:"88"`
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num104 int `json:"104"`
				Num106 int `json:"106"`
				Num108 int `json:"108"`
				Num110 int `json:"110"`
				Num112 int `json:"112"`
				Num114 int `json:"114"`
				Num118 int `json:"118"`
				Num176 int `json:"176"`
			} `json:"162"`
			Num164 struct {
				Num86  int `json:"86"`
				Num88  int `json:"88"`
				Num94  int `json:"94"`
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num104 int `json:"104"`
				Num106 int `json:"106"`
				Num108 int `json:"108"`
				Num114 int `json:"114"`
				Num116 int `json:"116"`
				Num174 int `json:"174"`
			} `json:"164"`
			Num166 struct {
				Num88  int `json:"88"`
				Num92  int `json:"92"`
				Num94  int `json:"94"`
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num104 int `json:"104"`
				Num112 int `json:"112"`
				Num114 int `json:"114"`
			} `json:"166"`
			Num168 struct {
				Num88  int `json:"88"`
				Num90  int `json:"90"`
				Num92  int `json:"92"`
				Num94  int `json:"94"`
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num102 int `json:"102"`
				Num110 int `json:"110"`
				Num112 int `json:"112"`
				Num114 int `json:"114"`
				Num174 int `json:"174"`
			} `json:"168"`
			Num170 struct {
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
				Num114 int `json:"114"`
				Num174 int `json:"174"`
			} `json:"170"`
			Num172 struct {
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
				Num114 int `json:"114"`
				Num174 int `json:"174"`
			} `json:"172"`
			Num174 struct {
				Num92  int `json:"92"`
				Num94  int `json:"94"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num102 int `json:"102"`
				Num104 int `json:"104"`
				Num106 int `json:"106"`
				Num108 int `json:"108"`
				Num110 int `json:"110"`
				Num112 int `json:"112"`
				Num114 int `json:"114"`
				Num174 int `json:"174"`
			} `json:"174"`
			Num176 struct {
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num102 int `json:"102"`
				Num108 int `json:"108"`
				Num110 int `json:"110"`
				Num174 int `json:"174"`
			} `json:"176"`
			Num178 struct {
				Num96  int `json:"96"`
				Num100 int `json:"100"`
				Num108 int `json:"108"`
				Num110 int `json:"110"`
				Num176 int `json:"176"`
			} `json:"178"`
			Num180 struct {
				Num178 int `json:"178"`
			} `json:"180"`
		} `json:"lane_pos,omitempty"`
		Obs4 struct {
			Num68 struct {
				Num130 int `json:"130"`
			} `json:"68"`
			Num72 struct {
				Num104 int `json:"104"`
			} `json:"72"`
			Num84 struct {
				Num100 int `json:"100"`
			} `json:"84"`
			Num94 struct {
				Num140 int `json:"140"`
			} `json:"94"`
			Num100 struct {
				Num78 int `json:"78"`
			} `json:"100"`
			Num106 struct {
				Num184 int `json:"184"`
			} `json:"106"`
			Num122 struct {
				Num144 int `json:"144"`
			} `json:"122"`
			Num132 struct {
				Num124 int `json:"124"`
			} `json:"132"`
			Num144 struct {
				Num90  int `json:"90"`
				Num120 int `json:"120"`
			} `json:"144"`
			Num160 struct {
				Num120 int `json:"120"`
			} `json:"160"`
		} `json:"obs,omitempty"`
		Purchase4 struct {
			WardDispenser      int `json:"ward_dispenser"`
			Branches           int `json:"branches"`
			Tango              int `json:"tango"`
			RingOfProtection   int `json:"ring_of_protection"`
			SobiMask           int `json:"sobi_mask"`
			WardObserver       int `json:"ward_observer"`
			WardSentry         int `json:"ward_sentry"`
			Clarity            int `json:"clarity"`
			EnchantedMango     int `json:"enchanted_mango"`
			Circlet            int `json:"circlet"`
			RecipeUrnOfShadows int `json:"recipe_urn_of_shadows"`
			SmokeOfDeceit      int `json:"smoke_of_deceit"`
			UrnOfShadows       int `json:"urn_of_shadows"`
			Boots              int `json:"boots"`
			MagicStick         int `json:"magic_stick"`
			MagicWand          int `json:"magic_wand"`
			RecipeMagicWand    int `json:"recipe_magic_wand"`
			VitalityBooster    int `json:"vitality_booster"`
			RecipeSpiritVessel int `json:"recipe_spirit_vessel"`
			SpiritVessel       int `json:"spirit_vessel"`
			TomeOfKnowledge    int `json:"tome_of_knowledge"`
			Gem                int `json:"gem"`
			EnergyBooster      int `json:"energy_booster"`
			ArcaneBoots        int `json:"arcane_boots"`
			Tpscroll           int `json:"tpscroll"`
			Platemail          int `json:"platemail"`
			RingOfHealth       int `json:"ring_of_health"`
			Pers               int `json:"pers"`
			VoidStone          int `json:"void_stone"`
			LotusOrb           int `json:"lotus_orb"`
			WindLace           int `json:"wind_lace"`
			Dust               int `json:"dust"`
			StaffOfWizardry    int `json:"staff_of_wizardry"`
			RecipeCyclone      int `json:"recipe_cyclone"`
			Cyclone            int `json:"cyclone"`
			RecipeWindWaker    int `json:"recipe_wind_waker"`
			MysticStaff        int `json:"mystic_staff"`
			WindWaker          int `json:"wind_waker"`
		} `json:"purchase,omitempty"`
		Sen4 struct {
			Num66 struct {
				Num132 int `json:"132"`
			} `json:"66"`
			Num68 struct {
				Num130 int `json:"130"`
			} `json:"68"`
			Num100 struct {
				Num78 int `json:"78"`
			} `json:"100"`
			Num108 struct {
				Num182 int `json:"182"`
			} `json:"108"`
			Num126 struct {
				Num146 int `json:"146"`
				Num164 int `json:"164"`
			} `json:"126"`
			Num130 struct {
				Num122 int `json:"122"`
			} `json:"130"`
			Num142 struct {
				Num122 int `json:"122"`
			} `json:"142"`
			Num148 struct {
				Num124 int `json:"124"`
				Num164 int `json:"164"`
			} `json:"148"`
			Num150 struct {
				Num124 int `json:"124"`
			} `json:"150"`
			Num152 struct {
				Num94 int `json:"94"`
			} `json:"152"`
			Num156 struct {
				Num134 int `json:"134"`
			} `json:"156"`
			Num172 struct {
				Num92 int `json:"92"`
			} `json:"172"`
		} `json:"sen,omitempty"`
		XpReasons4 struct {
			Num0 int `json:"0"`
			Num1 int `json:"1"`
			Num2 int `json:"2"`
			Num3 int `json:"3"`
			Num4 int `json:"4"`
			Num5 int `json:"5"`
		} `json:"xp_reasons,omitempty"`
		PurchaseTime4 struct {
			Branches         int `json:"branches"`
			Tango            int `json:"tango"`
			RingOfProtection int `json:"ring_of_protection"`
			SobiMask         int `json:"sobi_mask"`
			WardObserver     int `json:"ward_observer"`
			WardSentry       int `json:"ward_sentry"`
			Clarity          int `json:"clarity"`
			EnchantedMango   int `json:"enchanted_mango"`
			Circlet          int `json:"circlet"`
			SmokeOfDeceit    int `json:"smoke_of_deceit"`
			UrnOfShadows     int `json:"urn_of_shadows"`
			Boots            int `json:"boots"`
			MagicStick       int `json:"magic_stick"`
			MagicWand        int `json:"magic_wand"`
			VitalityBooster  int `json:"vitality_booster"`
			SpiritVessel     int `json:"spirit_vessel"`
			TomeOfKnowledge  int `json:"tome_of_knowledge"`
			Gem              int `json:"gem"`
			EnergyBooster    int `json:"energy_booster"`
			ArcaneBoots      int `json:"arcane_boots"`
			Tpscroll         int `json:"tpscroll"`
			Platemail        int `json:"platemail"`
			RingOfHealth     int `json:"ring_of_health"`
			Pers             int `json:"pers"`
			VoidStone        int `json:"void_stone"`
			LotusOrb         int `json:"lotus_orb"`
			WindLace         int `json:"wind_lace"`
			Dust             int `json:"dust"`
			StaffOfWizardry  int `json:"staff_of_wizardry"`
			Cyclone          int `json:"cyclone"`
			MysticStaff      int `json:"mystic_staff"`
			WindWaker        int `json:"wind_waker"`
		} `json:"purchase_time,omitempty"`
		FirstPurchaseTime4 struct {
			Branches         int `json:"branches"`
			Tango            int `json:"tango"`
			RingOfProtection int `json:"ring_of_protection"`
			SobiMask         int `json:"sobi_mask"`
			WardObserver     int `json:"ward_observer"`
			WardSentry       int `json:"ward_sentry"`
			Clarity          int `json:"clarity"`
			EnchantedMango   int `json:"enchanted_mango"`
			Circlet          int `json:"circlet"`
			SmokeOfDeceit    int `json:"smoke_of_deceit"`
			UrnOfShadows     int `json:"urn_of_shadows"`
			Boots            int `json:"boots"`
			MagicStick       int `json:"magic_stick"`
			MagicWand        int `json:"magic_wand"`
			VitalityBooster  int `json:"vitality_booster"`
			SpiritVessel     int `json:"spirit_vessel"`
			TomeOfKnowledge  int `json:"tome_of_knowledge"`
			Gem              int `json:"gem"`
			EnergyBooster    int `json:"energy_booster"`
			ArcaneBoots      int `json:"arcane_boots"`
			Tpscroll         int `json:"tpscroll"`
			Platemail        int `json:"platemail"`
			RingOfHealth     int `json:"ring_of_health"`
			Pers             int `json:"pers"`
			VoidStone        int `json:"void_stone"`
			LotusOrb         int `json:"lotus_orb"`
			WindLace         int `json:"wind_lace"`
			Dust             int `json:"dust"`
			StaffOfWizardry  int `json:"staff_of_wizardry"`
			Cyclone          int `json:"cyclone"`
			MysticStaff      int `json:"mystic_staff"`
			WindWaker        int `json:"wind_waker"`
		} `json:"first_purchase_time,omitempty"`
		ItemWin4 struct {
			Branches         int `json:"branches"`
			Tango            int `json:"tango"`
			RingOfProtection int `json:"ring_of_protection"`
			SobiMask         int `json:"sobi_mask"`
			WardObserver     int `json:"ward_observer"`
			WardSentry       int `json:"ward_sentry"`
			Clarity          int `json:"clarity"`
			EnchantedMango   int `json:"enchanted_mango"`
			Circlet          int `json:"circlet"`
			SmokeOfDeceit    int `json:"smoke_of_deceit"`
			UrnOfShadows     int `json:"urn_of_shadows"`
			Boots            int `json:"boots"`
			MagicStick       int `json:"magic_stick"`
			MagicWand        int `json:"magic_wand"`
			VitalityBooster  int `json:"vitality_booster"`
			SpiritVessel     int `json:"spirit_vessel"`
			TomeOfKnowledge  int `json:"tome_of_knowledge"`
			Gem              int `json:"gem"`
			EnergyBooster    int `json:"energy_booster"`
			ArcaneBoots      int `json:"arcane_boots"`
			Tpscroll         int `json:"tpscroll"`
			Platemail        int `json:"platemail"`
			RingOfHealth     int `json:"ring_of_health"`
			Pers             int `json:"pers"`
			VoidStone        int `json:"void_stone"`
			LotusOrb         int `json:"lotus_orb"`
			WindLace         int `json:"wind_lace"`
			Dust             int `json:"dust"`
			StaffOfWizardry  int `json:"staff_of_wizardry"`
			Cyclone          int `json:"cyclone"`
			MysticStaff      int `json:"mystic_staff"`
			WindWaker        int `json:"wind_waker"`
		} `json:"item_win,omitempty"`
		ItemUsage4 struct {
			Branches         int `json:"branches"`
			Tango            int `json:"tango"`
			RingOfProtection int `json:"ring_of_protection"`
			SobiMask         int `json:"sobi_mask"`
			WardObserver     int `json:"ward_observer"`
			WardSentry       int `json:"ward_sentry"`
			Clarity          int `json:"clarity"`
			EnchantedMango   int `json:"enchanted_mango"`
			Circlet          int `json:"circlet"`
			SmokeOfDeceit    int `json:"smoke_of_deceit"`
			UrnOfShadows     int `json:"urn_of_shadows"`
			Boots            int `json:"boots"`
			MagicStick       int `json:"magic_stick"`
			MagicWand        int `json:"magic_wand"`
			VitalityBooster  int `json:"vitality_booster"`
			SpiritVessel     int `json:"spirit_vessel"`
			TomeOfKnowledge  int `json:"tome_of_knowledge"`
			Gem              int `json:"gem"`
			EnergyBooster    int `json:"energy_booster"`
			ArcaneBoots      int `json:"arcane_boots"`
			Tpscroll         int `json:"tpscroll"`
			Platemail        int `json:"platemail"`
			RingOfHealth     int `json:"ring_of_health"`
			Pers             int `json:"pers"`
			VoidStone        int `json:"void_stone"`
			LotusOrb         int `json:"lotus_orb"`
			WindLace         int `json:"wind_lace"`
			Dust             int `json:"dust"`
			StaffOfWizardry  int `json:"staff_of_wizardry"`
			Cyclone          int `json:"cyclone"`
			MysticStaff      int `json:"mystic_staff"`
			WindWaker        int `json:"wind_waker"`
		} `json:"item_usage,omitempty"`
		AbilityTargets5 struct {
			QueenofpainShadowStrike struct {
				NpcDotaHeroEnchantress     int `json:"npc_dota_hero_enchantress"`
				NpcDotaHeroDeathProphet    int `json:"npc_dota_hero_death_prophet"`
				NpcDotaHeroPrimalBeast     int `json:"npc_dota_hero_primal_beast"`
				NpcDotaHeroTemplarAssassin int `json:"npc_dota_hero_templar_assassin"`
				NpcDotaHeroPuck            int `json:"npc_dota_hero_puck"`
			} `json:"queenofpain_shadow_strike"`
		} `json:"ability_targets,omitempty"`
		AbilityUses5 struct {
			QueenofpainScreamOfPain int `json:"queenofpain_scream_of_pain"`
			QueenofpainBlink        int `json:"queenofpain_blink"`
			QueenofpainSonicWave    int `json:"queenofpain_sonic_wave"`
			QueenofpainShadowStrike int `json:"queenofpain_shadow_strike"`
			AbilityCapture          int `json:"ability_capture"`
		} `json:"ability_uses,omitempty"`
		Actions5 struct {
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
			Num12 int `json:"12"`
			Num13 int `json:"13"`
			Num14 int `json:"14"`
			Num15 int `json:"15"`
			Num16 int `json:"16"`
			Num17 int `json:"17"`
			Num19 int `json:"19"`
			Num21 int `json:"21"`
			Num27 int `json:"27"`
			Num28 int `json:"28"`
			Num31 int `json:"31"`
			Num37 int `json:"37"`
			Num38 int `json:"38"`
		} `json:"actions,omitempty"`
		Damage5 struct {
			NpcDotaCreepGoodguysRanged             int `json:"npc_dota_creep_goodguys_ranged"`
			NpcDotaCreepGoodguysMelee              int `json:"npc_dota_creep_goodguys_melee"`
			NpcDotaCreepBadguysMelee               int `json:"npc_dota_creep_badguys_melee"`
			NpcDotaHeroPuck                        int `json:"npc_dota_hero_puck"`
			NpcDotaCreepBadguysRanged              int `json:"npc_dota_creep_badguys_ranged"`
			NpcDotaCreepGoodguysFlagbearer         int `json:"npc_dota_creep_goodguys_flagbearer"`
			NpcDotaCreepBadguysFlagbearer          int `json:"npc_dota_creep_badguys_flagbearer"`
			NpcDotaNeutralGnollAssassin            int `json:"npc_dota_neutral_gnoll_assassin"`
			NpcDotaHeroDeathProphet                int `json:"npc_dota_hero_death_prophet"`
			NpcDotaHeroPrimalBeast                 int `json:"npc_dota_hero_primal_beast"`
			NpcDotaCourier                         int `json:"npc_dota_courier"`
			NpcDotaNeutralPolarFurbolgUrsaWarrior  int `json:"npc_dota_neutral_polar_furbolg_ursa_warrior"`
			NpcDotaGoodguysTower1Bot               int `json:"npc_dota_goodguys_tower1_bot"`
			NpcDotaNeutralMudGolemSplit            int `json:"npc_dota_neutral_mud_golem_split"`
			NpcDotaHeroEnchantress                 int `json:"npc_dota_hero_enchantress"`
			NpcDotaNeutralSatyrTrickster           int `json:"npc_dota_neutral_satyr_trickster"`
			NpcDotaNeutralSatyrSoulstealer         int `json:"npc_dota_neutral_satyr_soulstealer"`
			NpcDotaGoodguysSiege                   int `json:"npc_dota_goodguys_siege"`
			NpcDotaNeutralSatyrHellcaller          int `json:"npc_dota_neutral_satyr_hellcaller"`
			NpcDotaNeutralMudGolem                 int `json:"npc_dota_neutral_mud_golem"`
			NpcDotaNeutralCentaurOutrunner         int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaNeutralCentaurKhan              int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaNeutralEnragedWildkin           int `json:"npc_dota_neutral_enraged_wildkin"`
			NpcDotaNeutralWildkin                  int `json:"npc_dota_neutral_wildkin"`
			NpcDotaHeroTemplarAssassin             int `json:"npc_dota_hero_templar_assassin"`
			NpcDotaNeutralDarkTrollWarlord         int `json:"npc_dota_neutral_dark_troll_warlord"`
			NpcDotaDarkTrollWarlordSkeletonWarrior int `json:"npc_dota_dark_troll_warlord_skeleton_warrior"`
			NpcDotaNeutralDarkTroll                int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaTemplarAssassinPsionicTrap      int `json:"npc_dota_templar_assassin_psionic_trap"`
			NpcDotaNeutralAlphaWolf                int `json:"npc_dota_neutral_alpha_wolf"`
			NpcDotaNeutralGiantWolf                int `json:"npc_dota_neutral_giant_wolf"`
			NpcDotaGoodguysTower2Bot               int `json:"npc_dota_goodguys_tower2_bot"`
			NpcDotaNeutralBigThunderLizard         int `json:"npc_dota_neutral_big_thunder_lizard"`
			NpcDotaNeutralSmallThunderLizard       int `json:"npc_dota_neutral_small_thunder_lizard"`
			NpcDotaObserverWards                   int `json:"npc_dota_observer_wards"`
			NpcDotaSentryWards                     int `json:"npc_dota_sentry_wards"`
			NpcDotaNeutralPolarFurbolgChampion     int `json:"npc_dota_neutral_polar_furbolg_champion"`
			NpcDotaNeutralOgreMauler               int `json:"npc_dota_neutral_ogre_mauler"`
			NpcDotaNeutralOgreMagi                 int `json:"npc_dota_neutral_ogre_magi"`
			NpcDotaNeutralWarpineRaider            int `json:"npc_dota_neutral_warpine_raider"`
			NpcDotaNeutralKoboldTaskmaster         int `json:"npc_dota_neutral_kobold_taskmaster"`
			NpcDotaNeutralKobold                   int `json:"npc_dota_neutral_kobold"`
			NpcDotaNeutralKoboldTunneler           int `json:"npc_dota_neutral_kobold_tunneler"`
			NpcDotaNeutralRockGolem                int `json:"npc_dota_neutral_rock_golem"`
			NpcDotaGoodguysTower3Mid               int `json:"npc_dota_goodguys_tower3_mid"`
			NpcDotaGoodguysMeleeRaxMid             int `json:"npc_dota_goodguys_melee_rax_mid"`
			NpcDotaGoodguysRangeRaxMid             int `json:"npc_dota_goodguys_range_rax_mid"`
			NpcDotaRoshan                          int `json:"npc_dota_roshan"`
			NpcDotaCreepBadguysFlagbearerUpgraded  int `json:"npc_dota_creep_badguys_flagbearer_upgraded"`
			NpcDotaGoodguysRangeRaxBot             int `json:"npc_dota_goodguys_range_rax_bot"`
			IllusionNpcDotaHeroQueenofpain         int `json:"illusion_npc_dota_hero_queenofpain"`
			NpcDotaCreepBadguysRangedUpgraded      int `json:"npc_dota_creep_badguys_ranged_upgraded"`
			NpcDotaCreepBadguysMeleeUpgraded       int `json:"npc_dota_creep_badguys_melee_upgraded"`
			NpcDotaGoodguysMeleeRaxTop             int `json:"npc_dota_goodguys_melee_rax_top"`
			NpcDotaGoodguysFillers                 int `json:"npc_dota_goodguys_fillers"`
			NpcDotaGoodguysTower4                  int `json:"npc_dota_goodguys_tower4"`
		} `json:"damage,omitempty"`
		DamageInflictor5 struct {
			Null                    int `json:"null"`
			QueenofpainScreamOfPain int `json:"queenofpain_scream_of_pain"`
			QueenofpainSonicWave    int `json:"queenofpain_sonic_wave"`
			Orchid                  int `json:"orchid"`
			QueenofpainShadowStrike int `json:"queenofpain_shadow_strike"`
			QueenofpainBlink        int `json:"queenofpain_blink"`
			ShivasGuard             int `json:"shivas_guard"`
			Bloodthorn              int `json:"bloodthorn"`
		} `json:"damage_inflictor,omitempty"`
		DamageInflictorReceived5 struct {
			PuckIllusoryOrb                    int `json:"puck_illusory_orb"`
			Null                               int `json:"null"`
			TemplarAssassinPsiBlades           int `json:"templar_assassin_psi_blades"`
			PolarFurbolgUrsaWarriorThunderClap int `json:"polar_furbolg_ursa_warrior_thunder_clap"`
			PuckDreamCoil                      int `json:"puck_dream_coil"`
			WitchBlade                         int `json:"witch_blade"`
			PrimalBeastPulverize               int `json:"primal_beast_pulverize"`
			TemplarAssassinTrap                int `json:"templar_assassin_trap"`
			PuckWaningRift                     int `json:"puck_waning_rift"`
			SatyrHellcallerShockwave           int `json:"satyr_hellcaller_shockwave"`
			PrimalBeastOnslaught               int `json:"primal_beast_onslaught"`
			PrimalBeastTrample                 int `json:"primal_beast_trample"`
			DragonScale                        int `json:"dragon_scale"`
			DeathProphetSpiritSiphon           int `json:"death_prophet_spirit_siphon"`
			PrimalBeastRockThrow               int `json:"primal_beast_rock_throw"`
			Stormcrafter                       int `json:"stormcrafter"`
			TemplarAssassinMeld                int `json:"templar_assassin_meld"`
			DeathProphetExorcism               int `json:"death_prophet_exorcism"`
			DeathProphetCarrionSwarm           int `json:"death_prophet_carrion_swarm"`
		} `json:"damage_inflictor_received,omitempty"`
		DamageTaken5 struct {
			NpcDotaHeroPuck                        int `json:"npc_dota_hero_puck"`
			NpcDotaCreepGoodguysMelee              int `json:"npc_dota_creep_goodguys_melee"`
			NpcDotaCreepGoodguysRanged             int `json:"npc_dota_creep_goodguys_ranged"`
			NpcDotaCreepGoodguysFlagbearer         int `json:"npc_dota_creep_goodguys_flagbearer"`
			NpcDotaHeroEnchantress                 int `json:"npc_dota_hero_enchantress"`
			NpcDotaGoodguysTower1Bot               int `json:"npc_dota_goodguys_tower1_bot"`
			NpcDotaNeutralMudGolemSplit            int `json:"npc_dota_neutral_mud_golem_split"`
			NpcDotaHeroTemplarAssassin             int `json:"npc_dota_hero_templar_assassin"`
			NpcDotaHeroDeathProphet                int `json:"npc_dota_hero_death_prophet"`
			NpcDotaNeutralSatyrTrickster           int `json:"npc_dota_neutral_satyr_trickster"`
			NpcDotaNeutralSatyrSoulstealer         int `json:"npc_dota_neutral_satyr_soulstealer"`
			NpcDotaHeroPrimalBeast                 int `json:"npc_dota_hero_primal_beast"`
			NpcDotaGoodguysSiege                   int `json:"npc_dota_goodguys_siege"`
			NpcDotaNeutralSatyrHellcaller          int `json:"npc_dota_neutral_satyr_hellcaller"`
			NpcDotaNeutralMudGolem                 int `json:"npc_dota_neutral_mud_golem"`
			NpcDotaNeutralEnragedWildkin           int `json:"npc_dota_neutral_enraged_wildkin"`
			NpcDotaNeutralWildkin                  int `json:"npc_dota_neutral_wildkin"`
			NpcDotaNeutralDarkTroll                int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaDarkTrollWarlordSkeletonWarrior int `json:"npc_dota_dark_troll_warlord_skeleton_warrior"`
			NpcDotaNeutralDarkTrollWarlord         int `json:"npc_dota_neutral_dark_troll_warlord"`
			NpcDotaNeutralGiantWolf                int `json:"npc_dota_neutral_giant_wolf"`
			NpcDotaNeutralAlphaWolf                int `json:"npc_dota_neutral_alpha_wolf"`
			NpcDotaGoodguysTower2Bot               int `json:"npc_dota_goodguys_tower2_bot"`
			NpcDotaNeutralSmallThunderLizard       int `json:"npc_dota_neutral_small_thunder_lizard"`
			NpcDotaNeutralBigThunderLizard         int `json:"npc_dota_neutral_big_thunder_lizard"`
			NpcDotaNeutralCentaurKhan              int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaNeutralCentaurOutrunner         int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaNeutralPolarFurbolgChampion     int `json:"npc_dota_neutral_polar_furbolg_champion"`
			NpcDotaNeutralPolarFurbolgUrsaWarrior  int `json:"npc_dota_neutral_polar_furbolg_ursa_warrior"`
			NpcDotaNeutralOgreMauler               int `json:"npc_dota_neutral_ogre_mauler"`
			NpcDotaNeutralWarpineRaider            int `json:"npc_dota_neutral_warpine_raider"`
			NpcDotaNeutralKoboldTaskmaster         int `json:"npc_dota_neutral_kobold_taskmaster"`
			NpcDotaNeutralKoboldTunneler           int `json:"npc_dota_neutral_kobold_tunneler"`
			NpcDotaNeutralKobold                   int `json:"npc_dota_neutral_kobold"`
			NpcDotaNeutralRockGolem                int `json:"npc_dota_neutral_rock_golem"`
			NpcDotaGoodguysTower3Mid               int `json:"npc_dota_goodguys_tower3_mid"`
			NpcDotaGoodguysTower2Top               int `json:"npc_dota_goodguys_tower2_top"`
			NpcDotaGoodguysTower4                  int `json:"npc_dota_goodguys_tower4"`
		} `json:"damage_taken,omitempty"`
		DamageTargets5 struct {
			Null struct {
				NpcDotaHeroPuck            int `json:"npc_dota_hero_puck"`
				NpcDotaHeroDeathProphet    int `json:"npc_dota_hero_death_prophet"`
				NpcDotaHeroPrimalBeast     int `json:"npc_dota_hero_primal_beast"`
				NpcDotaHeroEnchantress     int `json:"npc_dota_hero_enchantress"`
				NpcDotaHeroTemplarAssassin int `json:"npc_dota_hero_templar_assassin"`
			} `json:"null"`
			QueenofpainScreamOfPain struct {
				NpcDotaHeroPuck         int `json:"npc_dota_hero_puck"`
				NpcDotaHeroEnchantress  int `json:"npc_dota_hero_enchantress"`
				NpcDotaHeroDeathProphet int `json:"npc_dota_hero_death_prophet"`
				NpcDotaHeroPrimalBeast  int `json:"npc_dota_hero_primal_beast"`
			} `json:"queenofpain_scream_of_pain"`
			QueenofpainSonicWave struct {
				NpcDotaHeroDeathProphet    int `json:"npc_dota_hero_death_prophet"`
				NpcDotaHeroPrimalBeast     int `json:"npc_dota_hero_primal_beast"`
				NpcDotaHeroEnchantress     int `json:"npc_dota_hero_enchantress"`
				NpcDotaHeroPuck            int `json:"npc_dota_hero_puck"`
				NpcDotaHeroTemplarAssassin int `json:"npc_dota_hero_templar_assassin"`
			} `json:"queenofpain_sonic_wave"`
			Orchid struct {
				NpcDotaHeroPuck            int `json:"npc_dota_hero_puck"`
				NpcDotaHeroTemplarAssassin int `json:"npc_dota_hero_templar_assassin"`
				NpcDotaHeroPrimalBeast     int `json:"npc_dota_hero_primal_beast"`
				NpcDotaHeroEnchantress     int `json:"npc_dota_hero_enchantress"`
			} `json:"orchid"`
			QueenofpainShadowStrike struct {
				NpcDotaHeroEnchantress     int `json:"npc_dota_hero_enchantress"`
				NpcDotaHeroDeathProphet    int `json:"npc_dota_hero_death_prophet"`
				NpcDotaHeroPrimalBeast     int `json:"npc_dota_hero_primal_beast"`
				NpcDotaHeroTemplarAssassin int `json:"npc_dota_hero_templar_assassin"`
				NpcDotaHeroPuck            int `json:"npc_dota_hero_puck"`
			} `json:"queenofpain_shadow_strike"`
			QueenofpainBlink struct {
				NpcDotaHeroEnchantress     int `json:"npc_dota_hero_enchantress"`
				NpcDotaHeroDeathProphet    int `json:"npc_dota_hero_death_prophet"`
				NpcDotaHeroPrimalBeast     int `json:"npc_dota_hero_primal_beast"`
				NpcDotaHeroPuck            int `json:"npc_dota_hero_puck"`
				NpcDotaHeroTemplarAssassin int `json:"npc_dota_hero_templar_assassin"`
			} `json:"queenofpain_blink"`
			ShivasGuard struct {
				NpcDotaHeroEnchantress     int `json:"npc_dota_hero_enchantress"`
				NpcDotaHeroTemplarAssassin int `json:"npc_dota_hero_templar_assassin"`
				NpcDotaHeroPuck            int `json:"npc_dota_hero_puck"`
			} `json:"shivas_guard"`
			Bloodthorn struct {
				NpcDotaHeroEnchantress int `json:"npc_dota_hero_enchantress"`
			} `json:"bloodthorn"`
		} `json:"damage_targets,omitempty"`
		GoldReasons4 struct {
			Num0  int `json:"0"`
			Num6  int `json:"6"`
			Num11 int `json:"11"`
			Num12 int `json:"12"`
			Num13 int `json:"13"`
			Num14 int `json:"14"`
			Num15 int `json:"15"`
			Num16 int `json:"16"`
			Num17 int `json:"17"`
			Num19 int `json:"19"`
			Num21 int `json:"21"`
		} `json:"gold_reasons,omitempty"`
		HeroHits5 struct {
			Null                    int `json:"null"`
			QueenofpainScreamOfPain int `json:"queenofpain_scream_of_pain"`
			QueenofpainSonicWave    int `json:"queenofpain_sonic_wave"`
			Orchid                  int `json:"orchid"`
			QueenofpainShadowStrike int `json:"queenofpain_shadow_strike"`
			QueenofpainBlink        int `json:"queenofpain_blink"`
			ShivasGuard             int `json:"shivas_guard"`
			Bloodthorn              int `json:"bloodthorn"`
		} `json:"hero_hits,omitempty"`
		ItemUses5 struct {
			Branches      int `json:"branches"`
			Tango         int `json:"tango"`
			Bottle        int `json:"bottle"`
			MagicWand     int `json:"magic_wand"`
			Clarity       int `json:"clarity"`
			Tpscroll      int `json:"tpscroll"`
			PowerTreads   int `json:"power_treads"`
			ArcaneRing    int `json:"arcane_ring"`
			WardObserver  int `json:"ward_observer"`
			Orchid        int `json:"orchid"`
			WardDispenser int `json:"ward_dispenser"`
			BlackKingBar  int `json:"black_king_bar"`
			Gem           int `json:"gem"`
			ShivasGuard   int `json:"shivas_guard"`
			Sheepstick    int `json:"sheepstick"`
			Bloodthorn    int `json:"bloodthorn"`
		} `json:"item_uses,omitempty"`
		KillStreaks2 struct {
			Num3  int `json:"3"`
			Num4  int `json:"4"`
			Num5  int `json:"5"`
			Num6  int `json:"6"`
			Num7  int `json:"7"`
			Num8  int `json:"8"`
			Num9  int `json:"9"`
			Num10 int `json:"10"`
			Num11 int `json:"11"`
		} `json:"kill_streaks,omitempty"`
		Killed5 struct {
			NpcDotaCreepGoodguysMelee              int `json:"npc_dota_creep_goodguys_melee"`
			NpcDotaCreepGoodguysRanged             int `json:"npc_dota_creep_goodguys_ranged"`
			NpcDotaCreepBadguysMelee               int `json:"npc_dota_creep_badguys_melee"`
			NpcDotaCreepBadguysRanged              int `json:"npc_dota_creep_badguys_ranged"`
			NpcDotaCreepGoodguysFlagbearer         int `json:"npc_dota_creep_goodguys_flagbearer"`
			NpcDotaCreepBadguysFlagbearer          int `json:"npc_dota_creep_badguys_flagbearer"`
			NpcDotaNeutralGnollAssassin            int `json:"npc_dota_neutral_gnoll_assassin"`
			NpcDotaCourier                         int `json:"npc_dota_courier"`
			NpcDotaNeutralMudGolemSplit            int `json:"npc_dota_neutral_mud_golem_split"`
			NpcDotaHeroEnchantress                 int `json:"npc_dota_hero_enchantress"`
			NpcDotaNeutralSatyrSoulstealer         int `json:"npc_dota_neutral_satyr_soulstealer"`
			NpcDotaNeutralSatyrTrickster           int `json:"npc_dota_neutral_satyr_trickster"`
			NpcDotaGoodguysSiege                   int `json:"npc_dota_goodguys_siege"`
			NpcDotaHeroDeathProphet                int `json:"npc_dota_hero_death_prophet"`
			NpcDotaHeroPrimalBeast                 int `json:"npc_dota_hero_primal_beast"`
			NpcDotaNeutralSatyrHellcaller          int `json:"npc_dota_neutral_satyr_hellcaller"`
			NpcDotaNeutralMudGolem                 int `json:"npc_dota_neutral_mud_golem"`
			NpcDotaNeutralWildkin                  int `json:"npc_dota_neutral_wildkin"`
			NpcDotaNeutralEnragedWildkin           int `json:"npc_dota_neutral_enraged_wildkin"`
			NpcDotaHeroTemplarAssassin             int `json:"npc_dota_hero_templar_assassin"`
			NpcDotaNeutralDarkTrollWarlord         int `json:"npc_dota_neutral_dark_troll_warlord"`
			NpcDotaDarkTrollWarlordSkeletonWarrior int `json:"npc_dota_dark_troll_warlord_skeleton_warrior"`
			NpcDotaNeutralDarkTroll                int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaTemplarAssassinPsionicTrap      int `json:"npc_dota_templar_assassin_psionic_trap"`
			NpcDotaNeutralAlphaWolf                int `json:"npc_dota_neutral_alpha_wolf"`
			NpcDotaNeutralGiantWolf                int `json:"npc_dota_neutral_giant_wolf"`
			NpcDotaGoodguysTower2Bot               int `json:"npc_dota_goodguys_tower2_bot"`
			NpcDotaNeutralBigThunderLizard         int `json:"npc_dota_neutral_big_thunder_lizard"`
			NpcDotaNeutralSmallThunderLizard       int `json:"npc_dota_neutral_small_thunder_lizard"`
			NpcDotaObserverWards                   int `json:"npc_dota_observer_wards"`
			NpcDotaSentryWards                     int `json:"npc_dota_sentry_wards"`
			NpcDotaNeutralOgreMauler               int `json:"npc_dota_neutral_ogre_mauler"`
			NpcDotaNeutralOgreMagi                 int `json:"npc_dota_neutral_ogre_magi"`
			NpcDotaNeutralCentaurOutrunner         int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaNeutralCentaurKhan              int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaNeutralWarpineRaider            int `json:"npc_dota_neutral_warpine_raider"`
			NpcDotaNeutralKoboldTunneler           int `json:"npc_dota_neutral_kobold_tunneler"`
			NpcDotaNeutralKobold                   int `json:"npc_dota_neutral_kobold"`
			NpcDotaNeutralKoboldTaskmaster         int `json:"npc_dota_neutral_kobold_taskmaster"`
			NpcDotaGoodguysTower3Mid               int `json:"npc_dota_goodguys_tower3_mid"`
			NpcDotaGoodguysMeleeRaxMid             int `json:"npc_dota_goodguys_melee_rax_mid"`
			NpcDotaCreepBadguysFlagbearerUpgraded  int `json:"npc_dota_creep_badguys_flagbearer_upgraded"`
			NpcDotaNeutralPolarFurbolgUrsaWarrior  int `json:"npc_dota_neutral_polar_furbolg_ursa_warrior"`
			NpcDotaNeutralPolarFurbolgChampion     int `json:"npc_dota_neutral_polar_furbolg_champion"`
			NpcDotaCreepBadguysRangedUpgraded      int `json:"npc_dota_creep_badguys_ranged_upgraded"`
			NpcDotaCreepBadguysMeleeUpgraded       int `json:"npc_dota_creep_badguys_melee_upgraded"`
		} `json:"killed,omitempty"`
		KilledBy4 struct {
		} `json:"killed_by,omitempty"`
		LanePos5 struct {
			Num90 struct {
				Num170 int `json:"170"`
			} `json:"90"`
			Num92 struct {
				Num168 int `json:"168"`
				Num170 int `json:"170"`
			} `json:"92"`
			Num94 struct {
				Num168 int `json:"168"`
				Num172 int `json:"172"`
			} `json:"94"`
			Num96 struct {
				Num172 int `json:"172"`
			} `json:"96"`
			Num100 struct {
				Num162 int `json:"162"`
			} `json:"100"`
			Num102 struct {
				Num160 int `json:"160"`
			} `json:"102"`
			Num104 struct {
				Num160 int `json:"160"`
			} `json:"104"`
			Num106 struct {
				Num158 int `json:"158"`
			} `json:"106"`
			Num108 struct {
				Num154 int `json:"154"`
				Num158 int `json:"158"`
				Num176 int `json:"176"`
			} `json:"108"`
			Num110 struct {
				Num154 int `json:"154"`
				Num156 int `json:"156"`
				Num158 int `json:"158"`
				Num176 int `json:"176"`
			} `json:"110"`
			Num112 struct {
				Num156 int `json:"156"`
				Num158 int `json:"158"`
				Num176 int `json:"176"`
				Num178 int `json:"178"`
				Num180 int `json:"180"`
				Num182 int `json:"182"`
			} `json:"112"`
			Num114 struct {
				Num156 int `json:"156"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num176 int `json:"176"`
			} `json:"114"`
			Num116 struct {
				Num156 int `json:"156"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num174 int `json:"174"`
				Num176 int `json:"176"`
			} `json:"116"`
			Num118 struct {
				Num120 int `json:"120"`
				Num122 int `json:"122"`
				Num132 int `json:"132"`
				Num136 int `json:"136"`
				Num138 int `json:"138"`
				Num148 int `json:"148"`
				Num156 int `json:"156"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num172 int `json:"172"`
			} `json:"118"`
			Num120 struct {
				Num122 int `json:"122"`
				Num124 int `json:"124"`
				Num126 int `json:"126"`
				Num128 int `json:"128"`
				Num130 int `json:"130"`
				Num150 int `json:"150"`
				Num152 int `json:"152"`
				Num154 int `json:"154"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
			} `json:"120"`
			Num122 struct {
				Num122 int `json:"122"`
				Num124 int `json:"124"`
				Num126 int `json:"126"`
				Num128 int `json:"128"`
				Num154 int `json:"154"`
				Num156 int `json:"156"`
				Num158 int `json:"158"`
			} `json:"122"`
			Num124 struct {
				Num122 int `json:"122"`
				Num124 int `json:"124"`
				Num126 int `json:"126"`
				Num128 int `json:"128"`
				Num130 int `json:"130"`
				Num152 int `json:"152"`
				Num154 int `json:"154"`
				Num156 int `json:"156"`
				Num158 int `json:"158"`
			} `json:"124"`
			Num126 struct {
				Num122 int `json:"122"`
				Num124 int `json:"124"`
				Num126 int `json:"126"`
				Num128 int `json:"128"`
				Num130 int `json:"130"`
				Num132 int `json:"132"`
				Num134 int `json:"134"`
				Num136 int `json:"136"`
				Num144 int `json:"144"`
				Num146 int `json:"146"`
				Num148 int `json:"148"`
				Num150 int `json:"150"`
				Num160 int `json:"160"`
			} `json:"126"`
			Num128 struct {
				Num120 int `json:"120"`
				Num126 int `json:"126"`
				Num128 int `json:"128"`
				Num130 int `json:"130"`
				Num134 int `json:"134"`
				Num138 int `json:"138"`
				Num140 int `json:"140"`
				Num142 int `json:"142"`
				Num144 int `json:"144"`
				Num146 int `json:"146"`
				Num150 int `json:"150"`
				Num162 int `json:"162"`
			} `json:"128"`
			Num130 struct {
				Num124 int `json:"124"`
				Num126 int `json:"126"`
				Num128 int `json:"128"`
				Num130 int `json:"130"`
				Num132 int `json:"132"`
				Num134 int `json:"134"`
				Num138 int `json:"138"`
				Num140 int `json:"140"`
				Num150 int `json:"150"`
				Num162 int `json:"162"`
			} `json:"130"`
			Num132 struct {
				Num120 int `json:"120"`
				Num124 int `json:"124"`
				Num126 int `json:"126"`
				Num128 int `json:"128"`
				Num130 int `json:"130"`
				Num132 int `json:"132"`
				Num134 int `json:"134"`
				Num136 int `json:"136"`
				Num148 int `json:"148"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
			} `json:"132"`
			Num134 struct {
				Num78  int `json:"78"`
				Num80  int `json:"80"`
				Num120 int `json:"120"`
				Num124 int `json:"124"`
				Num126 int `json:"126"`
				Num130 int `json:"130"`
				Num132 int `json:"132"`
				Num134 int `json:"134"`
				Num148 int `json:"148"`
				Num168 int `json:"168"`
			} `json:"134"`
			Num136 struct {
				Num80  int `json:"80"`
				Num118 int `json:"118"`
				Num120 int `json:"120"`
				Num122 int `json:"122"`
				Num128 int `json:"128"`
				Num130 int `json:"130"`
				Num132 int `json:"132"`
				Num134 int `json:"134"`
				Num144 int `json:"144"`
				Num146 int `json:"146"`
				Num168 int `json:"168"`
			} `json:"136"`
			Num138 struct {
				Num82  int `json:"82"`
				Num116 int `json:"116"`
				Num130 int `json:"130"`
				Num132 int `json:"132"`
				Num134 int `json:"134"`
				Num136 int `json:"136"`
				Num138 int `json:"138"`
				Num140 int `json:"140"`
				Num142 int `json:"142"`
				Num170 int `json:"170"`
			} `json:"138"`
			Num140 struct {
				Num82  int `json:"82"`
				Num96  int `json:"96"`
				Num116 int `json:"116"`
				Num136 int `json:"136"`
				Num170 int `json:"170"`
			} `json:"140"`
			Num142 struct {
				Num84  int `json:"84"`
				Num86  int `json:"86"`
				Num90  int `json:"90"`
				Num92  int `json:"92"`
				Num94  int `json:"94"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num102 int `json:"102"`
				Num108 int `json:"108"`
				Num110 int `json:"110"`
				Num112 int `json:"112"`
				Num136 int `json:"136"`
				Num138 int `json:"138"`
				Num172 int `json:"172"`
			} `json:"142"`
			Num144 struct {
				Num98  int `json:"98"`
				Num102 int `json:"102"`
				Num106 int `json:"106"`
				Num140 int `json:"140"`
				Num172 int `json:"172"`
			} `json:"144"`
			Num146 struct {
				Num98  int `json:"98"`
				Num102 int `json:"102"`
				Num140 int `json:"140"`
				Num142 int `json:"142"`
			} `json:"146"`
			Num148 struct {
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num142 int `json:"142"`
				Num144 int `json:"144"`
				Num172 int `json:"172"`
			} `json:"148"`
			Num150 struct {
				Num98  int `json:"98"`
				Num144 int `json:"144"`
				Num146 int `json:"146"`
				Num174 int `json:"174"`
			} `json:"150"`
			Num152 struct {
				Num98  int `json:"98"`
				Num146 int `json:"146"`
				Num148 int `json:"148"`
				Num174 int `json:"174"`
			} `json:"152"`
			Num154 struct {
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num148 int `json:"148"`
				Num150 int `json:"150"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
			} `json:"154"`
			Num156 struct {
				Num94  int `json:"94"`
				Num96  int `json:"96"`
				Num152 int `json:"152"`
				Num172 int `json:"172"`
			} `json:"156"`
			Num158 struct {
				Num92  int `json:"92"`
				Num94  int `json:"94"`
				Num154 int `json:"154"`
				Num170 int `json:"170"`
			} `json:"158"`
			Num160 struct {
				Num90 int `json:"90"`
				Num92 int `json:"92"`
			} `json:"160"`
			Num162 struct {
				Num88  int `json:"88"`
				Num90  int `json:"90"`
				Num170 int `json:"170"`
			} `json:"162"`
			Num164 struct {
				Num86  int `json:"86"`
				Num170 int `json:"170"`
			} `json:"164"`
			Num166 struct {
				Num84  int `json:"84"`
				Num86  int `json:"86"`
				Num88  int `json:"88"`
				Num170 int `json:"170"`
			} `json:"166"`
			Num168 struct {
				Num84  int `json:"84"`
				Num86  int `json:"86"`
				Num170 int `json:"170"`
			} `json:"168"`
			Num170 struct {
				Num84  int `json:"84"`
				Num86  int `json:"86"`
				Num170 int `json:"170"`
			} `json:"170"`
			Num172 struct {
				Num172 int `json:"172"`
			} `json:"172"`
			Num174 struct {
				Num172 int `json:"172"`
			} `json:"174"`
			Num176 struct {
				Num174 int `json:"174"`
			} `json:"176"`
			Num178 struct {
				Num176 int `json:"176"`
			} `json:"178"`
			Num180 struct {
				Num176 int `json:"176"`
			} `json:"180"`
			Num182 struct {
				Num178 int `json:"178"`
			} `json:"182"`
		} `json:"lane_pos,omitempty"`
		LifeState0 struct {
			Num0 int `json:"0"`
		} `json:"life_state,omitempty"`
		MultiKills1 struct {
			Num2 int `json:"2"`
		} `json:"multi_kills,omitempty"`
		Obs5 struct {
			Num94 struct {
				Num140 int `json:"140"`
			} `json:"94"`
			Num98 struct {
				Num168 int `json:"168"`
			} `json:"98"`
		} `json:"obs,omitempty"`
		Purchase5 struct {
			Branches           int `json:"branches"`
			Tango              int `json:"tango"`
			FaerieFire         int `json:"faerie_fire"`
			Bottle             int `json:"bottle"`
			MagicStick         int `json:"magic_stick"`
			Gloves             int `json:"gloves"`
			RecipeMagicWand    int `json:"recipe_magic_wand"`
			InfusedRaindrop    int `json:"infused_raindrop"`
			Clarity            int `json:"clarity"`
			MagicWand          int `json:"magic_wand"`
			Robe               int `json:"robe"`
			Boots              int `json:"boots"`
			PowerTreads        int `json:"power_treads"`
			Tpscroll           int `json:"tpscroll"`
			VoidStone          int `json:"void_stone"`
			WardObserver       int `json:"ward_observer"`
			BlitzKnuckles      int `json:"blitz_knuckles"`
			Claymore           int `json:"claymore"`
			RecipeOrchid       int `json:"recipe_orchid"`
			Orchid             int `json:"orchid"`
			OgreAxe            int `json:"ogre_axe"`
			MithrilHammer      int `json:"mithril_hammer"`
			RecipeBlackKingBar int `json:"recipe_black_king_bar"`
			BlackKingBar       int `json:"black_king_bar"`
			Platemail          int `json:"platemail"`
			AghanimsShard      int `json:"aghanims_shard"`
			MysticStaff        int `json:"mystic_staff"`
			RecipeShivasGuard  int `json:"recipe_shivas_guard"`
			ShivasGuard        int `json:"shivas_guard"`
			UltimateOrb        int `json:"ultimate_orb"`
			Sheepstick         int `json:"sheepstick"`
			Cloak              int `json:"cloak"`
			Quarterstaff       int `json:"quarterstaff"`
			SobiMask           int `json:"sobi_mask"`
			OblivionStaff      int `json:"oblivion_staff"`
			MageSlayer         int `json:"mage_slayer"`
			RecipeMageSlayer   int `json:"recipe_mage_slayer"`
			Bloodthorn         int `json:"bloodthorn"`
			RecipeBloodthorn   int `json:"recipe_bloodthorn"`
		} `json:"purchase,omitempty"`
		Runes2 struct {
			Num0 int `json:"0"`
			Num1 int `json:"1"`
			Num2 int `json:"2"`
			Num3 int `json:"3"`
			Num4 int `json:"4"`
			Num5 int `json:"5"`
			Num6 int `json:"6"`
			Num7 int `json:"7"`
		} `json:"runes,omitempty"`
		Sen5 struct {
		} `json:"sen,omitempty"`
		XpReasons5 struct {
			Num0 int `json:"0"`
			Num1 int `json:"1"`
			Num2 int `json:"2"`
			Num3 int `json:"3"`
			Num5 int `json:"5"`
		} `json:"xp_reasons,omitempty"`
		PurchaseTime5 struct {
			Branches        int `json:"branches"`
			Tango           int `json:"tango"`
			FaerieFire      int `json:"faerie_fire"`
			Bottle          int `json:"bottle"`
			MagicStick      int `json:"magic_stick"`
			Gloves          int `json:"gloves"`
			InfusedRaindrop int `json:"infused_raindrop"`
			Clarity         int `json:"clarity"`
			MagicWand       int `json:"magic_wand"`
			Robe            int `json:"robe"`
			Boots           int `json:"boots"`
			PowerTreads     int `json:"power_treads"`
			Tpscroll        int `json:"tpscroll"`
			VoidStone       int `json:"void_stone"`
			WardObserver    int `json:"ward_observer"`
			BlitzKnuckles   int `json:"blitz_knuckles"`
			Claymore        int `json:"claymore"`
			Orchid          int `json:"orchid"`
			OgreAxe         int `json:"ogre_axe"`
			MithrilHammer   int `json:"mithril_hammer"`
			BlackKingBar    int `json:"black_king_bar"`
			Platemail       int `json:"platemail"`
			AghanimsShard   int `json:"aghanims_shard"`
			MysticStaff     int `json:"mystic_staff"`
			ShivasGuard     int `json:"shivas_guard"`
			UltimateOrb     int `json:"ultimate_orb"`
			Sheepstick      int `json:"sheepstick"`
			Cloak           int `json:"cloak"`
			Quarterstaff    int `json:"quarterstaff"`
			SobiMask        int `json:"sobi_mask"`
			OblivionStaff   int `json:"oblivion_staff"`
			MageSlayer      int `json:"mage_slayer"`
			Bloodthorn      int `json:"bloodthorn"`
		} `json:"purchase_time,omitempty"`
		FirstPurchaseTime5 struct {
			Branches        int `json:"branches"`
			Tango           int `json:"tango"`
			FaerieFire      int `json:"faerie_fire"`
			Bottle          int `json:"bottle"`
			MagicStick      int `json:"magic_stick"`
			Gloves          int `json:"gloves"`
			InfusedRaindrop int `json:"infused_raindrop"`
			Clarity         int `json:"clarity"`
			MagicWand       int `json:"magic_wand"`
			Robe            int `json:"robe"`
			Boots           int `json:"boots"`
			PowerTreads     int `json:"power_treads"`
			Tpscroll        int `json:"tpscroll"`
			VoidStone       int `json:"void_stone"`
			WardObserver    int `json:"ward_observer"`
			BlitzKnuckles   int `json:"blitz_knuckles"`
			Claymore        int `json:"claymore"`
			Orchid          int `json:"orchid"`
			OgreAxe         int `json:"ogre_axe"`
			MithrilHammer   int `json:"mithril_hammer"`
			BlackKingBar    int `json:"black_king_bar"`
			Platemail       int `json:"platemail"`
			AghanimsShard   int `json:"aghanims_shard"`
			MysticStaff     int `json:"mystic_staff"`
			ShivasGuard     int `json:"shivas_guard"`
			UltimateOrb     int `json:"ultimate_orb"`
			Sheepstick      int `json:"sheepstick"`
			Cloak           int `json:"cloak"`
			Quarterstaff    int `json:"quarterstaff"`
			SobiMask        int `json:"sobi_mask"`
			OblivionStaff   int `json:"oblivion_staff"`
			MageSlayer      int `json:"mage_slayer"`
			Bloodthorn      int `json:"bloodthorn"`
		} `json:"first_purchase_time,omitempty"`
		ItemWin5 struct {
			Branches        int `json:"branches"`
			Tango           int `json:"tango"`
			FaerieFire      int `json:"faerie_fire"`
			Bottle          int `json:"bottle"`
			MagicStick      int `json:"magic_stick"`
			Gloves          int `json:"gloves"`
			InfusedRaindrop int `json:"infused_raindrop"`
			Clarity         int `json:"clarity"`
			MagicWand       int `json:"magic_wand"`
			Robe            int `json:"robe"`
			Boots           int `json:"boots"`
			PowerTreads     int `json:"power_treads"`
			Tpscroll        int `json:"tpscroll"`
			VoidStone       int `json:"void_stone"`
			WardObserver    int `json:"ward_observer"`
			BlitzKnuckles   int `json:"blitz_knuckles"`
			Claymore        int `json:"claymore"`
			Orchid          int `json:"orchid"`
			OgreAxe         int `json:"ogre_axe"`
			MithrilHammer   int `json:"mithril_hammer"`
			BlackKingBar    int `json:"black_king_bar"`
			Platemail       int `json:"platemail"`
			AghanimsShard   int `json:"aghanims_shard"`
			MysticStaff     int `json:"mystic_staff"`
			ShivasGuard     int `json:"shivas_guard"`
			UltimateOrb     int `json:"ultimate_orb"`
			Sheepstick      int `json:"sheepstick"`
			Cloak           int `json:"cloak"`
			Quarterstaff    int `json:"quarterstaff"`
			SobiMask        int `json:"sobi_mask"`
			OblivionStaff   int `json:"oblivion_staff"`
			MageSlayer      int `json:"mage_slayer"`
			Bloodthorn      int `json:"bloodthorn"`
		} `json:"item_win,omitempty"`
		ItemUsage5 struct {
			Branches        int `json:"branches"`
			Tango           int `json:"tango"`
			FaerieFire      int `json:"faerie_fire"`
			Bottle          int `json:"bottle"`
			MagicStick      int `json:"magic_stick"`
			Gloves          int `json:"gloves"`
			InfusedRaindrop int `json:"infused_raindrop"`
			Clarity         int `json:"clarity"`
			MagicWand       int `json:"magic_wand"`
			Robe            int `json:"robe"`
			Boots           int `json:"boots"`
			PowerTreads     int `json:"power_treads"`
			Tpscroll        int `json:"tpscroll"`
			VoidStone       int `json:"void_stone"`
			WardObserver    int `json:"ward_observer"`
			BlitzKnuckles   int `json:"blitz_knuckles"`
			Claymore        int `json:"claymore"`
			Orchid          int `json:"orchid"`
			OgreAxe         int `json:"ogre_axe"`
			MithrilHammer   int `json:"mithril_hammer"`
			BlackKingBar    int `json:"black_king_bar"`
			Platemail       int `json:"platemail"`
			AghanimsShard   int `json:"aghanims_shard"`
			MysticStaff     int `json:"mystic_staff"`
			ShivasGuard     int `json:"shivas_guard"`
			UltimateOrb     int `json:"ultimate_orb"`
			Sheepstick      int `json:"sheepstick"`
			Cloak           int `json:"cloak"`
			Quarterstaff    int `json:"quarterstaff"`
			SobiMask        int `json:"sobi_mask"`
			OblivionStaff   int `json:"oblivion_staff"`
			MageSlayer      int `json:"mage_slayer"`
			Bloodthorn      int `json:"bloodthorn"`
		} `json:"item_usage,omitempty"`
		AbilityTargets6 struct {
			DragonKnightDragonTail struct {
				NpcDotaHeroTemplarAssassin int `json:"npc_dota_hero_templar_assassin"`
				NpcDotaHeroEnchantress     int `json:"npc_dota_hero_enchantress"`
				NpcDotaHeroDeathProphet    int `json:"npc_dota_hero_death_prophet"`
				NpcDotaHeroPuck            int `json:"npc_dota_hero_puck"`
				NpcDotaHeroPrimalBeast     int `json:"npc_dota_hero_primal_beast"`
			} `json:"dragon_knight_dragon_tail"`
			DragonKnightBreatheFire struct {
				NpcDotaHeroDeathProphet int `json:"npc_dota_hero_death_prophet"`
				NpcDotaHeroPuck         int `json:"npc_dota_hero_puck"`
				NpcDotaHeroEnchantress  int `json:"npc_dota_hero_enchantress"`
				NpcDotaHeroPrimalBeast  int `json:"npc_dota_hero_primal_beast"`
			} `json:"dragon_knight_breathe_fire"`
		} `json:"ability_targets,omitempty"`
		AbilityUses6 struct {
			DragonKnightDragonTail      int `json:"dragon_knight_dragon_tail"`
			PlusHighFive                int `json:"plus_high_five"`
			DragonKnightElderDragonForm int `json:"dragon_knight_elder_dragon_form"`
			DragonKnightBreatheFire     int `json:"dragon_knight_breathe_fire"`
			DragonKnightFireball        int `json:"dragon_knight_fireball"`
			AbilityCapture              int `json:"ability_capture"`
		} `json:"ability_uses,omitempty"`
		Actions6 struct {
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
			Num12 int `json:"12"`
			Num13 int `json:"13"`
			Num14 int `json:"14"`
			Num15 int `json:"15"`
			Num16 int `json:"16"`
			Num17 int `json:"17"`
			Num19 int `json:"19"`
			Num21 int `json:"21"`
			Num24 int `json:"24"`
			Num27 int `json:"27"`
			Num33 int `json:"33"`
			Num37 int `json:"37"`
			Num38 int `json:"38"`
		} `json:"actions,omitempty"`
		Damage6 struct {
			NpcDotaCreepGoodguysMelee              int `json:"npc_dota_creep_goodguys_melee"`
			NpcDotaHeroTemplarAssassin             int `json:"npc_dota_hero_templar_assassin"`
			NpcDotaCreepBadguysRanged              int `json:"npc_dota_creep_badguys_ranged"`
			NpcDotaCreepGoodguysRanged             int `json:"npc_dota_creep_goodguys_ranged"`
			NpcDotaCreepBadguysMelee               int `json:"npc_dota_creep_badguys_melee"`
			NpcDotaCreepGoodguysFlagbearer         int `json:"npc_dota_creep_goodguys_flagbearer"`
			NpcDotaCreepBadguysFlagbearer          int `json:"npc_dota_creep_badguys_flagbearer"`
			NpcDotaGoodguysTower1Bot               int `json:"npc_dota_goodguys_tower1_bot"`
			NpcDotaTemplarAssassinPsionicTrap      int `json:"npc_dota_templar_assassin_psionic_trap"`
			NpcDotaNeutralCentaurKhan              int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaHeroDeathProphet                int `json:"npc_dota_hero_death_prophet"`
			NpcDotaHeroEnchantress                 int `json:"npc_dota_hero_enchantress"`
			NpcDotaNeutralDarkTrollWarlord         int `json:"npc_dota_neutral_dark_troll_warlord"`
			NpcDotaNeutralPolarFurbolgChampion     int `json:"npc_dota_neutral_polar_furbolg_champion"`
			NpcDotaNeutralWarpineRaider            int `json:"npc_dota_neutral_warpine_raider"`
			NpcDotaNeutralDarkTroll                int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaDarkTrollWarlordSkeletonWarrior int `json:"npc_dota_dark_troll_warlord_skeleton_warrior"`
			NpcDotaGoodguysSiege                   int `json:"npc_dota_goodguys_siege"`
			NpcDotaNeutralEnragedWildkin           int `json:"npc_dota_neutral_enraged_wildkin"`
			NpcDotaHeroPrimalBeast                 int `json:"npc_dota_hero_primal_beast"`
			NpcDotaHeroPuck                        int `json:"npc_dota_hero_puck"`
			NpcDotaNeutralSatyrHellcaller          int `json:"npc_dota_neutral_satyr_hellcaller"`
			NpcDotaNeutralForestTrollBerserker     int `json:"npc_dota_neutral_forest_troll_berserker"`
			NpcDotaNeutralKoboldTaskmaster         int `json:"npc_dota_neutral_kobold_taskmaster"`
			NpcDotaNeutralPolarFurbolgUrsaWarrior  int `json:"npc_dota_neutral_polar_furbolg_ursa_warrior"`
			NpcDotaNeutralWildkin                  int `json:"npc_dota_neutral_wildkin"`
			NpcDotaGoodguysTower2Bot               int `json:"npc_dota_goodguys_tower2_bot"`
			NpcDotaNeutralSatyrSoulstealer         int `json:"npc_dota_neutral_satyr_soulstealer"`
			NpcDotaNeutralCentaurOutrunner         int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaNeutralMudGolem                 int `json:"npc_dota_neutral_mud_golem"`
			NpcDotaNeutralMudGolemSplit            int `json:"npc_dota_neutral_mud_golem_split"`
			NpcDotaNeutralRockGolem                int `json:"npc_dota_neutral_rock_golem"`
			NpcDotaNeutralGraniteGolem             int `json:"npc_dota_neutral_granite_golem"`
			NpcDotaGoodguysTower2Top               int `json:"npc_dota_goodguys_tower2_top"`
			NpcDotaRoshan                          int `json:"npc_dota_roshan"`
			NpcDotaNeutralBlackDrake               int `json:"npc_dota_neutral_black_drake"`
			NpcDotaNeutralBlackDragon              int `json:"npc_dota_neutral_black_dragon"`
			NpcDotaGoodguysTower3Bot               int `json:"npc_dota_goodguys_tower3_bot"`
			NpcDotaGoodguysMeleeRaxBot             int `json:"npc_dota_goodguys_melee_rax_bot"`
			NpcDotaGoodguysRangeRaxBot             int `json:"npc_dota_goodguys_range_rax_bot"`
			NpcDotaNeutralOgreMagi                 int `json:"npc_dota_neutral_ogre_magi"`
			NpcDotaGoodguysTower3Top               int `json:"npc_dota_goodguys_tower3_top"`
			NpcDotaGoodguysRangeRaxTop             int `json:"npc_dota_goodguys_range_rax_top"`
			NpcDotaGoodguysMeleeRaxTop             int `json:"npc_dota_goodguys_melee_rax_top"`
			NpcDotaGoodguysTower4                  int `json:"npc_dota_goodguys_tower4"`
		} `json:"damage,omitempty"`
		DamageInflictor6 struct {
			DragonKnightDragonTail      int `json:"dragon_knight_dragon_tail"`
			Null                        int `json:"null"`
			DragonKnightBreatheFire     int `json:"dragon_knight_breathe_fire"`
			DragonKnightElderDragonForm int `json:"dragon_knight_elder_dragon_form"`
			DragonKnightFireball        int `json:"dragon_knight_fireball"`
		} `json:"damage_inflictor,omitempty"`
		DamageInflictorReceived6 struct {
			Null                     int `json:"null"`
			TemplarAssassinMeld      int `json:"templar_assassin_meld"`
			TemplarAssassinPsiBlades int `json:"templar_assassin_psi_blades"`
			PuckDreamCoil            int `json:"puck_dream_coil"`
			PuckWaningRift           int `json:"puck_waning_rift"`
			CentaurKhanWarStomp      int `json:"centaur_khan_war_stomp"`
			EnchantressImpetus       int `json:"enchantress_impetus"`
			PuckIllusoryOrb          int `json:"puck_illusory_orb"`
			DeathProphetSpiritSiphon int `json:"death_prophet_spirit_siphon"`
			PrimalBeastOnslaught     int `json:"primal_beast_onslaught"`
			DeathProphetExorcism     int `json:"death_prophet_exorcism"`
			DeathProphetCarrionSwarm int `json:"death_prophet_carrion_swarm"`
			PrimalBeastTrample       int `json:"primal_beast_trample"`
			WarpineRaiderSeedShot    int `json:"warpine_raider_seed_shot"`
			WitchBlade               int `json:"witch_blade"`
			PrimalBeastPulverize     int `json:"primal_beast_pulverize"`
			SpiritVessel             int `json:"spirit_vessel"`
			DragonScale              int `json:"dragon_scale"`
			PrimalBeastRockThrow     int `json:"primal_beast_rock_throw"`
			Stormcrafter             int `json:"stormcrafter"`
			OverwhelmingBlink        int `json:"overwhelming_blink"`
		} `json:"damage_inflictor_received,omitempty"`
		DamageTaken6 struct {
			NpcDotaCreepGoodguysMelee              int `json:"npc_dota_creep_goodguys_melee"`
			NpcDotaHeroEnchantress                 int `json:"npc_dota_hero_enchantress"`
			NpcDotaHeroTemplarAssassin             int `json:"npc_dota_hero_templar_assassin"`
			NpcDotaCreepGoodguysFlagbearer         int `json:"npc_dota_creep_goodguys_flagbearer"`
			NpcDotaCreepGoodguysRanged             int `json:"npc_dota_creep_goodguys_ranged"`
			NpcDotaGoodguysSiege                   int `json:"npc_dota_goodguys_siege"`
			NpcDotaHeroPuck                        int `json:"npc_dota_hero_puck"`
			NpcDotaGoodguysTower1Bot               int `json:"npc_dota_goodguys_tower1_bot"`
			NpcDotaHeroDeathProphet                int `json:"npc_dota_hero_death_prophet"`
			NpcDotaHeroPrimalBeast                 int `json:"npc_dota_hero_primal_beast"`
			NpcDotaNeutralDarkTrollWarlord         int `json:"npc_dota_neutral_dark_troll_warlord"`
			NpcDotaNeutralDarkTroll                int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaNeutralPolarFurbolgChampion     int `json:"npc_dota_neutral_polar_furbolg_champion"`
			NpcDotaNeutralWarpineRaider            int `json:"npc_dota_neutral_warpine_raider"`
			NpcDotaNeutralWildkin                  int `json:"npc_dota_neutral_wildkin"`
			NpcDotaNeutralEnragedWildkin           int `json:"npc_dota_neutral_enraged_wildkin"`
			NpcDotaDarkTrollWarlordSkeletonWarrior int `json:"npc_dota_dark_troll_warlord_skeleton_warrior"`
			NpcDotaNeutralKoboldTaskmaster         int `json:"npc_dota_neutral_kobold_taskmaster"`
			NpcDotaNeutralForestTrollBerserker     int `json:"npc_dota_neutral_forest_troll_berserker"`
			NpcDotaGoodguysTower2Bot               int `json:"npc_dota_goodguys_tower2_bot"`
			NpcDotaNeutralCentaurKhan              int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaNeutralCentaurOutrunner         int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaNeutralMudGolem                 int `json:"npc_dota_neutral_mud_golem"`
			NpcDotaNeutralMudGolemSplit            int `json:"npc_dota_neutral_mud_golem_split"`
			NpcDotaNeutralRockGolem                int `json:"npc_dota_neutral_rock_golem"`
			NpcDotaNeutralGraniteGolem             int `json:"npc_dota_neutral_granite_golem"`
			NpcDotaGoodguysTower2Top               int `json:"npc_dota_goodguys_tower2_top"`
			NpcDotaGoodguysTower3Top               int `json:"npc_dota_goodguys_tower3_top"`
			NpcDotaNeutralBlackDragon              int `json:"npc_dota_neutral_black_dragon"`
			NpcDotaNeutralBlackDrake               int `json:"npc_dota_neutral_black_drake"`
			NpcDotaGoodguysTower3Bot               int `json:"npc_dota_goodguys_tower3_bot"`
			NpcDotaGoodguysTower4                  int `json:"npc_dota_goodguys_tower4"`
		} `json:"damage_taken,omitempty"`
		DamageTargets6 struct {
			DragonKnightDragonTail struct {
				NpcDotaHeroTemplarAssassin int `json:"npc_dota_hero_templar_assassin"`
				NpcDotaHeroEnchantress     int `json:"npc_dota_hero_enchantress"`
				NpcDotaHeroDeathProphet    int `json:"npc_dota_hero_death_prophet"`
				NpcDotaHeroPuck            int `json:"npc_dota_hero_puck"`
				NpcDotaHeroPrimalBeast     int `json:"npc_dota_hero_primal_beast"`
			} `json:"dragon_knight_dragon_tail"`
			Null struct {
				NpcDotaHeroTemplarAssassin int `json:"npc_dota_hero_templar_assassin"`
				NpcDotaHeroDeathProphet    int `json:"npc_dota_hero_death_prophet"`
				NpcDotaHeroPrimalBeast     int `json:"npc_dota_hero_primal_beast"`
				NpcDotaHeroPuck            int `json:"npc_dota_hero_puck"`
			} `json:"null"`
			DragonKnightBreatheFire struct {
				NpcDotaHeroDeathProphet    int `json:"npc_dota_hero_death_prophet"`
				NpcDotaHeroEnchantress     int `json:"npc_dota_hero_enchantress"`
				NpcDotaHeroPuck            int `json:"npc_dota_hero_puck"`
				NpcDotaHeroPrimalBeast     int `json:"npc_dota_hero_primal_beast"`
				NpcDotaHeroTemplarAssassin int `json:"npc_dota_hero_templar_assassin"`
			} `json:"dragon_knight_breathe_fire"`
			DragonKnightElderDragonForm struct {
				NpcDotaHeroPuck            int `json:"npc_dota_hero_puck"`
				NpcDotaHeroPrimalBeast     int `json:"npc_dota_hero_primal_beast"`
				NpcDotaHeroEnchantress     int `json:"npc_dota_hero_enchantress"`
				NpcDotaHeroTemplarAssassin int `json:"npc_dota_hero_templar_assassin"`
				NpcDotaHeroDeathProphet    int `json:"npc_dota_hero_death_prophet"`
			} `json:"dragon_knight_elder_dragon_form"`
			DragonKnightFireball struct {
				NpcDotaHeroPuck            int `json:"npc_dota_hero_puck"`
				NpcDotaHeroPrimalBeast     int `json:"npc_dota_hero_primal_beast"`
				NpcDotaHeroDeathProphet    int `json:"npc_dota_hero_death_prophet"`
				NpcDotaHeroEnchantress     int `json:"npc_dota_hero_enchantress"`
				NpcDotaHeroTemplarAssassin int `json:"npc_dota_hero_templar_assassin"`
			} `json:"dragon_knight_fireball"`
		} `json:"damage_targets,omitempty"`
		GoldReasons5 struct {
			Num0  int `json:"0"`
			Num1  int `json:"1"`
			Num6  int `json:"6"`
			Num11 int `json:"11"`
			Num12 int `json:"12"`
			Num13 int `json:"13"`
			Num14 int `json:"14"`
			Num15 int `json:"15"`
			Num16 int `json:"16"`
			Num17 int `json:"17"`
			Num19 int `json:"19"`
		} `json:"gold_reasons,omitempty"`
		HeroHits6 struct {
			DragonKnightDragonTail      int `json:"dragon_knight_dragon_tail"`
			Null                        int `json:"null"`
			DragonKnightBreatheFire     int `json:"dragon_knight_breathe_fire"`
			DragonKnightElderDragonForm int `json:"dragon_knight_elder_dragon_form"`
			DragonKnightFireball        int `json:"dragon_knight_fireball"`
		} `json:"hero_hits,omitempty"`
		ItemUses6 struct {
			QuellingBlade         int `json:"quelling_blade"`
			Tango                 int `json:"tango"`
			Tpscroll              int `json:"tpscroll"`
			PowerTreads           int `json:"power_treads"`
			WardSentry            int `json:"ward_sentry"`
			Blink                 int `json:"blink"`
			SmokeOfDeceit         int `json:"smoke_of_deceit"`
			SoulRing              int `json:"soul_ring"`
			OgreSealTotem         int `json:"ogre_seal_totem"`
			BlackKingBar          int `json:"black_king_bar"`
			HeavensHalberd        int `json:"heavens_halberd"`
			UltimateScepterRoshan int `json:"ultimate_scepter_roshan"`
			TricksterCloak        int `json:"trickster_cloak"`
			Cheese                int `json:"cheese"`
		} `json:"item_uses,omitempty"`
		KillStreaks3 struct {
			Num3 int `json:"3"`
		} `json:"kill_streaks,omitempty"`
		Killed6 struct {
			NpcDotaCreepBadguysRanged              int `json:"npc_dota_creep_badguys_ranged"`
			NpcDotaCreepGoodguysMelee              int `json:"npc_dota_creep_goodguys_melee"`
			NpcDotaCreepGoodguysRanged             int `json:"npc_dota_creep_goodguys_ranged"`
			NpcDotaCreepGoodguysFlagbearer         int `json:"npc_dota_creep_goodguys_flagbearer"`
			NpcDotaCreepBadguysMelee               int `json:"npc_dota_creep_badguys_melee"`
			NpcDotaTemplarAssassinPsionicTrap      int `json:"npc_dota_templar_assassin_psionic_trap"`
			NpcDotaGoodguysTower1Bot               int `json:"npc_dota_goodguys_tower1_bot"`
			NpcDotaNeutralCentaurKhan              int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaNeutralDarkTroll                int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaGoodguysSiege                   int `json:"npc_dota_goodguys_siege"`
			NpcDotaDarkTrollWarlordSkeletonWarrior int `json:"npc_dota_dark_troll_warlord_skeleton_warrior"`
			NpcDotaNeutralDarkTrollWarlord         int `json:"npc_dota_neutral_dark_troll_warlord"`
			NpcDotaHeroPuck                        int `json:"npc_dota_hero_puck"`
			NpcDotaNeutralForestTrollBerserker     int `json:"npc_dota_neutral_forest_troll_berserker"`
			NpcDotaNeutralKoboldTaskmaster         int `json:"npc_dota_neutral_kobold_taskmaster"`
			NpcDotaNeutralPolarFurbolgUrsaWarrior  int `json:"npc_dota_neutral_polar_furbolg_ursa_warrior"`
			NpcDotaNeutralWarpineRaider            int `json:"npc_dota_neutral_warpine_raider"`
			NpcDotaNeutralWildkin                  int `json:"npc_dota_neutral_wildkin"`
			NpcDotaNeutralSatyrHellcaller          int `json:"npc_dota_neutral_satyr_hellcaller"`
			NpcDotaNeutralCentaurOutrunner         int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaHeroPrimalBeast                 int `json:"npc_dota_hero_primal_beast"`
			NpcDotaNeutralMudGolem                 int `json:"npc_dota_neutral_mud_golem"`
			NpcDotaNeutralMudGolemSplit            int `json:"npc_dota_neutral_mud_golem_split"`
			NpcDotaNeutralRockGolem                int `json:"npc_dota_neutral_rock_golem"`
			NpcDotaNeutralGraniteGolem             int `json:"npc_dota_neutral_granite_golem"`
			NpcDotaNeutralEnragedWildkin           int `json:"npc_dota_neutral_enraged_wildkin"`
			NpcDotaNeutralBlackDragon              int `json:"npc_dota_neutral_black_dragon"`
			NpcDotaNeutralBlackDrake               int `json:"npc_dota_neutral_black_drake"`
			NpcDotaHeroEnchantress                 int `json:"npc_dota_hero_enchantress"`
			NpcDotaGoodguysTower3Bot               int `json:"npc_dota_goodguys_tower3_bot"`
			NpcDotaGoodguysMeleeRaxBot             int `json:"npc_dota_goodguys_melee_rax_bot"`
			NpcDotaGoodguysTower3Top               int `json:"npc_dota_goodguys_tower3_top"`
			NpcDotaHeroTemplarAssassin             int `json:"npc_dota_hero_templar_assassin"`
		} `json:"killed,omitempty"`
		KilledBy5 struct {
			NpcDotaHeroTemplarAssassin int `json:"npc_dota_hero_templar_assassin"`
			NpcDotaHeroDeathProphet    int `json:"npc_dota_hero_death_prophet"`
			NpcDotaHeroPuck            int `json:"npc_dota_hero_puck"`
			NpcDotaHeroPrimalBeast     int `json:"npc_dota_hero_primal_beast"`
		} `json:"killed_by,omitempty"`
		LanePos6 struct {
			Num106 struct {
				Num156 int `json:"156"`
				Num158 int `json:"158"`
			} `json:"106"`
			Num108 struct {
				Num156 int `json:"156"`
				Num158 int `json:"158"`
			} `json:"108"`
			Num110 struct {
				Num156 int `json:"156"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
			} `json:"110"`
			Num112 struct {
				Num156 int `json:"156"`
				Num158 int `json:"158"`
			} `json:"112"`
			Num114 struct {
				Num156 int `json:"156"`
				Num158 int `json:"158"`
			} `json:"114"`
			Num116 struct {
				Num160 int `json:"160"`
			} `json:"116"`
			Num118 struct {
				Num156 int `json:"156"`
				Num160 int `json:"160"`
			} `json:"118"`
			Num120 struct {
				Num154 int `json:"154"`
				Num162 int `json:"162"`
			} `json:"120"`
			Num122 struct {
				Num154 int `json:"154"`
				Num162 int `json:"162"`
			} `json:"122"`
			Num124 struct {
				Num152 int `json:"152"`
				Num162 int `json:"162"`
			} `json:"124"`
			Num126 struct {
				Num148 int `json:"148"`
				Num150 int `json:"150"`
			} `json:"126"`
			Num128 struct {
				Num144 int `json:"144"`
				Num146 int `json:"146"`
				Num164 int `json:"164"`
			} `json:"128"`
			Num130 struct {
				Num140 int `json:"140"`
				Num142 int `json:"142"`
				Num164 int `json:"164"`
			} `json:"130"`
			Num132 struct {
				Num130 int `json:"130"`
				Num132 int `json:"132"`
				Num138 int `json:"138"`
				Num166 int `json:"166"`
			} `json:"132"`
			Num134 struct {
				Num132 int `json:"132"`
				Num138 int `json:"138"`
				Num168 int `json:"168"`
			} `json:"134"`
			Num136 struct {
				Num122 int `json:"122"`
				Num124 int `json:"124"`
				Num132 int `json:"132"`
				Num134 int `json:"134"`
				Num136 int `json:"136"`
				Num168 int `json:"168"`
			} `json:"136"`
			Num138 struct {
				Num124 int `json:"124"`
				Num126 int `json:"126"`
				Num128 int `json:"128"`
				Num130 int `json:"130"`
				Num132 int `json:"132"`
				Num170 int `json:"170"`
			} `json:"138"`
			Num140 struct {
				Num78  int `json:"78"`
				Num126 int `json:"126"`
				Num132 int `json:"132"`
				Num170 int `json:"170"`
			} `json:"140"`
			Num142 struct {
				Num78  int `json:"78"`
				Num80  int `json:"80"`
				Num126 int `json:"126"`
				Num128 int `json:"128"`
				Num130 int `json:"130"`
				Num172 int `json:"172"`
			} `json:"142"`
			Num144 struct {
				Num78  int `json:"78"`
				Num80  int `json:"80"`
				Num128 int `json:"128"`
				Num172 int `json:"172"`
			} `json:"144"`
			Num146 struct {
				Num128 int `json:"128"`
				Num174 int `json:"174"`
			} `json:"146"`
			Num148 struct {
				Num78  int `json:"78"`
				Num126 int `json:"126"`
				Num174 int `json:"174"`
			} `json:"148"`
			Num150 struct {
				Num78  int `json:"78"`
				Num126 int `json:"126"`
			} `json:"150"`
			Num152 struct {
				Num124 int `json:"124"`
				Num174 int `json:"174"`
			} `json:"152"`
			Num154 struct {
				Num78  int `json:"78"`
				Num122 int `json:"122"`
				Num152 int `json:"152"`
				Num174 int `json:"174"`
			} `json:"154"`
			Num156 struct {
				Num80  int `json:"80"`
				Num120 int `json:"120"`
				Num154 int `json:"154"`
				Num156 int `json:"156"`
				Num176 int `json:"176"`
			} `json:"156"`
			Num158 struct {
				Num80  int `json:"80"`
				Num118 int `json:"118"`
				Num158 int `json:"158"`
				Num176 int `json:"176"`
			} `json:"158"`
			Num160 struct {
				Num118 int `json:"118"`
				Num160 int `json:"160"`
				Num176 int `json:"176"`
			} `json:"160"`
			Num162 struct {
				Num80  int `json:"80"`
				Num116 int `json:"116"`
				Num162 int `json:"162"`
				Num176 int `json:"176"`
			} `json:"162"`
			Num164 struct {
				Num80  int `json:"80"`
				Num114 int `json:"114"`
				Num164 int `json:"164"`
				Num176 int `json:"176"`
			} `json:"164"`
			Num166 struct {
				Num82  int `json:"82"`
				Num114 int `json:"114"`
				Num166 int `json:"166"`
				Num174 int `json:"174"`
			} `json:"166"`
			Num168 struct {
				Num80  int `json:"80"`
				Num82  int `json:"82"`
				Num84  int `json:"84"`
				Num94  int `json:"94"`
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num112 int `json:"112"`
				Num168 int `json:"168"`
			} `json:"168"`
			Num170 struct {
				Num80  int `json:"80"`
				Num82  int `json:"82"`
				Num84  int `json:"84"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num110 int `json:"110"`
				Num170 int `json:"170"`
				Num174 int `json:"174"`
			} `json:"170"`
			Num172 struct {
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
				Num108 int `json:"108"`
				Num170 int `json:"170"`
				Num174 int `json:"174"`
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
				Num172 int `json:"172"`
				Num174 int `json:"174"`
			} `json:"174"`
			Num176 struct {
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
				Num112 int `json:"112"`
				Num174 int `json:"174"`
			} `json:"176"`
			Num178 struct {
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
				Num172 int `json:"172"`
				Num176 int `json:"176"`
			} `json:"178"`
			Num180 struct {
				Num174 int `json:"174"`
				Num176 int `json:"176"`
				Num178 int `json:"178"`
			} `json:"180"`
			Num182 struct {
				Num174 int `json:"174"`
				Num176 int `json:"176"`
			} `json:"182"`
		} `json:"lane_pos,omitempty"`
		Obs6 struct {
		} `json:"obs,omitempty"`
		Purchase6 struct {
			Branches             int `json:"branches"`
			QuellingBlade        int `json:"quelling_blade"`
			Tango                int `json:"tango"`
			Circlet              int `json:"circlet"`
			Gauntlets            int `json:"gauntlets"`
			RecipeBracer         int `json:"recipe_bracer"`
			Bracer               int `json:"bracer"`
			Boots                int `json:"boots"`
			Gloves               int `json:"gloves"`
			BeltOfStrength       int `json:"belt_of_strength"`
			PowerTreads          int `json:"power_treads"`
			WardSentry           int `json:"ward_sentry"`
			Blink                int `json:"blink"`
			Tpscroll             int `json:"tpscroll"`
			AghanimsShard        int `json:"aghanims_shard"`
			RingOfProtection     int `json:"ring_of_protection"`
			RecipeSoulRing       int `json:"recipe_soul_ring"`
			Clarity              int `json:"clarity"`
			SoulRing             int `json:"soul_ring"`
			OgreAxe              int `json:"ogre_axe"`
			MithrilHammer        int `json:"mithril_hammer"`
			RecipeBlackKingBar   int `json:"recipe_black_king_bar"`
			BlackKingBar         int `json:"black_king_bar"`
			Sange                int `json:"sange"`
			RecipeSange          int `json:"recipe_sange"`
			RecipeHeavensHalberd int `json:"recipe_heavens_halberd"`
			TalismanOfEvasion    int `json:"talisman_of_evasion"`
			HeavensHalberd       int `json:"heavens_halberd"`
			Buckler              int `json:"buckler"`
			RecipeBuckler        int `json:"recipe_buckler"`
			RecipeAssault        int `json:"recipe_assault"`
			Hyperstone           int `json:"hyperstone"`
			Platemail            int `json:"platemail"`
			Assault              int `json:"assault"`
		} `json:"purchase,omitempty"`
		Sen6 struct {
			Num170 struct {
				Num86 int `json:"86"`
			} `json:"170"`
		} `json:"sen,omitempty"`
		XpReasons6 struct {
			Num0 int `json:"0"`
			Num1 int `json:"1"`
			Num2 int `json:"2"`
			Num3 int `json:"3"`
			Num5 int `json:"5"`
		} `json:"xp_reasons,omitempty"`
		PurchaseTime6 struct {
			Branches          int `json:"branches"`
			QuellingBlade     int `json:"quelling_blade"`
			Tango             int `json:"tango"`
			Circlet           int `json:"circlet"`
			Gauntlets         int `json:"gauntlets"`
			Bracer            int `json:"bracer"`
			Boots             int `json:"boots"`
			Gloves            int `json:"gloves"`
			BeltOfStrength    int `json:"belt_of_strength"`
			PowerTreads       int `json:"power_treads"`
			WardSentry        int `json:"ward_sentry"`
			Blink             int `json:"blink"`
			Tpscroll          int `json:"tpscroll"`
			AghanimsShard     int `json:"aghanims_shard"`
			RingOfProtection  int `json:"ring_of_protection"`
			Clarity           int `json:"clarity"`
			SoulRing          int `json:"soul_ring"`
			OgreAxe           int `json:"ogre_axe"`
			MithrilHammer     int `json:"mithril_hammer"`
			BlackKingBar      int `json:"black_king_bar"`
			Sange             int `json:"sange"`
			TalismanOfEvasion int `json:"talisman_of_evasion"`
			HeavensHalberd    int `json:"heavens_halberd"`
			Buckler           int `json:"buckler"`
			Hyperstone        int `json:"hyperstone"`
			Platemail         int `json:"platemail"`
			Assault           int `json:"assault"`
		} `json:"purchase_time,omitempty"`
		FirstPurchaseTime6 struct {
			Branches          int `json:"branches"`
			QuellingBlade     int `json:"quelling_blade"`
			Tango             int `json:"tango"`
			Circlet           int `json:"circlet"`
			Gauntlets         int `json:"gauntlets"`
			Bracer            int `json:"bracer"`
			Boots             int `json:"boots"`
			Gloves            int `json:"gloves"`
			BeltOfStrength    int `json:"belt_of_strength"`
			PowerTreads       int `json:"power_treads"`
			WardSentry        int `json:"ward_sentry"`
			Blink             int `json:"blink"`
			Tpscroll          int `json:"tpscroll"`
			AghanimsShard     int `json:"aghanims_shard"`
			RingOfProtection  int `json:"ring_of_protection"`
			Clarity           int `json:"clarity"`
			SoulRing          int `json:"soul_ring"`
			OgreAxe           int `json:"ogre_axe"`
			MithrilHammer     int `json:"mithril_hammer"`
			BlackKingBar      int `json:"black_king_bar"`
			Sange             int `json:"sange"`
			TalismanOfEvasion int `json:"talisman_of_evasion"`
			HeavensHalberd    int `json:"heavens_halberd"`
			Buckler           int `json:"buckler"`
			Hyperstone        int `json:"hyperstone"`
			Platemail         int `json:"platemail"`
			Assault           int `json:"assault"`
		} `json:"first_purchase_time,omitempty"`
		ItemWin6 struct {
			Branches          int `json:"branches"`
			QuellingBlade     int `json:"quelling_blade"`
			Tango             int `json:"tango"`
			Circlet           int `json:"circlet"`
			Gauntlets         int `json:"gauntlets"`
			Bracer            int `json:"bracer"`
			Boots             int `json:"boots"`
			Gloves            int `json:"gloves"`
			BeltOfStrength    int `json:"belt_of_strength"`
			PowerTreads       int `json:"power_treads"`
			WardSentry        int `json:"ward_sentry"`
			Blink             int `json:"blink"`
			Tpscroll          int `json:"tpscroll"`
			AghanimsShard     int `json:"aghanims_shard"`
			RingOfProtection  int `json:"ring_of_protection"`
			Clarity           int `json:"clarity"`
			SoulRing          int `json:"soul_ring"`
			OgreAxe           int `json:"ogre_axe"`
			MithrilHammer     int `json:"mithril_hammer"`
			BlackKingBar      int `json:"black_king_bar"`
			Sange             int `json:"sange"`
			TalismanOfEvasion int `json:"talisman_of_evasion"`
			HeavensHalberd    int `json:"heavens_halberd"`
			Buckler           int `json:"buckler"`
			Hyperstone        int `json:"hyperstone"`
			Platemail         int `json:"platemail"`
			Assault           int `json:"assault"`
		} `json:"item_win,omitempty"`
		ItemUsage6 struct {
			Branches          int `json:"branches"`
			QuellingBlade     int `json:"quelling_blade"`
			Tango             int `json:"tango"`
			Circlet           int `json:"circlet"`
			Gauntlets         int `json:"gauntlets"`
			Bracer            int `json:"bracer"`
			Boots             int `json:"boots"`
			Gloves            int `json:"gloves"`
			BeltOfStrength    int `json:"belt_of_strength"`
			PowerTreads       int `json:"power_treads"`
			WardSentry        int `json:"ward_sentry"`
			Blink             int `json:"blink"`
			Tpscroll          int `json:"tpscroll"`
			AghanimsShard     int `json:"aghanims_shard"`
			RingOfProtection  int `json:"ring_of_protection"`
			Clarity           int `json:"clarity"`
			SoulRing          int `json:"soul_ring"`
			OgreAxe           int `json:"ogre_axe"`
			MithrilHammer     int `json:"mithril_hammer"`
			BlackKingBar      int `json:"black_king_bar"`
			Sange             int `json:"sange"`
			TalismanOfEvasion int `json:"talisman_of_evasion"`
			HeavensHalberd    int `json:"heavens_halberd"`
			Buckler           int `json:"buckler"`
			Hyperstone        int `json:"hyperstone"`
			Platemail         int `json:"platemail"`
			Assault           int `json:"assault"`
		} `json:"item_usage,omitempty"`
		AbilityTargets7 struct {
			ChenPenitence struct {
				NpcDotaHeroEnchantress     int `json:"npc_dota_hero_enchantress"`
				NpcDotaHeroPuck            int `json:"npc_dota_hero_puck"`
				NpcDotaHeroTemplarAssassin int `json:"npc_dota_hero_templar_assassin"`
				NpcDotaHeroDeathProphet    int `json:"npc_dota_hero_death_prophet"`
			} `json:"chen_penitence"`
		} `json:"ability_targets,omitempty"`
		AbilityUses7 struct {
			PlusHighFive        int `json:"plus_high_five"`
			ChenHolyPersuasion  int `json:"chen_holy_persuasion"`
			ChenPenitence       int `json:"chen_penitence"`
			ChenHandOfGod       int `json:"chen_hand_of_god"`
			AbilityCapture      int `json:"ability_capture"`
			SeasonalTi11Balloon int `json:"seasonal_ti11_balloon"`
		} `json:"ability_uses,omitempty"`
		Actions7 struct {
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
			Num12 int `json:"12"`
			Num14 int `json:"14"`
			Num15 int `json:"15"`
			Num16 int `json:"16"`
			Num17 int `json:"17"`
			Num19 int `json:"19"`
			Num24 int `json:"24"`
			Num27 int `json:"27"`
			Num30 int `json:"30"`
			Num32 int `json:"32"`
			Num33 int `json:"33"`
			Num37 int `json:"37"`
			Num38 int `json:"38"`
		} `json:"actions,omitempty"`
		Damage7 struct {
			NpcDotaHeroDeathProphet                int `json:"npc_dota_hero_death_prophet"`
			NpcDotaSentryWards                     int `json:"npc_dota_sentry_wards"`
			NpcDotaCreepBadguysRanged              int `json:"npc_dota_creep_badguys_ranged"`
			NpcDotaHeroPrimalBeast                 int `json:"npc_dota_hero_primal_beast"`
			NpcDotaCreepBadguysMelee               int `json:"npc_dota_creep_badguys_melee"`
			NpcDotaNeutralKoboldTaskmaster         int `json:"npc_dota_neutral_kobold_taskmaster"`
			NpcDotaNeutralForestTrollBerserker     int `json:"npc_dota_neutral_forest_troll_berserker"`
			NpcDotaCreepGoodguysMelee              int `json:"npc_dota_creep_goodguys_melee"`
			NpcDotaNeutralHarpyStorm               int `json:"npc_dota_neutral_harpy_storm"`
			NpcDotaGoodguysTower1Bot               int `json:"npc_dota_goodguys_tower1_bot"`
			NpcDotaHeroEnchantress                 int `json:"npc_dota_hero_enchantress"`
			NpcDotaNeutralDarkTroll                int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaNeutralDarkTrollWarlord         int `json:"npc_dota_neutral_dark_troll_warlord"`
			NpcDotaDarkTrollWarlordSkeletonWarrior int `json:"npc_dota_dark_troll_warlord_skeleton_warrior"`
			NpcDotaNeutralGiantWolf                int `json:"npc_dota_neutral_giant_wolf"`
			NpcDotaNeutralWarpineRaider            int `json:"npc_dota_neutral_warpine_raider"`
			NpcDotaNeutralEnragedWildkin           int `json:"npc_dota_neutral_enraged_wildkin"`
			NpcDotaNeutralPolarFurbolgUrsaWarrior  int `json:"npc_dota_neutral_polar_furbolg_ursa_warrior"`
			NpcDotaNeutralCentaurOutrunner         int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaObserverWards                   int `json:"npc_dota_observer_wards"`
			NpcDotaNeutralCentaurKhan              int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaNeutralOgreMauler               int `json:"npc_dota_neutral_ogre_mauler"`
			NpcDotaCreepGoodguysFlagbearer         int `json:"npc_dota_creep_goodguys_flagbearer"`
			NpcDotaCreepGoodguysRanged             int `json:"npc_dota_creep_goodguys_ranged"`
			NpcDotaNeutralPolarFurbolgChampion     int `json:"npc_dota_neutral_polar_furbolg_champion"`
			NpcDotaHeroPuck                        int `json:"npc_dota_hero_puck"`
			NpcDotaNeutralWildkin                  int `json:"npc_dota_neutral_wildkin"`
			NpcDotaNeutralSatyrTrickster           int `json:"npc_dota_neutral_satyr_trickster"`
			NpcDotaNeutralSatyrSoulstealer         int `json:"npc_dota_neutral_satyr_soulstealer"`
			NpcDotaTemplarAssassinPsionicTrap      int `json:"npc_dota_templar_assassin_psionic_trap"`
			NpcDotaNeutralSatyrHellcaller          int `json:"npc_dota_neutral_satyr_hellcaller"`
			NpcDotaHeroTemplarAssassin             int `json:"npc_dota_hero_templar_assassin"`
			NpcDotaNeutralOgreMagi                 int `json:"npc_dota_neutral_ogre_magi"`
			NpcDotaNeutralMudGolem                 int `json:"npc_dota_neutral_mud_golem"`
			NpcDotaNeutralMudGolemSplit            int `json:"npc_dota_neutral_mud_golem_split"`
			NpcDotaGoodguysSiege                   int `json:"npc_dota_goodguys_siege"`
			NpcDotaGoodguysMeleeRaxMid             int `json:"npc_dota_goodguys_melee_rax_mid"`
			NpcDotaCourier                         int `json:"npc_dota_courier"`
			NpcDotaGoodguysRangeRaxTop             int `json:"npc_dota_goodguys_range_rax_top"`
			NpcDotaGoodguysMeleeRaxTop             int `json:"npc_dota_goodguys_melee_rax_top"`
			NpcDotaGoodguysFillers                 int `json:"npc_dota_goodguys_fillers"`
			NpcDotaGoodguysTower4                  int `json:"npc_dota_goodguys_tower4"`
		} `json:"damage,omitempty"`
		DamageInflictor7 struct {
			Null                     int `json:"null"`
			SatyrSoulstealerManaBurn int `json:"satyr_soulstealer_mana_burn"`
			MudGolemHurlBoulder      int `json:"mud_golem_hurl_boulder"`
			WarpineRaiderSeedShot    int `json:"warpine_raider_seed_shot"`
			CentaurKhanWarStomp      int `json:"centaur_khan_war_stomp"`
			WraithPact               int `json:"wraith_pact"`
		} `json:"damage_inflictor,omitempty"`
		DamageInflictorReceived7 struct {
			Null                     int `json:"null"`
			DeathProphetCarrionSwarm int `json:"death_prophet_carrion_swarm"`
			PrimalBeastTrample       int `json:"primal_beast_trample"`
			DeathProphetSpiritSiphon int `json:"death_prophet_spirit_siphon"`
			DeathProphetExorcism     int `json:"death_prophet_exorcism"`
			PrimalBeastPulverize     int `json:"primal_beast_pulverize"`
			EnchantressImpetus       int `json:"enchantress_impetus"`
			PuckWaningRift           int `json:"puck_waning_rift"`
			PuckIllusoryOrb          int `json:"puck_illusory_orb"`
			Dust                     int `json:"dust"`
			PuckDreamCoil            int `json:"puck_dream_coil"`
			WitchBlade               int `json:"witch_blade"`
			TemplarAssassinMeld      int `json:"templar_assassin_meld"`
			PrimalBeastRockThrow     int `json:"primal_beast_rock_throw"`
			PrimalBeastOnslaught     int `json:"primal_beast_onslaught"`
			SatyrHellcallerShockwave int `json:"satyr_hellcaller_shockwave"`
			TemplarAssassinTrap      int `json:"templar_assassin_trap"`
			OverwhelmingBlink        int `json:"overwhelming_blink"`
		} `json:"damage_inflictor_received,omitempty"`
		DamageTaken7 struct {
			NpcDotaHeroPrimalBeast                 int `json:"npc_dota_hero_primal_beast"`
			NpcDotaHeroDeathProphet                int `json:"npc_dota_hero_death_prophet"`
			NpcDotaCreepGoodguysMelee              int `json:"npc_dota_creep_goodguys_melee"`
			NpcDotaNeutralForestTrollBerserker     int `json:"npc_dota_neutral_forest_troll_berserker"`
			NpcDotaNeutralKoboldTaskmaster         int `json:"npc_dota_neutral_kobold_taskmaster"`
			NpcDotaNeutralSatyrTrickster           int `json:"npc_dota_neutral_satyr_trickster"`
			NpcDotaNeutralHarpyStorm               int `json:"npc_dota_neutral_harpy_storm"`
			NpcDotaNeutralMudGolem                 int `json:"npc_dota_neutral_mud_golem"`
			NpcDotaNeutralDarkTrollWarlord         int `json:"npc_dota_neutral_dark_troll_warlord"`
			NpcDotaDarkTrollWarlordSkeletonWarrior int `json:"npc_dota_dark_troll_warlord_skeleton_warrior"`
			NpcDotaHeroEnchantress                 int `json:"npc_dota_hero_enchantress"`
			NpcDotaHeroPuck                        int `json:"npc_dota_hero_puck"`
			NpcDotaHeroTemplarAssassin             int `json:"npc_dota_hero_templar_assassin"`
			NpcDotaNeutralSatyrHellcaller          int `json:"npc_dota_neutral_satyr_hellcaller"`
			NpcDotaNeutralPolarFurbolgUrsaWarrior  int `json:"npc_dota_neutral_polar_furbolg_ursa_warrior"`
			NpcDotaNeutralPolarFurbolgChampion     int `json:"npc_dota_neutral_polar_furbolg_champion"`
			NpcDotaNeutralOgreMagi                 int `json:"npc_dota_neutral_ogre_magi"`
			NpcDotaNeutralOgreMauler               int `json:"npc_dota_neutral_ogre_mauler"`
			NpcDotaNeutralCentaurOutrunner         int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaNeutralCentaurKhan              int `json:"npc_dota_neutral_centaur_khan"`
		} `json:"damage_taken,omitempty"`
		DamageTargets7 struct {
			Null struct {
				NpcDotaHeroDeathProphet    int `json:"npc_dota_hero_death_prophet"`
				NpcDotaHeroPrimalBeast     int `json:"npc_dota_hero_primal_beast"`
				NpcDotaHeroEnchantress     int `json:"npc_dota_hero_enchantress"`
				NpcDotaHeroPuck            int `json:"npc_dota_hero_puck"`
				NpcDotaHeroTemplarAssassin int `json:"npc_dota_hero_templar_assassin"`
			} `json:"null"`
			SatyrSoulstealerManaBurn struct {
				NpcDotaHeroDeathProphet int `json:"npc_dota_hero_death_prophet"`
			} `json:"satyr_soulstealer_mana_burn"`
			MudGolemHurlBoulder struct {
				NpcDotaHeroDeathProphet int `json:"npc_dota_hero_death_prophet"`
			} `json:"mud_golem_hurl_boulder"`
			WarpineRaiderSeedShot struct {
				NpcDotaHeroEnchantress int `json:"npc_dota_hero_enchantress"`
				NpcDotaHeroPrimalBeast int `json:"npc_dota_hero_primal_beast"`
			} `json:"warpine_raider_seed_shot"`
			CentaurKhanWarStomp struct {
				NpcDotaHeroPuck int `json:"npc_dota_hero_puck"`
			} `json:"centaur_khan_war_stomp"`
			WraithPact struct {
				NpcDotaHeroPuck            int `json:"npc_dota_hero_puck"`
				NpcDotaHeroEnchantress     int `json:"npc_dota_hero_enchantress"`
				NpcDotaHeroDeathProphet    int `json:"npc_dota_hero_death_prophet"`
				NpcDotaHeroPrimalBeast     int `json:"npc_dota_hero_primal_beast"`
				NpcDotaHeroTemplarAssassin int `json:"npc_dota_hero_templar_assassin"`
			} `json:"wraith_pact"`
		} `json:"damage_targets,omitempty"`
		GoldReasons6 struct {
			Num0  int `json:"0"`
			Num1  int `json:"1"`
			Num6  int `json:"6"`
			Num11 int `json:"11"`
			Num12 int `json:"12"`
			Num13 int `json:"13"`
			Num14 int `json:"14"`
			Num15 int `json:"15"`
			Num16 int `json:"16"`
			Num17 int `json:"17"`
			Num19 int `json:"19"`
			Num20 int `json:"20"`
			Num21 int `json:"21"`
		} `json:"gold_reasons,omitempty"`
		HeroHits7 struct {
			Null                     int `json:"null"`
			SatyrSoulstealerManaBurn int `json:"satyr_soulstealer_mana_burn"`
			MudGolemHurlBoulder      int `json:"mud_golem_hurl_boulder"`
			WarpineRaiderSeedShot    int `json:"warpine_raider_seed_shot"`
			CentaurKhanWarStomp      int `json:"centaur_khan_war_stomp"`
			WraithPact               int `json:"wraith_pact"`
		} `json:"hero_hits,omitempty"`
		ItemUses7 struct {
			WardDispenser   int `json:"ward_dispenser"`
			WardSentry      int `json:"ward_sentry"`
			Tango           int `json:"tango"`
			WardObserver    int `json:"ward_observer"`
			Tpscroll        int `json:"tpscroll"`
			ArcaneRing      int `json:"arcane_ring"`
			TomeOfKnowledge int `json:"tome_of_knowledge"`
			Clarity         int `json:"clarity"`
			Mekansm         int `json:"mekansm"`
			Dust            int `json:"dust"`
			SmokeOfDeceit   int `json:"smoke_of_deceit"`
			WraithPact      int `json:"wraith_pact"`
			Ghost           int `json:"ghost"`
			BootsOfBearing  int `json:"boots_of_bearing"`
		} `json:"item_uses,omitempty"`
		Killed7 struct {
			NpcDotaSentryWards                     int `json:"npc_dota_sentry_wards"`
			NpcDotaCreepBadguysRanged              int `json:"npc_dota_creep_badguys_ranged"`
			NpcDotaNeutralForestTrollBerserker     int `json:"npc_dota_neutral_forest_troll_berserker"`
			NpcDotaNeutralKoboldTaskmaster         int `json:"npc_dota_neutral_kobold_taskmaster"`
			NpcDotaCreepGoodguysMelee              int `json:"npc_dota_creep_goodguys_melee"`
			NpcDotaDarkTrollWarlordSkeletonWarrior int `json:"npc_dota_dark_troll_warlord_skeleton_warrior"`
			NpcDotaNeutralGiantWolf                int `json:"npc_dota_neutral_giant_wolf"`
			NpcDotaNeutralWarpineRaider            int `json:"npc_dota_neutral_warpine_raider"`
			NpcDotaNeutralDarkTrollWarlord         int `json:"npc_dota_neutral_dark_troll_warlord"`
			NpcDotaNeutralPolarFurbolgUrsaWarrior  int `json:"npc_dota_neutral_polar_furbolg_ursa_warrior"`
			NpcDotaObserverWards                   int `json:"npc_dota_observer_wards"`
			NpcDotaCreepGoodguysFlagbearer         int `json:"npc_dota_creep_goodguys_flagbearer"`
			NpcDotaNeutralCentaurOutrunner         int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaNeutralOgreMauler               int `json:"npc_dota_neutral_ogre_mauler"`
			NpcDotaNeutralDarkTroll                int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaNeutralWildkin                  int `json:"npc_dota_neutral_wildkin"`
			NpcDotaCreepGoodguysRanged             int `json:"npc_dota_creep_goodguys_ranged"`
			NpcDotaNeutralSatyrSoulstealer         int `json:"npc_dota_neutral_satyr_soulstealer"`
			NpcDotaNeutralSatyrTrickster           int `json:"npc_dota_neutral_satyr_trickster"`
			NpcDotaTemplarAssassinPsionicTrap      int `json:"npc_dota_templar_assassin_psionic_trap"`
			NpcDotaNeutralSatyrHellcaller          int `json:"npc_dota_neutral_satyr_hellcaller"`
			NpcDotaNeutralMudGolemSplit            int `json:"npc_dota_neutral_mud_golem_split"`
			NpcDotaNeutralEnragedWildkin           int `json:"npc_dota_neutral_enraged_wildkin"`
			NpcDotaCourier                         int `json:"npc_dota_courier"`
			NpcDotaGoodguysMeleeRaxTop             int `json:"npc_dota_goodguys_melee_rax_top"`
		} `json:"killed,omitempty"`
		KilledBy6 struct {
			NpcDotaHeroPuck            int `json:"npc_dota_hero_puck"`
			NpcDotaHeroTemplarAssassin int `json:"npc_dota_hero_templar_assassin"`
		} `json:"killed_by,omitempty"`
		LanePos7 struct {
			Num78 struct {
				Num150 int `json:"150"`
				Num152 int `json:"152"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
			} `json:"78"`
			Num80 struct {
				Num154 int `json:"154"`
				Num156 int `json:"156"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
			} `json:"80"`
			Num82 struct {
				Num156 int `json:"156"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
			} `json:"82"`
			Num84 struct {
				Num156 int `json:"156"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
			} `json:"84"`
			Num86 struct {
				Num156 int `json:"156"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
				Num186 int `json:"186"`
			} `json:"86"`
			Num88 struct {
				Num156 int `json:"156"`
				Num164 int `json:"164"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
				Num180 int `json:"180"`
				Num182 int `json:"182"`
				Num186 int `json:"186"`
			} `json:"88"`
			Num90 struct {
				Num164 int `json:"164"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
				Num180 int `json:"180"`
			} `json:"90"`
			Num92 struct {
				Num164 int `json:"164"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
				Num174 int `json:"174"`
				Num182 int `json:"182"`
			} `json:"92"`
			Num94 struct {
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
				Num174 int `json:"174"`
				Num178 int `json:"178"`
				Num180 int `json:"180"`
			} `json:"94"`
			Num96 struct {
				Num162 int `json:"162"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
				Num174 int `json:"174"`
			} `json:"96"`
			Num98 struct {
				Num162 int `json:"162"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
				Num174 int `json:"174"`
				Num176 int `json:"176"`
			} `json:"98"`
			Num100 struct {
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
				Num172 int `json:"172"`
				Num174 int `json:"174"`
				Num176 int `json:"176"`
			} `json:"100"`
			Num102 struct {
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
				Num174 int `json:"174"`
			} `json:"102"`
			Num104 struct {
				Num160 int `json:"160"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
			} `json:"104"`
			Num106 struct {
				Num156 int `json:"156"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num166 int `json:"166"`
			} `json:"106"`
			Num108 struct {
				Num156 int `json:"156"`
				Num158 int `json:"158"`
				Num166 int `json:"166"`
			} `json:"108"`
			Num110 struct {
				Num158 int `json:"158"`
				Num160 int `json:"160"`
			} `json:"110"`
			Num112 struct {
				Num156 int `json:"156"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num166 int `json:"166"`
			} `json:"112"`
			Num114 struct {
				Num138 int `json:"138"`
				Num140 int `json:"140"`
				Num156 int `json:"156"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num166 int `json:"166"`
			} `json:"114"`
			Num116 struct {
				Num140 int `json:"140"`
				Num142 int `json:"142"`
				Num144 int `json:"144"`
				Num156 int `json:"156"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
			} `json:"116"`
			Num118 struct {
				Num142 int `json:"142"`
				Num144 int `json:"144"`
				Num146 int `json:"146"`
				Num148 int `json:"148"`
				Num154 int `json:"154"`
				Num156 int `json:"156"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
			} `json:"118"`
			Num120 struct {
				Num150 int `json:"150"`
				Num152 int `json:"152"`
				Num154 int `json:"154"`
				Num156 int `json:"156"`
				Num158 int `json:"158"`
				Num162 int `json:"162"`
			} `json:"120"`
			Num122 struct {
				Num152 int `json:"152"`
				Num154 int `json:"154"`
				Num156 int `json:"156"`
				Num162 int `json:"162"`
			} `json:"122"`
			Num124 struct {
				Num146 int `json:"146"`
				Num148 int `json:"148"`
				Num150 int `json:"150"`
				Num152 int `json:"152"`
				Num154 int `json:"154"`
				Num156 int `json:"156"`
				Num162 int `json:"162"`
			} `json:"124"`
			Num126 struct {
				Num144 int `json:"144"`
				Num146 int `json:"146"`
				Num148 int `json:"148"`
				Num150 int `json:"150"`
				Num156 int `json:"156"`
				Num164 int `json:"164"`
			} `json:"126"`
			Num128 struct {
				Num164 int `json:"164"`
			} `json:"128"`
			Num130 struct {
				Num166 int `json:"166"`
			} `json:"130"`
			Num132 struct {
				Num166 int `json:"166"`
			} `json:"132"`
			Num134 struct {
				Num168 int `json:"168"`
			} `json:"134"`
			Num136 struct {
				Num120 int `json:"120"`
				Num170 int `json:"170"`
			} `json:"136"`
			Num138 struct {
				Num114 int `json:"114"`
				Num116 int `json:"116"`
				Num122 int `json:"122"`
				Num124 int `json:"124"`
				Num144 int `json:"144"`
				Num146 int `json:"146"`
				Num148 int `json:"148"`
				Num150 int `json:"150"`
				Num170 int `json:"170"`
			} `json:"138"`
			Num140 struct {
				Num110 int `json:"110"`
				Num112 int `json:"112"`
				Num126 int `json:"126"`
				Num128 int `json:"128"`
				Num140 int `json:"140"`
				Num142 int `json:"142"`
				Num150 int `json:"150"`
				Num172 int `json:"172"`
			} `json:"140"`
			Num142 struct {
				Num90  int `json:"90"`
				Num94  int `json:"94"`
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num104 int `json:"104"`
				Num106 int `json:"106"`
				Num128 int `json:"128"`
				Num138 int `json:"138"`
				Num150 int `json:"150"`
			} `json:"142"`
			Num144 struct {
				Num90  int `json:"90"`
				Num94  int `json:"94"`
				Num96  int `json:"96"`
				Num98  int `json:"98"`
				Num102 int `json:"102"`
				Num136 int `json:"136"`
				Num150 int `json:"150"`
				Num172 int `json:"172"`
			} `json:"144"`
			Num146 struct {
				Num90  int `json:"90"`
				Num96  int `json:"96"`
				Num128 int `json:"128"`
				Num132 int `json:"132"`
				Num150 int `json:"150"`
				Num172 int `json:"172"`
			} `json:"146"`
			Num148 struct {
				Num88  int `json:"88"`
				Num90  int `json:"90"`
				Num98  int `json:"98"`
				Num126 int `json:"126"`
				Num128 int `json:"128"`
				Num130 int `json:"130"`
				Num134 int `json:"134"`
				Num174 int `json:"174"`
			} `json:"148"`
			Num150 struct {
				Num88  int `json:"88"`
				Num90  int `json:"90"`
				Num98  int `json:"98"`
				Num136 int `json:"136"`
				Num174 int `json:"174"`
			} `json:"150"`
			Num152 struct {
				Num98  int `json:"98"`
				Num126 int `json:"126"`
				Num138 int `json:"138"`
				Num174 int `json:"174"`
			} `json:"152"`
			Num154 struct {
				Num88  int `json:"88"`
				Num90  int `json:"90"`
				Num96  int `json:"96"`
				Num138 int `json:"138"`
				Num176 int `json:"176"`
			} `json:"154"`
			Num156 struct {
				Num86  int `json:"86"`
				Num90  int `json:"90"`
				Num94  int `json:"94"`
				Num176 int `json:"176"`
			} `json:"156"`
			Num158 struct {
				Num86 int `json:"86"`
				Num92 int `json:"92"`
				Num94 int `json:"94"`
			} `json:"158"`
			Num160 struct {
				Num86  int `json:"86"`
				Num90  int `json:"90"`
				Num92  int `json:"92"`
				Num94  int `json:"94"`
				Num176 int `json:"176"`
			} `json:"160"`
			Num162 struct {
				Num86  int `json:"86"`
				Num90  int `json:"90"`
				Num96  int `json:"96"`
				Num176 int `json:"176"`
			} `json:"162"`
			Num164 struct {
				Num86  int `json:"86"`
				Num88  int `json:"88"`
				Num96  int `json:"96"`
				Num176 int `json:"176"`
			} `json:"164"`
			Num166 struct {
				Num98  int `json:"98"`
				Num176 int `json:"176"`
			} `json:"166"`
			Num168 struct {
				Num100 int `json:"100"`
				Num176 int `json:"176"`
			} `json:"168"`
			Num170 struct {
				Num100 int `json:"100"`
				Num104 int `json:"104"`
				Num106 int `json:"106"`
				Num108 int `json:"108"`
				Num174 int `json:"174"`
			} `json:"170"`
			Num172 struct {
				Num104 int `json:"104"`
				Num108 int `json:"108"`
				Num174 int `json:"174"`
			} `json:"172"`
			Num176 struct {
				Num174 int `json:"174"`
			} `json:"176"`
			Num178 struct {
				Num174 int `json:"174"`
			} `json:"178"`
			Num180 struct {
				Num176 int `json:"176"`
			} `json:"180"`
			Num182 struct {
				Num176 int `json:"176"`
			} `json:"182"`
		} `json:"lane_pos,omitempty"`
		Obs7 struct {
			Num76 struct {
				Num144 int `json:"144"`
			} `json:"76"`
			Num100 struct {
				Num78 int `json:"78"`
			} `json:"100"`
			Num106 struct {
				Num156 int `json:"156"`
			} `json:"106"`
			Num112 struct {
				Num146 int `json:"146"`
			} `json:"112"`
			Num122 struct {
				Num144 int `json:"144"`
			} `json:"122"`
			Num130 struct {
				Num106 int `json:"106"`
			} `json:"130"`
			Num132 struct {
				Num160 int `json:"160"`
			} `json:"132"`
			Num142 struct {
				Num92  int `json:"92"`
				Num152 int `json:"152"`
			} `json:"142"`
		} `json:"obs,omitempty"`
		Purchase7 struct {
			WardDispenser        int `json:"ward_dispenser"`
			Headdress            int `json:"headdress"`
			Tango                int `json:"tango"`
			WardSentry           int `json:"ward_sentry"`
			WardObserver         int `json:"ward_observer"`
			Boots                int `json:"boots"`
			SmokeOfDeceit        int `json:"smoke_of_deceit"`
			Chainmail            int `json:"chainmail"`
			TomeOfKnowledge      int `json:"tome_of_knowledge"`
			RecipeMekansm        int `json:"recipe_mekansm"`
			Mekansm              int `json:"mekansm"`
			Clarity              int `json:"clarity"`
			EnchantedMango       int `json:"enchanted_mango"`
			Dust                 int `json:"dust"`
			RingOfProtection     int `json:"ring_of_protection"`
			Buckler              int `json:"buckler"`
			RecipeBuckler        int `json:"recipe_buckler"`
			SobiMask             int `json:"sobi_mask"`
			RingOfBasilius       int `json:"ring_of_basilius"`
			RecipeRingOfBasilius int `json:"recipe_ring_of_basilius"`
			Lifesteal            int `json:"lifesteal"`
			BladesOfAttack       int `json:"blades_of_attack"`
			RecipeVladmir        int `json:"recipe_vladmir"`
			Vladmir              int `json:"vladmir"`
			RecipeWraithPact     int `json:"recipe_wraith_pact"`
			PointBooster         int `json:"point_booster"`
			WraithPact           int `json:"wraith_pact"`
			Ghost                int `json:"ghost"`
			WindLace             int `json:"wind_lace"`
			RingOfRegen          int `json:"ring_of_regen"`
			BeltOfStrength       int `json:"belt_of_strength"`
			Robe                 int `json:"robe"`
			AncientJanggo        int `json:"ancient_janggo"`
			RecipeAncientJanggo  int `json:"recipe_ancient_janggo"`
			RecipeBootsOfBearing int `json:"recipe_boots_of_bearing"`
			TranquilBoots        int `json:"tranquil_boots"`
			BootsOfBearing       int `json:"boots_of_bearing"`
		} `json:"purchase,omitempty"`
		Sen7 struct {
			Num72 struct {
				Num104 int `json:"104"`
			} `json:"72"`
			Num74 struct {
				Num110 int `json:"110"`
			} `json:"74"`
			Num80 struct {
				Num106 int `json:"106"`
			} `json:"80"`
			Num84 struct {
				Num168 int `json:"168"`
			} `json:"84"`
			Num96 struct {
				Num140 int `json:"140"`
			} `json:"96"`
			Num108 struct {
				Num158 int `json:"158"`
			} `json:"108"`
			Num110 struct {
				Num156 int `json:"156"`
			} `json:"110"`
			Num112 struct {
				Num144 int `json:"144"`
			} `json:"112"`
			Num116 struct {
				Num162 int `json:"162"`
			} `json:"116"`
			Num120 struct {
				Num144 int `json:"144"`
				Num156 int `json:"156"`
			} `json:"120"`
			Num122 struct {
				Num144 int `json:"144"`
				Num150 int `json:"150"`
				Num164 int `json:"164"`
			} `json:"122"`
			Num126 struct {
				Num146 int `json:"146"`
			} `json:"126"`
			Num128 struct {
				Num126 int `json:"126"`
			} `json:"128"`
			Num132 struct {
				Num118 int `json:"118"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
			} `json:"132"`
			Num134 struct {
				Num172 int `json:"172"`
			} `json:"134"`
			Num140 struct {
				Num136 int `json:"136"`
			} `json:"140"`
			Num142 struct {
				Num152 int `json:"152"`
			} `json:"142"`
			Num170 struct {
				Num152 int `json:"152"`
			} `json:"170"`
		} `json:"sen,omitempty"`
		PurchaseTime7 struct {
			Headdress        int `json:"headdress"`
			Tango            int `json:"tango"`
			WardSentry       int `json:"ward_sentry"`
			WardObserver     int `json:"ward_observer"`
			Boots            int `json:"boots"`
			SmokeOfDeceit    int `json:"smoke_of_deceit"`
			Chainmail        int `json:"chainmail"`
			TomeOfKnowledge  int `json:"tome_of_knowledge"`
			Mekansm          int `json:"mekansm"`
			Clarity          int `json:"clarity"`
			EnchantedMango   int `json:"enchanted_mango"`
			Dust             int `json:"dust"`
			RingOfProtection int `json:"ring_of_protection"`
			Buckler          int `json:"buckler"`
			SobiMask         int `json:"sobi_mask"`
			RingOfBasilius   int `json:"ring_of_basilius"`
			Lifesteal        int `json:"lifesteal"`
			BladesOfAttack   int `json:"blades_of_attack"`
			Vladmir          int `json:"vladmir"`
			PointBooster     int `json:"point_booster"`
			WraithPact       int `json:"wraith_pact"`
			Ghost            int `json:"ghost"`
			WindLace         int `json:"wind_lace"`
			RingOfRegen      int `json:"ring_of_regen"`
			BeltOfStrength   int `json:"belt_of_strength"`
			Robe             int `json:"robe"`
			AncientJanggo    int `json:"ancient_janggo"`
			TranquilBoots    int `json:"tranquil_boots"`
			BootsOfBearing   int `json:"boots_of_bearing"`
		} `json:"purchase_time,omitempty"`
		FirstPurchaseTime7 struct {
			Headdress        int `json:"headdress"`
			Tango            int `json:"tango"`
			WardSentry       int `json:"ward_sentry"`
			WardObserver     int `json:"ward_observer"`
			Boots            int `json:"boots"`
			SmokeOfDeceit    int `json:"smoke_of_deceit"`
			Chainmail        int `json:"chainmail"`
			TomeOfKnowledge  int `json:"tome_of_knowledge"`
			Mekansm          int `json:"mekansm"`
			Clarity          int `json:"clarity"`
			EnchantedMango   int `json:"enchanted_mango"`
			Dust             int `json:"dust"`
			RingOfProtection int `json:"ring_of_protection"`
			Buckler          int `json:"buckler"`
			SobiMask         int `json:"sobi_mask"`
			RingOfBasilius   int `json:"ring_of_basilius"`
			Lifesteal        int `json:"lifesteal"`
			BladesOfAttack   int `json:"blades_of_attack"`
			Vladmir          int `json:"vladmir"`
			PointBooster     int `json:"point_booster"`
			WraithPact       int `json:"wraith_pact"`
			Ghost            int `json:"ghost"`
			WindLace         int `json:"wind_lace"`
			RingOfRegen      int `json:"ring_of_regen"`
			BeltOfStrength   int `json:"belt_of_strength"`
			Robe             int `json:"robe"`
			AncientJanggo    int `json:"ancient_janggo"`
			TranquilBoots    int `json:"tranquil_boots"`
			BootsOfBearing   int `json:"boots_of_bearing"`
		} `json:"first_purchase_time,omitempty"`
		ItemWin7 struct {
			Headdress        int `json:"headdress"`
			Tango            int `json:"tango"`
			WardSentry       int `json:"ward_sentry"`
			WardObserver     int `json:"ward_observer"`
			Boots            int `json:"boots"`
			SmokeOfDeceit    int `json:"smoke_of_deceit"`
			Chainmail        int `json:"chainmail"`
			TomeOfKnowledge  int `json:"tome_of_knowledge"`
			Mekansm          int `json:"mekansm"`
			Clarity          int `json:"clarity"`
			EnchantedMango   int `json:"enchanted_mango"`
			Dust             int `json:"dust"`
			RingOfProtection int `json:"ring_of_protection"`
			Buckler          int `json:"buckler"`
			SobiMask         int `json:"sobi_mask"`
			RingOfBasilius   int `json:"ring_of_basilius"`
			Lifesteal        int `json:"lifesteal"`
			BladesOfAttack   int `json:"blades_of_attack"`
			Vladmir          int `json:"vladmir"`
			PointBooster     int `json:"point_booster"`
			WraithPact       int `json:"wraith_pact"`
			Ghost            int `json:"ghost"`
			WindLace         int `json:"wind_lace"`
			RingOfRegen      int `json:"ring_of_regen"`
			BeltOfStrength   int `json:"belt_of_strength"`
			Robe             int `json:"robe"`
			AncientJanggo    int `json:"ancient_janggo"`
			TranquilBoots    int `json:"tranquil_boots"`
			BootsOfBearing   int `json:"boots_of_bearing"`
		} `json:"item_win,omitempty"`
		ItemUsage7 struct {
			Headdress        int `json:"headdress"`
			Tango            int `json:"tango"`
			WardSentry       int `json:"ward_sentry"`
			WardObserver     int `json:"ward_observer"`
			Boots            int `json:"boots"`
			SmokeOfDeceit    int `json:"smoke_of_deceit"`
			Chainmail        int `json:"chainmail"`
			TomeOfKnowledge  int `json:"tome_of_knowledge"`
			Mekansm          int `json:"mekansm"`
			Clarity          int `json:"clarity"`
			EnchantedMango   int `json:"enchanted_mango"`
			Dust             int `json:"dust"`
			RingOfProtection int `json:"ring_of_protection"`
			Buckler          int `json:"buckler"`
			SobiMask         int `json:"sobi_mask"`
			RingOfBasilius   int `json:"ring_of_basilius"`
			Lifesteal        int `json:"lifesteal"`
			BladesOfAttack   int `json:"blades_of_attack"`
			Vladmir          int `json:"vladmir"`
			PointBooster     int `json:"point_booster"`
			WraithPact       int `json:"wraith_pact"`
			Ghost            int `json:"ghost"`
			WindLace         int `json:"wind_lace"`
			RingOfRegen      int `json:"ring_of_regen"`
			BeltOfStrength   int `json:"belt_of_strength"`
			Robe             int `json:"robe"`
			AncientJanggo    int `json:"ancient_janggo"`
			TranquilBoots    int `json:"tranquil_boots"`
			BootsOfBearing   int `json:"boots_of_bearing"`
		} `json:"item_usage,omitempty"`
		AbilityTargets8 struct {
			TrollWarlordWhirlingAxesRanged struct {
				NpcDotaHeroDeathProphet int `json:"npc_dota_hero_death_prophet"`
				NpcDotaHeroPrimalBeast  int `json:"npc_dota_hero_primal_beast"`
				NpcDotaHeroEnchantress  int `json:"npc_dota_hero_enchantress"`
			} `json:"troll_warlord_whirling_axes_ranged"`
		} `json:"ability_targets,omitempty"`
		AbilityUses8 struct {
			PlusGuildBanner                int `json:"plus_guild_banner"`
			TrollWarlordWhirlingAxesRanged int `json:"troll_warlord_whirling_axes_ranged"`
			TrollWarlordBerserkersRage     int `json:"troll_warlord_berserkers_rage"`
			TrollWarlordWhirlingAxesMelee  int `json:"troll_warlord_whirling_axes_melee"`
			TrollWarlordBattleTrance       int `json:"troll_warlord_battle_trance"`
			PlusHighFive                   int `json:"plus_high_five"`
		} `json:"ability_uses,omitempty"`
		Actions8 struct {
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
			Num13 int `json:"13"`
			Num14 int `json:"14"`
			Num15 int `json:"15"`
			Num16 int `json:"16"`
			Num17 int `json:"17"`
			Num19 int `json:"19"`
			Num20 int `json:"20"`
			Num27 int `json:"27"`
			Num31 int `json:"31"`
			Num33 int `json:"33"`
			Num37 int `json:"37"`
			Num38 int `json:"38"`
		} `json:"actions,omitempty"`
		Damage8 struct {
			NpcDotaCreepGoodguysRanged             int `json:"npc_dota_creep_goodguys_ranged"`
			NpcDotaCreepGoodguysMelee              int `json:"npc_dota_creep_goodguys_melee"`
			NpcDotaHeroPrimalBeast                 int `json:"npc_dota_hero_primal_beast"`
			NpcDotaCreepBadguysRanged              int `json:"npc_dota_creep_badguys_ranged"`
			NpcDotaHeroDeathProphet                int `json:"npc_dota_hero_death_prophet"`
			NpcDotaCreepBadguysMelee               int `json:"npc_dota_creep_badguys_melee"`
			NpcDotaCreepGoodguysFlagbearer         int `json:"npc_dota_creep_goodguys_flagbearer"`
			NpcDotaGoodguysSiege                   int `json:"npc_dota_goodguys_siege"`
			NpcDotaNeutralSatyrTrickster           int `json:"npc_dota_neutral_satyr_trickster"`
			NpcDotaNeutralSatyrSoulstealer         int `json:"npc_dota_neutral_satyr_soulstealer"`
			NpcDotaNeutralPolarFurbolgUrsaWarrior  int `json:"npc_dota_neutral_polar_furbolg_ursa_warrior"`
			NpcDotaNeutralPolarFurbolgChampion     int `json:"npc_dota_neutral_polar_furbolg_champion"`
			NpcDotaNeutralDarkTroll                int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaNeutralWarpineRaider            int `json:"npc_dota_neutral_warpine_raider"`
			NpcDotaNeutralDarkTrollWarlord         int `json:"npc_dota_neutral_dark_troll_warlord"`
			NpcDotaDarkTrollWarlordSkeletonWarrior int `json:"npc_dota_dark_troll_warlord_skeleton_warrior"`
			NpcDotaNeutralAlphaWolf                int `json:"npc_dota_neutral_alpha_wolf"`
			NpcDotaNeutralGiantWolf                int `json:"npc_dota_neutral_giant_wolf"`
			NpcDotaNeutralSatyrHellcaller          int `json:"npc_dota_neutral_satyr_hellcaller"`
			NpcDotaNeutralWildkin                  int `json:"npc_dota_neutral_wildkin"`
			NpcDotaNeutralEnragedWildkin           int `json:"npc_dota_neutral_enraged_wildkin"`
			NpcDotaNeutralFelBeast                 int `json:"npc_dota_neutral_fel_beast"`
			NpcDotaNeutralGhost                    int `json:"npc_dota_neutral_ghost"`
			NpcDotaNeutralMudGolem                 int `json:"npc_dota_neutral_mud_golem"`
			NpcDotaNeutralMudGolemSplit            int `json:"npc_dota_neutral_mud_golem_split"`
			NpcDotaNeutralIceShaman                int `json:"npc_dota_neutral_ice_shaman"`
			NpcDotaNeutralFrostbittenGolem         int `json:"npc_dota_neutral_frostbitten_golem"`
			NpcDotaNeutralKoboldTaskmaster         int `json:"npc_dota_neutral_kobold_taskmaster"`
			NpcDotaNeutralKoboldTunneler           int `json:"npc_dota_neutral_kobold_tunneler"`
			NpcDotaNeutralRockGolem                int `json:"npc_dota_neutral_rock_golem"`
			NpcDotaNeutralGraniteGolem             int `json:"npc_dota_neutral_granite_golem"`
			NpcDotaCreepBadguysFlagbearer          int `json:"npc_dota_creep_badguys_flagbearer"`
			NpcDotaNeutralHarpyScout               int `json:"npc_dota_neutral_harpy_scout"`
			NpcDotaNeutralHarpyStorm               int `json:"npc_dota_neutral_harpy_storm"`
			NpcDotaNeutralBigThunderLizard         int `json:"npc_dota_neutral_big_thunder_lizard"`
			NpcDotaNeutralSmallThunderLizard       int `json:"npc_dota_neutral_small_thunder_lizard"`
			NpcDotaNeutralCentaurOutrunner         int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaNeutralCentaurKhan              int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaNeutralForestTrollBerserker     int `json:"npc_dota_neutral_forest_troll_berserker"`
			NpcDotaNeutralBlackDrake               int `json:"npc_dota_neutral_black_drake"`
			NpcDotaNeutralBlackDragon              int `json:"npc_dota_neutral_black_dragon"`
			NpcDotaHeroEnchantress                 int `json:"npc_dota_hero_enchantress"`
			NpcDotaHeroTemplarAssassin             int `json:"npc_dota_hero_templar_assassin"`
			NpcDotaHeroPuck                        int `json:"npc_dota_hero_puck"`
			NpcDotaNeutralGnollAssassin            int `json:"npc_dota_neutral_gnoll_assassin"`
			NpcDotaNeutralOgreMagi                 int `json:"npc_dota_neutral_ogre_magi"`
			NpcDotaNeutralOgreMauler               int `json:"npc_dota_neutral_ogre_mauler"`
			NpcDotaSentryWards                     int `json:"npc_dota_sentry_wards"`
			NpcDotaTemplarAssassinPsionicTrap      int `json:"npc_dota_templar_assassin_psionic_trap"`
			NpcDotaNeutralKobold                   int `json:"npc_dota_neutral_kobold"`
			NpcDotaGoodguysTower1Mid               int `json:"npc_dota_goodguys_tower1_mid"`
			NpcDotaNeutralForestTrollHighPriest    int `json:"npc_dota_neutral_forest_troll_high_priest"`
			NpcDotaGoodguysTower2Mid               int `json:"npc_dota_goodguys_tower2_mid"`
			NpcDotaGoodguysTower3Mid               int `json:"npc_dota_goodguys_tower3_mid"`
			NpcDotaGoodguysMeleeRaxMid             int `json:"npc_dota_goodguys_melee_rax_mid"`
			NpcDotaGoodguysRangeRaxMid             int `json:"npc_dota_goodguys_range_rax_mid"`
			NpcDotaGoodguysTower3Top               int `json:"npc_dota_goodguys_tower3_top"`
			NpcDotaRoshan                          int `json:"npc_dota_roshan"`
			NpcDotaGoodguysTower3Bot               int `json:"npc_dota_goodguys_tower3_bot"`
			NpcDotaGoodguysMeleeRaxBot             int `json:"npc_dota_goodguys_melee_rax_bot"`
			NpcDotaObserverWards                   int `json:"npc_dota_observer_wards"`
			NpcDotaGoodguysRangeRaxBot             int `json:"npc_dota_goodguys_range_rax_bot"`
			NpcDotaGoodguysTower4                  int `json:"npc_dota_goodguys_tower4"`
		} `json:"damage,omitempty"`
		DamageInflictor8 struct {
			Null                           int `json:"null"`
			TrollWarlordWhirlingAxesRanged int `json:"troll_warlord_whirling_axes_ranged"`
			TrollWarlordWhirlingAxesMelee  int `json:"troll_warlord_whirling_axes_melee"`
			Bfury                          int `json:"bfury"`
		} `json:"damage_inflictor,omitempty"`
		DamageInflictorReceived8 struct {
			Null                     int `json:"null"`
			DeathProphetCarrionSwarm int `json:"death_prophet_carrion_swarm"`
			DeathProphetSpiritSiphon int `json:"death_prophet_spirit_siphon"`
			PrimalBeastTrample       int `json:"primal_beast_trample"`
			PrimalBeastOnslaught     int `json:"primal_beast_onslaught"`
			UrnOfShadows             int `json:"urn_of_shadows"`
			PrimalBeastPulverize     int `json:"primal_beast_pulverize"`
			SpiritVessel             int `json:"spirit_vessel"`
			PuckWaningRift           int `json:"puck_waning_rift"`
			DeathProphetExorcism     int `json:"death_prophet_exorcism"`
			PuckIllusoryOrb          int `json:"puck_illusory_orb"`
			WitchBlade               int `json:"witch_blade"`
			TemplarAssassinMeld      int `json:"templar_assassin_meld"`
			TemplarAssassinTrap      int `json:"templar_assassin_trap"`
			EnchantressImpetus       int `json:"enchantress_impetus"`
			SatyrHellcallerShockwave int `json:"satyr_hellcaller_shockwave"`
			CentaurKhanWarStomp      int `json:"centaur_khan_war_stomp"`
			DragonScale              int `json:"dragon_scale"`
			PrimalBeastRockThrow     int `json:"primal_beast_rock_throw"`
			Stormcrafter             int `json:"stormcrafter"`
			PuckDreamCoil            int `json:"puck_dream_coil"`
			TemplarAssassinPsiBlades int `json:"templar_assassin_psi_blades"`
			OverwhelmingBlink        int `json:"overwhelming_blink"`
		} `json:"damage_inflictor_received,omitempty"`
		DamageTaken8 struct {
			NpcDotaHeroDeathProphet                int `json:"npc_dota_hero_death_prophet"`
			NpcDotaCreepGoodguysMelee              int `json:"npc_dota_creep_goodguys_melee"`
			NpcDotaCreepGoodguysFlagbearer         int `json:"npc_dota_creep_goodguys_flagbearer"`
			NpcDotaCreepGoodguysRanged             int `json:"npc_dota_creep_goodguys_ranged"`
			NpcDotaHeroPrimalBeast                 int `json:"npc_dota_hero_primal_beast"`
			NpcDotaGoodguysSiege                   int `json:"npc_dota_goodguys_siege"`
			NpcDotaNeutralSatyrSoulstealer         int `json:"npc_dota_neutral_satyr_soulstealer"`
			NpcDotaNeutralSatyrTrickster           int `json:"npc_dota_neutral_satyr_trickster"`
			NpcDotaNeutralPolarFurbolgChampion     int `json:"npc_dota_neutral_polar_furbolg_champion"`
			NpcDotaNeutralPolarFurbolgUrsaWarrior  int `json:"npc_dota_neutral_polar_furbolg_ursa_warrior"`
			NpcDotaNeutralDarkTrollWarlord         int `json:"npc_dota_neutral_dark_troll_warlord"`
			NpcDotaNeutralBlackDragon              int `json:"npc_dota_neutral_black_dragon"`
			NpcDotaNeutralDarkTroll                int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaNeutralBigThunderLizard         int `json:"npc_dota_neutral_big_thunder_lizard"`
			NpcDotaNeutralSmallThunderLizard       int `json:"npc_dota_neutral_small_thunder_lizard"`
			NpcDotaNeutralWarpineRaider            int `json:"npc_dota_neutral_warpine_raider"`
			NpcDotaNeutralGiantWolf                int `json:"npc_dota_neutral_giant_wolf"`
			NpcDotaNeutralAlphaWolf                int `json:"npc_dota_neutral_alpha_wolf"`
			NpcDotaNeutralSatyrHellcaller          int `json:"npc_dota_neutral_satyr_hellcaller"`
			NpcDotaNeutralWildkin                  int `json:"npc_dota_neutral_wildkin"`
			NpcDotaNeutralEnragedWildkin           int `json:"npc_dota_neutral_enraged_wildkin"`
			NpcDotaEnragedWildkinTornado           int `json:"npc_dota_enraged_wildkin_tornado"`
			NpcDotaNeutralGhost                    int `json:"npc_dota_neutral_ghost"`
			NpcDotaNeutralFelBeast                 int `json:"npc_dota_neutral_fel_beast"`
			NpcDotaNeutralMudGolem                 int `json:"npc_dota_neutral_mud_golem"`
			NpcDotaNeutralMudGolemSplit            int `json:"npc_dota_neutral_mud_golem_split"`
			NpcDotaHeroPuck                        int `json:"npc_dota_hero_puck"`
			NpcDotaNeutralIceShaman                int `json:"npc_dota_neutral_ice_shaman"`
			NpcDotaNeutralFrostbittenGolem         int `json:"npc_dota_neutral_frostbitten_golem"`
			NpcDotaNeutralKoboldTaskmaster         int `json:"npc_dota_neutral_kobold_taskmaster"`
			NpcDotaNeutralGraniteGolem             int `json:"npc_dota_neutral_granite_golem"`
			NpcDotaNeutralRockGolem                int `json:"npc_dota_neutral_rock_golem"`
			NpcDotaNeutralHarpyStorm               int `json:"npc_dota_neutral_harpy_storm"`
			NpcDotaNeutralHarpyScout               int `json:"npc_dota_neutral_harpy_scout"`
			NpcDotaDarkTrollWarlordSkeletonWarrior int `json:"npc_dota_dark_troll_warlord_skeleton_warrior"`
			NpcDotaNeutralCentaurOutrunner         int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaNeutralCentaurKhan              int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaNeutralForestTrollBerserker     int `json:"npc_dota_neutral_forest_troll_berserker"`
			NpcDotaNeutralBlackDrake               int `json:"npc_dota_neutral_black_drake"`
			NpcDotaHeroTemplarAssassin             int `json:"npc_dota_hero_templar_assassin"`
			NpcDotaHeroEnchantress                 int `json:"npc_dota_hero_enchantress"`
			NpcDotaNeutralGnollAssassin            int `json:"npc_dota_neutral_gnoll_assassin"`
			NpcDotaNeutralOgreMauler               int `json:"npc_dota_neutral_ogre_mauler"`
			NpcDotaNeutralOgreMagi                 int `json:"npc_dota_neutral_ogre_magi"`
			NpcDotaGoodguysTower1Mid               int `json:"npc_dota_goodguys_tower1_mid"`
			NpcDotaGoodguysTower2Mid               int `json:"npc_dota_goodguys_tower2_mid"`
			NpcDotaGoodguysTower3Mid               int `json:"npc_dota_goodguys_tower3_mid"`
			NpcDotaGoodguysTower3Top               int `json:"npc_dota_goodguys_tower3_top"`
			NpcDotaRoshan                          int `json:"npc_dota_roshan"`
			NpcDotaGoodguysTower3Bot               int `json:"npc_dota_goodguys_tower3_bot"`
			NpcDotaGoodguysTower4                  int `json:"npc_dota_goodguys_tower4"`
		} `json:"damage_taken,omitempty"`
		DamageTargets8 struct {
			Null struct {
				NpcDotaHeroPrimalBeast     int `json:"npc_dota_hero_primal_beast"`
				NpcDotaHeroDeathProphet    int `json:"npc_dota_hero_death_prophet"`
				NpcDotaHeroTemplarAssassin int `json:"npc_dota_hero_templar_assassin"`
				NpcDotaHeroPuck            int `json:"npc_dota_hero_puck"`
				NpcDotaHeroEnchantress     int `json:"npc_dota_hero_enchantress"`
			} `json:"null"`
			TrollWarlordWhirlingAxesRanged struct {
				NpcDotaHeroPrimalBeast     int `json:"npc_dota_hero_primal_beast"`
				NpcDotaHeroDeathProphet    int `json:"npc_dota_hero_death_prophet"`
				NpcDotaHeroEnchantress     int `json:"npc_dota_hero_enchantress"`
				NpcDotaHeroTemplarAssassin int `json:"npc_dota_hero_templar_assassin"`
			} `json:"troll_warlord_whirling_axes_ranged"`
			TrollWarlordWhirlingAxesMelee struct {
				NpcDotaHeroPrimalBeast     int `json:"npc_dota_hero_primal_beast"`
				NpcDotaHeroDeathProphet    int `json:"npc_dota_hero_death_prophet"`
				NpcDotaHeroTemplarAssassin int `json:"npc_dota_hero_templar_assassin"`
				NpcDotaHeroEnchantress     int `json:"npc_dota_hero_enchantress"`
				NpcDotaHeroPuck            int `json:"npc_dota_hero_puck"`
			} `json:"troll_warlord_whirling_axes_melee"`
			Bfury struct {
				NpcDotaHeroEnchantress     int `json:"npc_dota_hero_enchantress"`
				NpcDotaHeroTemplarAssassin int `json:"npc_dota_hero_templar_assassin"`
				NpcDotaHeroPuck            int `json:"npc_dota_hero_puck"`
				NpcDotaHeroDeathProphet    int `json:"npc_dota_hero_death_prophet"`
			} `json:"bfury"`
		} `json:"damage_targets,omitempty"`
		GoldReasons7 struct {
			Num0  int `json:"0"`
			Num1  int `json:"1"`
			Num6  int `json:"6"`
			Num11 int `json:"11"`
			Num12 int `json:"12"`
			Num13 int `json:"13"`
			Num14 int `json:"14"`
			Num15 int `json:"15"`
			Num16 int `json:"16"`
			Num17 int `json:"17"`
			Num19 int `json:"19"`
		} `json:"gold_reasons,omitempty"`
		HeroHits8 struct {
			Null                           int `json:"null"`
			TrollWarlordWhirlingAxesRanged int `json:"troll_warlord_whirling_axes_ranged"`
			TrollWarlordWhirlingAxesMelee  int `json:"troll_warlord_whirling_axes_melee"`
			Bfury                          int `json:"bfury"`
		} `json:"hero_hits,omitempty"`
		ItemUses8 struct {
			QuellingBlade int `json:"quelling_blade"`
			Tango         int `json:"tango"`
			MagicStick    int `json:"magic_stick"`
			TangoSingle   int `json:"tango_single"`
			FaerieFire    int `json:"faerie_fire"`
			Tpscroll      int `json:"tpscroll"`
			PowerTreads   int `json:"power_treads"`
			Bfury         int `json:"bfury"`
			BlackKingBar  int `json:"black_king_bar"`
			MagicWand     int `json:"magic_wand"`
			AbyssalBlade  int `json:"abyssal_blade"`
			Blink         int `json:"blink"`
		} `json:"item_uses,omitempty"`
		KillStreaks4 struct {
			Num3 int `json:"3"`
			Num4 int `json:"4"`
			Num5 int `json:"5"`
			Num6 int `json:"6"`
		} `json:"kill_streaks,omitempty"`
		Killed8 struct {
			NpcDotaCreepGoodguysRanged             int `json:"npc_dota_creep_goodguys_ranged"`
			NpcDotaCreepGoodguysMelee              int `json:"npc_dota_creep_goodguys_melee"`
			NpcDotaCreepBadguysMelee               int `json:"npc_dota_creep_badguys_melee"`
			NpcDotaCreepGoodguysFlagbearer         int `json:"npc_dota_creep_goodguys_flagbearer"`
			NpcDotaCreepBadguysRanged              int `json:"npc_dota_creep_badguys_ranged"`
			NpcDotaNeutralSatyrSoulstealer         int `json:"npc_dota_neutral_satyr_soulstealer"`
			NpcDotaNeutralSatyrTrickster           int `json:"npc_dota_neutral_satyr_trickster"`
			NpcDotaNeutralPolarFurbolgChampion     int `json:"npc_dota_neutral_polar_furbolg_champion"`
			NpcDotaNeutralPolarFurbolgUrsaWarrior  int `json:"npc_dota_neutral_polar_furbolg_ursa_warrior"`
			NpcDotaNeutralWarpineRaider            int `json:"npc_dota_neutral_warpine_raider"`
			NpcDotaDarkTrollWarlordSkeletonWarrior int `json:"npc_dota_dark_troll_warlord_skeleton_warrior"`
			NpcDotaNeutralDarkTrollWarlord         int `json:"npc_dota_neutral_dark_troll_warlord"`
			NpcDotaNeutralDarkTroll                int `json:"npc_dota_neutral_dark_troll"`
			NpcDotaNeutralGiantWolf                int `json:"npc_dota_neutral_giant_wolf"`
			NpcDotaNeutralAlphaWolf                int `json:"npc_dota_neutral_alpha_wolf"`
			NpcDotaNeutralSatyrHellcaller          int `json:"npc_dota_neutral_satyr_hellcaller"`
			NpcDotaNeutralEnragedWildkin           int `json:"npc_dota_neutral_enraged_wildkin"`
			NpcDotaNeutralWildkin                  int `json:"npc_dota_neutral_wildkin"`
			NpcDotaNeutralGhost                    int `json:"npc_dota_neutral_ghost"`
			NpcDotaNeutralFelBeast                 int `json:"npc_dota_neutral_fel_beast"`
			NpcDotaNeutralMudGolem                 int `json:"npc_dota_neutral_mud_golem"`
			NpcDotaNeutralMudGolemSplit            int `json:"npc_dota_neutral_mud_golem_split"`
			NpcDotaNeutralIceShaman                int `json:"npc_dota_neutral_ice_shaman"`
			NpcDotaNeutralFrostbittenGolem         int `json:"npc_dota_neutral_frostbitten_golem"`
			NpcDotaGoodguysSiege                   int `json:"npc_dota_goodguys_siege"`
			NpcDotaNeutralKoboldTunneler           int `json:"npc_dota_neutral_kobold_tunneler"`
			NpcDotaNeutralKoboldTaskmaster         int `json:"npc_dota_neutral_kobold_taskmaster"`
			NpcDotaNeutralGraniteGolem             int `json:"npc_dota_neutral_granite_golem"`
			NpcDotaNeutralRockGolem                int `json:"npc_dota_neutral_rock_golem"`
			NpcDotaCreepBadguysFlagbearer          int `json:"npc_dota_creep_badguys_flagbearer"`
			NpcDotaNeutralHarpyStorm               int `json:"npc_dota_neutral_harpy_storm"`
			NpcDotaNeutralHarpyScout               int `json:"npc_dota_neutral_harpy_scout"`
			NpcDotaNeutralBigThunderLizard         int `json:"npc_dota_neutral_big_thunder_lizard"`
			NpcDotaNeutralSmallThunderLizard       int `json:"npc_dota_neutral_small_thunder_lizard"`
			NpcDotaNeutralCentaurOutrunner         int `json:"npc_dota_neutral_centaur_outrunner"`
			NpcDotaNeutralCentaurKhan              int `json:"npc_dota_neutral_centaur_khan"`
			NpcDotaNeutralForestTrollBerserker     int `json:"npc_dota_neutral_forest_troll_berserker"`
			NpcDotaNeutralBlackDrake               int `json:"npc_dota_neutral_black_drake"`
			NpcDotaNeutralBlackDragon              int `json:"npc_dota_neutral_black_dragon"`
			NpcDotaHeroEnchantress                 int `json:"npc_dota_hero_enchantress"`
			NpcDotaHeroDeathProphet                int `json:"npc_dota_hero_death_prophet"`
			NpcDotaNeutralGnollAssassin            int `json:"npc_dota_neutral_gnoll_assassin"`
			NpcDotaNeutralOgreMauler               int `json:"npc_dota_neutral_ogre_mauler"`
			NpcDotaNeutralOgreMagi                 int `json:"npc_dota_neutral_ogre_magi"`
			NpcDotaSentryWards                     int `json:"npc_dota_sentry_wards"`
			NpcDotaTemplarAssassinPsionicTrap      int `json:"npc_dota_templar_assassin_psionic_trap"`
			NpcDotaNeutralKobold                   int `json:"npc_dota_neutral_kobold"`
			NpcDotaGoodguysTower1Mid               int `json:"npc_dota_goodguys_tower1_mid"`
			NpcDotaNeutralForestTrollHighPriest    int `json:"npc_dota_neutral_forest_troll_high_priest"`
			NpcDotaHeroTemplarAssassin             int `json:"npc_dota_hero_templar_assassin"`
			NpcDotaGoodguysTower2Mid               int `json:"npc_dota_goodguys_tower2_mid"`
			NpcDotaGoodguysRangeRaxMid             int `json:"npc_dota_goodguys_range_rax_mid"`
			NpcDotaRoshan                          int `json:"npc_dota_roshan"`
			NpcDotaObserverWards                   int `json:"npc_dota_observer_wards"`
			NpcDotaGoodguysRangeRaxBot             int `json:"npc_dota_goodguys_range_rax_bot"`
			NpcDotaHeroPrimalBeast                 int `json:"npc_dota_hero_primal_beast"`
			NpcDotaGoodguysTower4                  int `json:"npc_dota_goodguys_tower4"`
		} `json:"killed,omitempty"`
		KilledBy7 struct {
			NpcDotaHeroDeathProphet int `json:"npc_dota_hero_death_prophet"`
			NpcDotaHeroPuck         int `json:"npc_dota_hero_puck"`
		} `json:"killed_by,omitempty"`
		LanePos8 struct {
			Num76 struct {
				Num146 int `json:"146"`
				Num148 int `json:"148"`
				Num150 int `json:"150"`
				Num156 int `json:"156"`
				Num158 int `json:"158"`
			} `json:"76"`
			Num78 struct {
				Num150 int `json:"150"`
				Num154 int `json:"154"`
				Num156 int `json:"156"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
			} `json:"78"`
			Num80 struct {
				Num154 int `json:"154"`
				Num156 int `json:"156"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
			} `json:"80"`
			Num82 struct {
				Num150 int `json:"150"`
				Num152 int `json:"152"`
				Num160 int `json:"160"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
				Num174 int `json:"174"`
			} `json:"82"`
			Num84 struct {
				Num166 int `json:"166"`
				Num168 int `json:"168"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
			} `json:"84"`
			Num86 struct {
				Num168 int `json:"168"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
			} `json:"86"`
			Num88 struct {
				Num170 int `json:"170"`
				Num172 int `json:"172"`
				Num174 int `json:"174"`
				Num176 int `json:"176"`
			} `json:"88"`
			Num90 struct {
				Num170 int `json:"170"`
				Num172 int `json:"172"`
				Num174 int `json:"174"`
				Num176 int `json:"176"`
			} `json:"90"`
			Num92 struct {
				Num170 int `json:"170"`
				Num172 int `json:"172"`
				Num174 int `json:"174"`
				Num176 int `json:"176"`
			} `json:"92"`
			Num94 struct {
				Num166 int `json:"166"`
				Num168 int `json:"168"`
				Num174 int `json:"174"`
				Num176 int `json:"176"`
			} `json:"94"`
			Num96 struct {
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num172 int `json:"172"`
				Num174 int `json:"174"`
				Num176 int `json:"176"`
			} `json:"96"`
			Num98 struct {
				Num162 int `json:"162"`
				Num172 int `json:"172"`
				Num174 int `json:"174"`
			} `json:"98"`
			Num100 struct {
				Num160 int `json:"160"`
				Num172 int `json:"172"`
				Num174 int `json:"174"`
				Num176 int `json:"176"`
			} `json:"100"`
			Num102 struct {
				Num160 int `json:"160"`
				Num170 int `json:"170"`
				Num176 int `json:"176"`
			} `json:"102"`
			Num104 struct {
				Num172 int `json:"172"`
				Num176 int `json:"176"`
			} `json:"104"`
			Num106 struct {
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num174 int `json:"174"`
				Num176 int `json:"176"`
			} `json:"106"`
			Num108 struct {
				Num158 int `json:"158"`
				Num174 int `json:"174"`
			} `json:"108"`
			Num110 struct {
				Num158 int `json:"158"`
				Num174 int `json:"174"`
				Num176 int `json:"176"`
			} `json:"110"`
			Num112 struct {
				Num158 int `json:"158"`
				Num160 int `json:"160"`
			} `json:"112"`
			Num114 struct {
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num174 int `json:"174"`
			} `json:"114"`
			Num116 struct {
				Num142 int `json:"142"`
				Num144 int `json:"144"`
				Num146 int `json:"146"`
				Num156 int `json:"156"`
				Num160 int `json:"160"`
				Num174 int `json:"174"`
			} `json:"116"`
			Num118 struct {
				Num146 int `json:"146"`
				Num148 int `json:"148"`
				Num158 int `json:"158"`
				Num160 int `json:"160"`
				Num172 int `json:"172"`
			} `json:"118"`
			Num120 struct {
				Num148 int `json:"148"`
				Num150 int `json:"150"`
				Num152 int `json:"152"`
				Num158 int `json:"158"`
				Num172 int `json:"172"`
			} `json:"120"`
			Num122 struct {
				Num150 int `json:"150"`
				Num152 int `json:"152"`
				Num154 int `json:"154"`
				Num156 int `json:"156"`
				Num158 int `json:"158"`
				Num170 int `json:"170"`
				Num172 int `json:"172"`
			} `json:"122"`
			Num124 struct {
				Num154 int `json:"154"`
				Num156 int `json:"156"`
				Num158 int `json:"158"`
			} `json:"124"`
			Num126 struct {
				Num154 int `json:"154"`
				Num160 int `json:"160"`
				Num172 int `json:"172"`
			} `json:"126"`
			Num128 struct {
				Num154 int `json:"154"`
				Num156 int `json:"156"`
				Num160 int `json:"160"`
				Num172 int `json:"172"`
			} `json:"128"`
			Num130 struct {
				Num152 int `json:"152"`
				Num158 int `json:"158"`
				Num162 int `json:"162"`
				Num164 int `json:"164"`
				Num170 int `json:"170"`
			} `json:"130"`
			Num132 struct {
				Num150 int `json:"150"`
				Num166 int `json:"166"`
				Num168 int `json:"168"`
			} `json:"132"`
			Num134 struct {
				Num132 int `json:"132"`
				Num150 int `json:"150"`
				Num168 int `json:"168"`
			} `json:"134"`
			Num136 struct {
				Num124 int `json:"124"`
				Num132 int `json:"132"`
				Num150 int `json:"150"`
				Num152 int `json:"152"`
				Num168 int `json:"168"`
			} `json:"136"`
			Num138 struct {
				Num124 int `json:"124"`
				Num126 int `json:"126"`
				Num128 int `json:"128"`
				Num130 int `json:"130"`
				Num134 int `json:"134"`
				Num136 int `json:"136"`
				Num148 int `json:"148"`
				Num150 int `json:"150"`
				Num170 int `json:"170"`
			} `json:"138"`
			Num140 struct {
				Num126 int `json:"126"`
				Num128 int `json:"128"`
				Num136 int `json:"136"`
				Num146 int `json:"146"`
				Num170 int `json:"170"`
			} `json:"140"`
			Num142 struct {
				Num138 int `json:"138"`
				Num144 int `json:"144"`
				Num172 int `json:"172"`
			} `json:"142"`
			Num144 struct {
				Num128 int `json:"128"`
				Num140 int `json:"140"`
				Num172 int `json:"172"`
			} `json:"144"`
			Num146 struct {
				Num128 int `json:"128"`
				Num130 int `json:"130"`
				Num172 int `json:"172"`
			} `json:"146"`
			Num148 struct {
				Num126 int `json:"126"`
				Num128 int `json:"128"`
				Num174 int `json:"174"`
			} `json:"148"`
			Num150 struct {
				Num124 int `json:"124"`
				Num126 int `json:"126"`
				Num174 int `json:"174"`
			} `json:"150"`
			Num152 struct {
				Num120 int `json:"120"`
				Num122 int `json:"122"`
				Num124 int `json:"124"`
				Num174 int `json:"174"`
			} `json:"152"`
			Num154 struct {
				Num84  int `json:"84"`
				Num120 int `json:"120"`
				Num174 int `json:"174"`
			} `json:"154"`
			Num156 struct {
				Num80  int `json:"80"`
				Num84  int `json:"84"`
				Num86  int `json:"86"`
				Num88  int `json:"88"`
				Num90  int `json:"90"`
				Num92  int `json:"92"`
				Num118 int `json:"118"`
				Num172 int `json:"172"`
			} `json:"156"`
			Num158 struct {
				Num80  int `json:"80"`
				Num96  int `json:"96"`
				Num100 int `json:"100"`
				Num102 int `json:"102"`
				Num104 int `json:"104"`
				Num116 int `json:"116"`
				Num172 int `json:"172"`
			} `json:"158"`
			Num160 struct {
				Num82  int `json:"82"`
				Num84  int `json:"84"`
				Num86  int `json:"86"`
				Num96  int `json:"96"`
				Num100 int `json:"100"`
				Num106 int `json:"106"`
				Num108 int `json:"108"`
				Num110 int `json:"110"`
				Num114 int `json:"114"`
				Num116 int `json:"116"`
			} `json:"160"`
			Num162 struct {
				Num80  int `json:"80"`
				Num98  int `json:"98"`
				Num100 int `json:"100"`
				Num110 int `json:"110"`
				Num112 int `json:"112"`
				Num114 int `json:"114"`
				Num174 int `json:"174"`
			} `json:"162"`
			Num164 struct {
				Num174 int `json:"174"`
			} `json:"164"`
			Num166 struct {
				Num174 int `json:"174"`
			} `json:"166"`
			Num168 struct {
				Num174 int `json:"174"`
			} `json:"168"`
			Num170 struct {
				Num174 int `json:"174"`
			} `json:"170"`
			Num172 struct {
				Num174 int `json:"174"`
			} `json:"172"`
			Num174 struct {
				Num174 int `json:"174"`
			} `json:"174"`
			Num178 struct {
				Num174 int `json:"174"`
			} `json:"178"`
			Num180 struct {
				Num174 int `json:"174"`
				Num178 int `json:"178"`
			} `json:"180"`
			Num182 struct {
				Num174 int `json:"174"`
			} `json:"182"`
			Num184 struct {
				Num174 int `json:"174"`
			} `json:"184"`
		} `json:"lane_pos,omitempty"`
		MultiKills2 struct {
			Num2 int `json:"2"`
			Num3 int `json:"3"`
			Num4 int `json:"4"`
		} `json:"multi_kills,omitempty"`
		Obs8 struct {
		} `json:"obs,omitempty"`
		Purchase8 struct {
			Branches            int `json:"branches"`
			MagicStick          int `json:"magic_stick"`
			Tango               int `json:"tango"`
			FaerieFire          int `json:"faerie_fire"`
			QuellingBlade       int `json:"quelling_blade"`
			Boots               int `json:"boots"`
			BootsOfElves        int `json:"boots_of_elves"`
			Gloves              int `json:"gloves"`
			PowerTreads         int `json:"power_treads"`
			RingOfHealth        int `json:"ring_of_health"`
			VoidStone           int `json:"void_stone"`
			Pers                int `json:"pers"`
			Broadsword          int `json:"broadsword"`
			Claymore            int `json:"claymore"`
			WindLace            int `json:"wind_lace"`
			Tpscroll            int `json:"tpscroll"`
			Bfury               int `json:"bfury"`
			RecipeMagicWand     int `json:"recipe_magic_wand"`
			MagicWand           int `json:"magic_wand"`
			MithrilHammer       int `json:"mithril_hammer"`
			OgreAxe             int `json:"ogre_axe"`
			RecipeBlackKingBar  int `json:"recipe_black_king_bar"`
			BlackKingBar        int `json:"black_king_bar"`
			BeltOfStrength      int `json:"belt_of_strength"`
			RecipeBasher        int `json:"recipe_basher"`
			Basher              int `json:"basher"`
			UltimateOrb         int `json:"ultimate_orb"`
			Skadi               int `json:"skadi"`
			PointBooster        int `json:"point_booster"`
			DemonEdge           int `json:"demon_edge"`
			Javelin             int `json:"javelin"`
			BlitzKnuckles       int `json:"blitz_knuckles"`
			RecipeMonkeyKingBar int `json:"recipe_monkey_king_bar"`
			MonkeyKingBar       int `json:"monkey_king_bar"`
			RecipeAbyssalBlade  int `json:"recipe_abyssal_blade"`
			Vanguard            int `json:"vanguard"`
			VitalityBooster     int `json:"vitality_booster"`
			AbyssalBlade        int `json:"abyssal_blade"`
			Blink               int `json:"blink"`
			RecipeSwiftBlink    int `json:"recipe_swift_blink"`
			Eagle               int `json:"eagle"`
		} `json:"purchase,omitempty"`
		Sen8 struct {
		} `json:"sen,omitempty"`
		XpReasons7 struct {
			Num0 int `json:"0"`
			Num1 int `json:"1"`
			Num2 int `json:"2"`
			Num3 int `json:"3"`
			Num5 int `json:"5"`
		} `json:"xp_reasons,omitempty"`
		PurchaseTime8 struct {
			Branches        int `json:"branches"`
			MagicStick      int `json:"magic_stick"`
			Tango           int `json:"tango"`
			FaerieFire      int `json:"faerie_fire"`
			QuellingBlade   int `json:"quelling_blade"`
			Boots           int `json:"boots"`
			BootsOfElves    int `json:"boots_of_elves"`
			Gloves          int `json:"gloves"`
			PowerTreads     int `json:"power_treads"`
			RingOfHealth    int `json:"ring_of_health"`
			VoidStone       int `json:"void_stone"`
			Pers            int `json:"pers"`
			Broadsword      int `json:"broadsword"`
			Claymore        int `json:"claymore"`
			WindLace        int `json:"wind_lace"`
			Tpscroll        int `json:"tpscroll"`
			Bfury           int `json:"bfury"`
			MagicWand       int `json:"magic_wand"`
			MithrilHammer   int `json:"mithril_hammer"`
			OgreAxe         int `json:"ogre_axe"`
			BlackKingBar    int `json:"black_king_bar"`
			BeltOfStrength  int `json:"belt_of_strength"`
			Basher          int `json:"basher"`
			UltimateOrb     int `json:"ultimate_orb"`
			Skadi           int `json:"skadi"`
			PointBooster    int `json:"point_booster"`
			DemonEdge       int `json:"demon_edge"`
			Javelin         int `json:"javelin"`
			BlitzKnuckles   int `json:"blitz_knuckles"`
			MonkeyKingBar   int `json:"monkey_king_bar"`
			Vanguard        int `json:"vanguard"`
			VitalityBooster int `json:"vitality_booster"`
			AbyssalBlade    int `json:"abyssal_blade"`
			Blink           int `json:"blink"`
			Eagle           int `json:"eagle"`
		} `json:"purchase_time,omitempty"`
		FirstPurchaseTime8 struct {
			Branches        int `json:"branches"`
			MagicStick      int `json:"magic_stick"`
			Tango           int `json:"tango"`
			FaerieFire      int `json:"faerie_fire"`
			QuellingBlade   int `json:"quelling_blade"`
			Boots           int `json:"boots"`
			BootsOfElves    int `json:"boots_of_elves"`
			Gloves          int `json:"gloves"`
			PowerTreads     int `json:"power_treads"`
			RingOfHealth    int `json:"ring_of_health"`
			VoidStone       int `json:"void_stone"`
			Pers            int `json:"pers"`
			Broadsword      int `json:"broadsword"`
			Claymore        int `json:"claymore"`
			WindLace        int `json:"wind_lace"`
			Tpscroll        int `json:"tpscroll"`
			Bfury           int `json:"bfury"`
			MagicWand       int `json:"magic_wand"`
			MithrilHammer   int `json:"mithril_hammer"`
			OgreAxe         int `json:"ogre_axe"`
			BlackKingBar    int `json:"black_king_bar"`
			BeltOfStrength  int `json:"belt_of_strength"`
			Basher          int `json:"basher"`
			UltimateOrb     int `json:"ultimate_orb"`
			Skadi           int `json:"skadi"`
			PointBooster    int `json:"point_booster"`
			DemonEdge       int `json:"demon_edge"`
			Javelin         int `json:"javelin"`
			BlitzKnuckles   int `json:"blitz_knuckles"`
			MonkeyKingBar   int `json:"monkey_king_bar"`
			Vanguard        int `json:"vanguard"`
			VitalityBooster int `json:"vitality_booster"`
			AbyssalBlade    int `json:"abyssal_blade"`
			Blink           int `json:"blink"`
			Eagle           int `json:"eagle"`
		} `json:"first_purchase_time,omitempty"`
		ItemWin8 struct {
			Branches        int `json:"branches"`
			MagicStick      int `json:"magic_stick"`
			Tango           int `json:"tango"`
			FaerieFire      int `json:"faerie_fire"`
			QuellingBlade   int `json:"quelling_blade"`
			Boots           int `json:"boots"`
			BootsOfElves    int `json:"boots_of_elves"`
			Gloves          int `json:"gloves"`
			PowerTreads     int `json:"power_treads"`
			RingOfHealth    int `json:"ring_of_health"`
			VoidStone       int `json:"void_stone"`
			Pers            int `json:"pers"`
			Broadsword      int `json:"broadsword"`
			Claymore        int `json:"claymore"`
			WindLace        int `json:"wind_lace"`
			Tpscroll        int `json:"tpscroll"`
			Bfury           int `json:"bfury"`
			MagicWand       int `json:"magic_wand"`
			MithrilHammer   int `json:"mithril_hammer"`
			OgreAxe         int `json:"ogre_axe"`
			BlackKingBar    int `json:"black_king_bar"`
			BeltOfStrength  int `json:"belt_of_strength"`
			Basher          int `json:"basher"`
			UltimateOrb     int `json:"ultimate_orb"`
			Skadi           int `json:"skadi"`
			PointBooster    int `json:"point_booster"`
			DemonEdge       int `json:"demon_edge"`
			Javelin         int `json:"javelin"`
			BlitzKnuckles   int `json:"blitz_knuckles"`
			MonkeyKingBar   int `json:"monkey_king_bar"`
			Vanguard        int `json:"vanguard"`
			VitalityBooster int `json:"vitality_booster"`
			AbyssalBlade    int `json:"abyssal_blade"`
			Blink           int `json:"blink"`
			Eagle           int `json:"eagle"`
		} `json:"item_win,omitempty"`
		ItemUsage8 struct {
			Branches        int `json:"branches"`
			MagicStick      int `json:"magic_stick"`
			Tango           int `json:"tango"`
			FaerieFire      int `json:"faerie_fire"`
			QuellingBlade   int `json:"quelling_blade"`
			Boots           int `json:"boots"`
			BootsOfElves    int `json:"boots_of_elves"`
			Gloves          int `json:"gloves"`
			PowerTreads     int `json:"power_treads"`
			RingOfHealth    int `json:"ring_of_health"`
			VoidStone       int `json:"void_stone"`
			Pers            int `json:"pers"`
			Broadsword      int `json:"broadsword"`
			Claymore        int `json:"claymore"`
			WindLace        int `json:"wind_lace"`
			Tpscroll        int `json:"tpscroll"`
			Bfury           int `json:"bfury"`
			MagicWand       int `json:"magic_wand"`
			MithrilHammer   int `json:"mithril_hammer"`
			OgreAxe         int `json:"ogre_axe"`
			BlackKingBar    int `json:"black_king_bar"`
			BeltOfStrength  int `json:"belt_of_strength"`
			Basher          int `json:"basher"`
			UltimateOrb     int `json:"ultimate_orb"`
			Skadi           int `json:"skadi"`
			PointBooster    int `json:"point_booster"`
			DemonEdge       int `json:"demon_edge"`
			Javelin         int `json:"javelin"`
			BlitzKnuckles   int `json:"blitz_knuckles"`
			MonkeyKingBar   int `json:"monkey_king_bar"`
			Vanguard        int `json:"vanguard"`
			VitalityBooster int `json:"vitality_booster"`
			AbyssalBlade    int `json:"abyssal_blade"`
			Blink           int `json:"blink"`
			Eagle           int `json:"eagle"`
		} `json:"item_usage,omitempty"`
	} `json:"players"`
	Patch         int `json:"patch"`
	Region        int `json:"region"`
	AllWordCounts struct {
		Sec            int `json:"sec"`
		Gl             int `json:"gl"`
		Hf             int `json:"hf"`
		Hfhf           int `json:"hfhf"`
		Rdy            int `json:"rdy"`
		NAMING_FAILED  int `json:"не"`
		NAMING_FAILED0 int `json:"зря"`
		NAMING_FAILED1 int `json:"игр"`
		NAMING_FAILED2 int `json:"на"`
		NAMING_FAILED3 int `json:"бисте"`
		NAMING_FAILED4 int `json:"семен"`
		G              int `json:"g"`
		Gg             int `json:"gg"`
		Glgl           int `json:"glgl"`
		Ggwp           int `json:"ggwp"`
		Ty             int `json:"ty"`
		Next           int `json:"next"`
		Thanks         int `json:"thanks"`
	} `json:"all_word_counts"`
	MyWordCounts struct {
	} `json:"my_word_counts"`
	Comeback  int    `json:"comeback"`
	Stomp     int    `json:"stomp"`
	ReplayURL string `json:"replay_url"`
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
