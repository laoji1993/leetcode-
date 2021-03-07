package main

func findDisappearedNumbers(nums []int) []int {
	i := 0
	for i < len(nums) {
		if nums[i] == i+1 {
			i++
			continue
		}
		curIndex := nums[i] - 1
		if nums[i] == nums[curIndex] {
			i++
			continue
		} else {
			nums[i], nums[curIndex] = nums[curIndex], nums[i]
		}
	}

	var res []int
	for k, v := range nums {
		if k != v-1 {
			res = append(res, k+1)
		}
	}
	return res
}
