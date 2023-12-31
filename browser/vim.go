package browser

import (
	"byline/common"
	"fmt"
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
	Bottom      *wasm.Wrapper
	MenuDiv     *wasm.Wrapper
	Menu        *Menu
	DebugLine   string
	X           int
	Y           int
	Location    int
	FocusY      int
	FocusStart  int
	FocusEnd    int
	FocusLevel  int
	InsertMode  bool
	VisualMode  bool
	DeleteMode  bool
	ReplaceMode bool
	BottomMode  bool
	DebugMode   bool
	GrowMode    bool
	BottomText  string
	StartY      int
	EndY        int
	Yanked      []string
	Stack       []*Operation
	Offset      int
}

var MAX_LINES = 20

var vim = Vim{}

func RegisterVimEvents() {
	windowHeight := Global.Window.GetInt("innerHeight")
	//1080 vs 214
	multiple := int(float64(windowHeight) / 214.0)
	_ = multiple
	MAX_LINES = 20 //MAX_LINES * multiple
	fmt.Println(windowHeight)
	Document.Document.Call("addEventListener", "paste", js.FuncOf(vimPaste))
	Document.Document.Call("addEventListener", "keydown", js.FuncOf(vimKeyPress))
	lines := Global.LocalStorage.GetItem("byline")
	if lines == "<null>" {
		lines = common.Sample
	}
	vim.SavedLines = strings.Split(lines, "\n")
	vim.OffsetLines = []string{}
	vim.Editor = Document.ByIdWrap("editor")
	vim.Preview = Document.ByIdWrap("preview")
	vim.Debug = Document.ByIdWrap("debug")
	vim.Bottom = Document.ByIdWrap("bottom")
	vim.MenuDiv = Document.ByIdWrap("menu")
	vim.Stack = []*Operation{}
	//vim.DebugMode = true
	vim.Render()
	leaveInsertMode()
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
		vim.ReplaceMode = false
		vim.BottomMode = false
		vim.GrowMode = false
		vim.DebugMode = false
		vim.MenuDiv.Hide()
		vim.Debug.Hide()
		vim.Bottom.Set("innerHTML", "&nbsp;")
		leaveInsertMode()
	}

	if vim.BottomMode {
		vim.BottomKeyPress(k)
		return nil
	}

	if vim.GrowMode {
		vim.GrowTag(k)
		return nil
	}

	if vim.ReplaceMode {
		vim.ReplaceMode = false
		vim.Replace(k)
		vim.Render()
		leaveInsertMode()
		return nil
	}
	if vim.DeleteMode && k == "d" {
		vim.DeleteMode = false
		line := vim.getLine()
		op := NewOperation("remove_lines")
		op.Data = []string{line}
		vim.Yanked = op.Data
		op.InsertY = vim.Y + vim.Offset

		saveBool := vim.hasDirectChildren()
		vim.RunOp(op)
		vim.X = 0
		if saveBool {
			vim.MoveChildrenLeft()
		}
		vim.Render()
		leaveInsertMode()
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
			vim.Offset--
			if vim.Offset < 0 {
				vim.Offset = 0
			}
			vim.Y++
		} else if vim.Y < 0 && vim.FocusY > 0 {
			vim.Refocus()
		}
		vim.X = len(vim.getLine()) - 1 - vim.FocusLevel
	} else if k == "ArrowDown" {
		vim.Y++
		size := vim.FocusEnd - vim.FocusStart
		if size == 0 {
			vim.Location = vim.Offset + vim.Y
			size = len(vim.SavedLines)
		} else {
			vim.Location = vim.Offset + vim.Y
		}
		if vim.Location >= size {
			vim.Y--
		}
		if vim.Y >= MAX_LINES {
			vim.Y--
			vim.Offset++
		}
		//if vim.X+vim.FocusLevel >= len(vim.getLine()) {
		vim.X = len(vim.getLine()) - 1 - vim.FocusLevel
		//}
	} else if k == "ArrowRight" {
		vim.X++
		if vim.X+vim.FocusLevel >= len(vim.getLine()) {
			vim.X--
		}
	} else if k == "ArrowLeft" {
		vim.X--
		if vim.X < 0 {
			vim.X++
		}
	} else if k == "i" {
		vim.InsertMode = true
	} else if k == "r" {
		vim.ReplaceMode = true
	} else if k == "o" {
		saveBool := vim.hasDirectChildren()
		if saveBool {
			fmt.Println("vim.AddOneNewLineAbove", vim.getLine())
			vim.AddOneNewLineAbove()
		} else {
			vim.AddOneNewLine()
		}
		if saveBool {
			vim.Y--
			vim.MoveChildrenRight()
		}
	} else if k == "x" {
		s := vim.getLine()
		vim.X++
		prefix := s[0 : vim.X+vim.FocusLevel]
		suffix := s[vim.X+vim.FocusLevel:]
		vim.SavedLines[vim.Y+vim.FocusStart+vim.Offset] = prefix[0:len(prefix)-1] + suffix
		vim.X--
		if vim.X < 0 {
			vim.X = 0
		}
		leaveInsertMode()
	} else if k == "V" {
		vim.VisualMode = true
		vim.StartY = vim.Y + vim.FocusStart + vim.Offset
		vim.EndY = vim.Y + vim.FocusStart + vim.Offset
	} else if k == ":" {
		vim.BottomMode = true
		vim.BottomText = ":"
		vim.Bottom.Set("innerHTML", ":")
	} else if k == "a" {
		vim.InsertMode = true
		vim.SavedLines[vim.Y+vim.FocusStart+vim.Offset] += " "
		vim.X++
	} else if k == "d" {
		vim.DeleteMode = true
	} else if k == "D" {
		s := vim.getLine()
		prefix := s[0 : vim.X+1+vim.FocusLevel]
		vim.SavedLines[vim.Y+vim.FocusStart+vim.Offset] = prefix[0:len(prefix)-1] + " "
		vim.X = len(prefix) - 1
		leaveInsertMode()
		vim.InsertMode = true
	} else if k == "u" {
		vim.Undo()
	} else if k == "Enter" {
		vim.Focus()
	} else if k == "p" {
		op := NewOperation("add_lines")
		op.Data = vim.Yanked
		op.InsertY = vim.Y + vim.Offset
		vim.RunOp(op)
	}

	vim.Render()

	return nil
}

func leaveInsertMode() {
	lines := strings.Join(vim.SavedLines, "\n")
	//fmt.Println(lines)
	Global.LocalStorage.SetItem("byline", lines)
	h := markup.ToHTMLFromLines(nil, vim.SavedLines)
	vim.Preview.Set("innerHTML", h)
}
