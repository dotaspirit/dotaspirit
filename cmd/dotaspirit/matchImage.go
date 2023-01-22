package main

import (
	"fmt"
	"image/png"
	"math"
	"os"
	"strconv"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
)

func makeMatchImage(matchData oDotaMatchData, isFull bool) {

	fontHypatiaSansPro := canvas.NewFontFamily("dejavu")
	if err := fontHypatiaSansPro.LoadFontFile("./assets/fonts/HypatiaSansPro-Regular.otf", canvas.FontRegular); err != nil {
		panic(err)
	}
	if err := fontHypatiaSansPro.LoadFontFile("./assets/fonts/HypatiaSansPro-Bold.otf", canvas.FontBold); err != nil {
		panic(err)
	}

	savePath := fmt.Sprintf("./tmp/%d.png", matchData.MatchID)

	c := canvas.New(510, 510)

	ctx := canvas.NewContext(c)

	ctx.SetCoordSystem(canvas.CartesianIV)

	colorBG := parseHexColor(cConfig.Bg)
	colorText := parseHexColor(cConfig.Text)
	colorRadiant := parseHexColor(cConfig.TextRadiant)
	colorDire := parseHexColor(cConfig.TextDire)
	colorPlayerBack := parseHexColor(cConfig.PlayerBack)
	colorOverlay := parseHexColor(cConfig.Overlay)
	colorGold := parseHexColor(cConfig.TextGold)
	colorDamage := parseHexColor(cConfig.TextDamage)
	colorHealing := parseHexColor(cConfig.TextHealing)
	colorGraphLine := parseHexColor(cConfig.GraphLine)
	colorGraphXP := parseHexColor(cConfig.GraphXP)
	colorGraphGold := parseHexColor(cConfig.GraphGold)

	ctx.SetFillColor(colorBG)
	ctx.DrawPath(0, 0, canvas.Rectangle(510, 510))

	matchDuration := matchData.Duration
	hours, minutes, seconds := secondsToTime(matchDuration)

	sMatchID := strconv.FormatInt(matchData.MatchID, 10)

	face := fontHypatiaSansPro.Face(60, colorText, canvas.FontRegular, canvas.FontNormal)

	grad, err := os.Open("./assets/images/gradient.png")
	if err != nil {
		panic(err)
	}
	defer grad.Close()
	gradImg, err := png.Decode(grad)
	if err != nil {
		panic(err)
	}

	// Corner texts
	ctx.DrawText(13, 30, canvas.NewTextLine(face, fmt.Sprintf("%d:%02d:%02d", hours, minutes, seconds), canvas.Left))
	ctx.DrawText(497, 30, canvas.NewTextLine(face, sMatchID, canvas.Right))
	ctx.DrawText(13, 491, canvas.NewTextLine(face, matchData.League.Name, canvas.Left))
	ctx.DrawImage(360, 460, gradImg, canvas.DPMM(6))
	ctx.DrawText(497, 491, canvas.NewTextLine(face, "#rsltdtk", canvas.Right))

	// Graph

	if isFull {
		graphScaleX := 483 / float64(matchDuration/60)

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

		maxAdv := 0

		if maxXpAdv > maxGoldAdv {
			maxAdv = maxXpAdv
		} else {
			maxAdv = maxGoldAdv
		}

		graphScaleY := (80 / 2) / math.Abs(float64(maxAdv))

		// Line
		p := &canvas.Path{}
		p.LineTo(483, 0)

		ctx.SetStrokeColor(colorGraphLine)
		ctx.SetFillColor(canvas.Transparent)
		ctx.DrawPath(13, 255, p)

		// XP

		p = &canvas.Path{}
		p.LineTo(0, 0)
		for i, adv := range xpAdv {
			sx := float64(i) * graphScaleX
			sy := float64(-adv) * graphScaleY
			p.LineTo(sx, sy)
		}

		ctx.SetStrokeColor(colorGraphXP)
		ctx.DrawPath(13, 255, p)

		// Gold

		p = &canvas.Path{}
		p.LineTo(0, 0)
		for i, adv := range goldAdv {
			sx := float64(i) * graphScaleX
			sy := float64(-adv) * graphScaleY
			p.LineTo(sx, sy)
		}

		ctx.SetStrokeColor(colorGraphGold)
		ctx.DrawPath(13, 255, p)
	}

	// Score text
	face = fontHypatiaSansPro.Face(100, colorText, canvas.FontRegular, canvas.FontNormal)
	ctx.DrawText(255, 242, canvas.NewTextLine(face, strconv.Itoa(matchData.RadiantScore), canvas.Center))
	ctx.DrawText(255, 292, canvas.NewTextLine(face, strconv.Itoa(matchData.DireScore), canvas.Center))

	radiantSeriesScore, direSeriesScore := getSeriesScore(matchData.RadiantTeamID, matchData.DireTeamID, matchData.SeriesID, matchData.MatchID)

	ctx.SetStrokeColor(colorGold)
	switch matchData.SeriesType {
	case 1:
		for i := 0; i < 2; i++ {
			if i < radiantSeriesScore {
				ctx.SetFillColor(colorGold)
				ctx.DrawPath(250+6*float64(i), 208, canvas.Rectangle(3, 5))
			} else {
				ctx.SetFillColor(canvas.Transparent)
				ctx.DrawPath(250+6*float64(i), 208, canvas.Rectangle(3, 5))
			}
		}
		for i := 0; i < 2; i++ {
			if i < direSeriesScore {
				ctx.SetFillColor(colorGold)
				ctx.DrawPath(250+6*float64(i), 297, canvas.Rectangle(3, 5))
			} else {
				ctx.SetFillColor(canvas.Transparent)
				ctx.DrawPath(250+6*float64(i), 208, canvas.Rectangle(3, 5))
			}
		}
	case 2:
		for i := 0; i < 3; i++ {
			if i < radiantSeriesScore {
				ctx.SetFillColor(colorGold)
				ctx.DrawPath(247+6*float64(i), 208, canvas.Rectangle(3, 5))
			} else {
				ctx.SetFillColor(canvas.Transparent)
				ctx.DrawPath(247+6*float64(i), 208, canvas.Rectangle(3, 5))
			}
		}
		for i := 0; i < 3; i++ {
			if i < direSeriesScore {
				ctx.SetFillColor(colorGold)
				ctx.DrawPath(247+6*float64(i), 297, canvas.Rectangle(3, 5))
			} else {
				ctx.SetFillColor(canvas.Transparent)
				ctx.DrawPath(247+6*float64(i), 297, canvas.Rectangle(3, 5))
			}
		}
	}

	if matchData.RadiantWin {
		face = fontHypatiaSansPro.Face(60, colorRadiant, canvas.FontRegular, canvas.FontNormal, canvas.FontUnderline)
		ctx.DrawText(13, 234, canvas.NewTextLine(face, matchData.RadiantTeam.Name, canvas.Left))
		face = fontHypatiaSansPro.Face(60, colorDire, canvas.FontRegular, canvas.FontNormal)
		ctx.DrawText(13, 285, canvas.NewTextLine(face, matchData.DireTeam.Name, canvas.Left))
	} else {
		face = fontHypatiaSansPro.Face(60, colorRadiant, canvas.FontRegular, canvas.FontNormal)
		ctx.DrawText(13, 234, canvas.NewTextLine(face, matchData.RadiantTeam.Name, canvas.Left))
		face = fontHypatiaSansPro.Face(60, colorDire, canvas.FontRegular, canvas.FontNormal, canvas.FontUnderline)
		ctx.DrawText(13, 285, canvas.NewTextLine(face, matchData.DireTeam.Name, canvas.Left))
	}

	// Players
	for i, player := range matchData.Players {
		ctx.SetStrokeColor(canvas.Transparent)
		face = fontHypatiaSansPro.Face(31.4285714286, colorText, canvas.FontRegular, canvas.FontNormal)
		hero, err := os.Open(fmt.Sprintf("./assets/heroes/%d.png", player.HeroID))
		if err != nil {
			panic(err)
		}
		defer hero.Close()
		heroImg, err := png.Decode(hero)
		if err != nil {
			panic(err)
		}
		stepDist := (95 + 2) * float64(i%5)
		startPosY := 0.0

		if i < 5 {
			startPosY = 50
		} else {
			startPosY = 333
		}
		// Background
		ctx.SetFillColor(colorPlayerBack)
		ctx.DrawPath(13+stepDist, startPosY-50+50, canvas.RoundedRectangle(95, 127, 6))
		// Player nickname
		ctx.DrawText(61+stepDist, startPosY-50+65, canvas.NewTextLine(face, player.Name, canvas.Center))
		// Hero image
		ctx.DrawImage(13+stepDist, startPosY-50+74, heroImg, canvas.DPMM(2.69473684211))
		ctx.SetFillColor(colorOverlay)
		ctx.DrawPath(13+stepDist, startPosY-50+112, canvas.Rectangle(95, 16))
		if isFull {
			ctx.DrawText(18+stepDist, startPosY-50+124, canvas.NewTextLine(face, getPlayerRole(player.LaneRole), canvas.Left))
		}
		// Player color
		ctx.SetStrokeColor(canvas.Transparent)
		ctx.SetFillColor(parseHexColor(cConfig.PlayerColor[i]))
		ctx.DrawPath(13+stepDist, startPosY-50+128, canvas.Rectangle(95, 3))
		// Level
		ctx.SetFillColor(canvas.Black)
		ctx.SetStrokeColor(colorGold)
		ctx.DrawPath(99+stepDist, startPosY-50+120, canvas.Circle(8))
		face = fontHypatiaSansPro.Face(31.4285714286, colorGold, canvas.FontRegular, canvas.FontNormal)
		ctx.DrawText(99+stepDist, startPosY-50+124, canvas.NewTextLine(face, strconv.Itoa(player.Level), canvas.Center))
		// Gold Damage Healing
		if player.HeroHealing > 0 {
			face = fontHypatiaSansPro.Face(31.4285714286, colorGold, canvas.FontRegular, canvas.FontNormal)
			ctx.DrawText(42+stepDist, startPosY-50+151, canvas.NewTextLine(face, prettifyBigValues(player.NetWorth), canvas.Right))

			face = fontHypatiaSansPro.Face(31.4285714286, colorDamage, canvas.FontRegular, canvas.FontNormal)
			ctx.DrawText(61+stepDist, startPosY-50+151, canvas.NewTextLine(face, prettifyBigValues(player.HeroDamage), canvas.Center))

			face = fontHypatiaSansPro.Face(31.4285714286, colorHealing, canvas.FontRegular, canvas.FontNormal)
			ctx.DrawText(78+stepDist, startPosY-50+151, canvas.NewTextLine(face, prettifyBigValues(player.HeroHealing), canvas.Left))
		} else {
			face = fontHypatiaSansPro.Face(31.4285714286, colorGold, canvas.FontRegular, canvas.FontNormal)
			ctx.DrawText(56+stepDist, startPosY-50+151, canvas.NewTextLine(face, prettifyBigValues(player.NetWorth), canvas.Right))

			face = fontHypatiaSansPro.Face(31.4285714286, colorDamage, canvas.FontRegular, canvas.FontNormal)
			ctx.DrawText(64+stepDist, startPosY-50+151, canvas.NewTextLine(face, prettifyBigValues(player.HeroDamage), canvas.Left))
		}
		// KDA
		face = fontHypatiaSansPro.Face(31.4285714286, colorText, canvas.FontRegular, canvas.FontNormal)
		ctx.DrawText(61+stepDist, startPosY-50+167, canvas.NewTextLine(face, fmt.Sprintf("%d / %d / %d", player.Kills, player.Deaths, player.Assists), canvas.Center))
	}

	// Picks and bans

	radiantPickBan := 0
	direPickBan := 0
	ctx.SetStrokeColor(canvas.Transparent)

	face = fontHypatiaSansPro.Face(17.1428571429, colorText, canvas.FontRegular, canvas.FontNormal)

	for _, pickBan := range matchData.PicksBans {
		if pickBan.IsPick {
			hero, err := os.Open(fmt.Sprintf("./assets/heroes/%d.png", pickBan.HeroID))
			if err != nil {
				panic(err)
			}
			defer hero.Close()
			heroImg, err := png.Decode(hero)
			if err != nil {
				panic(err)
			}
			if pickBan.Team == 0 {
				ctx.DrawImage(15+(39+1)*float64(radiantPickBan), 179, heroImg, canvas.DPMM(6.5641025641))
				ctx.SetFillColor(colorOverlay)
				ctx.DrawPath(15+(39+1)*float64(radiantPickBan), 195, canvas.Rectangle(39, 6))
				ctx.DrawText(35+(39+1)*float64(radiantPickBan), 200, canvas.NewTextLine(face, "PICK "+strconv.Itoa(pickBan.Order+1), canvas.Center))
				radiantPickBan++
			} else {
				ctx.DrawImage(15+(39+1)*float64(direPickBan), 309, heroImg, canvas.DPMM(6.5641025641))
				ctx.SetFillColor(colorOverlay)
				ctx.DrawPath(15+(39+1)*float64(direPickBan), 325, canvas.Rectangle(39, 6))
				ctx.DrawText(35+(39+1)*float64(direPickBan), 330, canvas.NewTextLine(face, "PICK "+strconv.Itoa(pickBan.Order+1), canvas.Center))
				direPickBan++
			}
		} else {
			hero, err := os.Open(fmt.Sprintf("./assets/heroes/bw/%d.png", pickBan.HeroID))
			if err != nil {
				panic(err)
			}
			defer hero.Close()
			heroImg, err := png.Decode(hero)
			if err != nil {
				panic(err)
			}
			if pickBan.Team == 0 {
				ctx.DrawImage(15+(39+1)*float64(radiantPickBan), 179, heroImg, canvas.DPMM(6.5641025641))
				ctx.SetFillColor(colorOverlay)
				ctx.DrawPath(15+(39+1)*float64(radiantPickBan), 195, canvas.Rectangle(39, 6))
				ctx.DrawText(35+(39+1)*float64(radiantPickBan), 200, canvas.NewTextLine(face, "BAN "+strconv.Itoa(pickBan.Order+1), canvas.Center))
				radiantPickBan++
			} else {
				ctx.DrawImage(15+(39+1)*float64(direPickBan), 309, heroImg, canvas.DPMM(6.5641025641))
				ctx.SetFillColor(colorOverlay)
				ctx.DrawPath(15+(39+1)*float64(direPickBan), 325, canvas.Rectangle(39, 6))
				ctx.DrawText(35+(39+1)*float64(direPickBan), 330, canvas.NewTextLine(face, "BAN "+strconv.Itoa(pickBan.Order+1), canvas.Center))
				direPickBan++
			}
		}
	}

	renderers.Write(savePath, c, canvas.DPMM(4))
}
