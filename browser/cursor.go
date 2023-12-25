package browser

import (
	"fmt"

	"github.com/andrewarrow/feedback/wasm"
)

type Cursor struct {
	Div *wasm.Wrapper
	X   int
	Y   int
}

func NewCursor(d *wasm.Wrapper) *Cursor {
	nc := Cursor{}
	nc.Div = d
	return &nc
}

func (c *Cursor) Clear() {
	left := fmt.Sprintf("left-%d", c.X)
	top := fmt.Sprintf("top-%d", c.Y)
	c.Div.RemoveClass(left)
	c.Div.RemoveClass(top)
}

func (c *Cursor) Render() {
	left := fmt.Sprintf("left-%d", c.X)
	top := fmt.Sprintf("top-%d", c.Y)
	c.Div.AddClass(left)
	c.Div.AddClass(top)
}
