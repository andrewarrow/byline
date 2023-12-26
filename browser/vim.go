package browser

import (
	"syscall/js"

	"github.com/andrewarrow/feedback/wasm"
)

type Vim struct {
	Lines        []string
	Editor       *wasm.Wrapper
	X            int
	Y            int
	InsertMode   bool
	VisualMode   bool
	DeleteMode   bool
	FromY        int
	ToY          int
	Deleted      string
	DeletedLines []string
	Stack        []*Operation
}

var vim = Vim{}

func RegisterVimEvents() {
	Document.Document.Call("addEventListener", "keydown", js.FuncOf(vimKeyPress))
	vim.Lines = []string{"001 tag hi", "  002 tag there", "  003 this is more", "  004 and this is even more"}
	vim.Editor = Document.ByIdWrap("editor")
	vim.Stack = []*Operation{}
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
		vim.DeleteMode = false
	}
	if vim.DeleteMode && k == "d" {
		vim.Deleted = vim.Lines[vim.Y]
		vim.DeleteMode = false
		vim.Lines = append(vim.Lines[0:vim.Y], vim.Lines[vim.Y+1:]...)
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
		op := NewOperation("add_lines")
		op.Data = []string{"  "}
		op.InsertY = vim.Y
		vim.RunOp(op)

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
	} else if k == "d" {
		vim.DeleteMode = true
	} else if k == "u" {
		vim.Undo()
	} else if k == "p" {
		yanked := vim.Lines[vim.FromY : vim.ToY+1]
		if vim.Deleted != "" {
			yanked = []string{vim.Deleted}
			vim.Deleted = ""
		}
		if len(vim.DeletedLines) > 0 {
			yanked = vim.DeletedLines
			vim.DeletedLines = []string{}
		}

		prefix := append(vim.Lines[0:vim.Y+1], yanked...)
		vim.Lines = append(prefix, vim.Lines[vim.Y+1:]...)

		vim.FromY = 0
		vim.ToY = 0
	}

	vim.Render()

	return nil
}
