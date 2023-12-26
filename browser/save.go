package browser

import "github.com/andrewarrow/feedback/wasm"

func saveLines(s string) {
	payload := map[string]any{}
	payload["lines"] = s
	wasm.DoPost("/space/save", payload)
}
