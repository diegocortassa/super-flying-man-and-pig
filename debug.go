package main

import (
	"log"
)

func DebugPrintf(str ...interface{}) {
	if debug {
		// fmt.Println(str...)
		log.Println(str...)
		return
	}
}
