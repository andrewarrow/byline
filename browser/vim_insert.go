package browser

func (v *Vim) Insert(k string) {
	//chars := []rune(v.Lines[v.Y])
	//replacementRune := []rune(k)[0]
	//chars[v.X] = replacementRune
	if k == "ArrowUp" || k == "ArrowDown" || k == "ArrowLeft" || k == "ArrowRight" {
		return
	}
	if k == "Enter" {
		return
	}

	s := v.Lines[v.Y]
	if k == "Backspace" {
		prefix := s[0 : v.X-1]
		suffix := s[v.X:]
		v.Lines[v.Y] = prefix + suffix
		v.X--
		return
	}

	prefix := s[0:v.X]
	suffix := s[v.X:]
	v.X++
	v.Lines[v.Y] = prefix + k + suffix
}
