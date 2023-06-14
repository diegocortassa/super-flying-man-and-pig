package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
	"os"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
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

type Mode int

const (
	ModeTitle Mode = iota
	ModeGame
	ModeAttract
	ModeHiscore
	ModeGameOver
)

var (
	debug bool

	fullscreen bool
	flagCRT    bool

	SpriteSheetImage *ebiten.Image
	//go:embed assets/SpriteSheet.png
	SpriteSheet_png []byte

	tilesImage *ebiten.Image
	//go:embed assets/Tiles.png
	Tiles_png []byte

	// EMBEDDED data
	//go:embed assets/ARCADEPI.TTF
	// 	The fonts must be used at the intended point size for optimal results.
	//  Arcadepix at at 10 points or multiplyed by whole number eg. 20, 30, 40
	arcadeFont_ttf []byte
	arcadeFont     font.Face

	//go:embed crt.go
	crtGo []byte

	audioContext *audio.Context
	//go:embed assets/2_Stage_lo_hi.mp3
	audioTheme_mp3 []byte

	iconImage *ebiten.Image
	//go:embed "assets/Icon.png"
	iconImage_png []byte

	// END EMBEDDED data

	lastUpdate time.Time
)

// Game controls overall gameplay.
type Game struct {
	crtShader *ebiten.Shader
	mode      Mode
	gameMap   []int
	position  int

	touchIDs   []ebiten.TouchID
	gamepadIDs []ebiten.GamepadID

	players              []*Entity
	playerOneBullettPool []*Entity
	playerTwoBullettPool []*Entity
	enemies              []*Entity
	enemiesBullettPool   []*Entity
}

func init() {

	rand.Seed(time.Now().UnixNano())

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

	// Load and play game theme
	audioContext = audio.NewContext(22050)

	soundTheme, err := mp3.DecodeWithoutResampling(bytes.NewReader(audioTheme_mp3))
	if err != nil {
		log.Fatal(err)
	}
	soundThemeSource := audio.NewInfiniteLoop(soundTheme, soundTheme.Length())
	audioThemePlayer, err := audio.NewPlayer(audioContext, soundThemeSource)
	if err != nil {
		log.Fatal(err)
	}
	audioThemePlayer.SetVolume(0.02)
	audioThemePlayer.Play()

	// Decode map tiles from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(Tiles_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)
}

func (g *Game) init() {

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
	g.playerOneBullettPool = initBulletPool(animSuperFlyingManPew)
	log.Println("main", len(g.playerOneBullettPool))
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

	g.playerTwoBullettPool = initBulletPool(animPigPew)
	shooter = NewKeyboardShooter(
		playerTwo,
		ebiten.KeyQ,
		g.playerTwoBullettPool,
		time.Millisecond*250,
	)
	playerTwo.addComponent(shooter)

	g.players = append(g.players, playerTwo)
	// #endregion player two

	// Enemies
	var enemy *Entity
	g.enemiesBullettPool = initBulletPool(animEnemyPew)

	enemy = newEntity(
		SpriteSheetImage,
		animEnemyThing,
		Vector{x: 30, y: 0},
	)
	cmover := NewConstantMover(enemy, Vector{x: 0.2, y: 1})
	enemy.addComponent(cmover)
	g.enemies = append(g.enemies, enemy)

	enemy = newEntity(
		SpriteSheetImage,
		animEnemyFlyingMan1,
		Vector{x: 0, y: 0},
	)
	cmover = NewConstantMover(enemy, Vector{x: 0.3, y: 0.3})
	enemy.addComponent(cmover)
	cshooter := NewConstantShooter(
		enemy,
		time.Millisecond*250,
		g.enemiesBullettPool,
	)
	enemy.addComponent(cshooter)
	g.enemies = append(g.enemies, enemy)

	// ne1 := newEnemyFlyingMan1(150, 0)
	// g.enemies = append(g.enemies, &ne1)
	// ne2 := newEnemyFlyingMan1(180, 0)
	// g.enemies = append(g.enemies, &ne2)
	// ne3 := newEnemyFlyingMan1(200, 0)
	// g.enemies = append(g.enemies, &ne3)
}

func (g *Game) reset() {
	g.position = 0
	g.enemies = nil
	g.init()
}

func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		log.Println("Bye")
		os.Exit(0)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		if ebiten.IsFullscreen() {
			ebiten.SetFullscreen(false)
		} else {
			ebiten.SetFullscreen(true)
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.reset()
	}

	for _, player := range g.players {
		player.Update(g)
	}

	for _, playerOneBullet := range g.playerOneBullettPool {
		playerOneBullet.Update(g)
	}

	for _, playerTwoBullet := range g.playerTwoBullettPool {
		playerTwoBullet.Update(g)
	}

	for _, entity := range g.enemies {
		entity.Update(g)
	}

	for _, enemyBullet := range g.enemiesBullettPool {
		enemyBullet.Update(g)
	}

	// p1 := g.entities[0]
	// p2 := g.entities[1]
	// if g.isCollision(p2.position.x, p2.position.y, spriteSize, spriteSize, p1.position.x, p1.position.y, spriteSize, spriteSize) {
	// 	log.Println("Game Over")
	// 	// ebiten.SetRunnableOnUnfocused(true)
	// 	// return nil
	// }

	if time.Since(lastUpdate) > scrollSpeed*time.Millisecond {
		lastUpdate = time.Now()
		g.position += 1 // pixel lines per scroll tick
		// tiles in a screen tilesScreenWidth*tilesScreenHeight
		// as g.position is the low line pixel index
		// we reset to 0 when we have only one screen left of tiles
		if (g.position/tileSize)*tilesScreenWidth >= len(gameMap)-(tilesScreenWidth*tilesScreenHeight) {
			g.position = 0
		}
		// g.position += 8
		// if g.position >= len(gameMap)-8*8 {
		// 	g.position = 0
		// }
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	activeBuls := 0
	activeEntities := 0
	tileIndex := 0
	// Draw world window
	rowPosition := (g.position / tileSize) * tilesScreenWidth
	screenPosition := (g.position % tileSize)
	// for rowIndex := tilesScreenHeight - 1; rowIndex >= 0; rowIndex-- { // use this to show scrolling trick
	for rowIndex := tilesScreenHeight; rowIndex >= 0; rowIndex-- {
		// fmt.Println("WPos: ", g.position, "SPos:", screenPosition, "RPos:", rowPosition, "Row:", rowIndex, "-")
		for columnIndex := 0; columnIndex < tilesScreenWidth; columnIndex++ {
			// fmt.Print("col:", columnIndex, " ")

			op := &ebiten.DrawImageOptions{}
			// op.GeoM.Translate(float64(tileSize*columnIndex), float64(tileSize*rowIndex+screenPosition)) // use this to show scrolling trick
			op.GeoM.Translate(float64(tileSize*columnIndex), float64(tileSize*rowIndex+screenPosition-tileSize))
			tileIndex = gameMap[rowPosition]
			rowPosition++
			screen.DrawImage(getTile(tileIndex).(*ebiten.Image), op)
		}
	}

	for _, player := range g.players {
		player.Draw(screen)
	}

	for _, playerOneBullet := range g.playerOneBullettPool {
		playerOneBullet.Draw(screen)
		if playerOneBullet.active {
			activeBuls += 1
		}
	}

	for _, playerTwoBullet := range g.playerTwoBullettPool {
		playerTwoBullet.Draw(screen)
		if playerTwoBullet.active {
			activeBuls += 1
		}
	}

	for _, entity := range g.enemies {
		entity.Draw(screen)
		if entity.active {
			activeEntities += 1
		}
	}

	for _, enemyBullet := range g.enemiesBullettPool {
		enemyBullet.Draw(screen)
		if enemyBullet.active {
			activeBuls += 1
		}
	}

	// Draw Score/Lives
	msg := fmt.Sprintf("1UP\nPRESS FIRE")
	text.Draw(screen, msg, arcadeFont, 5, 20, color.White)
	msg = fmt.Sprintf("HI-SCORE\n  12200")
	text.Draw(screen, msg, arcadeFont, 90, 20, color.White)
	msg = fmt.Sprintf("2UP\nPRESS FIRE")
	text.Draw(screen, msg, arcadeFont, 170, 20, color.White)

	// Draw debug data
	if debug {
		nl := "\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n"
		ebitenutil.DebugPrint(screen, fmt.Sprintf("%vTPS: %0.2f FPS: %0.2f \n Pos:%v Ent: %v Blt: %v", nl, ebiten.ActualTPS(), ebiten.ActualFPS(), g.position, activeEntities, activeBuls))
	}
}

type GameWithCRTEffect struct {
	ebiten.Game

	crtShader *ebiten.Shader
}

func (g *GameWithCRTEffect) DrawFinalScreen(screen ebiten.FinalScreen, offscreen *ebiten.Image, geoM ebiten.GeoM) {
	fmt.Println("DrawFinalScreen")

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

func (g *Game) isCollision(x1, y1, w1, h1, x2, y2, w2, h2 float64) bool {
	return x1 < x2+w2 &&
		x1+w1 > x2 &&
		y1 < y2+h2 &&
		y1+h1 > y2
}

func main() {
	flag.BoolVar(&fullscreen, "fullscreen", false, "run in fullscreen mode")
	flag.BoolVar(&flagCRT, "crt", false, "enable the CRT simulation")
	flag.Parse()

	debug = true

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
