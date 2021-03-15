package main

import "fmt"

func permute(nums []int) [][]int {
	var res [][]int
	var mid []int
	n := len(nums)
	backTrack(nums, mid, &res, n)
	return res
}

func backTrack(nums, mid []int, res *[][]int, n int) {
	if len(mid) == n {
		*res = append(*res, mid)
	}
	for k, v := range nums {
		mid1 := append(mid, v)
		temp := make([]int, len(nums))
		copy(temp, nums)
		nums1 := append(temp[:k], temp[k+1:]...)
		backTrack(nums1, mid1, res, n)

	}
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
