package browser

type Cursor struct {
	X int
	Y int
}

func NewCursor() *Cursor {
	nc := Cursor{}
	return &nc
}

func (c *Cursor) Clear() {
	//left := fmt.Sprintf("left-%d", c.X)
	//top := fmt.Sprintf("top-%d", c.Y)
	//c.Div.RemoveClass(left)
	//left: 1ch;
	//c.Div.RemoveClass(top)
}

func (c *Cursor) Render() {
	//left := fmt.Sprintf("left-%d", c.X)
	//top := fmt.Sprintf("top-%d", c.Y)
	//c.Div.AddClass(left)
	//c.Div.AddClass(top)
	//s.Set("left", fmt.Sprintf("%dch", c.X))
	//s.Set("top", fmt.Sprintf("%dch", c.Y))
}
