//go:build debug
// +build debug

package ginroutemg

import (
	"log"
	"os"
)

func init() {
	f, err := os.OpenFile("debug.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}

	log.SetOutput(f)
}

func debug(m ...any) {
	log.Println(m...)
}
