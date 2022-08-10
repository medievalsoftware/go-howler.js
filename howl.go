//go:build js && wasm

package howler

import (
	"syscall/js"
	"time"
)

var howl = js.Global().Get("Howl")

func New(opts HowlOptions) Howl {
	var tmp = js.Global().Get("Object").New()

	tmp.Set("src", opts.Source)
	tmp.Set("format", opts.Format)
	tmp.Set("volume", opts.Volume)
	tmp.Set("html5", opts.HTML5)
	tmp.Set("loop", opts.Loop)
	tmp.Set("preload", opts.Preload)
	tmp.Set("autoplay", opts.Autoplay)
	tmp.Set("mute", opts.Mute)
	tmp.Set("rate", opts.Rate)
	tmp.Set("pool", opts.Pool)
	tmp.Set("xhr", opts.XHR)
	tmp.Set("orientation", opts.Orientation)
	tmp.Set("stereo", opts.Stereo)
	tmp.Set("pos", opts.Pos)

	if opts.Sprites != nil {
		sprites := make(map[string]any)
		for name, sprite := range opts.Sprites {
			sprites[name] = []any{sprite.Offset.Milliseconds(), sprite.Duration.Milliseconds(), sprite.Loop}
		}
		tmp.Set("sprite", sprites)
	}

	setCallback(tmp, "onload", opts.OnLoad)
	setCallback(tmp, "onloaderror", opts.OnLoadError)
	setCallback(tmp, "onplayerror", opts.OnPlayError)
	setCallback(tmp, "onplay", opts.OnPlay)
	setCallback(tmp, "onend", opts.OnEnd)
	setCallback(tmp, "onpause", opts.OnPause)
	setCallback(tmp, "onstop", opts.OnStop)
	setCallback(tmp, "onmute", opts.OnMute)
	setCallback(tmp, "onvolume", opts.OnVolume)
	setCallback(tmp, "onrate", opts.OnRate)
	setCallback(tmp, "onseek", opts.OnSeek)
	setCallback(tmp, "onfade", opts.OnFade)
	setCallback(tmp, "onunlock", opts.OnUnlock)
	setCallback(tmp, "onstereo", opts.OnStereo)
	setCallback(tmp, "onpos", opts.OnPos)
	setCallback(tmp, "onorientation", opts.OnOrientation)

	return Howl{
		soundGroup{howl.New(tmp)},
	}
}

type Sprite struct {
	// The offset in milliseconds.
	Offset time.Duration `json:"offset,omitempty"`
	// The duration in milliseconds.
	Duration time.Duration `json:"duration,omitempty"`
	// Set to true to automatically loop the sprite forever.
	Loop bool `json:"loop,omitempty"`
}

type HowlOptions struct {
	// The sources to the track(s) to be loaded for the sound (URLs or base64 data
	// URIs). These should be in order of preference, howler.js will automatically
	// load the first one that is compatible with the current browser. If your files
	// have no extensions, you will need to explicitly specify the extension using
	// the format property.
	Source []OptionalString `json:"src,omitempty"`

	// The volume of the specific track, from 0.0 to 1.0.
	Volume OptionalFloat `json:"volume,omitempty"` // default=Howler's global volume

	// Set to true to force HTML5 Audio. This should be used for large audio files so
	// that you don't have to wait for the full file to be downloaded and decoded
	// before playing.
	HTML5 OptionalBool `json:"html5,omitempty"`

	// Set to true to automatically loop the sound forever.
	Loop OptionalBool `json:"loop,omitempty"`

	// Automatically begin downloading the audio file when the Howl is defined. If
	// using HTML5 Audio, you can set this to 'metadata' to only preload the file's
	// metadata (to get its duration without download the entire file, for example).
	Preload OptionalBool `json:"preload,omitempty"` // default=true

	// Set to true to automatically start playback when sound is loaded.
	Autoplay OptionalBool `json:"autoplay,omitempty"`

	// Set to true to load the audio muted.
	Mute OptionalBool `json:"mute,omitempty"`

	// Define a sound sprite for the sound.
	Sprites map[string]Sprite `json:"sprites,omitempty"`

	// The rate of playback. 0.5 to 4.0, with 1.0 being normal speed.
	Rate OptionalFloat `json:"rate,omitempty"` // default=1.0

	// The size of the inactive sounds pool. Once sounds are stopped or finish
	// playing, they are marked as ended and ready for cleanup. We keep a pool of
	// these to recycle for improved performance. Generally this doesn't need to be
	// changed. It is important to keep in mind that when a sound is paused, it won't
	// be removed from the pool and will still be considered active so that it can be
	// resumed later.
	Pool OptionalInt `json:"pool,omitempty"` // default=5

	// Howler.js automatically detects your file format from the extension, but you
	// may also specify a format in situations where extraction won't work (such as
	// with a SoundCloud stream).
	Format []OptionalString `json:"format,omitempty"`

	// When using Web Audio, howler.js uses an XHR request to load the audio files.
	// If you need to send custom headers, set the HTTP method or enable
	// withCredentials (see reference), include them with this parameter. Each is
	// optional (method defaults to GET, headers default to null and withCredentials
	// defaults to false).
	XHR js.Value `json:"-"`

	// Sets the stereo panning value of the audio source for this sound or group.
	// This makes it easy to setup left/right panning with a value of -1.0 being far
	// left and a value of 1.0 being far right.
	Stereo OptionalBool `json:"stereo,omitempty"`

	// Sets the 3D spatial position of the audio source for this sound or group
	// relative to the global listener.
	Pos []OptionalFloat `json:"pos,omitempty"`

	// Sets the direction the audio source is pointing in the 3D cartesian coordinate
	// space. Depending on how directional the sound is, based on the cone
	// attributes, a sound pointing away from the listener can be quiet or silent.
	Orientation []OptionalFloat `json:"orientation,omitempty"`

	// Sets the panner node's attributes for a sound or group of sounds. See the
	// pannerAttr method for all available options.
	PannerOptions PannerOptions `json:"panner_options"`

	// Fires when the sound is loaded.
	OnLoad CallbackFunc `json:"-"`
	// Fires when the sound is unable to load.
	OnLoadError CallbackErrorFunc `json:"-"`
	// Fires when the sound is unable to play.
	OnPlayError CallbackErrorFunc `json:"-"`
	// Fires when the sound begins playing.
	OnPlay CallbackFunc `json:"-"`
	// Fires when the sound finishes playing (if it is looping, it'll fire at the end
	// of each loop).
	OnEnd CallbackFunc `json:"-"`
	// Fires when the sound has been paused.
	OnPause CallbackFunc `json:"-"`
	// Fires when the sound has been stopped.
	OnStop CallbackFunc `json:"-"`
	// Fires when the sound has been muted/unmuted.
	OnMute CallbackFunc `json:"-"`
	// Fires when the sound's volume has changed.
	OnVolume CallbackFunc `json:"-"`
	//Fires when the sound's playback rate has changed.
	OnRate CallbackFunc `json:"-"`
	//Fires when the sound has been seeked.
	OnSeek CallbackFunc `json:"-"`
	// Fires when the current sound finishes fading in/out.
	OnFade CallbackFunc `json:"-"`
	// Fires when audio has been automatically unlocked through a touch/click event.
	OnUnlock CallbackFunc `json:"-"`
	// Fires when the current sound has the stereo panning changed.
	OnStereo CallbackFunc `json:"-"`
	// Fires when the current sound has the listener position changed.
	OnPos CallbackFunc `json:"-"`
	// Fires when the current sound has the direction of the listener changed.
	OnOrientation CallbackFunc `json:"-"`
}

type Howl struct {
	soundGroup
}

// Load is called by default, but if you set preload to false, you must call load
// before you can play any sounds.
func (h Howl) Load() {
	h.value.Call("load")
}

// PlaySprite will play the previously played sound (for example, after pausing
// it). However, if an id of a sound that has been drained from the pool is
// passed, nothing will play.
func (h Howl) PlaySprite(name string) Sound {
	if result := h.value.Call("play", name); result.Truthy() {
		return soundSpecific{
			id:    result.Int(),
			value: h.soundGroup.value,
		}
	}
	return soundSpecific{id: -1}
}

// Unload and destroy a Howl object. This will immediately stop all sounds
// attached to this sound and remove it from the cache.
func (h Howl) Unload() {
	h.value.Call("unload")
}
