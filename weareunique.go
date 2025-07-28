package piscine

import "strings"

func WeAreUnique(str1, str2 string) int {
	counter := 0
	temp := ""

	if str1 == "" && str2 == "" {
		return -1
	}
	for _, v := range str1 {
		if !strings.Contains(str2, string(v)) && !strings.Contains(temp, string(v)) {
			counter++
			temp += string(v)
		}
	}
	for _, v := range str1 {
		if !strings.Contains(str1, string(v)) && !strings.Contains(temp, string(v)) {
			counter++
			temp += string(v)
		}
	}
	return counter
}
