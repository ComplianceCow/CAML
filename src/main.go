package main

import "log"

func main() {

	if _, err := generateControlsData(); err != nil {
		log.Fatalf("Error generating ccm controls data\n")
	}

}
