package browser

import "fmt"

func (v *Vim) MoveChildrenLeft() {
	line := v.getLine()
	count := len(getSpaces(line))
	start := v.Y + v.Offset + v.FocusStart

	s := 0
	for i := start; i < len(v.SavedLines); i++ {
		line := v.SavedLines[i]
		s = len(getSpaces(line))
		fmt.Println(s, line, count)
		if s < count {
			break
		}
		v.SavedLines[i] = line[2:]
	}
	v.X = 0
}

func (v *Vim) MoveChildrenRight() {
	line := v.getLine()
	theSpaces := getSpaces(line)
	count := len(theSpaces) - 2
	start := v.Y + v.Offset + v.FocusStart
	s := 0
	for i := start + 1; i < len(v.SavedLines); i++ {
		line := v.SavedLines[i]
		s = len(getSpaces(line))
		fmt.Println(s, line, count)
		if s-2 > count {
			break
		}
		v.SavedLines[i] = "  " + line
	}

	v.SavedLines[start] = sp(s-2) + "div "

	if v.lineAtSameLevelAsChild() {
		v.SavedLines[start] = sp(s-4) + "div "
	}

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
