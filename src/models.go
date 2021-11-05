package main

import (
	"net/url"
	"time"
)

type ControlDefinition struct {
	*ControlDefinitionVO
	DependentControls []struct {
		*ControlDefinition   `json:"dependentControl,omitempty" yaml:"dependentControl,omitempty"`
		DependentControlType DependentControlType `json:"dependentControlType,omitempty" yaml:"dependentControlType,omitempty"`
	} `json:"dependentControls,omitempty" yaml:"dependentControls,omitempty"`
	ControlTags map[string]string `json:"controlTags,omitempty" yaml:"controlTags,omitempty"`
}

type CAMMetricDefinition struct {
	CCMControl           *ControlDefinition `json:"ccmControlInfo,omitempty" yaml:"ccmControlInfo,omitempty"`
	CAMMetricID          string             `json:"metricID,omitempty" yaml:"metricID,omitempty"`
	CAMMetricAlias       string             `json:"metricAlias,omitempty" yaml:"metricAlias,omitempty"`
	CAMMetricDescription string             `json:"metricDescription,omitempty" yaml:"metricDescription,omitempty"`
	// CAMMetrics
	CAMMetricMeasures []CAMMeasure `json:"measures,omitempty" yaml:"measures,omitempty"`
	// Metrics_Formula: Example, "{{A}}/{{B}}*100" where {{A}} and {{B}} corresponds to the Metrics Measures
	CAMMetricFormula  string `json:"metricFormula,omitempty" yaml:"metricFormula,omitempty"`
	CAMMetricDuration string `json:"metricDuration,omitempty" yaml:"metricDuration,omitempty"`
	// Metrics_Frequency in CRON format
	CAMMetricFrequency    string                 `json:"metricFrequency,omitempty" yaml:"metricFrequency,omitempty"`
	CAMSLORecommendations []CAMSLORecommendation `json:"camMetricSLORecommendations,omitempty" yaml:"camMetricSLORecommendations,omitempty"`
	CAMMetricTags         map[string]string      `json:"cammetricTags,omitempty" yaml:"cammetricTags,omitempty"`
}
type CAMMeasure struct {
	MeasureName        string `json:"measureName,omitempty" yaml:"measureName,omitempty"`
	MeasureAlias       string `json:"measureAlias,omitempty" yaml:"measureAlias,omitempty"`
	MeasureDescription string `json:"measureDescription,omitempty" yaml:"measureDescription,omitempty"`
	// Refer MeasureUnit
	MeasureUnit   MeasureUnit       `json:"measureUnit,omitempty" yaml:"measureUnit,omitempty"`
	MeasureType   MeasureType       `json:"measureType,omitempty" yaml:"measureType,omitempty"`
	MeasurePeriod MeasurePeriod     `json:"measurePeriod,omitempty" yaml:"measurePeriod,omitempty"`
	MeasureTags   map[string]string `json:"measureTags,omitempty" yaml:"measureTags,omitempty"`
}

type CAMSLORecommendation struct {
	SLOCondition            string `json:"sloCondition,omitempty" yaml:"sloCondition,omitempty"`
	SLOConditionDescription string `json:"sloConditionDescription,omitempty" yaml:"sloConditionDescription,omitempty"`

	/* SLOUnit:  Refer MeasureUnit. This is an enumerated variable and can be one of the following
	numericint - To indicate that the SLO output is a signed number
	numeric_float - To indicate that the SLO output can take decimal values
	*/
	SLOUnit MeasureUnit `json:"sloUnit,omitempty" yaml:"sloUnit,omitempty"`

	/* SLOType:  Refer MeasureType. This is an enumerated variable and can be one of the following
	percentage - To indicate that the SLOType is a percentage
	max - To indicate that the SLOType is a MAX value of
	min - To indicate that the SLOType is a MIN value of
	*/
	SLOType MeasureType `json:"sloType,omitempty" yaml:"sloType,omitempty"`

	/* SLOPeriod: Refer MeasurePeriod. This is an enumerated variable and can be one of the following
	Undefined - Not applicable. In the case of %
	Hour, Day, Week, Month, Year
	*/
	SLOPeriod MeasurePeriod `json:"sloPeriod,omitempty" yaml:"sloPeriod,omitempty"`

	// SLO_Range values will be strings that will be cast at runtime into appropriate data types
	// format: 90d, 1y, 20h etc.
	SLORangeMin string `json:"sloRangeMin,omitempty" yaml:"sloRangeMin,omitempty"`
	SLORangeMax string `json:"sloRangeMax,omitempty" yaml:"sloRangeMax,omitempty"`
}

type CAMProcessor struct {
	ProcessorName  string `json:"processorName,omitempty" yaml:"processorName,omitempty"`
	ProcessorAlias string `json:"processoAlias,omitempty" yaml:"processoAlias,omitempty"`
	ProcessorType  string `json:"processrType,omitempty" yaml:"processrType,omitempty"`
	// processor_Token: Can we use a JWT token to guarantee the integrity of the measures
	ProcessorToken        string  `json:"processorToken,omitempty" yaml:"processorToken,omitempty"`
	ProcessorClientID     string  `json:"proessorClientID,omitempty" yaml:"proessorClientID,omitempty"`
	ProcessorClientSecret string  `json:"pocessorClientSecret,omitempty" yaml:"pocessorClientSecret,omitempty"`
	ProcessorURL          url.URL `json:"processorURL,omitempty" yaml:"processorURL,omitempty"`
}

type CAMProcessorArtifacts struct {
	*CAMProcessor
	IsArtifactPresent bool `json:"isArtifactPresent,omitempty" yaml:"isArtifactPresent,omitempty"`
	Artifacts         []struct {
		ArtifactName        string       `json:"artifactName,omitempty" yaml:"artifactName,omitempty"`
		ArtifactDescription string       `json:"artifactDescription,omitempty" yaml:"artifactDescription,omitempty"`
		ArtifactType        ArtifactType `json:"artifactType,omitempty" yaml:"artifactType,omitempty"`
		Artifact_URL        url.URL      `json:"artifactURL,omitempty" yaml:"artifactURL,omitempty"`
	} `json:"artifacts,omitempty" yaml:"artifacts,omitempty"`
}

type CAMMeasureRuntime struct {
	*CAMMetricDefinition
	MeasureSource *CAMProcessor `json:"measureProcessor,omitempty" yaml:"measureProcessor,omitempty"`
	// CAMMeasure_Source_ID: GUID value generated by the source engine for the measure. This may be validated by the consumer
	MeasureSourceID       string                 `json:"measureID,omitempty" yaml:"measureID,omitempty"`
	MesurePeriodStartDate time.Time              `json:"measureStartDt,omitempty" yaml:"measureStartDt,omitempty"`
	MesurePeriodEndDate   time.Time              `json:"measureEndDt,omitempty" yaml:"measureEndDt,omitempty"`
	MeasureAuthzBoundary  *CAMAuthzBoundary      `json:"measureAuthzBoundary,omitempty" yaml:"measureAuthzBoundary,omitempty"`
	MeasureValue          string                 `json:"measureValue,omitempty" yaml:"measureValue,omitempty"`
	MeasureTags           map[string]string      `json:"measureTags,omitempty" yaml:"measureTags,omitempty"`
	MeasureReportDTM      time.Time              `json:"measureDTM,omitempty" yaml:"measureDTM,omitempty"`
	MeasureEvidences      *CAMProcessorArtifacts `json:"measureEvidences,omitempty" yaml:"measureEvidences,omitempty"`
	PreviousMeasure       *CAMMeasureRuntime     `json:"-" yaml:"-"`
}

type CAMMetricsRuntime struct {
	Metrics_Source            *CAMProcessor          `json:"metricsProcessor,omitempty" yaml:"metricsProcessor,omitempty"`
	Metrics_Source_ID         string                 `json:"metricsID,omitempty" yaml:"metricsID,omitempty"`
	Metrics_Period_Start_Date time.Time              `json:"metricsPeriodStartDate,omitempty" yaml:"metricsPeriodStartDate,omitempty"`
	Metrics_Period_End_Date   time.Time              `json:"metricsPeriodEndDate,omitempty" yaml:"metricsPeriodEndDate,omitempty"`
	Metrics_Authz_Boundary    *CAMAuthzBoundary      `json:"metricsAuthzBoundary,omitempty" yaml:"metricsAuthzBoundary,omitempty"`
	Metrics_Value             string                 `json:"metricsValue,omitempty" yaml:"metricsValue,omitempty"`
	Measures                  []CAMMeasureRuntime    `json:"measures,omitempty" yaml:"measures,omitempty"`
	MetricsTags               map[string]string      `json:"metricsTags,omitempty" yaml:"metricsTags,omitempty"`
	MetricsReportDTM          time.Time              `json:"metricsDTM,omitempty" yaml:"metricsDTM,omitempty"`
	MetricsEvidences          *CAMProcessorArtifacts `json:"metricsEvidences,omitempty" yaml:"metricsEvidences,omitempty"`
	PreviousMetrics           *CAMMetricsRuntime     `json:"-" yaml:"-"`
}

type CAMAuthzBoundary struct {
	AuthzBoundaryName  string                 `json:"authozBoundaryName,omitempty" yaml:"authozBoundaryName,omitempty"`
	AuthzBoundaryAlias string                 `json:"authzBoundaryAlias,omitempty" yaml:"authzBoundaryAlias,omitempty"`
	AuthzBoundaryTags  map[string]string      `json:"authzBoundaryTags,omitempty" yaml:"authzBoundaryTags,omitempty"`
	AuthzProcessor     *CAMProcessor          `json:"authzProcessor,omitempty" yaml:"authzProcessor,omitempty"`
	AuthzAssets        *CAMProcessorArtifacts `json:"authzAssets,omitempty" yaml:"authzAssets,omitempty"`
}
