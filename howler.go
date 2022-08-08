package howler

import (
	"errors"
	"syscall/js"
)

var howl = js.Global().Get("Howl")
var howler = js.Global().Get("Howler")
var muted bool

func New(opts HowlOptions) Howl {
	var tmp = js.Global().Get("Object").New()
	tmp.Set("src", opts.Source)

	if opts.Volume != nil {
		tmp.Set("volume", opts.Volume)
	}

	if opts.HTML5 {
		tmp.Set("html5", true)
	}

	if opts.Loop {
		tmp.Set("loop", true)
	}

	if opts.Preload != nil {
		tmp.Set("preload", opts.Preload)
	}

	if opts.Autoplay {
		tmp.Set("autoplay", true)
	}

	if opts.Mute {
		tmp.Set("mute", true)
	}

	if opts.Sprites != nil {
		sprites := make(map[string]any)
		for name, sprite := range opts.Sprites {
			sprites[name] = []any{sprite.Offset.Milliseconds(), sprite.Duration.Milliseconds(), sprite.Loop}
		}
		tmp.Set("sprite", sprites)
	}

	if opts.Rate != nil {
		tmp.Set("rate", opts.Rate)
	}

	if opts.Pool != nil {
		tmp.Set("pool", opts.Pool)
	}

	if opts.XHR.Truthy() {
		tmp.Set("xhr", opts.XHR)
	}

	if opts.Orientation != nil {
		tmp.Set("orientation", opts.Orientation)
	}

	if opts.Stereo != nil {
		tmp.Set("stereo", opts.Stereo)
	}

	if opts.Pos != nil {
		tmp.Set("pos", opts.Pos)
	}

	// TODO: PannerAttr

	if opts.OnLoad != nil {
		callback0(tmp, "onload", opts.OnLoad)
	}

	if opts.OnLoadError != nil {
		callbackerr(tmp, "onloaderror", opts.OnLoadError)
	}

	if opts.OnPlayError != nil {
		callbackerr(tmp, "onplayerror", opts.OnPlayError)
	}

	if opts.OnPlay != nil {
		callback1(tmp, "onplay", opts.OnPlay)
	}

	if opts.OnEnd != nil {
		callback1(tmp, "onend", opts.OnEnd)
	}

	if opts.OnPause != nil {
		callback1(tmp, "onpause", opts.OnPause)
	}

	if opts.OnStop != nil {
		callback1(tmp, "onstop", opts.OnStop)
	}

	if opts.OnMute != nil {
		callback1(tmp, "onmute", opts.OnMute)
	}

	if opts.OnVolume != nil {
		callback1(tmp, "onvolume", opts.OnVolume)
	}

	if opts.OnRate != nil {
		callback1(tmp, "onrate", opts.OnRate)
	}

	if opts.OnSeek != nil {
		callback1(tmp, "onseek", opts.OnSeek)
	}

	if opts.OnFade != nil {
		callback1(tmp, "onfade", opts.OnFade)
	}

	if opts.OnUnlock != nil {
		callback0(tmp, "onunlock", opts.OnUnlock)
	}

	if opts.OnStereo != nil {
		callback1(tmp, "onstereo", opts.OnStereo)
	}

	if opts.OnPos != nil {
		callback1(tmp, "onpos", opts.OnPos)
	}

	if opts.OnOrientation != nil {
		callback1(tmp, "onorientation", opts.OnOrientation)
	}

	return Howl{
		value: howl.New(tmp),
	}
}

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

// HTML5PoolSize gets the pool size. Each HTML5 Audio object must be unlocked
// individually, so we keep a global pool of unlocked nodes to share between all
// Howl instances. This pool gets created on the first user interaction and is
// set to the size of this property.
func HTML5PoolSize() int {
	return howler.Get("html5PoolSize").Int()
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

// Muted gets the muted state.
func Muted() bool {
	return muted
}

// Mute or unmute all sounds.
func Mute(mute bool) {
	howler.Call("mute", mute)
	muted = mute
}

// Volume gets the global volume for all sounds.
func Volume() float64 {
	return howler.Call("volume").Float()
}

// SetVolume sets the global volume for all sounds, relative to their own volume.
func SetVolume(volume float64) {
	howler.Call("volume", volume)
}

// Stop all sounds and reset their seek position to the beginning.
func Stop() {
	howler.Call("stop")
}

// Codecs checks supported audio codecs. Returns true if the codec is supported
// in the current browser.
func Codecs(extension string) bool {
	return howler.Call("codecs", extension).Truthy()
}

// Unload and destroy all currently loaded Howl objects. This will immediately
// stop all sounds and remove them from cache.
func Unload() {
	howler.Call("unload")
}

// ugly helper functions
func callback0(value js.Value, event string, cb func()) {
	value.Set(event, js.FuncOf(func(this js.Value, args []js.Value) any {
		cb()
		return nil
	}))
}

func callback1(value js.Value, event string, cb func(int)) {
	value.Set(event, js.FuncOf(func(this js.Value, args []js.Value) any {
		cb(args[0].Int())
		return nil
	}))
}

func callbackerr(value js.Value, event string, cb func(error)) {
	value.Set(event, js.FuncOf(func(this js.Value, args []js.Value) any {
		cb(errors.New(args[1].String()))
		return nil
	}))
}
