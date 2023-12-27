package browser

func (v *Vim) Focus() {
	// delete from 0 to x.Y-1
	// count spaces at x.Y
	// remove anything below <= spaces
	buffer := []string{}
	spaces := getSpaces(v.getLine())
	count := len(spaces)
	for i, line := range v.OffsetLines {
		if i < v.Y {
			continue
		}
		s := getSpaces(line)
		if len(s) <= count && i > v.Y {
			break
		}
		buffer = append(buffer, line[count:])
	}
	v.OffsetLines = buffer
}
