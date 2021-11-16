package commonlibs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func (controlsVO *ControlsVO) LoadControlsData(fileName string, fileType string) (err error) {

	fileByteArray, err := ioutil.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("error with loading the file: %s", err)
	}

	switch fileType {
	case "json":
		if err = json.Unmarshal(fileByteArray, controlsVO); err != nil {
			return fmt.Errorf("cannot unmarshal json file: %s", err)
		}
	case "yaml":
	default:
		return fmt.Errorf("specify a valid file type")

	}
	return nil
}

func (metricsVO *MetricsVO) LoadMetricsData(fileName string, fileType string) (err error) {

	fileByteArray, err := ioutil.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("error with loading the file: %s", err)
	}

	switch fileType {
	case "json":
		if err = json.Unmarshal(fileByteArray, metricsVO); err != nil {
			return fmt.Errorf("cannot unmarshal json file: %s", err)
		}
	case "yaml":
	default:
		return fmt.Errorf("specify a valid file type")

	}
	return nil
}
