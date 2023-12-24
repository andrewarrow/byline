package browser

import (
	"regexp"
	"strings"
)

func (s *Space) SetFlex() {
	buffer := []string{}
	for i, line := range s.Lines {
		if i == s.CurrentLine {
			spaces := getSpaces(line)
			tokens := strings.Split(strings.TrimSpace(line), " ")
			m := map[string]bool{}
			for _, token := range tokens[1:] {
				m[token] = true
			}
			m["flex"] = !m["flex"]
			buffer = append(buffer, spaces+tokens[0]+" "+makeClasses(m))
			continue
		}
		buffer = append(buffer, line)
	}
	s.Markup = strings.Join(buffer, "\n")
	s.Render()
}

var re = regexp.MustCompile(`^\s+`)

func getSpaces(s string) string {
	return re.FindString(s)
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
