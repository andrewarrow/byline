package browser

import (
	"byline/common"
	"fmt"
	"strings"

	"github.com/andrewarrow/feedback/markup"
	"github.com/brianvoe/gofakeit"
)

func (v *Vim) BottomKeyPress(k string) {

	if k == "Enter" {
		v.BottomCommand(v.BottomText[1:])
	} else if k == "Backspace" {
		v.BottomText = v.BottomText[0 : len(v.BottomText)-1]
	} else {
		v.BottomText += k
	}
	vim.Bottom.Set("innerHTML", vim.BottomText)
}

func sp(width int) string {
	return fmt.Sprintf("%-*s", width, " ")
}
func (v *Vim) AddImage(imgSize string) {
	op := NewOperation("add_lines")
	size := len(getSpaces(v.getLine())) + 2
	w := 90
	h := 60
	if imgSize == "md" {
		w *= 2
		h *= 2
	}
	op.Data = []string{
		fmt.Sprintf("%simg src=http://placekitten.com/%d/%d rounded-full",
			sp(size), w, h),
	}
	op.InsertY = vim.Y + vim.Offset + vim.FocusStart
	vim.RunOp(op)
}

func (v *Vim) BottomCommand(text string) {
	m := map[string]any{}
	if strings.HasPrefix(text, "w") {
		h := markup.ToHTMLFromLines(m, vim.SavedLines)
		v.Preview.Set("innerHTML", h)
		lines := strings.Join(vim.SavedLines, "\n")
		if len(text) > 1 {
			t := strings.Split(text, " ")
			filename := t[len(t)-1]
			Global.LocalStorage.SetItem(filename, lines)
		}
		Global.LocalStorage.SetItem("byline", lines)
	} else if strings.HasPrefix(text, "color ") {
		t := strings.Split(text, " ")
		color := t[len(t)-1]
		match := markup.ClosestColor(color)
		v.SavedLines[v.Y+v.FocusStart+v.Offset] = getSpaces(v.getLine()) + " " + match
		leaveInsertMode()
	} else if strings.HasPrefix(text, "o") {
		if len(text) > 1 {
			t := strings.Split(text, " ")
			filename := t[len(t)-1]
			lines := Global.LocalStorage.GetItem(filename)
			vim.SavedLines = strings.Split(lines, "\n")
		}
		leaveInsertMode()
	} else if text == "hacker" {
		text := gofakeit.HackerPhrase()
		v.SavedLines[v.Y+v.FocusStart+v.Offset] = getSpaces(v.getLine()) + text
		leaveInsertMode()
	} else if text == "lock" {
		m := v.getTokenMap()
		delete(m, "bg-r")
		color := markup.RandomColor()
		buffer := []string{}
		for k, _ := range m {
			buffer = append(buffer, k)
		}
		s := strings.Join(buffer, " ") + " " + color
		v.SavedLines[v.Y+v.FocusStart+v.Offset] = getSpaces(v.getLine()) + s
	} else if text == "grow" {
		v.GrowMode = true
	} else if text == "img sm" {
		v.AddImage("sm")
	} else if text == "img md" {
		v.AddImage("md")
	} else if text == "top" {
		op := NewOperation("add_lines")
		op.Data = strings.Split(common.Top, "\n")
		op.InsertY = vim.Y + vim.Offset + vim.FocusStart
		vim.RunOp(op)
	} else if text == "3" {
		op := NewOperation("add_lines")
		size := len(getSpaces(v.getLine())) + 2
		op.Data = []string{
			fmt.Sprintf("%sdiv flex w-full items-center", sp(size)),
			fmt.Sprintf("%sdiv bg-r", sp(size+2)),
			fmt.Sprintf("%sone", sp(size+4)),
			fmt.Sprintf("%sdiv bg-r w-full text-center", sp(size+2)),
			fmt.Sprintf("%stwo", sp(size+4)),
			fmt.Sprintf("%sdiv bg-r", sp(size+2)),
			fmt.Sprintf("%sthree", sp(size+4)),
		}
		op.InsertY = vim.Y + vim.Offset
		vim.RunOp(op)
	} else if text == "debug" {
		vim.DebugMode = true
	} else if text == "new" {
		vim.SavedLines = []string{"div min-h-full flex flex-col bg-gray-300 text-gray-600", "  div p-3", "    blank", "  div p-3", "    blank"}
		h := markup.ToHTMLFromLines(m, vim.SavedLines)
		vim.Preview.Set("innerHTML", h)
	}
	vim.BottomMode = false
	v.BottomText = "&nbsp;"
	v.Render()
}
