package browser

import (
	"strings"
)

func (s *Space) Add(t string) {
}

func (s *Space) RemoveAttr() {
	buffer := []string{}
	for i, line := range s.Lines {
		if i == s.CurrentLine {
			spaces := getSpaces(line)
			m, tag := makeClassMap(line)
			tokens := strings.Split(sortedList(m, s.AttrIndex), " ")
			rest := strings.Join(tokens[1:], " ")
			str := spaces + tag + " " + rest
			buffer = append(buffer, str)
			continue
		}
		buffer = append(buffer, line)
	}
	s.Markup = strings.Join(buffer, "\n")
}

func (s *Space) RemoveNode() {
}
