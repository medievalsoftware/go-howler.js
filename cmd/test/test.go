//go:build js && wasm

package main

import (
	"github.com/medievalsoftware/go-howler.js"
)

func main() {
	done := make(chan bool)

	howler.New(howler.HowlOptions{
		Source:   []any{"cheeky-buggers.mp3"},
		Volume:   0.5,
		Autoplay: true,
		OnLoad: func() {
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

	<-done

	println("done")
}
