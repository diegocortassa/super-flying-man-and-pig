package main

import (
	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"golang.org/x/image/font"
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
