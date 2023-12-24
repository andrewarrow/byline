package browser

import (
	"fmt"
	"syscall/js"
)

type Space struct {
	CurrentLine int
}

var space = Space{}

func RegisterSpaceEvents() {
	space.CurrentLine = 1
	Document.Document.Call("addEventListener", "keydown", js.FuncOf(keyPress))
}

func keyPress(this js.Value, p []js.Value) any {
	k := p[0].Get("key").String()
	if k == "ArrowUp" {
		space.CurrentLine--
	} else if k == "ArrowDown" {
		space.CurrentLine++
	}
	fmt.Println(space)
	return nil
}
