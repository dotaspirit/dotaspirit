package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"

	"gopkg.in/gographics/imagick.v3/imagick"
)

func prettifyBigValues(value int) string {
	if value < 1000 {
		return strconv.Itoa(value)
	} else {
		result := fmt.Sprintf("%.1fk", float64(value)/1000)
		return result
	}
}

func getPlayerRole(pos, role int) string {
	roleNames := [5]string{"Unknown", "Safe", "Mid", "Offlane", "Jungle"}
	return roleNames[role]
}

func getPlayerName(accountID int, playersData oDotaPlayersData) string {
	for _, player := range playersData {
		if player.AccountID == accountID {
			if player.Name != "" {
				return player.Name
			}
			return player.Personaname
		}
	}
	return "Player"
}

func getLeagueName(leagueID int, leaguesData oDotaLeaguesData) string {
	for _, league := range leaguesData {
		if league.Leagueid == leagueID {
			return league.Name
		}
	}
	return "League"
}

func guessFontSize(fontName string, maxFontSize, minFontSize float64, maxHeight, maxWidth float64, text string) float64 {
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	dw := imagick.NewDrawingWand()
	defer dw.Destroy()
	pw := imagick.NewPixelWand()
	defer pw.Destroy()

	pw.SetColor("none")
	mw.NewImage(1020, 1020, pw)

	fontSize := maxFontSize

	for fontSize >= minFontSize {
		dw.SetFont(fontName)
		dw.SetFontSize(maxFontSize)
		textMetrics := mw.QueryFontMetrics(dw, text)
		if textMetrics.TextHeight > maxHeight || textMetrics.TextWidth > maxWidth {
			fontSize = fontSize - 1
		} else {
			return fontSize
		}
	}

	return fontSize
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

func makeMatchImage(matchData oDotaMatchData, isFull bool) {
	var cConfig colorConfig
	err := loadConfig("config/colors.json", &cConfig)
	if err != nil {
		log.Println(err.Error())
	}

	var playersData oDotaPlayersData
	var leaguesData oDotaLeaguesData

	if isFull != true {
		playersData = getPlayersData()
		leaguesData = getLeaguesData()
	}

	var sMatchID = strconv.FormatInt(matchData.MatchID, 10)

	savePath := fmt.Sprintf("tmp/%d.png", matchData.MatchID)

	matchDuration := matchData.Duration

	hours, minutes, seconds := secondsToTime(matchDuration)

	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	dw := imagick.NewDrawingWand()
	defer dw.Destroy()
	pw := imagick.NewPixelWand()
	defer pw.Destroy()

	dw.SetFont("HypatiaSansPro-Regular")

	pw.SetColor(cConfig.Bg)
	mw.NewImage(1020, 1020, pw)

	// GRAPH

	if isFull {
		graphmw := imagick.NewMagickWand()
		defer graphmw.Destroy()
		graphdw := imagick.NewDrawingWand()
		defer graphdw.Destroy()
		graphpw := imagick.NewPixelWand()
		defer graphpw.Destroy()
		graphpw.SetColor("none")
		graphmw.NewImage(970, 114, graphpw)
		graphpw.SetColor(cConfig.Text + "77")

		graphdw.SetStrokeColor(graphpw)
		graphdw.SetStrokeWidth(2)
		graphdw.Line(0, 56, 970, 56)

		var scalex = 970 / float64(matchData.Duration)

		graphdw.SetStrokeWidth(7)

		for _, tf := range matchData.Teamfights {
			radiantDeaths := 0
			direDeaths := 0
			for i, player := range tf.Players {
				if i < 5 {
					direDeaths += player.Deaths
				} else {
					radiantDeaths += player.Deaths
				}
			}
			if radiantDeaths > direDeaths {
				graphpw.SetColor(cConfig.TextTFRadiant + "99")
			} else if direDeaths > radiantDeaths {
				graphpw.SetColor(cConfig.TextTFDire + "99")
			} else {
				graphpw.SetColor(cConfig.TextTF + "99")
			}
			graphdw.SetStrokeColor(graphpw)
			graphdw.Line(float64(tf.Start)*scalex, 56, float64(tf.End)*scalex, 56)
		}

		for _, obj := range matchData.Objectives {
			if obj.Type == "CHAT_MESSAGE_FIRSTBLOOD" {
				graphdw.SetStrokeWidth(15)
				if obj.PlayerSlot < 5 {
					graphpw.SetColor(cConfig.TextTFDire)
				} else {
					graphpw.SetColor(cConfig.TextTFRadiant)
				}
				graphdw.SetStrokeColor(graphpw)
				graphdw.Line(float64(obj.Time)*scalex-1, 56, float64(obj.Time)*scalex+1, 56)
			}
			if obj.Type == "CHAT_MESSAGE_ROSHAN_KILL" {
				graphdw.SetStrokeWidth(15)
				if obj.Team == 3 {
					graphpw.SetColor(cConfig.TextTFDire)
				} else {
					graphpw.SetColor(cConfig.TextTFRadiant)
				}
				graphdw.SetStrokeColor(graphpw)
				graphdw.Line(float64(obj.Time)*scalex-1, 56, float64(obj.Time)*scalex+1, 56)
			}
		}

		scalex = 970 / float64(matchDuration/60)

		xpAdv := matchData.RadiantXpAdv
		maxXpAdv := xpAdv[0]
		for i := 0; i <= matchDuration/60; i++ {
			posAdv := math.Abs(float64(xpAdv[i]))
			if posAdv > float64(maxXpAdv) {
				maxXpAdv = int(posAdv)
			}
		}

		goldAdv := matchData.RadiantGoldAdv
		maxGoldAdv := goldAdv[0]
		for i := 0; i <= matchDuration/60; i++ {
			posAdv := math.Abs(float64(goldAdv[i]))
			if posAdv > float64(maxGoldAdv) {
				maxGoldAdv = int(posAdv)
			}
		}

		var maxAdv int

		if maxXpAdv > maxGoldAdv {
			maxAdv = maxXpAdv
		} else {
			maxAdv = maxGoldAdv
		}

		scaley := (112 / 2) / math.Abs(float64(maxAdv))

		var sx float64
		var sy float64 = 56

		graphpw.SetColor(cConfig.TextXP)
		graphdw.SetStrokeWidth(2)
		graphdw.SetStrokeColor(graphpw)
		for i, adv := range xpAdv {
			nextx := float64(i) * scalex
			nexty := 56 + float64(adv)*scaley
			graphdw.Line(sx, sy, nextx, nexty)
			sx = nextx
			sy = nexty
		}

		sx = 0
		sy = 56

		graphpw.SetColor(cConfig.TextGold)
		graphdw.SetStrokeColor(graphpw)
		for i, adv := range goldAdv {
			nextx := float64(i) * scalex
			nexty := 56 + float64(adv)*scaley
			graphdw.Line(sx, sy, nextx, nexty)
			sx = nextx
			sy = nexty
		}

		graphmw.DrawImage(graphdw)
		mw.CompositeImage(graphmw, imagick.COMPOSITE_OP_OVER, true, 25, 453)
		// graphmw.WriteImage("./tmp/graph.png")
	}

	// GRAPH

	// TEXT

	pw.SetColor(cConfig.Text)
	dw.SetFontSize(cConfig.TextSize28)
	dw.SetFillColor(pw)

	dw.SetTextAlignment(imagick.ALIGN_LEFT)
	dw.Annotation(25, 63, fmt.Sprintf("%d:%02d:%02d", hours, minutes, seconds))
	if isFull == true {
		dw.Annotation(25, 986, matchData.League.Name)
	} else {
		dw.Annotation(25, 986, getLeagueName(matchData.Leagueid, leaguesData))
	}
	dw.SetTextAlignment(imagick.ALIGN_RIGHT)
	dw.Annotation(995, 63, sMatchID)
	dw.Annotation(995, 986, "#rsltdtk")

	dw.SetFontSize(cConfig.TextSize48)
	dw.SetTextAlignment(imagick.ALIGN_CENTER)
	dw.Annotation(510, 484, strconv.Itoa(matchData.DireScore))
	dw.Annotation(510, 584, strconv.Itoa(matchData.RadiantScore))

	radiantName := matchData.RadiantTeam.Name
	direName := matchData.DireTeam.Name

	if radiantName == "" {
		radiantName = matchData.RadiantName
	}
	if direName == "" {
		direName = matchData.DireName
	}

	if radiantName == "" {
		radiantName = "Radiant"
	}
	if direName == "" {
		direName = "Dire"
	}

	dw.SetFontSize(cConfig.TextSize28)
	pw.SetColor(cConfig.TextRadiant)
	dw.SetFillColor(pw)
	dw.SetTextAlignment(imagick.ALIGN_LEFT)
	if matchData.RadiantWin == true {
		dw.Annotation(64, 569, radiantName)
	} else {
		dw.Annotation(25, 569, radiantName)
	}
	pw.SetColor(cConfig.TextDire)
	dw.SetFillColor(pw)
	if matchData.RadiantWin == false {
		dw.Annotation(64, 469, direName)
	} else {
		dw.Annotation(25, 469, direName)
	}

	// TEXT

	// WINNER IMAGE

	winnerImage := imagick.NewMagickWand()
	defer winnerImage.Destroy()
	winnerImage.ReadImage("assets/images/winner.png")
	if matchData.RadiantWin == true {
		mw.CompositeImage(winnerImage, imagick.COMPOSITE_OP_OVER, true, 25, 541)
	} else {
		mw.CompositeImage(winnerImage, imagick.COMPOSITE_OP_OVER, true, 25, 441)
	}

	// WINNER IMAGE

	// PLAYERS

	direPlayersmw := imagick.NewMagickWand()
	defer direPlayersmw.Destroy()
	direPlayerspw := imagick.NewPixelWand()
	defer direPlayerspw.Destroy()
	radiantPlayersmw := imagick.NewMagickWand()
	defer radiantPlayersmw.Destroy()
	radiantPlayerspw := imagick.NewPixelWand()
	defer radiantPlayerspw.Destroy()

	direPlayerspw.SetColor("none")
	direPlayersmw.NewImage(970, 257, direPlayerspw)
	radiantPlayerspw.SetColor("none")
	radiantPlayersmw.NewImage(970, 257, radiantPlayerspw)
	for i := 0; i < 10; i++ {
		playermw := imagick.NewMagickWand()
		defer playermw.Destroy()
		playerdw := imagick.NewDrawingWand()
		defer playerdw.Destroy()
		playerpw := imagick.NewPixelWand()
		defer playerpw.Destroy()
		var gold = prettifyBigValues(matchData.Players[i].TotalGold)
		if gold == "0" {
			goldAPI := matchData.Players[i].Gold + matchData.Players[i].GoldSpent
			gold = prettifyBigValues(goldAPI)
		}
		var damage = prettifyBigValues(matchData.Players[i].HeroDamage)
		var healing = prettifyBigValues(matchData.Players[i].HeroHealing)
		var level = strconv.Itoa(matchData.Players[i].Level)
		if i < 5 {
			playerpw.SetColor(cConfig.BackRadiant)
		} else {
			playerpw.SetColor(cConfig.BackDire)
		}

		playermw.NewImage(190, 257, playerpw)

		heroImage := imagick.NewMagickWand()
		defer heroImage.Destroy()
		heroImage.ReadImage(fmt.Sprintf("assets/heroes/%d.png", matchData.Players[i].HeroID))
		heroImage.ResizeImage(190, 107, imagick.FILTER_CUBIC)
		playermw.CompositeImage(heroImage, imagick.COMPOSITE_OP_OVER, true, 0, 50)

		playerpw.SetColor(cConfig.PlayerColor[i])
		playerdw.SetFillColor(playerpw)
		playerdw.Rectangle(0, 157, 190, 162)

		kills := matchData.Players[i].Kills
		deaths := matchData.Players[i].Deaths
		assists := matchData.Players[i].Assists

		scalex := 190

		if kills+deaths+assists != 0 {
			scalex = 190 / (kills + deaths + assists)
		}

		playerpw.SetColor(cConfig.KDAColor[0])
		playerdw.SetFillColor(playerpw)
		playerdw.Rectangle(0, 252, float64(kills*scalex), 257)

		playerpw.SetColor(cConfig.KDAColor[1])
		playerdw.SetFillColor(playerpw)
		playerdw.Rectangle(float64(kills*scalex), 252, float64(kills*scalex+deaths*scalex), 257)

		playerpw.SetColor(cConfig.KDAColor[2])
		playerdw.SetFillColor(playerpw)
		playerdw.Rectangle(float64(kills*scalex+deaths*scalex), 252, 190, 257)

		playerdw.SetFontSize(cConfig.TextSize15)
		playerdw.SetTextAlignment(imagick.ALIGN_CENTER)
		playerpw.SetColor(cConfig.Text)
		playerdw.SetFillColor(playerpw)

		playerdw.SetFont("Noto-Sans-CJK-TC-Regular")
		var name string
		if isFull == true {
			if matchData.Players[i].Name == "" {
				name = matchData.Players[i].Personaname
			} else {
				name = matchData.Players[i].Name
			}
		} else {
			name = getPlayerName(matchData.Players[i].AccountID, playersData)
		}

		fontSize := guessFontSize("Noto-Sans-CJK-TC-Regular", cConfig.TextSize12, 10, 50, 180, name)
		playerdw.SetFontSize(fontSize)
		playerdw.Annotation(95, 30, name)

		playerdw.SetFont("HypatiaSansPro-Regular")

		playerpw.SetColor("#00000077")
		playerdw.SetFillColor(playerpw)
		playerdw.Rectangle(0, 125, 174, 156)

		playerpw.SetColor(cConfig.Text)
		playerdw.SetFillColor(playerpw)
		playerdw.SetFontSize(cConfig.TextSize15)

		playerdw.SetTextAlignment(imagick.ALIGN_LEFT)
		if isFull {
			playerdw.Annotation(10, 148, getPlayerRole(i, matchData.Players[i].LaneRole))
		}

		playerdw.SetTextAlignment(imagick.ALIGN_CENTER)
		playerdw.Annotation(95, 229, fmt.Sprintf("%d / %d / %d", kills, deaths, assists))

		playerpw.SetColor(cConfig.TextGold)
		playerdw.SetFillColor(playerpw)
		playerdw.Circle(173, 140, 188, 140)
		playerpw.SetColor("black")
		playerdw.SetFillColor(playerpw)
		playerdw.Circle(173, 140, 186, 140)
		playerpw.SetColor(cConfig.TextGold)
		playerdw.SetFillColor(playerpw)
		playerdw.Annotation(174, 148, level)

		if healing == "0" {
			playerdw.SetTextAlignment(imagick.ALIGN_RIGHT)
			playerdw.Annotation(85, 200, gold)

			playerpw.SetColor(cConfig.TextDamage)
			playerdw.SetFillColor(playerpw)
			playerdw.SetTextAlignment(imagick.ALIGN_LEFT)
			playerdw.Annotation(105, 200, damage)
		} else {
			playerdw.SetTextAlignment(imagick.ALIGN_RIGHT)
			playerdw.Annotation(65, 200, gold)

			playerpw.SetColor(cConfig.TextDamage)
			playerdw.SetFillColor(playerpw)
			playerdw.SetTextAlignment(imagick.ALIGN_CENTER)
			playerdw.Annotation(95, 200, damage)

			playerpw.SetColor(cConfig.TextHealing)
			playerdw.SetFillColor(playerpw)
			playerdw.SetTextAlignment(imagick.ALIGN_LEFT)
			playerdw.Annotation(125, 200, healing)
		}

		playermw.DrawImage(playerdw)
		// playermw.WriteImage("./tmp/player" + strconv.Itoa(i) + ".png")
		if i < 5 {
			radiantPlayersmw.CompositeImage(playermw, imagick.COMPOSITE_OP_OVER, true, 195*i, 0)
		} else {
			direPlayersmw.CompositeImage(playermw, imagick.COMPOSITE_OP_OVER, true, 195*(i-5), 0)
		}
	}
	// direPlayersmw.WriteImage("./tmp/dire_players.png")
	// radiantPlayersmw.WriteImage("./tmp/radiant_players.png")
	mw.CompositeImage(direPlayersmw, imagick.COMPOSITE_OP_OVER, true, 25, 100)
	mw.CompositeImage(radiantPlayersmw, imagick.COMPOSITE_OP_OVER, true, 25, 663)

	// PLAYERS

	// PICKS

	var dir int
	var rad int

	direPicksmw := imagick.NewMagickWand()
	defer direPicksmw.Destroy()
	direPickspw := imagick.NewPixelWand()
	defer direPickspw.Destroy()
	direPicksdw := imagick.NewDrawingWand()
	defer direPicksdw.Destroy()
	radiantPicksmw := imagick.NewMagickWand()
	defer radiantPicksmw.Destroy()
	radiantPickspw := imagick.NewPixelWand()
	defer radiantPickspw.Destroy()
	radiantPicksdw := imagick.NewDrawingWand()
	defer radiantPicksdw.Destroy()

	direPickspw.SetColor("none")
	direPicksmw.NewImage(965, 46, direPickspw)
	radiantPickspw.SetColor("none")
	radiantPicksmw.NewImage(965, 46, radiantPickspw)

	for i := range matchData.PicksBans {
		heroImage := imagick.NewMagickWand()
		if matchData.PicksBans[i].IsPick {
			heroImage.ReadImage(fmt.Sprintf("./assets/heroes/%d.png", matchData.PicksBans[i].HeroID))
		} else {
			heroImage.ReadImage(fmt.Sprintf("./assets/heroes/bw/%d.png", matchData.PicksBans[i].HeroID))
		}
		heroImage.ResizeImage(83, 46, imagick.FILTER_CUBIC)
		if matchData.PicksBans[i].Team == 1 {
			var text string
			direPicksmw.CompositeImage(heroImage, imagick.COMPOSITE_OP_OVER, true, dir*88, 0)
			heroImage.Destroy()
			direPickspw.SetColor("#000000AA")
			direPicksdw.SetFillColor(direPickspw)
			direPicksdw.Rectangle(float64(dir*88), float64(34), float64(dir*88+83), float64(46))

			direPickspw.SetColor(cConfig.Text)
			direPicksdw.SetFillColor(direPickspw)
			direPicksdw.SetFontSize(cConfig.TextSize8)
			direPicksdw.SetTextAlignment(imagick.ALIGN_CENTER)
			if matchData.PicksBans[i].IsPick == true {
				text = "PICK " + strconv.Itoa(matchData.PicksBans[i].Order+1)
			} else {
				text = "BAN " + strconv.Itoa(matchData.PicksBans[i].Order+1)
			}
			direPicksdw.Annotation(float64(dir*88+44), float64(44), text)
			dir++
		} else {
			var text string
			radiantPicksmw.CompositeImage(heroImage, imagick.COMPOSITE_OP_OVER, true, rad*88, 0)
			heroImage.Destroy()
			radiantPickspw.SetColor("#000000AA")
			radiantPicksdw.SetFillColor(radiantPickspw)
			radiantPicksdw.Rectangle(float64(rad*88), float64(34), float64(rad*88+83), float64(46))

			radiantPickspw.SetColor(cConfig.Text)
			radiantPicksdw.SetFillColor(radiantPickspw)
			radiantPicksdw.SetFontSize(cConfig.TextSize8)
			radiantPicksdw.SetTextAlignment(imagick.ALIGN_CENTER)
			if matchData.PicksBans[i].IsPick == true {
				text = "PICK " + strconv.Itoa(matchData.PicksBans[i].Order+1)
			} else {
				text = "BAN " + strconv.Itoa(matchData.PicksBans[i].Order+1)
			}
			radiantPicksdw.Annotation(float64(rad*88+44), float64(44), text)
			rad++
		}
	}
	radiantPicksmw.DrawImage(radiantPicksdw)
	direPicksmw.DrawImage(direPicksdw)
	// radiantPicksmw.WriteImage("./tmp/radiant_picks.png")
	// direPicksmw.WriteImage("./tmp/dire_picks.png")

	mw.CompositeImage(direPicksmw, imagick.COMPOSITE_OP_OVER, true, 28, 366)
	mw.CompositeImage(radiantPicksmw, imagick.COMPOSITE_OP_OVER, true, 28, 608)

	// PICKS

	mw.DrawImage(dw)
	mw.WriteImage(savePath)
}
