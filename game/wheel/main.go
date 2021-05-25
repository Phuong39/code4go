package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var pointerImage = ebiten.NewImage(8, 8)

func init() {
	pointerImage.Fill(color.RGBA{0xff, 0, 0, 0xff})
}

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	x float64
	y float64
}

func (g *Game) Update() error {
	dx, dy := ebiten.Wheel()
	g.x += dx
	g.y += dy
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.x, g.y)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	screen.DrawImage(pointerImage, op)

	ebitenutil.DebugPrint(screen,
		fmt.Sprintf("Move the red point by mouse wheel\n(%0.2f, %0.2f)", g.x, g.y))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	g := &Game{x: 0.0, y: 0.0}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Wheel (Ebiten Demo)")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
