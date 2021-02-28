package main

func twoSum(nums []int, target int) []int {
	var res []int
	temp := make(map[int]int)
	for k, v := range nums {
		if _, ok := temp[v];ok {
			res = append(res,temp[v],k)
			break
		} else {
			temp[target-v] = k
		}
	}
	return res
}
