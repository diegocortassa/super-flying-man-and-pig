package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/diegocortassa/super-flying-man-and-pig/assets"
	"github.com/diegocortassa/super-flying-man-and-pig/globals"
	"github.com/diegocortassa/super-flying-man-and-pig/version"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	const zoom = 2
	var startPosition int

	showVersion := flag.Bool("version", false, "Show version information")
	fullScreen := flag.Bool("fullscreen", false, "run in fullscreen mode")
	flag.BoolVar(&globals.MameKeys, "mamekeys", false, "Use MAME compatible key mapping")
	flagCRT := flag.Bool("crt", false, "enable the CRT simulation")
	flag.BoolVar(&globals.Debug, "debug", false, "enable debug")
	startAt := flag.String("startat", "Beach", "Start at coordinates Beach, Clouds, Desert, Forest or Castle")
	flag.Float64Var(&assets.SoundVolume, "soundvolume", assets.DefaultSoundVolume*10, "Set sound volume 0 to 10")
	startLives := flag.Int("startlives", 3, "Set lives number (default 3)")
	flag.Parse()

	if *showVersion {
		fmt.Printf("Version: %s\n", version.String())
		os.Exit(0)
	}

	// Volume must be a float64 from 0 to 1
	if assets.SoundVolume > 10 {
		assets.SoundVolume = 10
	}
	assets.SoundVolume /= float64(10)

	// initializations before creating the game
	assets.InitAssets()
	assets.InitSounds()

	fmt.Println("Super flying man and Pig - v 0.1")
	switch *startAt {
	case "Beach":
		fmt.Println("Starting at Beach")
		startPosition = globals.Beach
	case "Clouds":
		fmt.Println("Starting at Clouds")
		startPosition = globals.Clouds
	case "Desert":
		fmt.Println("Starting at Desert")
		startPosition = globals.Desert
	case "Badlands":
		fmt.Println("Starting at Badlands")
		startPosition = globals.Badlands
	case "Forest":
		fmt.Println("Starting at Forest")
		startPosition = globals.Forest
	case "Castle":
		fmt.Println("Starting at Castle")
		startPosition = globals.Castle
	default:
		fmt.Printf("Starting at Beach")
		startPosition = globals.Beach
	}

	g := NewGame(*flagCRT, startPosition, *startLives)

	ebiten.SetWindowSize(globals.ScreenWidth*zoom, globals.ScreenHeight*zoom)
	ebiten.SetWindowTitle("Super flying man and Pig")
	if *fullScreen {
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
