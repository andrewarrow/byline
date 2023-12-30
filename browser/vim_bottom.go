package browser

import (
	"strings"

	"github.com/andrewarrow/feedback/markup"
)

func (v *Vim) BottomKeyPress(k string) {

	if v.BottomTypeMode {
		v.BottomText += k
		vim.Bottom.Set("innerHTML", ":"+vim.BottomText)
	} else if k == "Enter" {
		m := map[string]any{}
		h := markup.ToHTMLFromLines(m, vim.SavedLines)
		vim.Preview.Set("innerHTML", h)
		go saveLines(strings.Join(vim.SavedLines, "\n"))
	}
}
