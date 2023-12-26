package browser

func (v *Vim) Undo() {
	if len(v.Stack) == 0 {
		return
	}
	lastOp := v.Stack[len(v.Stack)-1]
	v.Stack = v.Stack[0 : len(v.Stack)-1]
	v.UndoOp(lastOp)
}

func (v *Vim) UndoOp(op *Operation) {
	buffer := []string{}
	if op.Name == "add_lines" {
		for i, line := range v.Lines {
			if i > op.InsertY && i <= op.InsertY+len(op.Data) {
				continue
			}
			buffer = append(buffer, line)
		}
	} else if op.Name == "remove_lines" {
		for i, line := range v.Lines {
			buffer = append(buffer, line)
			if i == op.InsertY {
				buffer = append(buffer, op.Data...)
			}
		}
	}
	v.Lines = buffer
	v.Y = op.InsertY
}
