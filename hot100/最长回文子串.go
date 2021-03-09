package main

func longestPalindrome(s string) string {
	l := len(s)
	start, end := 0, 0
	for i := 0; i < l; i++ {
		left1, right1 := findPalindrome(i, i, s)
		left2, right2 := findPalindrome(i, i+1, s)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func findPalindrome(l, r int, s string) (int, int) {
	for ; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
	}
	return l + 1, r - 1
}
