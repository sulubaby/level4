package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	inWord := false
	printedWord := false

	for i := 0; i < len(input); i++ {
		ch := input[i]
		if ch != ' ' && ch != '\t' {
			if !inWord && printedWord {
				z01.PrintRune(' ')
				z01.PrintRune(' ')
				z01.PrintRune(' ')
			}
			z01.PrintRune(rune(ch))
			inWord = true
			printedWord = true
		} else {
			inWord = false
		}
	}
	z01.PrintRune('\n')
}
