package browser

import "syscall/js"

func vimPaste(this js.Value, p []js.Value) any {
	p[0].Call("preventDefault")
	//e := wasm.GetItemMap(p[0], 0)
	o := p[0].Get("clipboardData")
	paste := o.Call("getData", "text").String()
	s := vim.getLine()
	vim.SavedLines[vim.Y+vim.Offset] = s[0:len(s)-2] + paste
	vim.Render()
	/*
		for _, char := range paste {
			s := fmt.Sprintf("%c", char)
			params := []js.Value{}
			vimKeyPress(this, params)
		}*/

	return nil
}
