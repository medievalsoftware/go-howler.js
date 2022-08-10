//go:build js && wasm

package main

import (
	"fmt"
	"github.com/medievalsoftware/go-howler.js"
	"syscall/js"
	"time"
)

func timestamp(duration time.Duration) string {
	return fmt.Sprintf("%02d:%02d", int(duration.Minutes())%6, int(duration.Seconds())%60)
}

func main() {
	doc := js.Global().Get("document")
	progress := doc.Call("getElementById", "progress")
	seek := doc.Call("getElementById", "seek")
	duration := doc.Call("getElementById", "duration")

	var howl howler.Howl

	howl = howler.New(howler.HowlOptions{
		Source:   []any{"cheeky-buggers.mp3"},
		Volume:   0.5,
		Autoplay: false,

		OnLoad: func() {
			progress.Set("value", 0)
			progress.Set("max", howl.Duration().Milliseconds())
			duration.Set("innerText", timestamp(howl.Duration()))
			println("loaded")
		},

		OnLoadError: func(err error) {
			println(err)
		},

		OnPlayError: func(err error) {
			println(err)
		},

		OnUnlock: func() {
			println("unlock")
		},

		OnMute: func() {
			println("mute")
		},

		OnPlay: func() {
			println("playing")
		},

		OnPause: func() {
			println("paused")
		},

		OnStop: func() {
			println("stopped")
		},

		OnEnd: func() {
			println("done")
		},

		OnRate: func() {
			println("rate change")
		},

		OnFade: func() {
			println("fading")
		},

		OnStereo: func() {
			println("panning")
		},

		OnVolume: func() {
			println("volume change")
		},

		OnSeek: func() {
			println("seeking")
		},

		OnPos: func() {
			println("pos change")
		},

		OnOrientation: func() {
			println("orientation change")
		},
	})

	howler.SetPos(0, 0, 0)

	snd := howl.Play()
	snd.SetLoop(true)

	go func() {
		<-time.After(time.Second * 5)
		snd.Stop()
	}()

	go func() {
		<-time.After(time.Second * 7)
		snd.Play()
	}()

	go func() {
		<-time.After(time.Second * 9)
		snd.Mute()
	}()

	go func() {
		<-time.After(time.Second * 12)
		snd.Unmute()
	}()

	go func() {
		<-time.After(time.Second * 15)
		snd.Pause()
	}()

	go func() {
		<-time.After(time.Second * 17)
		snd.Play()
	}()

	go func() {
		<-time.After(time.Second * 24)
		snd.SetVolume(0.7)
	}()

	go func() {
		<-time.After(time.Second * 25)
		snd.SetSeek(time.Minute)
	}()

	go func() {
		<-time.After(time.Second * 28)
		snd.SetStereo(-1.0)
	}()

	go func() {
		<-time.After(time.Second * 31)
		snd.SetStereo(1.0)
	}()

	go func() {
		<-time.After(time.Second * 33)
		snd.SetStereo(0.0)
	}()

	go func() {
		<-time.After(time.Second * 36)
		snd.SetRate(2.0)
	}()

	go func() {
		<-time.After(time.Second * 41)
		snd.SetRate(0.5)
	}()

	go func() {
		<-time.After(time.Second * 50)
		snd.SetRate(1.0)
	}()

	go func() {
		t := time.Tick(time.Millisecond * 20) // 50fps should be fine

		for {
			<-t
			progress.Set("value", snd.Seek().Milliseconds())
			seek.Set("innerText", timestamp(snd.Seek()))
		}
	}()

	<-make(chan bool)
}
