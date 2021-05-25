package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth     = 640
	screenHeight    = 480
	initScreenScale = 1
)

var (
	dots       []byte
	dotsWidth  int
	dotsHeight int
)

func getDots(width, height int) []byte {
	if dotsWidth == width && dotsHeight == height {
		return dots
	}
	dotsWidth = width
	dotsHeight = height
	dots = make([]byte, width*height*4)
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			if (i+j)%2 == 0 {
				dots[(i+j*width)*4+0] = 0xff
				dots[(i+j*width)*4+1] = 0xff
				dots[(i+j*width)*4+2] = 0xff
				dots[(i+j*width)*4+3] = 0xff
			}
		}
	}
	return dots
}

type game struct {
	scale float64
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *game) Update() error {
	fullscreen := ebiten.IsFullscreen()

	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		switch g.scale {
		case 0.5:
			g.scale = 1
		case 1:
			g.scale = 1.5
		case 1.5:
			g.scale = 2
		case 2:
			g.scale = 0.5
		default:
			panic("not reached")
		}
		ebiten.SetWindowSize(int(screenWidth*g.scale), int(screenHeight*g.scale))
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		fullscreen = !fullscreen
		ebiten.SetFullscreen(fullscreen)
	}
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.ReplacePixels(getDots(screen.Size()))
}

func main() {
	g := &game{
		scale: initScreenScale,
	}
	ebiten.SetWindowSize(screenWidth*initScreenScale, screenHeight*initScreenScale)
	ebiten.SetWindowTitle("Moire (Ebiten Demo)")
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
