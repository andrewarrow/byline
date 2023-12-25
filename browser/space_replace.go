package browser

import "strings"

func (s *Space) Replace(prefix, value string) {
	buffer := []string{}
	for i, line := range s.Lines {
		if i == s.CurrentLine {
			spaces := getSpaces(line)
			m, tag := makeClassMap(line)
			for k, _ := range m {
				if strings.HasPrefix(k, prefix) {
					delete(m, k)
				}
			}
			m[value] = true
			buffer = append(buffer, spaces+tag+" "+makeClasses(m))
			continue
		}
		buffer = append(buffer, line)
	}
	s.Markup = strings.Join(buffer, "\n")
}
