package main

import (
	"fmt"
	"image"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/generaltso/vibrant"
	"gopkg.in/gographics/imagick.v3/imagick"
)

func getImageColors(file string) (fgColor, bgColor string) {
	checkErr := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	f, err := os.Open(file)
	checkErr(err)
	defer f.Close()

	img, _, err := image.Decode(f)
	checkErr(err)

	palette, err := vibrant.NewPaletteFromImage(img)
	checkErr(err)

	for name, swatch := range palette.ExtractAwesome() {
		if name == "Vibrant" {
			fgColor = swatch.Color.TitleTextColor().RGBHex()
			bgColor = swatch.Color.RGBHex()
		}
	}
	if fgColor == bgColor {
		fgColor = "#000"
	}
	return fgColor, bgColor
}

func createStoryText(fgColor, bgColor, text string) *imagick.MagickWand {
	mw := imagick.NewMagickWand()
	dw := imagick.NewDrawingWand()
	defer dw.Destroy()
	pw := imagick.NewPixelWand()
	defer pw.Destroy()

	pw.SetColor("none")
	mw.NewImage(1080, 103, pw)

	dw.SetFont("TTCommonsB")
	dw.SetFontSize(80)

	textMetrics := mw.QueryFontMetrics(dw, text)

	textWidth := textMetrics.TextWidth

	pw.SetColor(bgColor)
	dw.SetFillColor(pw)

	dw.RoundRectangle((1080-textWidth)/2-20, 0, (1080-textWidth)/2+textWidth+20, 103, 10, 10)

	pw.SetColor(fgColor)
	dw.SetFillColor(pw)
	dw.SetTextAlignment(imagick.ALIGN_CENTER)
	dw.Annotation(1080/2, 93-17, text)

	mw.DrawImage(dw)

	return mw
}

func downloadTeamLogo(filepath, url string) (err error) {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func getTeamsLogo(radiantTeamID, direTeamID int) {
	teamsData := getTeamsData()
	for _, team := range teamsData {
		teamID := team.TeamID
		if teamID == radiantTeamID || teamID == direTeamID {
			logo := team.LogoURL
			if logo == "" {
				logo = "https://github.com/dotaspirit/dotaspirit/raw/master/assets/images/nologo.png"
			}
			err := downloadTeamLogo(fmt.Sprintf("tmp/teams/%d.png", teamID), logo)
			if err != nil {
				log.Printf("Error downloading %d team logo", teamID)
			}
		}
	}
}

func guessResizeSize(width, height uint) (wOut, hOut uint) {
	if width > height {
		mult := 294 / width
		wOut = width * mult
		hOut = height * mult
	} else {
		mult := 294 / height
		wOut = width * mult
		hOut = height * mult
	}
	return uint(wOut), uint(hOut)
}

func makeStoryImage(matchData oDotaMatchData) {
	radiantTeamName := matchData.RadiantName
	direTeamName := matchData.DireName
	radiantTeamID := matchData.RadiantTeamID
	direTeamID := matchData.DireTeamID
	matchID := matchData.MatchID

	getTeamsLogo(radiantTeamID, direTeamID)

	direTeamFile := fmt.Sprintf("tmp/teams/%d.png", direTeamID)
	radiantTeamFile := fmt.Sprintf("tmp/teams/%d.png", radiantTeamID)

	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	dw := imagick.NewDrawingWand()
	defer dw.Destroy()
	pw := imagick.NewPixelWand()
	defer pw.Destroy()

	dw.SetFont("TTCommonsB")

	pw.SetColor("#242424")
	mw.NewImage(1080, 1920, pw)

	fgColor, bgColor := getImageColors(direTeamFile)

	direTeamText := createStoryText(fgColor, bgColor, direTeamName)

	mw.CompositeImage(direTeamText, imagick.COMPOSITE_OP_OVER, true, 0, ((1920/2)-294)/2+250)

	direTeamText.Destroy()

	fgColor, bgColor = getImageColors(radiantTeamFile)

	radiantTeamText := createStoryText(fgColor, bgColor, radiantTeamName)

	mw.CompositeImage(radiantTeamText, imagick.COMPOSITE_OP_OVER, true, 0, 1920/2+((1920/2)-294)/2+250)

	radiantTeamText.Destroy()

	direTeamLogo := imagick.NewMagickWand()
	direTeamLogo.ReadImage(direTeamFile)

	width := direTeamLogo.GetImageWidth()
	height := direTeamLogo.GetImageHeight()
	wGuessed, hGuessed := guessResizeSize(width, height)

	direTeamLogo.ResizeImage(wGuessed, hGuessed, imagick.FILTER_CUBIC)
	mw.CompositeImage(direTeamLogo, imagick.COMPOSITE_OP_OVER, true, (1080-294)/2, ((1920/2)-294)/2-100)
	direTeamLogo.Destroy()

	radiantTeamLogo := imagick.NewMagickWand()
	radiantTeamLogo.ReadImage(radiantTeamFile)

	width = radiantTeamLogo.GetImageWidth()
	height = radiantTeamLogo.GetImageHeight()
	wGuessed, hGuessed = guessResizeSize(width, height)

	radiantTeamLogo.ResizeImage(wGuessed, hGuessed, imagick.FILTER_CUBIC)
	mw.CompositeImage(radiantTeamLogo, imagick.COMPOSITE_OP_OVER, true, (1080-294)/2, 1920/2+((1920/2)-294)/2-100)
	radiantTeamLogo.Destroy()

	vs := createStoryText("#242424", "#ffffff", "VS")
	mw.CompositeImage(vs, imagick.COMPOSITE_OP_OVER, true, 0, 1920/2-51)
	vs.Destroy()

	mw.DrawImage(dw)
	storyFile := fmt.Sprintf("tmp/s%d.png", matchID)
	mw.WriteImage(storyFile)
}
