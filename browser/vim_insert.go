package browser

func (v *Vim) Insert(k string) {
	chars := []rune(v.Lines[v.Y])
	replacementRune := []rune(k)[0]
	chars[v.X] = replacementRune
	v.Lines[v.Y][v.X] = string(chars)
}
