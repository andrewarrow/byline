package browser

import (
	"strings"

	"github.com/andrewarrow/feedback/markup"
)

func (v *Vim) BottomKeyPress(k string) {

	if k == "Enter" {
		v.BottomCommand(v.BottomText[1:])
	} else if k == "Backspace" {
		v.BottomText = v.BottomText[0 : len(v.BottomText)-1]
	} else {
		v.BottomText += k
	}
	vim.Bottom.Set("innerHTML", vim.BottomText)
}

func (v *Vim) BottomCommand(text string) {
	m := map[string]any{}
	if text == "w" {
		h := markup.ToHTMLFromLines(m, vim.SavedLines)
		vim.Preview.Set("innerHTML", h)
		go saveLines(strings.Join(vim.SavedLines, "\n"))
	} else if text == "3" {
		op := NewOperation("add_lines")
		op.Data = []string{"    div flex w-full items-center",
			"      div bg-r",
			"        one",
			"      div bg-r w-full text-center",
			"        two",
			"      div bg-r",
			"        three"}
		op.InsertY = vim.Y + vim.Offset
		vim.RunOp(op)
	} else if text == "debug" {
		vim.DebugMode = true
	} else if text == "new" {
		vim.SavedLines = []string{"div min-h-full flex flex-col bg-gray-300 text-gray-600", "  div p-3", "    blank", "  div p-3", "    blank"}
		h := markup.ToHTMLFromLines(m, vim.SavedLines)
		vim.Preview.Set("innerHTML", h)
	}
	vim.BottomMode = false
	v.BottomText = "&nbsp;"
	v.Render()
}
