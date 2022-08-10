package howler

import (
	"syscall/js"
)

var howler = js.Global().Get("Howler")

// UsingWebAudio returns true if the Web Audio API is available.
func UsingWebAudio() bool {
	return howler.Get("usingWebAudio").Truthy()
}

// NoAudio returns true if no audio is available.
func NoAudio() bool {
	return howler.Get("noAudio").Truthy()
}

// AutoUnlock attempts to enable audio on mobile (iOS, Android, etc) devices and desktop Chrome/Safari.
func AutoUnlock() bool {
	return howler.Get("autoUnlock").Truthy()
}

// SetAutoUnlock sets the AutoUnlock property.
func SetAutoUnlock(autoUnlock bool) {
	howler.Set("autoUnlock", autoUnlock)
}

// HTML5PoolSize gets the pool size. Each HTML5 Audio object must be unlocked
// individually, so we keep a global pool of unlocked nodes to share between all
// Howl instances. This pool gets created on the first user interaction and is
// set to the size of this property.
func HTML5PoolSize() int {
	return howler.Get("html5PoolSize").Int()
}

// SetHTML5PoolSize sets the HTML5PoolSize property.
func SetHTML5PoolSize(size int) {
	howler.Set("html5PoolSize", size)
}

// AutoSuspend suspends the Web Audio AudioContext after 30 seconds of
// inactivity to decrease processing and energy usage. Automatically resumes upon
// new playback. Set this property to false to disable this behavior.
func AutoSuspend() bool {
	return howler.Get("autoSuspend").Truthy()
}

// SetAutoSuspend sets the AutoSuspend property.
func SetAutoSuspend(auto bool) {
	howler.Set("autoSuspend", auto)
}

// Volume gets the global volume for all sounds.
func Volume() float64 {
	return howler.Call("volume").Float()
}

// SetVolume sets the global volume for all sounds, relative to their own volume.
func SetVolume(volume float64) {
	howler.Call("volume", volume)
}

// Mute or unmute all sounds.
func Mute(mute bool) {
	howler.Call("mute", mute)
}

// Stop all sounds and reset their seek position to the beginning.
func Stop() {
	howler.Call("stop")
}

// Unload and destroy all currently loaded Howl objects. This will immediately
// stop all sounds and remove them from cache.
func Unload() {
	howler.Call("unload")
}

// Codecs checks supported audio codecs. Returns true if the codec is supported
// in the current browser.
func Codecs(extension string) bool {
	return howler.Call("codecs", extension).Truthy()
}

// SetStereo is a helper method to update the stereo panning position of all current
// Howls. Future Howls will not use this value unless explicitly set.
func SetStereo(stereo float64) {
	howler.Call("stereo", stereo)
}

// Pos gets the position of the listener in 3D cartesian space.
func Pos() (x, y, z float64) {
	arr := howler.Call("pos")
	x = arr.Index(0).Float()
	y = arr.Index(1).Float()
	z = arr.Index(2).Float()
	return
}

// SetPos sets the position of the listener in 3D cartesian space. Sounds using 3D
// position will be relative to the listener's position.
func SetPos(x, y, z float32) {
	howler.Call("pos", x, y, z)
}

// Orientation gets the direction the listener is pointing in the 3D cartesian
// space. A front and up vector must be provided. The front is the direction the
// face of the listener is pointing, and up is the direction the top of the
// listener is pointing. Thus, these values are expected to be at right angles
// from each other. [x, y, z, xUp, yUp, zUp]
func Orientation() (orientation []float64) {
	arr := howler.Call("orientation")
	orientation = make([]float64, 6)
	for i := 0; i < 6; i++ {
		orientation[i] = arr.Index(i).Float()
	}
	return
}

// SetOrientation sets the orientation.
func SetOrientation(x, y, z, upX, upY, upZ float64) {
	howler.Call("orientation", x, y, z, upX, upY, upZ)
}
