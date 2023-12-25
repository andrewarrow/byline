package browser

type ChangeMenu struct {
}

func NewChangeMenu() *ChangeMenu {
	m := ChangeMenu{}
	return &m
}

func (m *ChangeMenu) Template() string {
	return "color_menu"
}
