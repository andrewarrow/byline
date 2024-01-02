package browser

import "strings"

type Menu struct {
	Items    []string
	Selected int
	Search   string
}

func NewMenu(s string) *Menu {
	m := Menu{}
	m.Search = s
	m.Items = []string{}
	m.FillItems()
	m.Selected = 0
	return &m
}

func (m *Menu) Value() string {
	if len(m.Items) == 0 {
		return ""
	}
	return m.Items[m.Selected]
}

func (m *Menu) Filter(s string) {
	m.Search += s
	m.FillItems()
}
func (m *Menu) Backspace() {
	if m.Search == "" {
		return
	}
	m.Search = m.Search[0 : len(m.Search)-1]
	m.FillItems()
}
func (m *Menu) FillItems() {
	m.Items = []string{}
	for _, item := range allItems {
		if strings.HasPrefix(item, m.Search) {
			m.Items = append(m.Items, item)
			if len(m.Items) > 9 {
				break
			}
		}
	}
}

var allItems = []string{
	"line-through",
	"text-red-600",
	"mb-auto",
	"bg-white",
	"text-sm",
	"rounded-lg",
	"rounded-sm",
	"rounded-md",
	"text-2xl",
	"text-3xl",
	"text-4xl",
	"text-5xl",
	"text-6xl",
	"bg-r", "text-center", "text-left", "text-right", "text-white", "text-black",
	"cursor-pointer", "flex", "flex-grow", "flex-col", "items-start", "items-end",
	"items-center", "justify-left", "justify-center", "justify-right", "center-text",
	"rounded", "rounded-full", "p-3", "m-3", "w-full",
	"w-1/2", "w-64", "h-full", "h-64", "ml-auto", "space-y-3", "space-x-3",
	"ml-3", "mr-3", "mt-3", "mb-3",
	"pl-3", "pr-3", "pt-3", "pb-3",
	"border", "border-2", "border-black", "border-r-2", "border-b-2",
	"whitespace-nowrap",
}
