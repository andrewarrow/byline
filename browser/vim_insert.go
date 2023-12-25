package browser

func (v *Vim) Insert(k string) {
	//chars := []rune(v.Lines[v.Y])
	//replacementRune := []rune(k)[0]
	//chars[v.X] = replacementRune
	if k == "ArrowUp" || k == "ArrowDown" || k == "ArrowLeft" || k == "ArrowRight" {
		return
	}
	if k == "Enter" || k == "" || k == "" || k == "" {
		return
	}

	s := v.Lines[v.Y]
	prefix := s[0:v.X]
	suffix := s[v.X:]
	v.X++

	v.Lines[v.Y] = prefix + k + suffix
}
