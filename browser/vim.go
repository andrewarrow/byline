package browser

import (
	"strings"
	"syscall/js"

	"github.com/andrewarrow/feedback/markup"
	"github.com/andrewarrow/feedback/wasm"
)

type Vim struct {
	OffsetLines []string
	SavedLines  []string
	Editor      *wasm.Wrapper
	Preview     *wasm.Wrapper
	Debug       *wasm.Wrapper
	X           int
	Y           int
	FocusY      int
	FocusStart  int
	FocusEnd    int
	FocusLevel  int
	InsertMode  bool
	VisualMode  bool
	DeleteMode  bool
	StartY      int
	EndY        int
	Yanked      []string
	Stack       []*Operation
	Offset      int
}

const MAX_LINES = 20

var vim = Vim{}

func RegisterVimEvents() {
	Document.Document.Call("addEventListener", "paste", js.FuncOf(vimPaste))
	Document.Document.Call("addEventListener", "keydown", js.FuncOf(vimKeyPress))
	vim.OffsetLines = []string{"div p-3",
		"  div flex",
		"    div",
		"      left",
		"    div",
		"      right"}
	vim.SavedLines = append([]string{}, vim.OffsetLines...)
	vim.Editor = Document.ByIdWrap("editor")
	vim.Preview = Document.ByIdWrap("preview")
	vim.Debug = Document.ByIdWrap("debug")
	vim.Stack = []*Operation{}
	go func() {
		vim.OffsetLines = loadLines()
		vim.SavedLines = append([]string{}, vim.OffsetLines...)
		vim.Render()
	}()
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
		vim.DeleteMode = false
		op := NewOperation("remove_lines")
		op.Data = []string{string(vim.getLine())}
		vim.Yanked = op.Data
		op.InsertY = vim.Y
		vim.RunOp(op)
		vim.Render()
		return nil
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
		if vim.Y < 0 && vim.FocusY == 0 {
			vim.Y++
		} else if vim.Y < 0 && vim.FocusY > 0 {
			vim.Refocus()
		}
		if vim.X >= len(vim.getLine()) {
			vim.X = len(vim.getLine()) - 1
		}
	} else if k == "ArrowDown" {
		vim.Y++
		if vim.Y >= MAX_LINES {
			vim.Y--
			vim.Offset++
		}
		if vim.X >= len(vim.getLine()) {
			vim.X = len(vim.getLine()) - 1
		}
	} else if k == "ArrowRight" {
		vim.X++
		if vim.X >= len(vim.getLine()) {
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
		/*
			prefix := vim.Lines[vim.Y][0:vim.X]
			suffix := vim.Lines[vim.Y][vim.X+1:]
			vim.Lines[vim.Y] = prefix + suffix
			vim.X--
			if vim.X < 0 {
				vim.X = 0
			}*/
	} else if k == "V" {
		vim.VisualMode = true
		vim.StartY = vim.Y
		vim.EndY = vim.Y
	} else if k == "a" {
		vim.InsertMode = true
		//vim.Lines[vim.Y] += " "
		vim.X++
	} else if k == "Enter" {
		m := map[string]any{}
		h := markup.ToHTMLFromLines(m, vim.SavedLines)
		vim.Preview.Set("innerHTML", h)
		vim.Focus()
		go saveLines(strings.Join(vim.SavedLines, "\n"))
	} else if k == "d" {
		vim.DeleteMode = true
	} else if k == "D" {
		/*
			s := vim.Lines[vim.Y]
			vim.Lines[vim.Y] = s[0:vim.X] + " "
			vim.InsertMode = true
			vim.X = len(vim.Lines[vim.Y]) - 1*/
	} else if k == "u" {
		vim.Undo()
	} else if k == "p" {
		op := NewOperation("add_lines")
		op.Data = vim.Yanked
		op.InsertY = vim.Y
		vim.RunOp(op)
	}

	vim.Render()

	return nil
}
