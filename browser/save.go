package browser

import (
	"encoding/json"
	"strings"

	"github.com/andrewarrow/feedback/wasm"
)

func saveLines(s string) {
	payload := map[string]any{}
	payload["lines"] = s
	wasm.DoPost("/space/save", payload)
}

func loadLines() []string {
	jsonString := wasm.DoGet("/space/load/index.mu")
	var m map[string]any
	json.Unmarshal([]byte(jsonString), &m)
	lines := m["lines"].(string)
	return strings.Split(lines, "\n")
}
