package browser

import (
	"fmt"
	"regexp"
	"strconv"
)

var widthHightRegex = regexp.MustCompile(`/(\d+)/(\d+)`)

func (v *Vim) GrowTag(k string) {
	line := v.getLine()
	matches := widthHightRegex.FindStringSubmatch(line)
	width, _ := strconv.Atoi(matches[1])
	height, _ := strconv.Atoi(matches[2])

	if k == "ArrowUp" {
		height += 10
	} else if k == "ArrowDown" {
		height -= 10
	} else if k == "ArrowRight" {
		width += 10
	} else if k == "ArrowLeft" {
		width -= 10
	}
	spaces := getSpaces(line)
	newLine := fmt.Sprintf("%simg src=http://placekitten.com/%d/%d rounded-full",
		spaces, width, height)
	v.SavedLines[v.Y+v.FocusStart+v.Offset] = newLine
	v.Render()
	leaveInsertMode()
}
