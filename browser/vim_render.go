package browser

import (
	"fmt"
	"strings"
)

func (v *Vim) getLine() string {
	return v.SavedLines[v.Y+v.FocusStart+v.Offset]
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

func (v *Vim) pageLines2() []string {
	buffer := []string{}

	count := 0
	data := v.SavedLines
	if v.FocusStart > 0 {
		data = data[v.FocusStart:v.FocusEnd]
		count = v.FocusLevel
	}

	for _, line := range data {
		fixedLine := line
		if count > 0 {
			fixedLine = line[count:]
		}
		buffer = append(buffer, fixedLine)
		if len(buffer) > MAX_LINES {
			break
		}
	}
	return buffer
}

func (v *Vim) RenderDebug() {
	debug := `<p>Saved Lines: %d</p>
<p>FocusStart: %d</p>
<p>FocusEnd: %d</p>
<p>Y: %d</p>
<p>FocusY: %d</p>
<p>FocusLevel: %d</p>
<p class="font-mono">DebugLine: %s</p>
`
	v.Debug.Set("innerHTML", fmt.Sprintf(debug, len(v.SavedLines),
		v.FocusStart,
		v.FocusEnd,
		v.Y,
		v.FocusY,
		v.FocusLevel,
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
			if v.StartY == v.EndY && i == v.StartY {
				p.AddClass("bg-gray-600")
			} else if v.StartY > v.EndY && (i >= v.EndY && i <= v.StartY) {
				p.AddClass("bg-gray-600")
			} else if v.StartY < v.EndY && (i >= v.StartY && i <= v.EndY) {
				p.AddClass("bg-gray-600")
			}
		}
		v.Editor.AppendChild(p.JValue)
		//p.AddClass("whitespace-nowrap")
		if i != vim.Y {
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
