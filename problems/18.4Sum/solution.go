package problems

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	ans := make([][]int, 0)
	n := len(nums)
	sort.Ints(nums)

	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			k, l := j+1, n-1
			for k < l {
				sum := nums[i] + nums[j] + nums[k] + nums[l]

				if sum < target {
					k++
					continue
				}

				if sum > target {
					l--
					continue
				}

				ans = append(ans, []int{nums[i], nums[j], nums[k], nums[l]})

				for ok := true; ok; ok = nums[k] == nums[k-1] && k < l {
					k++
				}

				for ok := true; ok; ok = nums[l] == nums[l+1] && k < l {
					l--
				}
			}
		}
	}

	return ans
}
