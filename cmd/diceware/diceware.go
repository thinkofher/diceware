package main

import (
	"flag"
	"fmt"

	"github.com/thinkofher/diceware"
)

const shorthand = " (shorthand)"

var (
	wordsNum   = 5
	wordsUsage = "number of words in single password"

	passwordsNum   = 1
	passwordsUsage = "number of passwords to generate"
)

func main() {
	flag.IntVar(&wordsNum, "words", 5, wordsUsage)
	flag.IntVar(&wordsNum, "w", 5, wordsUsage+shorthand)
	flag.IntVar(&passwordsNum, "passwords", 1, passwordsUsage)
	flag.IntVar(&passwordsNum, "p", 1, passwordsUsage+shorthand)
	flag.Parse()

	for i := 0; i < passwordsNum; i++ {
		fmt.Println(diceware.Generate(wordsNum))
	}
}
