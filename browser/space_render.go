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

		spaces := getSpaces(line)
		m, tag := makeClassMap(line)
		str := fmt.Sprintf(strings.ReplaceAll(spaces, " ", "&nbsp;"))
		p := Document.NewTag("p", str+" "+tag+" "+sortedList(m))
		p.Set("id", fmt.Sprintf("line%d", i+1))
		p.AddClass("whitespace-nowrap")
		s.Left.AppendChild(p.JValue)
	}

	m := map[string]any{}
	h := markup.ToHTMLFromLines(m, lines)
	s.Right.Set("innerHTML", h)
}

func sortedList(m map[string]bool) string {
	buffer := []string{}

	for k, _ := range m {
		buffer = append(buffer, k)
	}
	sort.Strings(buffer)

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
