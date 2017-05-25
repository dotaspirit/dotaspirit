package images

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"

	"github.com/dotaspirit/dotaspirit/types"

	"gopkg.in/gographics/imagick.v2/imagick"
)

func getRplayerRole(pos, role int) string {
	roleNames := [5]string{"Unknown", "Safe", "Mid", "Offlane", "Jungle"}
	return roleNames[role]
}

func prettifyBigValues(value int) string {
	if value < 1000 {
		return strconv.Itoa(value)
	} else {
		result := fmt.Sprintf("%.1fk", float64(value)/1000)
		return result
	}
}

func loadColorConfig() colorConfig {
	buf := bytes.NewBuffer(nil)
	f, _ := os.Open("config/colors.json") // Error handling elided for brevity.
	io.Copy(buf, f)                       // Error handling elided for brevity.
	f.Close()

	var jsonobject colorConfig

	err := json.Unmarshal(buf.Bytes(), &jsonobject)

	if err != nil {
		fmt.Println("error:", err)
	}

	return jsonobject
}

func loadMatchData(matchID int64) types.Match {
	var sMatchID = strconv.FormatInt(matchID, 10)

	buf := bytes.NewBuffer(nil)
	f, _ := os.Open("./tmp/matches/" + sMatchID) // Error handling elided for brevity.
	io.Copy(buf, f)                              // Error handling elided for brevity.
	f.Close()

	var jsonobject types.Match
	err := json.Unmarshal(buf.Bytes(), &jsonobject)

	if err != nil {
		fmt.Println("error:", err)
	}

	return jsonobject
}

func MakeResultImage(matchID int64, ret map[int64]int) (err error) {
	sMatchID := strconv.FormatInt(matchID, 10)
	config := loadColorConfig()
	matchData := loadMatchData(matchID)
	retries := ret[matchID]

	if (len(matchData.RadiantGoldAdv) == 0 ||
		(len(matchData.PicksBans) == 0 && matchData.GameMode == 2)) &&
		retries < 180 {
		return fmt.Errorf("Match data is not full, %d retries", retries)
	}

	matchDuration := matchData.Duration
	minutes := matchDuration / 60
	seconds := matchDuration % 60
	hours := minutes / 60
	minutes = minutes % 60

	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	dw := imagick.NewDrawingWand()
	defer dw.Destroy()
	pw := imagick.NewPixelWand()
	defer pw.Destroy()

	dw.SetFont("HypatiaSansPro-Regular")

	pw.SetColor(config.Bg)
	mw.NewImage(1020, 1020, pw)

	// GRAPH
	if retries < 180 {
		graphmw := imagick.NewMagickWand()
		graphdw := imagick.NewDrawingWand()
		graphpw := imagick.NewPixelWand()
		graphpw.SetColor("none")
		graphmw.NewImage(970, 114, graphpw)
		graphpw.SetColor(config.TextDark)

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
				graphpw.SetColor(config.TextTFRadiant)
			} else if direDeaths > radiantDeaths {
				graphpw.SetColor(config.TextTFDire)
			} else {
				graphpw.SetColor(config.TextTF)
			}
			graphdw.SetStrokeColor(graphpw)
			graphdw.Line(float64(tf.Start)*scalex, 56, float64(tf.End)*scalex, 56)
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

		graphpw.SetColor(config.TextXP)
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

		graphpw.SetColor(config.TextGold)
		graphdw.SetStrokeColor(graphpw)
		for i, adv := range goldAdv {
			nextx := float64(i) * scalex
			nexty := 56 + float64(adv)*scaley
			graphdw.Line(sx, sy, nextx, nexty)
			sx = nextx
			sy = nexty
		}

		graphmw.DrawImage(graphdw)
		mw.CompositeImage(graphmw, imagick.COMPOSITE_OP_OVER, 25, 453)
		// graphmw.WriteImage("./tmp/graph.png")
		graphmw.Destroy()
		graphdw.Destroy()
		graphpw.Destroy()
	}
	// GRAPH
	// TEXT

	pw.SetColor(config.Text)
	dw.SetFontSize(config.TextSize28)
	dw.SetFillColor(pw)

	dw.SetTextAlignment(imagick.ALIGN_LEFT)
	dw.Annotation(25, 63, fmt.Sprintf("%d:%02d:%02d", hours, minutes, seconds))
	dw.Annotation(25, 986, matchData.League.Name)
	dw.SetTextAlignment(imagick.ALIGN_RIGHT)
	dw.Annotation(995, 63, sMatchID)
	dw.Annotation(995, 986, "#rsltdtk")

	dw.SetFontSize(config.TextSize48)
	dw.SetTextAlignment(imagick.ALIGN_CENTER)
	dw.Annotation(510, 484, strconv.Itoa(matchData.DireScore))
	dw.Annotation(510, 584, strconv.Itoa(matchData.RadiantScore))

	radiantName := matchData.RadiantTeam.Name
	direName := matchData.DireTeam.Name
	if radiantName == "" {
		radiantName = "Radiant"
	}
	if direName == "" {
		direName = "Dire"
	}

	dw.SetFontSize(config.TextSize28)
	pw.SetColor(config.TextRadiant)
	dw.SetFillColor(pw)
	dw.SetTextAlignment(imagick.ALIGN_LEFT)
	if matchData.RadiantWin == true {
		dw.Annotation(64, 569, radiantName)
	} else {
		dw.Annotation(25, 569, radiantName)
	}
	pw.SetColor(config.TextDire)
	dw.SetFillColor(pw)
	if matchData.RadiantWin == false {
		dw.Annotation(64, 469, direName)
	} else {
		dw.Annotation(25, 469, direName)
	}
	// TEXT

	winnerImage := imagick.NewMagickWand()
	winnerImage.ReadImage("./assets/images/winner.png")
	if matchData.RadiantWin == true {
		mw.CompositeImage(winnerImage, imagick.COMPOSITE_OP_OVER, 25, 541)
	} else {
		mw.CompositeImage(winnerImage, imagick.COMPOSITE_OP_OVER, 25, 441)
	}
	winnerImage.Destroy()

	//PLAYERS
	direPlayersmw := imagick.NewMagickWand()
	direPlayerspw := imagick.NewPixelWand()
	radiantPlayersmw := imagick.NewMagickWand()
	radiantPlayerspw := imagick.NewPixelWand()

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
		var damage = prettifyBigValues(matchData.Players[i].HeroDamage)
		var healing = prettifyBigValues(matchData.Players[i].HeroHealing)
		var level = strconv.Itoa(matchData.Players[i].Level)
		if i < 5 {
			playerpw.SetColor(config.BackRadiant)
		} else {
			playerpw.SetColor(config.BackDire)
		}

		playermw.NewImage(190, 257, playerpw)

		heroImage := imagick.NewMagickWand()
		heroImage.ReadImage(fmt.Sprintf("./assets/heroes/%d.png", matchData.Players[i].HeroID))
		heroImage.ResizeImage(190, 107, imagick.FILTER_CUBIC, 1)
		playermw.CompositeImage(heroImage, imagick.COMPOSITE_OP_OVER, 0, 50)
		heroImage.Destroy()

		playerpw.SetColor(config.PlayerColor[i])
		playerdw.SetFillColor(playerpw)
		playerdw.Rectangle(0, 157, 190, 162)

		kills := matchData.Players[i].Kills
		deaths := matchData.Players[i].Deaths
		assists := matchData.Players[i].Assists

		scalex := 190

		if kills+deaths+assists != 0 {
			scalex = 190 / (kills + deaths + assists)
		}

		playerpw.SetColor(config.KDAColor[0])
		playerdw.SetFillColor(playerpw)
		playerdw.Rectangle(0, 252, float64(kills*scalex), 257)

		playerpw.SetColor(config.KDAColor[1])
		playerdw.SetFillColor(playerpw)
		playerdw.Rectangle(float64(kills*scalex), 252, float64(kills*scalex+deaths*scalex), 257)

		playerpw.SetColor(config.KDAColor[2])
		playerdw.SetFillColor(playerpw)
		playerdw.Rectangle(float64(kills*scalex+deaths*scalex), 252, 190, 257)

		playerdw.SetFontSize(config.TextSize15)
		playerdw.SetTextAlignment(imagick.ALIGN_CENTER)
		playerpw.SetColor(config.Text)
		playerdw.SetFillColor(playerpw)

		playerdw.SetFont("Noto-Sans-CJK-TC-Regular")
		var name string
		if matchData.Players[i].Name == "" {
			name = matchData.Players[i].Personaname
		} else {
			name = matchData.Players[i].Name
		}
		if len(name) > 15 {
			playerdw.SetFontSize(config.TextSize12)
		}
		if len(name) > 25 {
			playerdw.SetFontSize(config.TextSize8)
		}
		playerdw.Annotation(95, 30, name)

		playerdw.SetFont("HypatiaSansPro-Regular")

		playerpw.SetColor("#00000077")
		playerdw.SetFillColor(playerpw)
		playerdw.Rectangle(0, 125, 174, 156)

		playerpw.SetColor(config.Text)
		playerdw.SetFillColor(playerpw)

		playerdw.SetTextAlignment(imagick.ALIGN_LEFT)
		playerdw.Annotation(10, 148, getRplayerRole(i, matchData.Players[i].LaneRole))

		playerdw.SetTextAlignment(imagick.ALIGN_CENTER)
		playerdw.SetFontSize(config.TextSize15)
		playerdw.Annotation(95, 229, fmt.Sprintf("%d / %d / %d", kills, deaths, assists))

		playerpw.SetColor(config.TextGold)
		playerdw.SetFillColor(playerpw)
		playerdw.Circle(173, 140, 188, 140)
		playerpw.SetColor("black")
		playerdw.SetFillColor(playerpw)
		playerdw.Circle(173, 140, 186, 140)
		playerpw.SetColor(config.TextGold)
		playerdw.SetFillColor(playerpw)
		playerdw.Annotation(175, 148, level)

		if healing == "0" {
			playerdw.SetTextAlignment(imagick.ALIGN_RIGHT)
			playerdw.Annotation(85, 200, gold)

			playerpw.SetColor(config.TextDamage)
			playerdw.SetFillColor(playerpw)
			playerdw.SetTextAlignment(imagick.ALIGN_LEFT)
			playerdw.Annotation(105, 200, damage)
		} else {
			playerdw.SetTextAlignment(imagick.ALIGN_RIGHT)
			playerdw.Annotation(65, 200, gold)

			playerpw.SetColor(config.TextDamage)
			playerdw.SetFillColor(playerpw)
			playerdw.SetTextAlignment(imagick.ALIGN_CENTER)
			playerdw.Annotation(95, 200, damage)

			playerpw.SetColor(config.TextHealing)
			playerdw.SetFillColor(playerpw)
			playerdw.SetTextAlignment(imagick.ALIGN_LEFT)
			playerdw.Annotation(125, 200, healing)
		}

		playermw.DrawImage(playerdw)
		// playermw.WriteImage("./tmp/player" + strconv.Itoa(i) + ".png")
		if i < 5 {
			radiantPlayersmw.CompositeImage(playermw, imagick.COMPOSITE_OP_OVER, 195*i, 0)
		} else {
			direPlayersmw.CompositeImage(playermw, imagick.COMPOSITE_OP_OVER, 195*(i-5), 0)
		}
	}
	// direPlayersmw.WriteImage("./tmp/dire_players.png")
	// radiantPlayersmw.WriteImage("./tmp/radiant_players.png")
	mw.CompositeImage(direPlayersmw, imagick.COMPOSITE_OP_OVER, 25, 100)
	mw.CompositeImage(radiantPlayersmw, imagick.COMPOSITE_OP_OVER, 25, 663)
	direPlayersmw.Destroy()
	direPlayerspw.Destroy()
	radiantPlayersmw.Destroy()
	radiantPlayerspw.Destroy()
	//PLAYERS

	// BANS
	var dir int
	var rad int

	direPicksmw := imagick.NewMagickWand()
	direPickspw := imagick.NewPixelWand()
	direPicksdw := imagick.NewDrawingWand()
	radiantPicksmw := imagick.NewMagickWand()
	radiantPickspw := imagick.NewPixelWand()
	radiantPicksdw := imagick.NewDrawingWand()

	direPickspw.SetColor("none")
	direPicksmw.NewImage(965, 52, direPickspw)
	radiantPickspw.SetColor("none")
	radiantPicksmw.NewImage(965, 52, radiantPickspw)

	for i := range matchData.PicksBans {
		heroImage := imagick.NewMagickWand()
		defer heroImage.Destroy()
		if matchData.PicksBans[i].IsPick == true {
			heroImage.ReadImage(fmt.Sprintf("./assets/heroes/%d.png", matchData.PicksBans[i].HeroID))
		} else {
			heroImage.ReadImage(fmt.Sprintf("./assets/heroes/bw/%d.png", matchData.PicksBans[i].HeroID))
		}
		heroImage.ResizeImage(92, 52, imagick.FILTER_CUBIC, 1)
		if matchData.PicksBans[i].Team == 1 {
			var text string
			direPicksmw.CompositeImage(heroImage, imagick.COMPOSITE_OP_OVER, dir*97, 0)
			direPickspw.SetColor("#000000AA")
			direPicksdw.SetFillColor(direPickspw)
			direPicksdw.Rectangle(float64(dir*97), float64(40), float64(dir*97+91), float64(52))

			direPickspw.SetColor(config.Text)
			direPicksdw.SetFillColor(direPickspw)
			direPicksdw.SetFontSize(config.TextSize8)
			direPicksdw.SetTextAlignment(imagick.ALIGN_CENTER)
			if matchData.PicksBans[i].IsPick == true {
				text = "PICK " + strconv.Itoa(matchData.PicksBans[i].Order+1)
			} else {
				text = "BAN " + strconv.Itoa(matchData.PicksBans[i].Order+1)
			}
			direPicksdw.Annotation(float64(dir*97+46), float64(50), text)
			dir++
		} else {
			var text string
			radiantPicksmw.CompositeImage(heroImage, imagick.COMPOSITE_OP_OVER, rad*97, 0)
			radiantPickspw.SetColor("#000000AA")
			radiantPicksdw.SetFillColor(radiantPickspw)
			radiantPicksdw.Rectangle(float64(rad*97), float64(40), float64(rad*97+91), float64(52))

			radiantPickspw.SetColor(config.Text)
			radiantPicksdw.SetFillColor(radiantPickspw)
			radiantPicksdw.SetFontSize(config.TextSize8)
			radiantPicksdw.SetTextAlignment(imagick.ALIGN_CENTER)
			if matchData.PicksBans[i].IsPick == true {
				text = "PICK " + strconv.Itoa(matchData.PicksBans[i].Order+1)
			} else {
				text = "BAN " + strconv.Itoa(matchData.PicksBans[i].Order+1)
			}
			radiantPicksdw.Annotation(float64(rad*97+46), float64(50), text)
			rad++
		}
	}
	radiantPicksmw.DrawImage(radiantPicksdw)
	direPicksmw.DrawImage(direPicksdw)
	// radiantPicksmw.WriteImage("./tmp/radiant_picks.png")
	// direPicksmw.WriteImage("./tmp/dire_picks.png")

	mw.CompositeImage(direPicksmw, imagick.COMPOSITE_OP_OVER, 28, 362)
	mw.CompositeImage(radiantPicksmw, imagick.COMPOSITE_OP_OVER, 28, 606)
	radiantPicksmw.Destroy()
	direPicksmw.Destroy()
	radiantPlayersmw.Destroy()
	direPlayerspw.Destroy()

	mw.DrawImage(dw)
	mw.WriteImage("./tmp/images/" + sMatchID + ".png")
	fmt.Printf("[IMAGES] Done for match #%d\n", matchID)
	delete(ret, matchID)

	return nil
}
