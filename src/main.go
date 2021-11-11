package main

import (
	"fmt"
	"log"
)

func main() {
	generateSampleData()
	// loadFileData()

}

func loadFileData() {
	controlsFile := "../etc/ccm_controls.json"
	metricsFile := "../etc/ccm_metrics.json"

	controlsVO := new(ControlsVO)
	if err := controlsVO.loadControlsData(controlsFile, "json"); err != nil {
		log.Printf("%s", err)
	} else {
		log.Printf("%s\n", controlsVO)
	}

	metricsVO := new(MetricsVO)
	if err := metricsVO.loadMetricsData(metricsFile, "json"); err != nil {
		log.Printf("%s", err)
	} else {
		fmt.Printf("%s\n", metricsVO)
	}

}

func generateSampleData() {

	// if _, err := generateControlsData(); err != nil {
	// 	log.Fatalf("Error generating ccm controls data\n")
	// }

	// if _, err := generateMetricsData(); err != nil {
	// 	log.Fatalf("Error generating ccm controls data\n")
	// }

	// if _, err := generateMetricsConfiguration(); err != nil {
	// 	log.Fatalf("Error generating ccm controls data\n")
	// }

	if _, err := generateMetricsRuntime(); err != nil {
		log.Fatalf("Error generating ccm controls data\n")
	}

}
