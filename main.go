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

	"github.com/hajimehoshi/ebiten/v2"
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

type State int // /possible scenes/states the game can be

const (
	StateInit State = iota
	StateTitle
	StateGame
	StateAttract
	StateHiscores
	StateCredits
	StateInsertName
	StateGameOver
	StateGameEnd
)

var (
	debug       bool    // command line flag
	fullScreen  bool    // command line flag
	flagCRT     bool    // command line flag
	MameKeys    bool    // command line flag
	SoundVolume float64 // command line argument

	//go:embed crt.go
	crtGo []byte
)

// Game controls overall gameplay.
type Game struct {
	crtShader *ebiten.Shader
	gameMap   []int // Tiled world map

	paused   bool // vertical position on screen map
	position int  // vertical position on screen map
	hiScores int

	touchIDs   []ebiten.TouchID
	gamepadIDs []ebiten.GamepadID

	playerOne            *Entity
	playerTwo            *Entity
	playerOneBullettPool []*Entity
	playerTwoBullettPool []*Entity
	enemies              []*Entity
	enemiesBullettPool   []*Entity

	CurrentState        State     // the state/scene the game is currently playing
	PreviousState       State     // the previous state/scene the game was
	lastStateTransition time.Time // last time the state/scene was changed
	lastEvent           time.Time

	lastSpawn time.Time // last time an enemy was spawned
	spawnHead int       // index on spawn script
}

func (g *Game) init() {
	// initializations before running the game
	g.CurrentState = StateInit
	g.hiScores = 1230
	g.PreviousState = StateInit
	g.lastStateTransition = time.Now()
	g.lastEvent = time.Now()
	g.reset()
	g.resetPlayerOne()
	g.resetPlayerTwo()
}

func (g *Game) reset() {
	g.position = 0
	// Enemies
	g.enemies = nil
	// Enemies Bullets
	g.enemiesBullettPool = initBulletPool("EnemyBullet", typeEnemyBullet, animEnemyBullet1, 20, 10, Vector{0, 2}, Box{9, 9, 6, 6})
	g.lastSpawn = time.Now()
	g.spawnHead = 0
}

func (g *Game) Update() error {

	g.UpdateDirector()
	if g.paused {
		return nil
	}

	switch g.CurrentState {
	case StateTitle:
		g.UpdateTileState()
	case StateAttract:
		g.UpdateAttractState()
	case StateGame:
		g.UpdateGameState()
	case StateGameOver:
		g.UpdateGameOverState()
	case StateInsertName:
		g.UpdateGameState()
	case StateHiscores:
		g.UpdateHiscoresState()
	case StateCredits:
		g.UpdateCreditsState()
	case StateGameEnd:
		g.UpdateGameState()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	switch g.CurrentState {
	case StateTitle:
		g.DrawTitleState(screen)
	case StateAttract:
		g.DrawAttractState(screen)
	case StateGame:
		g.DrawGameState(screen)
	case StateGameOver:
		g.DrawGameOverState(screen)
	case StateInsertName:
		g.DrawGameState(screen)
	case StateHiscores:
		g.DrawHiscoreState(screen)
	case StateCredits:
		g.DrawCreditsState(screen)
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

	os := offscreen.Bounds().Size()
	op := &ebiten.DrawRectShaderOptions{}
	op.Images[0] = offscreen
	op.GeoM = geoM
	screen.DrawRectShader(os.X, os.Y, g.crtShader, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// return screenWidth + 30, screenHeight + 30
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

	flag.BoolVar(&fullScreen, "fullscreen", false, "run in fullscreen mode")
	flag.BoolVar(&MameKeys, "mamekeys", false, "Use MAME compatible key mapping")
	flag.BoolVar(&flagCRT, "crt", false, "enable the CRT simulation")
	flag.BoolVar(&debug, "debug", false, "enable debug")
	flag.Float64Var(&SoundVolume, "soundvolume", SOUND_VOLUME*10, "Set sound volume 0 to 10")
	flag.Parse()

	// Volume must be a float64 from 0 to 1
	if SoundVolume > 10 {
		SoundVolume = 10
	}
	SoundVolume /= float64(10)

	// initializations before creating the game
	rand.Seed(time.Now().UnixNano())
	initAssets()
	initSounds()

	g := NewGame(flagCRT)

	ebiten.SetWindowSize(screenWidth*zoom, screenHeight*zoom)
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
	iconImg, _, _ := image.Decode(bytes.NewReader(iconImage_png))
	iconImgS := []image.Image{iconImg}
	ebiten.SetWindowIcon(iconImgS)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
