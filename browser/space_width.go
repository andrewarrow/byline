package browser

import (
	"strings"
)

func makeClassMap(line string) (map[string]bool, string) {
	tokens := strings.Split(strings.TrimSpace(line), " ")
	m := map[string]bool{}
	for _, token := range tokens[1:] {
		m[token] = true
	}
	return m, tokens[0]
}

func (s *Space) Width(val int) {
	buffer := []string{}
	for i, line := range s.Lines {
		if i == s.CurrentLine {
			spaces := getSpaces(line)
			m, tag := makeClassMap(line)
			w := findWidth(m)
			if w+val >= 0 && w+val < len(sizes) {
				newW := sizes[w+val]
				m["w-"+newW] = true
			}
			buffer = append(buffer, spaces+tag+" "+makeClasses(m))
			continue
		}
		buffer = append(buffer, line)
	}
	s.Markup = strings.Join(buffer, "\n")
	s.Render()
}

func findWidth(m map[string]bool) int {
	for i, s := range sizes {
		key := "w-" + s
		if m[key] == true {
			m[key] = false
			return i
		}
	}

	return 0
}

var spacing = []string{
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"8",
	"10",
	"12",
	"16",
	"20",
	"24",
	"32",
	"40",
	"48",
	"56",
	"64",
	"72",
	"80",
	"96",
}

var sizes = []string{
	"",
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
	"full",
}
