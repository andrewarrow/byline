package browser

func (s *Space) Render() {
	s.Left.Set("innerHTML", s.Markup)
}
