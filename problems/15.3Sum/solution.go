package problems

import "sort"

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	length := len(nums)
	sort.Ints(nums)

	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			continue
		}

		l, r := i+1, length-1
		for l < r {
			sum := num + nums[l] + nums[r]

			if sum < 0 {
				l++
				continue
			}

			if sum > 0 {
				r--
				continue
			}

			ans = append(ans, []int{num, nums[l], nums[r]})

			for ok := true; ok; ok = nums[l] == nums[l-1] && l < r {
				l++
			}

			for ok := true; ok; ok = nums[r] == nums[r+1] && l < r {
				r--
			}
		}
	}

	return ans
}
