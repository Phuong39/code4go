package main

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	raudio "github.com/hajimehoshi/ebiten/v2/examples/resources/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 640
	screenHeight = 480
	sampleRate   = 44100
)

type Game struct {
	audioContext *audio.Context
	audioPlayer  *audio.Player
}

var g Game

func init() {
	var err error
	// Initialize audio context.
	g.audioContext = audio.NewContext(sampleRate)

	// In this example, embedded resource "Jab_wav" is used.
	//
	// If you want to use a wav file, open this and pass the file stream to wav.Decode.
	// Note that file's Close() should not be closed here
	// since audio.Player manages stream state.
	//
	//     f, err := os.Open("jab.wav")
	//     if err != nil {
	//         return err
	//     }
	//
	//     d, err := wav.Decode(g.audioContext, f)
	//     ...

	// Decode wav-formatted data and retrieve decoded PCM stream.
	d, err := wav.Decode(g.audioContext, bytes.NewReader(raudio.Jab_wav))
	if err != nil {
		log.Fatal(err)
	}

	// Create an audio.Player that has one stream.
	g.audioPlayer, err = audio.NewPlayer(g.audioContext, d)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		// As audioPlayer has one stream and remembers the playing position,
		// rewinding is needed before playing when reusing audioPlayer.
		g.audioPlayer.Rewind()
		g.audioPlayer.Play()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.audioPlayer.IsPlaying() {
		ebitenutil.DebugPrint(screen, "Bump!")
	} else {
		ebitenutil.DebugPrint(screen, "Press P to play the wav")
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("WAV (Ebiten Demo)")
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
