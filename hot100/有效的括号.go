package main

func isValid(s string) bool {
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var temp []byte
	for _,v := range []byte(s) {
		switch v {
		case '(','[','{':
			temp = append(temp,v)
		case ')', ']','}':
			if len(temp) == 0 {
				return false
			}
			if temp[len(temp)-1] == m[v] {
				temp = temp[:len(temp)-1]
			} else {
				return false
			}
		}
	}
	if len(temp) != 0 {
		return false
	}
	return true
}
