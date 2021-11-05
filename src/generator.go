package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-yaml/yaml"
)

func generateControlsData() (controlsVO []ControlDefinitionVO, err error) {
	ccmControls := make([]ControlDefinitionVO, 0)

	ccmControl := ControlDefinitionVO{}
	ccmControl.ControlDomain = "Application & Interface Security"
	ccmControl.ControlTitle = "Automated Secure Application Deployment"
	ccmControl.ControlID = "AIS-06"
	ccmControl.ControlSpecification = `Establish and implement strategies and capabilities for secure, standardized, and compliant application deployment. Automate where possible.`
	ccmControl.DependentControls = []string{"DCS-06", "GRC-05"}

	ccmControls = append(ccmControls, ccmControl)

	ccmControl = ControlDefinitionVO{}
	ccmControl.ControlDomain = "Application & Interface Security"
	ccmControl.ControlTitle = "Automated Secure Application Deployment"
	ccmControl.ControlID = "AIS-07"
	ccmControl.ControlSpecification = `Define and implement a process to remediate application security vulnerabilities, automating remediation when possible.`
	ccmControl.DependentControls = []string{"AIS-02", "AIS-03", "AIS-06", "DCS-06", "GRC-02", "TVM-10", "TVM-07"}

	ccmControls = append(ccmControls, ccmControl)

	ccmControlsJSON, err := json.Marshal(ccmControls)
	if err != nil {
		log.Fatalf("Error marshaling json data. %s", err.Error())
		return
	}
	fmt.Printf("%s\n\n\n", ccmControlsJSON)

	ccmControlsYAML, err := yaml.Marshal(ccmControls)
	if err != nil {
		log.Fatalf("Error marshaling json data. %s", err.Error())
		return
	}
	fmt.Printf("%s", ccmControlsYAML)

	return ccmControls, nil
	// ccmControls = append(ccmControls, ccmControl)
}

func generateCAMData() {

}
