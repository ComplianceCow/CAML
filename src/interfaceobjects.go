package main

type ControlsVO struct {
	Controls []ControlDefinitionVO `json:"controls,omitempty" yaml:"controls,omitempty"`
}

type ControlDefinitionVO struct {
	ControlDomain        string            `json:"controlDomain,omitempty" yaml:"controlDomain,omitempty"`
	ControlDomainCode    string            `json:"controlDomainCode,omitempty" yaml:"controlDomainCode,omitempty"`
	ControlTitle         string            `json:"controlTitle,omitempty" yaml:"controlTitle,omitempty"`
	ControlID            string            `json:"controlID,omitempty" yaml:"controlID,omitempty"`
	ControlAlias         string            `json:"controlAlias,omitempty" yaml:"controlAlias,omitempty"`
	ControlSpecification string            `json:"controlSpec,omitempty" yaml:"controlSpec,omitempty"`
	DependentControls    []string          `json:"dependentControls,omitempty" yaml:"dependentControls,omitempty"`
	ControlTags          map[string]string `json:"controlTags,omitempty" yaml:"controlTags,omitempty"`
}

type MetricsVO struct {
	Metrics []MetricDefinitionVO `json:"metrics,omitempty" yaml:"metrics,omitempty"`
}

type MetricDefinitionVO struct {
	ControlID                      string      `json:"controID,omitempty" yaml:"controlID,omitempty"`
	MetricID                       string      `json:"metricID,omitempty" yaml:"metricID,omitempty"`
	MetricAlias                    string      `json:"metricAlias,omitempty" yaml:"metricAlias,omitempty"`
	MetricDescription              string      `json:"metricDescription,omitempty" yaml:"metricDescription,omitempty"`
	MetricExpression               string      `json:"metricExpression,omitempty" yaml:"metricExpression,omitempty"`
	MetricRule                     string      `json:"metricRule,omitempty" yaml:"metricRule,omitempty"`
	MetricImplementationGuidelines string      `json:"metricImplementationGuidelines,omitempty" yaml:"metricImplementationGuidelines,omitempty"`
	MetricSLODescription           string      `json:"metricSLODescription,omitempty" yaml:"metricSLODescription,omitempty"`
	MetricMeasures                 []MeasureVO `json:"measures,omitempty" yaml:"measures,omitempty"`
	// MetricFormula: Example, "{{A}}/{{B}}*100" where {{A}} and {{B}} corresponds to the Metrics Measures
	MetricFormula string `json:"metricFormula,omitempty" yaml:"metricFormula,omitempty"`
	MetricPeriod  string `json:"metricPeriod,omitempty" yaml:"metricPeriod,omitempty"`
	// Metrics_Frequency in CRON format
	MetricFrequency    string                `json:"metricFrequency,omitempty" yaml:"metricFrequency,omitempty"`
	SLORecommendations []SLORecommendationVO `json:"metricSLORecommendations,omitempty" yaml:"metricSLORecommendations,omitempty"`
	MetricTags         map[string]string     `json:"metricTags,omitempty" yaml:"metricTags,omitempty"`
}
type MeasureVO struct {
	MeasureName        string `json:"measureName,omitempty" yaml:"measureName,omitempty"`
	MeasureAlias       string `json:"measureAlias,omitempty" yaml:"measureAlias,omitempty"`
	MeasureDescription string `json:"measureDescription,omitempty" yaml:"measureDescription,omitempty"`
	MeasureNotes       string `json:"measureNotes,omitempty" yaml:"measureNotes,omitempty"`
	// Refer MeasureUnit
	MeasureUnit   string            `json:"measureUnit,omitempty" yaml:"measureUnit,omitempty"`
	MeasureType   string            `json:"measureType,omitempty" yaml:"measureType,omitempty"`
	MeasurePeriod string            `json:"measurePeriod,omitempty" yaml:"measurePeriod,omitempty"`
	MeasureTags   map[string]string `json:"measureTags,omitempty" yaml:"measureTags,omitempty"`
}

type SLORecommendationVO struct {
	SLOCondition            string `json:"sloCondition,omitempty" yaml:"sloCondition,omitempty"`
	SLOConditionDescription string `json:"sloConditionDescription,omitempty" yaml:"sloConditionDescription,omitempty"`
	// SLO_Range values will be strings that will be cast at runtime into appropriate data types
	// format: 90d, 1y, 20h etc.
	SLORangeMin string `json:"sloRangeMin,omitempty" yaml:"sloRangeMin,omitempty"`
	SLORangeMax string `json:"sloRangeMax,omitempty" yaml:"sloRangeMax,omitempty"`
	/* SLOUnit:  Refer MeasureUnit. This is an enumerated variable and can be one of the following
	percentage - To indicate that the SLOType is a percentage
	max - To indicate that the SLOType is a MAX value of
	min - To indicate that the SLOType is a MIN value of
	*/
	SLOUnit string `json:"sloUnit,omitempty" yaml:"sloUnit,omitempty"`
	/* SLOPeriod: Refer MeasurePeriod. This is an enumerated variable and can be one of the following
	Undefined - Not applicable. In the case of %
	Hour, Day, Week, Month, Year
	*/
	SLOPeriod string `json:"sloPeriod,omitempty" yaml:"sloPeriod,omitempty"`
}

// The following specify the run time constructs

type ProcessorVO struct {
	ProcessorName  string `json:"processorName,omitempty" yaml:"processorName,omitempty"`
	ProcessorID    string `json:"processorID,omitempty" yaml:"processorID,omitempty"`
	ProcessorAlias string `json:"processoAlias,omitempty" yaml:"processoAlias,omitempty"`
	ProcessorType  string `json:"processrType,omitempty" yaml:"processrType,omitempty"`
	// processor_Token: Can we use a JWT token to guarantee the integrity of the measures
	ProcessorToken        string `json:"processorToken,omitempty" yaml:"processorToken,omitempty"`
	ProcessorClientID     string `json:"proessorClientID,omitempty" yaml:"proessorClientID,omitempty"`
	ProcessorClientSecret string `json:"pocessorClientSecret,omitempty" yaml:"pocessorClientSecret,omitempty"`
	ProcessorURL          string `json:"processorURL,omitempty" yaml:"processorURL,omitempty"`
}

type ProcessorArtifactsVO struct {
	ProcessorName     string `json:"processorName,omitempty" yaml:"processorName,omitempty"`
	ProcessorID       string `json:"processorID,omitempty" yaml:"processorID,omitempty"`
	IsArtifactPresent bool   `json:"isArtifactPresent,omitempty" yaml:"isArtifactPresent,omitempty"`
	Artifacts         []struct {
		ArtifactName        string       `json:"artifactName,omitempty" yaml:"artifactName,omitempty"`
		ArtifactDescription string       `json:"artifactDescription,omitempty" yaml:"artifactDescription,omitempty"`
		ArtifactType        ArtifactType `json:"artifactType,omitempty" yaml:"artifactType,omitempty"`
		Artifact_URL        string       `json:"artifactURL,omitempty" yaml:"artifactURL,omitempty"`
	} `json:"artifacts,omitempty" yaml:"artifacts,omitempty"`
}

type MeasureRuntimeVO struct {
	*MetricDefinitionVO
	MeasureSource *ProcessorVO `json:"measureProcessor,omitempty" yaml:"measureProcessor,omitempty"`
	// CAMMeasure_Source_ID: GUID value generated by the source engine for the measure. This may be validated by the consumer
	MeasureSourceID       string                `json:"measureID,omitempty" yaml:"measureID,omitempty"`
	MesurePeriodStartDate string                `json:"measureStartDt,omitempty" yaml:"measureStartDt,omitempty"`
	MesurePeriodEndDate   string                `json:"measureEndDt,omitempty" yaml:"measureEndDt,omitempty"`
	MeasureAuthzBoundary  *AuthzBoundaryVO      `json:"measureAuthzBoundary,omitempty" yaml:"measureAuthzBoundary,omitempty"`
	MeasureValue          string                `json:"measureValue,omitempty" yaml:"measureValue,omitempty"`
	MeasureTags           map[string]string     `json:"measureTags,omitempty" yaml:"measureTags,omitempty"`
	MeasureReportDate     string                `json:"measureDate,omitempty" yaml:"measureDTM,omitempty"`
	MeasureEvidences      *ProcessorArtifactsVO `json:"measureEvidences,omitempty" yaml:"measureEvidences,omitempty"`
	MeasurePolicy         *PolicyVO             `json:"measurePolicy,omitempty" yaml:"measurePolicy,omitempty"`
}

type MetricsRuntimeVO struct {
	MetricsSource          *ProcessorVO          `json:"metricsProcessor,omitempty" yaml:"metricsProcessor,omitempty"`
	MetricsSourceID        string                `json:"metricsID,omitempty" yaml:"metricsID,omitempty"`
	MetricsPeriodStartDate string                `json:"metricsPeriodStartDate,omitempty" yaml:"metricsPeriodStartDate,omitempty"`
	MetricsPeriodEndDate   string                `json:"metricsPeriodEndDate,omitempty" yaml:"metricsPeriodEndDate,omitempty"`
	MetricsAuthzBoundary   *AuthzBoundaryVO      `json:"metricsAuthzBoundary,omitempty" yaml:"metricsAuthzBoundary,omitempty"`
	MetricsValue           string                `json:"metricsValue,omitempty" yaml:"metricsValue,omitempty"`
	Measures               []MeasureRuntimeVO    `json:"measures,omitempty" yaml:"measures,omitempty"`
	MetricsTags            map[string]string     `json:"metricsTags,omitempty" yaml:"metricsTags,omitempty"`
	MetricsReportDate      string                `json:"metricsDate,omitempty" yaml:"metricsDTM,omitempty"`
	MetricsEvidences       *ProcessorArtifactsVO `json:"metricsEvidences,omitempty" yaml:"metricsEvidences,omitempty"`
	MetricsPolicy          *PolicyVO             `json:"metricsPolicy,omitempty" yaml:"metricsPolicy,omitempty"`
}

type AuthzBoundaryVO struct {
	AuthzBoundaryName  string                `json:"authzBoundaryName,omitempty" yaml:"authozBoundaryName,omitempty"`
	AuthzBoundaryAlias string                `json:"authzBoundaryAlias,omitempty" yaml:"authzBoundaryAlias,omitempty"`
	AuthzBoundaryTags  map[string]string     `json:"authzBoundaryTags,omitempty" yaml:"authzBoundaryTags,omitempty"`
	AuthzProcessor     *ProcessorVO          `json:"authzProcessor,omitempty" yaml:"authzProcessor,omitempty"`
	AuthzAssets        *ProcessorArtifactsVO `json:"authzAssets,omitempty" yaml:"authzAssets,omitempty"`
}
type PolicyVO struct {
	PolicyName        string
	PolicyDescription string
}
