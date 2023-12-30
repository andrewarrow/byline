package browser

import (
	"strings"

	"github.com/andrewarrow/feedback/markup"
)

func (v *Vim) BottomKeyPress(k string) {

	if k == "Enter" {
		m := map[string]any{}
		h := markup.ToHTMLFromLines(m, vim.SavedLines)
		vim.Preview.Set("innerHTML", h)
		go saveLines(strings.Join(vim.SavedLines, "\n"))
		vim.BottomMode = false
		v.BottomText = "&nbsp;"
	} else {
		v.BottomText += k
	}
	vim.Bottom.Set("innerHTML", vim.BottomText)
}
