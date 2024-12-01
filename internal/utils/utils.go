package utils

import "strconv"

func ToIntSlice(lines []string) []int {
	var nums []int
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	return nums
}
