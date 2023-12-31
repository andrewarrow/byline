package browser

func (v *Vim) Focus() {
	if v.Y == 0 {
		return
	}
	spaces := getSpaces(v.getLine())
	count := len(spaces)
	first := 0
	last := 0
	offsetY := v.FocusStart + v.Y + v.Offset
	for i, line := range v.SavedLines {
		if i < offsetY {
			continue
		}
		if first == 0 {
			first = i
		}
		s := getSpaces(line)
		if len(s) <= count && i > offsetY {
			last = i
			break
		}
	}
	if last == 0 {
		last = len(v.SavedLines)
	}
	v.FocusY = v.Y + v.Offset
	v.FocusStart = first
	v.FocusEnd = last
	v.FocusLevel = count
	v.X = 0
	v.Y = 0
	v.Offset = 0
}

func (v *Vim) Refocus() {
	up := v.searchUp()

	v.Y = up
	v.FocusStart = 0
	v.FocusEnd = 0
	v.FocusY = 0
	v.FocusLevel = 0
	if v.Y == 0 {
		return
	}
	v.Focus()
}

func (v *Vim) searchUp() int {
	correct := 0
	for i := v.FocusStart - 1; i >= 0; i-- {
		line := v.SavedLines[i]
		s := len(getSpaces(line))
		if s < v.FocusLevel {
			correct = i
			break
		}
	}
	return correct
}
