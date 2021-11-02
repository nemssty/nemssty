package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func printStr(str string) {
	arrayStr := []rune(str)
	for i := 0; i < len(arrayStr); i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	args := os.Args[1:]
	if isEven(len(args)) == true {
		printStr(EvenMsg)
		if len(os.Args)-1%2 == 0 {
			fmt.Println("I have an even number of arguments")
		} else {
			printStr(OddMsg)
			fmt.Println("I have an odd number of arguments")
		}
	}
}
