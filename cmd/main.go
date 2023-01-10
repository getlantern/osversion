//go:build !android

package main

import (
	"log"

	"github.com/getlantern/osversion"
)

func main() {
	hstr, err := osversion.GetHumanReadable()
	if err != nil {
		log.Fatalf("Error getting OS version: %v", err)
	}

	log.Println(hstr)
}
