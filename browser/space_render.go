package browser

import (
	"fmt"
	"sort"
	"strings"

	"github.com/andrewarrow/feedback/markup"
)

func (s *Space) Render() {
	s.Left.Set("innerHTML", "")
	lines := strings.Split(s.Markup, "\n")
	for i, line := range lines {

		str := ""
		if i == s.CurrentLine {
			spaces := getSpaces(line)
			m, tag := makeClassMap(line)
			str = fmt.Sprintf(strings.ReplaceAll(spaces, " ", "&nbsp;"))
			str = str + " " + tag + " " + sortedList(m, s.AttrIndex)
		} else {
			str = fmt.Sprintf(strings.ReplaceAll(line, " ", "&nbsp;"))
		}
		p := Document.NewTag("p", str)
		p.Set("id", fmt.Sprintf("line%d", i+1))
		p.AddClass("whitespace-nowrap")
		s.Left.AppendChild(p.JValue)
	}

	m := map[string]any{}
	h := markup.ToHTMLFromLines(m, lines)
	s.Right.Set("innerHTML", h)
}

func sortedList(m map[string]bool, index int) string {
	buffer := []string{}

	for k, _ := range m {
		buffer = append(buffer, k)
	}
	sort.Strings(buffer)
	if index < len(buffer) {
		buffer = append(buffer[index:], buffer[0:index]...)
	}

	return strings.Join(buffer, " ")
}

func makeClassMap(line string) (map[string]bool, string) {
	tokens := strings.Split(strings.TrimSpace(line), " ")
	m := map[string]bool{}
	for _, token := range tokens[1:] {
		m[token] = true
	}
	return m, tokens[0]
}
