package main

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	ANIM_FPS     = 15
	MUSIC_VOLUME = 0.3
)

var (
	iconImage *ebiten.Image
	//go:embed "assets/Icon.png"
	iconImage_png []byte

	SpriteSheetImage *ebiten.Image
	//go:embed assets/SpriteSheet.png
	SpriteSheet_png []byte

	tilesImage *ebiten.Image
	//go:embed assets/Tiles.png
	Tiles_png []byte

	titleTextImage *ebiten.Image
	//go:embed "assets/TitleText.png"
	TitleTextImage_png []byte

	titleImage *ebiten.Image
	//go:embed "assets/Title.png"
	TitleImage_png []byte

	// Main FONT
	// 	The font must be used at the intended point size for optimal results.
	//  Arcadepix at at 10 points or multiplyed by whole number eg. 20, 30, 40

	//go:embed assets/ARCADEPI.TTF
	arcadeFont_ttf []byte
	arcadeFont     font.Face

	// AUDIO
	audioContext *audio.Context
	audioPlayer  *audio.Player

	//go:embed assets/1_Stage.mp3
	audio1StageTheme_mp3 []byte
	audio1StageTheme     *mp3.Stream

	//go:embed assets/2_Stage.mp3
	audio2StageTheme_mp3 []byte
	audio2StageTheme     *mp3.Stream

	//go:embed assets/3_Boss_Fight.mp3
	audioBossFightTheme_mp3 []byte
	audioBossFightTheme     *mp3.Stream

	//go:embed assets/4_Stage_Select.mp3
	audioStageSelectTheme_mp3 []byte
	audioStageSelectTheme     *mp3.Stream
)

// colors
var ColorYellow color.RGBA

// sprite animations
var (
	animSuperFlyingMan    = []int{0, 1, 2, 3, 4, 4, 3, 2, 1}                                              // SuperFlyingMan
	animSuperFlyingManPew = []int{5, 6, 7, 8, 7, 6, 5}                                                    // SuperFlyingManPew
	animPigPew            = []int{14, 15, 16, 15}                                                         // PigPew
	animPig               = []int{10, 11, 12, 13, 13, 12, 11, 10}                                         // Pig
	animSuperFlyingManDie = []int{17, 18, 19, 19, 20, 20, 21, 21, 22, 23, 24, 24, 25, 25, 25, 25}         // SuperFlyingManDie
	animPigDie            = []int{26, 27, 28, 28, 29, 29, 30, 30, 31, 31, 22, 23, 24, 24, 25, 25, 25, 25} // PigDie
	animEnemyPew          = []int{46, 47}                                                                 // EnemyPew
	animEnemyBaloon       = []int{32, 33, 34, 35, 36, 36, 35, 34, 33, 32}                                 // EnemyBaloon
	animEnemyBaloonFPS    = 5.0
	animEnemyBaloonDie    = []int{37, 38, 38, 39, 39, 40, 40, 46, 47, 48, 49, 49} // EnemyBaloonDie
	animExplosion         = []int{46, 47, 48, 49, 48, 47, 46}                     // Explosion
	animEnemyFlyingMan1   = []int{41, 42, 43, 44, 45, 45, 43, 42}                 // EnemyFlyingMan1
	animEnemyThing        = []int{50, 51, 52, 53, 52, 51, 50}                     // EnemyThing
	animEnemyCat          = []int{54, 55, 56, 57, 56, 55, 54}                     // EnemyCat
)

func initAssets() {
	// AUDIO
	audioContext = audio.NewContext(44100)
	var err error
	audio1StageTheme, err = mp3.DecodeWithoutResampling(bytes.NewReader(audio1StageTheme_mp3))
	if err != nil {
		log.Fatal(err)
	}
	audio2StageTheme, err = mp3.DecodeWithoutResampling(bytes.NewReader(audio2StageTheme_mp3))
	if err != nil {
		log.Fatal(err)
	}
	audioBossFightTheme, err = mp3.DecodeWithoutResampling(bytes.NewReader(audioBossFightTheme_mp3))
	if err != nil {
		log.Fatal(err)
	}
	audioStageSelectTheme, err = mp3.DecodeWithoutResampling(bytes.NewReader(audioStageSelectTheme_mp3))
	if err != nil {
		log.Fatal(err)
	}

	tt, err := opentype.Parse(arcadeFont_ttf)
	if err != nil {
		log.Fatal(err)
	}
	arcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    10,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	// GRAPHIC
	// Decode Title logo image
	img, _, err := image.Decode(bytes.NewReader(TitleTextImage_png))
	if err != nil {
		log.Fatal(err)
	}
	titleTextImage = ebiten.NewImageFromImage(img)

	// Decode Title logo texts
	img, _, err = image.Decode(bytes.NewReader(TitleImage_png))
	if err != nil {
		log.Fatal(err)
	}
	titleImage = ebiten.NewImageFromImage(img)

	// Decode sprite sheet from the image file's byte slice.
	img, _, err = image.Decode(bytes.NewReader(SpriteSheet_png))
	if err != nil {
		log.Fatal(err)
	}
	SpriteSheetImage = ebiten.NewImageFromImage(img)

	// Decode map tiles from the image file's byte slice.
	img, _, err = image.Decode(bytes.NewReader(Tiles_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)

	// Colors
	ColorYellow = color.RGBA{0xf6, 0xf4, 0x0d, 0xff}
}

func getSprite(frameIndex int) image.Image {
	frameOffset := frameIndex * spriteSize
	return SpriteSheetImage.SubImage(image.Rect(frameOffset, 0, frameOffset+spriteSize, spriteSize))
}
