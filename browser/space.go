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
	Lines       []string
	Left        *wasm.Wrapper
	Right       *wasm.Wrapper
	Markup      string
	TypeStart   bool
	Buffer      []string
	AttrIndex   int
}

var space = Space{}

func RegisterSpaceEvents() {
	Document.Document.Call("addEventListener", "keydown", js.FuncOf(keyPress))
	space.Markup = `div bg-gray-900 w-full min-h-screen
  div flex
    div w-1/4
      space
    div text-center m-3 p-3 w-full bg-blue-300 rounded-full text-4xl text-black font-bold
      Welcome To byLine!
    div w-1/4
      space`
	space.Lines = strings.Split(space.Markup, "\n")
	space.MaxLines = len(space.Lines)
	space.Left = Document.ByIdWrap("left")
	space.Right = Document.ByIdWrap("right")
	space.Buffer = []string{}
	space.Render()
}

func keyPress(this js.Value, p []js.Value) any {
	space.Lines = strings.Split(space.Markup, "\n")
	space.MaxLines = len(space.Lines)

	k := p[0].Get("key").String()
	if space.TypeStart {
		if k == "Enter" {
			text := strings.Join(space.Buffer, "")
			space.SetText(text)
			space.Buffer = []string{}
			space.TypeStart = false
			return nil
		}
		space.Buffer = append(space.Buffer, k)
		return nil
	}
	if k == "ArrowUp" && space.CurrentLine > 0 {
		space.CurrentLine--
	} else if k == "ArrowDown" && space.CurrentLine < space.MaxLines-1 {
		space.CurrentLine++
	} else if k == "f" {
		space.SetFlex()
	} else if k == "a" {
		space.AttrIndex++
	} else if k == "c" {
		space.Child()
	} else if k == "d" {
		space.Duplicate()
	} else if k == "r" {
		space.Color()
	} else if k == "w" {
		space.Width(1)
	} else if k == "W" {
		space.Width(-1)
	} else if k == "p" {
		space.Padding(1)
	} else if k == "P" {
		space.Padding(-1)
	} else if k == ":" {
		space.TypeStart = true
	}

	space.Render()
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
