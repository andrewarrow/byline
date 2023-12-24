package browser

import (
	"fmt"
	"syscall/js"
)

type Space struct {
	CurrentLine int
	MaxLines    int
}

var space = Space{}

func RegisterSpaceEvents() {
	Document.Document.Call("addEventListener", "keydown", js.FuncOf(keyPress))
	space.MaxLines = 6
}

func keyPress(this js.Value, p []js.Value) any {
	k := p[0].Get("key").String()
	if k == "ArrowUp" && space.CurrentLine > 0 {
		space.CurrentLine--
	} else if k == "ArrowDown" && space.CurrentLine < space.MaxLines-1 {
		space.CurrentLine++
	}

	for i := 0; i < space.MaxLines; i++ {
		w := Document.ByIdWrap(fmt.Sprintf("line%d", i+1))
		w.RemoveClass("bg-white")
		w.RemoveClass("text-black")
	}
	w := Document.ByIdWrap(fmt.Sprintf("line%d", space.CurrentLine+1))
	w.AddClass("bg-white")
	w.AddClass("text-black")

	return nil
}
