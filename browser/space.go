package browser

import (
	"fmt"
	"strings"
	"syscall/js"

	"github.com/andrewarrow/feedback/wasm"
)

type Space struct {
	CurrentLine int
	MaxLines    int
	Left        *wasm.Wrapper
	Right       *wasm.Wrapper
	Markup      string
}

var space = Space{}

func RegisterSpaceEvents() {
	Document.Document.Call("addEventListener", "keydown", js.FuncOf(keyPress))
	space.Markup = `div id=tag1 br-g w-64 h-64
  div id=tag2 br-g w-9 h-9
    div id=tag3 br-g w-9 h-9
  div id=tag4 br-g w-9 h-9
  div id=tag5 br-g w-9 h-9
  div id=tag6 br-g w-9 h-9`
	space.MaxLines = len(strings.Split(space.Markup, "\n"))
	space.Left = Document.ByIdWrap("left")
	space.Right = Document.ByIdWrap("right")
	space.Render()
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