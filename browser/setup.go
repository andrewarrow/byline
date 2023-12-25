package browser

import "github.com/andrewarrow/feedback/wasm"

var Global *wasm.Global
var Document *wasm.Document

func RegisterEvents() {
	if Global.Start == "welcome.html" {
		//RegisterLoginEvents()
	} else if Global.Start == "space.html" {
		RegisterSpaceEvents()
	} else if Global.Start == "vim.html" {
		RegisterVimEvents()
	}
}
