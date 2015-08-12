package main

import (
	"log"

	"github.com/getlantern/osversion"
)

func main() {
	str, err := osversion.GetString()
	if err != nil {
		log.Fatal("Error getting OS version")
	}
	log.Println(str)

	version, err := osversion.GetSemanticVersion()
	if err != nil {
		log.Fatal("Error getting OS version")
	}

	log.Println(version)
}
