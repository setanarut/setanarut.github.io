package main

import (
	"bytes"
	"image"
	"image/color"
	"log"

	"github.com/gio-eui/md3-fonts/fonts/robotomono/robotomonosemibold"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/setanarut/gog"
)

var roboto *text.GoTextFace
var letters []text.Glyph
var screenSize = image.Point{800, 400}
var dio = &ebiten.DrawImageOptions{Filter: ebiten.FilterLinear}
var path *gog.Path = gog.Lemniscate(300, 348)
var txt = `A dead simple 2D game engine for Go`
var tick float64
var pathlen float64

type Game struct{}

func main() {
	src, err := text.NewGoTextFaceSource(bytes.NewReader(robotomonosemibold.Data))
	if err != nil {
		log.Fatal(err)
	}
	roboto = &text.GoTextFace{Source: src, Size: 50}
	path.Translate(400, 200)
	path.Reverse()
	path.Close()
	pathlen = path.Length()
	dio.ColorScale.ScaleWithColor(color.RGBA{219, 86, 32, 255})
	letters = text.AppendGlyphs(letters, txt, roboto, &text.LayoutOptions{})
	ebiten.SetWindowSize(screenSize.X, screenSize.Y)
	ebiten.SetScreenClearedEveryFrame(false)
	ebiten.RunGameWithOptions(&Game{}, &ebiten.RunGameOptions{DisableHiDPI: true})
}

func (g *Game) Update() error {
	if tick > path.Length() {
		tick = 0
	}
	tick += 4
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{13, 17, 23, 255})

	for _, glyph := range letters {
		if glyph.Image != nil {
			l := (glyph.X + float64(glyph.Image.Bounds().Dx())/2)
			l += tick
			if l > pathlen {
				l = l - pathlen
			}
			point, angle := path.PointAngleAtLength(l)
			dio.GeoM.Reset()
			dio.GeoM.Translate(-float64(glyph.Image.Bounds().Dx())/2, -(glyph.OriginY - glyph.Y))
			dio.GeoM.Rotate(angle)
			dio.GeoM.Translate(point.X, point.Y)
			screen.DrawImage(glyph.Image, dio)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenSize.X, screenSize.Y
}
