package browser

import (
	"strings"
)

func (s *Space) SetFlex() {
	buffer := []string{}
	lines := strings.Split(s.Markup, "\n")
	for i, line := range lines {
		if i == s.CurrentLine {
			tokens := strings.Split(strings.TrimSpace(line), " ")
			m := map[string]bool{}
			for _, token := range tokens[1:] {
				m[token] = true
			}
			m["flex"] = !m["flex"]
			buffer = append(buffer, tokens[0]+" "+makeClasses(m))
			continue
		}
		buffer = append(buffer, line)
	}
	s.Markup = strings.Join(buffer, "\n")
	s.Render()
}

func makeClasses(m map[string]bool) string {
	buffer := []string{}
	for k, v := range m {
		if v == false {
			continue
		}
		buffer = append(buffer, k)
	}
	return strings.Join(buffer, " ")
}
