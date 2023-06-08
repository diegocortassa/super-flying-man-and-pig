package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
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
	screenWidth     = 256
	screenHeight    = 396
	tilesWidth      = 8
	tilesHeight     = 12
	tileSize        = 32
	zoom            = 2
	playerSpeed     = 2
	playerAnimSpeed = 4
	playerFrameNum  = 5
	playerFrameSize = 24
)

var (

	//go:embed assets/PublicPixel-z84yD.ttf
	arcadeFont_ttf []byte
	arcadeFont     font.Face

	audioContext *audio.Context
	//go:embed assets/moon_patrol.mp3
	audioTheme_mp3 []byte

	tilesImage *ebiten.Image
	//go:embed assets/map/Tiles.png
	Tiles_png []byte
	// position  = 0

	SuperFlyingManImage *ebiten.Image
	//go:embed assets/sprites/SuperFlyingMan.png
	SuperFlyingMan_png []byte

	lastUpdate time.Time
)

type Game struct {
	gameMap       []int
	position      int
	PlayerX       float64
	PlayerY       float64
	PlayerFrame   int
	PlayerCounter int
	PlayerDir     int
}

func init() {

	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 25
	arcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Load and play game theme
	audioContext = audio.NewContext(44100)
	s, err := mp3.DecodeWithoutResampling(bytes.NewReader(audioTheme_mp3))
	if err != nil {
		log.Fatal(err)
	}
	ss := audio.NewInfiniteLoop(s, s.Length())
	player, err := audio.NewPlayer(audioContext, ss)
	if err != nil {
		log.Fatal(err)
	}
	player.Play()

	// Decode map tiles from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(Tiles_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)

	// Decode player one spritesheet from the image file's byte slice.
	img, _, err = image.Decode(bytes.NewReader(SuperFlyingMan_png))
	if err != nil {
		log.Fatal(err)
	}
	SuperFlyingManImage = ebiten.NewImageFromImage(img)
}

func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		panic("Bye")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		if ebiten.IsFullscreen() {
			ebiten.SetFullscreen(false)
		} else {
			ebiten.SetFullscreen(true)
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.PlayerY -= playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.PlayerY += playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.PlayerX -= playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.PlayerX += playerSpeed
	}

	if g.PlayerX < 0 {
		g.PlayerX = 0
	} else if g.PlayerX > screenWidth-playerFrameSize {
		g.PlayerX = screenWidth - playerFrameSize
	}

	if g.PlayerY < 0 {
		g.PlayerY = 0
	} else if g.PlayerY > screenHeight-playerFrameSize {
		g.PlayerY = screenHeight - playerFrameSize
	}

	// compute Player one frame ping pong animation
	g.PlayerCounter++
	if g.PlayerCounter >= playerAnimSpeed {
		g.PlayerFrame += g.PlayerDir
		if g.PlayerFrame >= playerFrameNum-1 || g.PlayerFrame <= 0 {
			g.PlayerDir *= -1 // Reverse direction
		}
		g.PlayerCounter = 0
	}

	if time.Since(lastUpdate) > 800*time.Millisecond {
		lastUpdate = time.Now()
		g.position += 8
		if g.position >= len(gameMap)-8*8 {
			g.position = 0
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	tileIndex := 0

	// Draw world window
	rowPosition := g.position
	fmt.Println("------------------")
	for rowIndex := tilesHeight - 1; rowIndex >= 0; rowIndex-- {
		fmt.Println("Pos: ", g.position, "Row", rowIndex, "-")
		for columnIndex := 0; columnIndex < tilesWidth; columnIndex++ {
			fmt.Print("col:", columnIndex, " ")

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tileSize*columnIndex), float64(tileSize*rowIndex))
			tileIndex = gameMap[rowPosition]
			rowPosition++
			fmt.Print("til:", tileIndex, ", ")
			screen.DrawImage(getTile(tileIndex).(*ebiten.Image), op)
		}
		fmt.Println()
	}

	// Draw Player one
	playerFrameX := g.PlayerFrame * playerFrameSize
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.PlayerX, g.PlayerY)
	screen.DrawImage(SuperFlyingManImage.SubImage(image.Rect(playerFrameX, 0, playerFrameX+playerFrameSize, playerFrameSize)).(*ebiten.Image), opts)

	// Draw debug data
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f FPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS()))

	// Draw info
	msg := fmt.Sprintf("1UP %v", g.position)
	text.Draw(screen, msg, arcadeFont, 170, 15, color.White)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {

	g := &Game{
		PlayerX:   (screenWidth - playerFrameSize) / 2,
		PlayerY:   screenHeight - playerFrameSize,
		PlayerDir: 1, // Start with forward direction
	}

	ebiten.SetWindowSize(screenWidth*zoom, screenHeight*zoom)
	ebiten.SetWindowTitle("Super flying man and Pig")
	// ebiten.SetWindowIcon([]image.Image{loadImage("assets/icon.png")})
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
