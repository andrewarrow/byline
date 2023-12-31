package browser

type Operation struct {
	Name    string
	Data    []string
	InsertY int
	X       int
	Y       int
	From    int
	To      int
}

func NewOperation(name string) *Operation {
	no := Operation{}
	no.Name = name
	return &no
}

func (v *Vim) AddOneNewLine() {
	op := NewOperation("add_lines")
	size := len(getSpaces(vim.getLine())) + 2
	op.Data = []string{sp(size) + "  "}
	op.InsertY = vim.Y + vim.Offset
	vim.RunOp(op)

	vim.Y++
	vim.X = size
	vim.InsertMode = true
}

func (v *Vim) AddOneNewLineAbove() {
	op := NewOperation("add_line_above")
	size := len(getSpaces(vim.getLineAbove())) + 2
	op.Data = []string{sp(size) + "  "}
	op.InsertY = vim.Y + vim.Offset
	vim.RunOp(op)

	vim.Y++
	vim.X = size
	vim.InsertMode = true
}

func (v *Vim) RunOp(op *Operation) {
	buffer := []string{}
	if op.Name == "add_lines" {
		for i, line := range v.SavedLines {
			buffer = append(buffer, line)
			if i == op.InsertY {
				buffer = append(buffer, op.Data...)
			}
		}
	} else if op.Name == "add_line_above" {
		for i, line := range v.SavedLines {
			if i == op.InsertY+1 {
				buffer = append(buffer, op.Data...)
			}
			buffer = append(buffer, line)
		}
		/*
			for i := len(v.SavedLines) - 1; i >= 0; i-- {
				line := v.SavedLines[i]
				if i == op.InsertY {
					buffer = append(op.Data, buffer...)
				}
				buffer = append([]string{line}, buffer...)
			}*/
	} else if op.Name == "indent_lines" {
		/*
			for i, line := range v.Lines {
				spaces := ""
				if i >= op.From && i <= op.To {
					spaces = "  "
				}
				buffer = append(buffer, spaces+line)
			}*/
	} else if op.Name == "remove_lines" {
		for i, line := range v.SavedLines {
			if i >= op.InsertY && i < op.InsertY+len(op.Data) {
				continue
			}
			buffer = append(buffer, line)
		}
		op.InsertY--
	}
	v.SavedLines = buffer
	v.Stack = append(v.Stack, op)
	leaveInsertMode()
}
