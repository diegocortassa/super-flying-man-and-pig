package main

import (
	"bytes"
	_ "embed"
	"flag"
	"image"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/dcortassa/superflyingmanandpig/assets"
	"github.com/dcortassa/superflyingmanandpig/globals"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	zoom = 2
)

var (
	fullScreen bool
)

func main() {

	flag.BoolVar(&fullScreen, "fullscreen", false, "run in fullscreen mode")
	flag.BoolVar(&globals.MameKeys, "mamekeys", false, "Use MAME compatible key mapping")
	flag.BoolVar(&FlagCRT, "crt", false, "enable the CRT simulation")
	flag.BoolVar(&globals.Debug, "debug", false, "enable debug")
	flag.Float64Var(&assets.SoundVolume, "soundvolume", assets.DefaultSoundVolume*10, "Set sound volume 0 to 10")
	flag.Parse()

	// Volume must be a float64 from 0 to 1
	if assets.SoundVolume > 10 {
		assets.SoundVolume = 10
	}
	assets.SoundVolume /= float64(10)

	// initializations before creating the game
	rand.Seed(time.Now().UnixNano())
	assets.InitAssets()
	assets.InitSounds()

	g := NewGame(FlagCRT)

	ebiten.SetWindowSize(globals.ScreenWidth*zoom, globals.ScreenHeight*zoom)
	ebiten.SetWindowTitle("Super flying man and Pig")
	if fullScreen {
		ebiten.SetFullscreen(true)
	}
	// ebiten.SetWindowResizable(true)
	// ebiten.SetScreenTransparent(true)
	// ebiten.SetWindowDecorated(false)
	// Saves GPU if screen is not changed
	// ebiten.SetScreenClearedEveryFrame(false)
	// Decode player one spritesheet from the image file's byte slice.
	// var iconImg []image.Image
	iconImg, _, _ := image.Decode(bytes.NewReader(assets.IconImage_png))
	iconImgS := []image.Image{iconImg}
	ebiten.SetWindowIcon(iconImgS)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
