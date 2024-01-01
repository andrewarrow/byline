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
	if text == "w" {
		h := markup.ToHTMLFromLines(m, vim.SavedLines)
		vim.Preview.Set("innerHTML", h)
		lines := strings.Join(vim.SavedLines, "\n")
		Global.LocalStorage.SetItem("byline", lines)
	} else if text == "hacker" {
		text := gofakeit.HackerPhrase()
		v.SavedLines[v.Y+v.FocusStart+v.Offset] = getSpaces(v.getLine()) + text
		leaveInsertMode()
	} else if text == "grow" {
		vim.GrowMode = true
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
