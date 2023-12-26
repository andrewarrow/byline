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
			if i == op.InsertY {
				buffer = append(buffer, op.Data...)
			}
			buffer = append(buffer, line)
		}
	}
	v.Lines = buffer
	v.Operations = append(v.Operations, op)
}
