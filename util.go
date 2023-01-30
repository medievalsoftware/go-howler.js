package howler

import (
	"errors"
	"fmt"
	"reflect"
	"syscall/js"
)

type CallbackFunc func()
type CallbackErrorFunc func(error)

type OptionalInt = any
type OptionalFloat = any
type OptionalBool = any
type OptionalString = any

func setCallback(value js.Value, event string, callback any) {
	var fn js.Func

	if reflect.ValueOf(callback).IsNil() {
		return
	}

	switch callback := callback.(type) {
	case nil:
		fmt.Println(value, event, "nil")
		return
	case CallbackFunc:
		fn = js.FuncOf(func(this js.Value, args []js.Value) any {
			callback()
			return nil
		})
	case CallbackErrorFunc:
		fn = js.FuncOf(func(this js.Value, args []js.Value) any {
			callback(errors.New(args[1].String()))
			return nil
		})
	}

	if fn.Truthy() {
		value.Set(event, fn)
	}
}
