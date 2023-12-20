package assets

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
	AnimFPS            = 15
	DefaultSoundVolume = 0.3 // Default sound volume
	TileSize           = 32
	SpriteSize         = 24
)

var (
	IconImage *ebiten.Image
	//go:embed "Icon.png"
	IconImage_png []byte

	spriteSheetImage *ebiten.Image
	//go:embed SpriteSheet.png
	spriteSheet_png []byte

	tilesImage *ebiten.Image
	//go:embed Tiles.png
	tiles_png []byte

	TitleTextImage *ebiten.Image
	//go:embed "TitleText.png"
	titleTextImage_png []byte

	TitleImage *ebiten.Image
	//go:embed "Title.png"
	titleImage_png []byte

	// Main FONT
	// 	The font must be used at the intended point size for optimal results.
	//  Arcadepix at at 10 points or multiplyed by whole number eg. 20, 30, 40

	//go:embed ARCADEPI.TTF
	arcadeFont_ttf []byte
	ArcadeFont     font.Face

	// AUDIO
	SoundVolume        float64 // command line argument
	audioContext       *audio.Context
	AudioPlayerPlaying *audio.Player // points to currently playing audioplayer

	//go:embed audio/1_Stage.mp3
	Theme1Stage_mp3   []byte
	Theme1StagePlayer *audio.Player

	//go:embed audio/2_Stage.mp3
	Theme2Stage_mp3   []byte
	Theme2StageStream *mp3.Stream
	Theme2StagePlayer *audio.Player

	//go:embed audio/3_Boss_Fight.mp3
	ThemeBossFight_mp3   []byte
	ThemeBossFightPlayer *audio.Player

	//go:embed audio/4_Stage_Select.mp3
	ThemeStageSelect_mp3   []byte
	ThemeStageSelectPlayer *audio.Player

	// P1 fire
	//go:embed audio/sfx_wpn_laser1.wav
	sfx_wpn_laser1_wav   []byte
	Sfx_wpn_laser1Player *audio.Player

	// P1 Die
	//go:embed audio/sfx_exp_odd3.wav
	sfx_exp_odd3_wav   []byte
	Sfx_exp_odd3Player *audio.Player

	// P2 Fire
	// P2 Die

	// Enemy Fire
	//go:embed audio/sfx_wpn_laser8.wav
	sfx_wpn_laser8_wav   []byte
	Sfx_wpn_laser8Player *audio.Player

	// Vulcano fire
	//go:embed audio/sfx_exp_short_hard2.wav
	sfx_exp_short_hard2_wav   []byte
	Sfx_exp_short_hard2Player *audio.Player

	// Explosion
	//go:embed audio/sfx_exp_odd1.wav
	sfx_exp_odd1_wav   []byte
	Sfx_exp_odd1Player *audio.Player

	// EnemyBaloon Die
	//go:embed audio/sfx_exp_double3.wav
	sfx_exp_double3_wav   []byte
	Sfx_exp_double3Player *audio.Player
)

// colors
var ColorYellow color.RGBA
var ColorRed color.RGBA
var ColorBlack color.RGBA
var ColorWhite color.RGBA

// sprite animations
var (
	AnimSuperFlyingMan    = []int{0, 1, 2, 3, 4, 4, 3, 2, 1}                                              // SuperFlyingMan
	AnimSuperFlyingManPew = []int{5, 6, 7, 8, 7, 6, 5}                                                    // SuperFlyingManPew
	AnimPigPew            = []int{14, 15, 16, 15}                                                         // PigPew
	AnimPig               = []int{10, 11, 12, 13, 13, 12, 11, 10}                                         // Pig
	AnimSuperFlyingManDie = []int{17, 18, 19, 19, 20, 20, 21, 21, 22, 23, 24, 24, 25, 25, 25, 25}         // SuperFlyingManDie
	AnimPigDie            = []int{26, 27, 28, 28, 29, 29, 30, 30, 31, 31, 22, 23, 24, 24, 25, 25, 25, 25} // PigDie
	// animEnemyBullet          = []int{46, 47}                                                                 // animEnemyBullet
	AnimEnemyBullet1    = []int{58, 59, 60, 61, 62, 63}                 // EnemyBullet1
	AnimEnemyBullet2    = []int{64, 65}                                 // EnemyBullet2
	AnimEnemyBaloon     = []int{32, 33, 34, 35, 36, 36, 35, 34, 33, 32} // EnemyBaloon
	AnimEnemyBaloonFPS  = 5.0
	AnimEnemyBaloonDie  = []int{37, 38, 38, 39, 39, 40, 40, 46, 47, 48, 49, 49} // EnemyBaloonDie
	AnimExplosion       = []int{46, 47, 48, 49, 48, 47, 46}                     // Explosion
	AnimEnemyFlyingMan1 = []int{41, 42, 43, 44, 45, 45, 43, 42}                 // EnemyFlyingMan1
	AnimEnemyFlyingMan2 = []int{66, 67, 68, 69, 70, 70, 68, 67}                 // EnemyFlyingMan2
	AnimEnemyThing      = []int{50, 51, 52, 53, 52, 51, 50}                     // EnemyThing
	AnimEnemyCat        = []int{54, 55, 56, 57, 56, 55, 54}                     // EnemyCat
	AnimEnemyVulcano    = []int{200, 71, 72, 73, 72, 71}                        // EnemyVulcano
)

func InitAssets() {

	tt, err := opentype.Parse(arcadeFont_ttf)
	if err != nil {
		log.Fatal(err)
	}
	ArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    10,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	// GRAPHIC
	// Decode Title logo image
	img, _, err := image.Decode(bytes.NewReader(titleTextImage_png))
	if err != nil {
		log.Fatal(err)
	}
	TitleTextImage = ebiten.NewImageFromImage(img)

	// Decode Title logo texts
	img, _, err = image.Decode(bytes.NewReader(titleImage_png))
	if err != nil {
		log.Fatal(err)
	}
	TitleImage = ebiten.NewImageFromImage(img)

	// Decode sprite sheet from the image file's byte slice.
	img, _, err = image.Decode(bytes.NewReader(spriteSheet_png))
	if err != nil {
		log.Fatal(err)
	}
	spriteSheetImage = ebiten.NewImageFromImage(img)

	// Decode map tiles from the image file's byte slice.
	img, _, err = image.Decode(bytes.NewReader(tiles_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)

	// Colors
	ColorYellow = color.RGBA{0xf6, 0xf4, 0x0d, 0xff}
	ColorRed = color.RGBA{0xf6, 0x00, 0x00, 0xff}
	ColorBlack = color.RGBA{0x00, 0x00, 0x00, 0xff}
	ColorWhite = color.RGBA{0xff, 0xff, 0xff, 0xff}
}

func InitSounds() {
	// AUDIO
	audioContext = audio.NewContext(44100)

	Theme1StagePlayer = loadMp3Sound(Theme1Stage_mp3, true)
	Theme1StagePlayer.SetVolume(SoundVolume)

	Theme2StagePlayer = loadMp3Sound(Theme2Stage_mp3, true)
	Theme2StagePlayer.SetVolume(SoundVolume)

	ThemeBossFightPlayer = loadMp3Sound(ThemeBossFight_mp3, true)
	ThemeBossFightPlayer.SetVolume(SoundVolume)

	ThemeStageSelectPlayer = loadMp3Sound(ThemeStageSelect_mp3, true)
	ThemeStageSelectPlayer.SetVolume(SoundVolume)

	Sfx_wpn_laser1Player = loadWavSound(sfx_wpn_laser1_wav, false)
	Sfx_wpn_laser1Player.SetVolume(SoundVolume)

	Sfx_exp_odd3Player = loadWavSound(sfx_exp_odd3_wav, false)
	Sfx_exp_odd3Player.SetVolume(SoundVolume)

	Sfx_exp_odd1Player = loadWavSound(sfx_exp_odd1_wav, false)
	Sfx_exp_odd1Player.SetVolume(SoundVolume)

	Sfx_exp_double3Player = loadWavSound(sfx_exp_double3_wav, false)
	Sfx_exp_odd1Player.SetVolume(SoundVolume)

	Sfx_wpn_laser8Player = loadWavSound(sfx_wpn_laser8_wav, false)
	Sfx_wpn_laser8Player.SetVolume(SoundVolume)

	Sfx_exp_short_hard2Player = loadWavSound(sfx_exp_short_hard2_wav, false)
	Sfx_exp_short_hard2Player.SetVolume(SoundVolume)

}

func loadMp3Sound(contents []byte, loop bool) *audio.Player {
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

func loadWavSound(contents []byte, loop bool) *audio.Player {
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
	if AudioPlayerPlaying != nil {
		AudioPlayerPlaying.Pause()
		AudioPlayerPlaying.Rewind()
	}
	AudioPlayerPlaying = newPlayer
	AudioPlayerPlaying.Play()
}

func StopAudioPlayer() {
	if AudioPlayerPlaying != nil {
		AudioPlayerPlaying.Pause()
		AudioPlayerPlaying.Rewind()
	}
}

func GetSprite(frameIndex int) image.Image {
	frameOffset := frameIndex * SpriteSize
	return spriteSheetImage.SubImage(image.Rect(frameOffset, 0, frameOffset+SpriteSize, SpriteSize))
}
