package browser

type Operation struct {
	Name    string
	Data    []string
	InsertY int
	X       int
	Y       int
}

func NewOperation(name string) *Operation {
	no := Operation{}
	no.Name = name
	return &no
}

func (v *Vim) RunOp(op *Operation) {
	buffer := []string{}
	if op.Name == "add_lines" {
		for i, line := range v.Lines {
			buffer = append(buffer, line)
			if i == op.InsertY {
				buffer = append(buffer, op.Data...)
			}
		}
	} else if op.Name == "remove_lines" {
		for i, line := range v.Lines {
			if i >= op.InsertY && i < op.InsertY+len(op.Data) {
				continue
			}
			buffer = append(buffer, line)
		}
	}
	v.Lines = buffer
	v.Stack = append(v.Stack, op)
}
