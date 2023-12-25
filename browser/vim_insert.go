package browser

func (v *Vim) Insert(k string) {
	//chars := []rune(v.Lines[v.Y])
	//replacementRune := []rune(k)[0]
	//chars[v.X] = replacementRune

	s := v.Lines[v.Y]
	prefix := s[0 : v.X+1]
	suffix := s[v.X+1:]

	v.Lines[v.Y] = prefix + k + suffix
}
