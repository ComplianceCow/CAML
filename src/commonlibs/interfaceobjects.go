package commonlibs

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
	MetricWeight       int                   `json:"metricWeight,omitempty" yaml:"metricWeight,omitempty"`
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

type ProviderVO struct {
	ProviderName         string `json:"providerName,omitempty" yaml:"providerName,omitempty"`
	ProviderID           string `json:"providerID,omitempty" yaml:"providerID,omitempty"`
	ProviderAlias        string `json:"processoAlias,omitempty" yaml:"processoAlias,omitempty"`
	ProviderType         string `json:"processrType,omitempty" yaml:"processrType,omitempty"`
	ProviderToken        string `json:"providerToken,omitempty" yaml:"providerToken,omitempty"`
	ProviderClientID     string `json:"proessorClientID,omitempty" yaml:"proessorClientID,omitempty"`
	ProviderClientSecret string `json:"pocessorClientSecret,omitempty" yaml:"pocessorClientSecret,omitempty"`
	ProviderURL          string `json:"providerURL,omitempty" yaml:"providerURL,omitempty"`
}

type ProviderArtifactsVO struct {
	ProviderName      string             `json:"providerName,omitempty" yaml:"providerName,omitempty"`
	ProviderID        string             `json:"providerID,omitempty" yaml:"providerID,omitempty"`
	ProviderNotes     string             `json:"providerNotes,omitempty" yaml:"providerNotes,omitempty"`
	IsArtifactPresent bool               `json:"isArtifactPresent,omitempty" yaml:"isArtifactPresent,omitempty"`
	Artifacts         []ProviderArtifact `json:"artifacts,omitempty" yaml:"artifacts,omitempty"`
}

type ProviderArtifact struct {
	ArtifactName        string           `json:"artifactName,omitempty" yaml:"artifactName,omitempty"`
	ArtifactDescription string           `json:"artifactDescription,omitempty" yaml:"artifactDescription,omitempty"`
	ArtifactType        ArtifactType     `json:"artifactType,omitempty" yaml:"artifactType,omitempty"`
	ArtifactURL         string           `json:"artifactURL,omitempty" yaml:"artifactURL,omitempty"`
	ArtifactStatus      ProcessingStatus `json:"artifactStatus,omitempty" yaml:"artifactStatus,omitempty"`
}

type MeasureRuntimeVO struct {
	// *MetricDefinitionVO
	// providerName string `json:"providerName,omitempty" yaml:"providerName,omitempty"`
	// MeasureSourceID       string                `json:"measureID,omitempty" yaml:"measureID,omitempty"`
	MetricID              string               `json:"metricID,omitempty" yaml:"metricID,omitempty"`
	MeasureName           string               `json:"measureName,omitempty" yaml:"measureName,omitempty"`
	MeasureAlias          string               `json:"measureAlias,omitempty" yaml:"measureAlias,omitempty"`
	MeasureValue          string               `json:"measureValue,omitempty" yaml:"measureValue,omitempty"`
	MesurePeriodStartDate string               `json:"measureStartDt,omitempty" yaml:"measureStartDt,omitempty"`
	MesurePeriodEndDate   string               `json:"measureEndDt,omitempty" yaml:"measureEndDt,omitempty"`
	MeasureAuthzBoundary  *AuthzBoundaryVO     `json:"measureAuthzBoundary,omitempty" yaml:"measureAuthzBoundary,omitempty"`
	MeasureTags           map[string]string    `json:"measureTags,omitempty" yaml:"measureTags,omitempty"`
	MeasureReportDate     string               `json:"measureDate,omitempty" yaml:"measureDTM,omitempty"`
	MeasureEvidences      *ProviderArtifactsVO `json:"measureEvidences,omitempty" yaml:"measureEvidences,omitempty"`
	MeasurePolicy         *PolicyVO            `json:"measurePolicy,omitempty" yaml:"measurePolicy,omitempty"`
	MeasureExceptions     []ExceptionVO        `json:"exceptions,omitempty" yaml:"exceptions,omitempty"`
}

type MetricsRuntimeVO struct {
	// providerName          string                `json:"providerName,omitempty" yaml:"providerName,omitempty"`
	// MetricsSourceID        string                `json:"metricsID,omitempty" yaml:"metricsID,omitempty"`
	MetricID               string               `json:"metricID,omitempty" yaml:"metricID,omitempty"`
	MetricAlias            string               `json:"metricAlias,omitempty" yaml:"metricAlias,omitempty"`
	MetricsValue           string               `json:"metricsValue,omitempty" yaml:"metricsValue,omitempty"`
	MetricsPeriodStartDate string               `json:"metricsPeriodStartDate,omitempty" yaml:"metricsPeriodStartDate,omitempty"`
	MetricsPeriodEndDate   string               `json:"metricsPeriodEndDate,omitempty" yaml:"metricsPeriodEndDate,omitempty"`
	MetricsAuthzBoundary   *AuthzBoundaryVO     `json:"metricsAuthzBoundary,omitempty" yaml:"metricsAuthzBoundary,omitempty"`
	Measures               []MeasureRuntimeVO   `json:"measures,omitempty" yaml:"measures,omitempty"`
	MetricsTags            map[string]string    `json:"metricsTags,omitempty" yaml:"metricsTags,omitempty"`
	MetricsReportDate      string               `json:"metricsDate,omitempty" yaml:"metricsDTM,omitempty"`
	MetricsEvidences       *ProviderArtifactsVO `json:"metricsEvidences,omitempty" yaml:"metricsEvidences,omitempty"`
	MetricsPolicy          *PolicyVO            `json:"metricsPolicy,omitempty" yaml:"metricsPolicy,omitempty"`
	MetricsExceptions      []ExceptionVO        `json:"exceptions,omitempty" yaml:"exceptions,omitempty"`
}

type AuthzBoundaryVO struct {
	AuthzBoundaryName  string               `json:"authzBoundaryName,omitempty" yaml:"authozBoundaryName,omitempty"`
	AuthzBoundaryAlias string               `json:"authzBoundaryAlias,omitempty" yaml:"authzBoundaryAlias,omitempty"`
	AuthzBoundaryTags  map[string]string    `json:"authzBoundaryTags,omitempty" yaml:"authzBoundaryTags,omitempty"`
	AuthzAssets        *ProviderArtifactsVO `json:"authzAssets,omitempty" yaml:"authzAssets,omitempty"`
}
type PolicyVO struct {
	PolicyName        string `json:"policyName,omitempty" yaml:"policyName,omitempty"`
	PolicyDescription string `json:"policyDescription,omitempty" yaml:"policyDescription,omitempty"`
}

type ExceptionVO struct {
	Source   ArtifactType  `json:"source,omitempty" yaml:"source,omitempty"`
	Sequence int           `json:"seq,omitempty" yaml:"seq,omitempty"`
	Type     ExceptionType `json:"type,omitempty" yaml:"type,omitempty"`
	Message  string        `json:"message,omitempty" yaml:"message,omitempty"`
}
