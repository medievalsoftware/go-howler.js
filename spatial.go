package howler

import (
	"fmt"
	"syscall/js"
)

type DistanceModel int

const (
	DistanceModelUndefined DistanceModel = iota
	DistanceModelLinear
	DistanceModelInverse
	DistanceModelExponential
)

type PanningModel int

const (
	PanningModelUndefined PanningModel = iota
	PanningModelHRTF
	PanningModelEqualPower
)

type PannerOptions struct {
	// ConeInnerAngle is a parameter for directional audio sources, this is an angle, in
	// degrees, inside of which there will be no volume reduction. default: 360
	ConeInnerAngle OptionalFloat `json:"cone_inner_angle,omitempty"`

	// ConeOuterAngle is a parameter for directional audio sources, this is an angle, in
	// degrees, outside of which the volume will be reduced to a constant value of
	// ConeOuterGain. default: 360
	ConeOuterAngle OptionalFloat `json:"cone_outer_angle,omitempty"`

	// ConeOuterGain is a parameter for directional audio sources, this is the gain
	// outside of the coneOuterAngle. It is a linear value in the range [0, 1].
	ConeOuterGain OptionalFloat `json:"cone_outer_gain,omitempty"`

	// DistanceModel determines algorithm used to reduce volume as audio moves away
	// from listener. Can be DistanceModelLinear, DistanceModelInverse or DistanceModelExponential.
	DistanceModel DistanceModel `json:"distance_model,omitempty"`

	// The maximum distance between source and listener, after which the volume will
	// not be reduced any further. default: 10000
	MaxDistance OptionalFloat `json:"max_distance,omitempty"`

	// A reference distance for reducing volume as source moves further from the
	// listener. This is simply a variable of the distance model and has a different
	// effect depending on which model is used and the scale of your coordinates.
	// Generally, volume will be equal to 1 at this distance. default: 1
	RefDistance OptionalFloat `json:"ref_distance,omitempty"`

	// How quickly the volume reduces as source moves from listener. This is simply a
	// variable of the distance model and can be in the range of [0, 1] with linear
	// and [0, ∞] with inverse and exponential.
	RolloffFactor OptionalFloat `json:"rolloff_factor,omitempty"`

	// Determines which spatialization algorithm is used to position audio. Can be
	// PanningModelHRTF or PanningModelEqualPower.
	PanningModel PanningModel `json:"panning_model,omitempty"`
}

type PannerAttr struct {
	value js.Value
}

func (a PannerAttr) ConeInnerAngle() float64 {
	return a.value.Get("coneInnerAngle").Float()
}

func (a PannerAttr) SetConeInnerAngle(angle float64) {
	a.value.Set("coneInnerAngle", angle)
}

func (a PannerAttr) ConeOuterAngle() float64 {
	return a.value.Get("coneOuterAngle").Float()
}

func (a PannerAttr) SetConeOuterAngle(angle float64) {
	a.value.Set("coneOuterAngle", angle)
}

func (a PannerAttr) ConeOuterGain() float64 {
	return a.value.Get("coneOuterGain").Float()
}

func (a PannerAttr) SetConeOuterGain(gain float64) {
	a.value.Set("coneOuterGain", gain)
}

func (a PannerAttr) DistanceModel() DistanceModel {
	switch a.value.Get("distanceModel").String() {
	case "linear":
		return DistanceModelLinear
	case "inverse":
		return DistanceModelInverse
	case "exponential":
		return DistanceModelExponential
	default:
		return DistanceModelUndefined
	}
}

func (a PannerAttr) SetDistanceModel(model DistanceModel) {
	switch model {
	case DistanceModelLinear:
		a.value.Set("distanceModel", "linear")
	case DistanceModelInverse:
		a.value.Set("distanceModel", "inverse")
	case DistanceModelExponential:
		a.value.Set("distanceModel", "exponential")
	default:
		panic(fmt.Errorf("unknown distance model: %d", model))
	}
}

func (a PannerAttr) MaxDistance() float64 {
	return a.value.Get("maxDistance").Float()
}

func (a PannerAttr) SetMaxDistance(distance float64) {
	a.value.Set("maxDistance", distance)
}

func (a PannerAttr) RefDistance() float64 {
	return a.value.Get("refDistance").Float()
}

func (a PannerAttr) SetRefDistance(distance float64) {
	a.value.Set("refDistance", distance)
}

func (a PannerAttr) RolloffFactor() float64 {
	return a.value.Get("rolloffFactor").Float()
}

func (a PannerAttr) SetRolloffFactor(factor float64) {
	a.value.Set("rolloffFactor", factor)
}

func (a PannerAttr) PanningModel() PanningModel {
	switch a.value.Get("panningModel").String() {
	case "HRTF":
		return PanningModelHRTF
	case "equalpower":
		return PanningModelEqualPower
	default:
		return PanningModelUndefined
	}
}

func (a PannerAttr) SetPanningModel(model PanningModel) {
	switch model {
	case PanningModelHRTF:
		a.value.Set("panningModel", "HRTF")
	case PanningModelEqualPower:
		a.value.Set("panningModel", "equalpower")
	default:
		panic(fmt.Errorf("unknown panning model: %d", model))
	}
}
