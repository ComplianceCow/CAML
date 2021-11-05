package main

type DependentControlType int

type MeasureUnit int
type MeasureType int
type MeasurePeriod int

type ArtifactType int

const (
	CAMVersion   string = "0.0.1"
	CAMNamespace string = "/v1/alpha1"
)

const (
	UndefinedMeasureUnit MeasureUnit = iota
	NumericInt
	NumericaFloat
)

const (
	UndefinedMeasureType MeasureType = iota
	Percentage
	Max
	Min
	Count
	Distinct_Count
	Sum
)

const (
	UndefinedMeasurePeriod MeasurePeriod = iota
	Hour
	Day
	Week
	Month
	Year
)

const (
	Undefined_Depedency_Type DependentControlType = iota
	Parent
	Child
	Sibling
)

const (
	UndefinedArtifactType ArtifactType = iota
	Measure
	Metric
	Report
	Asset
)
