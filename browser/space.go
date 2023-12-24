package browser

import (
	"syscall/js"
)

type Space struct {
}

func RegisterSpaceEvents() {
	Document.Document.Call("addEventListener", "keydown", js.FuncOf(keyPress))
}

func keyPress(this js.Value, p []js.Value) any {
	k := p[0].Get("key")
	if k == "ArrowUp" {
	} else if k == "ArrowDown" {
	}
	return nil
}
