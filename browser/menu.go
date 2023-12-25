package browser

type Menu struct {
	Items    []string
	Selected int
}

func NewMenu() *Menu {
	m := Menu{}
	m.Items = []string{"wefwef", "wefwefwe"}
	m.Selected = 0
	return &m
}
