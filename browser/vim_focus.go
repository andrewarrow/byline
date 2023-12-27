package browser

func (v *Vim) Focus() {
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
	v.FocusY = v.Y
	v.X = 0
	v.Y = 0
	v.OffsetLines = buffer
}

func (v *Vim) Refocus() {
	v.OffsetLines = append([]string{}, v.SavedLines...)
	v.Y = v.FocusY - 1
	v.Focus()
}
