package main

type DependentControlType int

type MeasureUnit int
type MeasureType int
type MeasurePeriod int

type ArtifactType int

const (
	Version_   string = "0.0.1"
	Namespace_ string = "/v1/alpha1"
)

const (
	UndefinedMeasureUnit_ MeasureUnit = iota
	Percentage_
	Max_
	Min_
	Count_
	DistinctCount_
	Sum_
)

const (
	UndefinedMeasureType_ MeasureType = iota
	Calculate_
	Info_
)

const (
	UndefinedMeasurePeriod_ MeasurePeriod = iota
	Hour_
	Day_
	Week_
	Month_
	Year_
)

const (
	UndefinedDepedencyType_ DependentControlType = iota
	Parent_
	Child_
	Sibling_
)

const (
	UndefinedArtifactType_ ArtifactType = iota
	Measure_
	Metric_
	Report_
	Asset_
)
