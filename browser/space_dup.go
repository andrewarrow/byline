package browser

import (
	"strings"
)

func (s *Space) Duplicate() {
	spaces := getSpaces(s.Lines[s.CurrentLine])
	maxCount := len(spaces)

	bufferCopy := []string{}
	count := 0
	for i := s.CurrentLine; i < len(s.Lines); i++ {
		line := s.Lines[i]
		spaces := getSpaces(line)
		count = len(spaces)
		if count < maxCount {
			break
		}
		//fmt.Println(count, line, maxCount)
		bufferCopy = append(bufferCopy, line)
	}
	if count == maxCount {
		bufferCopy = bufferCopy[0 : len(bufferCopy)-1]
	}

	buffer := []string{}
	fired := false
	for i, line := range s.Lines {
		spaces := getSpaces(line)
		count := len(spaces)
		if count <= maxCount && i >= s.CurrentLine && fired == false {
			fired = true
			buffer = append(buffer, strings.Join(bufferCopy, "\n"))
		}
		buffer = append(buffer, line)
	}
	s.Markup = strings.Join(buffer, "\n")
}
