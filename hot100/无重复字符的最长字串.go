package main

func lengthOfLongestSubstring(s string) int {
	temp := make(map[byte]struct{})
	res, sign := 0, -1
	for i := 0; i < len(s); i++ {
		if i != 0 {
			delete(temp, s[i-1])
		}
		for j := sign + 1; j < len(s); j++ {
			if _, ok := temp[s[j]]; !ok {
				temp[s[j]] = struct{}{}
			} else {
				sign = j - 1
				break
			}
		}
		res = max(res,len(temp))
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
