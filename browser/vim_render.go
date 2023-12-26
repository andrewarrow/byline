package browser

import (
	"fmt"
	"strings"
)

func (v *Vim) getLine() string {
	return v.OffsetLines[v.Offset+v.Y]
}

func (v *Vim) pageLines() []string {
	if len(v.OffsetLines) < MAX_LINES {
		return v.OffsetLines
	}

	end := v.Offset + MAX_LINES
	return v.OffsetLines[v.Offset:end]
}

func (v *Vim) Render() {
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
