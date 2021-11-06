package main

import "log"

func main() {

	// if _, err := generateControlsData(); err != nil {
	// 	log.Fatalf("Error generating ccm controls data\n")
	// }

	if _, err := generateCAMData(); err != nil {
		log.Fatalf("Error generating ccm controls data\n")
	}
}
