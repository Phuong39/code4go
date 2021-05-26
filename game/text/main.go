package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

const sampleText = `The quick brown fox jumps
over the lazy dog.`

var (
	mplusNormalFont font.Face
	mplusBigFont    font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	mplusBigFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    32,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Game struct {
	counter        int
	kanjiText      []rune
	kanjiTextColor color.RGBA
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	gray := color.RGBA{0x80, 0x80, 0x80, 0xff}

	{
		const x, y = 20, 40
		b := text.BoundString(mplusNormalFont, sampleText)
		ebitenutil.DrawRect(screen, float64(b.Min.X+x), float64(b.Min.Y+y), float64(b.Dx()), float64(b.Dy()), gray)
		text.Draw(screen, sampleText, mplusNormalFont, x, y, color.White)
	}
	{
		const x, y = 20, 140
		b := text.BoundString(mplusBigFont, sampleText)
		ebitenutil.DrawRect(screen, float64(b.Min.X+x), float64(b.Min.Y+y), float64(b.Dx()), float64(b.Dy()), gray)
		text.Draw(screen, sampleText, mplusBigFont, x, y, color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Text (Ebiten Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
