package browser

func (v *Vim) MoveChildrenLeft() {
	line := v.getLine()
	count := len(getSpaces(line))
	start := v.Y + v.Offset + v.FocusStart
	for i := start; i < len(v.SavedLines); i++ {
		line := v.SavedLines[i]
		v.SavedLines[i] = line[2:]
		s := len(getSpaces(line))
		if s < count {
			break
		}
	}

}

func (v *Vim) searchDown(start, level int) int {
	correct := 0
	for i := start; i < len(v.SavedLines); i++ {
		line := v.SavedLines[i]
		s := len(getSpaces(line))
		if s < level {
			correct = i
			break
		}
	}
	return correct
}
