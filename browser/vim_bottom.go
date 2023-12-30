package browser

import (
	"strings"

	"github.com/andrewarrow/feedback/markup"
)

func (v *Vim) BottomKeyPress(k string) {

	if k == "Enter" {
		v.BottomCommand(v.BottomText[1:])
		return
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
	} else if text == "new" {
		vim.SavedLines = []string{"div p-3\n  div\n    blank"}
		h := markup.ToHTMLFromLines(m, vim.SavedLines)
		vim.Preview.Set("innerHTML", h)
	}
	vim.BottomMode = false
	v.BottomText = "&nbsp;"
}
