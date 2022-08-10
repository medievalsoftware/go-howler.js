package howler

import (
	"errors"
	"syscall/js"
)

type CallbackFunc func()
type CallbackErrorFunc func(error)

type OptionalInt = any
type OptionalFloat = any
type OptionalBool = any
type OptionalString = any

func setCallback(value js.Value, event string, fn any) {
	if fn == nil {
		return
	}

	var jsfunc js.Func

	switch fn := fn.(type) {
	case CallbackFunc:
		jsfunc = js.FuncOf(func(this js.Value, args []js.Value) any {
			fn()
			return nil
		})
	case CallbackErrorFunc:
		jsfunc = js.FuncOf(func(this js.Value, args []js.Value) any {
			fn(errors.New(args[1].String()))
			return nil
		})
	}

	if jsfunc.Truthy() {
		value.Set(event, jsfunc)
	}
}
