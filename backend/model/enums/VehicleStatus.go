package enum_models

type VehhicleStatus int

const (
	Inactive VehhicleStatus = iota
	Active
	Driving
)