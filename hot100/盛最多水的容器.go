package main

func maxArea(height []int) int {
	left, right, max := 0, len(height)-1, 0
	for left < right {
		area := min(height[left],height[right]) * (right-left)
		if max < area {
			max = area
		}
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return 	y
}
