package howler

import (
	"syscall/js"
	"time"
)

// State describes the load status of a given Howl.
type State int

const (
	StateUnloaded State = iota
	StateLoading
	StateLoaded
)

var states = map[string]State{
	"unloaded": StateUnloaded,
	"loading":  StateLoading,
	"loaded":   StateLoaded,
}

type Sound interface {
	ID() int
	State() State
	Playing() bool
	Duration() time.Duration
	Play() Sound
	Pause()
	Stop()
	Mute()
	Unmute()
	Fade(from float64, to float64, duration time.Duration)
	Volume() float64
	SetVolume(volume float64)
	Rate() float64
	SetRate(float64)
	Seek() time.Duration
	SetSeek(duration time.Duration)
	Loop() bool
	SetLoop(loop bool)
	Stereo() float64
	SetStereo(stereo float64)
	Pos() (x, y, z float64)
	SetPos(x, y, z float64)
	Orientation() (x, y, z float64)
	SetOrientation(x, y, z float64)
	PannerAttr() PannerAttr
	SetPannerAttr(attr PannerAttr)
}

type soundGroup struct {
	value js.Value
}

func (g soundGroup) ID() int {
	return -1
}

func (g soundGroup) State() State {
	return states[g.value.Call("state").String()]
}

func (g soundGroup) Playing() bool {
	return g.value.Call("playing").Bool()
}

func (g soundGroup) Duration() time.Duration {
	return time.Duration(g.value.Call("duration").Float() * float64(time.Second))
}

func (g soundGroup) Play() Sound {
	if result := g.value.Call("play"); result.Truthy() {
		return soundSpecific{
			id:    result.Int(),
			value: g.value,
		}
	}
	return soundGroup{}
}

func (g soundGroup) Pause() {
	g.value.Call("pause")
}

func (g soundGroup) Stop() {
	g.value.Call("stop")
}

func (g soundGroup) Mute() {
	g.value.Call("mute", true)
}

func (g soundGroup) Unmute() {
	g.value.Call("mute", false)
}

func (g soundGroup) Fade(from float64, to float64, duration time.Duration) {
	g.value.Call("fade", from, to, duration.Milliseconds())
}

func (g soundGroup) Volume() float64 {
	return g.value.Call("volume").Float()
}

func (g soundGroup) SetVolume(volume float64) {
	g.value.Call("volume", volume)
}

func (g soundGroup) Rate() float64 {
	return g.value.Call("rate").Float()
}

func (g soundGroup) SetRate(rate float64) {
	g.value.Call("rate", rate)
}

func (g soundGroup) Seek() time.Duration {
	return time.Duration(g.value.Call("seek").Float() * float64(time.Second))
}

func (g soundGroup) SetSeek(position time.Duration) {
	g.value.Call("seek", position.Seconds())
}

func (g soundGroup) Loop() bool {
	return g.value.Call("loop").Bool()
}

func (g soundGroup) SetLoop(loop bool) {
	g.value.Call("loop", loop)
}

func (g soundGroup) Stereo() float64 {
	return g.value.Call("stereo").Float()
}

func (g soundGroup) SetStereo(stereo float64) {
	g.value.Call("stereo", stereo)
}

func (g soundGroup) Pos() (x, y, z float64) {
	pos := g.value.Call("pos")
	x = pos.Index(0).Float()
	y = pos.Index(1).Float()
	z = pos.Index(2).Float()
	return
}

func (g soundGroup) SetPos(x, y, z float64) {
	g.value.Call("pos", x, y, z)
}

func (g soundGroup) Orientation() (x, y, z float64) {
	pos := g.value.Call("orientation")
	x = pos.Index(0).Float()
	y = pos.Index(1).Float()
	z = pos.Index(2).Float()
	return
}

func (g soundGroup) SetOrientation(x, y, z float64) {
	g.value.Call("orientation", x, y, z)
}

func (g soundGroup) PannerAttr() PannerAttr {
	return PannerAttr{value: g.value.Call("pannerAttr")}
}

func (g soundGroup) SetPannerAttr(attr PannerAttr) {
	g.value.Call("pannerAttr", attr.value)
}

type soundSpecific struct {
	id    int
	value js.Value
}

func (s soundSpecific) ID() int {
	return s.id
}

func (s soundSpecific) State() State {
	return states[s.value.Call("state").String()]
}

func (s soundSpecific) Playing() bool {
	return s.value.Call("playing", s.id).Bool()
}

func (s soundSpecific) Duration() time.Duration {
	return time.Duration(s.value.Call("duration", s.id).Float() * float64(time.Second))
}

func (s soundSpecific) Play() Sound {
	s.value.Call("play", s.id)
	return s
}

func (s soundSpecific) Pause() {
	s.value.Call("pause", s.id)
}

func (s soundSpecific) Stop() {
	s.value.Call("stop", s.id)
}

func (s soundSpecific) Mute() {
	s.value.Call("mute", true, s.id)
}

func (s soundSpecific) Unmute() {
	s.value.Call("mute", false, s.id)
}

func (s soundSpecific) Fade(from float64, to float64, duration time.Duration) {
	s.value.Call("fade", from, to, duration.Milliseconds(), s.id)
}

func (s soundSpecific) Volume() float64 {
	return s.value.Call("volume", s.id).Float()
}

func (s soundSpecific) SetVolume(volume float64) {
	s.value.Call("volume", volume, s.id)
}

func (s soundSpecific) Stereo() float64 {
	return s.value.Call("stereo", s.id).Float()
}

func (s soundSpecific) SetStereo(stereo float64) {
	s.value.Call("stereo", stereo, s.id)
}

func (s soundSpecific) Rate() float64 {
	return s.value.Call("rate", s.id).Float()
}

func (s soundSpecific) SetRate(rate float64) {
	s.value.Call("rate", rate, s.id)
}

func (s soundSpecific) Seek() time.Duration {
	return time.Duration(s.value.Call("seek", s.id).Float() * float64(time.Second))
}

func (s soundSpecific) SetSeek(position time.Duration) {
	s.value.Call("seek", position.Seconds(), s.id)
}

func (s soundSpecific) Loop() bool {
	return s.value.Call("loop", s.id).Bool()
}

func (s soundSpecific) SetLoop(loop bool) {
	s.value.Call("loop", loop, s.id)
}

func (s soundSpecific) Pos() (x, y, z float64) {
	pos := s.value.Call("pos", js.Null(), js.Null(), js.Null(), s.id)
	x = pos.Index(0).Float()
	y = pos.Index(1).Float()
	z = pos.Index(2).Float()
	return
}

func (s soundSpecific) SetPos(x, y, z float64) {
	s.value.Call("pos", x, y, z)
}

func (s soundSpecific) Orientation() (x, y, z float64) {
	pos := s.value.Call("orientation", js.Null(), js.Null(), js.Null(), s.id)
	x = pos.Index(0).Float()
	y = pos.Index(1).Float()
	z = pos.Index(2).Float()
	return
}

func (s soundSpecific) SetOrientation(x, y, z float64) {
	s.value.Call("orientation", x, y, z)
}

func (s soundSpecific) PannerAttr() PannerAttr {
	return PannerAttr{value: s.value.Call("pannerAttr", s.id)}
}

func (s soundSpecific) SetPannerAttr(attr PannerAttr) {
	s.value.Call("pannerAttr", attr.value, s.id)
}
