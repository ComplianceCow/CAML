package main

import (
	libs "commonlibs"
	"log"
)

var controlsVO *libs.ControlsVO
var metricsVO *libs.MetricsVO

func main() {

	//libs.GenerateSampleData()
	libs.GeneratePomMetrics(metricsVO)
	// libs.LoadFileData()
	// libs.DataDemo()

}

func init() {
	controlsVO = new(libs.ControlsVO)
	controlsFile := "../../configFiles/globalConfig/controls.json"
	if err := controlsVO.LoadControlsData(controlsFile, "json"); err != nil {
		log.Printf("%s", err)
	}

	metricsVO = new(libs.MetricsVO)
	metricsFile := "../../configFiles/globalConfig/metricsModel.yaml"
	if err := metricsVO.LoadMetricsData(metricsFile, "yaml"); err != nil {
		log.Printf("%s", err)
	}

}
