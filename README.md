# go-howler.js

Golang wrapper for [Howler.js](https://github.com/goldfire/howler.js) using the `syscall/js` package.

## Example

```go
//go:build js && wasm

package main

import (
	"github.com/medievalsoftware/go-howler.js"
)

func main() {
	howler.New(howler.HowlOptions{
		Source:   []any{"cheeky-buggers.mp3"},
		Volume:   0.5,
		Autoplay: true,
		Loop:     true,
		OnEnd: func() {
			println("finished!")
		},
	})
}
```