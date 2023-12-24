package browser

import (
	"fmt"
	"syscall/js"
)

type Space struct {
}

func RegisterSpaceEvents() {
	Document.Document.Call("addEventListener", "keydown", js.FuncOf(keyPress))
}

func keyPress(this js.Value, p []js.Value) any {
	fmt.Println(p)
	return nil
}
