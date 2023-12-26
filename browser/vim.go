package browser

import (
	"syscall/js"

	"github.com/andrewarrow/feedback/wasm"
)

type Vim struct {
	Lines      []string
	Editor     *wasm.Wrapper
	X          int
	Y          int
	InsertMode bool
	VisualMode bool
	FromY      int
	ToY        int
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
	if k == "Escape" {
		vim.InsertMode = false
		vim.VisualMode = false
	}
	if vim.VisualMode {
		vim.VisualArrows(k)
		vim.Render()
		return nil
	}
	if vim.InsertMode {
		vim.Insert(k)
		vim.Render()
		return nil
	}
	if k == "ArrowUp" {
		vim.Y--
		if vim.Y < 0 {
			vim.Y++
		}
		if vim.X >= len(vim.Lines[vim.Y]) {
			vim.X = len(vim.Lines[vim.Y]) - 1
		}
	} else if k == "ArrowDown" {
		vim.Y++
		if vim.Y >= len(vim.Lines) {
			vim.Y--
		}
		if vim.X >= len(vim.Lines[vim.Y]) {
			vim.X = len(vim.Lines[vim.Y]) - 1
		}
	} else if k == "ArrowRight" {
		vim.X++
		if vim.X >= len(vim.Lines[vim.Y]) {
			vim.X--
		}
	} else if k == "ArrowLeft" {
		vim.X--
		if vim.X < 0 {
			vim.X++
		}
	} else if k == "i" {
		vim.InsertMode = true
	} else if k == "o" {
		vim.Lines = append(vim.Lines, "  ")
		vim.Y++
		vim.X = 1
		vim.InsertMode = true
	} else if k == "x" {
		prefix := vim.Lines[vim.Y][0:vim.X]
		suffix := vim.Lines[vim.Y][vim.X+1:]
		vim.Lines[vim.Y] = prefix + suffix
		vim.X--
		if vim.X < 0 {
			vim.X = 0
		}
	} else if k == "V" {
		vim.VisualMode = true
		vim.FromY = vim.Y
		vim.ToY = vim.Y
	} else if k == " " {
	}

	vim.Render()

	return nil
}
