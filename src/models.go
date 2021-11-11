package main

type ControlDefinition struct {
	*ControlDefinitionVO
	// DependentControls []struct {
	// 	*ControlDefinition   `json:"dependentControl,omitempty" yaml:"dependentControl,omitempty"`
	// 	DependentControlType DependentControlType `json:"dependentControlType,omitempty" yaml:"dependentControlType,omitempty"`
	// } `json:"dependentControls,omitempty" yaml:"dependentControls,omitempty"`
}

type MetricDefinition struct {
	*MetricDefinitionVO
}
type Measure struct {
	*MeasureVO
}

type SLORecommendation struct {
	*SLORecommendationVO
}

type Provider struct {
	*ProviderVO
}

type ProviderArtifacts struct {
	*Provider
	*ProviderArtifactsVO
}

type MeasureRuntime struct {
	*MetricDefinition
	*MeasureRuntimeVO
	PreviousMeasure *MeasureRuntime `json:"-" yaml:"-"`
}

type MetricsRuntime struct {
	*MetricsRuntimeVO
	PreviousMetrics *MetricsRuntime `json:"-" yaml:"-"`
}

type AuthzBoundary struct {
	*AuthzBoundaryVO
}
