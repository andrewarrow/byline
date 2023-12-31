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
	v.X = 0
}

func (v *Vim) MoveChildrenRight() {
	line := v.getLine()
	theSpaces := getSpaces(line)
	count := len(theSpaces)
	start := v.Y + v.Offset + v.FocusStart
	for i := start; i < len(v.SavedLines); i++ {
		line := v.SavedLines[i]
		v.SavedLines[i] = "  " + line
		s := len(getSpaces(line))
		//fmt.Println(s, line, count)
		if s-2 > count {
			break
		}
	}

	v.SavedLines[start] = theSpaces[0:len(theSpaces)-2] + " div "
	v.X = len(v.SavedLines[start]) - 1
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
