package browser

import "strings"

func makeClassMap(line string) (map[string]bool, string) {
	tokens := strings.Split(strings.TrimSpace(line), " ")
	m := map[string]bool{}
	for _, token := range tokens[1:] {
		m[token] = true
	}
	return m, tokens[0]
}

func (s *Space) Width(i int) {
	buffer := []string{}
	for i, line := range s.Lines {
		if i == s.CurrentLine {
			spaces := getSpaces(line)
			m, tag := makeClassMap(line)
			buffer = append(buffer, spaces+tag+" "+makeClasses(m))
			continue
		}
		buffer = append(buffer, line)
	}
	s.Markup = strings.Join(buffer, "\n")
	s.Render()
}

var sizes = []string{
	"1/12",
	"1/6",
	"1/5",
	"1/4",
	"1/3",
	"2/5",
	"5/12",
	"3/5",
	"1/2",
	"7/12",
	"2/3",
	"3/4",
	"4/5",
	"5/6",
	"7/8",
	"11/12",
}
