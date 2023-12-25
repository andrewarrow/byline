package browser

import (
	"fmt"
	"strings"
	"syscall/js"

	"github.com/andrewarrow/feedback/markup"
	"github.com/andrewarrow/feedback/wasm"
)

type Space struct {
	CurrentLine int
	MaxLines    int
	Lines       []string
	Left        *wasm.Wrapper
	Right       *wasm.Wrapper
	Markup      string
	TypeStart   bool
	Buffer      []string
	AttrIndex   int
	Detail      bool
	Menu        *Menu
	ChangeMenu  *ChangeMenu
}

var space = Space{}

func RegisterSpaceEvents() {
	Document.Document.Call("addEventListener", "keydown", js.FuncOf(keyPress))
	space.Markup = `div bg-gray-900 w-full min-h-screen
  div flex
    div w-1/4
      space
    div text-center m-3 p-3 w-full bg-blue-300 rounded-full text-4xl text-black font-bold
      Welcome To byLine!
    div w-1/4
      space`
	space.Lines = strings.Split(space.Markup, "\n")
	space.MaxLines = len(space.Lines)
	space.Left = Document.ByIdWrap("left")
	space.Right = Document.ByIdWrap("right")
	space.Buffer = []string{}
	space.Render()
}

func keyPress(this js.Value, p []js.Value) any {
	k := p[0].Get("key").String()
	//fmt.Println(k)
	if k == "Meta" || k == "Shift" || k == "Control" {
		return nil
	}
	space.Lines = strings.Split(space.Markup, "\n")
	space.MaxLines = len(space.Lines)

	if space.ChangeMenu != nil {
		if k == "Enter" {
			Document.ByIdWrap("detail").Hide()
			space.Detail = false
			space.ChangeMenu = nil
		} else if k == "Escape" {
			Document.ByIdWrap("detail").Hide()
			space.ChangeMenu = nil
			space.Detail = false
		} else if k == "ArrowUp" {
			space.ChangeMenu.Selected--
			Document.RenderToId("menu", space.ChangeMenu.Template(), space.ChangeMenu)
		} else if k == "ArrowLeft" {
			space.ChangeMenu.Value -= 100
			if space.ChangeMenu.Value < 100 {
				space.ChangeMenu.Value = 100
			}
			Document.RenderToId("menu", space.ChangeMenu.Template(), space.ChangeMenu)
		} else if k == "ArrowRight" {
			space.ChangeMenu.Value += 100
			if space.ChangeMenu.Value > 900 {
				space.ChangeMenu.Value = 900
			}
			Document.RenderToId("menu", space.ChangeMenu.Template(), space.ChangeMenu)
		} else if k == "ArrowDown" {
			space.ChangeMenu.Selected++
			Document.RenderToId("menu", space.ChangeMenu.Template(), space.ChangeMenu)
		}
		color := markup.Colors[space.ChangeMenu.Selected]
		s := fmt.Sprintf("bg-%s-%d", color, space.ChangeMenu.Value)
		space.Replace("bg-", s)
		space.Render()
		return nil
	}

	if space.Detail {
		if k == "Enter" {
			Document.ByIdWrap("detail").Hide()
			space.Detail = false
			space.Add(space.Menu.Value())
			space.Render()
		} else if k == "Escape" {
			Document.ByIdWrap("detail").Hide()
			space.Detail = false
		} else if k == "Backspace" {
			space.Menu.Backspace()
			Document.RenderToId("menu", "menu", space.Menu)
		} else if k == "ArrowUp" {
			space.Menu.Selected--
			Document.RenderToId("menu", "menu", space.Menu)
		} else if k == "ArrowDown" {
			space.Menu.Selected++
			Document.RenderToId("menu", "menu", space.Menu)
		} else {
			space.Menu.Filter(k)
			Document.RenderToId("menu", "menu", space.Menu)
		}
		return nil
	}

	if space.TypeStart {
		if k == "Enter" {
			text := strings.Join(space.Buffer, "")
			space.SetText(text)
			space.Buffer = []string{}
			space.TypeStart = false
			return nil
		}
		space.Buffer = append(space.Buffer, k)
		return nil
	}
	if k == "ArrowUp" && space.CurrentLine > 0 {
		space.CurrentLine--
		space.AttrIndex = 0
	} else if k == "ArrowDown" && space.CurrentLine < space.MaxLines-1 {
		space.CurrentLine++
		space.AttrIndex = 0
	} else if k == "x" {
		space.RemoveAttr()
	} else if k == "a" {
		space.AttrIndex++
	} else if k == "A" {
		space.AttrIndex--
	} else if k == "c" {
		space.ChangeMenu = NewChangeMenu()
		Document.RenderToId("menu", space.ChangeMenu.Template(), space.ChangeMenu)
		Document.ByIdWrap("detail").Show()
		space.Detail = true
	} else if k == " " {
		space.Menu = NewMenu("")
		Document.RenderToId("menu", "menu", space.Menu)
		Document.ByIdWrap("detail").Show()
		space.Detail = true
	}

	space.Render()
	for i := 0; i < space.MaxLines; i++ {
		w := Document.ByIdWrap(fmt.Sprintf("line%d", i+1))
		w.RemoveClass("bg-white")
		w.RemoveClass("text-black")
	}
	w := Document.ByIdWrap(fmt.Sprintf("line%d", space.CurrentLine+1))
	w.AddClass("bg-white")
	w.AddClass("text-black")

	return nil
}

/*
	} else if k == "f" {
		space.SetFlex()
	} else if k == "c" {
		space.Child()
	} else if k == "d" {
		space.Duplicate()
	} else if k == "r" {
		space.Color()
	} else if k == "w" {
		space.Width(1)
	} else if k == "W" {
		space.Width(-1)
	} else if k == "p" {
		space.Padding(1)
	} else if k == "P" {
		space.Padding(-1)
	} else if k == "x" {
		space.RemoveAttr()
	} else if k == "X" {
		space.RemoveNode()
	} else if k == ":" {
		space.TypeStart = true
	} else if k == "." {
		space.Parentize()
*/
