package commonlibs

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-yaml/yaml"
)

func GenerateControlsData() (controlsVO ControlsVO, err error) {

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
	controlsVO.Controls = ccmControls

	controlsJSON, err := json.Marshal(controlsVO)
	if err != nil {
		log.Fatalf("Error marshaling json data. %s", err.Error())
		return
	}
	fmt.Printf("%s\n\n\n", controlsJSON)

	controlsYAML, err := yaml.Marshal(controlsVO)
	if err != nil {
		log.Fatalf("Error marshaling json data. %s", err.Error())
		return
	}
	fmt.Printf("%s", controlsYAML)

	return controlsVO, nil
}

func GenerateMetricsData() (metricsVO MetricsVO, err error) {

	metrics := make([]MetricDefinitionVO, 0)

	metricDefinition := MetricDefinitionVO{}
	metricDefinition.ControlID = "AIS-06"
	metricDefinition.MetricID = "AIS-06-M1"
	metricDefinition.MetricDescription = "This metric measures the percentage of running production code that can be directly traced back to automated security and quality tests that verify the compliance of each build."
	metricDefinition.MetricExpression = `Percentage of compliant code: 100 * A/B
	Where:
	A = Total number of pieces of Production Code that have an Associated Verification Step
	B = Total number of pieces of Production Code`
	metricDefinition.MetricRule = `"Production Code" is code deployed to the production runtime environment(s) within the scope of the Information Security Program defined in the CCMv4 GRC-05 control objective.
	An "Associated Verification Step" is a capability in the deployment process that ties production code back to a build with traceable results for quality, security, and privacy tests.`
	metricDefinition.MetricImplementationGuidelines = `There must be a Software Inventory of Deployed Production Code (DCS-06).  Production code must be quantified based on the organization's definition of deployed code running in production (e.g. microservices, builds, releases, packages, libraries, serverless functions, etc.). This should be the same number used to measure AIS-07.
	The definition of "deployed production code" used for the software inventory should be aligned with application security scanning, testing, and/or reporting methods where possible to simplify measurement.
	The likelihood of standardized deployments can decrease as the number of different deployment systems increases. If the Software Deployment Pipeline has multiple stages where change could be introduced, and end-to-end validation cannot be performed, then this metric may be more suitable for an organization:
	0%<=Percentage of steps in the Software Deployment Pipeline that have an associated verification step<=100%
	There should be a mechanism to identify deviations, and if deviations from the standard are approved, then the system should account for (and manage) the exception as approved.
	This metric should at least be aligned with an organization's development or release cycle to provide timely input for correction in the next deployment or release. For example, if an organization uses an Agile development methodology with two-week sprints, then the metric should be measured at least every two weeks to provide data for review at sprint retros.`
	metricDefinition.MetricSLODescription = "95%"

	metrics = append(metrics, metricDefinition)

	metricDefinition = MetricDefinitionVO{}
	metricDefinition.ControlID = "AIS-07"
	metricDefinition.MetricID = "AIS-07-M3"
	metricDefinition.MetricDescription = "This metric measures the coverage for application vulnerability remediation across the production code"
	metricDefinition.MetricExpression = `"Percentage: 100 * A/B 
	Where:
	A = Number of deployed production applications with acceptable level of risk from application security vulnerabilities 
	B = Total Number of deployed production applications"`
	metricDefinition.MetricRule = `"Production Application = Applications tracked within the software inventory established in CCMv4 DCS-06
	Acceptable level of risk from application security vulnerabilities: Vulnerabilities categorized as medium or low risk as well as critical or high vulnerabilities marked or identified as 'Accepted' (i.e. remediation not required).  Examples of accepted vulnerabilities can be false positives or vulnerabilities with compensating controls that make the residual risk of exploitation acceptable."`
	metricDefinition.MetricImplementationGuidelines = `"There must be a Software Inventory of Deployed Production Code (see DCS-06 for more info). Production code must be quantified based on the organization's definition of deployed code running in production (e.g. microservices, builds, releases, packages, libraries, serverless functions, etc.). This should be the same number used to measure AIS-06.
	The definition of ""deployed production application"" used for the software inventory should be aligned with application security scanning, testing, and/or reporting methods where possible to simplify measurement.
	""Acceptable Level of Risk"" should be defined by the organizations vulnerability management guidelines (eg. only ""Critical"" and ""High"" vulnerabilities, or ""Medium Vulnerabilities and Higher,"" etc.  Classification of vulnerabilities as 'High' or 'Critical' risk, etc, should be defined in the Vulnerability Management tool based on industry-accepted scoring system such as the Common Vulnerability Scoring System (CVSS) (https://nvd.nist.gov/vuln-metrics/cvss). For instance, vulnerabilities with a CVSS score of 9 or higher are 'Critical', and vulnerabilities with CVSS scores between 7 and 9 could be defined as 'High' risk.)"`
	metricDefinition.MetricSLODescription = `"80% 
	Rationale: The 2020 APPLICATION SECURITY OBSERVABILITY REPORT from Contrast Labs found 26% of applications had at least 1 serious vulnerability with 79% of those vulnerabilities remediated within 30 days. That leaves 20% of applications with serious vulnerabilities after 30 days, so the SLO to have 80% of production code with acceptable level of risk from application security vulnerabilities should be achievable for the average organization."`

	metrics = append(metrics, metricDefinition)

	metricsVO.Metrics = metrics

	metricsJSON, err := json.Marshal(metricsVO)
	if err != nil {
		log.Fatalf("Error marshaling json data. %s", err.Error())
		return
	}
	fmt.Printf("%s\n\n\n", metricsJSON)

	metricsYAML, err := yaml.Marshal(metricsVO)
	if err != nil {
		log.Fatalf("Error marshaling json data. %s", err.Error())
		return
	}
	fmt.Printf("%s", metricsYAML)

	return metricsVO, nil
}

func GenerateMetricsConfiguration() (metricsVO MetricsVO, err error) {

	metrics := make([]MetricDefinitionVO, 0)

	{
		metricDefinition := MetricDefinitionVO{}
		metricDefinition.MetricID = "AIS-06-M1"
		measures := make([]MeasureVO, 0)

		{
			measure := MeasureVO{}
			measure.MeasureName = "prod_apps_with_verification"
			measure.MeasureAlias = "A"
			measure.MeasureDescription = "Total number of pieces of Production Code that have an Associated Verification Step"
			measure.MeasureType = "calculate"
			measure.MeasureUnit = "count"
			measure.MeasurePeriod = "30d"

			measures = append(measures, measure)
		}
		{
			measure := MeasureVO{}
			measure.MeasureName = "prod_apps_deployed"
			measure.MeasureAlias = "B"
			measure.MeasureDescription = "Total number of pieces of Production Code"
			measure.MeasureType = "calculate"
			measure.MeasureUnit = "count"
			measure.MeasurePeriod = "30d"

			measures = append(measures, measure)
		}
		metricDefinition.MetricMeasures = measures

		metricDefinition.MetricFormula = "%v%%,({{A}}/{{B}*100)"
		metricDefinition.MetricPeriod = "30d"
		metricDefinition.MetricFrequency = "0 5 */30 * *"
		metricDefinition.SLORecommendations = make([]SLORecommendationVO, 0)
		{
			slo := SLORecommendationVO{}
			slo.SLOPeriod = "30d"
			slo.SLORangeMin = "95%"
			metricDefinition.SLORecommendations = append(metricDefinition.SLORecommendations, slo)
		}

		metrics = append(metrics, metricDefinition)
	}

	{
		metricDefinition := MetricDefinitionVO{}
		metricDefinition.MetricID = "AIS-07-M3"
		measures := make([]MeasureVO, 0)

		{
			measure := MeasureVO{}
			measure.MeasureName = "prod_apps_deployed_with_acceptable_vuls"
			measure.MeasureAlias = "A"
			measure.MeasureDescription = "Number of deployed production applications with acceptable level of risk from application security vulnerabilities"
			measure.MeasureType = "calculate"
			measure.MeasureUnit = "count"
			measure.MeasurePeriod = "30d"

			measures = append(measures, measure)
		}
		{
			measure := MeasureVO{}
			measure.MeasureName = "prod_apps_deployed"
			measure.MeasureAlias = "B"
			measure.MeasureDescription = "Total Number of deployed production applications"
			measure.MeasureType = "calculate"
			measure.MeasureUnit = "count"
			measure.MeasurePeriod = "30d"

			measures = append(measures, measure)
		}
		metricDefinition.MetricMeasures = measures

		metricDefinition.MetricFormula = "%v%%,({{A}}/{{B}*100)"
		metricDefinition.MetricPeriod = "30d"
		metricDefinition.MetricFrequency = "0 5 */30 * *"
		metricDefinition.SLORecommendations = make([]SLORecommendationVO, 0)
		{
			slo := SLORecommendationVO{}
			slo.SLOPeriod = "30d"
			slo.SLORangeMin = "80%"
			metricDefinition.SLORecommendations = append(metricDefinition.SLORecommendations, slo)
		}

		metrics = append(metrics, metricDefinition)
	}

	metricsVO.Metrics = metrics

	metricsJSON, err := json.Marshal(metricsVO)
	if err != nil {
		log.Fatalf("Error marshaling json data. %s", err.Error())
		return
	}
	fmt.Printf("%s\n\n\n", metricsJSON)

	metricsYAML, err := yaml.Marshal(metricsVO)
	if err != nil {
		log.Fatalf("Error marshaling json data. %s", err.Error())
		return
	}
	fmt.Printf("%s", metricsYAML)

	return metricsVO, nil
}

func GenerateProcessorConfiguration() (providerVO ProviderVO, err error) {

	providerVO.ProviderName = "compliancecow"
	providerVO.ProviderURL = "https://dev.compliancecow.live/api/v1/login"
	providerVO.ProviderToken = ""
	providerVO.ProviderClientID = ""
	providerVO.ProviderClientSecret = ""

	return providerVO, nil

}

func GenerateMetricsRuntime() (metricsRuntimeVO MetricsRuntimeVO, err error) {

	metricsRuntimeVO.MetricID = "AIS-06-M1"
	metricsRuntimeVO.Measures = make([]MeasureRuntimeVO, 0)
	{
		measure := MeasureRuntimeVO{}
		measure.MeasureEvidences = &ProviderArtifactsVO{
			Artifacts: make([]ProviderArtifact, 0),
		}

		measure.MeasureName = "prod_apps_with_verification"
		measure.MeasureAlias = "A"
		measure.MeasureValue = "90"
		measure.MesurePeriodStartDate = "10/01/2021"
		measure.MesurePeriodEndDate = "10/31/2021"
		measure.MeasureReportDate = "11/05/2021"
		{
			artifact := ProviderArtifact{}
			artifact.ArtifactType = Measure_
			artifact.ArtifactName = ""
			artifact.ArtifactDescription = ""
			artifact.ArtifactStatus = Success_
			artifact.ArtifactURL = "https://dev.compliancecow.live/measures/ais06m1/files/7117758f-b14a-4784-822d-44383f17180b/security_test_coverage.csv"
			measure.MeasureEvidences.Artifacts = append(measure.MeasureEvidences.Artifacts, artifact)
		}

		metricsRuntimeVO.Measures = append(metricsRuntimeVO.Measures, measure)
	}

	{
		measure := MeasureRuntimeVO{}
		measure.MeasureEvidences = &ProviderArtifactsVO{
			Artifacts: make([]ProviderArtifact, 0),
		}
		measure.MeasureName = "prod_apps_deployed"
		measure.MeasureAlias = "B"
		measure.MeasureValue = "120"
		measure.MesurePeriodStartDate = "10/01/2021"
		measure.MesurePeriodEndDate = "10/31/2021"
		measure.MeasureReportDate = "11/05/2021"

		{
			artifact := ProviderArtifact{}
			artifact.ArtifactType = Measure_
			artifact.ArtifactName = ""
			artifact.ArtifactDescription = ""
			artifact.ArtifactStatus = Success_
			artifact.ArtifactURL = "https://dev.compliancecow.live/measures/ais06m1/files/7baadd77-fa4e-49f6-8945-464c969358df/inventory_pmt_apps.csv"
			measure.MeasureEvidences.Artifacts = append(measure.MeasureEvidences.Artifacts, artifact)
		}
		metricsRuntimeVO.Measures = append(metricsRuntimeVO.Measures, measure)

	}

	metricsRuntimeVO.MetricsAuthzBoundary = &AuthzBoundaryVO{
		AuthzBoundaryName:  "composite payment processing apps",
		AuthzBoundaryAlias: "pmt",
		AuthzAssets: &ProviderArtifactsVO{
			ProviderName:      "compliancecow",
			ProviderNotes:     "The authorization boundary for the composite payment processing applications is defined through an application configuration in ComplianceCow. Refer to https://<<url>> for more details",
			IsArtifactPresent: true,
		},
	}
	{
		metricsRuntimeVO.MetricsAuthzBoundary.AuthzAssets.Artifacts = make([]ProviderArtifact, 0)
		{
			artifact := ProviderArtifact{}
			artifact.ArtifactType = Asset_
			artifact.ArtifactName = "payment app config"
			artifact.ArtifactDescription = "application configuration for the payment processing application in json format"
			artifact.ArtifactURL = "https://dev.compliancecow.live/apps/app-01?format=yaml"

			metricsRuntimeVO.MetricsAuthzBoundary.AuthzAssets.Artifacts = append(metricsRuntimeVO.MetricsAuthzBoundary.AuthzAssets.Artifacts, artifact)
		}
	}

	metricRuntimeJSON, err := json.Marshal(metricsRuntimeVO)
	if err != nil {
		log.Fatalf("Error marshaling json data. %s", err.Error())
		return
	}
	fmt.Printf("%s\n\n\n", metricRuntimeJSON)

	metricRuntimeYAML, err := yaml.Marshal(metricsRuntimeVO)
	if err != nil {
		log.Fatalf("Error marshaling json data. %s", err.Error())
		return
	}
	fmt.Printf("%s", metricRuntimeYAML)

	return metricsRuntimeVO, nil

}
