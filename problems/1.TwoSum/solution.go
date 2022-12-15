package problems

func twoSum(nums []int, target int) []int {
	table := make(map[int]int)

	for k, num := range nums {
		diff := target - num

		index, found := table[diff]

		if found {
			return []int{k, index}
		} else {
			table[num] = k
		}
	}

	return []int{}
}
