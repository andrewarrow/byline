package browser

import (
	"strings"

	"github.com/andrewarrow/feedback/wasm"
)

func saveLines(s string) {
	payload := map[string]any{}
	payload["lines"] = s
	wasm.DoPost("/space/save", payload)
}

func loadLines() []string {
	lines := wasm.DoGet("/space/load/index.mu")
	return strings.Split(lines, "\n")
}
