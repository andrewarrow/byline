package browser

import "github.com/andrewarrow/feedback/markup"

func (v *Vim) FullScreenPreview() {
	h := markup.ToHTMLFromLines(nil, vim.SavedLines)
	v.Left.Set("innerHTML", h)
	v.Left.Show()
	v.FullScreenMode = true
}
