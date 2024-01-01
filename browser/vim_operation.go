package browser

import (
	"strings"
)

type Operation struct {
	Name    string
	Data    []string
	InsertY int
	X       int
	Y       int
	From    int
	To      int
}

func NewOperation(name string) *Operation {
	no := Operation{}
	no.Name = name
	return &no
}

func (v *Vim) AddOneNewLine() {
	op := NewOperation("add_lines")
	size := len(getSpaces(vim.getLine())) + 2
	op.Data = []string{sp(size) + "  "}
	op.InsertY = vim.Y + vim.Offset + vim.FocusStart
	vim.RunOp(op)

	vim.Y++
	vim.X = size - vim.FocusLevel
	vim.InsertMode = true
}

func (v *Vim) AddOneNewLineAbove() {
	op := NewOperation("add_line_above")
	size := len(getSpaces(vim.getLine()))
	op.Data = []string{sp(size)}
	op.InsertY = vim.Y + vim.Offset + vim.FocusStart
	vim.RunOp(op)

	vim.Y++
	vim.X = 0
	if v.FocusStart > 0 {
		v.FocusEnd++
	}
	vim.InsertMode = true
}

func (v *Vim) RunOp(op *Operation) {
	buffer := []string{}
	if op.Name == "add_lines" {
		for i, line := range v.SavedLines {
			buffer = append(buffer, line)
			if i == op.InsertY {
				buffer = append(buffer, op.Data...)
			}
		}
		if v.FocusLevel > 0 {
			v.FocusEnd += len(op.Data)
		}
	} else if op.Name == "add_line_above" {
		for i, line := range v.SavedLines {
			if i == op.InsertY {
				buffer = append(buffer, op.Data...)
			}
			buffer = append(buffer, line)
		}
		if v.FocusLevel > 0 {
			v.FocusEnd += len(op.Data)
		}
	} else if op.Name == "indent_lines" {
		for i, line := range v.SavedLines {
			spaces := ""
			if i >= op.From && i <= op.To {
				spaces = "  "
			}
			buffer = append(buffer, spaces+line)
		}
	} else if op.Name == "unindent_lines" {
		for i, line := range v.SavedLines {
			spaces := line
			if i >= op.From && i <= op.To {
				spaces = spaces[2:]
			}
			buffer = append(buffer, spaces)
		}
	} else if op.Name == "remove_lines" {
		for i, line := range v.SavedLines {
			//fmt.Println("remove_lines", i, op.InsertY)
			if i >= op.InsertY && i < op.InsertY+len(op.Data) {
				continue
			}
			buffer = append(buffer, line)
		}
		v.FocusEnd -= len(op.Data)
	}
	lines := strings.Join(v.SavedLines, "\n")
	v.UndoStack = append(v.UndoStack, lines)
	v.SavedLines = buffer
	//for _, line := range v.SavedLines {
	//fmt.Println("_i", line)
	//}
	leaveInsertMode()
}
