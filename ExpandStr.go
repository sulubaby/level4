package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	str := os.Args[1]
	inWord := false
	spaceNeeded := false

	for _, ch := range str {
		if ch == ' ' || ch == '\t' {
			inWord = false
		} else {
			if spaceNeeded && !inWord {
				// print 3 spaces between words
				z01.PrintRune(' ')
				z01.PrintRune(' ')
				z01.PrintRune(' ')
			}
			z01.PrintRune(ch)
			inWord = true
			spaceNeeded = true
		}
	}
	z01.PrintRune('\n')
}

