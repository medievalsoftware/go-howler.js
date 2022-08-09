//go:build js && wasm

package howler

import (
	"syscall/js"
	"time"
)

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

type Sprite struct {
	// The offset in milliseconds.
	Offset time.Duration
	// The duration in milliseconds.
	Duration time.Duration
	// Set to true to automatically loop the sprite forever.
	Loop bool
}

type HowlOptions struct {
	// The sources to the track(s) to be loaded for the sound (URLs or base64 data
	// URIs). These should be in order of preference, howler.js will automatically
	// load the first one that is compatible with the current browser. If your files
	// have no extensions, you will need to explicitly specify the extension using
	// the format property.
	Source []any

	// The volume of the specific track, from 0.0 to 1.0.
	Volume any // default=Howler's global volume

	// Set to true to force HTML5 Audio. This should be used for large audio files so
	// that you don't have to wait for the full file to be downloaded and decoded
	// before playing.
	HTML5 bool

	// Set to true to automatically loop the sound forever.
	Loop bool

	// Automatically begin downloading the audio file when the Howl is defined. If
	// using HTML5 Audio, you can set this to 'metadata' to only preload the file's
	// metadata (to get its duration without download the entire file, for example).
	Preload any // default=true

	// Set to true to automatically start playback when sound is loaded.
	Autoplay bool

	// Set to true to load the audio muted.
	Mute bool

	// Define a sound sprite for the sound.
	Sprites map[string]Sprite

	// The rate of playback. 0.5 to 4.0, with 1.0 being normal speed.
	Rate any // default=1.0

	// The size of the inactive sounds pool. Once sounds are stopped or finish
	// playing, they are marked as ended and ready for cleanup. We keep a pool of
	// these to recycle for improved performance. Generally this doesn't need to be
	// changed. It is important to keep in mind that when a sound is paused, it won't
	// be removed from the pool and will still be considered active so that it can be
	// resumed later.
	Pool any // default=5

	// Howler.js automatically detects your file format from the extension, but you
	// may also specify a format in situations where extraction won't work (such as
	// with a SoundCloud stream).
	Format []any

	// When using Web Audio, howler.js uses an XHR request to load the audio files.
	// If you need to send custom headers, set the HTTP method or enable
	// withCredentials (see reference), include them with this parameter. Each is
	// optional (method defaults to GET, headers default to null and withCredentials
	// defaults to false).
	XHR js.Value

	// Sets the direction the audio source is pointing in the 3D cartesian coordinate
	// space. Depending on how directional the sound is, based on the cone
	// attributes, a sound pointing away from the listener can be quiet or silent.
	Orientation []any

	// Sets the stereo panning value of the audio source for this sound or group.
	// This makes it easy to setup left/right panning with a value of -1.0 being far
	// left and a value of 1.0 being far right.
	Stereo any

	// Sets the 3D spatial position of the audio source for this sound or group
	// relative to the global listener.
	Pos []any

	// Sets the panner node's attributes for a sound or group of sounds. See the
	// pannerAttr method for all available options.
	PannerAttr js.Value
	// TODO: PannerAttribute struct

	// Fires when the sound is loaded.
	OnLoad func()

	// Fires when the sound is unable to load.
	OnLoadError func(error)

	// Fires when the sound is unable to play.
	OnPlayError func(error)

	// Fires when the sound begins playing.
	OnPlay func(id int)

	// Fires when the sound finishes playing (if it is looping, it'll fire at the end
	// of each loop).
	OnEnd func(id int)

	// Fires when the sound has been paused.
	OnPause func(id int)

	// Fires when the sound has been stopped.
	OnStop func(id int)

	// Fires when the sound has been muted/unmuted.
	OnMute func(id int)

	// Fires when the sound's volume has changed.
	OnVolume func(id int)

	//Fires when the sound's playback rate has changed.
	OnRate func(id int)

	//Fires when the sound has been seeked.
	OnSeek func(id int)

	// Fires when the current sound finishes fading in/out.
	OnFade func(id int)

	// Fires when audio has been automatically unlocked through a touch/click event.
	OnUnlock func()

	// Fires when the current sound has the stereo panning changed.
	OnStereo func(id int)

	// Fires when the current sound has the listener position changed.
	OnPos func(id int)

	// Fires when the current sound has the direction of the listener changed.
	OnOrientation func(id int)
}

type Howl struct {
	value js.Value
}

// Play begins playback of a sound. Returns the sound id to be used with other
// methods.
func (h Howl) Play() int {
	return h.value.Call("play").Int()
}

// PlayID will play a new sound will play based on the sprite's definition.
func (h Howl) PlayID(id int) int {
	return h.value.Call("play", js.ValueOf(id)).Int()
}

// PlaySprite will play the previously played sound (for example, after pausing
// it). However, if an id of a sound that has been drained from the pool is
// passed, nothing will play.
func (h Howl) PlaySprite(name string) int {
	return h.value.Call("play", name).Int()
}

// Pause pauses playback of the group, saving the seek of playback.
func (h Howl) Pause() {
	h.value.Call("pause")
}

// PauseID pauses playback of the sound, saving the seek of playback.
func (h Howl) PauseID(id int) {
	h.value.Call("pause", id)
}

// Stop stops all sounds of the group, resetting seek to 0.
func (h Howl) Stop() {
	h.value.Call("stop")
}

// StopID stops playback of sound, resetting seek to 0.
func (h Howl) StopID(id int) {
	h.value.Call("stop", id)
}

func (h Howl) Muted() bool {
	return h.value.Call("mute").Bool()
}

// Mute all sounds in the group.
func (h Howl) Mute() {
	h.value.Call("mute", true)
}

// Unmute all sounds in the group.
func (h Howl) Unmute() {
	h.value.Call("mute", false)
}

// MuteID mutes a sound, but doesn't pause the playback.
func (h Howl) MuteID(id int) {
	h.value.Call("mute", true, id)
}

// UnmuteID unmutes a sound.
func (h Howl) UnmuteID(id int) {
	h.value.Call("mute", false, id)
}

// Volume gets the volume of the group.
func (h Howl) Volume() float64 {
	return h.value.Call("volume").Float()
}

// SetVolume sets the volume of all the sounds in the group relative to their own
// volume.
func (h Howl) SetVolume(volume float64) {
	h.value.Call("volume", volume)
}

// SetVolumeID sets the volume of a sound.
func (h Howl) SetVolumeID(id int, volume float64) {
	h.value.Call("volume", volume, id)
}

// Fade fades all sounds in the group.
func (h Howl) Fade(from, to float64, duration time.Duration) {
	h.value.Call("fade", from, to, duration.Milliseconds())
}

// FadeID fades a sound between two volumes. Fires the OnFade event when complete.
func (h Howl) FadeID(id int, from, to float64, duration time.Duration) {
	h.value.Call("fade", from, to, duration.Milliseconds(), id)
}

// Rate gets the rate of the group.
func (h Howl) Rate() float64 {
	return h.value.Call("rate").Float()
}

// SetRate sets the rate of the group. Playback rate of all sounds in group will
// change.
func (h Howl) SetRate(rate float64) {
	h.value.Call("rate", rate)
}

// SetRateID sets the rate of a sound in the group.
func (h Howl) SetRateID(id int, rate float64) {
	h.value.Call("rate", rate, id)
}

// Seek gets the position of playback for the first sound.
func (h Howl) Seek() time.Duration {
	return time.Duration(h.value.Call("seek").Float() * float64(time.Second))
}

// SetSeek sets the position of playback for the first sound.
func (h Howl) SetSeek(seek time.Duration) {
	h.value.Call("seek", seek.Seconds())
}

func (h Howl) SetSeekID(id int, seek time.Duration) {
	h.value.Call("seek", seek.Seconds(), id)
}

// Loop gets whether the first sound will loop.
func (h Howl) Loop() bool {
	return h.value.Call("loop").Bool()
}

// SetLoop sets whether all sounds in the group will loop.
func (h Howl) SetLoop(loop bool) {
	h.value.Call("loop", loop)
}

// SetLoopID sets whether a sound in the group will loop.
func (h Howl) SetLoopID(id int, loop bool) {
	h.value.Call("loop", loop, id)
}

// State checks the load status of the Howl, returns a unloaded, loading or loaded.
func (h Howl) State() State {
	return states[h.value.Call("state").String()]
}

// Playing checks if any sound in the group is playing.
func (h Howl) Playing() bool {
	return h.value.Call("playing").Bool()
}

// PlayingID checks if a sound in the group is playing.
func (h Howl) PlayingID(id int) bool {
	return h.value.Call("playing", id).Bool()
}

// Duration gets the duration of the audio source. Returns zero until the load event has fired.
func (h Howl) Duration() time.Duration {
	return time.Duration(h.value.Call("duration").Float() * float64(time.Second))
}

// DurationID gets return the duration of the sprite being played on this
// instance; otherwise, the full source duration is returned.
func (h Howl) DurationID(id int) time.Duration {
	return time.Duration(h.value.Call("duration", id).Float() * float64(time.Second))
}

// TODO: on(event, function, [id])
// TODO: once(event, function, [id])
// TODO: off(event, [function], [id])

// Load is called by default, but if you set preload to false, you must call load
// before you can play any sounds.
func (h Howl) Load() {
	h.value.Call("load")
}

// Unload and destroy a Howl object. This will immediately stop all sounds
// attached to this sound and remove it from the cache.
func (h Howl) Unload() {
	h.value.Call("unload")
}

// TODO: [Set]Stereo[ID]
// TODO: [Set]Pos[ID]
// TODO: [Set]Orientation[ID]
// TODO: [Set]PannerAttributes[ID]
