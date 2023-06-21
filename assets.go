package main

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	"io"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	ANIM_FPS     = 15
	SOUND_VOLUME = 0.3 // Default sound volume
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
	audioContext       *audio.Context
	audioPlayerPlaying *audio.Player // points to currently playing audioplayer

	//go:embed assets/1_Stage.mp3
	Theme1Stage_mp3   []byte
	Theme1StagePlayer *audio.Player

	//go:embed assets/2_Stage.mp3
	Theme2Stage_mp3   []byte
	Theme2StageStream *mp3.Stream
	Theme2StagePlayer *audio.Player

	//go:embed assets/3_Boss_Fight.mp3
	ThemeBossFight_mp3   []byte
	ThemeBossFightPlayer *audio.Player

	//go:embed assets/4_Stage_Select.mp3
	ThemeStageSelect_mp3   []byte
	ThemeStageSelectPlayer *audio.Player

	// P1 fire
	//go:embed assets/sfx_wpn_laser1.wav
	sfx_wpn_laser1_wav   []byte
	sfx_wpn_laser1Player *audio.Player

	// P1 Die
	//go:embed assets/sfx_exp_odd3.wav
	sfx_exp_odd3_wav   []byte
	sfx_exp_odd3Player *audio.Player

	// P2 Fire
	// P2 Die

	// Enemy Fire
	//go:embed assets/sfx_wpn_laser8.wav
	sfx_wpn_laser8_wav   []byte
	sfx_wpn_laser8Player *audio.Player

	// Explosion
	//go:embed assets/sfx_exp_odd1.wav
	sfx_exp_odd1_wav   []byte
	sfx_exp_odd1Player *audio.Player

	// EnemyBaloon Die
	//go:embed assets/sfx_exp_double3.wav
	sfx_exp_double3_wav   []byte
	sfx_exp_double3Player *audio.Player
)

// colors
var ColorYellow color.RGBA
var ColorRed color.RGBA

// sprite animations
var (
	animSuperFlyingMan    = []int{0, 1, 2, 3, 4, 4, 3, 2, 1}                                              // SuperFlyingMan
	animSuperFlyingManPew = []int{5, 6, 7, 8, 7, 6, 5}                                                    // SuperFlyingManPew
	animPigPew            = []int{14, 15, 16, 15}                                                         // PigPew
	animPig               = []int{10, 11, 12, 13, 13, 12, 11, 10}                                         // Pig
	animSuperFlyingManDie = []int{17, 18, 19, 19, 20, 20, 21, 21, 22, 23, 24, 24, 25, 25, 25, 25}         // SuperFlyingManDie
	animPigDie            = []int{26, 27, 28, 28, 29, 29, 30, 30, 31, 31, 22, 23, 24, 24, 25, 25, 25, 25} // PigDie
	// animEnemyBullet          = []int{46, 47}                                                                 // animEnemyBullet
	animEnemyBullet1    = []int{58, 59, 60, 61, 62, 63}                 // EnemyBullet1
	animEnemyBullet2    = []int{64, 65}                                 // EnemyBullet2
	animEnemyBaloon     = []int{32, 33, 34, 35, 36, 36, 35, 34, 33, 32} // EnemyBaloon
	animEnemyBaloonFPS  = 5.0
	animEnemyBaloonDie  = []int{37, 38, 38, 39, 39, 40, 40, 46, 47, 48, 49, 49} // EnemyBaloonDie
	animExplosion       = []int{46, 47, 48, 49, 48, 47, 46}                     // Explosion
	animEnemyFlyingMan1 = []int{41, 42, 43, 44, 45, 45, 43, 42}                 // EnemyFlyingMan1
	animEnemyThing      = []int{50, 51, 52, 53, 52, 51, 50}                     // EnemyThing
	animEnemyCat        = []int{54, 55, 56, 57, 56, 55, 54}                     // EnemyCat

)

func initAssets() {

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
	ColorRed = color.RGBA{0xf6, 0x00, 0x00, 0xff}
}

func initSounds() {
	// AUDIO
	audioContext = audio.NewContext(44100)

	Theme1StagePlayer = LoadMp3Sound(Theme1Stage_mp3, true)
	Theme1StagePlayer.SetVolume(SoundVolume)

	Theme2StagePlayer = LoadMp3Sound(Theme2Stage_mp3, true)
	Theme2StagePlayer.SetVolume(SoundVolume)

	ThemeBossFightPlayer = LoadMp3Sound(ThemeBossFight_mp3, true)
	ThemeBossFightPlayer.SetVolume(SoundVolume)

	ThemeStageSelectPlayer = LoadMp3Sound(ThemeStageSelect_mp3, true)
	ThemeStageSelectPlayer.SetVolume(SoundVolume)

	sfx_wpn_laser1Player = LoadWavSound(sfx_wpn_laser1_wav, false)
	sfx_wpn_laser1Player.SetVolume(SoundVolume)

	sfx_exp_odd3Player = LoadWavSound(sfx_exp_odd3_wav, false)
	sfx_exp_odd3Player.SetVolume(SoundVolume)

	sfx_exp_odd1Player = LoadWavSound(sfx_exp_odd1_wav, false)
	sfx_exp_odd1Player.SetVolume(SoundVolume)

	sfx_exp_double3Player = LoadWavSound(sfx_exp_double3_wav, false)
	sfx_exp_odd1Player.SetVolume(SoundVolume)

	sfx_wpn_laser8Player = LoadWavSound(sfx_wpn_laser8_wav, false)
	sfx_wpn_laser8Player.SetVolume(SoundVolume)
}

func LoadMp3Sound(contents []byte, loop bool) *audio.Player {
	stream, err := mp3.DecodeWithoutResampling(bytes.NewReader(contents))
	if err != nil {
		log.Fatal(err)
	}
	var audioSource io.Reader
	if loop {
		audioSource = audio.NewInfiniteLoop(stream, stream.Length()+1)
	} else {
		audioSource = stream
	}
	player, err := audio.NewPlayer(audioContext, audioSource)
	if err != nil {
		log.Fatal(err)
	}

	// Workaround to prevent delays when playing for the first time.
	// player.SetVolume(0)
	// player.Play()
	// player.Pause()
	// player.Rewind()
	// player.SetVolume(1)

	return player
}

func LoadWavSound(contents []byte, loop bool) *audio.Player {
	stream, err := wav.DecodeWithoutResampling(bytes.NewReader(contents))
	if err != nil {
		log.Fatal(err)
	}
	var audioSource io.Reader
	if loop {
		audioSource = audio.NewInfiniteLoop(stream, stream.Length()+1)
	} else {
		audioSource = stream
	}
	player, err := audio.NewPlayer(audioContext, audioSource)
	if err != nil {
		log.Fatal(err)
	}
	return player
}

func PlayTheme(newPlayer *audio.Player) {
	if audioPlayerPlaying != nil {
		audioPlayerPlaying.Pause()
		audioPlayerPlaying.Rewind()
	}
	audioPlayerPlaying = newPlayer
	audioPlayerPlaying.Play()
}

func getSprite(frameIndex int) image.Image {
	frameOffset := frameIndex * spriteSize
	return SpriteSheetImage.SubImage(image.Rect(frameOffset, 0, frameOffset+spriteSize, spriteSize))
}
