package problems

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func longestCommonSubsequence(text1 string, text2 string) int {
	length1, length2 := len(text1), len(text2)

	var lookup [1001][1001]int

	for i := 1; i <= length1; i++ {
		for j := 1; j <= length2; j++ {
			if text1[i-1] == text2[j-1] {
				lookup[i][j] = lookup[i-1][j-1] + 1
			} else {
				lookup[i][j] = max(lookup[i-1][j], lookup[i][j-1])
			}
		}
	}

	// for i := 0; i <= length1; i++ {
	// 	for j := 0; j <= length2; j++ {
	// 		fmt.Printf("%d ", lookup[i][j])
	// 	}
	// 	fmt.Println()
	// }

	return lookup[length1][length2]
}
