package main

func majorityElement(nums []int) int {
	var count,res int
	for _,v := range nums {
		if count == 0 {
			res = v
			count++
		} else if v == res {
			count++
		} else if v != res {
			count--
		}
	}
	return res
}
