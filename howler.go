//go:build js && wasm

package howler

import (
	"errors"
	"syscall/js"
)

var ErrSpriteNotFound = errors.New("sprite not found")

type ID int

type Sound struct {
	// the backing value
	value js.Value

	// The sources to the track(s) to be loaded for the sound (URLs or base64 data
	// URIs). These should be in order of preference, howler.js will automatically
	// load the first one that is compatible with the current browser. If your files
	// have no extensions, you will need to explicitly specify the extension using
	// the format property.
	Source []string

	// The volume of the specific track, from 0.0 to 1.0.
	Volume float32

	// Set to true to force HTML5 Audio. This should be used for large audio files so
	// that you don't have to wait for the full file to be downloaded and decoded
	// before playing.
	HTML5 bool

	// Set to true to automatically loop the sound forever.
	Loop bool

	// Automatically begin downloading the audio file when the Howl is defined. If
	// using HTML5 Audio, you can set this to 'metadata' to only preload the file's
	// metadata (to get its duration without download the entire file, for example).
	Preload bool

	// Set to true to automatically start playback when sound is loaded.
	Autoplay bool

	// Set to true to load the audio muted.
	Mute bool

	// Define a sound sprite for the sound.
	Sprites map[string]Sprite

	// The rate of playback. 0.5 to 4.0, with 1.0 being normal speed.
	Rate float32

	// The size of the inactive sounds pool. Once sounds are stopped or finish
	// playing, they are marked as ended and ready for cleanup. We keep a pool of
	// these to recycle for improved performance. Generally this doesn't need to be
	// changed. It is important to keep in mind that when a sound is paused, it won't
	// be removed from the pool and will still be considered active so that it can be
	// resumed later.
	Pool int

	// Howler.js automatically detects your file format from the extension, but you
	// may also specify a format in situations where extraction won't work (such as
	// with a SoundCloud stream).
	Formats []string

	// When using Web Audio, howler.js uses an XHR request to load the audio files.
	// If you need to send custom headers, set the HTTP method or enable
	// withCredentials (see reference), include them with this parameter. Each is
	// optional (method defaults to GET, headers default to null and withCredentials
	// defaults to false).
	XHR js.Value

	// Fires when the sound is loaded.
	OnLoad func()

	// Fires when the sound is unable to load. The first parameter is the ID of the
	// sound (if it exists) and the second is the error message/code.
	OnLoadError func(id ID, err js.Value)

	// Fires when the sound is unable to play. The first parameter is the ID of the
	// sound and the second is the error message/code.
	OnPlayError func(id ID, err js.Value)

	// Fires when the sound begins playing. The first parameter is the ID of the
	// sound.
	OnPlay func(id ID)

	// Fires when the sound finishes playing (if it is looping, it'll fire at the end
	// of each loop). The first parameter is the ID of the sound.
	OnEnd func(id ID)

	// Fires when the sound has been paused. The first parameter is the ID of the
	// sound.
	OnPause func(id ID)

	// Fires when the sound has been stopped. The first parameter is the ID of the
	// sound.
	OnStop func(id ID)

	// Fires when the sound has been muted/unmuted. The first parameter is the ID of the sound.
	OnMute func(id ID)

	// Fires when the sound's volume has changed. The first parameter is the ID of the sound.
	OnVolume func(id ID)

	//Fires when the sound's playback rate has changed. The first parameter is the ID of the sound.
	OnRate func(id ID)

	//Fires when the sound has been seeked. The first parameter is the ID of the sound.
	OnSeek func(id ID)

	// Fires when the current sound finishes fading in/out. The first parameter is the ID of the sound.
	OnFade func(id ID)

	// Fires when audio has been automatically unlocked through a touch/click event.
	OnUnlock func()
}

// Play begins playback of a sound. Returns the sound id to be used with other
// methods.
func (s *Sound) Play() ID {
	return ID(s.value.Call("play").Int())
}

// PlayID will play a new sound will play based on the sprite's definition.
func (s *Sound) PlayID(id ID) ID {
	return ID(s.value.Call("play", js.ValueOf(id)).Int())
}

// PlaySprite will play the previously played sound (for example, after pausing
// it). However, if an ID of a sound that has been drained from the pool is
// passed, nothing will play.
func (s *Sound) PlaySprite(name string) (ID, error) {
	if sprite, ok := s.Sprites[name]; ok {
		return ID(s.value.Call("play", sprite.value).Int()), nil
	}
	return -1, ErrSpriteNotFound
}

type Sprite struct {
	// the backing value
	value js.Value

	// The offset in milliseconds.
	Offset int

	// The duration in milliseconds.
	Duration int

	// Set to true to automatically loop the sprite forever.
	Loop bool
}
