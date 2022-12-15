package problems

func reverseString(s []byte) {
	length := len(s)
	half := length / 2
	for i := 0; i < half; i++ {
		left, right := i, length-i-1
		s[left], s[right] = s[right], s[left]
	}
}
