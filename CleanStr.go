package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	arr := []string{}
	temp := ""

	for _, v := range os.Args[1] + " " {
		if v == ' ' || v == '\t' {
			if temp != "" {
				arr = append(arr, temp)
				temp = ""
			}
		} else {
			temp += string(v)
		}
	}
	for i, v := range arr {
		for _, char := range v {
			z01.PrintRune(char)
		}
		if i != len(arr)-1 {
			z01.PrintRune(' ')
		}
	}
}
