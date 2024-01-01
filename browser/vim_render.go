package browser

import (
	"fmt"
	"strings"
)

func (v *Vim) getTokenMap() (string, map[string]int) {
	s := v.getLine()
	tokens := strings.Fields(s)
	m := map[string]int{}
	for i, token := range tokens[1:] {
		m[token] = i
	}
	return tokens[0], m
}
func (v *Vim) getFirstToken() string {
	s := v.getLine()
	tokens := strings.Fields(s)
	if len(tokens) > 0 {
		return tokens[0]
	}
	return ""
}
func (v *Vim) endOfLine() int {
	s := v.getLine()
	return len(s)
}
func (v *Vim) getLine() string {
	return v.SavedLines[v.Y+v.FocusStart+v.Offset]
}
func (v *Vim) getLineBelow() string {
	val := v.Y + v.FocusStart + v.Offset + 1
	if val < len(v.SavedLines) {
		return v.SavedLines[val]
	}
	return ""
}
func (v *Vim) getLineAbove() string {
	return v.SavedLines[v.Y+v.FocusStart+v.Offset-1]
}
func (v *Vim) lineAtSameLevelAsChild() bool {
	line := v.getLine()
	below := v.getLineBelow()
	belowCount := len(getSpaces(below))
	return belowCount == len(getSpaces(line))
}
func (v *Vim) hasDirectChildren() bool {
	line := v.getLine()
	below := v.getLineBelow()
	belowCount := len(getSpaces(below))
	return belowCount > len(getSpaces(line))
}

func windowOfLines(offset int, lines []string) []string {
	buffer := []string{}

	for _, line := range lines[offset:] {
		buffer = append(buffer, line)
		if len(buffer) > MAX_LINES {
			break
		}
	}
	return buffer
}

func (v *Vim) pageLines() []string {
	if v.FocusStart == 0 {
		return windowOfLines(v.Offset, v.SavedLines)
	}

	buffer := []string{}
	count := 0
	data := v.SavedLines
	data = data[v.FocusStart:v.FocusEnd]
	count = v.FocusLevel

	for _, line := range data {
		fixedLine := line
		if count > 0 {
			fixedLine = line[count:]
		}
		buffer = append(buffer, fixedLine)
	}

	return windowOfLines(v.Offset, buffer)

}

func (v *Vim) RenderDebug() {
	if v.DebugMode == false {
		return
	}
	v.Debug.Show()
	debug := `<p>Saved Lines: %d</p>
<p>FocusStart: %d</p>
<p>FocusEnd: %d</p>
<p>Y: %d</p>
<p>X: %d</p>
<p>FocusY: %d</p>
<p>FocusLevel: %d</p>
<p>Offset: %d</p>
<p>Location: %d</p>
<p>VisualStartY: %d</p>
<p>VisualEndY: %d</p>
<p>UndoStack: %d</p>
<p>RedoStack: %d</p>
<p class="font-mono">DebugLine: %s</p>
`
	v.Debug.Set("innerHTML", fmt.Sprintf(debug, len(v.SavedLines),
		v.FocusStart,
		v.FocusEnd,
		v.Y,
		v.X,
		v.FocusY,
		v.FocusLevel,
		v.Offset,
		v.Location,
		v.StartY,
		v.EndY,
		len(v.UndoStack),
		len(v.RedoStack),
		strings.ReplaceAll(v.DebugLine, " ", "_")))
}

func (v *Vim) Render() {
	v.RenderDebug()
	v.Editor.Set("innerHTML", "")
	page := v.pageLines()
	for i, line := range page {
		p := Document.NewTag("p", "")
		p.Set("id", fmt.Sprintf("p%d", i+1))
		if v.VisualMode {
			offsetI := i + v.Offset + v.FocusStart
			if v.StartY == v.EndY && offsetI == v.StartY {
				p.AddClass("bg-gray-600")
			} else if v.StartY > v.EndY && (offsetI >= v.EndY && offsetI <= v.StartY) {
				p.AddClass("bg-gray-600")
			} else if v.StartY < v.EndY && (offsetI >= v.StartY && offsetI <= v.EndY) {
				p.AddClass("bg-gray-600")
			}
		}
		v.Editor.AppendChild(p.JValue)
		//p.AddClass("whitespace-nowrap")
		if i != vim.Y || vim.BottomMode {
			str := fmt.Sprintf(strings.ReplaceAll(line, " ", "&nbsp;"))
			p.Set("innerHTML", str)
			continue
		}

		for j, char := range line {
			s := fmt.Sprintf("%c", char)
			if s == " " {
				s = "&nbsp;"
			}
			span := Document.NewTag("span", s)
			span.Set("id", fmt.Sprintf("s%d-%d", i+1, j+1))
			if v.X == j && v.Y == i {
				span.AddClass("bg-white")
				span.AddClass("text-black")
			}
			p.AppendChild(span.JValue)
		}
	}
}
