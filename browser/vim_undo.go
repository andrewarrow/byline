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
	}
	v.Lines = buffer
	v.Y = op.InsertY
}
