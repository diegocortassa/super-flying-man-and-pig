package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

// 256, 320 # 8x10 32 pixel tiles
// 200, 320 # qvga 6,25x10 32 pixel tiles
// 320, 480 # hvga 10x15 32 pixel tiles
// 256, 384 # SEUCK Amiga 8x12 32 pixel tiles
// Xevious      288x224@60 7x9
// Terra cresta 256x224@60 7x8
// Commando     256x224@60 7x8
// 1942         256x224@59 7x8
// Alcon        296x240@57 7,5x
const (
	screenWidth  = 256
	screenHeight = 384
	// screenHeight      = 416 // use this to show scrolling trick
	tilesScreenWidth  = 8
	tilesScreenHeight = 12
	tileSize          = 32
	zoom              = 2
	spriteSize        = 24
	scrollSpeed       = 30 // speed in 1 pixel per microsecond
)

type State int

const (
	StateTitle State = iota
	StateGame
	StateAttract
	StateHiscore
	StateInsertName
	StateGameOver
	StateGameEnd
)

var (
	debug      bool
	fullscreen bool
	flagCRT    bool

	CurrentState  State
	PreviousState State

	lastUpdate time.Time

	//go:embed crt.go
	crtGo []byte
)

// Game controls overall gameplay.
type Game struct {
	crtShader *ebiten.Shader
	state     State
	gameMap   []int
	position  int

	touchIDs   []ebiten.TouchID
	gamepadIDs []ebiten.GamepadID

	players              []*Entity
	playerOneBullettPool []*Entity
	playerTwoBullettPool []*Entity
	enemies              []*Entity
	enemiesBullettPool   []*Entity
	lastEvent            time.Time
}

func init() {

	rand.Seed(time.Now().UnixNano())

	// Prepare audio
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

	soundThemeSource := audio.NewInfiniteLoop(audio2StageTheme, audio2StageTheme.Length()+1)
	audioPlayer, err = audio.NewPlayer(audioContext, soundThemeSource)
	if err != nil {
		log.Fatal(err)
	}
	audioPlayer.SetVolume(0.05)
	audioPlayer.Play()

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

	// Decode map tiles from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(Tiles_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(TitleTextImage_png))
	if err != nil {
		log.Fatal(err)
	}
	titleTextImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(TitleImage_png))
	if err != nil {
		log.Fatal(err)
	}
	titleImage = ebiten.NewImageFromImage(img)
}

func (g *Game) init() {
	g.state = StateTitle

	g.lastEvent = time.Now()

	// Decode sprite sheet from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(SpriteSheet_png))
	if err != nil {
		log.Fatal(err)
	}
	SpriteSheetImage = ebiten.NewImageFromImage(img)

	// #region player one
	playerOne := newEntity(
		SpriteSheetImage,
		animSuperFlyingMan,
		Vector{x: (screenWidth - spriteSize) / 4, y: screenHeight - spriteSize - 20},
	)
	playerOne.hitBoxes = append(playerOne.hitBoxes, Box{5, 2, 15, 20})
	playerOne.name = "P1"
	playerOne.lives = 3
	mover := NewKeyboardMover(
		playerOne,
		Keybinds{
			Up:        ebiten.KeyArrowUp,
			Down:      ebiten.KeyArrowDown,
			Left:      ebiten.KeyArrowLeft,
			Right:     ebiten.KeyArrowRight,
			Fire:      ebiten.KeyControlLeft,
			Secondary: ebiten.KeyAltLeft,
		},
		Vector{1, 1},
	)
	playerOne.addComponent(mover)
	g.playerOneBullettPool = initBulletPool("P1Bullet", typePlayerOneBullet, animSuperFlyingManPew, 5, Vector{0, -4}, Box{8, 2, 8, 8})
	shooter := NewKeyboardShooter(
		playerOne,
		ebiten.KeyControlLeft,
		g.playerOneBullettPool,
		time.Millisecond*250,
	)
	playerOne.addComponent(shooter)

	g.players = append(g.players, playerOne)
	// #endregion player one

	// #region player two
	playerTwo := newEntity(
		SpriteSheetImage,
		animPig,
		Vector{x: (screenWidth - spriteSize) / 4 * 3, y: screenHeight - spriteSize - 20},
	)
	playerTwo.hitBoxes = append(playerTwo.hitBoxes, Box{5, 2, 15, 20})
	playerTwo.name = "P2"
	playerTwo.lives = 3
	mover = NewKeyboardMover(
		playerTwo,
		Keybinds{
			Up:        ebiten.KeyW,
			Down:      ebiten.KeyS,
			Left:      ebiten.KeyA,
			Right:     ebiten.KeyD,
			Fire:      ebiten.KeyQ,
			Secondary: ebiten.KeyAltLeft,
		},
		Vector{1, 1},
	)
	playerTwo.addComponent(mover)

	g.playerTwoBullettPool = initBulletPool("P2Bullet", typePlayerTwoBullet, animPigPew, 5, Vector{0, -4}, Box{8, 2, 8, 8})
	shooter = NewKeyboardShooter(
		playerTwo,
		ebiten.KeyQ,
		g.playerTwoBullettPool,
		time.Millisecond*250,
	)
	playerTwo.addComponent(shooter)

	g.players = append(g.players, playerTwo)
	// #endregion player two

	// Enemies Bullets
	g.enemiesBullettPool = initBulletPool("EnemyBullet", typeEnemyBullet, animEnemyPew, 10, Vector{0, 0}, Box{8, 8, 8, 8})
}

func (g *Game) reset() {

	g.enemies = nil

	g.state = StateTitle
	g.gameMap = nil
	g.position = 0
	g.players = nil
	g.playerOneBullettPool = nil
	g.playerTwoBullettPool = nil
	g.enemies = nil
	g.enemiesBullettPool = nil
	g.lastEvent = time.Now()
	g.init()
}

func (g *Game) Update() error {

	g.UpdateSequencer()

	switch g.state {
	case StateTitle:
		g.UpdateTileState()
	case StateAttract:
		g.UpdateGameState()
	case StateGame:
		g.UpdateGameState()
	case StateGameOver:
		g.UpdateGameOverState()
	case StateInsertName:
		g.UpdateGameState()
	case StateHiscore:
		g.UpdateGameState()
	case StateGameEnd:
		g.UpdateGameState()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	switch g.state {
	case StateTitle:
		g.DrawTitleState(screen)
	case StateAttract:
		g.DrawGameState(screen)
	case StateGame:
		g.DrawGameState(screen)
	case StateGameOver:
		g.DrawGameOverState(screen)
	case StateInsertName:
		g.DrawGameState(screen)
	case StateHiscore:
		g.DrawGameState(screen)
	case StateGameEnd:
		g.DrawGameState(screen)
	}

}

type GameWithCRTEffect struct {
	ebiten.Game

	crtShader *ebiten.Shader
}

func (g *GameWithCRTEffect) DrawFinalScreen(screen ebiten.FinalScreen, offscreen *ebiten.Image, geoM ebiten.GeoM) {
	DebugPrintf("DrawFinalScreen")

	if g.crtShader == nil {
		s, err := ebiten.NewShader(crtGo)
		if err != nil {
			panic(fmt.Sprintf("failed to compiled the CRT shader: %v", err))
		}
		g.crtShader = s
	}

	doubledOffscreen := ebiten.NewImage(screenWidth*2, screenHeight*2)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(2.0, 2.0)
	doubledOffscreen.DrawImage(offscreen, opts)

	os := doubledOffscreen.Bounds().Size()
	op := &ebiten.DrawRectShaderOptions{}
	op.Images[0] = doubledOffscreen
	op.GeoM = geoM
	screen.DrawRectShader(os.X, os.Y, g.crtShader, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	if flagCRT {
		return screenWidth * 2, screenHeight * 2
	}
	return screenWidth, screenHeight
}

func NewGame(flagCRT bool) ebiten.Game {
	g := &Game{}
	g.init()
	if flagCRT {
		return &GameWithCRTEffect{Game: g}
	}
	return g
}

func main() {
	flag.BoolVar(&fullscreen, "fullscreen", false, "run in fullscreen mode")
	flag.BoolVar(&flagCRT, "crt", false, "enable the CRT simulation")
	flag.BoolVar(&debug, "debug", false, "enable debug")
	flag.Parse()

	// g := &Game{}
	// g.init()
	g := NewGame(flagCRT)

	ebiten.SetWindowSize(screenWidth*zoom, screenHeight*zoom)
	ebiten.SetWindowTitle("Super flying man and Pig")
	if fullscreen {
		ebiten.SetFullscreen(true)
	}
	// ebiten.SetWindowResizable(true)
	// ebiten.SetScreenTransparent(true)
	// ebiten.SetWindowDecorated(false)
	// Saves GPU if screen is not changed
	// ebiten.SetScreenClearedEveryFrame(false)
	// Decode player one spritesheet from the image file's byte slice.
	// var iconImg []image.Image
	iconImg, _, _ := image.Decode(bytes.NewReader(iconImage_png))
	iconImgS := []image.Image{iconImg}
	ebiten.SetWindowIcon(iconImgS)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
