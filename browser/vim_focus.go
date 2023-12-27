package browser

func (v *Vim) Focus() {
	if v.FocusY > 0 && v.Y == 0 {
		return
	}
	spaces := getSpaces(v.getLine())
	count := len(spaces)
	first := 0
	last := 0
	for i, line := range v.SavedLines {
		if i < v.Y+v.Offset {
			continue
		}
		if first == 0 {
			first = i
		}
		s := getSpaces(line)
		if len(s) <= count && i > v.Y+v.Offset {
			last = i
			break
		}
	}
	v.FocusY = v.Y + v.Offset
	v.FocusStart = first
	v.FocusEnd = last
	v.X = 0
	v.Y = 0
}

func (v *Vim) Refocus() {
	v.OffsetLines = append([]string{}, v.SavedLines...)
	v.Y = v.FocusY - 1
	v.Focus()
}
