package browser

import (
	"syscall/js"

	"github.com/andrewarrow/feedback/wasm"
)

type Vim struct {
	Lines  []string
	Editor *wasm.Wrapper
	X      int
	Y      int
}

var vim = Vim{}

func RegisterVimEvents() {
	Document.Document.Call("addEventListener", "keydown", js.FuncOf(vimKeyPress))
	vim.Lines = []string{"tag hi", "  tag there"}
	vim.Editor = Document.ByIdWrap("editor")
	vim.Render()
}

func vimKeyPress(this js.Value, p []js.Value) any {
	k := p[0].Get("key").String()
	//fmt.Println(k)
	if k == "Meta" || k == "Shift" || k == "Control" {
		return nil
	}
	if k == "ArrowUp" {
		vim.Y--
	} else if k == "ArrowDown" {
		vim.Y++
	} else if k == "ArrowRight" {
		vim.X++
	} else if k == "ArrowLeft" {
		vim.X--
	} else if k == "a" {
	} else if k == "k" {
	} else if k == "d" {
	} else if k == "A" {
	} else if k == "c" {
	} else if k == " " {
	}

	vim.Render()

	return nil
}
