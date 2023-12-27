package browser

import (
	"fmt"
	"strings"
)

func (v *Vim) getLine() string {
	return v.SavedLines[v.Offset+v.Y]
}

func (v *Vim) pageLines() []string {
	buffer := []string{}

	count := 0
	data := v.SavedLines
	if v.FocusStart > 0 {
		data = data[v.FocusStart:v.FocusEnd]
		spaces := getSpaces(data[0])
		count = len(spaces)
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
`
	v.Debug.Set("innerHTML", fmt.Sprintf(debug, len(v.SavedLines),
		v.FocusStart,
		v.FocusEnd,
		v.Y,
		v.FocusY))
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
