package browser

func (v *Vim) AddNewLineBelow() {
	v.AddOneNewLine()
}

func (v *Vim) AddNewLineAbove() {
	saveBool := vim.hasDirectChildren()
	vim.AddOneNewLineAbove()
	if saveBool {
		vim.Y--
		vim.MoveChildrenRight()
	}
}

func (v *Vim) CreateOneLineOp(dir string, amount int) {
	op := NewOperation("add_lines")
	size := len(getSpaces(vim.getLine())) - 2
	op.Data = []string{sp(size) + "div "}
	op.InsertY = vim.Y + vim.Offset + vim.FocusStart
	vim.RunOp(op)

	vim.Y++
	vim.X = size - vim.FocusLevel + 3
	if v.FocusStart > 0 {
		v.FocusEnd++
	}
	vim.InsertMode = true
}
