package main

func hammingDistance(x int, y int) int {
	temp := x ^ y

	var res int
	for temp != 0 {
		if temp & 1 == 1{
			res ++
		}
		temp >>= 1
	}
	return res

}
