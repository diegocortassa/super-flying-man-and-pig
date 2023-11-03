package globals

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
	ScreenWidth  = 256
	ScreenHeight = 384
	// screenHeight      = 416 // use this to show scrolling trick
	TilesScreenWidth  = 8
	TilesScreenHeight = 12
	ScrollSpeed       = 30 // speed milliseconds per 1 pixel (33 pixels/sec)
)

var (
	Debug    bool // command line flag
	MameKeys bool // command line flag
)
