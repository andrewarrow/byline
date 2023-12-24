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
	space.Markup = `div bg-gray-900 w-full min-h-screen
  div flex
    div w-1/8
      |
    div w-full bg-blue-300 rounded-full
      Welcome To byLine!
    div w-1/8
      |`
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
	} else if k == "f" {
		space.SetFlex()
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
