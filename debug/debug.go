package debug

import (
	"log"

	"github.com/dcortassa/superflyingmanandpig/globals"
)

func DebugPrintf(str ...interface{}) {
	if globals.Debug {
		// fmt.Println(str...)
		log.Println(str...)
		return
	}
}
