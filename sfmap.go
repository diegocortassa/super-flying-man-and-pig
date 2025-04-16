package main

import (
	_ "embed"
	"fmt"
	"runtime"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/dcortassa/super-flying-man-and-pig/assets"
	"github.com/dcortassa/super-flying-man-and-pig/debug"
	"github.com/dcortassa/super-flying-man-and-pig/globals"
)

type State int // /possible scenes/states the game can be

const (
	StateInit State = iota
	StateTitle
	StateGame
	StateAttract
	StateHiscores
	StateHiscoresInsert
	StateCredits
	StateInsertName
	StateGameOver
	StateGameEnd
)

var (
	FullScreen bool // command line flag
	FlagCRT    bool // command line flag

	//go:embed crt.go
	crtGo []byte
)

// Game controls overall gameplay.
type Game struct {
	paused                       bool // vertical position on screen map
	Position                     int  // vertical position on screen map
	StartPosition                int
	HiScores                     int
	HiScoresTable                []HiScoresEntry
	HiScoresInsertCursorPosition int

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

	startLives int
}

type HiScoresEntry struct {
	Scores int
	Name   string
}

func (g *Game) init() {
	// initializations before running the game
	g.CurrentState = StateInit
	g.PreviousState = StateInit
	g.lastStateTransition = time.Now()
	g.lastEvent = time.Now()
	g.reset()
	g.resetPlayerOne()
	g.resetPlayerTwo()

	g.HiScoresTable = []HiScoresEntry{
		{Name: "DIE", Scores: 700},
		{Name: "LIV", Scores: 600},
		{Name: "AND", Scores: 500},
		{Name: "NOR", Scores: 400},
		{Name: "MRJ", Scores: 300},
	}
	g.HiScores = g.HiScoresTable[0].Scores
}

func (g *Game) reset() {
	g.Position = g.StartPosition
	// Enemies
	g.enemies = nil
	// Enemies Bullets
	g.enemiesBullettPool = initBulletPool("EnemyBullet", TypeEnemyBullet, assets.AnimEnemyBullet1, 20, 12, Vector{X: 0, Y: 2}, Box{X: 9, Y: 9, W: 6, H: 6})
	g.lastSpawn = time.Now()
	g.spawnHead = 0
	g.HiScoresInsertCursorPosition = 0
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
	case StateHiscoresInsert:
		g.UpdateHiscoresInsertState()
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
	case StateHiscoresInsert:
		g.DrawHiscoreInsertState(screen)
	case StateCredits:
		g.DrawCreditsState(screen)
	case StateGameEnd:
		g.DrawGameState(screen)
	}

	// Draw debug data
	if globals.Debug {

		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		ebitenutil.DebugPrintAt(screen,
			fmt.Sprintf("Alloc: %v TotAlloc: %v Sys: %v NumGC: %v",
				m.Alloc/1024/1024, m.TotalAlloc/1024/1024, m.Sys/1024/1024, m.NumGC),
			2, globals.ScreenHeight-45)

	}

}

type GameWithCRTEffect struct {
	ebiten.Game

	crtShader *ebiten.Shader
}

func (g *GameWithCRTEffect) DrawFinalScreen(screen ebiten.FinalScreen, offscreen *ebiten.Image, geoM ebiten.GeoM) {
	debug.DebugPrintf("DrawFinalScreen")

	if g.crtShader == nil {
		s, err := ebiten.NewShader(crtGo)
		if err != nil {
			panic(fmt.Sprintf("failed to compile the CRT shader: %v", err))
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
	return globals.ScreenWidth, globals.ScreenHeight
}

func NewGame(flagCRT bool, startPosition int, startLives int) ebiten.Game {
	g := &Game{}
	g.StartPosition = startPosition
	g.startLives = startLives
	g.init()
	if flagCRT {
		return &GameWithCRTEffect{Game: g}
	}
	return g
}
