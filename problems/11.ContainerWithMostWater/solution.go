package problems

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(h []int) int {
	ans := 0
	length := len(h)
	l, r := 0, length-1

	for l < r {
		height := min(h[l], h[r])
		width := r - l
		rect := height * width

		ans = max(ans, rect)

		if h[l] < h[r] {
			l++
		} else {
			r--
		}
	}

	return ans
}
