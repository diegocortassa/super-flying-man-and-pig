package debug

import (
	"log"

	"github.com/dcortassa/super-flying-man-and-pig/globals"
)

func DebugPrintf(str ...interface{}) {
	if globals.Debug {
		// fmt.Println(str...)
		log.Println(str...)
		return
	}
}
