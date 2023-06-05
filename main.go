package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// 256, 320 # 8x10 32 pixel tiles
// 200, 320 # qvga 6,25x10 32 pixel tiles
// 320, 480 # hvga 10x15 32 pixel tiles
// 480, 800 # wvga800 15x25 32 pixel tiles
// 240, 400 # hwvga800 7.5x12.5 32 pixel tiles
// 256, 384 # SEUCK Amiga 8x12 32 pixel tiles
// Xevious      288x224@60 7x9
// Terra cresta 256x224@60 7x8
// Commando     256x224@60 7x8
// 1942         256x224@59 7x8
// Alcon        296x240@57 7,5x
const (
	screenWidth  = 256
	screenHeight = 256
	tileSize     = 32
	zoom         = 3
)

var (
	tilesImage *ebiten.Image
	//go:embed assets/map/Tiles.png
	Tiles_png []byte
)

func init() {
	// Decode an image from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(Tiles_png))
	// img, _, err := ebitenutil.NewImageFromFile("assets/MapTiles.png")
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)
}

type Game struct {
	gameMap []int
}

func (g *Game) Update() error {

	inpututil.IsKeyJustPressed(ebiten.KeyArrowUp)
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		if ebiten.IsFullscreen() {
			ebiten.SetFullscreen(false)
		} else {
			ebiten.SetFullscreen(true)
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	w := tilesImage.Bounds().Dx()
	tileXCount := w / tileSize

	// Draw each tile with each DrawImage call.
	// As the source images of all DrawImage calls are always same,
	// this rendering is done very efficiently.
	// For more detail, see https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#Image.DrawImage
	const xCount = screenWidth / tileSize

	for i, t := range g.gameMap {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64((i%xCount)*tileSize), float64((i/xCount)*tileSize))

		sx := (t % tileXCount) * tileSize
		sy := (t / tileXCount) * tileSize

		screen.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {

	g := &Game{

		gameMap: []int{
			0, 1, 1, 0, 1, 1, 1, 0,
			0, 1, 0, 1, 0, 1, 1, 1,
			0, 0, 1, 1, 0, 1, 0, 15,
			2, 3, 4, 5, 6, 7, 10, 11,
			18, 17, 83, 62, 63, 8, 9, 14,
			17, 16, 62, 60, 61, 16, 12, 13,
			18, 60, 16, 17, 16, 83, 16, 17,
			83, 18, 17, 16, 18, 64, 83, 59,
		},
	}
	reverse(g.gameMap)

	ebiten.SetWindowSize(screenWidth*zoom, screenHeight*zoom)
	ebiten.SetWindowTitle("Super flying man and pig")
	// ebiten.SetWindowIcon([]image.Image{loadImage("assets/icon.png")})
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func reverse(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
