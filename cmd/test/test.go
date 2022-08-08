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
	done := make(chan bool)

	doc := js.Global().Get("document")
	progress := doc.Call("getElementById", "progress")
	seek := doc.Call("getElementById", "seek")
	duration := doc.Call("getElementById", "duration")

	var snd howler.Howl

	snd = howler.New(howler.HowlOptions{
		Source:   []any{"cheeky-buggers.mp3"},
		Volume:   0.5,
		Autoplay: true,

		OnLoad: func() {
			progress.Set("value", 0)
			progress.Set("max", snd.Duration().Milliseconds())
			duration.Set("innerText", timestamp(snd.Duration()))
			println("loaded")
		},

		OnLoadError: func(err error) {
			println(err)
		},

		OnPlay: func(id int) {
			println("playing", id)
		},

		OnEnd: func(id int) {
			done <- true
		},
	})

	go func() {
		t := time.Tick(time.Millisecond * 20) // 50fps should be fine

		for {
			<-t
			progress.Set("value", snd.Seek().Milliseconds())
			seek.Set("innerText", timestamp(snd.Seek()))
		}
	}()

	<-done

	println("done")
}
