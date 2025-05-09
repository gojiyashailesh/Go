package syntax

func CountClosures() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
