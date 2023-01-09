package main

import (
	"log"

	"github.com/getlantern/osversion"

	_ "golang.org/x/mobile/app"
)

func main() {
	str, err := osversion.GetHumanReadable()
	if err != nil {
		log.Printf("Error in osversion.GetHumanReadable: %v", err)
	}
	// PINEAPPLE is just to make grepping through logcat easier
	log.Printf("PINEAPPLE: %s\n", str)
}
