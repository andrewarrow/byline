package browser

func (v *Vim) VisualArrows(k string) {
	if k == "ArrowDown" {
		v.Y++
		v.EndY++
	} else if k == "y" {
		v.VisualMode = false
		start, end := blockOfLines(v.StartY, v.EndY)
		//v.DebugLine = fmt.Sprintf("%d,%d", start, end)
		//v.Render()
		lines := v.SavedLines[start : end+1]
		v.Yanked = append([]string{}, lines...)
		v.X = 0
	} else if k == ">" {
		v.VisualMode = false
		start, end := blockOfLines(v.StartY, v.EndY)
		op := NewOperation("indent_lines")
		op.From = start
		op.To = end
		vim.RunOp(op)
	} else if k == "d" {
		v.VisualMode = false
		start, end := blockOfLines(v.StartY, v.EndY)
		lines := v.SavedLines[start : end+1]
		v.Yanked = append([]string{}, lines...)

		op := NewOperation("remove_lines")
		op.Data = v.Yanked
		op.InsertY = start
		vim.RunOp(op)
		v.Y -= len(lines)
	} else if k == "ArrowUp" {
		v.Y--
		//fmt.Println(v.Y, v.FromY, v.ToY)
		v.EndY--
		//fmt.Println(v.Y, v.FromY, v.ToY)
	}
	if k == "Enter" {
		return
	}
}

func blockOfLines(start, end int) (int, int) {
	if start > end {
		return end, start
	} else if start < end {
		return start, end
	}
	return start, start
}
